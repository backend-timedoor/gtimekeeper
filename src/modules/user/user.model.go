package user

import (
	"time"

	"github.com/backend-timedoor/gtimekeeper-framework/utils/helper"
	"github.com/backend-timedoor/gtimekeeper-framework/utils/paginate"
	"gorm.io/gorm"
)

type (
	User struct {
		ID           int       `json:"id"`
		Name         string    `json:"name"`
		Email        string    `json:"email"`
		Phone        string    `json:"phone"`
		Username     string    `json:"username"`
		Password     string    `json:"password"`
		Status       string    `json:"status"`
		UserableID   int       `gorm:"default:null" json:"userable_id,omitempty"`
		UserableType string    `gorm:"default:null" json:"userable_type,omitempty"`
		ReferalCode  string    `gorm:"default:null" json:"referal_code"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	Response struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		Phone       string `json:"phone"`
		Username    string `json:"username"`
		Status      string `json:"status"`
		ReferalCode string `json:"referal_code"`
	}

	CreateUserRequest struct {
		Name        string `json:"name" validate:"required"`
		Email       string `json:"email" validate:"required,email,unique=users:email"`
		Phone       string `json:"phone" validate:"required,unique=users:phone"`
		Username    string `json:"username" validate:"required,unique=users:username"`
		Password    string `json:"password" validate:"required,min=8"`
		Status      string `json:"status" validate:"required,oneof=active banned inactive"`
		ReferalCode string `json:"referal_code"`
	}

	UpdateUserRequest struct {
		Name        string `json:"name" validate:"required"`
		Email       string `json:"email" validate:"required,email"`
		Phone       string `json:"phone" validate:"required"`
		Username    string `json:"username" validate:"required"`
		Status      string `json:"status" validate:"required,oneof=active banned inactive"`
		ReferalCode string `json:"referal_code"`
	}

	QueryRequest struct {
		paginate.PaginationRequest
		FilterUsername string `query:"username"`
		FilterName     string `query:"name"`
	}
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	//

	if u.Password != "" {
		u.Password, err = helper.Hash(u.Password)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *QueryRequest) FilterByUsername(db *gorm.DB) *gorm.DB {
	if m.FilterUsername == "" {
		return db
	}

	return db.Where("LOWER(username) = LOWER(?)", m.FilterUsername)
}

func (m *QueryRequest) FilterByName(db *gorm.DB) *gorm.DB {
	if m.FilterName == "" {
		return db
	}

	return db.Where("name ILIKE ?", "%"+m.FilterName+"%")
}
