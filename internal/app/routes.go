package app

import "net/http"

func (app *Application) configureRouter() {
	app.router.HandleFunc("/", app.home)
	app.router.HandleFunc("/account", app.account)
	app.router.HandleFunc("/create_transaction", app.createTransaction)
	app.router.HandleFunc("/add_categories", app.getCategories)
	

	//подключаю CSS стили
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    app.router.Handle("/static/", http.StripPrefix("/static", fileServer))
}