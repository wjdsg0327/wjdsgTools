package FileUtils

import (
	"fmt"
	"io"
	"os"
)

type FileUtils struct {
}

// Copy 文件拷贝
/*
 *target 目标文件
 *source 被拷贝的文件
 */
func (FileUtils) Copy(target *os.File, source *os.File) (bool, string) {

	written, err := io.Copy(target, source)
	if err != nil {
		sprintf := fmt.Sprintf("拷贝错误：%s", err)
		panic(fmt.Sprintf("拷贝错误：%s", sprintf))
		return false, sprintf
	}
	sprintf := fmt.Sprintf("拷贝成功，文件大小：%d", written)
	return true, sprintf
}

// Read 读取文件内容
/*
 * @params file
 * @params size:需要读取的大小，如果要全部读取，填写 0
 * return : 读取的文件字节，读取大小
 */
func (FileUtils) Read(file *os.File, size int) ([]byte, int) {

	var bs = make([]byte, size)
	if size == 0 {
		stat, err := file.Stat()
		if err != nil {
			panic(fmt.Sprintf("文件信息读取失败:%s", err))
		}
		bs = make([]byte, stat.Size()) // 初始化缓冲区
	}
	n, err := file.Read(bs)
	if err != nil {
		panic(fmt.Sprintf("读取错误:%s", err))

	}
	return bs, n
}
