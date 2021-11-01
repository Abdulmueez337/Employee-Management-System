package db

//DeleteEmployeeDetails update employee job type status resign or terminate
func DeleteEmployeeDetails(userId string, jobType string) (int, error) {
	db := database.SqlClient()
	_, err := db.Exec("UPDATE employeeOfficial SET jobType='" + jobType + "' where userid='" + userId + "'")
	if err != nil {
		return 500, err
	}
	return 200, err
}
