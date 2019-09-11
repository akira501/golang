package main

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

//--Hàm này viết để xem mình chạy câu lệnh nào (1)
type dbLogger struct{}

func (d dbLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d dbLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

//--Func main
func main() {
	//--Kết nối database
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123",
		Database: "Shop_database", // --Tên database
		Addr:     "localhost:5432",
	})
	//--In ra câu lệnh mình đã thực hiện trong lần chạy (1)
	db.AddQueryHook(dbLogger{})
	//Tạo nháp
	tx, err = db.Begin()
	if err != nil {
		return err
	}
	//--Đóng kết nối
	defer db.Close() //--Thực hiện ngay trước khi chương trình kết thúc

	//--Create Table
	type Book struct {
		Id        int32
		Name      string
		Author_id int32
		Publisher string
	}
	type Author struct {
		Id    int32
		Name  string
		Phone int32
		Email string
	}
	type Table struct {
		BookID     int32
		BookName   string
		AuthorName string
	}

	var product Book
	err := db.CreateTable(&product, &orm.CreateTableOptions{ //--Gọi đến hàm orm
		Temp:          false,
		FKConstraints: true,
		//--Nếu tạo lại bảng vẫn không bị lỗi
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}

	var authors Author
	err = db.CreateTable(&authors, &orm.CreateTableOptions{ //--Gọi đến hàm orm
		Temp:          false,
		FKConstraints: true,
		//--Nếu tạo lại bảng vẫn không bị lỗi
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("ALTER TABLE post ADD CONSTRAINT fk_author FOREIGN KEY (Author_id) REFERENCES author (Id)")

	//--Querry
	var newTable []Table
	_, err = db.Query(&newTable, `SELECT books.id as book_id, books.name as book_name, authors.name as author_name 
									FROM public.books 
									INNER JOIN public.authors 
									ON books.author_id=authors.id
									`)
	if err != nil {
		panic(err)
	}
	fmt.Println(newTable)
	//--Add Book
	// productAdd := []Author{
	// 	{
	// 		Id:    1,
	// 		Name:  "Nhật Minh",
	// 		Phone: 123,
	// 		Email: "123@gmail.com",
	// 	},
	// 	{
	// 		Id:    2,
	// 		Name:  "123@gmail.com",
	// 		Phone: 123,
	// 		Email: "",
	// 	},
	// 	{
	// 		Id:    3,
	// 		Name:  "Nhật Minh",
	// 		Phone: 123,
	// 		Email: "123@gmail.com",
	// 	},
	// }
	// productAdd := []Book{
	// 	{
	// 		Id:        1,
	// 		Name:      "Book1",
	// 		Author_id: 1,
	// 		Publisher: "Nhật Minh",
	// 	},
	// 	{
	// 		Id:        2,
	// 		Name:      "Book2",
	// 		Author_id: 2,
	// 		Publisher: "Nhật Minh",
	// 	},
	// 	{
	// 		Id:        3,
	// 		Name:      "Book3",
	// 		Author_id: 1,
	// 		Publisher: "Nhật Minh",
	// 	},
	// }

	//--Insert Product
	// err = db.Insert(&productAdd) //--Truyền con trỏ vào
	// if err != nil {
	//đặt câu rollback
	//	tx.Rollback()
	// 	panic(err)
	// }

}
