package domain

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/samhaque1504106/banking_hexagonal/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/samhaque1504106/banking_hexagonal/errs"
	//"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	customers := []Customer{}
	//var rows *sql.Rows //after introducing sqlx to the project no need of rows
	var err error

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
		//rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "select customer_id,name,date_of_birth, city,zipcode,status from customers where status= ?"
		err = d.client.Select(&customers, findAllSql, status)
		//rows, err = d.client.Query(findAllSql, status)
	}

	//rows, err := d.client.Query(findAllSql)
	if err != nil {
		//log.Println("Error while querying " + err.Error())
		logger.Error("Error while querying ")
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// after introducing sqlx no need to use this boilerplate, sqlx does it with Select
	//for rows.Next() {
	//	var c Customer
	//	err = rows.Scan(&c.ID, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	//	if err != nil {
	//		//log.Println("Error while scanning " + err.Error())
	//		logger.Error("Error while scanning")
	//		return nil, errs.NewUnexpectedError("Unexpected database error")
	//	}
	//	customers = append(customers, c)
	//}
	return customers, nil
}
func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	var c Customer
	cuatomerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	//row := d.client.QueryRow(cuatomerSql, id)
	//err := row.Scan(&c.ID, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)  // sqlx reduces this 2 lines by Get
	err := d.client.Get(&c, cuatomerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			//log.Println("Error while querying by id " + err.Error())
			logger.Error("Error while querying by id")
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:1234@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// settings
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
