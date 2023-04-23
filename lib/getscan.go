package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

func Getscan(u string, thread int) {

	fmt.Println(Yellow("\n参数获取开始"))

	url := u

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	re := regexp.MustCompile(`(?m)^\s*var\s+(\w+)\b`)
	varNames := make(map[string]bool)

	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 || err != nil {
			break
		}

		matches := re.FindAllStringSubmatch(string(buf[:n]), -1)

		for _, match := range matches {
			varNames[match[1]] = true
		}
	}
	fmt.Println(Yellow("\n参数获取结束"))
	fmt.Println(Yellow("\nGET测试开始："))
	dirct := make(chan string, thread)
	var wg sync.WaitGroup
	for i := 0; i < cap(dirct); i++ {
		go worker(u, dirct, &wg)

	}
	for v := range varNames {
		wg.Add(1)
		dirct <- v

	}
	wg.Wait()
	close(dirct)
}

func worker(u string, dirct chan string, wg *sync.WaitGroup) {

	for d := range dirct {
		address := u + "?" + d + "=wa1ki0gwa1ki0gwa1ki0gwa1ki0gwa1ki0gwa1ki0gwa1ki0gwa1ki0gwa1ki0ghhhhhhh"
		resp, err := http.Get(address)
		if err != nil {
			fmt.Printf(Red("[-] " + address + " InternetError\n"))
			continue
		}

		body, _ := ioutil.ReadAll(resp.Body)
		res123 := strings.Contains(string(body), "wa1ki0gwa1ki0gwa1ki0gwa1ki0gwa1ki0gwa1ki0gwa1ki0gwa1ki0gwa1ki0ghhhhhhh")
		if resp.StatusCode != 404 {
			if res123 {
				fmt.Printf(Green("[+] %s status:%d len:%d\n"), d+" 参数值被直接输出在页面   "+u+"?"+d+"=xxxxxx", resp.StatusCode, len(body))
			}

		}

		resp.Body.Close()
		wg.Done()
	}
}
