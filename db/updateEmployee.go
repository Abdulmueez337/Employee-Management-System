package db

import "fmt"

//UpdateEmployeeDetails update the record of employee
func UpdateEmployeeDetails(healthInsurance, lifeInsurance int, userId string, employee *models.Employee) (int, error) {
	db := database.SqlClient()
	_, err := db.Query("UPDATE employeeOfficial SET designation='" + employee.Designation + "' ,department='" + employee.Department + "', teamLead='" + employee.TeamLead + "',jobType='" + employee.JobType + "',healthInsurance='" + fmt.Sprintf("%v", healthInsurance) + "',lifeInsurance='" + fmt.Sprintf("%v", lifeInsurance) + "',salary='" + fmt.Sprintf("%v", employee.Salary) + "' where userId='" + userId + "'")
	if err != nil {
		return 500, err
	}
	defer db.Close()
	return 200, nil
}

