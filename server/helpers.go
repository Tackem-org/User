package server

import (
	"time"

	"github.com/google/uuid"
)

func newSession(userID uint64, ipAddress string, expiryTime time.Duration) string {
	new := session{
		UserID:       userID,
		SessionToken: uuid.New().String(),
		IPAddress:    ipAddress,
		ExpireTime:   time.Now().Add(expiryTime),
	}
	Sessions = append(Sessions, new)
	return new.SessionToken
}
