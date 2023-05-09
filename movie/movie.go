package movie

import (
	"fmt"
)

func init() {
	fmt.Println("init: movie")
}


// ขึ้นต้นด้วยตัวใหญ่ = public
func Review(name string, rating float64) {
	fmt.Printf("!!! I reviewed %s and it's rating is %f\n", name, rating)
}
