package model

import "gorm.io/gen"

// ToDo:  通过添加注释生成自定义方法 自定义SQL查询方法.

type Querier interface {
	//SELECT * FROM @@table WHERE id=@id 中的 @@table 是一个占位符，应该替换为实际的表名，而 @id 是另一个占位符，应该替换为实际的条件值

	// SELECT * FROM @@table WHERE id=@id
	GetByID(id int) (gen.T, error)

	// SELECT * FROM @@table WHERE id=@id
	GetByIDReturnMap(id int) (gen.M, error)

	// SELECT * FROM @@table WHERE author=@author
	GetBooksByAuthor(author string) ([]*gen.T, error)
}

type Searcher interface {
	//ToDo :这是一个根据id或者author进行数据查询的操作 注意{{end}}结尾的操作表示方法

	//select *from book
	//where publish_date is not null
	// {{if book!=nil}}
	//	 {{if book.ID>0}}
	//	 and id=@book.ID
	// {{else if book.Author !=""}}
	// 	 and author=@book.Author
	// {{end}}
	//{{end}}
	Search(book *gen.T) ([]*gen.T, error)
}

type Filter interface {
	// select * from @@table where @@column=@value
	FilterWriteColumn(column string, value string) ([]*gen.T, error)
}
