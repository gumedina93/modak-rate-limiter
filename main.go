package main

import (
	"fmt"
	"sync"
)

func main() {
	notifier := NewNotificationService()
	user1 := "user1"

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := notifier.SendNotification(News, user1)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	wg.Wait()
}
