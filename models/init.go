package models

import "AlifTask/base"

func Init() error {
	if err := base.Migrate(&Client{}); err != nil {
		return err
	}
	if err := base.Migrate(&Replenishment{}); err != nil {
		return err
	}
	return nil
}
