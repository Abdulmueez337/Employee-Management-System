package service

import (
	"fmt"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/client"
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/database/db"
)

//DeleteDetails call delete employee db function
func DeleteDetails(UserId, jobtype, tokenAuth string) (int, error) {
	//Extract userId from token
	newUserId, check := UserIdExtract(tokenAuth)
	if check == 500 {
		fmt.Println("UserID service Error:", check)
		return 500, nil
	}

	//Check user is existed
	designation, err := db.CheckEmployee(UserId)
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
		return 500, err
	}
	if designation != "" {
		newDesignation, _ := db.CheckEmployee(newUserId)
		fmt.Println("Token User Designation is :", newDesignation)

		//Check the user is authorized for this api
		newClient := client.NewRollBaseClient()
		authorizeResponse := (*client.RollBaseClient).GetRole(newClient, newDesignation, "DeleteDetails")
		if authorizeResponse == 200 {
			return db.DeleteEmployeeDetails(UserId, jobtype)
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
