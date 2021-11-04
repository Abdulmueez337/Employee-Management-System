package service

import (
	"fmt"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/client"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/database/db"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/models"
	"strconv"
)

//AddDetails call add details db function
func AddDetails(employee *models.Employee, tokenAuth string) (int, error) {
	//Extract userId from token
	newUserId, check := UserIdExtract(tokenAuth)
	if check == 500 {
		fmt.Println("UserID service Error:", check)
		return 500, nil
	}
	//Check user is existed
	designation, err := db.CheckEmployee(employee.UserId)
	if err != nil {
		fmt.Println("designation error")
		fmt.Errorf("failed to convert string to int, error : %v", err)
		//return 500, err
	}
	if designation != "" {
		fmt.Println("designation null")
		return 400, nil
	} else {
		//jest for testing given this designation name
		newDesignation, _ := db.CheckEmployee(newUserId)
		fmt.Println("Token User Designation is :", newDesignation)

		//Check the user is authorized for this api
		newClient := client.NewRollBaseClient()
		authorizeResponse := (*client.RollBaseClient).GetRole(newClient, newDesignation, "AddDetails")
		if authorizeResponse == 200 {
			healthInsurance, err := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
			if err != nil {
				fmt.Errorf("failed to convert Bool to String, error : %v", err)
			}
			lifeInsurance, _ := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
			if err != nil {
				fmt.Errorf("failed to convert Bool to String, error : %v", err)
			}
			return db.AddEmployeeDetails(healthInsurance, lifeInsurance, employee)
		} else if authorizeResponse == 400 {
			return 400, nil
		} else if authorizeResponse == 403 {
			return 401, nil
		} else {
			return 500, nil
		}
	}
}
