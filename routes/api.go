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
	tresholdHandlers := handlers.Treshold()
	parameterHandlers := handlers.Parameter()

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

		backOffice.GET("treshold", tresholdHandlers.GetTresholdList)
		backOffice.POST("treshold", tresholdHandlers.AddTreshold)
		backOffice.PUT("treshold", tresholdHandlers.UpdateTreshold)
		backOffice.DELETE("treshold/:id", tresholdHandlers.DeleteTreshold)

		backOffice.GET("department", departmentHandlers.GetDepartments)
		backOffice.POST("department", departmentHandlers.AddDepartment)
		backOffice.PUT("department", departmentHandlers.UpdateDepartment)
		backOffice.DELETE("department/:id", departmentHandlers.DeleteDepartment)

		backOffice.GET("department/:id/teams", departmentHandlers.GetDepartmentTeams)
		backOffice.POST("department/:id/teams", departmentHandlers.AddDepartmentTeam)
		backOffice.PUT("department/:id/teams", departmentHandlers.UpdateDepartmentTeam)
		backOffice.DELETE("department/:id/teams/:teamId", departmentHandlers.DeleteDepartmentTeam)

		backOffice.GET("department/:id/units", departmentHandlers.GetDepartmentUnits)
		backOffice.POST("department/:id/units", departmentHandlers.AddDepartmentUnit)
		backOffice.PUT("department/:id/units", departmentHandlers.UpdateDepartmentUnit)
		backOffice.DELETE("department/:id/units/:unitId", departmentHandlers.DeleteDepartmentUnit)

		backOffice.GET("parameter", parameterHandlers.GetParameterList)
		backOffice.POST("parameter/score", parameterHandlers.AddParameterScore)
		backOffice.POST("parameter/difficulty", parameterHandlers.AddParameterDifficulty)
		backOffice.PUT("parameter/score", parameterHandlers.UpdateParameterScore)
		backOffice.PUT("parameter/difficulty", parameterHandlers.UpdateParameterDifficulty)
		backOffice.DELETE("parameter/score/:id", parameterHandlers.DeleteParameterScore)
		backOffice.DELETE("parameter/difficulty/:id", parameterHandlers.DeleteParameterDifficulty)
	}

}
