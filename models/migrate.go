package models

import "github.com/seccret404/BE-payment-gateway-Go-Fiber/config"

func Migrate() {
	err := config.DB.AutoMigrate(
		&Payment{},
		&PaymentLog{},
	)

	if err != nil{
		panic("Gagal migrate table" + err.Error())
	}
}