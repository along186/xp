package Model

type Setting struct {
	Id string	`json:"id"`
	Key string `json:"key"`
	Value string `json:"value"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetSettingInfoByType(t int) (settings []Setting) {
	Map := make(map[string]interface{})
	Map["type"] = 1
	Map["status"] = 1
	db.Where(Map).Find(&settings)
	return
}

func ChangeOrderOpenStatus(open int) bool  {
	s := Setting{}
	Map := make(map[string]interface{})
	Map["type"] = 1
	Map["status"] = 1
	Map["key"] = "is_open"
	db.Where(Map).Last(&s)
	if db.Model(&s).Update("value", open).Error == nil {
		return true
	} else {
		return false
	}
}