package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"Project2/construct"

	"github.com/gorilla/mux"
)

func (handler *ApplicationHandler) UpdateUsers(rw http.ResponseWriter, r *http.Request) {
	/*w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")*/
	var uIns construct.Users

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)
	// convert the id type from string to int
	email := params["email"]

	uIns.FullName = r.FormValue("fullname")
	uIns.Password = r.FormValue("password")
	uIns.PromoCode = r.FormValue("promocode")

	// create an empty user of type models.User
	//var u_ins users
	// decode the json request to user
	/*err := json.NewDecoder(r.Body).Decode(&u_ins)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}*/
	print(uIns.FullName)
	val := ValidateUser(email)
	var response = construct.JSONResponse{}
	if email == "" || !val {
		response = construct.JSONResponse{Type: "error", Message: "You are missing user parameter or you have given the mail id that doesnt exist."}
	} else {
		/*if err != nil {
			log.Fatalf("Unable to decode the request body.  %v", err)
		}*/
		// call update user to update the user
		updateUser(email, uIns)
		// format the message string
		/*msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)
		// format the response message
		response = JsonResponse{
			Type:    "success",
			Message: msg}*/
		response = construct.JSONResponse{Type: "success", Message: "The user has been updated successfully!"}
	}
	// send the response
	json.NewEncoder(rw).Encode(response)

}
func updateUser(email string, uIns construct.Users) int64 {
	// create the postgres db connection
	db := setupDB()
	// close the db connection
	defer db.Close()
	// create the update sql query
	fmt.Print(uIns.FullName, uIns.PhoneNumber)
	sqlStatement := "UPDATE Signupinv SET full_name=$2, password=$3, promo_code=$4 WHERE email=$1;"
	res, err := db.Exec(sqlStatement, email, uIns.FullName, uIns.Password, uIns.PromoCode)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}
