package services

import (
	"fmt"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/models"
	"sv-sfia/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ParticipantService struct {
	db *gorm.DB
}

func newParticipantService(db *gorm.DB) *ParticipantService {
	return &ParticipantService{
		db: db,
	}
}

func (service ParticipantService) checkPassword(plain string, hashed string) *dto.ApiError {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))

	if err != nil {
		return &dto.ApiError{
			Typ:          dto.ErrorBadData,
			Err:          err,
			ErrorMessage: "Email atau Password salah",
		}
	}

	return nil
}

func (service ParticipantService) hashPassword(pass string) (string, *dto.ApiError) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), 10)

	if err != nil {
		zap.L().Error("failed to hash password", zap.Error(err))

		return "", dto.InternalError(err)
	}

	return string(hashed), nil
}

func (service ParticipantService) genereteToken(userId string) (string, *dto.ApiError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"expire":  time.Now().Add(time.Hour * (24 * 30)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(""))
	if err != nil {
		zap.L().Error("error creating token participant", zap.Error(err))

		return "", dto.InternalError(err)
	}

	return tokenString, nil
}

func (service ParticipantService) GetByParticipantId(participantId uuid.UUID) *dto.ApiError {
	participant := models.Participant{}

	err := service.db.Where("uuid = ?", participantId).
		First(&participant).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &dto.ApiError{
				Typ:          dto.ErrorUnauthorized,
				ErrorMessage: "participant not found",
				Err:          fmt.Errorf("participant not found"),
			}
		}

		return dto.InternalError(err)
	}

	return nil
}

func (service ParticipantService) Login(req requests.LoginRequest) (*responses.ParticipantAuthResponse, *dto.ApiError) {
	participant := models.Participant{}
	err := service.db.
		Preload("ParticipantRole").
		Preload("ParticipantDepartment").
		Preload("ParticipantSkills").
		Where("email = ?", req.Email).
		First(&participant).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &dto.ApiError{
				Typ:          dto.ErrorBadData,
				Err:          err,
				ErrorMessage: "Email atau Password salah",
			}
		}

		return nil, dto.InternalError(err)
	}

	passErr := service.checkPassword(req.Password, participant.Password)
	if passErr != nil {
		return nil, passErr
	}

	token, tokenErr := service.genereteToken(participant.Uuid.String())
	if tokenErr != nil {
		return nil, tokenErr
	}

	return &responses.ParticipantAuthResponse{
		Token: token,
		Participant: responses.Participant{
			Name: participant.Name,
			Pn:   participant.Pn,
		},
		IsProfileComplete: service.checkProfileCompleted(participant),
	}, nil
}

func (service ParticipantService) Register(req requests.RegisterRequest) (*responses.AuthResponse, *dto.ApiError) {
	user := models.User{
		Email:      req.Email,
		Name:       req.Name,
		Pn:         req.Pn,
		RoleAccess: models.USER,
	}

	err := service.db.
		Where("email = ?", req.Email).
		First(&models.User{}).
		Error

	if err == nil {
		return nil, &dto.ApiError{
			Typ:          dto.ErrorBadData,
			Err:          fmt.Errorf("participant already registered"),
			ErrorMessage: "Email sudah terdaftar",
		}
	} else {
		if err != gorm.ErrRecordNotFound {

			zap.L().Error("error register participant", zap.Error(err))

			return nil, dto.InternalError(err)
		}
	}

	hashed, errHashed := service.hashPassword(req.Password)
	if errHashed != nil {
		return nil, errHashed
	}

	user.Password = hashed

	err = service.db.Create(&user).Error
	if err != nil {
		return nil, dto.InternalError(err)
	}

	token, errToken := service.genereteToken(user.Uuid.String())
	if errToken != nil {
		return nil, errToken
	}

	return &responses.AuthResponse{
		Token: token,
		User: responses.UserResponse{
			Name:  user.Name,
			Pn:    user.Pn,
			Email: user.Email,
			Role:  user.RoleAccess,
		},
	}, nil
}

func (service ParticipantService) GetPersonalInformation(participantId uuid.UUID) (*responses.ParticipantProfileResponse, *dto.ApiError) {
	participant := models.Participant{}

	err := service.db.Where("uuid = ?", participantId).
		Preload("ParticipantDepartment").
		Preload("ParticipantRole").
		Preload("ParticipantSkills").
		Preload("ParticipantRole.MainRole").
		Preload("ParticipantRole.MainRole.Group").
		Preload("ParticipantRole.SecondaryRole").
		Preload("ParticipantRole.SecondaryRole.Group").
		Preload("ParticipantRole.InterestRole").
		Preload("ParticipantRole.InterestRole.Group").
		First(&participant).Error

	if err != nil {
		zap.L().Error("error querying particpant", zap.Error(err))

		return nil, dto.InternalError(err)
	}

	res := &responses.ParticipantProfileResponse{
		Name:               participant.Name,
		Pn:                 participant.Pn,
		Email:              participant.Email,
		IsProfileCompleted: service.checkProfileCompleted(participant),
	}

	if participant.ParticipantDepartment != nil {
		departmentId := participant.ParticipantDepartment.DepartmentId.String()
		departmentUnitId := participant.ParticipantDepartment.DepartmentUnitId.String()
		departmentTeamId := participant.ParticipantDepartment.DepartmentTeamId.String()

		res.DepartmentId = &departmentId
		res.DepartmentRoleId = &departmentUnitId
		res.DepartmentTeamId = &departmentTeamId
	}

	if participant.ParticipantRole != nil {
		roleInformation := responses.NewParticipantRoleInformation(*participant.ParticipantRole)

		res.RoleInformation = roleInformation
	}

	return res, nil
}

