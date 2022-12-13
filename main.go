package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

type response struct {
}

var userIds = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

var wg sync.WaitGroup

func main() {
	runHttpServer()
}

func runHttpServer() {
	e := echo.New()
	registerRoutes(e)

	address := fmt.Sprintf("%s:%s", "0.0.0.0", "8585")

	err := e.Start(address)
	if err != nil {
		e := fmt.Sprintf("Failed to start HTTP server")
		panic(e)
	}
}

func registerRoutes(e *echo.Echo) {
	e.GET("/report", Report)
}

func Report(c echo.Context) error {
	price := c.FormValue("price")

	wg.Add(len(userIds) * 3)

	for i := range userIds {
		go callShort(price, i)
		go callLong(price, i)
		go callCancel(price, i)
	}
	wg.Wait()

	return c.JSON(200, &response{})
}

func callShort(price string, i int) {
	postBody, _ := json.Marshal(map[string]string{
		"price": price,
	})
	responseBody := bytes.NewBuffer(postBody)
	_, err := http.Post("http://127.0.0.1:8000/testShort/"+fmt.Sprintf("%v", i), "application/json", responseBody)
	//Handle Error
	if err != nil {
	}
	//
	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//
	//	}
	//}(resp.Body)

	wg.Done()
}

func callLong(price string, i int) {
	postBody, _ := json.Marshal(map[string]string{
		"price": price,
	})
	responseBody := bytes.NewBuffer(postBody)
	_, err := http.Post("http://127.0.0.1:8000/testLong"+fmt.Sprintf("%v", i), "application/json", responseBody)
	//Handle Error
	if err != nil {
	}

	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//
	//	}
	//}(resp.Body)

	wg.Done()
}

func callCancel(price string, i int) {
	postBody, _ := json.Marshal(map[string]string{
		"price": price,
	})
	responseBody := bytes.NewBuffer(postBody)
	_, err := http.Post("http://127.0.0.1:8000/testCancel"+fmt.Sprintf("%v", i), "application/json", responseBody)
	//Handle Error
	if err != nil {
	}

	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//
	//	}
	//}(resp.Body)

	wg.Done()
}
