package logging

import (
	"fmt"
	"go-gin/pkg/setting"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = setting.LogPath
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// openLogFile 打开日志文件
// os.Stat 会返回文件信息结构描述文件，如果出现错误会返回*PathError
func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	//os.IsExists 能够接受ErrNotExist、syscall的一些错误 能够得知文件不存在或者目录不存在
	case os.IsNotExist(err):
		mkDir()
	//os.IsPermission 能够接受ErrPermission syscall错误 判断权限是否满足
	case os.IsPermission(err):
		log.Fatalf("Permission failed. %v", err)
	}

	// os.OpenFile 指定模式调用文件、文件权限。返回的方法*File可以用于IO。
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile : %v", err)
	}
	return handle
}

func mkDir() {
	//os.Getwd 返回当前文件根路径
	dir, _ := os.Getwd()
	logDir := dir + "/" + getLogFilePath()
	_, err := os.Stat(logDir)
	if !os.IsNotExist(err) {
		return
	}
	//os.ModePerm   ModePerm FileMode = 0777 // Unix permission bits
	err = os.Mkdir(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}

}
