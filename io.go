package toolkit

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// 获取目录下面的所有文件，subdir是否要包含子目录, filter过滤的扩展名
func AllFiles(dir string, subdir bool, filter []string) []string {
	curdir, _ := os.Open(dir)
	all, _ := curdir.Readdir(0)
	files := make([]string, 0, 32)

	for _, f := range all {
		p := dir + "/" + f.Name()
		if !f.IsDir() {
			if filter != nil && len(filter) > 0 {
				ext := strings.ToLower(path.Ext(p))
				for _, f := range filter {
					if ext == f {
						files = append(files, p)
						break
					}
				}
				continue
			}
			files = append(files, p)
			continue
		}
		if subdir {
			sub := AllFiles(p, subdir, filter)
			files = append(files, sub...)
		}
	}

	return files
}

// 读文件
func ReadFile(f string) ([]byte, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	return ioutil.ReadAll(file)
}