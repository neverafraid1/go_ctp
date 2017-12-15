// Created by cgo - DO NOT EDIT

//line /root/go/src/crane/src/goctp/goctp.go:1
package main
//line /root/go/src/crane/src/goctp/goctp.go:14

//line /root/go/src/crane/src/goctp/goctp.go:13
import (
	"fmt"
	"unsafe"
)
//line /root/go/src/crane/src/goctp/goctp.go:19

//line /root/go/src/crane/src/goctp/goctp.go:18
type CTPTradeApi struct {
	mPath			string
	mAddress		string
	mNameServer		string
	mBrokerID		string
	mInvestorID		string
	mPassword		string
	mPrivateResumeType	int32
	mPublicResumeType	int32
	mUserProductInfo	string
	mAutoCode		string
}
//line /root/go/src/crane/src/goctp/goctp.go:32

//line /root/go/src/crane/src/goctp/goctp.go:31
var TraderApiMap map[unsafe.Pointer]*CTPTradeApi
//line /root/go/src/crane/src/goctp/goctp.go:34

//line /root/go/src/crane/src/goctp/goctp.go:33
func (mApi *CTPTradeApi) Connect() {
//line /root/go/src/crane/src/goctp/goctp.go:59

//line /root/go/src/crane/src/goctp/goctp.go:58
}
//line /root/go/src/crane/src/goctp/goctp.go:61

//line /root/go/src/crane/src/goctp/goctp.go:60
func (mApi *CTPTradeApi) OnFrontConnect() {
	fmt.Println("123123123")
}
//line /root/go/src/crane/src/goctp/goctp.go:66

//line /root/go/src/crane/src/goctp/goctp.go:65
func GoExOnFrontConnected(api unsafe.Pointer) {
	v, exist := TraderApiMap[api]
	if !exist {
		return
	}
	v.OnFrontConnect()
}
//line /root/go/src/crane/src/goctp/goctp.go:74

//line /root/go/src/crane/src/goctp/goctp.go:73
func main() {
	api := CTPTradeApi{
		"", "tcp://180.168.146.187:10001", "",
		"9999", "090678", "123456", 2, 2,
		"", ""}
//line /root/go/src/crane/src/goctp/goctp.go:80

//line /root/go/src/crane/src/goctp/goctp.go:79
	(&api).Connect()
}
