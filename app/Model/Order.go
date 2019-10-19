package Model

import (
	"fmt"
	"time"
)

type Order struct {
	Id        int       `json:"order_id"`
	OrderNo   string    `json:"order_no"`
	ProductId int64     `json:"product_id"`
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

func GetEveryTodayOrder(date time.Time) []Order {
	orders := []Order{}
	end := date.AddDate(0, 0, 1).Format("2006-01-02 00:00:00")
	db.Model(Order{}).
		Where("created_at > ?", date).
		Where("created_at < ?", end).
		Find(&orders)
	return orders
}

type p struct {
	ProductId int64
	Title     string
	Count     int64
}

func ReportCount(date time.Time) (int, int, []p) {
	// 今日点餐人数
	userCount := 0
	start := date
	end := date.AddDate(0, 0, 1).Format("2006-01-02 00:00:00")
	db.Model(Order{}).
		Where("created_at > ?", start).
		Where("created_at < ?", end).
		Count(&userCount)
	// 今日签到人数
	signCount := 0
	db.Model(Order{}).
		Where("created_at > ?", start).
		Where("created_at < ?", end).
		Where("status = ?", FINISHED).
		Count(&signCount)
	// 获取点餐product_id
	orders := GetEveryTodayOrder(date)
	var productIds []int64
	for _, order := range orders {
		productIds = append(productIds, order.ProductId)
	}
	fmt.Println(productIds)
	// 获取商品
	products := []Product{}
	db.Model(Product{}).Where("id in (?)", productIds).Find(&products)
	count := []p{}
	for _, product := range products {
		var temp int64 = 0
		for _, order := range orders {
			if order.ProductId == product.Id {
				temp = temp + 1
			}
		}
		count = append(count, p{
			ProductId: product.Id,
			Title:     product.Title,
			Count:     temp,
		})
	}

	return userCount, signCount, count
}
