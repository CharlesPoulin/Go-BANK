package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
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

func (s *MySQLStore) init() error {
	return s.CreateAccountTable()
	// drop table, migration ect.
}

// IDONTKNOW HOW TO USE ORM AND I NEED TO STUDY MYSQL FOR MY EXAM
func (s *MySQLStore) CreateAccountTable() error {
	query := `
    CREATE TABLE IF NOT EXISTS accounts (
        id INT AUTO_INCREMENT PRIMARY KEY,
        first_name VARCHAR(50),
        last_name VARCHAR(50),
        number BIGINT,
        balance BIGINT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    )`

	_, err := s.db.Exec(query)
	if err != nil {
		// Optionally log the error here as well
		return fmt.Errorf("error creating accounts table: %w", err)
	}
	return nil
}

func (s *MySQLStore) CreateAccount(a *Account) error {
	query := "INSERT INTO accounts (first_name, last_name, number, balance) VALUES (?, ?, ?, ?)"
	resp, err := s.db.Exec(query, a.FirstName, a.LastName, a.Number, a.Balance)
	if err != nil {
		log.Printf("Failed to create account: %v", err) // Assuming 'log' is properly initialized
		return fmt.Errorf("error creating account: %w", err)
	}

	id, err := resp.LastInsertId()
	if err != nil {
		log.Printf("Failed to retrieve last insert ID: %v", err)
		return fmt.Errorf("error getting last insert ID: %w", err)
	}

	fmt.Printf("Account created with ID: %d\n", id)
	return nil
}

func (s *MySQLStore) UpdateAccount(a *Account) error {
	//_, err := s.db.Exec("INSERT INTO accounts (first_name, last_name, number, balance) VALUES (?, ?, ?, ?)", a.FirstName, a.LastName, a.Number, a.Balance)
	return nil
}

func (s *MySQLStore) DeleteAccount(id int) error {

	//todo find the Cleanway to do it for prod (this looks barbarian)
	_, err := s.db.Query("DELETE FROM accounts WHERE id = ?", id)
	return err
}

func (s *MySQLStore) GetAccountByID(id int) (*Account, error) {
	rows, err := s.db.Query("SELECT id, first_name, last_name, number, balance, created_at, updated_at FROM accounts WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

// implement page limit
func (s *MySQLStore) GetAccounts() ([]*Account, error) {
	// Query the database for all accounts
	rows, err := s.db.Query("SELECT id, first_name, last_name, number, balance, created_at, updated_at FROM accounts")
	if err != nil {
		return nil, fmt.Errorf("error querying accounts: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var accounts []*Account

	for rows.Next() {
		account := new(Account)
		if err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt, &account.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning account: %w", err)
		}
		account, _ = scanIntoAccount(rows)
		accounts = append(accounts, account)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return accounts, nil
}

func scanIntoAccount(row *sql.Rows) (*Account, error) {
	account := new(Account)
	if err := row.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt, &account.UpdatedAt); err != nil {
		return nil, fmt.Errorf("error scanning account: %w", err)
	}
	return account, nil
}
