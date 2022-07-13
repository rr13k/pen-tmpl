package suger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// 判断文件路径是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Police(title string, msg string, erps []string) {
	data := map[string]interface{}{
		"erps":   strings.Join(erps, ","),
		"title":  title,
		"msg":    msg,
		"url":    "http://xqp.jd.com/#task/item",
		"source": "ads-xqp@Hw1B92Wn",
	}
	byteData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	reader := bytes.NewReader(byteData)

	client := &http.Client{}
	url := "http://signal-api.jd.local/sendTimeline"
	req, err := http.NewRequest("POST", url, reader)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("jdos查询镜像失败:", err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	defer res.Body.Close()
}
