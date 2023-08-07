package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type NotificationType string

const (
	Status    NotificationType = "status"
	News                       = "news"
	Marketing                  = "marketing"
)

type NotificationService struct {
	mu         sync.Mutex
	rateLimits map[string]map[NotificationType]rateLimit
}

type rateLimit struct {
	count         int
	lastTimeStamp time.Time
	duration      time.Duration
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		rateLimits: make(map[string]map[NotificationType]rateLimit),
	}
}

func (ns *NotificationService) SendNotification(notificationType NotificationType, user string) error {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	if _, ok := ns.rateLimits[user]; !ok {
		ns.rateLimits[user] = make(map[NotificationType]rateLimit)
	}

	rLimit, ok := ns.rateLimits[user][notificationType]
	if !ok {
		duration, err := getDurationByNotificationType(notificationType)
		if err != nil {
			return err
		}
		rLimit.duration = duration
	}

	if isRateLimitExceeded(rLimit, notificationType) {
		return fmt.Errorf("rate limit exceeded for user %s and notification %s", user, notificationType)
	}

	if time.Since(rLimit.lastTimeStamp) >= rLimit.duration {
		rLimit.count = 0
	}

	rLimit.count++
	rLimit.lastTimeStamp = time.Now()
	ns.rateLimits[user][notificationType] = rLimit

	fmt.Printf("sending notification %s to user %s...\n", notificationType, user)
	return nil
}

func getDurationByNotificationType(notificationType NotificationType) (time.Duration, error) {
	switch notificationType {
	case Status:
		return 1 * time.Minute, nil
	case News:
		return 24 * time.Hour, nil
	case Marketing:
		return 1 * time.Hour, nil
	default:
		return 0, errors.New("unknown notification type provided")
	}
}

func getRateByNotificationType(notificationType NotificationType) int {
	switch notificationType {
	case Status:
		return 2
	case News:
		return 1
	case Marketing:
		return 3
	default:
		return 0
	}
}

func isRateLimitExceeded(rateLimit rateLimit, notificationType NotificationType) bool {
	return time.Since(rateLimit.lastTimeStamp) < rateLimit.duration &&
		rateLimit.count >= getRateByNotificationType(notificationType)
}
