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
	w.Write([]byte("ngepark"))
}

func LeaveHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ngeleave"))
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
	w.Write([]byte("get status"))
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
