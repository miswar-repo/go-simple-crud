package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//main function
func main() {
	port := 8181

	router := mux.NewRouter()
	router.HandleFunc("/user", GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}

//Struct user
type userData struct {
	ID        string
	Firstname string
	Lastname  string
	City      string
	Country   string
}

type msgResponse struct {
	Status  string
	Message string
}

//Connection Database
func ConnDB() (*sql.DB, error) {

	//access database
	//Note : adjust your user and password access. for this code, we use user : root, password : password
        //Use IP MySQL Server, for this code we use 192.168.73.100 
	db, err := sql.Open("mysql", "root:password@tcp(192.168.73.100:3306)/cruddb")

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

//GetAllUser
func GetUsers(w http.ResponseWriter, r *http.Request) {

	var myUser userData
	// get connection database
	db, err := ConnDB()
	if err != nil {
		log.Panic(err)
	}

	// run query to get data all user
	rows, err := db.Query("select ID,Firstname,Lastname,City,Country from user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// looping over the rows to get data each row
	for rows.Next() {
		//Read the columns in each row into variables
		err := rows.Scan(&myUser.ID, &myUser.Firstname, &myUser.Lastname, &myUser.City, &myUser.Country)
		if err != nil {
			log.Fatal(err)
		}

		//encode to json format and send as  response
		json.NewEncoder(w).Encode(&myUser)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

//Get Specific User with id
func GetUser(w http.ResponseWriter, r *http.Request) {

	var myUser userData

	// get connection database
	db, err := ConnDB()
	if err != nil {
		log.Panic(err)
	}

	// get id from request http
	params := mux.Vars(r)
	var ID = params["id"]

	// run query to get data user with ID is params["id"]
	err = db.QueryRow("select ID,Firstname,Lastname,City,Country from user where ID = ?", ID).Scan(&myUser.ID, &myUser.Firstname, &myUser.Lastname, &myUser.City, &myUser.Country)
	if err != nil {
		log.Fatal(err)
	}

	//encode to json format and send as  response
	json.NewEncoder(w).Encode(&myUser)

}

//Create User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var myUser userData
	var msgR msgResponse
	// get connection database
	db, err := ConnDB()
	if err != nil {
		log.Panic(err)
	}

	//decode data from response body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&myUser)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	//prepare query to insert data
	stmt, err := db.Prepare("INSERT INTO user (ID,Firstname,Lastname,City,Country)VALUES(?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}

	//execute with parameter data
	_, err = stmt.Exec(&myUser.ID, &myUser.Firstname, &myUser.Lastname, &myUser.City, &myUser.Country)
	if err != nil {
		//log.Panic(err)
		msgR = msgResponse{Status: "Error", Message: err.Error()}
	} else {
		msgR = msgResponse{Status: "Success"}
	}

	//encode to json format and send status as  response
	json.NewEncoder(w).Encode(msgR)

}

//Delete User
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var msgR msgResponse
	// get connection database
	db, err := ConnDB()
	if err != nil {
		log.Panic(err)
	}

	// get id from request http
	params := mux.Vars(r)
	var ID = params["id"]

	//prepare query to delete data
	stmt, err := db.Prepare("DELETE FROM user where ID=?;")
	if err != nil {
		log.Fatal(err)
	}

	//execute with parameter data
	res, err := stmt.Exec(&ID)
	if err != nil {
		//log.Panic(err)
		msgR = msgResponse{Status: "Error", Message: err.Error()}
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	//encode to json format and send status as  response
	if rowCount > 0 {
		msgR = msgResponse{Status: "Success"}
	} else {
		msgR = msgResponse{Status: "Error", Message: "Not Found Data with ID " + ID}
	}

	json.NewEncoder(w).Encode(msgR)

}

//Update User
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var myUser userData
	var msgR msgResponse

	// get connection database
	db, err := ConnDB()
	if err != nil {
		log.Panic(err)
	}

	// get id from request http
	params := mux.Vars(r)
	var ID = params["id"]

	//decode data from response body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&myUser)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	//prepare query to update data
	stmt, err := db.Prepare("UPDATE user SET ID=?,Firstname=?,Lastname=?,City=?,Country=? where ID=?;")
	if err != nil {
		log.Fatal(err)
	}

	//execute with parameter data
	res, err := stmt.Exec(&myUser.ID, &myUser.Firstname, &myUser.Lastname, &myUser.City, &myUser.Country, &ID)
	if err != nil {
		//log.Panic(err)
		msgR = msgResponse{Status: "Error", Message: err.Error()}
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	//encode to json format and send status as  response
	if rowCount > 0 {
		msgR = msgResponse{Status: "Success"}
	} else {
		msgR = msgResponse{Status: "Error", Message: "Not Found Data with ID " + ID}
	}

	//encode to json format and send status as  response
	json.NewEncoder(w).Encode(msgR)

}

