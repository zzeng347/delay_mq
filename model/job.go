package model

type JobResp struct {
	PushJobReq
	RetryCount  int64 `json:"retry_count"`
}

type PushJobReq struct {
	Container string `json:"container" binding:"required,max=30"`
	Id    string `json:"id" binding:"required"`    // job唯一标识ID
	Delay int64  `json:"delay" binding:"required,gte=1"` // 延迟时间, 秒
	TTR   int64  `json:"ttr"`
	Body  string `json:"body" binding:"required"`
	Route string `json:"route" binding:"required"`
}

type PopJob struct {

}

type FinishJob struct {

}

type DelJob struct {

}