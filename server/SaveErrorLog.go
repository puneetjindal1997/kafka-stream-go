package main

import (
	"context"
	"disastermanagement/constants"
	"disastermanagement/database"
	"disastermanagement/pb"
	"errors"
	"log"
)

func (s *server) SaveUserLogs(ctx context.Context, req *pb.UserErrorRequest) (*pb.UserErrorResponse, error) {
	log.Println("----------------------------------- in server -------------------------")

	dbErr := database.Mgr.InsertUserLog(req)
	if dbErr != nil {
		return &pb.UserErrorResponse{}, dbErr
	}
	return &pb.UserErrorResponse{Status: true}, nil
}

func (s *server) GetUserLogs(ctx context.Context, req *pb.Null) (*pb.UserLogResp, error) {
	log.Println("----------------------------------- in server -------------------------")

	dbResp, dbErr := database.Mgr.GetUserLogs(constants.LogType(1).String())
	if dbErr != nil {
		return &pb.UserLogResp{}, errors.New(constants.InternalServerError)
	}
	return &pb.UserLogResp{UserResp: dbResp}, nil
}
