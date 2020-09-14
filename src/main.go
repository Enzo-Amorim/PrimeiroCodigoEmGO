package main

import (
	"fmt"
	"newsfeeder/src/model/newsfeed"
)

func main() {
	/*r := gin.Default()

	r.GET("/ping", controller.PingGet)

	r.Run()*/

	feed := newsfeed.New()

	fmt.Println(feed)

	feed.Add(newsfeed.Item{"Hello", "How ya' doing mate/"})

	feed.Add(newsfeed.Item{"Coe", "Como tu ta?"})

	fmt.Println(feed)
}
