package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	config *Config
	db *sql.DB
	accountRepository *AcoountRepository
	categoryRepository *CategoryRepository
	transactionRepository *TransactionRepository
}

func New(config *Config) *Store{
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error{
	db, err:= sql.Open("mysql", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
        return err
	}

	s.db = db

	return nil
}

func (s *Store) Close(){
	s.db.Close()
}

func (s *Store) Account() *AcoountRepository{
	if s.accountRepository != nil{
		return s.accountRepository
	}

	s.accountRepository = &AcoountRepository{
		store: s,
	}

	return s.accountRepository
}

func (s *Store) Category() *CategoryRepository{
	if s.categoryRepository != nil{
		return s.categoryRepository
	}

	s.categoryRepository = &CategoryRepository{
		store: s,
	}

	return s.categoryRepository
}

func (s *Store) Transaction() *TransactionRepository{
	if s.transactionRepository != nil{
		return s.transactionRepository
	}

	s.transactionRepository = &TransactionRepository{
		store: s,
	}

	return s.transactionRepository
}

