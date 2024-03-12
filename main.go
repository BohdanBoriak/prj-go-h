package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"prj-go-h/domain"
	"sort"
	"strconv"
	"time"
)

const (
	rohalyky            = 10
	rohalykyPerQuestion = 5
)

var id uint64 = 1

func main() {
	fmt.Println("Вітаємо у грі \"РОГАЛИКИ\"")
	time.Sleep(1 * time.Second)

	users := showRate()
	for _, u := range users {
		if u.Id > id {
			id = u.Id
		}
	}
	id++

	for {
		menu()
		punct := ""
		fmt.Scan(&punct)

		switch punct {
		case "1":
			u := play()
			users := showRate()
			users = append(users, u)
			sortAndSave(users)
		case "2":
			fmt.Println("Список гравців:")
			users := showRate()
			for _, user := range users {
				fmt.Printf("Id: %v, Name: %s, Time: %v\n",
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
		fmt.Printf("%v + %v = ", x, y)

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

func sortAndSave(users []domain.User) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].Time < users[j].Time
	})

	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("Сталась помилка Т_Т: %s\n", err)
		return
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		fmt.Printf("Сталась помилка Т_Т: %s\n", err)
		return
	}
}

func showRate() []domain.User {
	info, err := os.Stat("users.json")
	if err != nil {
		fmt.Printf("Сталась помилка Т_Т: %s\n", err)
		return nil
	}

	var users []domain.User
	if info.Size() != 0 {
		file, err := os.Open("users.json")
		if err != nil {
			fmt.Printf("Сталась помилка Т_Т: %s\n", err)
			return nil
		}

		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		}(file)

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&users)
		if err != nil {
			fmt.Printf("Сталась помилка Т_Т: %s\n", err)
			return nil
		}
	}

	return users
}
