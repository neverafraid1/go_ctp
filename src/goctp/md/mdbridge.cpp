#include "ThostFtdcMdApi.h"
#include "MdSpi.hpp"
#include "mdbridge.h"
#include "_cgo_export.h"
#include <iostream>
#include <cstring>
#include <cstdlib>
#include <cstdio>
using namespace std;

///创建MdSpi
///@return 创建出的UserSpi
MdSpi CreateFtdcMdSpi()
{
    CMdSpi *pCMdSpi = new CMdSpi();
    return (void *)pCMdSpi;
}

///创建MdApi
///@param pszFlowPath 存贮订阅信息文件的目录，默认为当前目录
///@return 创建出的UserApi
MdApi CreateFtdcMdApi(char *pszFlowPath)
{
    CThostFtdcMdApi *pCMdApi = CThostFtdcMdApi::CreateFtdcMdApi(pszFlowPath);
    return (void *)pCMdApi;
}

/*
void RegisterMdSpi(MdApi pMdApi, MdSpi pMdSpi);
void SubscribePublicTopic(MdApi pMdApi);
void SubscribeMarketData(MdApi pMdApi);
void RegisterFront(MdApi pMdApi);
void Init(MdApi pMdApi);
void JoinMdApi(MdApi pMdApi);
void Release(MdApi pMdApi);


void SetRegisterMdSpiUserID(MdSpi pMdSpi,char *userId);
int ReqUserLogin(MdApi pMdApi, char *account,char *passwd,char *brokerID,int requestID);
*/

///注册回调接口
///@param pSpi 派生自回调接口类的实例
void RegisterMdSpi(MdApi pMdApi, MdSpi pMdSpi)
{
    CThostFtdcMdApi *pCMdApi = (CThostFtdcMdApi *)pMdApi;
    CMdSpi *pCMdSpi = (CMdSpi *)pMdSpi;
    pCMdSpi->SetMdUserApi(pCMdApi);
    pCMdApi->RegisterSpi((CThostFtdcMdSpi *)pCMdSpi);
//    pCMdSpi->setGoMdSpi(pGoMdSpi);
}

///注册前置机网络地址
///@param pszFrontAddress：前置机网络地址。
///@remark 网络地址的格式为：“protocol://ipaddress:port”，如：”tcp://127.0.0.1:17001”
void RegisterMdFront(MdApi pMdApi,char* front_addr)
{
    CThostFtdcMdApi *pCMdApi = (CThostFtdcMdApi *)pMdApi;
    pCMdApi->RegisterFront(front_addr); //connect 前置机
}

///初始化
///@remark 初始化运行环境,只有调用后,接口才开始工作
void Init(MdApi pMdApi)
{
    CThostFtdcMdApi *pCMdApi = (CThostFtdcMdApi *)pMdApi;
    pCMdApi->Init();
}

///等待接口线程结束运行
///@return 线程退出代码
void JoinMdApi(MdApi pMdApi)
{
    CThostFtdcMdApi *pCMdApi = (CThostFtdcMdApi *)pMdApi;
    pCMdApi->Join();
}

///删除接口对象本身
///@remark 不再使用本接口对象时,调用该函数删除接口对象
void Release(MdApi pMdApi)
{
    CThostFtdcMdApi *pCMdApi = (CThostFtdcMdApi *)pMdApi;
    pCMdApi->Release();
}


int ReqUserLogin(MdApi pMdApi, char *account,char *passwd,char *brokerID,int requestID)
{
    CThostFtdcMdApi *pUserApi = (CThostFtdcMdApi *)pMdApi;
    CThostFtdcReqUserLoginField req;
    memset(&req, 0, sizeof(req));
    strcpy(req.BrokerID, brokerID);
    strcpy(req.UserID, account);
    strcpy(req.Password, passwd);
    int iResult = pUserApi->ReqUserLogin(&req, requestID);
    cerr << "--->>> 发送用户登录请求: " << ((iResult == 0) ? "成功" : "失败")
            << "\tBrokerID: " << req.BrokerID
            << "\tInvestorID: " << req.UserID
            << "\tPasswd: " << req.Password << endl;
    return iResult;
};


///订阅行情。
///@parma pMdApi  pMdUserApi
///@param ppInstrumentID 合约ID
///@param nCount 要订阅/退订行情的合约个数
///@remark
int SubscribeMarketData(MdApi pMdApi,char** ppInstrumentID, int iInstrumentID)
{
    cerr << "--->>> 发送行情订阅请求: " << ppInstrumentID<< "size: " << iInstrumentID << endl;
    CThostFtdcMdApi *pUserApi = (CThostFtdcMdApi *)pMdApi;
    int iResult = pUserApi->SubscribeMarketData(ppInstrumentID, iInstrumentID);
    cerr << "--->>> 发送行情订阅请求: " << ((iResult == 0) ? "成功" : "失败") << endl;
}

///退订行情。
///@parma pMdApi  pMdUserApi
///@param ppInstrumentID 合约ID
///@param nCount 要订阅/退订行情的合约个数
///@remark
int UnSubscribeMarketData(MdApi pMdApi ,char *ppInstrumentID[], int nCount)
{
	cerr << "--->>> " << __FUNCTION__ << endl;
	CThostFtdcMdApi *pCMdApi = (CThostFtdcMdApi *)pMdApi;
    int iResult =pCMdApi->UnSubscribeMarketData(ppInstrumentID,nCount);
    cerr << "--->>> 发送退订行情请求: " << ((iResult == 0) ? "成功" : "失败") << endl;

}

void SetMdUserApi(MdApi pMdApi, MdSpi pMdSpi)
{
    CThostFtdcMdApi *pCMdApi = (CThostFtdcMdApi *)pMdApi;
    CMdSpi *pCMdSpi = (CMdSpi *)pMdSpi;
    pCMdSpi->SetMdUserApi(pCMdApi);
}

///获取当前交易日
///@retrun 获取到的交易日
///@remark 只有登录成功后,才能得到正确的交易日
const char* GetTradingDay(MdApi pMdApi)
{
    CThostFtdcMdApi *pCMdApi = (CThostFtdcMdApi *)pMdApi;
    return pCMdApi->GetTradingDay();
}