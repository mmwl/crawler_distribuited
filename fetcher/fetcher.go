/**
	只是负责将网站文本转换为utf-8
 */
package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"golang.org/x/net/html/charset"
	"bufio"
	"regexp"
	"log"
	"time"
)

var rateLimiter = time.Tick(60 * time.Millisecond)

func Fetch(url string) ([]byte, error) {

	<- rateLimiter

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	//增加header选项
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9.0.3) Gecko/2008092417 Firefox/3.0.3")
	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, _ := client.Do(reqest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)

	//监测网页是什么编码
	e := determineEncoding(bodyReader)

	//要下载gopm get -g -v golang.org/x/text 这个库
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, v := range matches {
		fmt.Println(string(v[1]), " | ", string(v[2]))
	}
	fmt.Println(len(matches))
}
