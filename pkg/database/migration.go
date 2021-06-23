package database

import "full-stack-go/models"

func MigrationTables() error {
	return DBInstance.AutoMigrate(
		&models.Bank{},
		&models.Policy{},
		&models.ContractType{},
		&models.Contract{},
		&models.Email{},
		&models.ContractTimeline{},
		&models.Position{},
		&models.InterestRate{},
		&models.User{},
		&models.Customer{},
		)
}
