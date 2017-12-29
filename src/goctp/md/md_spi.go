package md

// #include "mdbridge.h"
import "C"
import (
	//"unsafe"
	. "crane/src/goctp/types"
	"log"
)

//GoMdSPI golang 行情 SPI
type GoMdSPI interface {
	/*
		//Spi 回调接口
	*/

	//OnFrontConnected 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用
	OnFrontConnected()
	//OnFrontDisconnected 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
	OnFrontDisconnected(nReason int)
	//OnHeartBeatWarning 心跳超时警告。当长时间未收到报文时，该方法被调用。
	OnHeartBeatWarning(nTimeLapse int)
	//OnRspError 错误应答
	OnRspError(pRspInfo *RspInfoField, nRequestID int, bIsLast bool)
	//OnRspUserLogin 登录请求响应
	OnRspUserLogin(pRspUserLogin *RspUserLoginInfoField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)
	//OnRspUserLogout 登出请求响应
	OnRspUserLogout(pUserLogout *ReqUserLogoutInfoField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)
	//OnRspSubMarketData 订阅行情应答
	OnRspSubMarketData(pSpecificInstrument *SpecificInstrumentField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)
	//OnRspUnSubMarketData 取消订阅行情应答
	OnRspUnSubMarketData(pSpecificInstrument *SpecificInstrumentField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)
	//OnRtnDepthMarketData 深度行情通知
	OnRtnDepthMarketData(pDepthMarketData *DepthMarketDataField)

	InitCMdSPI()

	CMdSPI() C.MdSpi
}

var mdUserSpi GoMdSPI

//GoMdSPISimpleImpl 行情 SPI 实现
type GoMdSPISimpleImpl struct {
	MdSPI    C.MdSpi
	pGoMdAPI uintptr
}

//NewGoMdSimpleSPI new GoMdSimpleSPI
func NewGoMdSimpleSPI() *GoMdSPISimpleImpl {
	return &GoMdSPISimpleImpl{}
}

//InitCMdSPI 初始化 C md SPI
func (spi *GoMdSPISimpleImpl) InitCMdSPI() {
	spi.MdSPI = C.CreateFtdcMdSpi()

}

//CMdSPI C md SPI
func (spi GoMdSPISimpleImpl) CMdSPI() C.MdSpi {
	return spi.MdSPI
}

/*///当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
virtual void OnFrontConnected(){};*/

//OnFrontConnected 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用
func (spi *GoMdSPISimpleImpl) OnFrontConnected() {

}

/*
///当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
///@param nReason 错误原因
///        0x1001 网络读失败
///        0x1002 网络写失败
///        0x2001 接收心跳超时
///        0x2002 发送心跳失败
///        0x2003 收到错误报文

virtual void OnFrontDisconnected(int nReason){};
*/

//OnFrontDisconnected 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
func (spi *GoMdSPISimpleImpl) OnFrontDisconnected(nReason int) {
	log.Println("GoMdSPISimpleImpl::", "OnFrontDisconnected")

}

/*
//错误应答
virtual void OnRspError(pRspInfo *C.CThostFtdcRspUserLoginField, nRequestID C.int, bIsLast bool) {};
*/

//OnRspError 错误应答
func (spi *GoMdSPISimpleImpl) OnRspError(pRspInfo *RspInfoField, nRequestID int, bIsLast bool) {

}

/*

//登录请求响应
virtual void OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin, RspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};
*/

//OnRspUserLogin 登录请求响应
func (spi *GoMdSPISimpleImpl) OnRspUserLogin(pRspUserLogin *RspUserLoginInfoField,
	pRspInfo *RspInfoField, nRequestID int, bIsLast bool) {
	log.Println("GoMdSPISimpleImpl::", "OnRspUserLogin")
}

/*
///登出请求响应
virtual void OnRspUserLogout(CThostFtdcUserLogoutField *pUserLogout, pRspInfo *C.CThostFtdcRspInfoField, nRequestID C.int, bIsLast bool) {};
*/

//OnRspUserLogout 登出请求响应
func (spi *GoMdSPISimpleImpl) OnRspUserLogout(pUserLogout *ReqUserLogoutInfoField, pRspInfo *RspInfoField,
	nRequestID int, bIsLast bool) {

}

/*
///心跳超时警告。当长时间未收到报文时，该方法被调用。
///@param nTimeLapse 距离上次接收报文的时间
virtual void OnHeartBeatWarning(int nTimeLapse){};
*/

//OnHeartBeatWarning 心跳超时警告。当长时间未收到报文时，该方法被调用。
func (spi *GoMdSPISimpleImpl) OnHeartBeatWarning(nTimeLapse int) {

}

/*
///订阅行情应答
virtual void OnRspSubMarketData(CThostFtdcSpecificInstrumentField *pSpecificInstrument, pRspInfo *C.CThostFtdcRspInfoField, nRequestID C.int, bIsLast bool) {};
*/

//OnRspSubMarketData 订阅行情应答
func (spi *GoMdSPISimpleImpl) OnRspSubMarketData(pSpecificInstrument *SpecificInstrumentField,
	pRspInfo *RspInfoField, nRequestID int, bIsLast bool) {

}

/*
///取消订阅行情应答
virtual void OnRspUnSubMarketData(CThostFtdcSpecificInstrumentField *pSpecificInstrument, pRspInfo *C.CThostFtdcRspInfoField, nRequestID C.int, bIsLast bool) {};
*/

//OnRspUnSubMarketData 取消订阅行情应答
func (spi *GoMdSPISimpleImpl) OnRspUnSubMarketData(pSpecificInstrument *SpecificInstrumentField,
	pRspInfo *RspInfoField, nRequestID int, bIsLast bool) {

}

/*
///深度行情通知
virtual void OnRtnDepthMarketData(CThostFtdcDepthMarketDataField *pDepthMarketData) {};
*/

//OnRtnDepthMarketData 深度行情通知
func (spi *GoMdSPISimpleImpl) OnRtnDepthMarketData(pDepthMarketData *DepthMarketDataField) {
	log.Println("%+v", pDepthMarketData)
}
