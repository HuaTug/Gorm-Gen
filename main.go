package main

import (
	"context"
	"fmt"
	"github.com/HuaTug/Gorm-Gen/dal"
	"github.com/HuaTug/Gorm-Gen/dal/model"
	"github.com/HuaTug/Gorm-Gen/dal/query"
)

// gen demo

// MySQLDSN MySQL data source name
const MySQLDSN = "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"

func init() {
	dal.DB = dal.ConnectDB(MySQLDSN).Debug()
}

func main() {
	// 设置默认DB对象
	query.SetDefault(dal.DB)

	// 创建
	/*	b1 := model.Book{
			Title:       "Gen of using",
			Author:      "NigTusg",
			PublishDate: time.Now(),
			Price:       100,
		}
		err := query.Book.WithContext(context.Background()).Create(&b1)
		//ToDo: 这些数据库的操作方法已经被提供
		if err != nil {
			fmt.Printf("create book fail, err:%v\n", err)
			return
		}*/
	//创建操作
	user := model.User{
		Name:  "XuZh",
		Phone: "15663568888",
	}
	err := query.User.WithContext(context.Background()).Create(&user)
	if err != nil {
		fmt.Printf("create user fail, err:%v\n", err)
		return
	}
	// 更新
	ret, err := query.Book.WithContext(context.Background()).
		Where(query.Book.ID.Eq(2)).
		Update(query.Book.Price, 200)
	if err != nil {
		fmt.Printf("update book fail, err:%v\n", err)
		return
	}

	fmt.Printf("RowsAffected:%v\n", ret.RowsAffected)

	users, err := query.User.WithContext(context.Background()).Where(query.User.ID.Eq(1)).
		Update(query.User.Phone, "15718568888")
	if err != nil {
		fmt.Printf("update book fail, err:%v\n", err)
		return
	}

	fmt.Printf("RowsAffected:%v\n", users.RowsAffected)
	// 查询
	book, err := query.Book.WithContext(context.Background()).First()
	// 也可以使用全局Q对象查询
	//book, err := query.Q.Book.WithContext(context.Background()).First()
	if err != nil {
		fmt.Printf("query book fail, err:%v\n", err)
		return
	}
	fmt.Printf("book:%v\n", book.Title)

	// 删除
	ret, err = query.Book.WithContext(context.Background()).Where(query.Book.ID.Eq(1)).Delete()
	if err != nil {
		fmt.Printf("delete book fail, err:%v\n", err)
		return
	}
	fmt.Printf("RowsAffected:%v\n", ret.RowsAffected)

	rets, err := query.Book.WithContext(context.Background()).GetBooksByAuthor("NigTusg")
	if err != nil {
		fmt.Printf("GetBooksByAuthor fail,err:%vn", err)
		return
	}
	for i, b := range rets {
		fmt.Printf("%d:%v\n", i, b)
	}
	//ToDo: 对filter这个切片进行遍历
	filter, err := query.Book.WithContext(context.Background()).FilterWriteColumn("price", "100")
	for i, b := range filter {
		fmt.Printf("The Result%d are %v\n", i, b)
	}
	b := &model.Book{
		Author: "Xui",
		ID:     2,
	}
	rets, err = query.Book.WithContext(context.Background()).Search(b)
	if err != nil {
		fmt.Printf("Search fail,err:%v\n", err)
		return
	}
	for i, b := range rets {
		fmt.Printf("%d:%v\n", i, b)
	}
}
