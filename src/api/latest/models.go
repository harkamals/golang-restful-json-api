package latest

import (
	"github.com/jinzhu/gorm"
)

type Orders struct {
	gorm.Model
	Project     string `gorm:"not null"`
	Environment string `gorm:"not null"` // poc,dev,staging,prod
	CostCentre  int    `gorm:"not null"`
	Owner       string `gorm:"not null"`
	Service     string `gorm:"not null"` // aws_account,aws_ami, aws_j5
	Tier        string // Tier: 1,2,3
	Status      string
	Tracking    string
	Account     Accounts
	Image       Images
}

type Accounts struct {
	gorm.Model
	Environment string `gorm:"not null"`
	Version     int    `gorm:"not null"`
	AccountId   int    `gorm:"not null;unique"`
	Status      string
	Email       string `gorm:"not null;unique"`
}

type Images struct {
	gorm.Model
	Reference   string `gorm:"not null;unique"`
	Description string `gorm:"not null"`
}
