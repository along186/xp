package Model

import "time"

type Order struct {
	Id int	`json:"order_id"`
	OrderNo string `json:"order_no"`
	Uid string `json:"user_id"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

const (
	DEFAULT = 0
	FINISHED = 1
	CANCEL = 2
)
func GetTodayOrderListByUid(uid int)  (orders []Order){
	Map := make(map[string]interface{})
	Map["uid"] = uid
	Map["status"] = FINISHED
	db.Where(Map).Where("created_at > " + time.Now().Format("2006-01-02")).Find(&orders)
	return
}