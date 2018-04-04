package contract_pay
// 微信委托代扣
// 委托代扣可应用于定期扣款或需事后扣款以提高效率的场景。例如但不限于：会员制缴费、水电煤缴费、黄钻绿钻增值服务、打车类软件、停车场或高速公路无人缴费、理财通基金定投、信用卡还款等通过用户授权给商户，进行委托扣款的场景

import (
	"github.com/go-wechat/util/httper"
)

const (
	SignContractUrl string = "https://api.mch.weixin.qq.com/papay/entrustweb"
)

//公众号和APP调起签约api
type SignContractReq struct {

}
type SignContractRes struct {

}
func SignContract() {

}
