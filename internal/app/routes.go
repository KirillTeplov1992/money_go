package app

import "net/http"

func (app *Application) configureRouter() {
	app.router.HandleFunc("/", app.home)
	app.router.HandleFunc("/account", app.account)
	app.router.HandleFunc("/create_transaction", app.createTransaction)
	app.router.HandleFunc("/get_incoms", app.getIncoms)
	app.router.HandleFunc("/get_accounts", app.getAccounts)
	app.router.HandleFunc("/get_expenses", app.getExpenses)
	

	//подключаю CSS стили
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    app.router.Handle("/static/", http.StripPrefix("/static", fileServer))
}