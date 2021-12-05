package prep

import (
	"container/list"
	"sync"
)

type VarType uint8

const (
	Bool = iota
	Int
	String
)

const INCLUDE string = "#include"
const DEFINE string = "#define"
const SCRIPT_EXT string = ".txt" // 脚本扩展名

var SCRIPT_ROOT string = "" // 脚本根目录

type Script struct {
	Filename  string     // 脚本所属文件
	TextList  list.List  // 脚本全部内容
	Functions []Function // 函数集合
	Script    []Script   // 所包含的其他脚本
}

type Variable struct {
	Name  string      // 变量名
	Type  VarType     // 变量类型
	Value interface{} // 变量值
	// #define 常量处理
	// block块 void
	// 先解析获得块
	// 中间代码文件保存，文件头源文件md5值，避免重复解析
	// 展开指令转为类汇编方式，ip寄存器处理判断循环？
	// 寻址 map["文件名"]？
	// 怎么寻址指定方法，跳转到指定位置
}

type CodeBlock struct {
	LocalVars []Variable // 局部变量，代码块执行完毕清空
}

// 全局变量
type GlobalVariable struct {
	Variable
	sync.RWMutex
}

type Function struct {
	Param    []Variable  // 参数列表
	RetValue []Variable  // 返回值
	Blocks   []CodeBlock // 代码块
}
