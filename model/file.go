package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "time"
	"fmt"
)

type File struct {
	Id         int
	Name       string
	Url        string
	Created_at string
}

func NewFile() *File {
	return &File{}
}

// func Init_db() {
// 	db, err := sql.Open("mysql", "sxuesong:miami,wade123@tcp(123.206.79.49:3306)/fileserve?charset=utf8")
// 	checkErr(err)
// 	fmt.Println(db)
// }
func (this *File) Create() error {
	db, err := sql.Open("mysql", "sxuesong:miami,wade123@tcp(123.206.79.49:3306)/fileserve?charset=utf8")
	if err != nil {
		return fmt.Errorf("database open error,details:%v", err)
	}
	//插入数据
	stmt, _ := db.Prepare("INSERT files SET name=?,url=?,created_at=?")

	_, err = stmt.Exec(this.Name, this.Url, this.Created_at)
	if err != nil {
		return fmt.Errorf("database insert error,details:%v", err)
	}

	// id, _ := res.LastInsertId()
	db.Close()
	// fmt.Println(id)
	return nil
}

func Get() ([]File, error) {
	db, err := sql.Open("mysql", "sxuesong:miami,wade123@tcp(123.206.79.49:3306)/fileserve?charset=utf8")
	if err != nil {
		return nil, fmt.Errorf("database open error,details:%v", err)
	}
	rows, err := db.Query("SELECT id,name,url,created_at FROM files")
	if err != nil {
		return nil, fmt.Errorf("database query error,details:%v", err)
	}
	var objs []File
	for rows.Next() {
		var id int
		var name string
		var url string
		var created_at string
		rows.Scan(&id, &name, &url, &created_at)
		obj := File{id, name, url, created_at}
		// checkErr(err)
		objs = append(objs, obj)
	}
	db.Close()
	return objs, nil
}

func DestroyOne(id int) error {
	db, err := sql.Open("mysql", "sxuesong:miami,wade123@tcp(123.206.79.49:3306)/fileserve?charset=utf8")
	if err != nil {
		return fmt.Errorf("database open error,details:%v", err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("delete from files where id=?")
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	if _, err := res.RowsAffected(); err != nil {
		return err
	}
	return nil

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
