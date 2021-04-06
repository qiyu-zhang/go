package models

/*数据库相关处理函数*/
//建表与建立数据库用其他软件执行
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type user struct {
	name string
	pwd  string
}

//查询函数，用于查询输入账户，返回对应密码，无则返回空值
//输入:账号
//返回:密码

func QueryRow(name string) string {
	//连接数据库
	db, err := sql.Open("mysql", "用户:密码@tcp(127.0.0.1:3306)/数据库?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStr := "select pwd from users where name =?"
	var u string
	row := db.QueryRow(sqlStr, name)
	if row == nil {
		return ""
	}
	row.Scan(&u)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

//插入函数；用于将注册的账户插入数据库
func InsertRow(name string, pwd string) bool {
	//连接数据库
	db, err := sql.Open("mysql", "root:666234@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		return false
		log.Fatal(err)
	}
	defer db.Close()
	sqlStr := "insert into users(name, pwd) values(?,?)"

	_, err = db.Exec(sqlStr, name, pwd)
	fmt.Printf("insert success.\n")
	if err != nil {
		return false
	}
	return true
}
