package models

import "app/config"

func Migrate() {
	err := config.DB.AutoMigrate(
		&Payment{},
		&PaymentLog{},
	)

	if err != nil{
		panic("Gagal migrate table" + err.Error())
	}
}