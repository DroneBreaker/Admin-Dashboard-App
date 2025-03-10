package repository

import (
	"database/sql"

	"github.com/DroneBreaker/Admin-Dashboard-App/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	Create(user *models.User) error
	GetByID(id int) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Update(user *models.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	users := []models.User{}
	query := `SELECT id, firstName, lastName, username, email FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) Create(user *models.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (id, firstName, lastName, username, email, password) VALUES (?,?,?,?,?,?)`
	result, err := r.db.Exec(query, user.ID, user.FirstName, user.LastName, user.Username, user.Email, hashedPassword)

	if err != nil {
		return err
	}

	// Retrieve the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id) // Set the ID in the user object
	return nil
	// return r.db.QueryRow(query, user.Name, user.Username, user.BusinessTIN, user.Password).Scan(&user.ID)
	// fmt.Println("Inserted record ID:", result.LastInsertId())
}

func (r *userRepository) GetByID(id int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, firstName, lastName, username, email, password FROM users WHERE id =?`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email)

	return user, err
}

func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, firstName, lastName, username, email, password FROM users WHERE username =?`
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(user *models.User) error {
	query := `UPDATE users SET firstName = ?, lastName = ?, username = ?, email = ? WHERE id = ?`
	_, err := r.db.Exec(query, user.FirstName, user.LastName, user.Username, user.Email, user.ID)
	return err
	// fmt.Println("User updated")
}
