package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// import "fmt"

type HomeController struct{}

func (h *HomeController) Index(c echo.Context) error {
	data := map[string]interface{}{
		"Title":   "Trang học Golang",
		"Message": "Chào mừng bạn đến với hành trình học Go!",
	}
	// fmt.Printf("%#v\n", data)
	return c.Render(http.StatusOK, "index.html", data)
}

func (h *HomeController) HelloGolang(c echo.Context) error {
	data := map[string]interface{}{
		"Title":   "Hello Golang",
		"Message": "Giới thiệu Golang và cấu trúc project.",
	}
	return c.Render(http.StatusOK, "hello-golang.html", data)
}

func (h *HomeController) RestAPI(c echo.Context) error {
	data := map[string]interface{}{
		"Title":   "REST API với Echo",
		"Message": "Cách xây dựng REST API với Echo framework.",
	}
	return c.Render(http.StatusOK, "rest-api-echo.html", data)
}

func (h *HomeController) DBConnection(c echo.Context) error {
	dsn := "vietvv:pass@tcp(mysql-e54fbe6-rooney-e8ac.b.aivencloud.com:28586)/mastersns2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return c.String(http.StatusInternalServerError, "Kết nối DB lỗi: "+err.Error())
	}

	// Model ví dụ
	type Admin struct {
		ID   uint
		Name string
	}

	var admins []Admin
	if err := db.Find(&admins).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Lỗi truy vấn: "+err.Error())
	}

	data := map[string]interface{}{
		"Title":   "Kết nối Database",
		"Message": "Sử dụng PostgreSQL hoặc MySQL với Go.",
		"Admins":  admins,
	}

	return c.Render(http.StatusOK, "db-connection.html", data)
}

func (h *HomeController) JWTAuth(c echo.Context) error {
	data := map[string]interface{}{
		"Title":   "Xác thực JWT",
		"Message": "Hướng dẫn sử dụng middleware và xác thực bằng JWT.",
	}
	return c.Render(http.StatusOK, "jwt-auth.html", data)
}

func (h *HomeController) MVCPattern(c echo.Context) error {
	data := map[string]interface{}{
		"Title":   "Mô hình MVC",
		"Message": "Tổ chức code rõ ràng với mô hình MVC.",
	}
	return c.Render(http.StatusOK, "mvc-pattern.html", data)
}

func (h *HomeController) DeployHTTPS(c echo.Context) error {
	data := map[string]interface{}{
		"Title":   "Triển khai HTTPS",
		"Message": "Hướng dẫn deploy ứng dụng Golang với HTTPS.",
	}
	return c.Render(http.StatusOK, "deploy-https.html", data)
}
