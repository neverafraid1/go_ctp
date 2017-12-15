#include <stdlib.h>

#include "ctpwrapper.h"
#include "TraderApi.h"


void* GetApiInstance()
{
    TraderApi* p = (TraderApi*)malloc(sizeof(TraderApi));

    return p;
}

void ReleaseApiInstance(void* p)
{
    free((TraderApi*)p);
}

void ReqConnect(void* pApi,
                const char* szPath,
                const char* szAddresses,
                const char* szNameServer,
                const char* szBrokerId,
                const char* szInvestorId,
                const char* szPassword,
                int nPrivateResumeType,
                int nPublicResumeType,
                const char* szUserProductInfo,
                const char* szAuthCode)
{
    ((TraderApi*)pApi)->ReqConnect(szPath, szAddresses, szNameServer, szBrokerId, szInvestorId,
     szPassword, (THOST_TE_RESUME_TYPE)nPrivateResumeType, (THOST_TE_RESUME_TYPE)nPublicResumeType, szUserProductInfo, szAddresses);
     printf("ReqConnect\n");
}

void ReqLogin(void* pApi)
{
    ((TraderApi*)pApi)->ReqLogin();
}

void ExOnConnected(void* p)
{
    printf("ExOnConnected\n");
    GoExOnFrontConnected(p);
}

void ExOnRspUserLogin(void* p, void* pRspUserLogin, void* pRspInfo, int nRequestID, bool bIsLast)
{
//    GoExOnRspUserLogin(p, pRspUserLogin, pRspInfo, nRequestID, bIsLast);
}