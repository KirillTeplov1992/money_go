package app

import (
	"encoding/json"
	"fmt"
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

	cat_id, err := strconv.Atoi(r.FormValue("category"))
	if err != nil{
		panic(err)
	}

	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil{
		panic(err)
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
		fmt.Println("action", action)
		fmt.Println("acc id", transaction.AccountID)
		fmt.Println("amount", transaction.Amount )
		fmt.Println("type of category", type_of_category)
		fmt.Println("date", transaction.Date)
		fmt.Println("cat id", transaction.CategoryID)
		fmt.Println("comment", transaction.Comment)

		http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
	}

	if action == "add+" {
		fmt.Println("action", action)
		fmt.Println("acc id", transaction.AccountID)
		fmt.Println("amount", transaction.Amount )
		fmt.Println("type of category", type_of_category)
		fmt.Println("date", transaction.Date)
		fmt.Println("cat id", transaction.CategoryID)
		fmt.Println("comment", transaction.Comment)

		http.Redirect(w, r, fmt.Sprintf("/create_transaction"), http.StatusSeeOther)
	}

	

	
	
}


