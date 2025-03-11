package handlers

import "sv-sfia/services"

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
	}
}

func (provider *HandlerProvider) RoleGroup() *roleGroupHandler {
	return &roleGroupHandler{
		roleGroupService: provider.serviceProvider.RoleGroupService,
	}
}

func (provider *HandlerProvider) Treshold() *parameterHandler {
	return &parameterHandler{
		parameterService: provider.serviceProvider.TresholdService,
	}
}

func (provider *HandlerProvider) Parameter() *parameterHandler {
	return &parameterHandler{
		parameterService: provider.serviceProvider.ParameterService,
	}
}