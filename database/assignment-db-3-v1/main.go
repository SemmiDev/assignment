package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "embed"

	_ "github.com/lib/pq"
)

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

type FinalScore struct {
	ID           int     `sql:"id"`
	Fullname     string  `sql:"fullname"`
	Class        string  `sql:"class"`
	AverageScore float64 `sql:"average_score"`
}

func Connect(creds *Credential) (*sql.DB, error) {
	// this is only an example, please modify it to your need
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	// connect using database/sql + pq
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

//go:embed select.sql
var queryStr string

func QuerySQL(db *sql.DB) ([]FinalScore, error) {
	var res []FinalScore

	rows, err := db.Query(queryStr)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var f FinalScore

		if err := rows.Scan(&f.ID, &f.Fullname, &f.Class, &f.AverageScore); err != nil {
			return res, err
		}

		res = append(res, f)
	}

	// sort.Slice(res, func(i, j int) bool {
	// 	return res[i].AverageScore > res[j].AverageScore
	// })

	return res, nil
}

var (
	CreateSQL = `CREATE TABLE IF NOT EXISTS final_scores (
	id SERIAL PRIMARY KEY,
	exam_id VARCHAR(255) NOT NULL,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL,
	bahasa_indonesia INT NOT NULL,
	bahasa_inggris INT NOT NULL,
	matematika INT NOT NULL,
	ipa INT NOT NULL,
	exam_status VARCHAR(50) NOT NULL,
	fee_status VARCHAR(50) NOT NULL );`

	InsertSQL = `INSERT INTO final_scores (exam_id, first_name, last_name, bahasa_indonesia, bahasa_inggris, matematika, ipa, exam_status, fee_status) VALUES
	('1A-001', 'John', 'Doe', 80, 90, 70, 80, 'pass', 'full'),
	('1A-002', 'Jane', 'Doe', 90, 80, 90, 80, 'pass', 'full'),
	('1B-003', 'John', 'Smith', 70, 80, 70, 80, 'pass', 'full'),
	('1B-004', 'Jane', 'White', 80, 70, 80, 80, 'pass', 'full'),
	('1B-005', 'John', 'Bernard', 80, 90, 70, 80, 'pass', 'full'),
	('1B-006', 'Jane', 'Abrams', 90, 80, 90, 80, 'pass', 'full'),
	('1B-007', 'John', 'Albert', 70, 80, 70, 80, 'pass', 'full');
	`
)

func SQLExecute(db *sql.DB, insert string) error {
	_, err := db.Exec(CreateSQL)
	if err != nil {
		return err
	}

	fmt.Println("success create table")

	_, err = db.Exec(insert)
	if err != nil {
		return err
	}
	fmt.Println("success insert data")

	return nil
}

func main() {
	dbCredential := Credential{
		Host:         "localhost",
		Username:     "root",
		Password:     "secret",
		DatabaseName: "assignments",
		Port:         5432,
	}
	dbConn, err := Connect(&dbCredential)
	if err != nil {
		log.Fatal(err)
	}

	_, err = dbConn.Exec("DROP TABLE IF EXISTS final_scores CASCADE")
	if err != nil {
		log.Fatal("error drop table: " + err.Error())
	}

	err = SQLExecute(dbConn, `INSERT INTO final_scores (exam_id, first_name, last_name, bahasa_indonesia, bahasa_inggris, matematika, ipa, exam_status, fee_status)
	VALUES ('1A-001', 'John', 'Doe', 80, 90, 70, 80, 'pass', 'full'),
	('1A-002', 'Jane', 'Doe', 90, 80, 90, 80, 'pass', 'not paid'),
	('1A-003', 'John', 'Smith', 70, 80, 70, 80, 'pass', 'installment'),
	('1A-004', 'Jane', 'White', 80, 70, 80, 80, 'pass', 'full'),
	('1A-005', 'Abrams', 'White', 80, 70, 80, 80, 'pass', 'full'),
	('1A-006', 'Herdi', 'White', 80, 70, 80, 80, 'fail', 'not paid'),
	('1A-007', 'Wendy', 'White', 100, 95, 80, 80, 'fail', 'installment'),
	('1A-008', 'Ardi', 'White', 100, 95, 80, 80, 'pass', 'not paid'),
	('1A-009', 'Abrams', 'Smith', 95, 93, 80, 80, 'fail', 'not paid'),
	('1A-010', 'Welly', 'White', 95, 93, 80, 80, 'fail', 'not paid'),
	('1B-001', 'Indah', 'Sudarni', 95, 93, 80, 80, 'fail', 'full'),
	('1B-002', 'Aren', 'White', 80, 70, 80, 80, 'pass', 'full'),
	('1B-003', 'John', 'Bernard', 80, 90, 70, 80, 'fail', 'installment'),
	('1B-004', 'Jane', 'Abrams', 90, 80, 90, 80, 'pass', 'full'),
	('1B-005', 'John', 'Albert', 70, 80, 70, 80, 'pass', 'installment');`)
	if err != nil {
		log.Fatal("error SQL execute: " + err.Error())
	}

	res, err := QuerySQL(dbConn)
	if err != nil {
		log.Fatal("query error: " + err.Error())
	}

	fmt.Print("\033[H\033[2J")

	fmt.Println("OUTPUTNYA GINI")
	for _, s := range res {
		fmt.Println(s)
	}

	fmt.Println("\n\nMAUNYA GINI")

	// clean screen
	a := ([]FinalScore{
		{ID: 14, Fullname: "Jane Abrams", Class: "1B", AverageScore: 85},
		{ID: 1, Fullname: "John Doe", Class: "1A", AverageScore: 80},
		{ID: 12, Fullname: "Aren White", Class: "1B", AverageScore: 77},
		{ID: 5, Fullname: "Abrams White", Class: "1A", AverageScore: 77},
		{ID: 4, Fullname: "Jane White", Class: "1A", AverageScore: 77},
	})

	for _, s := range a {
		fmt.Println(s)
	}

	// 80 + 90 + 70 + 80 = 320 / 4 = 80
	// 70 + 80+ 70+ 80 = 300 / 4 = 75
}
