package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
)

//flag 支持三种形式命令
//1. -flag xxx 仅仅支持bool类型
//2. -flag=xxx 支持所有类型
//3. -flag xxx

var (
	intflag    int
	boolflag   bool
	stringflag string
)

func main() {
	// 需要最后调用
	flag.Parse()
	//fmt.Println(flag.Lookup("init").Value)
	// 获取每一个参数未解析的参数（这tm 啥sb方法,一般逻辑应该是返回flag 含有啥参数才对啊）
	fmt.Println(flag.Args())
	flag.NArg()
	for i := 0; i < flag.NFlag(); i++ {
		fmt.Println(flag.Arg(i))
	}
	// 获取flag的值 包括未解析和解析过的
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Println("所有参数", f.Name, f.Value)
	})

	flag.Visit(func(f *flag.Flag) {
		fmt.Println("已经解析的", f.Name, f.Value)
	})
	// 未解析的
	flag.Args()

	// flag 是flagSet 一个简易用法
	// 参数，这个参数可以来源于配置文件之类的
	args := []string{"-intflag", "12", "-stringflag", "test"}
	// 重置flag
	// param 1: 程序名称；param 2: 错误处理
	flagset := flag.NewFlagSet("demo", flag.ExitOnError)
	flagset.IntVar(&intflag, "intflag", 0, "int flag value")
	flagset.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	flagset.StringVar(&stringflag, "stringflag", "default", "string flag value")
	// 手动解析参数
	flagset.Parse(args)

	pflagSet := pflag.NewFlagSet("demo2", pflag.ExitOnError)
	// 从命令行中找到intflag  如果有就加入道pflagset中
	pflagSet.AddGoFlag(flag.CommandLine.Lookup("intflag"))
	//pflagSet
	pflagSet.Parse(args)
}

func init() {
	flag.Int("init", 123, "init flag")

	flag.IntVar(&intflag, "iniflag", 123, "init flag")
	flag.BoolVar(&boolflag, "boolflag", true, "bool flag")
	flag.StringVar(&stringflag, "strflag", "hello", "str flag")
}
