package store

import "money/internal/models"

type TransactionRepository struct{
	store *Store
}

func (tr *TransactionRepository) AddTransaction(transaction models.Transaction){
	stmt := `
	INSERT INTO transactions (
		date,
		account_id,
		category_id,
		amount,
		comment)
	VALUES (
		?,
		?,
		?,
		?,
		?)
	`
	_, err := tr.store.db.Exec(stmt,
									transaction.Date,
									transaction.AccountID,
									transaction.CategoryID,
									transaction.Amount,
									transaction.Comment)
	if err != nil{
		panic(err)
	}
}
