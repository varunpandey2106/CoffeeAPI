package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/varunpandey2106/CoffeeAPI/helpers"
	"github.com/varunpandey2106/CoffeeAPI/services"
	"github.com/go-chi/chi/v5"

)

var coffee services.Coffee

//GET COFFEEES

func GetAllCoffees(w http.ResponseWriter, r*http.Request){
	all, err :=coffee.GetAllCoffees()
	if err !=nil{
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffees": all})
}

// GET//coffees/coffee/{id}
func GetCoffeeById(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    coffee, err := coffee.GetCoffeeById(id)
    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
        return
    }
    helpers.WriteJSON(w, http.StatusOK, coffee)
}

func CreateCoffee(w http.ResponseWriter, r*http.Request){

	var coffeeData services.Coffee
	err:=json.NewDecoder(r.Body).Decode(&coffeeData)

	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
		return 
	}

	coffeeCreated , err:=coffee.CreateCoffee(coffeeData)

	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, coffeeCreated)

}