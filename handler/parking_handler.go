package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/skyezon/parking-lot/service"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func ParkHandler(w http.ResponseWriter, r *http.Request) {
	regisNumber := chi.URLParam(r, "platNumber")
	color := chi.URLParam(r, "color")
	slot, err := service.ParkParkingLot(regisNumber, color)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(fmt.Sprintf("Allocated slot number: %d", slot)))

}

func LeaveHandler(w http.ResponseWriter, r *http.Request) {
	slot := chi.URLParam(r, "slotNumber")
	slotInt, err := strconv.Atoi(slot)
	if err != nil {
		w.Write([]byte("Slot number is invalid"))
		return
	}
	absoluteSlot := slotInt - 1
	err = service.LeaveParkingLot(absoluteSlot)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(fmt.Sprintf("Slot number %d is free", slotInt)))
}

func CreateParkingHandler(w http.ResponseWriter, r *http.Request) {
	totalLot := chi.URLParam(r, "totalParkingLot")
	totalLotInt, err := strconv.Atoi(totalLot)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	err = service.CreateParkingLot(totalLotInt)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(fmt.Sprintf("Created a parking lot with %s slots", totalLot)))
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
    res, err := service.StatusParkingLot()
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }
    w.Write([]byte(res))
}

func FindRegisNumberByColor(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("find regis number by color"))
}

func FindCarSlotsByColor(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("find slot by color"))
}

func FindSlotNumberbyRegisNumber(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("find slot number by regis number"))
}
