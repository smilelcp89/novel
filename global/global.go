package global

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	RootPath string
	Logger   *logrus.Logger
)

const (
	LoginSession = "loginUserInfo"
)

type MapData map[string]interface{}

func SetRootPath() {
	RootPath, _ = os.Getwd()
}

func SetLogger() {
	Logger = logrus.New()
	file, err := os.OpenFile(RootPath+"/runtime/"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.SetOutput(file)
	} else {
		panic(fmt.Sprintf("Log init error: %v", err))
	}
}

func Setup() {
	SetRootPath()
	SetLogger()
}
