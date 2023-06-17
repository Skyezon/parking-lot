package model

import (
	"fmt"
	"math"
	"strconv"

	"github.com/skyezon/parking-lot/common/errors"
)

var GlobalLot *ParkingLot

type ParkingLot struct {
	TotalLot int
	Lots     []Car
}

func GetLotInstance() (*ParkingLot,error){
    if GlobalLot != nil {
        return GlobalLot,nil
    }
    return &ParkingLot{}, errors.LogErr(fmt.Errorf("Please initialize Lot first"))
}

func NewParkingLot(totalLot int) (error) {
	if totalLot <= 0 || totalLot >= math.MaxInt32{
		return  errors.LogErr(fmt.Errorf("Maximum lot number is invalid"))
	}
    GlobalLot = &ParkingLot{
        TotalLot : totalLot,
        Lots: make([]Car,totalLot),
    }
    return nil
}

func (lot *ParkingLot)Park(theCar Car) (int,error) {
	slot, err := lot.findNextFreeSlot()
	if err != nil {
		return 0, errors.LogErr(err)
	}
	lot.Lots[slot] = theCar
	return slot, nil
}

// absolute means real location in array, e.g : user input 4, then absolute is 3
func (lot *ParkingLot) Leave(absoluteSlotNumber int) error {
	if err := lot.validateSlotNumber(absoluteSlotNumber); err != nil {
		return errors.LogErr(err) 
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
		if lot.Lots[i] == (Car{}) {
			return i, nil
		}
	}
	return 0, fmt.Errorf("No empty lot")
}

