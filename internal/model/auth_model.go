package model

import pb "github.com/vnnyx/auth-service/pb/auth"

type TokenDetails struct {
	AccessToken string
	AccessUUID  string
	AtExpires   int64
}

type JwtPayload struct {
	UserID     string
	Username   string
	Email      string
	AccessUUID string
}

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	AccessToken string
}

type User struct {
	ID       string
	Username string
}

func (l LoginResponse) ToGRPCResponse() *pb.Token {
	return &pb.Token{
		AccessToken: l.AccessToken,
	}
}
