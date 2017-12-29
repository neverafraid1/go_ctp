#ifdef __cplusplus
extern "C" {
#endif

#if !defined(GOCTP_MD_API_H)
#define GOCTP_MD_API_H
#endif

//#if _MSC_VER > 1000
//#pragma once
//#endif // _MSC_VER > 1000

#include <stdio.h>
#include <stdlib.h>
#include "ThostFtdcUserApiStruct.h"
#include "types.h"

//核心API SPI
typedef void *MdApi;
typedef void *MdSpi;
typedef void *GoMdApi;


///初始化
///@remark 初始化运行环境,只有调用后,接口才开始工作
MdSpi CreateFtdcMdSpi();

///创建MdApi
///@param pszFlowPath 存贮订阅信息文件的目录，默认为当前目录
///@return 创建出的UserApi
MdApi CreateFtdcMdApi(char *pszFlowPath);
///注册回调接口
///@param pSpi 派生自回调接口类的实例
//void RegisterMdSpi(void* pGoMdSpi,MdApi pMdApi, MdSpi pMdSpi);
void RegisterMdSpi(MdApi pMdApi, MdSpi pMdSpi);
///注册前置机网络地址
///@parma pMdApi  pMdUserApi
///@param pszFrontAddress：前置机网络地址。
///@remark 网络地址的格式为：“protocol://ipaddress:port”，如：”tcp://127.0.0.1:17001”。
///@remark “tcp”代表传输协议，“127.0.0.1”代表服务器地址。”17001”代表服务器端口号。
void RegisterMdFront(MdApi pMdApi,char* front_addr);
///初始化
///@remark 初始化运行环境,只有调用后,接口才开始工作
void Init(MdApi pMdApi);
///等待接口线程结束运行
///@return 线程退出代码
void JoinMdApi(MdApi pMdApi);
///删除接口对象本身
///@parma pMdApi  pMdUserApi
///@remark 不再使用本接口对象时,调用该函数删除接口对象
void Release(MdApi pMdApi);


void SetMdUserApi(MdApi pMdApi, MdSpi pMdSpi);

///MdApi
int ReqUserLogin(MdApi pMdApi, char *account,char *passwd,char *brokerID,int requestID);
///订阅行情。
///@parma pMdApi  pMdUserApi
///@param ppInstrumentID 合约ID
///@param nCount 要订阅/退订行情的合约个数
///@remark
int SubscribeMarketData(MdApi pMdApi,char** ppInstrumentID, int iInstrumentID);

///获取当前交易日
///@retrun 获取到的交易日
///@parma pMdApi  pMdUserApi
///@remark 只有登录成功后,才能得到正确的交易日
const char *GetTradingDay(MdApi pMdApi);

///退订行情。
///@parma pMdApi  pMdUserApi
///@param ppInstrumentID 合约ID
///@param nCount 要订阅/退订行情的合约个数
///@remark
int UnSubscribeMarketData(MdApi pMdApi ,char *ppInstrumentID[], int nCount);

#ifdef __cplusplus
}
#endif