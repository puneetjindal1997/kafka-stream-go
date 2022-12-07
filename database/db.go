package database

import "disastermanagement/pb"

const (
	UserLog = "user_log"
)

/*
 *	Function to save the log data
 *
 *	response error
 */
func (mgr *manager) InsertUserLog(log *pb.UserErrorRequest) error {
	resp := mgr.connection.Table(UserLog).Create(&log)
	return resp.Error
}

/*
 *	Function to get all the logs by type filter
 *
 *	response []*pb.UserErrorRequest, error
 */
func (mgr *manager) GetUserLogs(filterBy string) (logs []*pb.UserErrorRequest, err error) {
	resp := mgr.connection.Table(UserLog).Where("type=?", filterBy).Offset(0).Limit(10).Order("updated_at desc").Find(&logs)
	err = resp.Error
	return logs, err
}
