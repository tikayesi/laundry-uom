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
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	dbHost 		= "localhost"
	dbPort 		= "5432"
	dbName 		= "laundry2"
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
	runConsole()
}

// #region console menu
func runConsole() {
    db := connectDb()
    defer db.Close()
    mainMenuForm()
    for {
        var selectedMenu string
        fmt.Scanln(&selectedMenu)
        switch selectedMenu {
        case "1":
            uom := uomCreateForm()
            err := insertUom(db, &uom)
            checkErr(err)
            mainMenuForm()
        case "2":
            //delivery.ListProductForm(repo)
        case "3":
            //delivery.SearchProductForm(repo)
        case "4":
            //delivery.SearchProductForm(repo)
        case "5":
            billRequest := billCreateForm()
            fmt.Println(billRequest)
            err := createBill(db, &billRequest)
            checkErr(err)
            mainMenuForm()
        case "0":
            os.Exit(0)
        }
    }
}

func mainMenuForm() {
	fmt.Println(strings.Repeat("*", 30))
	fmt.Println("Enigma Laundry")
	fmt.Println(strings.Repeat("*", 30))
	fmt.Println("1. Master UOM")
	fmt.Println("2. Master Produk")
	fmt.Println("3. Master Staf")
	fmt.Println("4. Master Pelanggan")
	fmt.Println("5. Transaksi Baru")
	fmt.Println("0. Keluar")
	fmt.Println("Pilih Menu (0-5): ")
}

func uomCreateForm() Uom {
	var uomName string
	var saveUOMConfirmation string
	fmt.Print("UOM Name: ")
	fmt.Scanln(&uomName)
	fmt.Printf("UOM %s akan disimpan (y/t) ?", uomName)
	fmt.Scanln(&saveUOMConfirmation)

	if saveUOMConfirmation == "y" {
		var uom Uom
		uom.Name = uomName
		return uom
	}

	return Uom{}
}

func billDetailForm(billDetail *[]BillItemRequest) {
	for {
		var productId string
		var qty int
		var saveBillDetConfirmation string
		fmt.Println("Produk Id:")
		fmt.Scanln(&productId)
		fmt.Println("Jumlah:")
		fmt.Scanln(&qty)
		fmt.Print("simpan produk (y/t) ?")
		fmt.Scanln(&saveBillDetConfirmation)
		if saveBillDetConfirmation == "y" {
			*billDetail = append(*billDetail, BillItemRequest{
				ProductId: productId,
				Qty:       qty,
			})
		}
		var finishBillDetConfirmation string
		fmt.Print("selesai (y/t) ?")
		fmt.Scanln(&finishBillDetConfirmation)
		if finishBillDetConfirmation == "y" {
			fmt.Println(billDetail)
			break
		}
	}
}

func billCreateForm() BillRequest {
	var employeeId string
	var customerId string
	var billDetail []BillItemRequest
	var saveBillConfirmation string
	fmt.Print("Id Staff: ")
	fmt.Scanln(&employeeId)
	fmt.Print("Id Pelanggan: ")
	fmt.Scanln(&customerId)
	billDetailForm(&billDetail)
	fmt.Print("Buat Struk (y/t) ?")
	fmt.Scanln(&saveBillConfirmation)

	if saveBillConfirmation == "y" {
		var billReq BillRequest
		billReq.EmployeeId = employeeId
		billReq.CustomerId = customerId
		billReq.Items = billDetail
		return billReq
	}
	return BillRequest{}
}
// #endregion

// #region UOM
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

// #endregion

// #region Employee
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
// #endregion

// #region Customer
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
// #endregion

// #region Product
func insertProduct(db *sql.DB, newProduct *Product) error {
	newId := uuid.New().String()
	stmt := `INSERT INTO public.products (id,name,price,uom_id) VALUES ($1,$2,$3,$4)`
	_, err := db.Exec(stmt, newId, newProduct.Name, newProduct.Price, newProduct.Uom.Id)

	if err == nil {
		log.Printf("Product %s added successfully", newProduct.Name)
	} else {
		log.Println(err)
	}

	return err
}

func getProduct(db *sql.DB, id string) (*Product, error) {
	stmt := `SELECT p.name, p.price, u.id, u.name FROM public.products p INNER JOIN public.uom u ON p.uom_id = u.id WHERE p.id = $1`
	row := db.QueryRow(stmt, id)

	var product Product
	var uom Uom
	err := row.Scan(&product.Name, &product.Price, &uom.Id, &uom.Name)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Product with ID %s not found", id)
	} else if err != nil {
		return nil, err
	}

	product.Id = id
	product.Uom = uom

	return &product, nil
}

func getAllProducts(db *sql.DB) ([]Product, error) {
	stmt := `SELECT p.id, p.name, p.price, u.id, u.name FROM public.products p INNER JOIN public.uom u ON p.uom_id = u.id`
	rows, err := db.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var product Product
		var uom Uom
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &uom.Id, &uom.Name)
		if err != nil {
			return nil, err
		}
		product.Uom = uom
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func updateProduct(db *sql.DB, updatedProduct *Product) error {
	stmt := `UPDATE public.products SET name = $2, price = $3, uom_id = $4 WHERE id = $1`
	_, err := db.Exec(stmt, updatedProduct.Id, updatedProduct.Name, updatedProduct.Price, updatedProduct.Uom.Id)

	if err == nil {
		log.Printf("Product with ID %s updated successfully", updatedProduct.Id)
	} else {
		log.Println(err)
	}

	return err
}

func deleteProduct(db *sql.DB, id string) error {
	stmt := `UPDATE products SET is_deleted = true where id=$1`
	_, err := db.Exec(stmt, id)

	if err == nil {
		log.Printf("Product with ID %s deleted successfully", id)
	} else {
		log.Println(err)
	}

	return err
}
// #endregion

// #region transaction
func createBill(db *sql.DB, newBill *BillRequest) error {
	tx, err := db.Begin()
	checkErr(err)
	billStmt := `INSERT into bill (id,bill_date,finish_date, employee_id,customer_id) VALUES($1,$2,$3,$4,$5)`
	billId := uuid.New().String()
	billDate := time.Now()
	finishDate := billDate.AddDate(0, 0, 7)

	_, err = tx.Exec(billStmt, billId, billDate, finishDate, newBill.EmployeeId, newBill.CustomerId)
	validateTransaction(err, tx)

	billDetailStmt := `INSERT into bill_detail (id,bill_id,product_id, product_price,qty) VALUES($1,$2,$3,$4,$5)`
	for _, prodReq := range newBill.Items {
		billDetailId := uuid.New().String()
		product, _ := getProduct(db, prodReq.ProductId)
		_, err = tx.Exec(billDetailStmt, billDetailId, billId, prodReq.ProductId, product.Price, prodReq.Qty)
		validateTransaction(err, tx)
	}
	tx.Commit()
	return err
}

func validateTransaction(err error, tx *sql.Tx) {
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	}
}
// #endregion

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