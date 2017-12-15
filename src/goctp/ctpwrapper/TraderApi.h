#pragma once

#include "./include/ThostFtdcTraderApi.h"
#include <stdio.h>
//#include "ctpwrapper.h"

class TraderApi : public CThostFtdcTraderSpi
{
public:
    TraderApi();
    void ReqConnect(const char* szPath,
                    const char* szAddresses,
                    const char* szNameServer,
                    const char* szBrokerId,
                    const char* szInvestorId,
                    const char* szPassword,
                    THOST_TE_RESUME_TYPE nPrivateResumeType,
                    THOST_TE_RESUME_TYPE nPublicResumeType,
                    const char* szUserProductInfo,
                    const char* szAuthCode);
    void ReqLogin();

private:
    virtual void OnFrontConnected();
    virtual void OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

private:
    CThostFtdcTraderApi*		m_pApi;

    char*						m_szPath;
    char                        m_szAddresses[40];
    char*                       m_szNameServer;
    TThostFtdcBrokerIDType		m_szBrokerId;
    TThostFtdcInvestorIDType    m_szInvestorId;
    TThostFtdcPasswordType      m_szPassword;
    THOST_TE_RESUME_TYPE        m_nPrivateResumeType;
    THOST_TE_RESUME_TYPE        m_nPublicResumeType;
    TThostFtdcProductInfoType   m_szUserProductInfo;
    TThostFtdcAuthCodeType      m_szAuthCode;

    volatile int reqid;
};