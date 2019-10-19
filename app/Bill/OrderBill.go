package Bill

import (
	"errors"
	"time"

	"xp/app/Model"
)

func CheckTodaykHasOrdered(uid int) bool {
	orderList := Model.GetTodayOrderListByUid(uid)
	if len(orderList) > 0 {
		return true
	}

	return false
}

func SaveOrder(order Model.Order) (Model.Order, error) {
	count := Model.CountTodayOrderByUid(order.Uid)
	if count > 0 {
		return order, errors.New("请先取消，然后再点餐")
	}
	return Model.SaveOrder(order), nil
}

func DeleteOrderByUid(uid int) bool {
	return Model.DeleteTodayOrder(uid)
}

func GetTodayEveryBodyOrder(date time.Time) []Model.Order {
	return Model.GetEveryTodayOrder(date)
}

func GetReportCount(date time.Time) map[string]interface{} {
	userCount, signCount, count := Model.ReportCount(date)
	data := make(map[string]interface{})
	data["user_count"] = userCount
	data["sign_count"] = signCount
	data["count"] = count
	return data
}
