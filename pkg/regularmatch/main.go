package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// 讀取文件所有行
func readLineAll(path string) []string {
	rs := make([]string, 0)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("【提示】failed to open the file, ", err)
		return rs
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		rs = append(rs, fileScanner.Text())
	}
	return rs
}

func main()  {
	var (
		errMsg string
		msg string
	)
	/*args := os.Args
	if len(args) < 2 {
		fmt.Println("【提示】缺少参数, 参数1：匹配字符串，参数2：正则表达式。")
		os.Exit(0)
	}*/
	// 读取文件
	lines := readLineAll("args.txt")
	for i:=0; i<len(lines); i++ {
		if i%2 != 0 {
			fmt.Println("1", "regex:", lines[i], "str:", lines[i-1], "end")
			fmt.Println("2", "regex:", strings.Trim(lines[i], string(rune(32))), "str:", strings.Trim(lines[i-1], string(rune(32))), "end")
			matched, err := regexp.MatchString(strings.Trim(lines[i], string(rune(32))), strings.Trim(lines[i-1], string(rune(32))))
			if err != nil {
				errMsg += lines[i-1] + ":匹配错误，"+err.Error()+"；"
			}
			if matched {
				msg += lines[i-1] + ":匹配成功；"
			} else {
				errMsg += lines[i-1] + ":匹配失败；"
			}
		}
	}
	if len(msg) > 0 {
		fmt.Println("【成功】", msg)
	}
	if len(errMsg) > 0 {
		fmt.Println("【失败】", errMsg)
	}


}