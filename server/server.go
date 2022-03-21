package server

import (
	"time"

	pb "github.com/Tackem-org/Global/pb/user"
)

var (
	Sessions []session
)

type UserServer struct {
	pb.UnimplementedUserServer
}

type session struct {
	UserID       uint64
	SessionToken string
	IPAddress    string
	ExpireTime   time.Time
}
