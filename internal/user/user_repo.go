package user

import (
	"context"
	"fmt"

	"github.com/bookofshame/bookofshame/pkg/database"
	"github.com/bookofshame/bookofshame/pkg/logging"
	"go.uber.org/zap"
)

type Repository struct {
	db     *database.Sql
	logger *zap.SugaredLogger
}

func NewRepository(ctx context.Context, sql *database.Sql) Repository {
	return Repository{
		db:     sql,
		logger: logging.FromContext(ctx),
	}
}

func (r *Repository) GetAll() ([]User, error) {
	users := []User{}
	err := r.db.Select(&users, "SELECT * FROM user")

	if err != nil {
		r.logger.Errorf("query error: %w", err)
		return nil, fmt.Errorf("failed to fetch users")
	}

	return users, nil
}

func (r *Repository) GetByEmail(email string) (*User, error) {
	user := []User{}
	err := r.db.Select(&user, "SELECT * FROM user WHERE email=?", email)

	if err != nil {
		r.logger.Errorf("query error: %w", email, err)
		return nil, fmt.Errorf("failed to fetch user (email = %s)", email)
	}

	if len(user) == 0 {
		return nil, nil
	}

	return &user[0], nil
}

func (r *Repository) GetById(id int) (*User, error) {
	user := []User{}
	err := r.db.Select(&user, "SELECT * FROM user WHERE id=?", id)

	if err != nil {
		r.logger.Errorf("query error: %w", id, err)
		return nil, fmt.Errorf("failed to fetch user (id = %d)", id)
	}

	if len(user) == 0 {
		return nil, nil
	}

	return &user[0], nil
}

func (r *Repository) GetByPhone(phone string) (*User, error) {
	var user []User
	err := r.db.Select(&user, "SELECT * FROM user WHERE phone=?", phone)

	if err != nil {
		r.logger.Errorf("query error: %w", err)
		return nil, fmt.Errorf("failed to fetch user (phone = %s)", phone)
	}

	if len(user) == 0 {
		return nil, nil
	}

	return &user[0], nil
}

func (r *Repository) GetIdByActivationCode(code string) (int, error) {
	id := []int{}
	err := r.db.Select(&id, "SELECT id FROM user WHERE activationCode=?", code)

	if err != nil {
		r.logger.Errorf("query error: %w", err)
		return 0, fmt.Errorf("failed to fetch user by activation code")
	}

	if len(id) == 0 {
		return 0, fmt.Errorf("failed to fetch user by activation code")
	}

	return id[0], nil
}

func (r *Repository) Create(user User) (int64, error) {
	res, err := r.db.Exec(`
        INSERT INTO user (fullName, address, phone, email, password, activationCode, isActive) VALUES (?, ?, ?, ?, ?, ?, ?)
    `, user.FullName, user.Address, user.Phone, user.Email, user.Password, user.ActivationCode, false)

	if err != nil {
		r.logger.Errorf("query error: %w", err)
		return 0, fmt.Errorf("query failed to create user")
	}

	id, _ := res.LastInsertId()

	return id, nil
}

func (r *Repository) PhoneExists(phone string) (bool, error) {
	id := []int{}
	if err := r.db.Select(&id, "SELECT id FROM user WHERE phone=?", phone); err != nil {
		r.logger.Errorf("query error: %w", err)
		return false, fmt.Errorf("failed to fetch existing user")
	}

	exists := len(id) > 0 && id[0] > 0
	return exists, nil
}

func (r *Repository) EmailExists(email string) (bool, error) {
	id := []int{}
	if err := r.db.Select(&id, "SELECT id FROM user WHERE email=?", email); err != nil {
		r.logger.Errorf("query error: %w", err)
		return false, fmt.Errorf("failed to fetch existing user")
	}

	exists := len(id) > 0 && id[0] > 0
	return exists, nil
}

func (r *Repository) Activate(id int) error {
	_, err := r.db.Exec(`UPDATE user SET isActive = true, activationCode = NULL WHERE id = ?`, id)

	if err != nil {
		r.logger.Errorf("query error: %w", err)
		return fmt.Errorf("failed activate user")
	}

	return nil
}

func (r *Repository) Update(user User) error {
	_, err := r.db.Exec(`
        UPDATE user SET fullName = ?, address = ?, phone = ?, email = ?, password = ?, activationCode = ?, isActive = ? WHERE id = ?
    `, user.FullName, user.Address, user.Phone, user.Email, user.Password, user.ActivationCode, user.IsActive, user.Id)

	if err != nil {
		r.logger.Errorf("query error: %w", err)
		return fmt.Errorf("query failed to update user")
	}

	return nil
}

func (r *Repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM user WHERE id=?", id)

	if err != nil {
		r.logger.Errorf("query error: %w", err)
		return fmt.Errorf("failed to delete user")
	}

	return nil
}
