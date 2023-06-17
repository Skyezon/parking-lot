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

func Test_LeaveParkingLot(t * testing.T){
    tests := []struct{
        name string
        absoluteSlotNumber int
        isErr bool
    }{
        {
            name: "valid",
            absoluteSlotNumber: 0,
            isErr: false,
        },
        {
            name: "invalid high absolute slot number",
            absoluteSlotNumber: 100,
            isErr: true,
        },
        {
            name: "invalid low absolute slot number",
            absoluteSlotNumber: -1,
            isErr: true,
        },
    }

    model.NewParkingLot(10)
    lot, _ := model.GetLotInstance()
    lot.Park(model.Car{})
    
    for _,tt := range tests{
        t.Run(tt.name,func(t *testing.T) {
            err := LeaveParkingLot(tt.absoluteSlotNumber)
            if (err != nil) != tt.isErr{
                t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE,tt.isErr,err)
            }
            
        })
    }

}

func Test_StatusParkingLot(t *testing.T){
    tests := []struct{
        name string
        isParkingLotCreated bool
        isErr bool
    }{
        {
            name: "valid",
            isParkingLotCreated: true,
            isErr: false,
        },
        {
            name: "invalid",
            isParkingLotCreated: false,
            isErr: true,
        },
    }

    for _,tt := range tests{
        t.Run(tt.name,func(t *testing.T) {
            if tt.isParkingLotCreated{
                model.NewParkingLot(10)
            }else{
                model.GlobalLot = nil
            }
            _, err := StatusParkingLot()
            if (err != nil) != tt.isErr{
                t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE,tt.isErr,err)
            }
        })
    }
}

