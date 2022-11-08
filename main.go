package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 获取n位随机数
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	// 获取当前目录路径
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("获取当前目录失败", err.Error())
		return
	}

	fmt.Print("说明：将当前工作路径中第 n 层的`所有文件及文件夹`重命名为`所在文件夹名-8 位随机字符串.后缀`，迁出`所在文件夹`并删除`所在文件夹`\n输入指令：\n\teg：查看示例\n\t2~9：文件在第2~9层\n请输入指令：")
	var str string
	fmt.Scan(&str)
	if str == "eg" {
		fmt.Println(`
==================举例==================
当前工作路径：
	此程序.exe
	文件夹A(第一层文件夹)
		20220901(第二层文件夹)
			asdohasisaf.jpg(第三层文件)
			qweqwrasdfas.jpg(第三层文件)
			sandaofdna.png(第三层文件)
		20220902(第二层文件夹)
			adasknfadnf.exe(第三层文件)
			sadmpanfda.gif(第三层文件)
			asdjbocasda.txt(第三层文件)
	文件夹B(第一层文件夹)
		20220903(第二层文件夹)
			nfdosfdaif(第三层文件夹)
			dapsojfpaf.jpeg(第三层文件)
			asdsakfadfa.md(第三层文件)
		20220904(第二层文件夹)
			panfpandva(第三层文件夹)
			onasfdabsad.go(第三层文件)
			csnoafncoda.c(第三层文件)
==========运行程序，输入层数 3 进行确认==========
说明：将当前工作路径中第 3 层的所有文件及文件夹重命名为“所在文件夹名-8位随机字符串.后缀”，迁出所在文件夹并删除所在文件夹
当前工作路径：
	此程序.exe
	文件夹A(第一层文件夹)
		20220901-1oascder.jpg
		20220901-idjg4ltd.jpg
		20220901-91jS0scd.png
		20220902-8sJXNsax.exe
		20220902-fosJaq2s.gif
		20220902-asmCmse2.txt
	文件夹B(第一层文件夹)
		20220903-fandpkfn(文件夹)
		20220904-sadn3fma(文件夹)
		20220903-qasfpoid.jpeg
		20220903-AS9jsfhg.md
		20220904-sdaclfgf.go
		20220904-qwedsavx.c`)
		return
	}
	layerNum, err := strconv.Atoi(str)
	if err != nil && strings.Contains("invalid syntax", err.Error()) {
		fmt.Println("指令输入错误，请输入eg或2~9整数")
		return
	}
	if err != nil {
		log.Println("转整数失败", err.Error())
		return
	}
	if layerNum < 2 || layerNum > 9 {
		fmt.Println("指令输入错误，请输入eg或2~9整数")
		return
	}
	perform(currentDir, layerNum, "")
}

// 执行
func perform(dir string, layerNum int, dirName string) {
	// 获取路径下的所有文件
	matches, err := filepath.Glob(path.Join(dir, "*"))
	if err != nil {
		log.Println("获取文件夹信息失败", err.Error())
		return
	}
	// 如果层数不等于1
	if layerNum != 1 {
		// 循环路径下的所有文件及文件夹
		for _, v := range matches {
			// 获取信息
			fileInfo, err := os.Stat(v)
			if err != nil {
				log.Println("获取v信息失败", err.Error())
				return
			}
			// 如果是文件夹就继续往下执行，如果层数已经到了倒数第二层，就传入文件夹名
			if fileInfo.IsDir() {
				if layerNum == 2 {
					perform(path.Join(dir, fileInfo.Name()), layerNum-1, fileInfo.Name())
				} else {
					perform(path.Join(dir, fileInfo.Name()), layerNum-1, "")
				}
			}
		}
	} else {
		// 如果层数到了最后一层
		for _, v := range matches {
			randomStr := RandStringBytes(8)
			os.Rename(v, path.Join(dir[:len(dir)-len(dirName)-1], dirName+"-"+randomStr+path.Ext(v)))
		}
		// 删除文件夹
		os.Remove(dir)
	}
}
