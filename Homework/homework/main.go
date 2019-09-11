package main

import "github.com/kataras/iris"

type newProduct struct {
	Id       string
	Image    string
	Rating   string
	Name     string
	NewPrice string
	OldPrice string
}

var Products = []newProduct{
	{
		Id: "1",
		//Chú ý đường dẫn tuyệt đối và đường dẫn tương đối
		Image:    "/resources/image/ic1.png",
		Rating:   "3.5",
		Name:     "Kem hoa quả",
		NewPrice: "100000",
		OldPrice: "120000",
	},
	{
		Id:       "2",
		Image:    "/resources/image/ic2.png",
		Rating:   "4",
		Name:     "Kem ngon",
		NewPrice: "100000",
		OldPrice: "0",
	},
	{
		Id:       "3",
		Image:    "/resources/image/ic3.png",
		Rating:   "3.5",
		Name:     "Kem hoa quả 1",
		NewPrice: "100000",
		OldPrice: "120000",
	},
	{
		Id:       "4",
		Image:    "/resources/image/ic4.png",
		Rating:   "4",
		Name:     "Kem ngon 1",
		NewPrice: "100000",
		OldPrice: "0",
	},
}

func main() {
	app := iris.New()
	// lấy các file html trong thư mục view
	tmpl := iris.HTML("./view", ".html")
	tmpl.Layout("layout.html")
	//Chỉnh sửa ko cần reset server
	tmpl.Reload(true)
	app.RegisterView(tmpl)
	//để tải đc ảnh
	app.HandleDir("/resources", "./resources")
	//trả về file shop.html
	app.Get("/san-pham", GetShop)
	app.Get("/tao-san-pham", GetCreateProductPage)
	app.Post("/tao-san-pham", CreateProduct)

	// Trả về file ảnh
	app.Get("/image/{file}", func(ctx iris.Context) {
		fileName := ctx.Params().Get("file")
		filePath := "/image/" + fileName
		ctx.ServeFile(filePath, true)
	})
	app.Get("/san-pham/{id}", func(ctx iris.Context) {
		newsID := ctx.Params().Get("id")
		var newsReturn newProduct
		for _, post := range Products {
			if post.Id == newsID {
				newsReturn = post
				break
			}
		}
		ctx.ViewData("newsReturn", newsReturn)
		ctx.View("product.html")
	})
	app.Run(iris.Addr(":8080"))
}

func GetShop(ctx iris.Context) {
	ctx.ViewData("Products", Products)
	ctx.View("shop.html")
}

func GetCreateProductPage(ctx iris.Context) {
	//Không sử dụng layout dùng chung.
	ctx.ViewLayout(iris.NoLayout)
	ctx.View("create.html")
}

func CreateProduct(ctx iris.Context) {
	//Tạo slice data để hứng dữ liệu từ form Register
	data := newProduct{}
	err := ctx.ReadForm(&data)
	if err != nil && !iris.IsErrPath(err) {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
	ctx.Writef("Tạo sản phẩm thành công!")
	//thêm data vào slice
	Products = append(Products, data)
}
