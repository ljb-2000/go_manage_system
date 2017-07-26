package util

import (
	"os"
)

// @Title 读取小文件
func ReadFile(path string) ([]byte, error) {
	// 1. 打开文件
	fl, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fl.Close()

	// 2. 读取文件
	chunks := make([]byte, 0, 2048)
	buf := make([]byte, 512)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}

	return chunks, nil
}