package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func printFile(file os.FileInfo, root string) {
	// 文件全路径
	filename := path.Join(root, file.Name())
	fmt.Println(filename)
}

func printDir(dir []os.FileInfo, root string, deep int) {
	// 如果深度超过3，则返回
	// 此处仅为测试时防止控制台打印输出太多
	if deep > 3 {
		return
	}

	for _, v := range dir {

		name := v.Name()

		// 忽略.和..
		if strings.HasPrefix(name, ".") {
			continue
		}

		isDir := v.IsDir()
		if !isDir {
			printFile(v, root)
		} else {
			_path := path.Join(root, name)
			// 打印目录名称
			fmt.Println(_path)
			_dirpath, err := os.Open(_path)
			if err != nil {
				log.Fatal(err)
			}
			defer _dirpath.Close()

			_dir, err := _dirpath.Readdir(0)
			if err != nil {
				log.Fatal(err)
			}
			// 递归目录
			printDir(_dir, _path, deep+1)
		}
	}
}

func main() {
	// 读取go目录, just for example!
	rootPath := "/home/oslet"
	rootDir, err := os.Open(rootPath)
	if err != nil {
		log.Fatal(err)
	}
	defer rootDir.Close()

	fs, err := rootDir.Readdir(0)
	printDir(fs, rootPath, 0)
}
