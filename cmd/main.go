package main

import (
	"flag"
	"fmt"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/client"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/gen/restapi"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/gen/restapi/operations"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/handler"
	"github.com/go-openapi/loads"
	"log"
)

//Specify Server port
var portFlag = flag.Int("port", 3001, "Port to run this service on")

func main() {
	//Link with swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	//Connect with swagger api and server
	api := operations.NewEmployeeManagemntSystemAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()
	flag.Parse()
	server.Port = *portFlag
	//connecting swagger handler with domain handler
	api.AdminAddEmployeeHandler = handler.AddEmployee()
	api.AdminShowEmployeeHandler = handler.ShowEmployeeAdmin()
	api.AdminUpdateEmployeeHandler = handler.UpdateEmployee()
	api.AdminDeleteEmployeeHandler = handler.DeleteEmployee()
	api.EmployeeShowEmployeeSelfHandler = handler.ShowEmployeeSelf()
	api.TeamLeadShowEmployeeTeamHandler = handler.ShowEmployeeTeamLead()
	//Verify token from the login api
	api.BearerAuth = func(token string) (interface{}, error) {

		authClient := client.NewTokenBaseClient()
		authresponse := (*client.TokenBaseClient).GetValidate(authClient, token)
		if authresponse == 200 {
			fmt.Println("Token is valid")
			return api, nil
		} else if authresponse == 400 {
			fmt.Println("Invalid Token")
			return api, nil
		} else {
			fmt.Println("Token > INTERNAL SERVER ERROR")
			return api, nil
		}
	}
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
