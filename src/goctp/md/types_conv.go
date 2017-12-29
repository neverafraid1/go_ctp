package md

// #include "mdbridge.h"
import "C"
import (
	"gopkg.in/iconv.v1"

	. "crane/src/goctp/types"
)

func toGoRspInfoField(pCRspInfoField *C.CThostFtdcRspInfoField) (pRspInfoField *RspInfoField) {
	if pCRspInfoField == nil {
		return
	}

	cd, err := iconv.Open("utf-8", "gbk") // convert gbk to utf-8
	if err != nil {
		pRspInfoField = &RspInfoField{
			IntMax, // set to int_max
			err.Error()}
		return
	}

	defer cd.Close()

	pRspInfoField = &RspInfoField{
		int(pCRspInfoField.ErrorID),
		cd.ConvString(C.GoString(&pCRspInfoField.ErrorMsg[0]))}
	return
}

func toGoRspUserLoginField(pCLoginFiled *C.CThostFtdcRspUserLoginField) (pLoginField *RspUserLoginInfoField) {
	if pCLoginFiled == nil {
		return
	}
	pLoginField = &RspUserLoginInfoField{
		C.GoString(&pCLoginFiled.TradingDay[0]),
		C.GoString(&pCLoginFiled.LoginTime[0]),
		C.GoString(&pCLoginFiled.BrokerID[0]),
		C.GoString(&pCLoginFiled.UserID[0]),
		C.GoString(&pCLoginFiled.SystemName[0]),
		int(pCLoginFiled.FrontID),
		int(pCLoginFiled.SessionID),
		C.GoString(&pCLoginFiled.MaxOrderRef[0]),
		C.GoString(&pCLoginFiled.SHFETime[0]),
		C.GoString(&pCLoginFiled.DCETime[0]),
		C.GoString(&pCLoginFiled.CZCETime[0]),
		C.GoString(&pCLoginFiled.FFEXTime[0]),
		C.GoString(&pCLoginFiled.INETime[0])}
	return
}

func toGoUserLogoutField(pUserLogout *C.CThostFtdcUserLogoutField) (pGoUserLogout *ReqUserLogoutInfoField) {
	if pUserLogout == nil {
		return
	}
	pGoUserLogout = &ReqUserLogoutInfoField{
		C.GoString(&pUserLogout.BrokerID[0]),
		C.GoString(&pUserLogout.UserID[0])}
	return
}

func toGoSpecificInstrumentField(pSpecificInstrument *C.CThostFtdcSpecificInstrumentField) (pGoSpecificInstrument *SpecificInstrumentField) {
	if pSpecificInstrument == nil {
		return
	}
	pGoSpecificInstrument = &SpecificInstrumentField{
		C.GoString(&pSpecificInstrument.InstrumentID[0])}
	return
}

func toGoDepthMarketDataField(pDepthMarketData *C.CThostFtdcDepthMarketDataField) (pGoDepthMarketData *DepthMarketDataField) {
	if pDepthMarketData == nil {
		return
	}
	pGoDepthMarketData = &DepthMarketDataField{
		 C.GoString(&pDepthMarketData.TradingDay[0]),
		 C.GoString(&pDepthMarketData.InstrumentID[0]),
		 C.GoString(&pDepthMarketData.ExchangeID[0]),
		 C.GoString(&pDepthMarketData.ExchangeInstID[0]),
		 float64(pDepthMarketData.LastPrice),
		 float64(pDepthMarketData.PreSettlementPrice),
		 float64(pDepthMarketData.PreClosePrice),
		 float64(pDepthMarketData.PreOpenInterest),
		 float64(pDepthMarketData.OpenPrice),
		 float64(pDepthMarketData.HighestPrice),
		 float64(pDepthMarketData.LowestPrice),
		 int(pDepthMarketData.Volume),
		 float64(pDepthMarketData.Turnover),
		 float64(pDepthMarketData.OpenInterest),
		 float64(pDepthMarketData.ClosePrice),
		 float64(pDepthMarketData.SettlementPrice),
		 float64(pDepthMarketData.UpperLimitPrice),
		 float64(pDepthMarketData.LowerLimitPrice),
		 float64(pDepthMarketData.PreDelta),
		 float64(pDepthMarketData.CurrDelta),
		 C.GoString(&pDepthMarketData.UpdateTime[0]),
		 int(pDepthMarketData.UpdateMillisec),
		 float64(pDepthMarketData.BidPrice1),
		 float64(pDepthMarketData.BidVolume1),
		 float64(pDepthMarketData.AskPrice1),
		 float64(pDepthMarketData.AskVolume1),
		 float64(pDepthMarketData.BidPrice2),
		 float64(pDepthMarketData.BidVolume2),
		 float64(pDepthMarketData.AskPrice2),
		 float64(pDepthMarketData.AskVolume2),
		 float64(pDepthMarketData.BidPrice3),
		 float64(pDepthMarketData.BidVolume3),
		 float64(pDepthMarketData.AskPrice3),
		 float64(pDepthMarketData.AskVolume3),
		 float64(pDepthMarketData.BidPrice4),
		 float64(pDepthMarketData.BidVolume4),
		 float64(pDepthMarketData.AskPrice4),
		 float64(pDepthMarketData.AskVolume4),
		 float64(pDepthMarketData.BidPrice5),
		 float64(pDepthMarketData.BidVolume5),
		 float64(pDepthMarketData.AskPrice5),
		 float64(pDepthMarketData.AskVolume5),
		 float64(pDepthMarketData.AveragePrice)}

	return
}

