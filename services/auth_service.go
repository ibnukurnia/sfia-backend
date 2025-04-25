package services

import (
	"fmt"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func newAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (service AuthService) hashPassword(pass string) (string, *dto.ApiError) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), 10)

	if err != nil {
		zap.L().Error("failed to hash password", zap.Error(err))

		return "", dto.InternalError(err)
	}

	return string(hashed), nil
}

func (service AuthService) genereteToken(userId string) (string, *dto.ApiError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"expire":  time.Now().Add(time.Hour * (24 * 30)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(""))
	if err != nil {
		zap.L().Error("error creating jwt token", zap.Error(err))

		return "", dto.InternalError(err)
	}

	return tokenString, nil
}

func (service AuthService) checkPassword(plain string, hashed string) *dto.ApiError {
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

func (service AuthService) Register(req requests.RegisterRequest) (*responses.AuthResponse, *dto.ApiError) {
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
			Err:          fmt.Errorf("user already registered"),
			ErrorMessage: "Email sudah terdaftar",
		}
	} else {
		if err != gorm.ErrRecordNotFound {

			zap.L().Error("error register user", zap.Error(err))

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

func (service AuthService) Login(req requests.LoginRequest) (*responses.AuthResponse, *dto.ApiError) {
	user := models.User{}

	err := service.db.
		Where("email = ?", req.Email).
		First(&user).
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

	passErr := service.checkPassword(req.Password, user.Password)
	if passErr != nil {
		return nil, passErr
	}

	token, tokenErr := service.genereteToken(user.Uuid.String())
	if tokenErr != nil {
		return nil, tokenErr
	}

	return &responses.AuthResponse{
		Token: token,
		User: responses.UserResponse{
			Name:  user.Name,
			Email: user.Email,
			Pn:    user.Pn,
			Role:  user.RoleAccess,
		},
	}, nil
}
