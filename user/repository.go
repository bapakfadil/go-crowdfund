package user

import "gorm.io/gorm"

// Repository is a contract what user can do with this service
type Repository interface {
	// Save is a function to save user data to the database
	Save(user User) (User, error)
}

// repository is a struct to define the repository
type repository struct {
	// db is an instance of gorm.DB
	db *gorm.DB
}

// NewRepository is a constructor
func NewRepository(db *gorm.DB) *repository {
  return &repository{db}
}

// r *repository is a receiver
func (r *repository) Save(user User) (User, error) {
	// Create user data to the database
	err := r.db.Create(&user).Error
	
	// If there is an error, return the error
	if err != nil {
		return user, err
	}

	return user, nil
}