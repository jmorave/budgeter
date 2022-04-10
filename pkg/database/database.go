package database

import (
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
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

type Migrator interface {
	// AutoMigrate
	AutoMigrate(dst ...interface{}) error
  
	// Database
	CurrentDatabase() string
  
	// Tables
	CreateTable(dst ...interface{}) error
	DropTable(dst ...interface{}) error
	HasTable(dst interface{}) bool
	RenameTable(oldName, newName interface{}) error
  
	// Columns
	AddColumn(dst interface{}, field string) error
	DropColumn(dst interface{}, field string) error
	AlterColumn(dst interface{}, field string) error
	HasColumn(dst interface{}, field string) bool
	RenameColumn(dst interface{}, oldName, field string) error
  
	// Constraints
	CreateConstraint(dst interface{}, name string) error
	DropConstraint(dst interface{}, name string) error
	HasConstraint(dst interface{}, name string) bool
  
	// Indexes
	CreateIndex(dst interface{}, name string) error
	DropIndex(dst interface{}, name string) error
	HasIndex(dst interface{}, name string) bool
	RenameIndex(dst interface{}, oldName, newName string) error
  }

func Database_connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=A1qwerty! dbname=budgeter port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}
	return db

}

func Database_create() {
	if Database_connect().Migrator().HasTable(&Bill{}) == false {
		Database_connect().AutoMigrate(&Bill{}, &NatashaExpense{}, &JustinExpense{}, &FuelExpense{}, &HygeineExpense{}, &Groceries{}, &Savings{}, &Budget{})
		fmt.Printf("Creating Database")
	} else {
		fmt.Printf("Database already exists")
	}
}