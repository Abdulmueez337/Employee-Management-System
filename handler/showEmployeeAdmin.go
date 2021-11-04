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
func (e *showEmployeeAdmin) Handle(params admin.ShowEmployeeParams, i interface{}) middleware.Responder {
	tokenAuth := params.HTTPRequest.Header.Get("Authorization")
	//Call service layer
	employee, result, err := service.ShowDetailsAdmin(params.UserID, tokenAuth)
	if err != nil {
		fmt.Errorf("INTERNAL SERVER ERROR: %s", err)
		return admin.NewShowEmployeeInternalServerError().WithPayload("Internal Server Error")
	} else if result == 400 {
		return admin.NewShowEmployeeBadRequest().WithPayload("Bad Request")
	} else if result == 401 {
		return admin.NewShowEmployeeUnauthorized().WithPayload("Not Authorized")
	} else if result == 404 {
		return admin.NewShowEmployeeNotFound().WithPayload("Employee not Exist")
	} else if result == 200 {
		return admin.NewShowEmployeeOK().WithPayload(toEmployeeGen(employee))
	} else {
		return admin.NewShowEmployeeInternalServerError()
	}
}
