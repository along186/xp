package Bill

import(
	"github.com/Unknwon/com"
	"strings"
	"time"
	"xp/app/Model"
)

func CheckSystemAvailable() bool {
	data := Model.GetSettingInfoByType(1)
	for _, s := range data {
		// 检查系统是否开放
		if s.Key == "is_open" && s.Value != "1" {
			return false
		}

		// 检查系统开放时间
		if s.Key == "open_time" {
			str := strings.Split(s.Value,"-")
			startSlice := strings.Split(str[0],":")
			startInt := com.StrTo(startSlice[0]).MustInt()*3600 + com.StrTo(startSlice[1]).MustInt()*60 + com.StrTo(startSlice[2]).MustInt()
			endSlice := strings.Split(str[1],":")
			endInt := com.StrTo(endSlice[0]).MustInt()*3600 + com.StrTo(endSlice[1]).MustInt()*60 + com.StrTo(endSlice[2]).MustInt()
			curHour := time.Now().Hour()*3600 + time.Now().Minute()*60 + time.Now().Second()

			if curHour < startInt || curHour > endInt {
				return false
			}
		}
	}

	return true
}
