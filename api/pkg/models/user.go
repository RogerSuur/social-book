package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	Birthday  time.Time
	Nickname  string
	About     string
	ImagePath string
	CreatedAt time.Time
	IsPublic  bool
}

type IUserRepository interface {
	Insert(*User) (int64, error)
	Update(*User) error
	GetById(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByUserName(userName string) (*User, error)
	CheckIfNicknameExists(nickname string, id int64) error
	GetAllUserFollowers(id int) ([]*User, error)
	GetAllFollowedBy(id int) ([]*User, error)
}

type UserRepository struct {
	Logger *log.Logger
	DB     *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		DB:     db,
	}
}

func (repo UserRepository) Insert(user *User) (int64, error) {
	query := `INSERT INTO users (forname, surname, email, password, birthday, nickname, about, image_path, is_public, created_at)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	args := []interface{}{
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Birthday,
		user.Nickname,
		user.About,
		user.ImagePath,
		user.IsPublic,
		time.Now(),
	}

	result, err := repo.DB.Exec(query, args...)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	repo.Logger.Printf("Inserted user: %s / %s %s (last insert ID: %d)", user.Email, user.FirstName, user.LastName, lastId)

	return lastId, nil
}

func (repo UserRepository) Update(user *User) error {
	query := `UPDATE users SET forname = ?, surname = ?, email = ?, password = ?, birthday = ?, 
	nickname = ?, about = ?, image_path = ?, isPublic = ? WHERE id = ?`

	args := []interface{}{
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Birthday,
		user.Nickname,
		user.About,
		user.ImagePath,
		user.IsPublic,
		user.Id,
	}

	_, err := repo.DB.Exec(query, args...)

	return err
}

func (repo UserRepository) GetById(id int64) (*User, error) {
	query := `SELECT id, forname, surname, email, password, birthday, nickname, about, image_path, created_at, is_public FROM users WHERE id = ?`
	row := repo.DB.QueryRow(query, id)
	user := &User{}

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Birthday, &user.Nickname, &user.About, &user.ImagePath, &user.CreatedAt, &user.IsPublic)

	return user, err
}

func (repo UserRepository) GetByEmail(email string) (*User, error) {
	query := `SELECT id, forname, surname, email, password, birthday, nickname, about, image_path, created_at, is_public  FROM users WHERE email = ?`
	row := repo.DB.QueryRow(query, email)
	user := &User{}

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Birthday, &user.Nickname, &user.About, &user.ImagePath, &user.CreatedAt, &user.IsPublic)

	return user, err
}

func (repo UserRepository) GetByUserName(name string) (*User, error) {
	query := `SELECT id, forname, surname, email, password, birthday, nickname, about, image_path, created_at, is_public FROM users WHERE nickname = ?`
	row := repo.DB.QueryRow(query, name)
	user := &User{}

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Birthday, &user.Nickname, &user.About, &user.ImagePath, &user.CreatedAt, &user.IsPublic)

	return user, err
}

func (repo UserRepository) CheckIfNicknameExists(nickname string, id int64) error {
	query := `SELECT id FROM users WHERE nickname = ? AND id != ?`
	row := repo.DB.QueryRow(query, nickname, id)
	user := &User{}

	err := row.Scan(&user.Id)

	return err
}

// Return all user followers, who follow user with given id
func (repo UserRepository) GetAllUserFollowers(id int) ([]*User, error) {
	stmt := `SELECT users.id, users.forname, users.surname, users.email, users.password, birthday, nickname, about, image_path, created_at, is_public FROM users
	 INNER JOIN followers f on f.follower_id = users.id AND f.following_id = ?`

	rows, err := repo.DB.Query(stmt, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		user := &User{}

		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Birthday, &user.Nickname, &user.About, &user.ImagePath, &user.CreatedAt, &user.IsPublic)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Return all followed users by user id
func (repo UserRepository) GetAllFollowedBy(id int) ([]*User, error) {

	stmt := `SELECT users.id, users.forname, users.surname, users.email, users.password, birthday, nickname, about, image_path, created_at, is_public FROM users
	 INNER JOIN followers f on f.following_id = users.id AND f.follower_id = ?`

	rows, err := repo.DB.Query(stmt, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		user := &User{}

		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Birthday, &user.Nickname, &user.About, &user.ImagePath, &user.CreatedAt, &user.IsPublic)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
