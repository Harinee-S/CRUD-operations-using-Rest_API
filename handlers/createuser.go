package handlers

import (
	"Project2/construct"
	"encoding/json"
	"fmt"
	"net/http"
)

func (handler *ApplicationHandler) CreateUsers(rw http.ResponseWriter, r *http.Request) {
	fullName := r.FormValue("fullname")
	email := r.FormValue("eemail")
	password := r.FormValue("pword")
	num := r.FormValue("number")
	promoCode := r.FormValue("code")
	reff := r.FormValue("referal")

	fmt.Println(fullName, email, password, num, promoCode, reff)

	var response = construct.JSONResponse{}

	if fullName == "" || email == "" || password == "" || num == "" || promoCode == "" || reff == "" {
		response = construct.JSONResponse{Type: "error", Message: "You are missing any of the parameter."}
	} else {
		db := setupDB()

		//printMessage("Creating users..")

		fmt.Println("Inserting new user with details: " + fullName + " with mail " + email + "with password" + password + "with phone number" + num + "promocode given by" + promoCode + "referred by" + reff)

		var lastInsertID string
		err := db.QueryRow("INSERT INTO Signupinv(full_name, email, password, num, promo_code, refer) VALUES($1, $2, $3, $4, $5, $6) returning full_name;", fullName, email, password, num, promoCode, reff).Scan(&lastInsertID)
		// check errors
		checkErr(err)

		response = construct.JSONResponse{Type: "success", Message: "The details has been inserted successfully!"}
	}

	json.NewEncoder(rw).Encode(response)
}
