package stats

type Summary struct {
	Count       int      `json:"count"`
	Top5Words   []string `json:"top5words"`
	Top5Letters []string `json:"top5letters"`
}
