package main

import (
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Car struct {
	ID string `json:"id"`
	Name string `json:"name"`
    Model string `json:"model"`
	Owner *Owner `json:"owner"`
}

type Owner struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var cars []Car

func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func getOneCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range cars {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Car{})
}

func createCar(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var car Car
    _ = json.NewDecoder(r.Body).Decode(&car)
    car.ID = strconv.Itoa(rand.Intn(1000000))
    cars = append(cars, car) 
    json.NewEncoder(w).Encode(car)
}


func updateCar(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range cars {
        if item.ID == params["id"] {
            cars = append(cars[:index], cars[index+1:]...)
            var car Car
            _ = json.NewDecoder(r.Body).Decode(&car)
            car.ID = params["id"]
            cars = append(cars, car) 
            json.NewEncoder(w).Encode(car)
            return
        }
    }
    json.NewEncoder(w).Encode(cars)
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range cars {
        if item.ID == params["id"] {
            cars = append(cars[:index], cars[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(cars)
}

func main() {
    r := mux.NewRouter()
    cars = append(cars, Car{ID: "1", Name: "Opel", Model: "Astra", Owner: &Owner{Firstname: "Egor", Lastname: "Glukov"}})
    cars = append(cars, Car{ID: "2", Name: "BMW", Model: "M6", Owner: &Owner{Firstname: "Georgy", Lastname: "Bobov"}})
    cars = append(cars, Car{ID: "3", Name: "Lada", Model: "Granta", Owner: &Owner{Firstname: "Vlad", Lastname: "Ruman"}})
 
    r.HandleFunc("/cars", getCars).Methods("GET")
    r.HandleFunc("/cars/{id}", getOneCar).Methods("GET")
    r.HandleFunc("/cars", createCar).Methods("POST")
    r.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
    r.HandleFunc("/cars/{id}", deleteCar).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", r))
}