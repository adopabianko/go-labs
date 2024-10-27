package main

import (
	"fmt"
	"log"

	"github.com/go-faker/faker/v4"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
	CREATE TABLE IF NOT EXISTS customers (
		id serial primary key not null,
		username varchar(30) not null,
		first_name varchar(30) not null,
		last_name varchar(30) not null,
		email varchar(100) not null,
		password varchar(100) not null,
		phone_number varchar(30) not null
	);
`

type Customer struct {
	UserName    string `db:"username" faker:"username"`
	FirstName   string `db:"first_name" faker:"first_name"`
	LastName    string `db:"last_name" faker:"last_name"`
	Email       string `db:"email" faker:"email"`
	Password    string `db:"password" faker:"password"`
	PhoneNumber string `db:"phone_number" faker:"phone_number"`
}

type config struct {
	DB *sqlx.DB
}

func main() {
	cfg := config{}
	cfg.dbConn()

	// create table customers if not exists in database
	cfg.DB.MustExec(schema)

	cfg.process()
}

func (c *config) process() {
	var (
		fakeCus   = Customer{}
		customers = []Customer{}
	)

	for i := 0; i < 1000; i++ {
		err := faker.FakeData(&fakeCus)
		if err != nil {
			log.Println(err)
		}

		customers = append(customers, Customer{
			UserName:    fakeCus.UserName,
			FirstName:   fakeCus.FirstName,
			LastName:    fakeCus.LastName,
			Email:       fakeCus.Email,
			Password:    fakeCus.Password,
			PhoneNumber: fakeCus.PhoneNumber,
		})
	}

	batch := 20

	for i := 0; i < len(customers); i += batch {
		j := i + batch
		if j > len(customers) {
			j = len(customers)
		}

		// insert batch data
		_, err := c.DB.NamedExec(`insert into customers (username, first_name, last_name, email, password, phone_number) values(:username, :first_name, :last_name, :email, :password, :phone_number)`,
			customers[i:j])
		if err != nil {
			log.Println(err)
		}

		log.Println(customers[i:j])
	}

}

func (c *config) dbConn() {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "faker-db", 5432, "faker", "faker", "faker_customers")
	db, err := sqlx.Connect("postgres", conn)

	if err != nil {
		log.Println(err)
	}

	c.DB = db
}
