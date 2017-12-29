package md

// #cgo CFLAGS: -std=c99
// #cgo linux LDFLAGS: -fPIC -L${SRCDIR}/../ThostTraderApi_v6.3.6/linux64 -Wl,-rpath=${SRCDIR}/../ThostTraderApi_v6.3.6/linux64 -lthostmduserapi -lstdc++
// #cgo linux CPPFLAGS: -fPIC -I${SRCDIR}/../ThostTraderApi_v6.3.6/include
// #cgo windows LDFLAGS: -fPIC -L${SRCDIR}/../ThostTraderApi_v6.3.6/win64 -Wl,-rpath=${SRCDIR}/../ThostTraderApi_v6.3.6/win64 ${SRCDIR}/../ThostTraderApi_v6.3.6/win64/thostmduserapi.lib -lthostmduserapi -lstdc++
// #cgo windows CPPFLAGS: -fPIC -I${SRCDIR}/../ThostTraderApi_v6.3.6/include -DISLIB -DWIN32 -DLIB_MD_API_EXPORT
// #cgo CPPFLAGS: -fPIC -I${SRCDIR}/../include
// #cgo windows CPPFLAGS: -DISLIB -DWIN32 -DLIB_MD_API_EXPORT
// #include "mdbridge.h"
import "C"
import (
	"unsafe"

	"crane/src/goctp/util"
	"log"
)

//GoMdAPI for API
type GoMdAPI struct {
	MdAPI			C.MdApi
	//Spi            GoMdSpi
	Spi            	uintptr
	SessionID      	int    //typedef int TThostFtdcSessionIDType;       //TFtdcSessionIDType是一个会话编号类型
	FrontID        	int    //typedef int TThostFtdcFrontIDType;         //TFtdcFrontIDType是一个前置编号类型
	UserID         	string //typedef char TThostFtdcUserIDType[16];     //TFtdcUserIDType是一个用户代码类型
	BrokerID       	string //typedef char TThostFtdcBrokerIDType[11];   //TFtdcBrokerIDType是一个经纪公司代码类型
	SystemName     	string //typedef char TThostFtdcSystemNameType[41]; //TFtdcSystemNameType是一个系统名称类型
	FrontAddr      	string //前置地址
	password       	string //登录密码
	RequestID      	int
	FrontConnected 	bool //前置连接是否成功
	//lock           *sync.RWMutex
	//waitGroup      *sync.WaitGroup
}

//NewMdAPI init Md user API which hold some key info
//folwpath is start /data/md/flowpath + userId
//make sure mdkir flowpath begin starting
func NewMdAPI(frontAddr, brokerID, userID, password string) *GoMdAPI {

	//TODO 参数校验
	userAPI := &GoMdAPI{}
	userAPI.FrontAddr = frontAddr
	userAPI.BrokerID = brokerID
	userAPI.UserID = userID
	userAPI.password = password
	//If dir not exist create
	flowPath := util.CreateFlowPath("md", userID)
	cFlowPath := C.CString(flowPath)
	defer C.free(unsafe.Pointer(cFlowPath))
	userAPI.MdAPI = C.CreateFtdcMdApi(cFlowPath)

	return userAPI
}

//RegisterSpi SPI
func (api *GoMdAPI) RegisterSpi(spi GoMdSPI) {
	log.Println("%+v", *api)
	C.RegisterMdSpi(api.MdAPI, spi.CMdSPI())
	mdUserSpi = spi
}

//RegisterFront 注册前置机网络地址
//@param pszFrontAddress：前置机网络地址。
//@remark 网络地址的格式为：“protocol://ipaddress:port”，如：”tcp://127.0.0.1:17001”
func (api *GoMdAPI) RegisterFront() {
	log.Println("连接前置服务器：", api.FrontAddr)
	C.RegisterMdFront(api.MdAPI, C.CString(api.FrontAddr))
}

//Connect connect front
func (api *GoMdAPI) Connect() {
	C.Init(api.MdAPI)
}

//Release release connect
func (api *GoMdAPI) Release() {
	C.Release(api.MdAPI)
}

//Join join API
func (api *GoMdAPI) Join() {
	C.JoinMdApi(api.MdAPI)
}

//UserLogin user account login
func (api *GoMdAPI) UserLogin() (result int, err error) {
	var account *C.char = C.CString(api.UserID)
	var pwd *C.char = C.CString(api.password)
	var brokerID *C.char = C.CString(api.BrokerID)
	return int(C.ReqUserLogin(api.MdAPI, account, pwd, brokerID, C.int(api.generateRequestID()))), nil

}

/*//char *ppInstrumentID[], char* iInstrumentID
///订阅行情。
///@param ppInstrumentID 合约ID
///@param nCount 要订阅/退订行情的合约个数
///@remark */

//SubscribeMarketData 订阅行情
func (api *GoMdAPI) SubscribeMarketData(instrumentIDs []string, instrumentIDSize int) (result int, err error) {
	ppInstrumentID := make([]*C.char, len(instrumentIDs))
	for i := range instrumentIDs {
		ppInstrumentID[i] = C.CString(instrumentIDs[i])
		defer C.free(unsafe.Pointer(ppInstrumentID[i]))
	}

	result = int(C.SubscribeMarketData(api.MdAPI, (**C.char)(unsafe.Pointer(&ppInstrumentID[0])), C.int(instrumentIDSize)))
	return
}

//UnSubscribeMarketData 退订行情。
///@param ppInstrumentID 合约ID
///@param nCount 要订阅/退订行情的合约个数
///@remark
func (api *GoMdAPI) UnSubscribeMarketData(instrumentIDs []string, instrumentIDSize int) (result int, err error) {
	ppInstrumentID := make([]*C.char, len(instrumentIDs))
	for i := range instrumentIDs {
		ppInstrumentID[i] = C.CString(instrumentIDs[i])
		defer C.free(unsafe.Pointer(ppInstrumentID[i]))
	}

	result = int(C.UnSubscribeMarketData(api.MdAPI, (**C.char)(unsafe.Pointer(&ppInstrumentID[0])), C.int(instrumentIDSize)))
	return
}

//IsFrontConnected check front is connected or not
func (api *GoMdAPI) IsFrontConnected() bool {
	return api.FrontConnected
}
func (api *GoMdAPI) getCUserID() (cUserID *C.char) {
	cUserID = C.CString(api.UserID)
	return
}

func (api *GoMdAPI) getCBrokerID() (cBrokerID *C.char) {
	cBrokerID = C.CString(api.BrokerID)
	return
}
func (api *GoMdAPI) getCRequestID() C.int {
	return C.int(api.generateRequestID())
}

func (api *GoMdAPI) generateRequestID() int {
	api.RequestID += 1
	return api.RequestID
}
