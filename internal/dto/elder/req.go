package elder

type UpdateMyElderReq struct {
	BlockChainAddress string `json:"blockChainAddress" binding:"required"` // binding:"required" 增加了校验，确保该字段不为空
}
