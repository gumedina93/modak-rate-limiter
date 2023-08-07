package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSendNotification_News_Success(t *testing.T) {
	// Given
	notificationSvc := NewNotificationService()
	user := "user1"
	notificationType := News

	// When
	err := notificationSvc.SendNotification(notificationType, user)

	// Then
	require.NoError(t, err)
}

func TestSendNotification_Status_Success(t *testing.T) {
	// Given
	notificationSvc := NewNotificationService()
	user := "user1"
	notificationType := Status

	// When
	err := notificationSvc.SendNotification(notificationType, user)

	// Then
	require.NoError(t, err)
}

func TestSendNotification_Marketing_Success(t *testing.T) {
	// Given
	notificationSvc := NewNotificationService()
	user := "user1"
	notificationType := Marketing

	// When
	err := notificationSvc.SendNotification(notificationType, user)

	// Then
	require.NoError(t, err)
}

func TestSendNotification_Exceeded_RateLimit(t *testing.T) {
	// Given
	notificationSvc := NewNotificationService()
	user := "user1"
	notificationType := News

	// When
	err := notificationSvc.SendNotification(notificationType, user)
	require.NoError(t, err)

	// Then
	err = notificationSvc.SendNotification(notificationType, user)
	require.Error(t, err)
	require.Equal(t, "rate limit exceeded for user user1 and notification news", err.Error())
}

func TestSendNotification_Invalid_Notification_Type(t *testing.T) {
	// Given
	notificationSvc := NewNotificationService()
	user := "user1"
	notificationType := "random-type"

	// When
	err := notificationSvc.SendNotification(NotificationType(notificationType), user)

	// Then
	require.Error(t, err)
	require.Equal(t, "unknown notification type provided", err.Error())
}
