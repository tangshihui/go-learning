package main

import (
	"bufio"
	"bytes"
	"example/go-learning/gormTest/handler"
	"example/go-learning/gormTest/repo"
	"github.com/labstack/echo-contrib/prometheus"
	"io"
	"log"
	"net"

	"github.com/labstack/echo/v4"
	"io/ioutil"

	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func ping(c echo.Context) error {
	//panic("Failed ping")
	return c.String(http.StatusOK, "Pong")
}

func hello(c echo.Context) error {
	//panic("Failed ping")
	name := c.Param("name")
	return c.String(http.StatusOK, "Hello,"+name)
}

type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *bodyDumpResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *bodyDumpResponseWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

func (w *bodyDumpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}

func selfMiddle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Request
		reqBody := []byte{}
		if c.Request().Body != nil { // Read
			reqBody, _ = ioutil.ReadAll(c.Request().Body)
		}
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody)) // Reset

		// Response
		resBody := new(bytes.Buffer)
		mw := io.MultiWriter(c.Response().Writer, resBody)
		writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
		c.Response().Writer = writer

		if err := next(c); err != nil {
			c.Error(err)
		}
		if c.Response().Status >= 300 {
			log.Printf("bad status %#v with req: %#v resp: %#v", c.Response().Status, string(reqBody), resBody.String())
		}
		return nil
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(selfMiddle)
	prometheus.NewPrometheus("go-learning", nil).Use(e)

	e.GET("/ping", ping)
	e.GET("/hello/:name", hello)

	userRepo := repo.NewUserRepo(repo.DB)
	//defer repo.DB.clo
	userHandler := handler.NewUserHandler(userRepo)
	e.GET("/users", userHandler.ListAllUsers)
	e.GET("/user/:id", userHandler.QueryUser)
	e.DELETE("/user/:id", userHandler.DeleteUser)
	e.PUT("/user", userHandler.SaveUser)
	e.POST("/user", userHandler.UpdateUser)

	creditCardRepo := repo.NewCreditCardRepo(repo.DB)
	creditCardHandler := handler.NewCreditCardHandler(creditCardRepo)
	e.GET("/credit-card", creditCardHandler.ListAll)
	e.GET("/credit-card/:id", creditCardHandler.QueryById)
	e.GET("/credit-card/:userId", creditCardHandler.QueryByUserId)
	e.POST("/credit-card", creditCardHandler.SaveCreditCard)

	e.Start(":8080")
}
