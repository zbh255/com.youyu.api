package main

import (
	"bytes"
	"com.youyu.api/lib/utils/version"
	"encoding/json"
	"log"
	"os"
)

func main() {
	appInfo := version.GetInfo()
	result, err := json.Marshal(appInfo)
	if err != nil {
		log.Fatal(err)
	}
	out := bytes.Buffer{}
	err = json.Indent(&out, result, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	_, err = out.WriteTo(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}