package main

import (
	"github.com/kataras/iris"
)

func homePage(ctx iris.Context) {
	ctx.View("index.html")
}
func contactPage(ctx iris.Context) {
	ctx.View("contact.html")
}
func imagePage(ctx iris.Context) {
	ctx.View("image.html")
}

type newsDetail struct {
	Id      string
	Image   string
	Title   string
	Content string
	Author  string
}

var listNews = []newsDetail{
	{
		Id:      "1",
		Image:   "img/anh1.jpg",
		Title:   "Anh 1",
		Content: "Ahihi",
		Author:  "Storm.J.M",
	},
	{
		Id:      "2",
		Image:   "img/anh2.jpg",
		Title:   "Anh 2",
		Content: "Ahihi",
		Author:  "Storm.J.M",
	},
	{
		Id:      "3",
		Image:   "img/anh3.jpg",
		Title:   "Anh 3",
		Content: "Ahihi",
		Author:  "Storm.J.M",
	},
}

func newsPage(ctx iris.Context) {
	ctx.ViewData("listNews", listNews)
	ctx.View("news.html")
}
func detailPage(ctx iris.Context) {
	ctx.ViewData("listNews", listNews)
	ctx.View("detail.html")
}

func main() {
	app := iris.New()
	tmpl := iris.HTML("./view", ".html")
	app.RegisterView(tmpl)
	// Mỗi lần chỉnh sửa file HTML thì chỉ cần refresh lại trang
	tmpl.Reload(true)
	// app.Get("/", func(ctx iris.Context) {
	// 	ctx.View("index.html")
	// })
	// app.Get("/contact", func(ctx iris.Context) {
	// 	ctx.View("contact.html")
	// })
	app.Get("/", homePage)
	app.Get("/lienhe", contactPage)
	app.Get("/anh", imagePage)
	app.Get("/news", newsPage)

	// Trả về file style.css
	app.Get("/css/{file}", func(ctx iris.Context) {
		fileName := ctx.Params().Get("file")
		filePath := "css/" + fileName
		ctx.ServeFile(filePath, true)
	})
	// Trả về file ảnh
	app.Get("/img/{file}", func(ctx iris.Context) {
		fileName := ctx.Params().Get("file")
		filePath := "img/" + fileName
		ctx.ServeFile(filePath, true)
	})
	app.Get("/detail/{id}", func(ctx iris.Context) {
		newsID := ctx.Params().Get("id")
		var newsReturn newsDetail
		for _, post := range listNews {
			if post.Id == newsID {
				newsReturn = post
				break
			}
		}
		ctx.ViewData("newsReturn", newsReturn)
		ctx.View("detail.html")
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

// /////////////////
// func SigninHandler(ctx iris.Context) {
// 	user := User{}
// 	err := ctx.ReadForm(&user)
// 	if err != nil && !iris.IsErrPath(err) /* see: https://github.com/kataras/iris/issues/1157 */ {
// 		ctx.StatusCode(iris.StatusInternalServerError)
// 		ctx.WriteString(err.Error())
// 	}

// 	isRight := false
// 	var userID int
// 	for _, v := range users {
// 		if user.Name == v.Name && user.Password == v.Password {
// 			isRight = true
// 			userID = v.ID
// 			break
// 		}
// 	}
// 	if isRight {
// 		ctx.Writef("Sign in successfully \nUser id: %d", userID)
// 	} else {
// 		ctx.Writef("Failed to sign in")
// 	}
// }
