package main

import (
	"github.com/kataras/iris"
)

type newUsers struct {
	Name     string
	Password string
}

var Users = []newUsers{
	{
		Name:     "duy",
		Password: "123",
	},
	{
		Name:     "minh",
		Password: "123",
	},
}

func homePage(ctx iris.Context) {
	ctx.View("index.html")
}
func loginPage(ctx iris.Context) {
	ctx.View("login.html")
}
func registerPage(ctx iris.Context) {
	ctx.View("register.html")
}
func listPage(ctx iris.Context) {
	ctx.ViewData("Users", Users)
	ctx.View("listUsers.html")
}
func registerFunc(ctx iris.Context) {
	//Tạo slice data để hứng dữ liệu từ form Register
	data := newUsers{}
	err := ctx.ReadForm(&data)
	if err != nil && !iris.IsErrPath(err) /* see: https://github.com/kataras/iris/issues/1157 */ {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
	ctx.Writef("Đăng ký thành công! Xin chào ", data.Name)
	//thêm data vào slice
	Users = append(Users, data)
}
func loginFunc(ctx iris.Context) {
	//Tạo slice để hứng thông tin đăng nhập vào, sau đó so sánh với mảng Users
	dataUser := newUsers{}
	err := ctx.ReadForm(&dataUser)
	if err != nil && !iris.IsErrPath(err) /* see: https://github.com/kataras/iris/issues/1157 */ {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
	isRight := false

	//var userID int
	for _, v := range Users {
		if dataUser.Name == v.Name && dataUser.Password == v.Password {
			isRight = true
			break
		}
	}
	if isRight {
		ctx.Writef("Sign in successfully \nUser id: %d")
	} else {
		ctx.Writef("Failed to sign in")
	}
}

func main() {
	app := iris.New()
	tmpl := iris.HTML("./view", ".html")
	app.RegisterView(tmpl)
	// Mỗi lần chỉnh sửa file HTML thì chỉ cần refresh lại trang
	tmpl.Reload(true)

	app.Get("/", homePage)
	app.Get("/login", loginPage)
	app.Get("/register", registerPage)
	app.Get("/listUsers", listPage)

	app.Post("/register", registerFunc)
	app.Post("/login", loginFunc)
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
