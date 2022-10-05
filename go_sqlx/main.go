package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	initDB()
	defer db.Close()
	// find(1)
	// user := User{
	// 	Name: "张三",
	// 	Age:  25,
	// }
	// deleteRow(3)
	// insertDemo()
	// namedQuery()
	// transactionDemo()

	user1 := User{Name: "张三", Age: 20}
	user2 := User{Name: "李四", Age: 20}
	user3 := User{Name: "王五", Age: 20}
	user4 := User{Name: "赵六", Age: 20}
	user5 := User{Name: "钱七", Age: 20}
	userSlice := []interface{}{user1, user2, user3, user4, user5}
	err := BatchInsertUsers2(userSlice)
	if err != nil {
		fmt.Printf("err is %v\n", err)
	}
}

// 全局变量db，并发安全的对象
var db *sqlx.DB

type User struct {
	Id   uint   `db:"id"`
	Name string `db:"name"`
	Age  uint   `db:"age"`
}

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/testabc?charset=utf8mb4&parseTime=true"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err: %v\n", err)
		return
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return
}

func find(id int) {
	sqlStr := "select id,name,age from user where id = ?"
	var u User
	err := db.Get(&u, sqlStr, id)
	if err != nil {
		fmt.Println("查询单行失败,err:", err)
		return
	}
	fmt.Println(u)
}

func insertRow(user User) {
	sqlStr := "insert into user(name,age) values (?,?)"
	name := user.Name
	age := user.Age
	ret, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", err)
		return
	}
	id, err := ret.LastInsertId() // 新插入数据的Id
	if err != nil {
		fmt.Printf("get last insert id failed, err: %v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", id)
}

func updateRow(user User, id int) {
	sqlStr := "update user set name = ?, age = ? where id = ?"
	name := user.Name
	age := user.Age
	ret, err := db.Exec(sqlStr, name, age, id)
	if err != nil {
		fmt.Printf("更新失败,id:%d\n", id)
		return
	}
	// 受影响的行数
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("受影响行数获取失败,%d\n", rows)
		return
	}
	fmt.Printf("更新id: %d 成功，受影响行数：%d\n", id, rows)
}

func deleteRow(id int) {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("删除失败, err:%d\n", err)
		return
	}
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("id: %d 获取受影响的行数失败,err: %v\n", id, err)
		return
	}
	fmt.Printf("删除 id: %d 成功，受影响的行数: %d\n", id, rows)
}

func insertDemo() (err error) {
	sqlStr := "insert into user (name, age) values (:name, :age)"
	res, err := db.NamedExec(sqlStr, map[string]any{
		"name": "张三",
		"age":  28,
	})
	rows, err := res.RowsAffected()
	fmt.Printf("插入成功，rows:%d\n", rows)
	return
}

func namedQuery() {
	sqlStr := "select * from user where name = :name"
	// 使用map做命名查询
	rows, err := db.NamedQuery(sqlStr, map[string]any{"name": "贺敬宇"})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err: %v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err %v\n", err)
			continue
		}
		fmt.Printf("user: %v\n", u)
	}
}

func transactionDemo() (err error) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin trans failed, err: %v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() // err is non-nil;don't change it
		} else {
			err = tx.Commit() // err is nil;if commit returns err, update err
			fmt.Println("commit")
		}
	}()
	sqlStr1 := "update user set age = 20 where id = ?"
	res, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	sqlStr2 := "update user set age = 50 where id = ?"
	res, err = tx.Exec(sqlStr2, 5)
	if err != nil {
		return err
	}
	rows, err = res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("exec sqlStr2 failed")
	}
	return err
}

// BatchInsertUsers 自行构造批量插入的语句，比较麻烦，且不通用
func BatchInsertUsers(users []*User) error {
	// 存放 (?, ?) 的slice
	valueStrings := make([]string, 0, len(users))
	// 存放values的slice
	valueArgs := make([]interface{}, 0, len(users)*2)
	// 遍历users准备相关数据
	for _, u := range users {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	// 自行拼接要执行的具体语句
	stmt := fmt.Sprintf("INSERT INTO user (name, age) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	return err
}

func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

// BatchInsertUsers2 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
func BatchInsertUsers2(users []interface{}) error {
	query, args, err := sqlx.In(
		"INSERT INTO user (name, age) VALUES (?),(?), (?), (?), (?)", // 注意占位符要和你传入的参数的个数一致
		users..., // 如果arg实现了 driver.Value, sqlx.In 会通过调用 Value()来展开它
	)
	if err != nil {
		return err
	}
	fmt.Println(query) // 查看生成的querystring
	fmt.Println(args)  // 查看生成的args
	_, err = db.Exec(query, args...)
	return err
}

// QueryByIDs 根据给定ID查询
func QueryByIDs(ids []int) (users []User, err error) {
	// 动态填充id
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?)", ids)
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)

	err = db.Select(&users, query, args...)
	return
}
