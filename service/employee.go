package service

import (
	"fmt"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/database/db"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/models"
	"strconv"
)

//AddDetails call add details db function
func AddDetails(employee *models.Employee) (int, error) {
	//Need to add authorization function for response 401
	designation, err := db.CheckEmployee(employee.UserId)
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
		return 500, err
	}
	if designation != "" {
		return 400, nil
	}
	healthInsurance, err := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert Bool to String, error : %v", err)
	}
	lifeInsurance, _ := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert Bool to String, error : %v", err)
	}
	return db.AddEmployeeDetails(healthInsurance, lifeInsurance, employee)
}

//DeleteDetails call delete employee db function
func DeleteDetails(UserId string, jobtype string) (int, error) {
	//Need to add authorization function for response 401
	designation, err := db.CheckEmployee(UserId)
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
		return 500, err
	}
	if designation != "" {
		return db.DeleteEmployeeDetails(UserId, jobtype)
	} else {
		return 404, nil
	}
}

//ShowDetailsAdmin call show details to admin db function
func ShowDetailsAdmin(userid string) (models.Employee, int, error) {
	//Need to add authorization function for response 401
	designation, err := db.CheckEmployee(userid)
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
		return models.Employee{}, 500, err
	}
	if designation != "" {
		return db.ShowEmployeeDetails(userid)
	}
	return models.Employee{}, 404, nil
}

//ShowDetailsEmployeeSelf call show team members details to team lead db function
func ShowDetailsEmployeeSelf(userId string) (models.Employee, int, error) {
	//Need to add authorization function for response 401
	designation, err := db.CheckEmployee(userId)
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
		return models.Employee{}, 500, err
	}
	if designation != "" {
		return db.ShowEmployeeDetails(userId)
	}
	return models.Employee{}, 404, nil
}

//ShowDetailsTeamLead call show details to employee db function
func ShowDetailsTeamLead(userId string) (models.Employee, int, error) {
	//Need to add authorization function for response 401
	designation, err := db.CheckEmployee(userId)
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
		return models.Employee{}, 500, err
	}
	if designation != "" {
		return db.ShowDetailsTeamLead(userId)
	}
	return models.Employee{}, 404, nil
}

//UpdateDetails call update details db function
func UpdateDetails(userId string, employee *models.Employee) (int, error) {
	//Need to add authorization function for response 401
	designation, err := db.CheckEmployee(userId)
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
		return 500, err
	}
	if designation != "" {
		healthInsurance, err := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
		if err != nil {
			fmt.Errorf("failed to convert string to int, error : %v", err)
		}
		lifeInsurance, _ := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
		if err != nil {
			fmt.Errorf("failed to convert string to int, error : %v", err)
		}
		return db.UpdateEmployeeDetails(healthInsurance, lifeInsurance, userId, employee)
	} else {
		return 404, nil
	}
}
