package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
)

func main () {
	// ⌚👓️➕📹🤓️🎸🧙🐶
	script, err := ioutil.ReadFile(parseConf())
	strScript := strings.ReplaceAll(string(script), "\n", "")
	strScript = strings.ReplaceAll(strScript, " ", "")

	if err != nil {
		panic("[Error] 长者对你的脚本路径很不满意, 你的生命减少了一秒钟……")
	}

	StartInterpreter(string(script))
}

func parseConf () string {
	pHelpFlag := flag.Bool("h", false, "长者的指点")
	pFilePath := flag.String("f", "", "放膜蛤脚本的路径是坠吼的！")
	flag.Parse()

	if *pHelpFlag {
		flag.Usage()
		os.Exit(0)
	}

	return *pFilePath
}
