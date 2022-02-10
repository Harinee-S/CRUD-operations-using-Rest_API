package handlers

import (
	"encoding/json"
	"net/http"

	"Project2/construct"

	"github.com/gorilla/mux"
)

func (handler *ApplicationHandler) DeleteUsers(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	email := params["email"]

	var response = construct.JSONResponse{}

	val := ValidateUser(email)

	if email == "" || !val {
		response = construct.JSONResponse{Type: "error", Message: "You are missing user parameter or you have given the mail id that doesnt exist."}
	} else {
		//ValidateUser(email)
		db := setupDB()

		//printMessage("Deleting user from DB")

		_, err := db.Exec("DELETE FROM Signupinv where email = $1", email)

		// check errors
		checkErr(err)

		response = construct.JSONResponse{Type: "success", Message: "The user has been deleted successfully!"}
	}

	json.NewEncoder(rw).Encode(response)
}
