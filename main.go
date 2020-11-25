package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

var SpineMap map[string]string
var spinePath string = "/Applications/Spine/Spine.app/Contents/MacOS/Spine "
var exportPath string ="/Users/cabrite/Desktop/go/batchExportSpine/export"
var jsonFile string = "./spine_export_setting.json"

func main() {
	SpineMap = make(map[string]string)
	GetAllFile("/Users/cabrite/Desktop/go/batchExportSpine/spine")
	for k,v := range(SpineMap) {
		//fmt.Println("name = "+k+"  path = " + v)
		exportName := exportPath + "/" + k
		fullPath := spinePath+"-i " + v +" -o " + exportName + " -e "+jsonFile
		fmt.Println("fullPath = " + string(fullPath))
		var cmd = exec.Command("sh","export.sh",v,exportName,jsonFile)
		output,err := cmd.Output()
		if err != nil {
			fmt.Println("output = " + string(output))
		}
	}
}
func GetAllFile(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			GetSpinePath(pathname + "/" + fi.Name(),fi.Name())
		}
	}
	return err
}

func GetSpinePath(pathname string,key string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if filepath.Ext(fi.Name()) == ".spine" {
			SpineMap[key] = pathname + "/" + fi.Name()
		}
	}
	return err
}