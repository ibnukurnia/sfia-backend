package routes

import (
	"sv-sfia/handlers"
	"sv-sfia/middleware"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
)

func InitApiRouter(e *gin.Engine, services *services.ServiceProvider) {
	handlers := handlers.NewHandlerProvider(services)

	assessmentHandlers := handlers.Assesment()

	participantHandlers := handlers.Participant()
	skillHandlers := handlers.Skill()
	roleHandlers := handlers.Role()
	departmentHandlers := handlers.Department()
	roleGroupHandlers := handlers.RoleGroup()

	v1 := e.Group("api/v1")

	// roles
	v1.GET("roles", roleHandlers.GetRoles)

	// departments
	v1.GET("departments", departmentHandlers.GetDepartments)
	v1.GET("departments/:id/teams", departmentHandlers.GetDepartmentTeams)
	v1.GET("departments/:id/units", departmentHandlers.GetDepartmentUnits)

	// assesments

	authEp := v1.Group("auth")
	{
		authEp.POST("participant/sign-up", participantHandlers.Register)
		authEp.POST("participant/sign-in", participantHandlers.Login)
	}

	v1.GET("skills", skillHandlers.GetSkills)

	v1.Use(middleware.AssessmentJWT(services.ParticipantService.GetByParticipantId))
	v1.GET("participant/profile", participantHandlers.Profile)
	v1.POST("onboarding/general-information", participantHandlers.StorePersonalInformation)
	v1.POST("onboarding/role", participantHandlers.CreateParticipantRole)
	v1.POST("onboarding/skills", participantHandlers.StoreParticipantSkill)

	v1.POST("participant/personal-information", participantHandlers.StorePersonalInformation)
	v1.POST("participants/role", participantHandlers.CreateParticipantRole)
	v1.GET("participants/role/skills", participantHandlers.GetParticipantRoleSkills)

	v1.POST("participants/assign-skills", participantHandlers.StoreParticipantSkill)
	v1.GET("participants/tools", participantHandlers.GetParticipantTool)
	v1.POST("participants/tools", participantHandlers.CreateParticipantTool)

	assessmentsEp := v1.Group("assessments")
	{
		assessmentsEp.POST("new", assessmentHandlers.CreateNewAssessment)

		assessmentsEp.GET("self-assessment", assessmentHandlers.GetSelfAssessments)
		assessmentsEp.POST("self-assessment", assessmentHandlers.SaveSelfAssessmentAnswer)

		assessmentsEp.GET("duj", assessmentHandlers.GetDujAssesments)
		assessmentsEp.POST("duj", assessmentHandlers.SaveDujAnswer)

		assessmentsEp.GET("tool", assessmentHandlers.GetToolAssessment)
		assessmentsEp.POST("tool", assessmentHandlers.SaveToolAssessmentAnswers)

		assessmentsEp.GET("trainings", participantHandlers.GetParticipantRoleTraining)
		assessmentsEp.POST("trainings", participantHandlers.CreateParticipantTraining)
	}

	backOffice := v1.Group("backoffice")
	{
		backOffice.GET("role-group", roleGroupHandlers.GetRoleGroup)
		backOffice.POST("role-group", roleGroupHandlers.AddRoleGroup)
		backOffice.PUT("role-group", roleGroupHandlers.UpdateRoleGroup)
		backOffice.DELETE("role-group/:id", roleGroupHandlers.DeleteRoleGroup)

		backOffice.GET("role", roleHandlers.GetRoleList)
		backOffice.POST("role", roleHandlers.AddRole)
		backOffice.PUT("role", roleHandlers.UpdateRole)
		backOffice.DELETE("role/:id", roleHandlers.DeleteRole)

		backOffice.GET("skillset", skillHandlers.GetSkillsetList)
		backOffice.POST("skillset", skillHandlers.AddSkillset)
		backOffice.PUT("skillset", skillHandlers.UpdateSkillset)
		backOffice.DELETE("skillset/:id", skillHandlers.DeleteSkillset)
	}

}
