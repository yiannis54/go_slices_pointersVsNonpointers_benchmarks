package main

import (
	"time"

	"gorm.io/datatypes"
)

func main() {}

type Gender int

const (
	Female Gender = iota + 1
	Male
)

type MyStruct struct {
	ID         uint64
	ExternalID uint64
	FirstName  string
	LastName   string
	Phone      string
	Email      string
	Gender     Gender
	StartDate  string
	EndDate    string
	Birthday   datatypes.Date
	IsActive   bool
	IsGuest    bool
	Newsletter bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type PointerList []*MyStruct
type NoPointerList []MyStruct

func (dto *MyStruct) toID() uint64 {
	return dto.ID
}
