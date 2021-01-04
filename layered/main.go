package main
import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/test/layered/handler"
	"github.com/test/layered/service"
	"github.com/test/layered/infra"

	"github.com/go-xorm/xorm"
	"xorm.io/core"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"

)


var engine *xorm.Engine

func NewDB() *xorm.Engine {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("NewDB:", err)
	}
	driverName := os.Getenv("DRIVER_NAME")
	user:= os.Getenv("USER_NAME")
	pass:= os.Getenv("PASS")
	tcp := os.Getenv("TCP")
	dbName := os.Getenv("DB_NAME")

	engine, err = xorm.NewEngine(driverName, user+":"+pass+"@"+tcp+"/"+dbName)
	if err != nil {
		log.Fatal("NewEngine: ", err)
	}
	//defer engine.Close()
	engine.ShowSQL(true)
	engine.SetMapper(core.GonicMapper{})

	return engine
}

func NewHandler(db *xorm.Engine) handler.Handler{
	log.Println("db:", db)

	// 明日はinfraでDBに接続する処理を書く
	// main ではすべての層にアクセスする（最初だけ
	repoVal := infra.NewSampleInfra(db)
	serviceVal := service.NewSampleService(repoVal)

	h:= handler.NewHandler(serviceVal)
	return h

}


func main() {
	db := NewDB()
	h := NewHandler(db)
	log.Println("sdlkjhfgjakljgh:", h)

	//handler

	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/aaaa", h.Sample)
	e.Logger.Fatal(e.Start(":8080"))
}

func test(c echo.Context) error{
	return c.JSON(http.StatusOK, "aaaa")
}