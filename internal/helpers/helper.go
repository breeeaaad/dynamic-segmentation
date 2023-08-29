package helpers

type User struct {
	Id int `uri:"id" binding:"required"`
}

type Segment struct {
	Name    string  `uri:"segment" json:"segment"`
	Percent float32 `json:"percent"`
}

type Add struct {
	Id          int      `json:"id"`
	Addsegments []string `json:"addsegments"`
	Delsegments []string `json:"delsegments"`
}
