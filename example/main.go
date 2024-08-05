package main

import (
	"bytes"
	"compress/gzip"
	"database/sql"
	"log"
)

func main() {
	// DSN: Data Source Name
	dsn := "root:guanghua@tcp(127.0.0.1:13306)/yytemp?parseTime=true&interpolateParams=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查数据库连接是否成功
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	valueBlob := Gzipify("Hello World")
	sqlStatement := "INSERT INTO `meta_field` (`value_blob`) VALUES (?)"

	result, err := db.Exec(sqlStatement, valueBlob)
	if err != nil {
		log.Fatal(err)
	}

	// 获取新插入行的 ID
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Inserted user with ID: %d", id)
}

func Gzipify(value string) []byte {
	var b bytes.Buffer
	gzWriter := gzip.NewWriter(&b)
	gzWriter.Write([]byte(value))
	gzWriter.Close() // 关闭写入器以确保数据被刷新
	return b.Bytes()
}
