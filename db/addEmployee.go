package db

import (
	"EmployeeManagementSystemProject/database"
	"fmt"
)

//AddEmployeeDetails create a new employee record
func AddEmployeeDetails(healthInsurance, lifeInsurance int, employee *models.Employee) (int, error) {
	db := database.SqlClient()

	result, err := db.Query("INSERT INTO employeeOfficial (designation,department,teamLead,jobType,healthInsurance,lifeInsurance,userId,salary)VALUES ( '" + employee.Designation + "' ,'" + employee.Department + "','" + employee.TeamLead + "','" + employee.JobType + "','" + fmt.Sprintf("%v", healthInsurance) + "','" + fmt.Sprintf("%v", lifeInsurance) + "','" + employee.UserId + "','" + fmt.Sprintf("%v", employee.Salary) + "')")
	if err != nil {
		return 500, err
	}
	defer result.Close()
	return 200, nil
}
