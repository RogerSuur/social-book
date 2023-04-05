package models

import (
	"database/sql"
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

type UserModel struct {
	DB *sql.DB
}

func (u UserModel) Insert(user *User) (int64, error) {
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

	result, err := u.DB.Exec(query, args...)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (u UserModel) Update(user *User) error {
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

	_, err := u.DB.Exec(query, args...)

	return err
}

func (u UserModel) GetById(id int64) (*User, error) {
	query := `SELECT id, forname, surname, email, password, birthday, nickname, about, image_path, created_at, is_public FROM users WHERE id = ?`
	row := u.DB.QueryRow(query, id)
	user := &User{}

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Birthday, &user.Nickname, &user.About, &user.ImagePath, &user.CreatedAt, &user.IsPublic)

	return user, err
}

func (u UserModel) GetByEmail(email string) (*User, error) {
	query := `SELECT id, forname, surname, email, password, birthday, nickname, about, image_path, created_at, is_public  FROM users WHERE email = ?`
	row := u.DB.QueryRow(query, email)
	user := &User{}

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Birthday, &user.Nickname, &user.About, &user.ImagePath, &user.CreatedAt, &user.IsPublic)

	return user, err
}

func (u UserModel) GetByUserName(name string) (*User, error) {
	query := `SELECT id, forname, surname, email, password, birthday, nickname, about, image_path, created_at, is_public FROM users WHERE nickname = ?`
	row := u.DB.QueryRow(query, name)
	user := &User{}

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Birthday, &user.Nickname, &user.About, &user.ImagePath, &user.CreatedAt, &user.IsPublic)

	return user, err
}

// Return all user followers, who follow user with given id
func (m UserModel) GetAllUserFollowers(id int) ([]*User, error) {
	stmt := `SELECT users.id, users.forname, users.surname, users.email, users.password, birthday, nickname, about, image_path, created_at, is_public FROM users
	 INNER JOIN followers f on f.follower_id = users.id AND f.following_id = ?`

	rows, err := m.DB.Query(stmt, id)
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
func (m UserModel) GetAllFollowedBy(id int) ([]*User, error) {

	stmt := `SELECT users.id, users.forname, users.surname, users.email, users.password, birthday, nickname, about, image_path, created_at, is_public FROM users
	 INNER JOIN followers f on f.following_id = users.id AND f.follower_id = ?`

	rows, err := m.DB.Query(stmt, id)
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
