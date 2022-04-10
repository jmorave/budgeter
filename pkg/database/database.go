package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Bill struct {
	Amount_actual   float64   `gorm:"column:amount_actual;not null;", json:amount_actual`
	Amount_budgeted float64   `gorm:"column:amount_budgeted;not null;"`
	Date_paid       time.Time `gorm:"column:date_paid;not null;"`
	Name            string    `gorm:"column:name;size:128;not null;"`
	Due_date        time.Time `gorm:"column:due_date;not null;"`
}

type NatashaExpense struct {
	Amount_actual float64   `gorm:"column:amount_actual;not null;"`
	Date_paid     time.Time `gorm:"column:date_paid;not null;"`
	Name          string    `gorm:"column:name;size:128;not null;"`
}

type JustinExpense struct {
	Amount_actual float64   `gorm:"column:amount_actual;not null;"`
	Date_paid     time.Time `gorm:"column:date_paid;not null;"`
	Name          string    `gorm:"column:name;size:128;not null;"`
}

type FuelExpense struct {
	Amount_actual float64   `gorm:"column:amount_actual;not null;"`
	Date_paid     time.Time `gorm:"column:date_paid;not null;"`
	Name          string    `gorm:"column:name;size:128;not null;"`
}

type HygeineExpense struct {
	Amount_actual float64   `gorm:"column:amount_actual;not null;"`
	Date_paid     time.Time `gorm:"column:date_paid;not null;"`
	Name          string    `gorm:"column:name;size:128;not null;"`
}

type Groceries struct {
	Amount_actual float64   `gorm:"column:amount_actual;not null;"`
	Date_paid     time.Time `gorm:"column:date_paid;not null;"`
	Name          string    `gorm:"column:name;size:128;not null;"`
}

type Savings struct {
	Amount_actual float64   `gorm:"column:amount_actual;not null;"`
	Date_paid     time.Time `gorm:"column:date_paid;not null;"`
	Name          string    `gorm:"column:name;size:128;not null;"`
}

type Budget struct {
	Amount_actual float64 `gorm:"column:amount_actual;not null;"`
	Name          string  `gorm:"column:name;size:128;not null;"`
}

func database_create() {
	dsn := "host=localhost user=postgres password=A1qwerty! dbname=budgeter port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	db.AutoMigrate(&Bill{}, NatashaExpense{}, &JustinExpense{}, &FuelExpense{}, &HygeineExpense{}, &Groceries{}, &Savings{}, &Budget{})

}