func (service ParticipantService) StorePersonalInformation(userId uuid.UUID, req requests.PersonalInformationRequest) *dto.ApiError {
	departmentId, errParse := utils.ParseUUid(req.DepartmentId)

	if errParse != nil {
		return errParse
	}

	unitId, errParse := utils.ParseUUid(req.DepartmentUnitId)
	if errParse != nil {
		return errParse
	}

	teamId, errParse := utils.ParseUUid(req.DepartmentTeamId)
	if errParse != nil {
		return errParse
	}

	participantDepartment := models.ParticipantDepartment{
		UserId:           userId,
		DepartmentId:     departmentId,
		DepartmentUnitId: unitId,
		DepartmentTeamId: teamId,
	}

	err := service.db.Create(&participantDepartment).Error

	if err != nil {
		zap.L().Error("error store personal information", zap.Error(err))

		return dto.InternalError(err)
	}

	return nil
}

func (service ParticipantService) UpdatePersonalInformation(userId uuid.UUID, req requests.PersonalInformationRequest) *dto.ApiError {
	departmentId, errParse := utils.ParseUUid(req.DepartmentId)

	if errParse != nil {
		return errParse
	}

	unitId, errParse := utils.ParseUUid(req.DepartmentUnitId)
	if errParse != nil {
		return errParse
	}

	teamId, errParse := utils.ParseUUid(req.DepartmentTeamId)
	if errParse != nil {
		return errParse
	}

	participantDepartment := models.ParticipantDepartment{
		UserId:           userId,
		DepartmentId:     departmentId,
		DepartmentUnitId: unitId,
		DepartmentTeamId: teamId,
	}

	err := service.db.Where("user_id = ?", userId).Updates(&participantDepartment).Error

	if err != nil {
		zap.L().Error("error update personal information", zap.Error(err))

		return dto.InternalError(err)
	}

	return nil
}

func (service ParticipantService) StoreParticipantRole(particpantId uuid.UUID, request requests.CreateParticipantRoleRequest) *dto.ApiError {
	participantRole := models.ParticipantRole{ParticipantId: particpantId}

	exist := service.db.Where("participant_id = ?", particpantId).
		First(&participantRole).RowsAffected > 0

	mainRoleId, err := utils.ParseUUid(request.MainRoleId)
	if err != nil {
		return err
	}

	if request.SecondaryRoleId != nil {
		id, err := utils.ParseUUid(*request.SecondaryRoleId)

		if err != nil {
			return err
		}

		participantRole.SecondaryRoleId = &id
	} else {
		participantRole.SecondaryRoleId = nil
	}

	if request.InterestRoleId != nil {
		id, err := utils.ParseUUid(*request.InterestRoleId)

		if err != nil {
			return err
		}

		participantRole.InterestRoleId = &id
	} else {
		participantRole.InterestRoleId = nil
	}

	participantRole.MainRoleId = mainRoleId

	if !exist {
		if err := service.db.Create(&participantRole).Error; err != nil {
			return dto.InternalError(fmt.Errorf("failed to create participant role: %s", err.Error()))
		}

		return nil
	}

	if err := service.db.Model(&models.ParticipantRole{}).
		Where("participant_id = ?", particpantId).
		Updates(&participantRole).Error; err != nil {
		return dto.InternalError(fmt.Errorf("failed to create participant role: %s", err.Error()))
	}

	return nil
}

