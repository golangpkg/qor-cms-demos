package main

import (
	"testing"
	"log"
	"os"
	"io/ioutil"
	"strings"
	"sort"
)

var (
	i18nPath = os.Getenv("GOPATH") + "/src/github.com/qor/admin/views"
	i18nMap  = make(map[string]string)
)

func readDir(base, dir string) {
	files, err := ioutil.ReadDir(base + "/" + dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		mode := f.Mode()
		if mode.IsDir() {
			//递归循环，base目录，加上相对目录即可。
			readDir(base, dir+"/"+f.Name())
		} else if strings.Contains(f.Name(), ".tmpl") {
			//println(f.Name())
			ReadFile(base + "/" + dir + "/" + f.Name())
		}
	}
}

func Test_Find(t *testing.T) {
	println("i18nPath: ", i18nPath)
	readDir(i18nPath, "")
	var keys []string
	for k := range i18nMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		k2 := strings.Replace(k, ".", "\t", -1)
		println(k2, ": \"", i18nMap[k], "\"")
	}
}

func ReadFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	println("read file :", fileName)
	if err == nil {
		content := string(data)
		print(len(content))
		result := GetStringInBetween(content, "{{", "}}")
		//println(result)
		//fmt.Print(string(data))
		for _, val := range result {
			val = strings.Replace(val, "\"", "", -1)
			//会有多个空格，只split 2 个。
			vals := strings.SplitN(val, " ", 2)
			println(val, "len:", len(vals))
			if len(vals) == 2 {
				//println(vals[0], ":", vals[1])
				i18nMap[vals[0]] = vals[1]
			}
		}
	}
}

// GetStringInBetween Returns empty string if no start string found
//https://stackoverflow.com/questions/26916952/go-retrieve-a-string-from-between-two-characters-or-other-strings
func GetStringInBetween(str string, start string, end string) (result []string) {
	s := 1
	for s > 0 {
		s = strings.Index(str, start)
		//println(s)
		if s == -1 {
			break
		}
		s += len(start)
		e := strings.Index(str, end)
		//println("start:", s, ",end:", e)
		if e <= s { //有可数组越界的可能。
			break
		}
		tmpStr := str[s:e]
		//判断里面是否存中 t 前缀。
		if strings.Contains(tmpStr, "qor_admin") && strings.HasPrefix(tmpStr, "t ") {
			//println(tmpStr)
			//println(tmpStr)
			//替换字符串。
			result = append(result, strings.Replace(tmpStr, "t ", "", -1))
		}
		str = strings.Replace(str, start+tmpStr+end, "", -1)
	}
	return
}

func Test_FindStr(t *testing.T) {
	println("i18nPath: ", i18nPath)
	fileName := i18nPath + "/index/pagination.tmpl"
	println("fileName: ", fileName)
	ReadFile(fileName)
}
