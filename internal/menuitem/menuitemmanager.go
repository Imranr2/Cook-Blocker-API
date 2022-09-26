package menuitem

import (
	"gorm.io/gorm"
)

type MenuItemManager interface {
	GetMenuItems() (GetResponse, error)
	GetMenuItemWithID(GetWithIDRequest) (GetWithIDResponse, error)
	CreateMenuItem(CreateRequest) (CreateResponse, error)
	DeleteMenuItem(DeleteRequest) (DeleteResponse, error)
}

type MenuItemManagerImpl struct {
	database *gorm.DB
}

func NewMenuItemManager(database *gorm.DB) MenuItemManager {
	return &MenuItemManagerImpl{
		database: database,
	}
}

func (m *MenuItemManagerImpl) GetMenuItems() (resp GetResponse, err error) {
	var menuItems []MenuItem
	err = m.database.Model(&MenuItem{}).Preload("Ingredients").Preload("Image").Find(&menuItems).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
		return
	}

	resp.MenuItems = menuItems
	return
}

func (m *MenuItemManagerImpl) GetMenuItemWithID(req GetWithIDRequest) (resp GetWithIDResponse, err error) {
	var menuItem MenuItem
	err = m.database.Model(&MenuItem{}).Preload("Ingredients").Preload("Image").First(&menuItem, req.ID).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
		return
	}

	resp.MenuItem = &menuItem
	return
}

func (m *MenuItemManagerImpl) CreateMenuItem(req CreateRequest) (resp CreateResponse, err error) {
	var ingredients []*Ingredient
	for _, ingredient := range req.Ingredients {
		ingredient := ingredient
		ingredients = append(ingredients, &ingredient)
	}

	newMenuItem := &MenuItem{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Steps:       req.Steps,
		CreatedBy:   req.UserID,
		Ingredients: ingredients,
		Image:       req.Image,
	}

	err = m.database.Create(&newMenuItem).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
	}
	return
}

func (m *MenuItemManagerImpl) DeleteMenuItem(req DeleteRequest) (resp DeleteResponse, err error) {
	var menuItem MenuItem
	err = m.database.First(&menuItem, req.ID).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
		return
	}

	err = m.database.Model(&menuItem).Association("Ingredients").Clear()
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
		return
	}

	err = m.database.Delete(&menuItem, req.ID).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
	}
	return
}
