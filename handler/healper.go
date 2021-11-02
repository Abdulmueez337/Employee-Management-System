package handler

import (
	"github.com/Abdulmueez337/EmployeeManagementSystemProject/gen/models"
	domain "github.com/Abdulmueez337/EmployeeManagementSystemProject/models"
	"github.com/go-openapi/swag"
)

// toEmployeeDomain converts gen to domain model
func toEmployeeDomain(employee *models.EmployeeOfficial) *domain.Employee{
	return &domain.Employee{
		UserId:          swag.StringValue(employee.UserID),
		Salary:          int(swag.Int64Value(employee.Salary)),
		Designation:     swag.StringValue(employee.Designation),
		Department:      swag.StringValue(employee.Department),
		TeamLead:        swag.StringValue(employee.TeamLead),
		JobType:         swag.StringValue(employee.JobType),
		HealthInsurance: swag.BoolValue(employee.HealthInsurance),
		LifeInsurance:   swag.BoolValue(employee.LifeInsurance),
	}

}

// toEmployeeGen converts domain to gen model
func toEmployeeGen(domain domain.Employee) *models.EmployeeOfficial {
	return &models.EmployeeOfficial{
		UserID:          swag.String(domain.UserId),
		Salary:          swag.Int64(int64(domain.Salary)),
		Designation:     swag.String(domain.Designation),
		Department:      swag.String(domain.Department),
		TeamLead:        swag.String(domain.TeamLead),
		JobType:         swag.String(domain.JobType),
		HealthInsurance: swag.Bool(domain.HealthInsurance),
		LifeInsurance:   swag.Bool(domain.LifeInsurance),
	}
}

func toUpdateEmployeeDomain(employee *models.UpdateEmployeeOfficial) *domain.Employee {
	return &domain.Employee{
		Salary:          int(swag.Int64Value(employee.Salary)),
		Designation:     swag.StringValue(employee.Designation),
		Department:      swag.StringValue(employee.Department),
		TeamLead:        swag.StringValue(employee.TeamLead),
		JobType:         swag.StringValue(employee.JobType),
		HealthInsurance: swag.BoolValue(employee.HealthInsurance),
		LifeInsurance:   swag.BoolValue(employee.LifeInsurance),
	}

}

/*func toUpdateEmployeeGen(domain domain.Employee) *models.UpdateEmployeeOfficial {
	return &models.UpdateEmployeeOfficial{
		Salary:          swag.Int64(int64(domain.Salary)),
		Designation:     swag.String(domain.Designation),
		Department:      swag.String(domain.Department),
		TeamLead:        swag.String(domain.TeamLead),
		JobType:         swag.String(domain.JobType),
		HealthInsurance: swag.Bool(domain.HealthInsurance),
		LifeInsurance:   swag.Bool(domain.LifeInsurance),
	}
}*/

// toEmployeeGen converts domain to gen model
func toTeamLeadEmployeeGen(domain domain.Employee) *models.TeamLeadEmployee {
	return &models.TeamLeadEmployee{
		Designation: swag.String(domain.Designation),
		Department:  swag.String(domain.Department),
		JobType:     swag.String(domain.JobType),
	}
}
