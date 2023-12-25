package mysql

type TRecord struct {
	*Meta
	Name       string  `json:"name"`
	Weight     float32 `json:"weight"`
	BF         string  `json:"bf"`
	LUN        string  `json:"lun"`
	DIN        string  `json:"din"`
	EXTRA      string  `json:"extra"`
	Core       int     `json:"core"`
	Runner     int     `json:"runner"`
	Support    int     `json:"support"`
	Squat      int     `json:"squat"`
	EasyBurpee int     `json:"easy_burpee"`
	Chair      int     `json:"chair"`
	Stretch    int     `json:"stretch"`
	Vol1       string  `json:"vol1"`
	Vol2       string  `json:"vol2"`
	Vol3       string  `json:"vol3"`
	Vol4       string  `json:"vol4"`
	Content    string  `json:"content"`

	Region string `json:"region"`
}