package dto

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func Success() *Resp {
	return &Resp{
		Code: 0,
		Msg:  "success",
	}
}

type PageParam struct {
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}
