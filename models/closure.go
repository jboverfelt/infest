package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/pop/nulls"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Closure struct {
	ID             uuid.UUID    `json:"id" db:"id"`
	CreatedAt      time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at" db:"updated_at"`
	Name           string       `json:"name" db:"name"`
	Address        string       `json:"address" db:"address"`
	Reason         string       `json:"reason" db:"reason"`
	ClosureDate    time.Time    `json:"closureDate" db:"closuredate"`
	ReopenDate     nulls.Time   `json:"reopenDate" db:"reopendate"`
	ReopenComments nulls.String `json:"reopenComments" db:"reopencomments"`
}

// String is not required by pop and may be deleted
func (c Closure) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Closures is not required by pop and may be deleted
type Closures []Closure

// String is not required by pop and may be deleted
func (c Closures) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (c *Closure) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Name, Name: "Name"},
		&validators.StringIsPresent{Field: c.Address, Name: "Address"},
		&validators.StringIsPresent{Field: c.Reason, Name: "Reason"},
		&validators.TimeIsPresent{Field: c.ClosureDate, Name: "ClosureDate"},
	), nil
}

// ValidateSave gets run everytime you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (c *Closure) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run everytime you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (c *Closure) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