func toGoOrderField(pOrder *C.CThostFtdcOrderField) (pGoOrder *OrderField) {
	if pOrder == nil {
		return
	}

	pGoOrder = &OrderField{
		C.GoString(&pOrder.BrokerID[0]),
		C.GoString(&pOrder.InvestorID[0]),
		C.GoString(&pOrder.InstrumentID[0]),
		C.GoString(&pOrder.OrderRef[0]),
		C.GoString(&pOrder.UserID[0]),
		byte(pOrder.OrderPriceType),
		byte(pOrder.Direction),
		C.GoString(&pOrder.CombOffsetFlag[0]),
		C.GoString(&pOrder.CombHedgeFlag[0]),
		float64(pOrder.LimitPrice),
		int(pOrder.VolumeTotalOriginal),
		byte(pOrder.TimeCondition),
		C.GoString(&pOrder.GTDDate[0]),
		byte(pOrder.VolumeCondition),
		int(pOrder.MinVolume),
		byte(pOrder.ContingentCondition),
		float64(pOrder.StopPrice),
		byte(pOrder.ForceCloseReason),
		(pOrder.IsAutoSuspend != 0),
		C.GoString(&pOrder.BusinessUnit[0]),
		int(pOrder.RequestID),
		C.GoString(&pOrder.OrderLocalID[0]),
		C.GoString(&pOrder.ExchangeID[0]),
		C.GoString(&pOrder.ParticipantID[0]),
		C.GoString(&pOrder.ClientID[0]),
		C.GoString(&pOrder.ExchangeInstID[0]),
		C.GoString(&pOrder.TraderID[0]),
		int(pOrder.InstallID),
		byte(pOrder.OrderSubmitStatus),
		int(pOrder.NotifySequence),
		C.GoString(&pOrder.TradingDay[0]),
		int(pOrder.SettlementID),
		C.GoString(&pOrder.OrderSysID[0]),
		byte(pOrder.OrderSource),
		byte(pOrder.OrderStatus),
		byte(pOrder.OrderType),
		int(pOrder.VolumeTraded),
		int(pOrder.VolumeTotal),
		C.GoString(&pOrder.InsertDate[0]),
		C.GoString(&pOrder.InsertTime[0]),
		C.GoString(&pOrder.ActiveTime[0]),
		C.GoString(&pOrder.SuspendTime[0]),
		C.GoString(&pOrder.UpdateTime[0]),
		C.GoString(&pOrder.CancelTime[0]),
		C.GoString(&pOrder.ActiveTraderID[0]),
		C.GoString(&pOrder.ClearingPartID[0]),
		int(pOrder.SequenceNo),
		int(pOrder.FrontID),
		int(pOrder.SessionID),
		C.GoString(&pOrder.UserProductInfo[0]),
		C.GoString(&pOrder.StatusMsg[0]),
		(pOrder.UserForceClose != 0),
		C.GoString(&pOrder.ActiveUserID[0]),
		int(pOrder.BrokerOrderSeq),
		C.GoString(&pOrder.RelativeOrderSysID[0]),
		int(pOrder.ZCETotalTradedVolume),
		(pOrder.IsSwapOrder != 0)	}
	return
}

func toGoTradeField(pTrade *C.CThostFtdcTradeField) (pGoTrade *TradeField) {
	if pTrade == nil {
		return
	}

	pGoTrade = &TradeField{
		C.GoString(&pTrade.BrokerID[0]),
		C.GoString(&pTrade.InvestorID[0]),
		C.GoString(&pTrade.InstrumentID[0]),
		C.GoString(&pTrade.OrderRef[0]),
		C.GoString(&pTrade.UserID[0]),
		C.GoString(&pTrade.UserID[0]),
		C.GoString(&pTrade.UserID[0]),
		byte(pTrade.Direction),
		C.GoString(&pTrade.OrderSysID[0]),
		C.GoString(&pTrade.ParticipantID[0]),
		C.GoString(&pTrade.ClientID[0]),
		byte(pTrade.TradingRole),
		C.GoString(&pTrade.ExchangeInstID[0]),
		byte(pTrade.OffsetFlag),
		byte(pTrade.HedgeFlag),
		float64(pTrade.Price),
		int(pTrade.Volume),
		C.GoString(&pTrade.TradeDate[0]),
		C.GoString(&pTrade.TradeTime[0]),
		byte(pTrade.TradeType),
		byte(pTrade.PriceSource),
		C.GoString(&pTrade.TraderID[0]),
		C.GoString(&pTrade.OrderLocalID[0]),
		C.GoString(&pTrade.ClearingPartID[0]),
		C.GoString(&pTrade.BusinessUnit[0]),
		int(pTrade.SequenceNo),
		C.GoString(&pTrade.TradingDay[0]),
		int(pTrade.SettlementID),
		int(pTrade.BrokerOrderSeq),
		byte(pTrade.TradeSource),
	}

	return
}
