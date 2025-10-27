package repositories

import (
	"database/sql"

	"github.com/cperdiansyah/gin-rest-learn/internal/app/rest_api/database"
	"github.com/cperdiansyah/gin-rest-learn/internal/app/rest_api/entities"
)

type User struct {
	database.BaseSQLRepository[entities.User]
}

func NewUserRepository(db *sql.DB) *User {
	return &User{
		BaseSQLRepository: database.BaseSQLRepository[entities.User]{DB: db},
	}
}

func mapUser(rows *sql.Row, u *entities.User) error {
	return rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.PhoneNumber)
}

func mapUsers(rows *sql.Rows, u *entities.User) error {
	return rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.PhoneNumber)
}

func (r *User) FindByEmail(email string) (*entities.User, error) {
	return r.SelectSingle(
		mapUser,
		"SELECT u.id, u.first_name, u.last_name, u.email, u.phone_number FROM users u WHERE u.email = $1",
		email,
	)
}

func (r *User) FindById(id int) (*entities.User, error) {
	return r.SelectSingle(
		mapUser,
		"SELECT u.id, u.first_name, u.last_name, u.email, u.phone_number FROM users u WHERE u.id = $1",
		id,
	)
}

func (r *User) GetAllUsers() ([]*entities.User, error) {
	return r.SelectMultiple(
		mapUsers,
		"SELECT u.id, u.first_name, u.last_name, u.email, u.phone_number FROM users u",
	)
}

func (r *User) Create(user *entities.User) error {
	_, err := r.Insert(
		"INSERT INTO users (first_name, last_name, email, phone_number) VALUES ($1, $2, $3, $4)",
		user.FirstName, user.LastName, user.Email, user.PhoneNumber,
	)

	return err
}

func (r *User) Update(user *entities.User) error {
	_, err := r.ExecuteQuery(
		"UPDATE users SET first_name = $1, last_name = $2, email = $3, phone_number = $4 WHERE id = $5",
		user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.ID,
	)

	return err
}

func (r *User) DeleteUser(id int) error {
	_, err := r.ExecuteQuery("DELETE FROM users WHERE id = $1", id)

	return err
}
