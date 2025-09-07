package app

import (
	"money/internal/models"
	"html/template"
	"path/filepath"
	
)

type templateData struct {
	AccountsListWithBalance []*models.AccountData
	TotalBalance *models.TotalBalance
	TransactionList []*models.AccountTransaction
	Account *models.Account
	AccountList []*models.Account
	CategoryList []*models.Category
	Transaction *models.Transaction
	Errors []string
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	
	pages, err := filepath.Glob(filepath.Join(dir, "*page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		
		ts, err = ts.ParseGlob(filepath.Join(dir, "*layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*partail.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