func (service ParticipantService) StoreParticipantSkill(participantId uuid.UUID, request requests.StoreParticipantSkillRequest) *dto.ApiError {
	skills := []models.ParticipantSkill{}

	exist := service.db.Where("participant_id = ?", participantId).
		Find(&skills).RowsAffected > 0

	if !exist {
		skillRequest := []models.ParticipantSkill{}

		for _, skill := range request.Skills {
			id, err := utils.ParseUUid(skill.Id)

			if err != nil {
				return err
			}

			skillRequest = append(skillRequest, models.ParticipantSkill{
				SkillId:       id,
				ParticipantId: participantId,
				IsMastered:    skill.IsMastered,
				UsedFor:       int8(skill.UsedFor),
			})
		}

		err := service.db.Create(&skillRequest).Error

		if err != nil {
			zap.L().Error("error storing skills", zap.Error(err))

			return &dto.ApiError{
				Typ:          dto.ErrorExec,
				ErrorMessage: "Gagal menyimpan skill",
				Err:          err,
			}
		}

		return nil
	}

	newSkills := []models.ParticipantSkill{}
	skillIds := []uuid.UUID{}

	for _, skill := range request.Skills {
		skillId, err := utils.ParseUUid(skill.Id)

		if err != nil {
			return err
		}

		s := models.ParticipantSkill{
			ParticipantId: participantId,
			SkillId:       skillId,
			UsedFor:       int8(skill.UsedFor),
			IsMastered:    skill.IsMastered,
		}

		skillIds = append(skillIds, skillId)

		skillExist := service.db.
			Where("participant_id = ?", participantId).
			Where("skill_id = ?", skillId).
			Find(&s).RowsAffected > 0

		if skillExist {
			s.UsedFor = int8(skill.UsedFor)
			if err := service.db.Save(&s).Error; err != nil {
				return &dto.ApiError{
					Typ:          dto.ErrorExec,
					ErrorMessage: "Gagal menyimpan skill",
					Err:          err,
				}
			}

			continue
		}

		newSkills = append(newSkills, s)
	}

	if len(newSkills) > 0 {
		if err := service.db.Create(&newSkills).Error; err != nil {
			return &dto.ApiError{
				Typ:          dto.ErrorExec,
				ErrorMessage: "Gagal menyimpan skill",
				Err:          err,
			}
		}
	}

	if len(skillIds) > 0 {
		err := service.db.
			Where("participant_id = ?", participantId).
			Where("skill_id not in ?", skillIds).
			Unscoped().
			Delete(&models.ParticipantSkill{}).
			Error

		if err != nil {
			return &dto.ApiError{
				Typ:          dto.ErrorExec,
				ErrorMessage: "Gagal menyimpan skill",
				Err:          err,
			}
		}
	}

	return nil
}

func (service ParticipantService) checkProfileCompleted(participant models.Participant) bool {
	return participant.ParticipantRole != nil &&
		participant.ParticipantDepartment != nil &&
		participant.ParticipantSkills != nil
}

func (service ParticipantService) AssignRoles(userId uuid.UUID, req requests.CreateParticipantRoleRequest) *dto.ApiError {
	mainRoleId, errParse := utils.ParseUUid(req.MainRoleId)
	if errParse != nil {
		return errParse
	}

	participantRole := models.ParticipantRole{
		MainRoleId:      mainRoleId,
		UserId:          userId,
		SecondaryRoleId: nil,
		InterestRoleId:  nil,
	}

	if req.SecondaryRoleId != nil {
		secondaryRoleId, errParse := utils.ParseUUid(*req.SecondaryRoleId)
		if errParse != nil {
			return errParse
		}

		*participantRole.SecondaryRoleId = secondaryRoleId
	}

	if req.InterestRoleId != nil {
		interestRoleId, errParse := utils.ParseUUid(*req.InterestRoleId)
		if errParse != nil {
			return errParse
		}

		*participantRole.InterestRoleId = interestRoleId
	}

	err := service.db.Create(&participantRole).Error

	if err != nil {
		zap.L().Error("error assign participant role", zap.Error(err))

		return dto.InternalError(err)
	}

	return nil
}

func (service ParticipantService) UpdateRoles(userId uuid.UUID, req requests.UpdateParticipantRoleRequest) *dto.ApiError {
	mainRoleId, errParse := utils.ParseUUid(req.MainRoleId)
	if errParse != nil {
		return errParse
	}

	participantRole := models.ParticipantRole{
		MainRoleId:      mainRoleId,
		UserId:          userId,
		SecondaryRoleId: nil,
		InterestRoleId:  nil,
	}

	if req.SecondaryRoleId != nil {
		secondaryRoleId, errParse := utils.ParseUUid(*req.SecondaryRoleId)
		if errParse != nil {
			return errParse
		}

		*participantRole.SecondaryRoleId = secondaryRoleId
	}

	if req.InterestRoleId != nil {
		interestRoleId, errParse := utils.ParseUUid(*req.InterestRoleId)
		if errParse != nil {
			return errParse
		}

		*participantRole.InterestRoleId = interestRoleId
	}

	err := service.db.
		Where("user_id = ?", userId).
		Updates(&participantRole).Error

	if err != nil {
		zap.L().Error("error update participant role", zap.Error(err))

		return dto.InternalError(err)
	}

	return nil
}

func (service ParticipantService) AssignDepartment(userId uuid.UUID, req requests.CreateParticipantDepartmentRequest) *dto.ApiError {
	departmentId, errParse := utils.ParseUUid(req.DepartmentId)
	if errParse != nil {
		return errParse
	}

	unitId, errParse := utils.ParseUUid(req.DepartmentUnitId)
	if errParse != nil {
		return errParse
	}

	teamId, errParse := utils.ParseUUid(req.DepartmentTeamId)
	if errParse != nil {
		return errParse
	}

	participantDepartment := models.ParticipantDepartment{
		UserId:           userId,
		DepartmentId:     departmentId,
		DepartmentUnitId: unitId,
		DepartmentTeamId: teamId,
	}

	err := service.db.Create(&participantDepartment).Error

	if err != nil {
		zap.L().Error("error assign participant department", zap.Error(err))

		return dto.InternalError(err)
	}

	return nil
}
