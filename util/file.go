package util

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	FilePath = getCurrentPath() + "/assets/img/" //文件默认存储路径
	fileType = []string{                         //文件类型头
		"FFD8FFE0",                 //jpg
		"FFD8FFE1",                 //jpg
		"FFD8FFE8",                 //jpg
		"89504E47",                 //png
		"47494638",                 //gif
		"424D",                     //bmp
		"49492A00",                 //tif
		"41433130",                 //dwg
		"38425053",                 //psd
		"7B5C727466",               //rtf
		"3C3F786D6C",               //xml
		"68746D6C3E",               //html
		"44656C69766572792D646174", //eml
		"CFAD12FEC5FD746F",         //dbx
		"2142444E",                 //pst
		"D0CF11E0",                 //xls/doc
		"5374616E64617264204A",     //mdb
		"FF575043",                 //wpd
		"252150532D41646F6265",     //eps/ps
		"255044462D312E",           //pdf
		"E3828596",                 //pwl
		"504B0304",                 //zip
		"52617221",                 //rar
		"57415645",                 //wav
		"41564920",                 //avi
		"2E7261FD",                 //ram
		"2E524D46",                 //rm
		"000001BA",                 //mpg
		"000001B3",                 //mpg
		"6D6F6F76",                 //mov
		"3026B2758E66CF11",         //asf
		"4D546864",                 //mid
	}
	imgType = []string{ //图片文件类型
		"FFD8FFE0", //jpg
		"FFD8FFE1", //jpg
		"FFD8FFE8", //jpg
		"89504E47", //png
		"47494638", //gif
		"424D",     //bmp
	}
)

/*
保存头像文件
params:
	file：文件
	header: 文件信息
*/
func SaveFile(imgBase64 string) (filename string, err error) {
	timestamp := time.Now().UnixNano()

	//若文件夹不存在，则创建文件夹
	_, e := os.Stat(FilePath)
	if e != nil {
		if os.IsNotExist(e) {
			os.Mkdir(FilePath, 666)
		}
	}

	filename = strconv.FormatInt(timestamp, 10)

	file := FilePath + filename

	//成图片文件并把文件写入到buffer
	imgData, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		LoggerError(err)
		return
	}

	//判断文件头信息是否是合法文件,获取文件16进制编码的前15个字节
	imgBytes := []byte(hex.EncodeToString(imgData))
	if len(imgBytes) < 15 {
		err = errors.New("Illegal Tile")
		return
	}
	imgHeadInfo := string(imgBytes[:15])
	if !IsImage(imgHeadInfo) {
		err = errors.New("Illegal Tile Type")
		return
	}

	//buffer输出到jpg文件中（不做处理，直接写到文件）
	err = ioutil.WriteFile(file, imgData, 0666)
	if err != nil {
		LoggerError(err)
		return
	}

	return
}

//判断文件头是否合法
func IsImage(imgHeadStr string) (b bool) {
	imgHeadStr = strings.ToUpper(imgHeadStr)
	for _, v := range imgType {
		if strings.Contains(imgHeadStr, v) {
			b = true
			break
		}
	}
	return
}

// 获取当前运行目录
func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		LoggerError(err)
		return ""
	}
	fmt.Println(dir)
	return dir
}
