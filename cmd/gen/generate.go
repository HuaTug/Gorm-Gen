package main

// gorm gen configure

import (
	"fmt"
	"github.com/HuaTug/Gorm-Gen/dal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gorm.io/gen"
)

const MySQLDSN = "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}

func main() {
	// 指定生成代码的具体相对目录(相对当前文件)，默认为：./query
	// 默认生成需要使用WithContext之后才可以查询的代码，但可以通过设置gen.WithoutContext禁用该模式
	g := gen.NewGenerator(gen.Config{
		// 默认会在 OutPath 目录生成CRUD代码，并且同目录下生成 model 包
		// 所以OutPath最终package不能设置为model，在有数据库表同步的情况下会产生冲突
		// 若一定要使用可以通过ModelPkgPath单独指定model package的名称
		OutPath: "../../dal/query",
		/* ModelPkgPath: "dal/model"*/

		// gen.WithoutContext：禁用WithContext模式
		// gen.WithDefaultQuery：生成一个全局Query对象Q
		// gen.WithQueryInterface：生成Query接口
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// 通常复用项目中已有的SQL连接配置db(*gorm.DB)
	// 非必需，但如果需要复用连接时的gorm.Config或需要连接数据库同步表信息则必须设置
	g.UseDB(connectDB(MySQLDSN))

	// 从连接的数据库为所有表生成Model结构体和CRUD代码
	// 也可以手动指定需要生成代码的数据表
	g.ApplyBasic(g.GenerateAllTable()...)
	/*
		ToDo:代码中的 g.ApplyInterface 函数是 Gen 框架提供的一个方法，它接受一个接口类型和一个函数作为参数。在这里，您传递了一个匿名函数给 g.ApplyInterface，该函数具有一个 model.Querier 参数，函数体为空。
			同时，使用 g.GenerateModel("book") 生成了一个与 "book" 相关的 SQL 语句。
	*/
	/*
		ToDo； 在 Gen 框架中，您可以定义接口，并使用 g.ApplyInterface 方法将接口应用于特定的数据库操作。通过这种方式，您可以在 Gen 框架生成的代码中插入自定义的逻辑，并与数据库交互。
		       g.ApplyInterface 方法接受两个参数：
		       第一个参数是一个函数或闭包，该函数或闭包需要实现您定义的接口。
		       第二个参数是一个可选参数，用于传递给接口方法的参数。
	*/
	g.GenerateModelAs("user", "User")

	g.ApplyInterface(func(querier model.Querier) {}, g.GenerateModel("book"))
	g.ApplyInterface(func(filter model.Filter) {}, g.GenerateModel("book"))
	g.ApplyInterface(func(searcher model.Searcher) {}, g.GenerateModel("book"))

	// 执行并生成代码
	g.Execute()
}
