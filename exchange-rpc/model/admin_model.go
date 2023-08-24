package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ AdminModel = (*customAdminModel)(nil)

type (
	// AdminModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminModel.
	AdminModel interface {
		adminModel
	}

	customAdminModel struct {
		*defaultAdminModel
	}
)

// NewAdminModel returns a model for the mongo.
func NewAdminModel(url, db, collection string) AdminModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customAdminModel{
		defaultAdminModel: newDefaultAdminModel(conn),
	}
}
