package md

import (
	"os"
	"testing"
	"time"
	"log"
)

var mdUserApi *GoMdAPI

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	time.Sleep(2 * time.Second)
	tearDown()
	os.Exit(retCode)
}
func setUp() {

	log.Println("start")
	mdUserApi = NewMdAPI("tcp://180.168.146.187:10010", "9999", "090678", "1256")
	spi := NewGoMdSimpleSPI()
	spi.InitCMdSPI()
	mdUserApi.RegisterSpi(spi)
	mdUserApi.RegisterFront()
	mdUserApi.Connect()


	go func() {
		mdUserApi.Join()
	}()

}
func tearDown() {
	mdUserApi.Release()
}

func TestGoMdApi_UserLogin(t *testing.T) {

	for {
		result, err := mdUserApi.UserLogin()
		if err != nil {
			//todo
		}
		if result == 0 {
			break
		}
		time.Sleep(2 * time.Second)
		result, err = mdUserApi.UserLogin()

	}
}

func TestGoMdAPI_SubscribeMarketData(t *testing.T) {
	ins := [] string{"rb1805"}
	mdUserApi.SubscribeMarketData(ins, 1)

	for true {
		time.Sleep(time.Second)
	}
}
