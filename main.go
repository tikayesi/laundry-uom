/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Mon May 08 2023 11:32:12 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbHost 		= "localhost"
	dbPort 		= "5432"
	dbName 		= "enigma_laundry"
	dbUser 		= "postgres"
	dbPassword 	= "postgres"
	sslMode 	= "disable"
)

type Customer struct {
	Id          string
	Name        string
	PhoneNumber string
}

type Employee struct {
	Id   string
	Name string
}

type Uom struct {
	Id   string
	Name string
}

type Product struct {
	Id    string
	Name  string
	Price int
	Uom   Uom
}

type BillRequest struct {
	EmployeeId string
	CustomerId string
	Items      []BillItemRequest
}

type BillItemRequest struct {
	ProductId string
	Qty       int
}

func main() {
	db := connectDb()

	// test CRUD
	newOum := Uom {Name: "Test Unit"}
	insertUom(db, &newOum)
}

//=================== Create, Read, Update, Delete (CRUD) ========================

// ----------------- UOM Master -------------------------
func insertUom(db *sql.DB, newUom *Uom) {
	stmt := `INSERT INTO uom (id, name) VALUES ($1,$2)`
	_, err := db.Exec(stmt, newUom.Id, newUom.Name)
	checkErr(err)
	log.Println("UOM added successfully")
}

// func deleteUom(db *sql.DB, id string) error {
// 	stmt := `UPDATE uom SET is_delete = true WHERE id=$1`
// 	_, err := db.Exec(stmt, id)
// 	return err
// }

// func findUomById(db *sql.DB, id string) (Uom, error) {
// 	stmt := `SELECT id, name FROM uom WHERE id=$1`
// 	row := db.QueryRow(stmt, id)
// 	var uom Uom
// 	switch err := row.Scan(&uom.Id, &uom.Name); err {
// 	case sql.ErrNoRows:
// 		return Uom{}, err
// 	case nil:
// 		return uom, nil
// 	default:
// 		panic(err)
// 	}
// }

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func connectDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)
	db, err := sql.Open("postgres", psqlInfo)
	checkErr(err)
	fmt.Println("connected to db")
	return db
}