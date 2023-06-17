package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/skyezon/parking-lot/service"
)

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
		
		w.Write([]byte("slot number invalid"))
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
		
		w.Write([]byte("Maximum lot number is invalid"))
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

func FindRegisNumberByColorHandler(w http.ResponseWriter, r *http.Request) {
	color := chi.URLParam(r, "color")
	res, err := service.GetRegisNumberByColor(color)
	if err != nil {
		
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(res))
}

func FindCarSlotsByColorHandler(w http.ResponseWriter, r *http.Request) {
	color := chi.URLParam(r, "color")
	res, err := service.GetSlotByColor(color)
	if err != nil {
		
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(res))
}

func FindSlotNumberbyRegisNumberHandler(w http.ResponseWriter, r *http.Request) {
	regisNumber := chi.URLParam(r, "regisNumber")
	res, err := service.GetSlotByRegisNum(regisNumber)
	if err != nil {
		
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(res))
}

func BulkCommandHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		
		w.Write([]byte(err.Error()))
		return
	}
	payload := string(body)
	res := service.BulkCommander(payload)

	if err != nil {
	
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(res))
}
