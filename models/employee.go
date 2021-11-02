package models

//Employee use as a domain model
type Employee struct {
	UserId          string `json:"user_id" structs:"user_id"`
	Salary          int    `json:"salary" structs:"salary"`
	Designation     string `json:"designation" structs:"designation"`
	Department      string `json:"department" structs:"department"`
	TeamLead        string `json:"team_lead" structs:"team_lead"`
	JobType         string `json:"job_type" structs:"job_type"`
	HealthInsurance bool   `json:"health_insurance" structs:"health_insurance"`
	LifeInsurance   bool   `json:"life_insurance" structs:"life_insurance"`
}
