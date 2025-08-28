package app

import (
	"encoding/json"
	"net/http"
	"strconv"
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


