package ryutil

import (
	"net/url"
	"net/http"
	"io/ioutil"
)

//POST and GET

// get 网络请求
func Get(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		Error(err.Error())
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	if params != nil {
		Url.RawQuery = params.Encode()
	}
	resp, err := http.Get(Url.String())
	if err != nil {
		Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// post 网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values) (rs []byte, err error) {
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}


