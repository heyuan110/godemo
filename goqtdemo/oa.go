package main

import (
	"github.com/heyuan110/gorepertory/logger"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"os"
	"time"
)

type QmlBridge struct {
	core.QObject
	_ func() `constructor:"init"`
	_ func(data string) `signal:"sendToQml"`
	_ func(data string) `slot:"sendToGo"`
}

var login_bridge *QmlBridge

func (bridge *QmlBridge)init() {
	//get message from qml
	bridge.ConnectSendToGo(func(data string) {
		logger.Info("r->",data)
		mapQmlCallToFunction(data)
	})
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	gui.NewQGuiApplication(len(os.Args), os.Args)
	login_bridge = NewQmlBridge(nil)
	var app = qml.NewQQmlApplicationEngine(nil)
	app.RootContext().SetContextProperty("bridge", login_bridge)
	app.Load(core.NewQUrl3("qrc:/qml/login.qml", 0))
	gui.QGuiApplication_Exec()
}

func mapQmlCallToFunction(json string)  {
	logger.Info(json)
	action := "login"
	if action == "login" {
		go login("","")
	}
}

func login(u string,p string) {
	time.Sleep(5*time.Second)
	//send message to qml
	login_bridge.SendToQml("login_error")
}