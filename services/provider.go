package services

import (
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type ServiceProvider struct {
	AssessmentService  *AssessmentService
	RoleService        *RoleService
	DepartmentService  *DepartmentService
	ParticipantService *ParticipantService
	SkillService       *SkillService
	SfiaService        *SfiaService
	DujService         *DujService
	ToolService        *ToolService
	TrainingService    *TrainingService
	RoleGroupService   *RoleGroupService
	TresholdService	   *TresholdService
	ParameterService   *ParameterService
	DujAdminService	*DujAdminService
	UserAdminService *UserAdminService
	TrainingMasterService *TrainingMasterService
	ToolsMasterServices *ToolsMasterService
}

func NewServiceProvider(db *gorm.DB, minioClient *minio.Client) *ServiceProvider {
	return &ServiceProvider{
		AssessmentService:  newAssessmentService(db),
		RoleService:        newRoleService(db),
		DepartmentService:  newDeparmentService(db),
		ParticipantService: newParticipantService(db),
		SkillService:       newSkillService(db),
		SfiaService:        newSfiaService(db),
		DujService:         newDujService(db),
		ToolService:        newToolService(db),
		TrainingService:    newTrainingService(db),
		RoleGroupService: newRoleGroupService(db),
		TresholdService: newTresholdService(db),
		ParameterService:  newParameterService(db),
		DujAdminService: newDujAdminService(db),
		UserAdminService: newUserAdminService(db),
		TrainingMasterService: newTrainingMasterService(db),
		ToolsMasterServices: newToolsMasterService(db, minioClient),
	}
}
