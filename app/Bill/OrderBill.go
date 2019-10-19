package Bill

import "xp/app/Model"

func CheckTodaykHasOrdered(uid int) bool {
	orderList := Model.GetTodayOrderListByUid(uid)
	if len(orderList) > 0 {
		return true
	}

	return false
}

func SaveOrder(order Model.Order) Model.Order {



	return Model.SaveOrder(order)
}