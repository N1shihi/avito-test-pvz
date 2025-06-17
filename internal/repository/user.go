package repository

import (
	"avito-test-pvz/internal/models"
	_ "github.com/lib/pq"
)

func (r UserRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO "users" (id,email, password_hash, role) 
		VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, user.ID, user.Email, user.Password, user.Role)
	if err != nil {
		return err
	}
	return nil
}
