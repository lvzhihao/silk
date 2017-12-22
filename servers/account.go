package servers

import (
	"github.com/jinzhu/gorm"
	"github.com/lvzhihao/silk/models"
)

func CreateAccount(db *gorm.DB, account *models.Account) error {
	//todo check
	return account.Create(db)
}
