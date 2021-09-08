package main

import (
	"com.youyu.api/app/uts"
	"flag"
	"fmt"
	"os"
)

func main() {
	_ = flag.String("v", "", "查看应用程序版本信息")
	create := flag.String("c", "", "-c fileType 创建的文件类型，err_msg/ras_key")
	filePath := flag.String("p", "", "-p filePath 创建文件的路径，不指定则默认为./dir")

	// 版本信息-v没有赋值参数
	// 处理单个参数
	if len(os.Args) == 1 {
		fmt.Println("参数不能为空")
		fmt.Println("具体支持的参数请查看帮助信息,--h")
		flag.PrintDefaults()
		return
	} else if len(os.Args) == 2 {
		switch os.Args[1] {
		case "-v", "--v":
			uts.GetAppInfoToJson()
			return
		}
	}
	flag.Parse()
	// 创建对应的文件
	if *create != "" {
		switch *create {
		case "err_msg":
			uts.CreateBusinessErrInfoToJson(*filePath)
			break
		case "rsa_key":
			break
		}
	}
}
