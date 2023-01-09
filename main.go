package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type response struct {
}

var userIds = map[string]string{
	"3":  "https://go-finance-robot.kadoopin.com/bot",
	"4":  "https://go-finance-robot.kadoopin.com/bot",
	"5":  "https://go-finance-robot.kadoopin.com/bot",
	"6":  "https://go-finance-robot.kadoopin.com/bot",
	"14": "https://go-finance-robot.kadoopin.com/bot",
	"7":  "http://go-finance-robot-1.kadoopin.com/bot",
	"8":  "http://go-finance-robot-1.kadoopin.com/bot",
	"9":  "http://go-finance-robot-1.kadoopin.com/bot",
	"10": "http://go-finance-robot-1.kadoopin.com/bot",
	"11": "http://go-finance-robot-1.kadoopin.com/bot",
}

type priceRequest struct {
	Price json.RawMessage `json:"price"`
}

func main() {
	runHttpServer()
}

func runHttpServer() {
	fmt.Println("run server")
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
	fmt.Println("short report")
	pr := &priceRequest{}
	err := c.Bind(pr)

	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(pr.Price)

	for id, url := range userIds {
		go callShort(string(pr.Price), id, url)
	}

	return c.JSON(200, &response{})
}

func ReportLong(c echo.Context) error {
	fmt.Println("long report")
	pr := &priceRequest{}
	err := c.Bind(pr)

	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(pr.Price)

	for id, url := range userIds {
		go callLong(string(pr.Price), id, url)
	}

	return c.JSON(200, &response{})
}

func ReportCancel(c echo.Context) error {
	fmt.Println("cancel report")
	pr := &priceRequest{}
	err := c.Bind(pr)

	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(pr.Price)

	for id, url := range userIds {
		go callCancel(string(pr.Price), id, url)
	}

	return c.JSON(200, &response{})
}

func callShort(price string, id string, baseUrl string) {
	postBody, _ := json.Marshal(map[string]string{
		"price": price,
	})
	responseBody := bytes.NewBuffer(postBody)
	_, err := http.Post(fmt.Sprintf("%v", baseUrl)+"/pro/feature/stop-limit/short/"+fmt.Sprintf("%v", id), "application/json", responseBody)
	//Handle Error
	if err != nil {
		fmt.Println(err)
	}
	//
	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//
	//	}
	//}(resp.Body)

}

func callLong(price string, id string, baseUrl string) {
	postBody, _ := json.Marshal(map[string]string{
		"price": price,
	})
	responseBody := bytes.NewBuffer(postBody)
	_, err := http.Post(fmt.Sprintf("%v", baseUrl)+"/pro/feature/stop-limit/long/"+fmt.Sprintf("%v", id), "application/json", responseBody)
	//Handle Error
	if err != nil {
		fmt.Println(err)
	}

	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//
	//	}
	//}(resp.Body)

}

func callCancel(price string, id string, baseUrl string) {
	postBody, _ := json.Marshal(map[string]string{
		"price": price,
	})
	responseBody := bytes.NewBuffer(postBody)
	_, err := http.Post(fmt.Sprintf("%v", baseUrl)+"/pro/feature/stop-limit/cancel/"+fmt.Sprintf("%v", id), "application/json", responseBody)
	//Handle Error
	if err != nil {
		fmt.Println(err)
	}

	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//
	//	}
	//}(resp.Body)

}
