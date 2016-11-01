package main

import (
	"bufio"
	//"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	//"strings"
	//"reflect"
)

func CheckError(err error) {
	if err != nil {
		panic("ERROR:" + err.Error())
	}
}

func ReadCode(FileName string) []string {
	f, err := os.Open(FileName)
	CheckError(err)
	defer f.Close()
	//buf := bufio.NewReader(f)
	buf := bufio.NewScanner(f)
	codeslice := make([]string, 0)
	for buf.Scan() {

		//code, err := fmt.Scan(FileName)
		code := buf.Text()
		//fmt.Printf("code:%s", code)
		codeslice = append(codeslice, code)
		if err == io.EOF {
			break
		}
		CheckError(err)
	}
	return codeslice
}

func UrlJoin(code string) string {
	return "http://fundgz.1234567.com.cn/js/" + code + ".js"
}

func EncodeGbk(code string) string {
	enc := mahonia.NewEncoder("gbk")
	return enc.ConvertString(string(code))
}

type Fundinfo struct {
	Fundcode, Name, Jzrq, Dwjz, Gsz, Gszzl, Gztime string
}

func main() {
	code := ReadCode("fundcode.txt")
	fmt.Println("基金代码    基金名称        净值日期    单位净值    估算值    估算增长率   估值时间")
	for i := 0; i < len(code); i++ {
		codeinfo, err := http.Get(UrlJoin(code[i]))
		CheckError(err)
		data, err := ioutil.ReadAll(codeinfo.Body)
		//enc := mahonia.NewEncoder("gbk")
		CheckError(err)
		//fmt.Println(UrlJoin(code[i]))
		data = data[8 : len(data)-2]
		//datas := enc.ConvertString(string(data))
		//fmt.Printf("%v\n", datas)
		var fundinfo1 Fundinfo
		json.Unmarshal(data, &fundinfo1)
		//fmt.Println(EncodeGbk(fundinfo1.Name))
		// fmt.Println(reflect.TypeOf(data))
		// fmt.Println(reflect.TypeOf(datas))
		fmt.Printf("%s | %s\n", fundinfo1, fundinfo1.Gszzl)

	}
}
