package table

import (
	"gorm.io/gorm"
)

type TableManager interface {
	GetTables() (GetResponse, error)
}

type TableManagerImpl struct {
	database *gorm.DB
}

func NewTableManager(database *gorm.DB) TableManager {
	return &TableManagerImpl{
		database: database,
	}
}

func (m *TableManagerImpl) GetTables() (resp GetResponse, err error) {
	var tables []Table
	err = m.database.Model(&Table{}).Find(&tables).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
		return
	}

	resp.Tables = tables
	return
}
