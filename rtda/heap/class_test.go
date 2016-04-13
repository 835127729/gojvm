package heap

import (
	"fmt"
	_ "gojvm/classfile"
	"gojvm/classpath"
	"gojvm/cmdline"
	"testing"
)

//测试头，用于读取命令行参数
func testHeader() *cmdline.Cmd {
	cmd := cmdline.ParseCmd()

	if cmd.VersionFlag() {
		fmt.Println("version 0.0.1")
		panic("version 0.0.1")
	} else if cmd.HelpFlag() || cmd.Class() == "" {
		cmdline.PrintUsage()
		panic("")
	}
	return cmd
}

//用于打印classmap
func printClassMap(classmap map[string]*Class) {
	for key, val := range classmap {
		fmt.Println("********", key, "********")
		fmt.Println(val.Name())
		for _, m := range val.methods {
			fmt.Println("********ExceptionTable********")
			for _, e := range m.exceptionTable {
				fmt.Println(e)
			}
		}
	}
	fmt.Println("****************")
}

//测试由classfile转化成class
func TestClass(t *testing.T) {

	cmd := cmdline.ParseCmd()

	if cmd.VersionFlag() {
		fmt.Println("version 0.0.1")
		return
	} else if cmd.HelpFlag() || cmd.Class() == "" {
		cmdline.PrintUsage()
		return
	}
	cp := classpath.Parse(cmd.XjreOption(), cmd.CpOption())
	loader := &ClassLoader{
		cp:          cp,
		verboseFlag: true,
		classMap:    make(map[string]*Class),
	}
	loader.LoadClass(cmd.Class())
	printClassMap(loader.classMap)
}

/*
//测试LoadBasicClasses()函数
func TestLoadBasicClasses(t *testing.T) {
	cmd := testHeader()
	cp := classpath.Parse(cmd.XjreOption(), cmd.CpOption())
	loader := &ClassLoader{
		cp:          cp,
		verboseFlag: true,
		classMap:    make(map[string]*Class),
	}
	loader.loadBasicClasses()
	printClassMap(loader.classMap)
}

func TestLoadPrimitiveClasses(t *testing.T) {
	cmd := testHeader()
	cp := classpath.Parse(cmd.XjreOption(), cmd.CpOption())
	loader := &ClassLoader{
		cp:          cp,
		verboseFlag: true,
		classMap:    make(map[string]*Class),
	}
	loader.loadPrimitiveClasses()
	printClassMap(loader.classMap)
}

func TestLoadPrimitiveArrayClasses(t *testing.T) {
	cmd := testHeader()
	cp := classpath.Parse(cmd.XjreOption(), cmd.CpOption())
	loader := &ClassLoader{
		cp:          cp,
		verboseFlag: true,
		classMap:    make(map[string]*Class),
	}
	loader.loadPrimitiveArrayClasses()
	printClassMap(loader.classMap)
}
*/
