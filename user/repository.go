package user

import "gorm.io/gorm"

// Repository is a contract what user can do with this service
type Repository interface {
	// Save is a function to save user data to the database
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

// repository is an object to connect to the database
type repository struct {
	// db is an instance of gorm.DB
	db *gorm.DB
}

// Function to create a new repository
func NewRepository(db *gorm.DB) *repository {
  return &repository{db}
}

// Function to save user data to the database
func (r *repository) Save(user User) (User, error) {
	// Create user data to the database
	err := r.db.Create(&user).Error
	
	// If there is an error, return the error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Function to find user data by email
func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	// Find user data by email
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}