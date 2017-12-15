package main

/*
#cgo LDFLAGS: -L./ctpwrapper -L./ -lctpwrapper
#include <stdio.h>
#include <stdlib.h>
#include "./ctpwrapper/ctpwrapper.h"

extern void GoExOnFrontConnected(void* api);

*/
import "C"
import (
	"fmt"
	"unsafe"
)

type CTPTradeApi struct {
	mPath              string
	mAddress           string
	mNameServer        string
	mBrokerID          string
	mInvestorID        string
	mPassword          string
	mPrivateResumeType int32
	mPublicResumeType  int32
	mUserProductInfo   string
	mAutoCode          string
}

var TraderApiMap map[unsafe.Pointer]*CTPTradeApi

func (mApi *CTPTradeApi) Connect() {

	path := C.CString(mApi.mPath)
	address := C.CString(mApi.mAddress)
	nameserver := C.CString(mApi.mNameServer)
	brokerid := C.CString(mApi.mBrokerID)
	investorid := C.CString(mApi.mInvestorID)
	passwd := C.CString(mApi.mPassword)
	userproductinfo := C.CString(mApi.mUserProductInfo)
	authcode := C.CString(mApi.mAutoCode)

	defer func() {
		C.free(unsafe.Pointer(path))
		C.free(unsafe.Pointer(address))
		C.free(unsafe.Pointer(nameserver))
		C.free(unsafe.Pointer(brokerid))
		C.free(unsafe.Pointer(investorid))
		C.free(unsafe.Pointer(passwd))
		C.free(unsafe.Pointer(userproductinfo))
		C.free(unsafe.Pointer(authcode))
	}()

	api := C.GetApiInstance();

	fmt.Printf("%+v\n", mApi)


	C.ReqConnect(api, path, address, nameserver, brokerid, investorid, passwd,
		C.int(mApi.mPrivateResumeType), C.int(mApi.mPublicResumeType), userproductinfo, authcode)

	TraderApiMap[api] = mApi
}

func (mApi *CTPTradeApi) OnFrontConnect() {
	fmt.Println("123123123")
}

//export GoExOnFrontConnected
func GoExOnFrontConnected(api unsafe.Pointer) {
	v, exist := TraderApiMap[api]
	if !exist {
		return
	}
	v.OnFrontConnect()
}

func main() {
	TraderApiMap = make(map[unsafe.Pointer]*CTPTradeApi)

	api := CTPTradeApi{
		"", "tcp://180.168.212.79:31205", "",
		"8000", "41008389", "33905075", 2, 2,
		"", ""}

	(&api).Connect()
}
