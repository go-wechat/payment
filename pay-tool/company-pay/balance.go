package company_pay

//企业付款为企业提供付款至用户零钱的能力，支持通过API接口付款

import (
	"github.com/go-wechat/util/httper"
)

const (
	TransferToBalanceUrl string = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"
	GetTransferToBalanceInfoUrl string = "https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo"
)

// 企业付款到零钱接口
type TransferToBalanceReq struct {
	MchAppid string `json:"mch_appid", xml:"mch_appid"`
	MchID string `json:"mchid", xml:"mchid"`
}
type TransferToBalanceRes struct {

}
func TransferTOBalance(client httper.HttpClient, request *TransferToBalanceReq) (*TransferToBalanceRes, error){
	var response = &TransferToBalanceRes{}
	return response, nil
}

// 查询企业付款到零钱接口
type GetTransferToBalanceInfoReq struct {

}
type GetTransferToBalanceInfoRes struct {

}
func GetTransferToBalanceInfo(client httper.HttpClient, request *GetTransferToBalanceInfoReq) (*GetTransferToBalanceInfoRes, error){
	var response = &GetTransferToBalanceInfoRes{}
	return response, nil
}