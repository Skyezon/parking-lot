package model

import (
	"fmt"
	"strconv"

	"github.com/skyezon/parking-lot/common/errors"
)

type ParkingLot struct {
	TotalLot int
	Lots     []Car
}

func (lot ParkingLot) Park(theCar Car) error {
	slot, err := lot.findNextFreeSlot()
	if err != nil {
		return errors.LogErr(err)
	}
	lot.Lots[slot] = theCar
	return nil
}

// absolute means real location in array, e.g : user input 4, then absolute is 3
func (lot ParkingLot) Leave(absoluteSlotNumber int) error {
	if err := lot.validateSlotNumber(absoluteSlotNumber); err != nil {
		return err
	}
	lot.Lots[absoluteSlotNumber] = Car{}
	return nil
}

func (lot ParkingLot) validateSlotNumber(absoluteSlotNumber int) error {
	if absoluteSlotNumber < 0 || absoluteSlotNumber >= lot.TotalLot {
		return errors.LogErr(fmt.Errorf("slot number invalid"), strconv.Itoa(absoluteSlotNumber))
	}
	return nil

}

func (lot ParkingLot) findNextFreeSlot() (int, error) {
	for i := 0; i < lot.TotalLot; i++ {
		if lot.Lots[i] != (Car{}) {
			return i, nil
		}
	}
	return 0, fmt.Errorf("No empty lot")
}

func NewParkingLot(totalLot int) (ParkingLot, error) {
	if totalLot <= 0 {
		return ParkingLot{}, fmt.Errorf("Maximum lot number is invalid")
	}

	return ParkingLot{
		TotalLot: totalLot,
		Lots:     make([]Car, 0, totalLot),
	}, nil
}
