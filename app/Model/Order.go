package Model

import (
	"time"
)

type Order struct {
	Id        int       `json:"order_id"`
	OrderNo   string    `json:"order_no"`
	ProductId string    `json:"product_id"`
	Uid       int       `json:"uid"`
	UserName  string    `json:"user_name"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	DEFAULT  = 0
	ORDER    = 1
	FINISHED = 2
	CANCEL   = 3
)

func GetTodayOrderListByUid(uid int) (orders []Order) {
	Map := make(map[string]interface{})
	Map["uid"] = uid
	Map["status"] = FINISHED
	db.Where(Map).Where("created_at > " + time.Now().Format("2006-01-02")).Find(&orders)
	return
}

func CountTodayOrderByUid(uid int) int64 {
	var rows int64
	db.Model(Order{}).
		Where("created_at > ?", time.Now().Format("2006-01-02 00:00:00")).
		Where("uid = ?", uid).
		Count(&rows)
	return rows
}

func SaveOrder(order Order) Order {
	db.Model(Product{}).Create(&order)
	return order
}

func DeleteTodayOrder(uid int) bool {
	order := Order{}
	db.Model(Order{}).Unscoped().
		Where("created_at > ?", time.Now().Format("2006-01-02 00:00:00")).
		Where("created_at < ?", time.Now().AddDate(0, 0, 1).Format("2006-01-02 00:00:00")).
		Where("uid = ?", uid).
		Delete(&order)
	return true
}
