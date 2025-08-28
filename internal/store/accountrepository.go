package store

import (
	"money/internal/models"
)

type AcoountRepository struct{
	store *Store
}

func (ar *AcoountRepository) GetAccounts() []*models.AccountData {
	stmt := `
	SELECT
		A.id,
		A.name,
    	SUM(amount)
	FROM
		transactions T
	INNER JOIN
		accounts A
	ON
		T.account_id = A.id
	GROUP BY
		A.name
	ORDER BY
		A.id`
	
	res, err := ar.store.db.Query(stmt)
	if err != nil{
		panic(err)
	}
	
	var accountList []*models.AccountData

	for res.Next(){
		account := &models.AccountData{}
		err = res.Scan(&account.ID,
						&account.Name,
						&account.Amount)
		if err != nil{
			panic(err)
		}

		accountList = append(accountList, account)
	}

	return accountList
}

func (ar *AcoountRepository) GetBalance()(*models.TotalBalance, error){
	stmt := `
	SELECT
		SUM(amount)
	FROM
		transactions`
	
	res := ar.store.db.QueryRow(stmt)
	
	balance := &models.TotalBalance{}

	err := res.Scan(&balance.Balance)
	if err != nil{
		return nil, err
	}

	return balance, nil
}

func (ar *AcoountRepository) GetAccountData (acc_id int) []*models.AccountTransaction{
	stmt := `
	SELECT
    	T.id,
    	T.date,
    	C.name,
    	T.amount
	FROM
		transactions T
	INNER JOIN
		categories C
	ON
		T.category_id = C.id
	WHERE
		account_id = ?
	ORDER BY
		T.date DESC
	`
	res, err := ar.store.db.Query(stmt, acc_id)
	if err != nil{
		panic(err)
	}

	var accDataList []*models.AccountTransaction

	for res.Next(){
		transaction := &models.AccountTransaction{}
		err = res.Scan(&transaction.ID,
						&transaction.Date,
						&transaction.Category,
						&transaction.Amount)
		if err != nil{
			panic(err)
		}

		accDataList = append(accDataList, transaction)
	}

	return accDataList
}

func (ar *AcoountRepository) GetAccountName(acc_id int) *models.Account{
	stmt :=`
	SELECT 
		name
	FROM
		accounts
	WHERE
		id = ?
	`
	res := ar.store.db.QueryRow(stmt, acc_id)
	
	account := &models.Account{}

	err := res.Scan(&account.Name)
	if err != nil{
		panic(err)
	}

	return account
}

func (ar *AcoountRepository) GetAccountsList() []*models.Account{
	stmt := `
	SELECT
		id,
		name
	FROM
		accounts`

	res, err := ar.store.db.Query(stmt)
	if err != nil{
		panic(err)
	}

	var accList []*models.Account

	for res.Next(){
		acc := &models.Account{}
		err = res.Scan(&acc.ID, &acc.Name)
		if err != nil{
			panic(err)
		}

		accList = append(accList, acc)
	}

	return accList

}