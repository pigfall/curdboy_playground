// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "user_id"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldRace holds the string denoting the race field in the database.
	FieldRace = "race"
	// EdgeCar holds the string denoting the car edge name in mutations.
	EdgeCar = "car"
	// EdgeDept holds the string denoting the dept edge name in mutations.
	EdgeDept = "dept"
	// CarFieldID holds the string denoting the ID field of the Car.
	CarFieldID = "car_id"
	// DeptFieldID holds the string denoting the ID field of the Dept.
	DeptFieldID = "dept_id"
	// Table holds the table name of the user in the database.
	Table = "users"
	// CarTable is the table that holds the car relation/edge.
	CarTable = "cars"
	// CarInverseTable is the table name for the Car entity.
	// It exists in this package in order to avoid circular dependency with the "car" package.
	CarInverseTable = "cars"
	// CarColumn is the table column denoting the car relation/edge.
	CarColumn = "user_id"
	// DeptTable is the table that holds the dept relation/edge. The primary key declared below.
	DeptTable = "dept_user"
	// DeptInverseTable is the table name for the Dept entity.
	// It exists in this package in order to avoid circular dependency with the "dept" package.
	DeptInverseTable = "depts"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldPhone,
	FieldRace,
}

var (
	// DeptPrimaryKey and DeptColumn2 are the table columns denoting the
	// primary key for the dept relation (M2M).
	DeptPrimaryKey = []string{"dept_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// Race defines the type for the "race" enum field.
type Race string

// Race values.
const (
	RaceYELLOW Race = "YELLOW"
	RaceWHITE  Race = "WHITE"
	RaceBLACK  Race = "BLACK"
)

func (r Race) String() string {
	return string(r)
}

// RaceValidator is a validator for the "race" field enum values. It is called by the builders before save.
func RaceValidator(r Race) error {
	switch r {
	case RaceYELLOW, RaceWHITE, RaceBLACK:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for race field: %q", r)
	}
}