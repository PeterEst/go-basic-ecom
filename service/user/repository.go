package user

import (
	"database/sql"
	"fmt"

	"github.com/peterest/go-basic-ecom/types"
)

type UserRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	u := new(types.User)

	err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *UserRepository) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)

	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)

		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *UserRepository) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (s *UserRepository) CreateUser(user types.User) error {
	return nil
}
