package request

type GroupBuy struct {
	Pid int64 `json:"pid" xml:"pid" form:"pid"`
	Num int64 `json:"num" xml:"num" form:"num"`
}

type JoinGroupBuy struct {
	Pid    int64  `json:"pid" xml:"pid" form:"pid"`
	PinkId string `json:"pinkId" xml:"pinkId" form:"pinkId"`
}
type CombinationInfo struct {
	Cid int64 `json:"cid" xml:"cid" form:"cid"`
}
