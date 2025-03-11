package services

import "gorm.io/gorm"

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
	TresholdService	   *ParameterService
	ParameterService   *ParameterService
}

func NewServiceProvider(db *gorm.DB) *ServiceProvider {
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
	}
}
