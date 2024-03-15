package main

import (
	"database/sql"
	"fmt"
	"fun-with-gin/internal/repository/pgsql/models"
	"fun-with-gin/pkg/config"
	"fun-with-gin/pkg/utils"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"os"
	"time"
)

func main() {

	dbConf := config.DatabasePGSQL()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConf.Host, dbConf.Port, dbConf.UserName, dbConf.Password, dbConf.Database)

	fmt.Println(connStr)

	pgDb, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer pgDb.Close()

	err = pgDb.Ping()
	if err != nil {
		panic(err)
	}

	_, err = pgDb.Exec(fmt.Sprintf(`set search_path to "%s"`, dbConf.Schema))
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(pgDb, &postgres.Config{
		SchemaName: dbConf.Schema,
	})
	if err != nil {
		panic(err)
	}
	defer driver.Close()

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+os.Getenv("MIGRATION_DIR"),
		"postgres",
		driver,
	)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil {
		panic(err)
	}

	fmt.Println("success migrate")

	needSeeder := utils.DefaultValueBool("NEED_MIGRATION", false)

	password, _ := utils.HashPassword("whosyourdaddy")
	if needSeeder {
		firstUser := models.User{
			ID:        1,
			Name:      "fan",
			Email:     "rfanazhari@gmail.com",
			Password:  password,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		stmt, err := pgDb.Prepare(fmt.Sprintf("INSERT INTO %s (id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", firstUser.TableName()))
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		_, errInsert := stmt.Exec(firstUser)
		if errInsert != nil {
			panic(errInsert)
		}

		fmt.Println("success seed")
	}
	os.Exit(0)
}
