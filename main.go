package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
)

type response struct {
}

var userIds = map[string]string{}

var userIdsV2 = map[string]string{
	"18": "https://go-finance-robot.kadoopin.com/bot",
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
	e.POST("/v2/report/short", ReportShortV2)
	e.POST("/v2/report/long", ReportLongV2)
	e.POST("/v2/report/cancel", ReportCancelV2)
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

func ReportShortV2(c echo.Context) error {
	fmt.Println("short report v2")
	pr := &priceRequest{}
	err := c.Bind(pr)

	if err != nil {
		log.Println(err)
		return err
	}

	for id, url := range userIdsV2 {
		go callShort(string(pr.Price), id, url)
	}

	return c.JSON(200, &response{})
}

func ReportLongV2(c echo.Context) error {
	fmt.Println("long report v2")
	pr := &priceRequest{}
	err := c.Bind(pr)

	if err != nil {
		log.Println(err)
		return err
	}

	for id, url := range userIdsV2 {
		go callLong(string(pr.Price), id, url)
	}

	return c.JSON(200, &response{})
}

func ReportCancelV2(c echo.Context) error {
	fmt.Println("cancel report v2")
	pr := &priceRequest{}
	err := c.Bind(pr)

	if err != nil {
		log.Println(err)
		return err
	}

	for id, url := range userIdsV2 {
		go callCancel(string(pr.Price), id, url)
	}

	return c.JSON(200, &response{})
}

func callShort(price string, id string, baseUrl string) {
	postBody, _ := json.Marshal(map[string]string{
		"price": price,
	})
	responseBody := bytes.NewBuffer(postBody)
	response, err := http.Post(fmt.Sprintf("%v", baseUrl)+"/pro/feature/parcham/stop-limit/short/"+fmt.Sprintf("%v", id), "application/json", responseBody)
	//Handle Error
	if err != nil {
		fmt.Println(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

}

func callLong(price string, id string, baseUrl string) {
	postBody, _ := json.Marshal(map[string]string{
		"price": price,
	})
	responseBody := bytes.NewBuffer(postBody)
	_, err := http.Post(fmt.Sprintf("%v", baseUrl)+"/pro/feature/parcham/stop-limit/long/"+fmt.Sprintf("%v", id), "application/json", responseBody)
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
	_, err := http.Post(fmt.Sprintf("%v", baseUrl)+"/pro/feature/parcham/stop-limit/cancel/"+fmt.Sprintf("%v", id), "application/json", responseBody)
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
