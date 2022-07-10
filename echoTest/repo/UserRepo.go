package repo

import (
	"database/sql"
	"errors"
	"example/go-learning/echoTest/model"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) ListAllUsers() ([]model.User, error) {
	rows, err := r.db.Query("select id, name, address, created_at, updated_at from user;")
	if err != nil {
		log.Printf("failed list all users.%v", err)
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			if mysqlError.Number == 1062 {
				log.Printf("已有该分类，请修改分类名~")
			}
		}
		return nil, err
	}

	users := make([]model.User, 0)

	for rows.Next() {
		u := model.User{}
		err := rows.Scan(&u.Id, &u.Name, &u.Address, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			log.Printf("failed set user")
			continue
		}

		users = append(users, u)
	}

	return users, nil
}

func (r *UserRepo) SaveUser(user model.User) (*model.User, error) {
	sql := `insert into user (name,address) values(?,?)`
	stmt, err2 := r.db.Prepare(sql)
	if err2 != nil {
		return nil, err2
	}
	defer stmt.Close()

	res, err2 := stmt.Exec(user.Name, user.Address)
	if err2 != nil {
		return nil, err2
	}
	id, _ := res.LastInsertId()
	user.Id = int(id)

	return &user, nil
}

func (r *UserRepo) DeleteUser(id int) (int64, error) {
	sql := `delete from user where id = ?`
	stmt, err2 := r.db.Prepare(sql)
	if err2 != nil {
		return 0, err2
	}
	defer stmt.Close()

	res, err2 := stmt.Exec(id)
	if err2 != nil {
		return 0, err2
	}

	affected, err2 := res.RowsAffected()

	return affected, nil
}

func (r *UserRepo) UpdateUser(user *model.User) (int64, error) {
	sql := `update user set name =?,address=? where id = ?`
	stmt, err2 := r.db.Prepare(sql)
	if err2 != nil {
		return 0, err2
	}
	defer stmt.Close()

	res, err2 := stmt.Exec(user.Name, user.Address, user.Id)
	if err2 != nil {
		return 0, err2
	}

	affected, err2 := res.RowsAffected()

	return affected, nil
}

func (r *UserRepo) QueryUser(id int) (*model.User, error) {
	sql := `select id,name,address,created_at,updated_at from user where id = ?`
	stmt, err2 := r.db.Prepare(sql)
	if err2 != nil {
		return nil, err2
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(id)
	if err2 != nil {
		return nil, err2
	}

	u := model.User{}

	if rows.Next() {
		err := rows.Scan(&u.Id, &u.Name, &u.Address, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			log.Printf("failed set user")
		}
	} else {
		return nil, errors.New(fmt.Sprintf("No data found for id:%v", id))
	}

	return &u, nil
}
