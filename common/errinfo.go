package common

import "log"

func NewErrInfo(errinfo *ErrInfo, more string) *ErrInfo {
	return &ErrInfo{info: errinfo.info, code: errinfo.code, more: more}
}

func AppendMore(errinfo *ErrInfo, more string) *ErrInfo {
	errinfo.more += "\n" + more
	return errinfo
}

func (err ErrInfo) Print() {
	log.Printf("[error]{\n%s %d\n%s}", err.info, err.code, err.more)
}

var ErrFileNotExist = &ErrInfo{info: "文件不存在", code: -1}
var ErrOpenFile = &ErrInfo{info: "打开文件失败", code: -2}
var ErrRepeatInclude = &ErrInfo{info: "重复包含", code: -3}
