package main

import "github.com/labstack/echo/v4"
//import "net/http"
import "math/rand"
import "os"
import "time"

//http://google.com/search?q="sometext"

type quote struct {
	Title string
	Author string
}

var quotes []quote = []quote{
	{"Learn to Lead", "Sai Vidya Campus"},
	{"jkljklj", "Sai Vidya Campus"},
	{"m m m, m, m, ", "Sai Vidya Campus"},
	{"jkhkjhjhkjh", "Sai Vidya Campus"},
}
func main(){


	rand.Seed(time.Now().Unix())
	api := echo.New()
	api.GET("/quotes",getQuotes)
	api.GET("/quotes/random",getRandomQuote)

	port := os.Getenv("PORT")
	if port ==""{
		port="32445"
	}

	//api.POST("/home",anotherHandler)
	//api.PUT("/home",hello)
	//api.GET()
	api.Start(":"+port)
}

 /*func hello(c echo.Context) error{
	return c.JSON(200,"nnnnnn")
}

func anotherHandler(c echo.Context) error{
	return c.NoContent(http.StatusOK)
}*/

func getQuotes(c echo.Context) error{
	return c.JSON(200,quotes)
}

func getRandomQuote(c echo.Context) error{
	index :=rand.Intn((len(quotes)))
	return c.JSON(200,quotes[index])
}