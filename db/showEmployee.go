package db

//ShowEmployeeDetails get employee record from bd and show it to who'll be able to access
func ShowEmployeeDetails(userId string) (models.Employee, int, error) {
	db := database.SqlClient()
	var employee models.Employee
	query, err := db.Query("SELECT * FROM employeeOfficial WHERE userid='" + userId + "'")
	if err != nil {
		return employee, 500, err
	}
	defer query.Close()
	for query.Next() {
		err := query.Scan(&employee.Designation, &employee.Department, &employee.TeamLead, &employee.JobType, &employee.HealthInsurance,
			&employee.LifeInsurance, &employee.UserId, &employee.Salary)
		if err != nil {
			return employee, 500, err
		}
	}
	return employee, 200, nil
}

