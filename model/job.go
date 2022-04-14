package model

type JobResp struct {
	Topic string `json:"topic"`
	Id    string `json:"id"`
	Delay int64  `json:"delay"`
	TTR   int64  `json:"ttr"`
	Body  string `json:"body"`
}

type PushJobReq struct {
	Topic string `json:"topic" binding:"required,max=30"`
	Id    string `json:"id" binding:"required"`    // job唯一标识ID
	Delay int64  `json:"delay" binding:"required,gte=1"` // 延迟时间, 秒
	TTR   int64  `json:"ttr"`
	Body  string `json:"body" binding:"required"`
}

type PopJob struct {

}

type FinishJob struct {

}

type DelJob struct {

}