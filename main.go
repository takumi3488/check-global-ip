package main

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	req, _ := http.NewRequest("GET", "https://inet-ip.info", nil)
	req.Header.Set("User-Agent", "curl/8.6.0")
	client := &http.Client{}

	e.GET("/", func(c echo.Context) error {
		resp, err := client.Do(req)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		defer resp.Body.Close()
		byteArray, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, string(byteArray))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
