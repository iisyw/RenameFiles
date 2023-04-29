package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"renameFile/controller"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config") // 设置配置文件名为“config”
	viper.SetConfigType("ini")    // 设置配置文件类型为“ini”
	viper.AddConfigPath(".")      // 设置在当前目录中查找配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Panic("没有找到配置文件")
		} else {
			// Config file was found but another error was produced
			log.Panic("初始化配置出错", err.Error())
		}
	}
	controller.RenameType = viper.GetInt("fileInfo.rename_type")
	controller.RandomStringDigit = viper.GetInt("fileInfo.random_string_digit")
}

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
	rand.Seed(time.Now().Unix())
	// 获取当前目录路径
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("获取当前目录失败", err.Error())
		return
	}

	fmt.Print("说明：将当前工作路径中第 n 层的`所有文件及文件夹`重命名后，迁出`所在文件夹`并删除`所在文件夹`\n当前工作路径为第一层，不可迁出当前目录文件及文件夹到上一级目录，因此层数为2~9层\n输入指令：\n\teg：查看示例\n\t2~9：文件在第2~9层\n请输入指令：")
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
		20220901-1oascder.jpg(第二层文件)
		20220901-idjg4ltd.jpg(第二层文件)
		20220901-91jS0scd.png(第二层文件)
		20220902-8sJXNsax.exe(第二层文件)
		20220902-fosJaq2s.gif(第二层文件)
		20220902-asmCmse2.txt(第二层文件)
	文件夹B(第一层文件夹)
		20220903-fandpkfn(第二层文件夹)
		20220904-sadn3fma(第二层文件夹)
		20220903-qasfpoid.jpeg(第二层文件)
		20220903-AS9jsfhg.md(第二层文件)
		20220904-sdaclfgf.go(第二层文件)
		20220904-qwedsavx.c(第二层文件)`)
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
	renameMoveOut(currentDir, layerNum, "")
}

// 重命名并迁出
func renameMoveOut(dir string, layerNum int, dirName string) {
	// 获取dir路径下的所有文件
	matches, err := filepath.Glob(path.Join(dir, "*"))
	if err != nil {
		log.Println("获取文件夹信息失败", err.Error())
		return
	}
	// 如果层数不等于1
	if layerNum != 1 {
		// 循环dir路径下的所有文件及文件夹
		for _, v := range matches {
			// 获取每个文件或文件夹的信息
			fileInfo, err := os.Stat(v)
			if err != nil {
				log.Println("获取v信息失败", err.Error())
				return
			}
			// 如果是文件夹就继续往下执行，如果层数已经到了倒数第二层，就传入文件夹名
			if fileInfo.IsDir() {
				if layerNum == 2 {
					renameMoveOut(path.Join(dir, fileInfo.Name()), layerNum-1, fileInfo.Name())
				} else {
					renameMoveOut(path.Join(dir, fileInfo.Name()), layerNum-1, "")
				}
			}
		}
	} else {
		// 如果层数到了最后一层
		for _, v := range matches {
			filename := filepath.Base(v)
			randomStr := RandStringBytes(controller.RandomStringDigit)
			if controller.RenameType == 1 {
				os.Rename(v, path.Join(dir[:len(dir)-len(dirName)-1], dirName+"-"+filename))
			} else if controller.RenameType == 2 {
				os.Rename(v, path.Join(dir[:len(dir)-len(dirName)-1], dirName+"-"+randomStr+path.Ext(v)))
			}
		}
		// 删除文件夹
		os.Remove(dir)
	}
}
