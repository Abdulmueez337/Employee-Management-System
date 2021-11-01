package db

//CheckEmployee check from db user exist or not
func CheckEmployee(userId string) (string, error) {
	db := database.SqlClient()
	query, err := db.Query("SELECT designation FROM employeeOfficial WHERE userid='" + userId + "'")
	if err != nil {
		return "", err
	}
	var designationEmployee string
	for query.Next() {
		err = query.Scan(&designationEmployee)
		if err != nil {
			return "", err
		}
	}
	return designationEmployee, nil
}
