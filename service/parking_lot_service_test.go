package service

import (
	"testing"

	"github.com/skyezon/parking-lot/common/errors"
	"github.com/skyezon/parking-lot/db/model"
)

func Test_CreateParkingLot(t *testing.T) {
	tests := []struct {
		name     string
		totalLot int
		isErr    bool
	}{
		{
			name:     "valid",
			totalLot: 10,
			isErr:    false,
		},
		{
			name:     "invalid",
			totalLot: -1,
			isErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateParkingLot(tt.totalLot); (err != nil) != tt.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.isErr, err)
			}
		})

	}
}

func Test_ParkParkingLot(t *testing.T) {
	tests := []struct {
		name        string
		regisNumber string
		color       string
		isErr       bool
	}{
		{
			name:        "valid",
			regisNumber: "BK-5432-abc",
			color:       "Red",
			isErr:       false,
		},
		{
			name:        "invalid",
			regisNumber: "",
			color:       "",
			isErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model.NewParkingLot(10)
			_, err := ParkParkingLot(tt.regisNumber, tt.color)
			if (err != nil) != tt.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.isErr, err)
			}
		})
	}
}

func Test_LeaveParkingLot(t *testing.T) {
	tests := []struct {
		name               string
		absoluteSlotNumber int
		isErr              bool
	}{
		{
			name:               "valid",
			absoluteSlotNumber: 0,
			isErr:              false,
		},
		{
			name:               "invalid high absolute slot number",
			absoluteSlotNumber: 100,
			isErr:              true,
		},
		{
			name:               "invalid low absolute slot number",
			absoluteSlotNumber: -1,
			isErr:              true,
		},
	}

	model.NewParkingLot(10)
	lot, _ := model.GetLotInstance()
	lot.Park(model.Car{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := LeaveParkingLot(tt.absoluteSlotNumber)
			if (err != nil) != tt.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.isErr, err)
			}

		})
	}

}

func Test_StatusParkingLot(t *testing.T) {
	tests := []struct {
		name                string
		isParkingLotCreated bool
		isErr               bool
	}{
		{
			name:                "valid",
			isParkingLotCreated: true,
			isErr:               false,
		},
		{
			name:                "invalid",
			isParkingLotCreated: false,
			isErr:               true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isParkingLotCreated {
				model.NewParkingLot(10)
			} else {
				model.GlobalLot = nil
			}
			_, err := StatusParkingLot()
			if (err != nil) != tt.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.isErr, err)
			}
		})
	}
}

func Test_GetRegisNumberByColor(t *testing.T) {
	tests := []struct {
		name                string
		isParkingLotCreated bool
		isErr               bool
	}{
		{
			name:                "valid",
			isParkingLotCreated: true,
			isErr:               false,
		},
		{
			name:                "invalid",
			isParkingLotCreated: false,
			isErr:               true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isParkingLotCreated {
				model.NewParkingLot(10)
			} else {
				model.GlobalLot = nil
			}
			_, err := GetRegisNumberByColor("abc")
			if (err != nil) != tt.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.isErr, err)
			}
		})
	}
}

func Test_GetSlotByColor(t *testing.T) {
	tests := []struct {
		name                string
		isParkingLotCreated bool
		isErr               bool
	}{
		{
			name:                "valid",
			isParkingLotCreated: true,
			isErr:               false,
		},
		{
			name:                "invalid",
			isParkingLotCreated: false,
			isErr:               true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isParkingLotCreated {
				model.NewParkingLot(10)
			} else {
				model.GlobalLot = nil
			}
			_, err := GetSlotByColor("abc")
			if (err != nil) != tt.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.isErr, err)
			}
		})
	}
}

func Test_GetSlotByRegisNum(t *testing.T) {
	tests := []struct {
		name                string
		isParkingLotCreated bool
		isErr               bool
		isFilled            bool
	}{
		{
			name:                "error not found",
			isParkingLotCreated: true,
			isErr:               true,
			isFilled:            false,
		},
		{
			name:                "invalid",
			isParkingLotCreated: false,
			isErr:               true,
			isFilled:            false,
		},
		{
			name:                "valid",
			isParkingLotCreated: true,
			isErr:               false,
			isFilled:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isParkingLotCreated {
				model.NewParkingLot(10)
			} else {
				model.GlobalLot = nil
			}
			if tt.isFilled {
				lot, _ := model.GetLotInstance()
				lot.Park(model.Car{Color: "blue", RegisNumber: "bk-1234-abc"})

			}
			_, err := GetSlotByRegisNum("BK-1234-ABC")
			if (err != nil) != tt.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.isErr, err)
			}
		})
	}
}

func Test_BulkCommander(t *testing.T) {
	tests := []struct {
		name          string
		payload       string
		expectedValue string
	}{
		{
			name: "valid",
			payload: `create_parking_lot 6
park B-1234-RFS Black
park B-1999-RFD Green
park B-1000-RFS Black
park B-1777-RFU BlueSky
park B-1701-RFL Blue
park B-1141-RFS Black
leave 4
status
park B-1333-RFS Black
park B-1989-RFU BlueSky
registration_numbers_for_cars_with_colour Black
slot_numbers_for_cars_with_colour Black
slot_number_for_registration_number B-1701-RFL
slot_number_for_registration_number RI-1
`,
			expectedValue: `Created a parking lot with 6 slots
Allocated slot number: 1
Allocated slot number: 2
Allocated slot number: 3
Allocated slot number: 4
Allocated slot number: 5
Allocated slot number: 6
Slot number 4 is free
Slot No. Registration No Colour
1 B-1234-RFS Black
2 B-1999-RFD Green
3 B-1000-RFS Black
5 B-1701-RFL Blue
6 B-1141-RFS Black
Allocated slot number: 4
Sorry, parking lot is full
B-1234-RFS, B-1000-RFS, B-1333-RFS, B-1141-RFS
1, 3, 4, 6
5
Not found
`,
		},
		{
			name: "invalid",
			payload: `create_parking_lot 6
park B-1234-RFS Black
park B-1999-RFDb Green`,
			expectedValue: `Created a parking lot with 6 slots
Allocated slot number: 1
Registration number is invalid`,
		},
		{
			name: " error insufficient param",
			payload: `create_parking_lot 6
park B-1234-RFS 
leave 
status
park B-1333-RFS 
park 1989-RFU BlueSky
registration_numbers_for_cars_with_colour 
slot_numbers_for_cars_with_colour 
slot_number_for_registration_number 1701-RFL
slot_number_for_registration_number 
`,
			expectedValue: `Created a parking lot with 6 slots
color cannot be empty
invalid slot number
Slot No. Registration No Colour

color cannot be empty
Registration number is invalid
color cannot be empty
1, 2, 3, 4, 5, 6
Not found
Not found
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := BulkCommander(tt.payload)
			if res != tt.expectedValue {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.expectedValue, res)
			}
		})
	}
}

func Test_executeCommand(t *testing.T) {
	tests := []struct {
		name          string
		oneCommand    string
		isErr         bool
		expectedValue string
	}{
		{
			name:          "valid",
			oneCommand:    "create_parking_lot 6",
			isErr:         false,
			expectedValue: "Created a parking lot with 6 slots",
		},
		{
			name:          "invalid",
			oneCommand:    "create_parking_lot -1",
			isErr:         true,
			expectedValue: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := executeCommand(tt.oneCommand)
			if (err != nil) != tt.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.isErr, err)
			}
			if err == nil && res != tt.expectedValue {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.expectedValue, res)
			}
		})
	}
}
