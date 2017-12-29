#ifdef __cplusplus
extern "C" {
#endif

#if !defined(GOCTP_TYPES_H)
#define GOCTP_TYPES_H
#endif

#if _MSC_VER > 1000
#pragma once
#endif // _MSC_VER > 1000



//请求查询资金账户报文类型
typedef struct CThostFtdcQryTradingAccountField CThostFtdcQryTradingAccountField;
typedef struct CThostFtdcTradingAccountField CThostFtdcTradingAccountField;

typedef struct CThostFtdcRspInfoField CThostFtdcRspInfoField;
typedef struct CThostFtdcRspUserLoginField CThostFtdcRspUserLoginField;
typedef struct CThostFtdcSettlementInfoConfirmField CThostFtdcSettlementInfoConfirmField;
typedef struct CThostFtdcUserLogoutField CThostFtdcUserLogoutField;
typedef struct CThostFtdcSpecificInstrumentField CThostFtdcSpecificInstrumentField;
typedef struct CThostFtdcDepthMarketDataField CThostFtdcDepthMarketDataField;
typedef struct CThostFtdcQrySettlementInfoField CThostFtdcQrySettlementInfoField;
typedef struct CThostFtdcSettlementInfoField CThostFtdcSettlementInfoField;
typedef struct CThostFtdcOrderField CThostFtdcOrderField;
typedef struct CThostFtdcTradeField CThostFtdcTradeField;


#ifdef __cplusplus
}
#endif // GOCTP_TYPES_H