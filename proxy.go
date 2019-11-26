package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Proxy struct {
	Ip       string `json:"ip"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
	Error    string `json:"error"`
}

func GetProxy() (*url.URL, error) {
	res, err := http.Get("https://api.getproxylist.com/proxy?protocol[]=http")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	proxy := new(Proxy)
	err = json.Unmarshal(data, proxy)
	if err != nil {
		return nil, err
	}
	if proxy.Error != "" {
		return nil, fmt.Errorf("%s", proxy.Error)
	}

	return url.Parse(proxy.String())
}

func (p *Proxy) String() string {
	return fmt.Sprintf("%s://%s:%d", p.Protocol, p.Ip, p.Port)
}
