package service

import (
	"fmt"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/client"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/database/db"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/models"
)

//ShowDetailsTeamLead call show details to employee db function
func ShowDetailsTeamLead(userId, tokenAuth string) (models.Employee, int, error) {
	//Extract userId from token
	newUserId, check := UserIdExtract(tokenAuth)
	if check == 500 {
		fmt.Println("UserID service Error:", check)
		return models.Employee{}, 500, nil
	}
	//Check user is existed
	designation, err := db.CheckEmployee(userId)
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
		return models.Employee{}, 500, err
	}
	if designation != "" {
		newDesignation, _ := db.CheckEmployee(newUserId)
		fmt.Println("Token User Designation is :", newDesignation)
		//Check the user is authorized for this api
		newClient := client.NewRollBaseClient()
		authorizeResponse := (*client.RollBaseClient).GetRole(newClient, newDesignation, "showdetailsteamlead")
		if authorizeResponse == 200 {
			//var checkTeamLead models.Employee
			checkTeamLeadName, _ := db.DataUpdateEmployee(newUserId)
			checkEmplTL, _ := db.DataUpdateEmployee(userId)
			if checkTeamLeadName.TeamLead == checkEmplTL.TeamLead {
				return db.ShowDetailsTeamLead(userId)
			} else {
				return models.Employee{}, 401, nil
			}
		} else if authorizeResponse == 400 {
			return models.Employee{}, 400, nil
		} else if authorizeResponse == 403 {
			return models.Employee{}, 401, nil
		} else {
			return models.Employee{}, 500, nil
		}
	}
	return models.Employee{}, 404, nil
}
