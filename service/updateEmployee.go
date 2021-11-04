package service

import (
	"fmt"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/client"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/database/db"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/models"
	"strconv"
)

//UpdateDetails call update details db function
func UpdateDetails(userId string, swaggerEmployee *models.Employee, tokenAuth string) (int, error) {
	//Extract userId from token
	newUserId, check := UserIdExtract(tokenAuth)
	if check == 500 {
		fmt.Println("UserID service Error:", check)
		return 500, nil
	}
	designation, err := db.CheckEmployee(userId)
	var newEmployee models.Employee
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
		return 500, err
	}
	if designation != "" {
		newDesignation, _ := db.CheckEmployee(newUserId)
		fmt.Println("Token User Designation is :", newDesignation)

		//Check the user is authorized for this api
		newClient := client.NewRollBaseClient()
		authorizeResponse := (*client.RollBaseClient).GetRole(newClient, newDesignation, "UpdateDetails")
		if authorizeResponse == 200 {
			dbEmployee, err := db.DataUpdateEmployee(userId)
			if err != nil {
				fmt.Errorf("failed to convert string to int, error : %v", err)
				return 500, err
			} else {
				//Check and Update Department
				if dbEmployee.Department != swaggerEmployee.Department {
					if swaggerEmployee.Department != "" {
						dbEmployee.Department = swaggerEmployee.Department
					}
				}
				newEmployee.Department = dbEmployee.Department

				//Check and Update Designation
				if dbEmployee.Designation != swaggerEmployee.Designation {
					if swaggerEmployee.Designation != "" {
						dbEmployee.Designation = swaggerEmployee.Designation
					}
				}
				newEmployee.Designation = dbEmployee.Designation

				//Check and Update JobType
				if dbEmployee.JobType != swaggerEmployee.JobType {
					if swaggerEmployee.JobType != "" {
						dbEmployee.JobType = swaggerEmployee.JobType
					}
				}
				newEmployee.JobType = dbEmployee.JobType

				//Check and Update TeamLead
				if dbEmployee.TeamLead != swaggerEmployee.TeamLead {
					if swaggerEmployee.TeamLead != "" {
						dbEmployee.TeamLead = swaggerEmployee.TeamLead
					}
				}
				newEmployee.TeamLead = dbEmployee.TeamLead

				//Check and Update Salary
				if dbEmployee.Salary != swaggerEmployee.Salary {
					if swaggerEmployee.Salary > dbEmployee.Salary {
						dbEmployee.Salary = swaggerEmployee.Salary
					}
				}
				newEmployee.Salary = dbEmployee.Salary

				//Check and Update HealthInsurance
				if dbEmployee.HealthInsurance != swaggerEmployee.HealthInsurance {
					dbEmployee.HealthInsurance = swaggerEmployee.HealthInsurance
				}
				newEmployee.HealthInsurance = dbEmployee.HealthInsurance

				//Check and Update LifeInsurance
				if dbEmployee.LifeInsurance != swaggerEmployee.LifeInsurance {
					dbEmployee.LifeInsurance = swaggerEmployee.LifeInsurance
				}
				newEmployee.LifeInsurance = dbEmployee.LifeInsurance
			}
			healthInsurance, err := strconv.Atoi(strconv.FormatBool(newEmployee.HealthInsurance))
			if err != nil {
				fmt.Errorf("failed to convert string to int, error : %v", err)
			}
			lifeInsurance, _ := strconv.Atoi(strconv.FormatBool(newEmployee.HealthInsurance))
			if err != nil {
				fmt.Errorf("failed to convert string to int, error : %v", err)
			}
			return db.UpdateEmployeeDetails(healthInsurance, lifeInsurance, userId, newEmployee)
		} else if authorizeResponse == 400 {
			return 400, nil
		} else if authorizeResponse == 403 {
			return 401, nil
		} else {
			return 500, nil
		}
	} else {
		return 404, nil
	}

}
