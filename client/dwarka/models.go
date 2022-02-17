package dwarka

// Building -
type Building struct {
	Name        string  `json:"name"`
	Lat         float64 `json:"lat"`
	Lan         float64 `json:"lan"`
	Description string  `json:"description"`
}

// Floor -
type Floor struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       int    `json:"level"`
}
