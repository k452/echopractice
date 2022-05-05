package persistence

import (
	"github.com/sirupsen/logrus"
	"sake_io_api/domain/model"
	"sake_io_api/domain/repository"
	"sake_io_api/infra/RDB"
	"time"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (up userPersistence) GetUserAll() []model.User {
	db := RDB.MySQLConnect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user;")
	if err != nil {
		logrus.Error(err)
	}

	var users []model.User
	for rows.Next() {
		user := model.User{}
		if err := rows.Scan(&user.User_id, &user.Name, &user.Mail, &user.Level, &user.Pass, &user.Created_at, &user.Updated_at); err != nil {
			logrus.Error(err)
		}
		users = append(users, user)
	}
	return users
}

func (up userPersistence) GetUser(user_id int) []model.User {
	db := RDB.MySQLConnect()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM user WHERE id=?;")
	if err != nil {
		logrus.Error(err)
	}

	rows, err := stmt.Query(user_id)
	if err != nil {
		logrus.Error(err)
	}

	var users []model.User
	for rows.Next() {
		user := model.User{}
		if err := rows.Scan(&user.User_id, &user.Name, &user.Mail, &user.Level, &user.Pass, &user.Created_at, &user.Updated_at); err != nil {
			logrus.Error(err)
		}
		users = append(users, user)
	}
	return users
}

func (up userPersistence) RegisterUser(u *model.RegisterUserType) {
	db := RDB.MySQLConnect()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO user(name, mail, level, pass, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?);")
	if err != nil {
		logrus.Error(err)
	}

	_, err = stmt.Exec(u.Name, u.Mail, 1, u.Pass, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}
}

func (up userPersistence) UpdateUser(u *model.UpdateUserType) {
	db := RDB.MySQLConnect()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE user SET name=?, mail=?, level=?, pass=?, updated_at=? WHERE id=?;")
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Name, u.Mail, u.Level, u.Pass, time.Now(), u.User_id)
	if err != nil {
		panic(err)
	}
}

func (up userPersistence) DeleteUser(user_id int) {
	db := RDB.MySQLConnect()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id)
	if err != nil {
		panic(err)
	}
}
