package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
	DATE_FORMAT = "2006-01-02"
)

type User struct {
	ID       uint      `gorm:"column:id;primaryKey"`
	Name     string    `gorm:"type:varchar(20);not null" json:"name" binding:"required"`
	Age      uint      `gorm:"type:tinyint;not null" json:"age"`
	Birthday time.Time `gorm:"type:date" json:"birthday"`
	Address  string    `gorm:"type:varchar(100)" json:"address"`
}

func main() {
	InitDb()
	var user User
	DB.First(&user)
	fmt.Printf("%v\n", user)
	fmt.Println("=============")
	// 清空，避免前后影响
	user = User{}
	DB.Find(&user, 2)
	// DB.Where("id=?", 2).Find(&users)
	fmt.Printf("%v\n", user)
	fmt.Println("==============")
	var users []User
	DB.Find(&users, []int{1, 2})
	// DB.Where("id in", []int{1,2}).Find(&users)
	for _, u := range users {
		fmt.Printf("%v\n", u)
	}
	fmt.Println("===============")
	users = []User{}
	DB.Where("address=?", "湖北省").Find(&users)
	for _, u := range users {
		fmt.Printf("%v\n", u)
	}
	fmt.Println("===============")
	user = User{}
	users = []User{}
	DB.Where(map[string]any{"address": "湖北省"}).Find(&users)
	for _, u := range users {
		fmt.Printf("%v\n", u)
	}
	fmt.Println("================")
	users = []User{}
	DB.Order("age").Limit(1).Offset(1).Find(&users)
	for _, u := range users {
		fmt.Printf("%v\n", u)
	}
	fmt.Println("==================")
	users = []User{}
	DB.Select("name", "age", "address").Take(&users)
	for _, u := range users {
		fmt.Printf("%v\n", u)
	}
	fmt.Println("=================")

	// updateRow()
	// Insert()
	Delete()
}

// 连接数据库
func InitDb() (err error) {
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 191,
	}), &gorm.Config{
		SkipDefaultTransaction: false, // 跳过默认事务
		NamingStrategy: schema.NamingStrategy{ // 表命名规则
			TablePrefix:   "t_", // 表前缀
			SingularTable: true, // 表名是否复数带s
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用逻辑外键
	})
	return
}

// AutoMigrate 会创建表、缺失的外键、约束、列和索引。
// 如果大小、精度、是否为空可以更改，则 AutoMigrate 会改变列的类型。
// 出于保护您数据的目的，它 不会 删除未使用的列
func MigrateTable() (err error) {
	err = DB.AutoMigrate(&User{})
	return
}

// 插入记录
func InsertUser(user User) (err error) {
	result := DB.Create(&user)
	err = result.Error
	fmt.Println("插入用户的数量", result.RowsAffected)
	return
}

func Insert() (err error) {
	// 插入一条
	user := User{Name: "李四", Address: "湖南省", Age: 30, Birthday: time.Now()}
	DB.Create(&user)

	// 插入多条
	users := []User{{Name: "王五", Address: "陕西省", Age: 29, Birthday: time.Now()}, {Name: "赵六", Address: "河北省", Age: 21, Birthday: time.Now()}}
	DB.Create(users)

	// 量太大时分批插入
	users = []User{{Name: "王五五", Address: "陕西省", Age: 29, Birthday: time.Now()}, {Name: "赵六六", Address: "河北省", Age: 21, Birthday: time.Now()}, {Name: "钱七七", Address: "海南省", Age: 26, Birthday: time.Now()}}
	DB.CreateInBatches(users, 2)
	return
}

func GetOne() {
	var user User
	DB.First(&user)
	fmt.Printf("user is %v\n", user)
}

func updateRow() (err error) {
	// 根据where更新一列
	DB.Model(&User{}).Where("address = ?", "浙江省").Update("age", 26)
	// 更新多列
	DB.Model(&User{}).Where("address = ?", "浙江省").Updates(map[string]any{"name": "张三三", "age": 24})
	// 更新特定记录
	user := User{ID: 3, Name: "张三四"}
	// 相当于 where id = 3 and name = "张三三"
	DB.Model(&user).Where("name=?", "张三三").Updates(map[string]any{"name": "张三"})
	return
}

func Delete() (err error) {
	// where 删除
	DB.Where("address in ?", []string{"陕西省", "河北省"}).Delete(&User{})
	// 主键删除
	DB.Delete(&User{}, 10)
	DB.Delete(&User{}, []int{3, 4, 5})
	return
}
