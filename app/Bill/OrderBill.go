package Bill

import (
	"errors"

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
