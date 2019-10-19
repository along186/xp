package Model

import "time"

type Order struct {
	Id        int       `json:"order_id"`
	OrderNo   string    `json:"order_no"`
	ProductId string    `json:"product_id"`
	Uid       int       `json:"user_id"`
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



func SaveOrder(order Order) Order {
	db.Model(Product{}).Create(&order)
	return order
}
