package handlers

import (
	"Project2/construct"
	"encoding/json"
	"fmt"
	"net/http"
)

func (handler *ApplicationHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {

	db := setupDB()

	//printMessage("Gettxing users...")
	//log.Log.Infof("Getting users: %v", r)

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM Signupinv")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var details []construct.Users

	// Foreach movie
	for rows.Next() {
		var fullName string
		var email string
		var password string
		var num int
		var promoCode string
		var refer string

		err = rows.Scan(&fullName, &email, &password, &num, &promoCode, &refer)

		// check errors
		checkErr(err)

		details = append(details, construct.Users{FullName: fullName, Email: email, Password: password, PhoneNumber: num, PromoCode: promoCode, Reference: refer})
	}

	var response = construct.JSONResponse{Type: "success", Data: details}
	fmt.Println(response)

	json.NewEncoder(rw).Encode(response)
}
