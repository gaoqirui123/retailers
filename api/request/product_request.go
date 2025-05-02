package request

type GroupBuy struct {
	Pid int64 `json:"pid" xml:"pid" form:"pid"`
	Num int64 `json:"num" xml:"num" form:"num"`
}

type JoinGroupBuy struct {
	Pid    int64  `json:"pid" xml:"pid" form:"pid"`
	PinkId string `json:"pinkId" xml:"pinkId" form:"pinkId"`
}

type AddSeckillProduct struct {
	ProductId   int64   `json:"productId" xml:"productId" form:"productId"`
	Num         int64   `json:"num" xml:"num" form:"num"`
	Price       float64 `json:"price" xml:"price" form:"price"`
	Description string  `json:"description" xml:"description" form:"description"`
	StartTime   string  `json:"startTime" xml:"startTime" form:"startTime"`
	StopTime    string  `json:"stopTime" xml:"stopTime" form:"stopTime"`
}

type ReverseStock struct {
	ProductId int64 `json:"productId" xml:"productId" form:"productId"`
}
