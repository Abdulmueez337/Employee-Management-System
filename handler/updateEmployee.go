package handler

import (
	"fmt"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/gen/restapi/operations/admin"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/service"
	"github.com/go-openapi/runtime/middleware"
)

type updateEmployee struct {
}

func UpdateEmployee() admin.UpdateEmployeeHandler {
	return &updateEmployee{}
}

//Handle call update details service layer function
func (e *updateEmployee) Handle(params admin.UpdateEmployeeParams, i interface{}) middleware.Responder {
	tokenAuth := params.HTTPRequest.Header.Get("Authorization")
	//Call service layer
	result, err := service.UpdateDetails(params.UserID, toUpdateEmployeeDomain(params.Body),tokenAuth)
	if err != nil {
		fmt.Errorf("INTERNAL SERVER ERROR: %s", err)
		return admin.NewUpdateEmployeeInternalServerError().WithPayload("Internal Server Error")
	} else if result == 401 {
		return admin.NewUpdateEmployeeUnauthorized().WithPayload("Not Authorized")
	} else if result == 404 {
		return admin.NewUpdateEmployeeNotFound().WithPayload("Employee Not Found")
	} else if result == 200 {
		return admin.NewUpdateEmployeeOK().WithPayload("Successfully Update")
	} else {
		return admin.NewUpdateEmployeeInternalServerError()
	}
}
