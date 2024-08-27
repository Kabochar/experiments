package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type AlbumInfo struct {
	Id     int32   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := InitDB()
	if err != nil {
		log.Fatal("Error connecting to database", err)
		return
	}
	defer db.Close()

	fmt.Println("start getting albums")
	albumList, err := GetAlbumList(db)
	if err != nil {
		log.Fatal("Error getting album list", err)
	}
	for _, album := range albumList {
		fmt.Printf("album val=%+v\n", album)
	}
	fmt.Println("------")

	fmt.Println("start getting idx=1 album")
	targetId := 1
	album, err := GetAlbum(db, targetId)
	if err != nil {
		log.Fatal("Error getting album", err)
	}
	if album != nil {
		fmt.Printf("album idx=%d, val=%+v\n", targetId, album)
	}
	fmt.Println("------")

	fmt.Println("start inserting album")
	insertAlbum := &AlbumInfo{
		Title:  time.Now().Format("20060102150405"),
		Artist: "Alda",
		Price:  25.00,
	}
	if err = InsertAlbum(db, insertAlbum); err != nil {
		log.Fatal("Error inserting album", err)
	}
	fmt.Println("------")

	fmt.Println("start updating album")
	targetUpdate := AlbumInfo{Title: time.Now().Format("20060102150405"), Artist: "Alda", Price: 23.99}
	query := AlbumInfo{Id: 1}
	if err = UpdateAlbum(db, targetUpdate, query); err != nil {
		log.Fatal("Error updating album", err)
	}
	fmt.Println("------")

	fmt.Println("start deleting album")
	targetDeleteId := 7
	if err = DeleteAlbum(db, targetDeleteId); err != nil {
		log.Fatal("Error deleting album", err)
	}
}

func InitDB() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ping db fail: %v", err)
	}

	return db, nil
}

func GetAlbumList(db *sql.DB) ([]*AlbumInfo, error) {
	execSQL := `select id, title, artist, price from album`
	rows, err := db.Query(execSQL)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("executing query error: %w", err)
	}
	defer rows.Close()

	var albumInfoSet []*AlbumInfo
	for rows.Next() {
		var albumInfo AlbumInfo
		err = rows.Scan(&albumInfo.Id, &albumInfo.Title, &albumInfo.Artist, &albumInfo.Price)
		if err != nil {
			return nil, fmt.Errorf("scan rows error: %v", err)
		}
		albumInfoSet = append(albumInfoSet, &albumInfo)
	}

	return albumInfoSet, nil
}

func GetAlbum(db *sql.DB, id int) (*AlbumInfo, error) {
	var albumInfo AlbumInfo
	execSQL := `select id, title, artist, price from album where id = ?`
	if err := db.QueryRow(execSQL, id).Scan(&albumInfo.Id, &albumInfo.Title, &albumInfo.Artist, &albumInfo.Price); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("db query fail: %v", err)
	}

	return &albumInfo, nil
}

func InsertAlbum(db *sql.DB, albumInfo *AlbumInfo) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin tx error: %w", err)
	}
	defer func() {
		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	execSQL := "insert into album(title, artist, price) values(?, ?, ?)"
	result, err := tx.Exec(execSQL, albumInfo.Title, albumInfo.Artist, albumInfo.Price)
	if err != nil {
		return fmt.Errorf("insert fail %v", err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("rows affected fail %v", err)
	}
	fmt.Println("rows affected ", insertId)

	return nil
}

func UpdateAlbum(db *sql.DB, info AlbumInfo, query AlbumInfo) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin tx fail: %w", err)
	}
	defer func() {
		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	execSQL := `update album set title = ?, artist = ?, price = ? where id = ?`
	result, err := tx.Exec(execSQL, info.Title, info.Artist, info.Price, query.Id)
	if err != nil {
		return fmt.Errorf("update fail %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected fail %v", err)
	}
	fmt.Println("rows affected ", rowsAffected)
	return nil
}

func DeleteAlbum(db *sql.DB, id int) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin tx fail: %w", err)
	}
	defer func() {
		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	execSQL := "delete from album where id = ?"
	result, err := tx.Exec(execSQL, id)
	if err != nil {
		return fmt.Errorf("delete fail %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected fail %v", err)
	}
	fmt.Println("rows affected ", rowsAffected)
	return nil
}
