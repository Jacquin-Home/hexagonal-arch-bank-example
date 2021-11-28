package repositories

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"hexagonal-example/internal/domain"
	"log"
)

type sqliteDB struct {
	sqlite *sql.DB
}

func NewSqliteDB() *sqliteDB {
	db, err := sql.Open("sqlite3", "../../internal/repositories/local.db")
	if err != nil {
		panic(err)
	}

	// temporary:
	ex, err := db.Exec(`
		DROP TABLE IF EXISTS account;
		CREATE TABLE account (
					  id TEXT,
					  money INT
					 );
	`)
	fmt.Println(ex.RowsAffected())
	// end temporary

	return &sqliteDB{
		sqlite: db,
	}
}

func (db *sqliteDB) GetAccount(accountId uuid.UUID) (*domain.Account, error) {
	stmt := `
		SELECT id, money
		  FROM account
		 WHERE id = $1;
	`
	rows := db.sqlite.QueryRow(stmt, accountId)

	var account domain.Account
	err := rows.Scan(&account.Id, &account.Money)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &account, nil
}

func (db *sqliteDB) SaveAccount(account *domain.Account) error {
	fmt.Println("save:")
	fmt.Println(account.Id)
	stmt := `
		INSERT INTO account (id, money)
			 VALUES ($1, $2);
	`
	exec, err := db.sqlite.Exec(stmt, account.Id, account.Money)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(exec.RowsAffected())
	return nil
}

func (db *sqliteDB) SavePayment(payment domain.Payment) error {
	return nil
}
