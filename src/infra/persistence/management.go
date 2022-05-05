package persistence

import (
	"github.com/sirupsen/logrus"
	"sake_io_api/domain/model"
	"sake_io_api/domain/repository"
	"sake_io_api/infra/RDB"
	"time"
)

type managementPersistence struct{}

func NewManagementPersistence() repository.ManagementRepository {
	return &managementPersistence{}
}

func (up managementPersistence) GetManagementAll() []model.Management {
	db := RDB.MySQLConnect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM management;")
	if err != nil {
		logrus.Error(err)
	}

	var managements []model.Management
	for rows.Next() {
		management := model.Management{}
		if err := rows.Scan(&management.Management_id, &management.User_id, &management.Sake_name, &management.Amount, &management.Date, &management.Created_at, &management.Updated_at); err != nil {
			logrus.Error(err)
		}
		managements = append(managements, management)
	}
	return managements
}

func (up managementPersistence) GetManagement(user_id int) []model.Management {
	db := RDB.MySQLConnect()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM management WHERE user_id=?;")
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(user_id)
	if err != nil {
		logrus.Error(err)
	}

	var managements []model.Management
	for rows.Next() {
		management := model.Management{}
		if err := rows.Scan(&management.Management_id, &management.User_id, &management.Sake_name, &management.Amount, &management.Date, &management.Created_at, &management.Updated_at); err != nil {
			logrus.Error(err)
		}
		managements = append(managements, management)
	}
	return managements
}

func (up managementPersistence) RegisterManagement(m *model.RegisterManagementType) {
	db := RDB.MySQLConnect()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO management(user_id, sake_name, amount, date, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?);")
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.User_id, m.Sake_name, m.Amount, m.Date, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}
}

func (up managementPersistence) UpdateManagement(m *model.UpdateManagementType) {
	db := RDB.MySQLConnect()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE management SET amount=?, updated_at=? WHERE id=?;")
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.Amount, time.Now(), m.Management_id)
	if err != nil {
		panic(err)
	}
}

func (up managementPersistence) DeleteManagement(management_id int) {
	db := RDB.MySQLConnect()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM management WHERE id=?")
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(management_id)
	if err != nil {
		panic(err)
	}
}
