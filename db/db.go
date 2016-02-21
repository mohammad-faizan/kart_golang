package db

import(
	"simple-server/config"
	_"github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	d "simple-server/data"
	"database/sql"
	"errors"
	"time"
)

func NewDbConnection() (DbAdapter, error) {
	db, err := sqlx.Connect("postgres", config.GetDbConnectionString())

	if err != nil {
		return DbAdapter{}, err
	}

	return DbAdapter{db}, nil
}

type DbAdapter struct {
	Connection *sqlx.DB
}

func (this *DbAdapter) UserList() ([]d.User, error) {
	query := `SELECT id, email, first_name, last_name, COALESCE(dob, NOW()) as dob, gender from users`

	var users []d.User

	err := this.Connection.Select(&users, query)

	if err == sql.ErrNoRows {
		err = nil
	}

	return users, err
}

func (this *DbAdapter) SearchByEmail(email string) (d.User, error) {
	query := `SELECT id, first_name, last_name, dob, gender, email, encrypted_password from users WHERE email = $1`

	var user d.User

	err := this.Connection.Get(&user, query, email)

	if err == sql.ErrNoRows {
		err = errors.New("username/password invalid")
	}

	return user, err
}

func (this *DbAdapter) CreateUser(u d.User) (sql.Result, error) {
	query := getInsertQuery(`users`)
	query = query + ` VALUES ($1, $2, $3, $4, $5, $6, $7)`

	pass_hash, err := u.PasswordHash()

	if err != nil {
		return sql.Result(nil), err
	}

	return this.Connection.Exec(query, u.FirstName, u.LastName, u.Email, u.Gender, pass_hash, time.Now(), time.Now())
}

func getInsertQuery(table string) string {
	q := `INSERT INTO ` + table + ` `
	columns := ``
	switch table {
	case `users`:
		columns = `(first_name, last_name, email, gender, encrypted_password, created_at, updated_at)`
	}
	q += columns
	return q
}