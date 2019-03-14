package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type SampleData struct {
	Hoge string  `json:"hoge"`
	Foo  MapData `json:"foo"`
}

type MapData struct {
	Bar  string `json:"bar"`
	Fuga FugaData
}

type FugaData struct {
	Piyo string `json:"piyo"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		data, err := ioutil.ReadFile("./data.json")
		if err != nil {
			fmt.Print(err)
		}

		var obj SampleData
		err = json.Unmarshal(data, &obj)
		if err != nil {
			fmt.Println("error:", err)
		}

		c.JSON(200, obj)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
