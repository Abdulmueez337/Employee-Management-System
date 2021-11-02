package handler

import (
	"fmt"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/gen/restapi/operations/admin"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/service"
	"github.com/go-openapi/runtime/middleware"
)

type showEmployeeAdmin struct {
}

func ShowEmployeeAdmin() admin.ShowEmployeeHandler {
	return &showEmployeeAdmin{}
}

//Handle call show details to admin service function
func (e *showEmployeeAdmin) Handle(params admin.ShowEmployeeParams) middleware.Responder {
	//Call service layer
	employee, result, err := service.ShowDetailsAdmin(params.UserID)
	if err != nil {
		fmt.Errorf("INTERNAL SERVER ERROR: %s", err)
		return admin.NewShowEmployeeInternalServerError().WithPayload("Internal Server Error")
	} else if result == 401 {
		return admin.NewShowEmployeeUnauthorized().WithPayload("Not Authorized")
	} else if result == 404 {
		return admin.NewShowEmployeeNotFound().WithPayload("Employee not Exist")
	} else {
		return admin.NewShowEmployeeOK().WithPayload(toEmployeeGen(employee))
	}
}