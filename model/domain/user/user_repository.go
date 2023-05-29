package user

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	db "github.com/captrep/gin-simple-crud/datasource/mysql"
	"github.com/captrep/gin-simple-crud/utils/res"
	"github.com/google/uuid"
)

type UserRepository interface {
	Save(user User) (User, *res.Err)
	GetAll() ([]User, *res.Err)
	FindById(userId string) (User, *res.Err)
	Update(user User) (User, *res.Err)
	Delete(user User) *res.Err
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

const (
	queryInsertUser = "INSERT INTO users (id, first_name, last_name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	queryFindAll    = "SELECT * FROM users"
	queryGetById    = "SELECT id, first_name, last_name, email, created_at FROM users WHERE id=?"
	queryUpdate     = "UPDATE users SET first_name = ?, last_name = ?, email = ?, updated_at = ? WHERE id=?"
	queryDelete     = "DELETE FROM users WHERE id=?"
	errorNoRows     = "no rows in result set"
)

func (repository *UserRepositoryImpl) Save(user User) (User, *res.Err) {
	id := uuid.New()
	statement, err := db.Mysql.Prepare(queryInsertUser)
	if err != nil {
		return user, res.NewRestErr(http.StatusInternalServerError, "internal server error", err.Error())
	}
	defer statement.Close()

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Id = id.String()

	result, err := statement.Exec(user.Id, user.Firstname, user.Lastname, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return user, res.NewRestErr(http.StatusInternalServerError, "internal server error", err.Error())
	}
	log.Println(result)
	if err != nil {
		return user, res.NewRestErr(http.StatusInternalServerError, "internal server error", err.Error())
	}

	return user, nil
}

func (repository *UserRepositoryImpl) GetAll() ([]User, *res.Err) {
	statement, err := db.Mysql.Prepare(queryFindAll)
	if err != nil {
		return nil, res.NewRestErr(http.StatusInternalServerError, "internal server error", err.Error())
	}

	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		return nil, res.NewRestErr(http.StatusInternalServerError, "internal server error", "Not found")
	}

	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, res.NewRestErr(http.StatusInternalServerError, "internal server error", err.Error())
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository *UserRepositoryImpl) FindById(userId string) (User, *res.Err) {
	statement, err := db.Mysql.Prepare(queryGetById)
	user := User{}
	if err != nil {
		return user, res.NewRestErr(http.StatusInternalServerError, "internal server error", err.Error())
	}
	defer statement.Close()

	result := statement.QueryRow(userId)
	if err := result.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.CreatedAt); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return user, res.NewRestErr(http.StatusNotFound, "not found", fmt.Sprintf("user with id %s not found", user.Id))
		}
		return user, res.NewRestErr(http.StatusInternalServerError, "internal server error", err.Error())
	}

	return user, nil

}

func (repository *UserRepositoryImpl) Update(user User) (User, *res.Err) {
	statement, err := db.Mysql.Prepare(queryUpdate)
	if err != nil {
		return user, res.NewRestErr(http.StatusInternalServerError, "internal server error", err.Error())
	}
	defer statement.Close()

	user.UpdatedAt = time.Now()
	_, err = statement.Exec(user.Firstname, user.Lastname, user.Email, user.UpdatedAt, user.Id)
	if err != nil {
		return user, res.NewRestErr(http.StatusInternalServerError, "internal server error", err.Error())
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Delete(user User) *res.Err {
	statement, err := db.Mysql.Prepare(queryDelete)
	if err != nil {
		return res.NewRestErr(http.StatusInternalServerError, err.Error(), "internal_server_error")
	}
	defer statement.Close()

	_, err = statement.Exec(user.Id)
	if err != nil {
		return res.NewRestErr(http.StatusInternalServerError, err.Error(), "internal_server_error")
	}

	return nil
}
