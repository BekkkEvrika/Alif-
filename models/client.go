package models

import (
	"AlifTask/base"
	"gorm.io/gorm"
	"time"
)

const (
	identified   = 100000
	unidentified = 10000
)

type Client struct {
	ID             int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string  `json:"name" gorm:"size:100"`
	Account        string  `json:"account" gorm:"size:20"`
	Balance        float64 `json:"balance"`
	Identification bool    `json:"identification"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type Clients []Client

func (cl *Client) Exists() bool {
	var count int64
	db := base.GetDB()
	db = db.Model(cl).Where(cl).Count(&count)
	return count > 0
}

func (cl *Client) GetBalance() (float64, error) {
	db := base.GetDB()
	db = db.Model(cl).Where(cl).Scan(cl)
	return cl.Balance, db.Error
}

func (cl *Client) Replenishment(amount float64) error {
	db := base.GetDB()
	db = db.Model(cl).Where(cl).Scan(cl)
	if db.Error != nil {
		return db.Error
	}
	if cl.checkBalance(amount) {
		cl.Balance += amount
		return cl.update()
	}
	return nil
}

func (cl *Client) update() error {
	db := base.GetDB()
	return db.Model(cl).Save(cl).Error
}

func (cl *Client) checkBalance(amount float64) bool {
	if cl.Identification {
		return cl.Balance+amount <= identified
	} else {
		return cl.Balance+amount <= unidentified
	}
}
