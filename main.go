package main

import (
	"fmt"
	"math/rand"
	"prj-go-h/domain"
	"strconv"
	"time"
)

const (
	rohalyky            = 10
	rohalykyPerQuestion = 5
)

var id uint64 = 1

func main() {
	var users []domain.User

	fmt.Println("Вітаємо у грі \"РОГАЛИКИ\"")
	time.Sleep(1 * time.Second)

	for {
		menu()
		punct := ""
		fmt.Scan(&punct)

		switch punct {
		case "1":
			u := play()
			users = append(users, u)
		case "2":
			fmt.Println("Список гравців:")
			for _, user := range users {
				fmt.Printf("Id: %v, Name: %s, Time: %v",
					user.Id,
					user.Name,
					user.Time)
			}
		case "3":
			return
		default:
			fmt.Println("Та що ви таке ввели!")
		}
	}
}

func menu() {
	fmt.Println("Виберіть пункт меню:")
	fmt.Println("1. Грати в рогалики!")
	fmt.Println("2. Рейтинг гравців...")
	fmt.Println("3. Залишити гру =(")
}

func play() domain.User {
	fmt.Println("Відповідай і вигравай!")

	moiiRohalyky := 0
	totalRohalyky := rohalyky
	now := time.Now()

	for totalRohalyky > 0 {
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
				totalRohalyky -= rohalykyPerQuestion
				fmt.Printf("Чудово, у тебе: %v рогаликів\n", moiiRohalyky)
			} else {
				fmt.Println("Яке розчарування... Спробуй ще!")
			}
		}
	}
	then := time.Now()
	timeSpent := then.Sub(now)

	fmt.Printf("ЧУДОВО! Ти зібрав усі рогалики за %v секунд!", timeSpent)

	fmt.Println("Введіть ім'я: ")
	name := ""
	fmt.Scan(&name)

	user := domain.User{
		Id:   id,
		Name: name,
		Time: timeSpent,
	}
	id++

	return user
}
