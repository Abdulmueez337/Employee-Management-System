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
func UpdateDetails(userId string, swaggerEmployee *models.Employee) (int, error) {
	//Need to add authorization function for response 401
	designation, err := db.CheckEmployee(userId)
	var newEmployee models.Employee
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
		return 500, err
	}
	if designation != "" {
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
	} else {
		return 404, nil
	}
}
