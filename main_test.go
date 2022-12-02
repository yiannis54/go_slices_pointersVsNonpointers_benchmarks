package main

import (
	"testing"
	"time"

	"gorm.io/datatypes"
)

func BenchmarkSlicePointers(b *testing.B) {
	b.ReportAllocs()

	pointerList := make(PointerList, 1000)
	for i := range pointerList {
		pointerList[i] = &MyStruct{
			ID:         uint64(i),
			ExternalID: uint64(i),
			FirstName:  "John",
			LastName:   "Doe",
			Phone:      "123456789",
			Email:      "john@doe.com",
			Gender:     Male,
			StartDate:  "2022-01-01",
			EndDate:    "2022-01-01",
			Birthday:   datatypes.Date(time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local)),
			IsActive:   true,
			IsGuest:    false,
			Newsletter: true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
	}

	finalList := make([]uint64, len(pointerList))
	for _, dto := range pointerList {
		finalList = append(finalList, dto.toID())
	}
}

func BenchmarkSliceNoPointers(b *testing.B) {
	b.ReportAllocs()

	noPointerList := make(NoPointerList, 1000)
	for i := range noPointerList {
		noPointerList[i] = MyStruct{
			ID:         uint64(i),
			ExternalID: uint64(i),
			FirstName:  "John",
			LastName:   "Doe",
			Phone:      "123456789",
			Email:      "john@doe.com",
			Gender:     Male,
			StartDate:  "2022-01-01",
			EndDate:    "2022-01-01",
			Birthday:   datatypes.Date(time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local)),
			IsActive:   true,
			IsGuest:    false,
			Newsletter: true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
	}

	finalList := make([]uint64, len(noPointerList))
	for i := range noPointerList {
		finalList[i] = noPointerList[i].toID()
	}
}
