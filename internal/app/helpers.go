package app

import (
	"fmt"
	"money/internal/models"
)

func workingWithTransaction (transaction *models.Transaction){
	fmt.Println("acc id", transaction.AccountID)
	fmt.Println("amount", transaction.Amount)
	fmt.Println("date", transaction.Date)
	fmt.Println("cat id", transaction.CategoryID)
	fmt.Println("comment", transaction.Comment)
}