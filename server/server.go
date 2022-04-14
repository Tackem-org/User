package server

import (
	"time"

	pb "github.com/Tackem-org/Global/pb/user"
)

var (
	Sessions []Session
)

type UserServer struct {
	pb.UnimplementedUserServer
}

type Session struct {
	UserID       uint64
	SessionToken string
	IPAddress    string
	ExpireTime   time.Time
}
