package shared

type Spice struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Level      float32 `json:"level"`
	Substitute string  `json:"substitute"`
}
