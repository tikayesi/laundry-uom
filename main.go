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

	"github.com/google/uuid"
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

	// CREATE
	newOum := Uom {Name: "Test Unit"}
	insertUom(db, &newOum)

	// READ
	fmt.Println(getUom(db, "8ddefbbc-717f-423e-8d66-391a0ceedecb"))
	fmt.Println(getAllUom(db))
	
	// UPDATE
	updatedUom := Uom{Id: "8ddefbbc-717f-423e-8d66-391a0ceedecb", Name: "Test Unit2"}
	updateUom(db, &updatedUom)

	// DELETE
	deleteUom(db, "8ddefbbc-717f-423e-8d66-391a0ceedecb")
}

//=================== Create, Read, Update, Delete (CRUD) ========================

// ----------------- UOM Master -------------------------
func insertUom(db *sql.DB, newUom *Uom) error {
	newId := uuid.New().String()
	stmt := `INSERT INTO public.uom (id,name) VALUES ($1,$2)`
	_, err := db.Exec(stmt, newId, newUom.Name)

	if err == nil {
		log.Printf("UOM %s added successfully", newUom.Name)
	} else {
		log.Println(err)
	}

	return err
}

func getUom(db *sql.DB, id string) (*Uom, error) {
	stmt := `SELECT name FROM public.uom WHERE id = $1`
	row := db.QueryRow(stmt, id)

	var uom Uom
	err := row.Scan(&uom.Name)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("UOM with ID %s not found", id)
	} else if err != nil {
		return nil, err
	}

	uom.Id = id

	return &uom, nil
}

func getAllUom(db *sql.DB) ([]Uom, error) {
	stmt := `SELECT id, name FROM public.uom`
	rows, err := db.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	uoms := []Uom{}
	for rows.Next() {
		var uom Uom
		err := rows.Scan(&uom.Id, &uom.Name)
		if err != nil {
			return nil, err
		}
		uoms = append(uoms, uom)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return uoms, nil
}

func updateUom(db *sql.DB, updatedUom *Uom) error {
	stmt := `UPDATE public.uom SET name = $2 WHERE id = $1`
	_, err := db.Exec(stmt, updatedUom.Id, updatedUom.Name)

	if err == nil {
		log.Printf("UOM with ID %s updated successfully", updatedUom.Id)
	} else {
		log.Println(err)
	}

	return err
}

func deleteUom(db *sql.DB, id string) error {
	stmt := `UPDATE uom SET is_deleted = true where id=$1`
	_, err := db.Exec(stmt, id)

	if err == nil {
		log.Printf("UOM with ID %s deleted successfully", id)
	} else {
		log.Println(err)
	}

	return err
}


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