package server

import (
	yahooapi "dashboard/yahooApi"
	"encoding/json"
	"errors"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	address string
	port    string

	Database *Database

	is_running bool
}

func NewServer() (server *Server, err error) {
	server = &Server{}
	server.port = "7444"
	server.address = "127.0.0.1"

	server.Database = &Database{path: "./dashboard.db"}
	err = server.Database.InitDb()

	return
}

func (self *Server) Serve() {
	app := gin.Default()
	app.Use(ErrorHandler)
	app.POST("/api/add", self.ApiAddSymbol)
	app.POST("/api/search/:symbol", ApiSearchSymbol)
	app.POST("/api/notify/:symbol", self.ApiNotify)
	app.POST("/api/denotify/:symbol", self.ApiDenotify)
	app.GET("/api/tracked", self.ApiListTracked)
	app.Run(self.address + ":" + self.port)
}

func (self *Server) ApiAddSymbol(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		ctx.Error(err)
		ctx.Status(400)
		return
	}

	var symbol yahooapi.Symbol
	err = json.Unmarshal(body, &symbol)
	if err != nil {
		ctx.Error(err)
		ctx.Status(400)
		return
	}

	err = self.Database.AddSymbol(symbol)
	if err != nil {
		ctx.Error(err)
		ctx.Status(500)
		return
	}

	return
}

func (self *Server) ApiListTracked(ctx *gin.Context) {
	list, err := self.Database.Symbols()
	if err != nil {
		ctx.Error(err)
		log.Println(err)
		ctx.Status(500)
		return
	}

	json, err := json.Marshal(list)

	if err != nil {
		ctx.Error(err)
		log.Println(err)
		ctx.Status(500)
		return
	}

	ctx.Writer.Write(json)
}

func (self *Server) ApiNotify(ctx *gin.Context) {
	symbol := ctx.Params.ByName("symbol")
	query := `
	update stocks
	set notify = true
	where symbol = ?
	`
	tx, err := self.Database.Prepare(query)
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(500, err)
		return
	}

	res, err := tx.Exec(symbol)
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(500, err)
		return
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(500, err)
		return
	}

	if rows != 1 {
		ctx.AbortWithError(404, errors.New("Unknown Symbol"))
		return
	}
}

func (self *Server) ApiDenotify(ctx *gin.Context) {
	symbol := ctx.Params.ByName("symbol")
	query := `
	update stocks
	set notify = false
	where symbol = ?
	`
	tx, err := self.Database.Prepare(query)
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(500, err)
		return
	}

	res, err := tx.Exec(symbol)
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(500, err)
		return
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(500, err)
		return
	}

	if rows != 1 {
		ctx.AbortWithError(404, errors.New("Unknown Symbol"))
		return
	}
}

func ApiSearchSymbol(ctx *gin.Context) {
	result, err := yahooapi.SearchStockSymbol(ctx.Params.ByName("symbol"))
	if err != nil {
		ctx.AbortWithError(500, err)
	}

	json, err := json.Marshal(result)
	if err != nil {
		ctx.AbortWithError(500, err)
	}

	ctx.Writer.Write(json)
}

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	for _, err := range ctx.Errors {
		ctx.JSON(ctx.Writer.Status(), "Error: "+err.Error())
	}

}
