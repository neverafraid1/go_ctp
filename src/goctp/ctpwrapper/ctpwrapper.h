#pragma once

#ifdef __cplusplus
extern "C" {
#endif

    #include <stdbool.h>
    #include "./include/ThostFtdcUserApiStruct.h"
    #include "../_obj/_cgo_export.h"

    void* GetApiInstance();

    void ReleaseApiInstance(void* p);

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
                    const char* szAuthCode);

    void ReqLogin(void* pApi);

    void ExOnConnected(void* p);

    void ExOnRspUserLogin(void* p, void* pRspUserLogin, void* pRspInfo, int nRequestID, bool bIsLast);

#ifdef __cplusplus
}
#endif

