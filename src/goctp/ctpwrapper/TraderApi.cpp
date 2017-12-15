#include "TraderApi.h"
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include "ctpwrapper.h"

extern void ExOnConnected(void* api);
extern void ExOnRspUserLogin(void* api, void* pRspUserLogin, void* pRspInfo, int nRequestID, bool bIsLast);

TraderApi::TraderApi()
{
    m_pApi = NULL;
    reqid = 0;
}

void TraderApi::ReqConnect(const char* szPath,
                           const char* szAddresses,
                           const char* szNameServer,
                           const char* szBrokerId,
                           const char* szInvestorId,
                           const char* szPassword,
                           THOST_TE_RESUME_TYPE nPrivateResumeType,
                           THOST_TE_RESUME_TYPE nPublicResumeType,
                           const char* szUserProductInfo,
                           const char* szAuthCode)
{
//    if(m_pApi != NULL)
//        return;

    memset(m_szAddresses, 0, 40);
    memcpy(m_szAddresses, szAddresses, strlen(szAddresses));
	memcpy(m_szBrokerId, szBrokerId, sizeof(TThostFtdcBrokerIDType));
	memcpy(m_szInvestorId, szInvestorId, sizeof(TThostFtdcInvestorIDType));
	memcpy(m_szPassword, szPassword, sizeof(TThostFtdcPasswordType));
	m_nPrivateResumeType = nPrivateResumeType;
	m_nPublicResumeType = nPublicResumeType;
	memcpy(m_szUserProductInfo, szUserProductInfo, sizeof(TThostFtdcProductInfoType));
	memcpy(m_szAuthCode, szAuthCode, sizeof(TThostFtdcAuthCodeType));

	m_pApi = CThostFtdcTraderApi::CreateFtdcTraderApi(szPath);

	if (m_pApi)
	{
		m_pApi->RegisterSpi(this);

        m_pApi->RegisterFront(m_szAddresses);

        printf("%s\n", m_szAddresses);

		m_pApi->SubscribePublicTopic(nPublicResumeType);
		m_pApi->SubscribePrivateTopic(nPrivateResumeType);

		m_pApi->Init();

        sleep(1);

		printf("123123123123123123\n");
		fflush(stdout);
	}
}

void TraderApi::ReqLogin()
{

    CThostFtdcReqUserLoginField body;

    strncpy(body.BrokerID, m_szBrokerId, sizeof(TThostFtdcBrokerIDType));
    strncpy(body.UserID, m_szInvestorId, sizeof(TThostFtdcInvestorIDType));
    strncpy(body.Password, m_szPassword, sizeof(TThostFtdcPasswordType));
    strncpy(body.UserProductInfo,m_szUserProductInfo, sizeof(TThostFtdcProductInfoType));

    m_pApi->ReqUserLogin(&body, ++reqid);
}

void TraderApi::OnFrontConnected()
{

    printf("22222222222\n");
    fflush(stdout);
    ExOnConnected((void*)this);
}

void TraderApi::OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
//    ExOnRspUserLogin((void*)this, (void*)pRspUserLogin, (void*)pRspInfo, nRequestID, bIsLast);
}