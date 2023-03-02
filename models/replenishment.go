package models

import (
	"AlifTask/base"
	"gorm.io/gorm"
	"time"
)

type Replenishment struct {
	ID            int     `json:"id" gorm:"primaryKey;autoIncrement"`
	OperationTime string  `json:"operationTime"`
	Account       string  `json:"account" gorm:"size:20"`
	Amount        float64 `json:"amount"`
	Status        int     `json:"status"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type SumCount struct {
	Count float64 `json:"count"`
	Sum   float64 `json:"sum"`
}

func (r *Replenishment) Create() error {
	r.Status = 1
	db := base.GetDB()
	db = db.Model(r).Create(r)
	if db.Error != nil {
		return db.Error
	}
	cl := Client{Account: r.Account}
	if err := cl.Replenishment(r.Amount); err != nil {
		return err
	}
	r.Status = 10
	return db.Model(r).Save(r).Error
}

func GetCountSum(account string) (SumCount, error) {
	sc := SumCount{}
	db := base.GetDB()
	db = db.Model(Replenishment{}).Where("date_part('month', operation_time)=? and status=10 and account=?", time.Now().Format("02"), account).
		Select("sum(amount) as sum,count(1) as count").Scan(&sc)
	return sc, db.Error
}
