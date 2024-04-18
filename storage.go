package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore() (*MySQLStore, error) {
	// Replace these details with your MySQL user, password, database name, and other connection details as necessary.
	connStr := "root:seq940@tcp(localhost:3306)/mysqldb?parseTime=true"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MySQLStore{
		db: db,
	}, nil
}

func (s *MySQLStore) CreateAccount(a *Account) error {
	_, err := s.db.Exec("INSERT INTO accounts (first_name, last_name, number, balance) VALUES (?, ?, ?, ?)", a.FirstName, a.LastName, a.Number, a.Balance)
	return err
}

func (s *MySQLStore) UpdateAccount(a *Account) error {
	//_, err := s.db.Exec("INSERT INTO accounts (first_name, last_name, number, balance) VALUES (?, ?, ?, ?)", a.FirstName, a.LastName, a.Number, a.Balance)
	return nil
}

func (s *MySQLStore) DeleteAccount(id int) error {
	//_, err := s.db.Exec("INSERT INTO accounts (first_name, last_name, number, balance) VALUES (?, ?, ?, ?)", a.FirstName, a.LastName, a.Number, a.Balance)
	return nil
}

func (s *MySQLStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
