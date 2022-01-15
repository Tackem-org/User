package server

import (
	"time"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/google/uuid"
)

func newSession(userID uint64, ipAddress string, expiryTime time.Duration) string {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[server.newSession(userID uint64, iPAddress string, expiryTime time.Duration) string] {userID=%d, ipAddress=%s, expiryTime=%s}", userID, ipAddress, helpers.DurationToString(expiryTime))
	new := session{
		UserID:       userID,
		SessionToken: uuid.New().String(),
		IPAddress:    ipAddress,
		ExpireTime:   time.Now().Add(expiryTime),
	}
	Sessions = append(Sessions, new)
	return new.SessionToken
}
