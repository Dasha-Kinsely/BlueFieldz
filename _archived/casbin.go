package models

import (
	"github.com/casbin/casbin"
)

type CustomCasbinModel struct {
	CasbinID string `gorm:"column:casbin_id;primaryKey"`
	Role string `gorm:"role"`
	Policy string `gorm:"column:policy"`
	Path string `gorm:"column:path"`
	Method string `gorm:"column:method"`
}

func Casbin() *casbin.Enforcer {
	adapter := mysqladapter.NewAdapter("mysql", )
	e, err := casbin.NewEnforcer("/casbins/admin_model.conf", adapter)
	if err != nil {
		return err
	}
	e.LoadPolicy()
	return e
}

