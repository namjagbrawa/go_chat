package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func httpGet(url string){
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func httpPostForm() {
	resp, err := http.PostForm("http://127.0.0.1:5000/login",
		url.Values{"username": {"admin"}, "password": {"111111"}})

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func main() {
	// https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	httpGet("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxb1a05ddc69040dff&secret=448b77cd710abad65950c57e85174760")
}

// 29_cyivhJbuwvTN3wh2WUlWtojGEW-EqtgBgjDKlXR6x69akIwXfDkajTrmQQLuer-yrcyRatuG0jkJZiaZIK7tmHrwkX1KWEFCyZDfABPEdW_WUSVAvHNdfuQ34qPtVx2BEMXaS0_i-f7XFxGjOYGdAFAZJR