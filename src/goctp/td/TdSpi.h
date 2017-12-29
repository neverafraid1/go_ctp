#pragma once
#include "ThostFtdcTraderApi.h"

class CTdSpi : public CThostFtdcTraderSpi
{
public:
    ///当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
    virtual void OnFrontConnected();

    ///当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
    ///@param nReason 错误原因
    ///        0x1001 网络读失败
    ///        0x1002 网络写失败
    ///        0x2001 接收心跳超时
    ///        0x2002 发送心跳失败
    ///        0x2003 收到错误报文
    virtual void OnFrontDisconnected(int nReason);

    ///心跳超时警告。当长时间未收到报文时，该方法被调用。
    ///@param nTimeLapse 距离上次接收报文的时间
    virtual void OnHeartBeatWarning(int nTimeLapse);

    ///客户端认证响应
    virtual void OnRspAuthenticate(CThostFtdcRspAuthenticateField *pRspAuthenticateField, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);


    ///登录请求响应
    virtual void OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    ///登出请求响应
    virtual void OnRspUserLogout(CThostFtdcUserLogoutField *pUserLogout, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    ///请求查询报单响应
    virtual void OnRspQryOrder(CThostFtdcOrderField *pOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

    virtual void SetTdUserApi(CThostFtdcTraderApi* pUserApi);

    virtual void SetGoTdSpi(void* pGoTdSpi);

private:
    TThostFtdcFrontIDType	FrontID;
    ///会话编号
    TThostFtdcSessionIDType	SessionID;
    ///经纪公司代码
    TThostFtdcBrokerIDType	BrokerID;
    ///用户代码
    TThostFtdcUserIDType	UserID;
    ///投资者帐号
    TThostFtdcInvestorIDType InvestorID;
    ///密码
    TThostFtdcPasswordType  Password;

    CThostFtdcTraderApi* pUserApi;

    void *GoTdSpi;

}