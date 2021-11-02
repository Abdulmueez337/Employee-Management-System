package handler

import (
	"fmt"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/gen/restapi/operations/admin"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/service"
	"github.com/go-openapi/runtime/middleware"
)

type addEmployee struct {
}

func AddEmployee() admin.AddEmployeeHandler {
	return &addEmployee{}
}

//Handle call add details service layer function
func (e *addEmployee) Handle(params admin.AddEmployeeParams) middleware.Responder {
	//Call service layer
	result, err := service.AddDetails(toEmployeeDomain(params.Body))
	if err != nil {
		fmt.Errorf("INTERNAL SERVER ERROR: %s", err)
		return admin.NewAddEmployeeInternalServerError()
	} else if result == 400 {
		return admin.NewUpdateEmployeeBadRequest().WithPayload("Employee Already Exist")
	} else if result == 401 {
		return admin.NewShowEmployeeUnauthorized().WithPayload("Not Authorized")
	} else if result == 404 {
		return admin.NewShowEmployeeNotFound().WithPayload("Employee Not Found")
	} else {
		return admin.NewAddEmployeeCreated().WithPayload("Employee Successfully Added")
	}
}
