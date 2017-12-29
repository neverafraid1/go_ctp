package md

//#include "mdbridge.h"
import "C"
import (
	"log"
)

//export __go__on_connected__
func __go__on_connected__() {
	log.Println("Run go c++ bridge function:__go__on_connected__")
	mdUserSpi.OnFrontConnected()
}

//export __go_on_rsp_user_login__
func __go_on_rsp_user_login__(pRspUserLogin *C.CThostFtdcRspUserLoginField, pRspInfo *C.CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool) {
	log.Println("Run go c++ bridge function:__go_on_rsp_user_login__")

	loginInfo := toGoRspUserLoginField(pRspUserLogin)
	pGoRspInfo := toGoRspInfoField(pRspInfo)

	if pGoRspInfo.ErrorID != 0 {
		log.Println("登陆失败！ ErrorID : ", pGoRspInfo.ErrorID, "ErrorMsg : ", pGoRspInfo.ErrorMsg)
	} else {
		log.Println("登陆成功")
	}
	mdUserSpi.OnRspUserLogin(loginInfo, pGoRspInfo, nRequestID, bIsLast)

}

//export __go_on_rtn_depth_market_data__
func __go_on_rtn_depth_market_data__(pDepthMarketData *C.CThostFtdcDepthMarketDataField) {
	log.Println("Run go c++ bridge function:__go_on_rtn_depth_market_data__")

	pGoDepthMarketData := toGoDepthMarketDataField(pDepthMarketData)
	mdUserSpi.OnRtnDepthMarketData(pGoDepthMarketData)
}

//export __go_on_rsp_sub_market_data__
func __go_on_rsp_sub_market_data__(pSpecificInstrument *C.CThostFtdcSpecificInstrumentField,
	pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)  {
	log.Println("Run go c++ bridge function:__go_on_rsp_sub_market_data__")
	pGoSepcificInstrument := toGoSpecificInstrumentField(pSpecificInstrument)
	pGoRspInfo := toGoRspInfoField(pRspInfo)
	mdUserSpi.OnRspSubMarketData(pGoSepcificInstrument, pGoRspInfo, nRequestID, bIsLast)
}

//export __go_on_rsp_un_sub_market_data__
func __go_on_rsp_un_sub_market_data__(pSpecificInstrument *C.CThostFtdcSpecificInstrumentField,
	pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	log.Println("Run go c++ bridge function:__go_on_rsp_un_sub_market_data__")
	pGoSpecificInstrument := toGoSpecificInstrumentField(pSpecificInstrument)
	pGoRspInfo := toGoRspInfoField(pRspInfo)
	mdUserSpi.OnRspUnSubMarketData(pGoSpecificInstrument, pGoRspInfo, nRequestID, bIsLast)
}
