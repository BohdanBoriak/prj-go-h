package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var rohalyky int = 50

func main() {
	fmt.Println("Вітаємо у грі \"РОГАЛИКИ\"")
	time.Sleep(1 * time.Second)
	fmt.Println("Відповідай і вигравай!")

	rohalykyPerQuestion := 5
	moiiRohalyky := 0
	now := time.Now()

	for rohalyky > 0 {
		x, y := rand.Intn(100), rand.Intn(100)
		res := x + y

		fmt.Printf("%v + %v = ?\n", x, y)
		var ans string

		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		} else {
			if ansInt == res {
				moiiRohalyky += rohalykyPerQuestion
				rohalyky -= rohalykyPerQuestion
				fmt.Printf("Чудово, у тебе: %v рогаликів\n", moiiRohalyky)
			} else {
				fmt.Println("Яке розчарування... Спробуй ще!")
			}
		}
	}
	then := time.Now()
	fmt.Printf("ЧУДОВО! Ти зібрав усі рогалики за %v секунд!", then.Sub(now))
	time.Sleep(10 * time.Second)
}
