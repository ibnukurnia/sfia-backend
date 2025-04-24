package handlers

import (
	"sv-sfia/handlers/management"
	"sv-sfia/services"
)

type HandlerProvider struct {
	serviceProvider *services.ServiceProvider
}

func NewHandlerProvider(serviceProvider *services.ServiceProvider) *HandlerProvider {
	return &HandlerProvider{
		serviceProvider: serviceProvider,
	}
}

func (provider *HandlerProvider) Role() *roleHandler {
	return &roleHandler{
		roleService: provider.serviceProvider.RoleService,
	}
}

func (provider *HandlerProvider) Department() *departmentHandler {
	return &departmentHandler{
		deparmentService: provider.serviceProvider.DepartmentService,
	}
}

func (provider *HandlerProvider) Assesment() *assessmentHandler {
	return &assessmentHandler{
		assessmentService: provider.serviceProvider.AssessmentService,
		sfiaService:       provider.serviceProvider.SfiaService,
		skillService:      provider.serviceProvider.SkillService,
		dujService:        provider.serviceProvider.DujService,
		toolService:       provider.serviceProvider.ToolService,
	}
}

func (provider *HandlerProvider) Participant() *participantHandler {
	return &participantHandler{
		participantService: provider.serviceProvider.ParticipantService,
		roleService:        provider.serviceProvider.RoleService,
		skillService:       provider.serviceProvider.SkillService,
		toolService:        provider.serviceProvider.ToolService,
		trainingService:    provider.serviceProvider.TrainingService,
	}
}

func (provider *HandlerProvider) Skill() *skillHandler {
	return &skillHandler{
		skillService: provider.serviceProvider.SkillService,
		roleSevice:   provider.serviceProvider.RoleService,
	}
}

func (provider *HandlerProvider) RoleGroup() *roleGroupHandler {
	return &roleGroupHandler{
		roleGroupService: provider.serviceProvider.RoleGroupService,
	}
}

func (provider *HandlerProvider) Treshold() *tresholdHandler {
	return &tresholdHandler{
		tresholdService: provider.serviceProvider.TresholdService,
	}
}

func (provider *HandlerProvider) Parameter() *parameterHandler {
	return &parameterHandler{
		parameterService: provider.serviceProvider.ParameterService,
	}
}

func (provider *HandlerProvider) DujAdmin() *dujAdminHandler {
	return &dujAdminHandler{
		dujAdminService: provider.serviceProvider.DujAdminService,
	}
}

func (provider *HandlerProvider) UserAdmin() *userAdminHandler {
	return &userAdminHandler{
		userAdminService: provider.serviceProvider.UserAdminService,
	}
}

func (provider *HandlerProvider) TrainingMaster() *trainingMasterHandler {
	return &trainingMasterHandler{
		trainingMasterService: provider.serviceProvider.TrainingMasterService,
	}
}

func (provider *HandlerProvider) Tools() *toolsHandler {
	return &toolsHandler{
		toolsService: provider.serviceProvider.ToolsMasterServices,
	}
}
func (provider *HandlerProvider) Tool() *toolHandler {
	return &toolHandler{
		toolService: provider.serviceProvider.ToolService,
	}
}

func (provider *HandlerProvider) ManagementTalent() *management.ManagementTalentHandler {
	return &management.ManagementTalentHandler{}
}

func (provider *HandlerProvider) ManagementResume() *management.ManagementResumeHandler {
	return &management.ManagementResumeHandler{}
}

func (provider *HandlerProvider) ManagementAplikasi() *management.ManagementAplikasiHandler {
	return &management.ManagementAplikasiHandler{}
}

func (provider *HandlerProvider) ManagementUseCase() *management.ManagementUseCaseHandler {
	return &management.ManagementUseCaseHandler{}
}

func (provider *HandlerProvider) ManagemenRoleAndSkill() *management.RoleAndSkillManagement {
	return &management.RoleAndSkillManagement{}
}
