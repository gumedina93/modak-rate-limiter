package main

import (
	"fmt"
	"sync"
)

func main() {
	ns := NewNotificationService()
	user := "user@example.com"

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := ns.SendNotification(News, user)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	wg.Wait()
}
