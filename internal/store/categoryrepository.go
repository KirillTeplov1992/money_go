package store

import "money/internal/models"

type CategoryRepository struct{
	store *Store
}

func (cr *CategoryRepository) GetExpenses() []*models.Category{
	stmt := `
	SELECT
		id,
		name
 	FROM 
		categories
	WHERE
		is_public and not type_of_category`
	
		res, err := cr.store.db.Query(stmt)
		if err != nil{
		panic(err)
		}

		var categoryList []*models.Category

		for res.Next(){
			category := &models.Category{}
			err = res.Scan(&category.ID,
							&category.Name)

			if err != nil{
				panic(err)
			}

			categoryList = append(categoryList, category)
		}

		return categoryList
}