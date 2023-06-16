package handler

import "net/http"

func HelloWorldHandler(w http.ResponseWriter,r *http.Request){
    w.Write([]byte("Hello world"))
}

func ParkHandler(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("ngepark"))
}

func LeaveHandler(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("ngeleave"))
}

func CreateParkingHandler(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("create parking"))
}

func StatusHandler(w http.ResponseWriter,r *http.Request){
    w.Write([]byte("get status"))
}

func FindRegisNumberByColor(w http.ResponseWriter, r * http.Request ){
    w.Write([]byte("find regis number by color"))
} 

func FindCarSlotsByColor(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("find slot by color"))
}

func FindSlotNumberbyRegisNumber(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("find slot number by regis number"))
}
