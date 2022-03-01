package entities

import (
	"gorm-postgres/enums"
)

type Order struct {
	Entity
	Name                 string
	EmailAddress         string                   `json:"email_address"`
	PhoneNumber          string                   `json:"phone_number"`
	PersonalUrl          string                   `json:"personal_url"`
	YearsOfAge           int                      `json:"years_of_age"`
	IsExternalContractor bool                     `json:"is_external_contractor"`
	RelationshipStatus   enums.RelationshipStatus `json:"relationship_status"`
}
