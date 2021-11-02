package handler

import (
	"fmt"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/gen/restapi/operations/admin"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/service"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

type deleteEmployee struct {
}

func DeleteEmployee() admin.DeleteEmployeeHandler {
	return &deleteEmployee{}
}

//Handle call delete details service layer function
func (e *deleteEmployee) Handle(params admin.DeleteEmployeeParams) middleware.Responder {
	//Call service layer
	result, err := service.DeleteDetails(params.UserID, swag.StringValue(params.Body.JobType))
	if err != nil {
		fmt.Errorf("INTERNAL SERVER ERROR FOUND: %s", err)
		return admin.NewDeleteEmployeeInternalServerError().WithPayload("Internal Server Error")
	} else if result == 401 {
		return admin.NewDeleteEmployeeUnauthorized().WithPayload("Not Authorized")
	} else if result == 404 {
		return admin.NewDeleteEmployeeNotFound().WithPayload("Employee Not Exist")
	} else {
		return admin.NewDeleteEmployeeOK().WithPayload("Successfully Deleted")
	}
}
