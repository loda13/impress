package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
 * @Author: tang
 * @mail: yuetang2
 * @Date: 2022/6/16 20:55
 * @Desc: 详情见06课程作业.md，启动服务后，可实时修改目录及文件来刷新验证
 */

func _fileCount(countType string, path string, i *int, j *int) int {
	f, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("open dir error: %s", err)
		return 0
	}
	for _, file := range f {
		if file.IsDir() {
			filename := path + "/" + file.Name()
			_fileCount(countType, filename, i, j)
			*j++
		} else {
			*i++
		}
	}
	if countType == "file" {
		return *i
	} else if countType == "path" {
		return *j
	} else {
		return 0
	}
}

func Httpserver(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//fmt.Fprintln(w, r.Form)
	i := 0
	j := 0
	countType := r.FormValue("countType")
	fmt.Fprintln(w, _fileCount(countType, ".", &i, &j))
}

func main() {
	http.HandleFunc("/count", Httpserver)
	http.ListenAndServe("127.0.0.1:80", nil)
}
