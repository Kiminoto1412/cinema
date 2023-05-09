package main

import (
	"fmt"

	"github.com/Kiminoto1412/cinema/movie"
	"github.com/Kiminoto1412/cinema/ticket"
)

// เป็นbuit in functionืี่มีให้ในทุกfile ของ go ซึ่งจะเรียกใช้ตั้งแต่เริ่มเมื่อ package นั้นถูกเรียกใช้
func init() {
	fmt.Println("init: main")
}

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
