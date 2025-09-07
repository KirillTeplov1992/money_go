package app

import (
	"encoding/json"
	"money/internal/models"
	"net/http"
	"strconv"
	"time"
)

func (app *Application) home (w http.ResponseWriter, r *http.Request) {
	accountList := app.store.Account().GetAccounts()

	balance, err := app.store.Account().GetBalance()
	if err != nil{
		panic(err)
	}
	
	app.render(w, r, "home.page.tmpl", &templateData{
		AccountsListWithBalance : accountList,
		TotalBalance: balance,
	})
}

func (app *Application) account (w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	transactionsList := app.store.Account().GetAccountData(id)
	account := app.store.Account().GetAccountName(id)

	app.render(w, r, "account.page.tmpl", &templateData{
		TransactionList: transactionsList,
		Account: account,
	})
}

// createTransaction
// Функция создает окно, в которое вводятся данные о транзакции
func (app *Application) createTransaction (w http.ResponseWriter, r *http.Request){
	accList := app.store.Account().GetAccountsList()
	catList := app.store.Category().GetExpenses()

	app.render(w, r, "create_transaction.page.tmpl", &templateData{
		AccountList: accList,
		CategoryList: catList,
	})
}

func (app *Application) getIncoms (w http.ResponseWriter, r *http.Request){
	incomList := app.store.Category().GetIncoms()

	w.Header().Set("incoms list", "application/json")

	json.NewEncoder(w).Encode(incomList)
}

func (app *Application) getAccounts (w http.ResponseWriter, r *http.Request){
	accList := app.store.Account().GetAccountsList()

	w.Header().Set("incoms list", "application/json")

	json.NewEncoder(w).Encode(accList)
}

func (app *Application) getExpenses (w http.ResponseWriter, r *http.Request){
	expensesList := app.store.Category().GetExpenses()

	w.Header().Set("incoms list", "application/json")

	json.NewEncoder(w).Encode(expensesList)
}

// addTransaction
// Функция добавляет транзакцию в базу, в зависимости от её типа
func (app *Application) addTransaction (w http.ResponseWriter, r *http.Request) {
	action := r.FormValue("action")
	type_of_category := r.FormValue("type_of_category")

	date, err := time.Parse("2006-01-02" ,r.FormValue("date"))
	if err != nil {
		panic(err)
	}

	acc_id, err := strconv.Atoi(r.FormValue("account"))
	if err != nil {
		panic(err)
	}

	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil{
		panic(err)
	}

	if type_of_category != "Перевод" {
		cat_id, err := strconv.Atoi(r.FormValue("category"))
		if err != nil{
			panic(err)
		}	

		if type_of_category == "Расход" {
			amount = amount*(-1)
		}
	

		//var errList []string

		transaction := models.Transaction {
			Date: date,
			AccountID: acc_id,
			CategoryID: cat_id,
			Amount: amount,
			Comment: r.FormValue("comment"),
		}
	

		if action == "add" {
			app.store.Transaction().AddTransaction(transaction)
			http.Redirect(w, r,"/", http.StatusSeeOther)
		}

		if action == "add+" {
			app.store.Transaction().AddTransaction(transaction)
			http.Redirect(w, r, "/create_transaction", http.StatusSeeOther)
		}
	} else {
		first_acc_id := acc_id
		second_acc_id, err := strconv.Atoi(r.FormValue("category"))
		if err != nil{
			panic(err)
		}

		transactionMinus := models.Transaction {
			Date: date,
			AccountID: first_acc_id,
			CategoryID: 2,
			Amount: amount*(-1),
			Comment: r.FormValue("comment"),
		}

		transactionPlus := models.Transaction {
			Date: date,
			AccountID: second_acc_id,
			CategoryID: 1,
			Amount: amount,
			Comment: r.FormValue("comment"),
		}

		app.store.Transaction().AddTransaction(transactionMinus)
		app.store.Transaction().AddTransaction(transactionPlus)

		if action == "add"{
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		if action == "add+"{
			http.Redirect(w, r, "/create_transaction", http.StatusSeeOther)
		}
	}	
}

func (app *Application) getTransaction (w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	transaction, err := app.store.Transaction().GetTransaction(id)
	if err != nil{
		panic(err)
	}

	category := app.store.Category().GetTypeOfCategory(transaction.CategoryID)
	accList := app.store.Account().GetAccountsList()

	if category.Type_of_category {
		catList := app.store.Category().GetIncoms()
		
		app.render(w, r, "get_transaction.page.tmpl", &templateData{
		AccountList: accList,
		CategoryList: catList,
		Transaction: transaction,
		})
	} else {
		catList := app.store.Category().GetExpenses()

		transaction.Amount = transaction.Amount*(-1)
	
		app.render(w, r, "get_transaction.page.tmpl", &templateData{
		AccountList: accList,
		CategoryList: catList,
		Transaction: transaction,
		})
	}

	




	
	
}


