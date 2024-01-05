package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type User struct {
	id    int
	email string
}

func (u User) getData() string {
	out := ""

	for i := 0; i < 250; i++ {
		out += fmt.Sprintf("ID: %d | Email: %s\n", u.id, u.email)
	}
	return out
}

func generateUsers(count int) []User {
	users := make([]User, count)
	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@gmail.com", i+1),
		}
	}
	return users
}

func main() {
	users := generateUsers(10000)

	t := time.Now()

	wg := &sync.WaitGroup{}

	for _, user := range users {
		wg.Add(1)
		go writeFile(user, wg)
	}

	wg.Wait()

	fmt.Printf("Time Elapsed: %s\n", time.Since(t).String())

}

func writeFile(user User, wg *sync.WaitGroup) error {
	fmt.Printf("Writing file to: %d\n", user.id)
	fileName := fmt.Sprintf("logs/uid_%d", user.id)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)

	defer file.Close()

	if err != nil {
		return err
	}

	_, err = file.WriteString(user.getData())

	if err != nil {
		return err
	}

	wg.Done()
	return nil
}
