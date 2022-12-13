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

var userIds = [10]int{3, 4, 5}

var wg sync.WaitGroup

type priceRequest struct {
	Price string `json:"price"`
}

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
	e.POST("/report/short", ReportShort)
	e.POST("/report/long", ReportLong)
	e.POST("/report/cancel", ReportCancel)
}

func ReportShort(c echo.Context) error {
	pr := &priceRequest{}
	err := c.Bind(pr)

	if err != nil {
		return err
	}

	wg.Add(len(userIds))

	for i := range userIds {
		go callShort(pr.Price, userIds[i])
	}
	wg.Wait()

	return c.JSON(200, &response{})
}

func ReportLong(c echo.Context) error {
	pr := &priceRequest{}
	err := c.Bind(pr)

	if err != nil {
		return err
	}

	wg.Add(len(userIds))

	for i := range userIds {
		go callLong(pr.Price, userIds[i])
	}
	wg.Wait()

	return c.JSON(200, &response{})
}

func ReportCancel(c echo.Context) error {
	pr := &priceRequest{}
	err := c.Bind(pr)

	if err != nil {
		return err
	}

	wg.Add(len(userIds))

	for i := range userIds {
		go callCancel(pr.Price, userIds[i])
	}
	wg.Wait()

	return c.JSON(200, &response{})
}

func callShort(price string, i int) {
	postBody, _ := json.Marshal(map[string]string{
		"price": price,
	})
	responseBody := bytes.NewBuffer(postBody)
	_, err := http.Post("https://go-finance-robot.kadoopin.com/bot/pro/feature/stop-limit/short/"+fmt.Sprintf("%v", i), "application/json", responseBody)
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
	_, err := http.Post("https://go-finance-robot.kadoopin.com/bot/pro/feature/stop-limit/long/"+fmt.Sprintf("%v", i), "application/json", responseBody)
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
	_, err := http.Post("https://go-finance-robot.kadoopin.com/bot/pro/feature/stop-limit/cancel/"+fmt.Sprintf("%v", i), "application/json", responseBody)
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
