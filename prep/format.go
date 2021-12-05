package prep

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	c "github.com/drunkz/parser/common"
)

var singleLoadFiles []string // 单个脚本载入文件列表

func FormatStringVar(text string) (string, error) {
	return text, nil
}

func FormatScript(filename string) (*list.List, *c.ErrInfo) {
	singleLoadFiles = singleLoadFiles[0:0]
	return FormatFirst(filename)

}

// 处理文件包含#include、//注释、首尾空格
func FormatFirst(filename string) (*list.List, *c.ErrInfo) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, c.NewErrInfo(c.ErrOpenFile, err.Error())
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	retList := new(list.List) // 处理完毕后的内容
	lineNum := 0              // 读入文件行号
	singleLoadFiles = append(singleLoadFiles, filename)
	retList.PushBack(fmt.Sprintf("<file:%s>", filename)) // 保存文件名，用于二次分析时溯源
	for {
		lineNum++
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		// 去注释、去首尾空格
		index := bytes.Index(line, []byte("//"))
		lineText := ""
		if index != -1 {
			lineText = strings.TrimSpace(string(line[:index]))
		} else {
			lineText = strings.TrimSpace(string(line[:]))
		}
		// 处理文件包含
		if strings.HasPrefix(lineText, INCLUDE) {
			includeFile := GetIncludeFileName(lineText)
			if c.IsContain(singleLoadFiles, includeFile) {
				// 重复包含
				return nil, c.NewErrInfo(c.ErrRepeatInclude, fmt.Sprintf("位于 %s 第 %d 行", filename, lineNum))
			}
			includeList, err := FormatFirst(includeFile)
			if err != nil {
				// 包含文件错误
				return nil, c.AppendMore(err, fmt.Sprintf("文件引用错误，位于 %s 第 %d 行", filename, lineNum))
			} else {
				retList.PushFrontList(includeList)
			}
			retList.PushBack("") // 替换#include为空，用以占位行
		} else {
			retList.PushBack(lineText)
		}
	}
	return retList, nil
}

// 处理#define、//注释
func FormatSecond() {

}

func GetIncludeFileName(cmd string) string {
	index := strings.Index(cmd, `"`)
	lastIndex := strings.LastIndex(cmd, `"`)
	if index != -1 && lastIndex > index {
		return filepath.Join(SCRIPT_ROOT, cmd[index+1:lastIndex]+SCRIPT_EXT)
	} else {
		return ""
	}
}
