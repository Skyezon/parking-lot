package model

import (
	"math"
	"testing"

	"github.com/skyezon/parking-lot/common/errors"
)

func Test_NewParkingLot(t *testing.T) {
	tests := []struct {
		name     string
		totalLot int
		isErr    bool
	}{
		{
			name:     "valid",
			totalLot: 1,
			isErr:    false,
		},
		{
			name:     "err too high",
			totalLot: math.MaxInt32 + 1,
			isErr:    true,
		},
		{
			name:     "err zero",
			totalLot: 0,
			isErr:    true,
		},
		{
			name:     "err too low",
			totalLot: -1,
			isErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewParkingLot(tt.totalLot)
			if (err != nil) != tt.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.isErr, err)
			}
		})
	}
}

func Test_validateSlotNumber(t *testing.T) {
	tests := []struct {
		name               string
		absoluteSlotNumber int
		maxTotalSlot       int
		isErr              bool
	}{
		{
			name:               "valid",
			absoluteSlotNumber: 5,
			maxTotalSlot:       10,
			isErr:              false,
		},
		{
			name:               "err over",
			absoluteSlotNumber: 11,
			maxTotalSlot:       10,
			isErr:              true,
		},
		{
			name:               "err too low",
			absoluteSlotNumber: -1,
			maxTotalSlot:       10,
			isErr:              true,
		},
		{
			name:               "err same as size",
			absoluteSlotNumber: 10,
			maxTotalSlot:       10,
			isErr:              true,
		},
		{
			name:               "valid first",
			absoluteSlotNumber: 0,
			maxTotalSlot:       10,
			isErr:              false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewParkingLot(tt.maxTotalSlot)
			lot, _ := GetLotInstance()
			err := lot.validateSlotNumber(tt.absoluteSlotNumber)
			if (err != nil) != tt.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.isErr, err)
			}

		})
	}
}

func Test_findNextFreeSlot(t *testing.T) {
	tests := []struct {
		name  string
		lot   ParkingLot
		isErr bool
	}{
		{
			name: "valid",
			lot: ParkingLot{
				TotalLot: 10,
				Lots:     make([]Car, 10),
			},
			isErr: false,
		},
		{
			name: "valid pointer",
			lot: *&ParkingLot{
				TotalLot: 10,
				Lots:     make([]Car, 10),
			},
			isErr: false,
		},
		{
			name:  "invalid empty lot",
			lot:   ParkingLot{},
			isErr: true,
		},
		{
			name: "invalid a lot full",
			lot: ParkingLot{
				TotalLot: 2,
				Lots:     []Car{{"red", "bk-1234-abc"}, {"blue", "bk-1234-cde"}},
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.lot.findNextFreeSlot()
			if (err != nil) != tt.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.isErr, err)
			}
		})
	}
}

func Test_Park(t *testing.T) {
	tests := []struct {
		name  string
		lot   ParkingLot
		isErr bool
	}{
		{
			name: "valid",
			lot: ParkingLot{
				TotalLot: 10,
				Lots:     make([]Car, 10),
			},
			isErr: false,
		},
		{
			name:  "invalid",
			lot:   ParkingLot{},
			isErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.lot.Park(Car{"abc", "abc"})
			if (err != nil) != tt.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.isErr, err)
			}
		})
	}
}

func Test_Leave(t *testing.T) {
	tests := []struct {
		name         string
		lot          ParkingLot
		isErr        bool
		absoluteSlot int
	}{
		{
			name: "valid",
			lot: ParkingLot{
				10,
				[]Car{{"Red", "bk-1235-abc"}},
			},
			absoluteSlot: 0,
			isErr:        false,
		},
		{
			name:         "invalid",
			lot:          ParkingLot{},
			absoluteSlot: 1,
			isErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.lot.Leave(tt.absoluteSlot)
			if (err != nil) != tt.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.isErr, err)
			}
		})
	}
}

func Test_GetLotInstance(t *testing.T) {
	tests := []struct {
		name         string
		isLotCreated bool
		isErr        bool
	}{
		{
			name:         "invalid",
			isLotCreated: false,
			isErr:        true,
		},
		{
			name:         "valid",
			isLotCreated: true,
			isErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isLotCreated {
				NewParkingLot(10)
			} else {
				GlobalLot = nil
			}
			_, err := GetLotInstance()
			if (err != nil) != tt.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.isErr, err)
			}
		})
	}
}


