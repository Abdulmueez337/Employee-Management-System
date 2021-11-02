package handler

import (
	"fmt"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/gen/restapi/operations/team_lead"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/service"
	"github.com/go-openapi/runtime/middleware"
)

type showEmployeeTeamLead struct {
}

func ShowEmployeeTeamLead() team_lead.ShowEmployeeTeamHandler {
	return &showEmployeeTeamLead{}
}

//Handle call show team members details to team lead service function
func (e *showEmployeeTeamLead) Handle(params team_lead.ShowEmployeeTeamParams) middleware.Responder {
	//Call service layer
	result, status, err := service.ShowDetailsTeamLead(params.UserID)
	if err != nil {
		fmt.Errorf("INTERNAL SERVER ERROR: %s", err)
		return team_lead.NewShowEmployeeTeamInternalServerError().WithPayload("Internal Server Error")
	} else if status == 401 {
		return team_lead.NewShowEmployeeTeamUnauthorized().WithPayload("Not Authorized")
	} else if status == 404 {
		return team_lead.NewShowEmployeeTeamNotFound().WithPayload("Not Found")
	} else {
		return team_lead.NewShowEmployeeTeamOK().WithPayload(toTeamLeadEmployeeGen(result))
	}
}