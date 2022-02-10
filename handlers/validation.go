package handlers

import "fmt"

func ValidateUser(email string) bool {
	db := setupDB()

	//printMessage("Getting users...")
	// Get all movies from movies table that don't have movieID = "1"
	var val bool
	err := db.QueryRow("select exists(select full_name from Signupinv where email = $1);", email).Scan(&val)
	//val = IsValid.Scan(&result)
	fmt.Print(email)
	fmt.Print(val)
	// check errors
	checkErr(err)
	return val

}
