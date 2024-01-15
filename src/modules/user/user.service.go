package user

import (
	"github.com/backend-timedoor/gtimekeeper-framework/app"
	"github.com/backend-timedoor/gtimekeeper-framework/utils/helper"
	"github.com/backend-timedoor/gtimekeeper-framework/utils/paginate"
)

type ServiceUser struct{}

type ServiceUserInterface interface {
	FindAll(QueryRequest) ([]User, paginate.Pagination, error)
	Store(CreateUserRequest) (*User, error)
	FindById(int) (*User, error)
	Update(UpdateUserRequest, int) (*User, error)
	Delete(int) error
}

func (s *ServiceUser) FindAll(req QueryRequest) (users []User, pagination paginate.Pagination, err error) {
	err = app.DB.Scopes(
		paginate.Paginate(&users, &pagination, &req.PaginationRequest),
		req.FilterByUsername,
		req.FilterByName,
	).Find(&users).Error

	if err != nil {
		return users, pagination, err
	}

	return users, pagination, nil
}

func (s *ServiceUser) Store(req CreateUserRequest) (*User, error) {
	user := &User{}

	helper.Clone(&user, &req)

	if err := app.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *ServiceUser) FindById(id int) (*User, error) {
	user := &User{}

	if err := app.DB.Find(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *ServiceUser) Update(req UpdateUserRequest, id int) (*User, error) {
	user, err := s.FindById(id)
	if err != nil {
		return nil, err
	}

	helper.Clone(&user, &req)

	if err := app.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *ServiceUser) Delete(id int) error {
	// user, err := s.FindById(id)
	// if err != nil {
	// 	return err
	// }

	if err := app.DB.Delete(&User{}, id).Error; err != nil {
		return err
	}

	return nil
}
