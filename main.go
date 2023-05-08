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
	dbName 		= "laundry1"
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

	// --- OUM ---
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

	// --- Employee ---
	// CREATE
	newEmployee := Employee{Name: "John Doe"}
	insertEmployee(db, &newEmployee)

	// READ
	retrievedEmployee, err := getEmployee(db, newEmployee.Id)
	if err != nil {
        log.Fatal(err)
    }
	fmt.Println(retrievedEmployee)
	fmt.Println(getAllEmployees(db))

	// UPDATE
	updatedEmployee := Employee{Id: newEmployee.Id, Name: "Jane Doe"}
	updateEmployee(db, &updatedEmployee)

	// DELETE
	deleteEmployee(db, newEmployee.Id)

	// --- Customer ---
	// CREATE
    newCustomer := Customer{Name: "John Doe", PhoneNumber: "+1 123-456-7890"}
    insertCustomer(db, &newCustomer)

    // READ
    retrievedCustomer, err := getCustomer(db, newCustomer.Id)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(retrievedCustomer)
	fmt.Println(getAllCustomers(db))

    // UPDATE
    retrievedCustomer.Name = "Jane Doe"
    updateCustomer(db, retrievedCustomer)

    // DELETE
    deleteCustomer(db, retrievedCustomer.Id)
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

// --------------------- Employee Master -------------------
func insertEmployee(db *sql.DB, newEmployee *Employee) error {
	newId := uuid.New().String()
	stmt := `INSERT INTO public.employee (id,name) VALUES ($1,$2)`
	_, err := db.Exec(stmt, newId, newEmployee.Name)

	if err == nil {
		log.Printf("Employee %s added successfully", newEmployee.Name)
	} else {
		log.Println(err)
	}

	return err
}

func getEmployee(db *sql.DB, id string) (*Employee, error) {
	stmt := `SELECT name FROM public.employee WHERE id = $1`
	row := db.QueryRow(stmt, id)

	var employee Employee
	err := row.Scan(&employee.Name)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Employee with ID %s not found", id)
	} else if err != nil {
		return nil, err
	}

	employee.Id = id

	return &employee, nil
}

func getAllEmployees(db *sql.DB) ([]Employee, error) {
	stmt := `SELECT id, name FROM public.employee`
	rows, err := db.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := []Employee{}
	for rows.Next() {
		var employee Employee
		err := rows.Scan(&employee.Id, &employee.Name)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func updateEmployee(db *sql.DB, updatedEmployee *Employee) error {
	stmt := `UPDATE public.employee SET name = $2 WHERE id = $1`
	_, err := db.Exec(stmt, updatedEmployee.Id, updatedEmployee.Name)

	if err == nil {
		log.Printf("Employee with ID %s updated successfully", updatedEmployee.Id)
	} else {
		log.Println(err)
	}

	return err
}

func deleteEmployee(db *sql.DB, id string) error {
	stmt := `UPDATE employee SET is_deleted = true where id=$1`
	_, err := db.Exec(stmt, id)

	if err == nil {
		log.Printf("Employee with ID %s deleted successfully", id)
	} else {
		log.Println(err)
	}

	return err
}

// ----------------- Customer -------------------------
func insertCustomer(db *sql.DB, newCustomer *Customer) error {
	newId := uuid.New().String()
	stmt := `INSERT INTO public.customer (id,name,phone_number) VALUES ($1,$2,$3)`
	_, err := db.Exec(stmt, newId, newCustomer.Name, newCustomer.PhoneNumber)

	if err == nil {
		log.Printf("Customer %s added successfully", newCustomer.Name)
	} else {
		log.Println(err)
	}

	return err
}

func getCustomer(db *sql.DB, id string) (*Customer, error) {
	stmt := `SELECT name, phone_number FROM public.customer WHERE id = $1`
	row := db.QueryRow(stmt, id)

	var customer Customer
	err := row.Scan(&customer.Name, &customer.PhoneNumber)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Customer with ID %s not found", id)
	} else if err != nil {
		return nil, err
	}

	customer.Id = id

	return &customer, nil
}

func getAllCustomers(db *sql.DB) ([]Customer, error) {
	stmt := `SELECT id, name, phone_number FROM public.customer WHERE is_deleted = false`
	rows, err := db.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := []Customer{}
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func updateCustomer(db *sql.DB, updatedCustomer *Customer) error {
	stmt := `UPDATE public.customer SET name = $2, phone_number = $3 WHERE id = $1`
	_, err := db.Exec(stmt, updatedCustomer.Id, updatedCustomer.Name, updatedCustomer.PhoneNumber)

	if err == nil {
		log.Printf("Customer with ID %s updated successfully", updatedCustomer.Id)
	} else {
		log.Println(err)
	}

	return err
}

func deleteCustomer(db *sql.DB, id string) error {
	stmt := `UPDATE customer SET is_deleted = true where id=$1`
	_, err := db.Exec(stmt, id)

	if err == nil {
		log.Printf("Customer with ID %s deleted successfully", id)
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