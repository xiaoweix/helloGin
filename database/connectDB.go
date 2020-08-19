package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

const (
	userName = "root"
	password = "xxw2020"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "student"
)

func GetDataBase() *sql.DB {

	//postgres数据库
	//url := "postgres://postgres:123456@localhost/gosql?sslmode=disable"
	//log.Println(">>>> get database connection action start <<<<")
	//db, err := sql.Open("postgres", url)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 返回数据库连接
	//return db

	//mysql 数据库
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	fmt.Println(path)
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ := sql.Open("mysql", path)
	if DB == nil {
		log.Fatal("连接失败！")
		return nil
	}
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(10)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(5)
	//验证连接
	if err := DB.Ping(); err != nil{
		log.Fatal("opon database fail")
		return nil
	}
	return DB
}
