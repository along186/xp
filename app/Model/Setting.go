package Model

type Setting struct {
	Key string `json:"key"`
	Value string `json:"value"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetSettingInfoByType(t int) (settings []Setting) {
	Map := make(map[string]interface{})
	Map["type"] = t
	Map["status"] = 1
	db.Where(Map).Find(&settings)
	return
}