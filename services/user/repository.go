package user

import (
	"ourstartup/entities"

	"gorm.io/gorm"
)

// used for interacting with service
type Repository interface {
	Save(user entities.User) (entities.User, error)
	FindByQuery(query string) (entities.User, error)
	FindById(id int) (entities.User, error)
	Update(user entities.User) (entities.User, error)
}

// private
type repository struct {
	//variable with type gorm.DB
	db *gorm.DB
}

// constructor
func CreateRepository(db *gorm.DB) *repository {
	// return the repository where db is db from params
	return &repository{db}
}

// create a function called save for "repository" that the
// parameter is user object and the return is entities.User object or error
func (r *repository) Save(user entities.User) (entities.User, error) {
	//create user object on db with user data from params
	//and return assign error to err if error
	err := r.db.Create(&user).Error
	//if error return user and the error
	if err != nil {
		return user, err
	}
	//if not error return user
	return user, nil
}

func (r *repository) FindByQuery(query string) (entities.User, error) {
	// init variable with type user
	var user entities.User
	// find in table user where email = email and save it to variable user
	err := r.db.Where("email = ?", query).Or("username = ?", query).Find(&user).Error
	//if error return user and the error
	if err != nil {
		return user, err
	}
	//if not error return user
	return user, nil
}

func (r *repository) FindById(id int) (entities.User, error) {
	// init variable with type user
	var user entities.User
	// find in table user where id = id and save it to variable user
	err := r.db.Where("id = ?", id).Find(&user).Error
	//if error return user and the error
	if err != nil {
		return user, err
	}
	//if not error return user
	return user, nil
}

func (r *repository) Update(user entities.User) (entities.User, error) {
	// update user data in db that have the parameter user's id
	err := r.db.Save(&user).Error

	//if error return user and the error
	if err != nil {
		return user, err
	}
	//if not error return user
	return user, nil
}
