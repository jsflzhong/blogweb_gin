package database

import (
	"blogweb_gin/config"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

/*
该方法虽然不是内置的初始化方法,但是会在main入口那边的第一行被调用,等于是起到了初始化的作用.

有个地方需要注意的是，Query、Exec操作用法有些差异：
	a.Exec(update、insert、delete等无结果集返回的操作)调用完后会自动释放连接；
	b.Query(返回sql.Rows)则不会释放连接，调用完后仍然占有连接，它将连接的所属权转移给了sql.Rows，
		所以需要手动调用close归还连接，即使不用Rows也得调用rows.Close()，否则可能导致后续使用出错，如下的用法是错误的
*/
func InitMysql() {
	fmt.Println("[InitMysql]InitMysql....")
	if db == nil {
		db, _ = sql.Open(config.DriverName, buildDBString())
		CreateTableWithUser()
		CreateTableWithArticle()
		CreateTableWithAlbum()
	}
}

func buildDBString() string {
	dbPassword := flag.String("db_password", "nil", "The password of db connection")
	dbUrl := flag.String("db_url", "nil", "The url of db connection")
	flag.Parse()
	dbString := fmt.Sprintf("root:%s@tcp(%s:3306)/test1?charset=utf8", *dbPassword, *dbUrl)
	log.Println("@@@dbString:", dbString)
	return dbString
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`

	ModifyDB(sql)

}

//查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

//创建文章表
func CreateTableWithArticle() {
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

//--------图片--------
func CreateTableWithAlbum() {
	sql := `create table if not exists album(
		id int(4) primary key auto_increment not null,
		filepath varchar(255),
		filename varchar(64),
		status int(4),
		createtime int(10)
		);`
	ModifyDB(sql)
}
