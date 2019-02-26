package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
	"net/http"
)

func Fetcher(url string) ([]byte, error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err!=nil {
		fmt.Printf("wrong http request: %s", err.Error())
		return nil, fmt.Errorf("wrong http request: %s", err.Error())
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error status code:", resp.StatusCode)
		return nil, fmt.Errorf("wrong http status code: %d", resp.StatusCode)
	}

	// 编码转换
	//htmlEncoding:=determineEncoding(resp.Body)
	//reader := transform.NewReader(resp.Body, htmlEncoding.NewDecoder())
	//body, err := ioutil.ReadAll(reader)

	return ioutil.ReadAll(resp.Body)
}

// 自动判断html charset 并返回encoding.Encoding
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
