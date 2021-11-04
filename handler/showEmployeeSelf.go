package handler

import (
	"fmt"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/gen/restapi/operations/employee"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/service"
	"github.com/go-openapi/runtime/middleware"
)

type showEmployeeSelf struct {
}

func ShowEmployeeSelf() employee.ShowEmployeeSelfHandler {
	return &showEmployeeSelf{}
}

//Handle call show details to employee  service function
func (e *showEmployeeSelf) Handle(params employee.ShowEmployeeSelfParams, i interface{}) middleware.Responder {
	tokenAuth := params.HTTPRequest.Header.Get("Authorization")
	//Call service layer
	result, status, err := service.ShowDetailsEmployeeSelf(params.UserID, tokenAuth)
	if err != nil {
		fmt.Errorf("INTERNAL SERVER ERROR: %s", err)
		return employee.NewShowEmployeeSelfInternalServerError().WithPayload("Internal Server Error")
	} else if status == 400 {
		return employee.NewShowEmployeeSelfBadRequest().WithPayload("Bad Request")
	} else if status == 401 {
		return employee.NewShowEmployeeSelfUnauthorized().WithPayload("Not Authorized")
	} else if status == 404 {
		return employee.NewShowEmployeeSelfNotFound().WithPayload("Employee Not Found")
	} else if status == 200 {
		return employee.NewShowEmployeeSelfOK().WithPayload(toEmployeeGen(result))
	} else {
		return employee.NewShowEmployeeSelfInternalServerError()
	}
}
