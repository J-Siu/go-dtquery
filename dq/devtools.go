/*
Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// dq - Devtools Query
package dq

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/J-Siu/go-helper/v2/basestruct"
	"github.com/J-Siu/go-helper/v2/ezlog"
)

// DevTools ws/url info
type DevTools struct {
	*basestruct.Base

	Host string `json:"Host"`
	Port int    `json:"Port"`
	Url  string `json:"Url"`

	Pages []DevtoolsInfo `json:"Pages"` // Tabs with Page type only
	Tabs  []DevtoolsInfo `json:"Tabs"`  // From http://[Host]:[Port]/json
	Ver   DevtoolsInfo   `json:"Ver"`   // From http://[Host]:[Port]/json/version
}

func (t *DevTools) New(host string, port int) *DevTools {
	t.Base = new(basestruct.Base)
	t.MyType = "Devtools"
	t.Host = host
	t.Port = port
	t.Url = net.JoinHostPort(t.Host, strconv.Itoa(t.Port))
	t.Initialized = true
	return t
}

// Get json info from http://<host>:<port/json/version
//
// Populate
func (t *DevTools) GetVer() *DevTools {
	prefix := t.MyType + ".getVer"
	if t.CheckErrInit(prefix) {
		urlVer, _ := url.JoinPath("http://", t.Url, "json", "version")
		t.Err = httpGetJson(urlVer, &t.Ver, 2)
	}
	return t
}

// Get json info from http://<host>:<port/json
//
// Populate both
func (t *DevTools) GetTabs() *DevTools {
	prefix := t.MyType + ".getTabs"

	ezlog.Trace().N(prefix).TxtStart().Out()
	if t.CheckErrInit(prefix) {
		urlTab, _ := url.JoinPath("http://", t.Url, "json")
		t.Err = httpGetJson(urlTab, &t.Tabs, 2)
		if t.Err == nil {
			t.getPages()
		}
		ezlog.Debug().Nn(prefix).M(t.Tabs).Out()
	}

	ezlog.Trace().N(prefix).TxtEnd().Out()
	return t
}

// Filter page type from d.Tabs into d.Pages
func (t *DevTools) getPages() *DevTools {
	prefix := t.MyType + ".getPages"
	ezlog.Trace().N(prefix).TxtStart().Out()

	if t.CheckErrInit(prefix) {
		if t.Err == nil {
			// Only Keep "Page"
			t.Pages = []DevtoolsInfo{}
			for _, tab := range t.Tabs {
				if tab.Type == "page" {
					t.Pages = append(t.Pages, tab)
				}
			}
		}
		ezlog.Trace().Nn(prefix).M(t.Pages).Out()
	}

	ezlog.Trace().N(prefix).TxtEnd().Out()
	return t
}

func httpGetJson[T any](urlStr string, jsonObjP *T, timeout int) (err error) {
	prefix := "httpGetJson"

	var body []byte
	var req *http.Request
	var res *http.Response
	client := http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

	req, err = http.NewRequest(http.MethodGet, urlStr, nil)
	if err == nil {
		res, err = client.Do(req)
	}

	if err == nil && res.Body != nil {
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		err = json.Unmarshal(body, jsonObjP)
	}

	ezlog.Trace().Nn(prefix).M(jsonObjP).Out()
	return err
}
