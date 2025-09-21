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

	"github.com/J-Siu/go-basestruct"
	"github.com/J-Siu/go-ezlog"
)

func Get(host string, port int) *DevTools {
	d := new(DevTools)
	d.MyType = "DevtoolsInfo"
	prefix := "dq.Get"
	d.Initialized = true

	d.Host = host
	d.Port = port
	d.Url = net.JoinHostPort(d.Host, strconv.Itoa(d.Port))

	d.getVer().getPages()
	if d.Err != nil {
		ezlog.Err(d.Err.Error())
	}
	ezlog.Debug(prefix + ": Pages:")
	ezlog.DebugP(MustToJsonStrP(d.Pages))
	return d
}

type DevtoolsInfo struct {
	Browser     string `json:"Browser,omitempty"`
	ProtocolVer string `json:"Protocol-Version,omitempty"`
	UserAgent   string `json:"User-Agent,omitempty"`
	V8Ver       string `json:"V8-Version,omitempty"`
	WebKitVer   string `json:"WebKit-Version,omitempty"`
	WsUrl       string `json:"webSocketDebuggerUrl,omitempty"`

	Description string `json:"description,omitempty"`
	Title       string `json:"title,omitempty"`
	Type        string `json:"type,omitempty"`
	Url         string `json:"url,omitempty"`
}

// DevTools ws/url info
type DevTools struct {
	basestruct.Base

	Host string `json:"host,omitempty"`
	Port int    `json:"port,omitempty"`
	Url  string `json:"url,omitempty"`

	Pages []DevtoolsInfo `json:"pages,omitempty"` // Tabs with Page type only
	Tabs  []DevtoolsInfo `json:"tabs,omitempty"`  // From http://[Host]:[Port]/json
	Ver   DevtoolsInfo   `json:"ver,omitempty"`   // From http://[Host]:[Port]/json/version
}

func (d *DevTools) getVer() *DevTools {
	prefix := d.MyType + ".getVer"
	if d.CheckErrInit(prefix) {
		urlVer, _ := url.JoinPath("http://", d.Url, "json", "version")
		d.Err = HttpGetJson(urlVer, &d.Ver, 2)
	}
	return d
}

func (d *DevTools) getTabs() *DevTools {
	prefix := d.MyType + ".getTabs"
	ezlog.Trace(prefix + ": Start")

	if d.CheckErrInit(prefix) {
		urlTab, _ := url.JoinPath("http://", d.Url, "json")
		d.Err = HttpGetJson(urlTab, &d.Tabs, 2)
		ezlog.Debug(prefix)
		ezlog.DebugP(MustToJsonStrP(d.Tabs))
	}

	ezlog.Trace(prefix + ": End")
	return d
}

// Filter page type from d.Tabs into d.Pages
func (d *DevTools) getPages() *DevTools {
	prefix := d.MyType + ".getPages"
	ezlog.Trace(prefix + ": Start")

	if d.CheckErrInit(prefix) {
		d.getTabs()
		if d.Err == nil {
			// Only Keep "Page"
			d.Pages = []DevtoolsInfo{}
			for _, tab := range d.Tabs {
				if tab.Type == "page" {
					d.Pages = append(d.Pages, tab)
				}
			}
		}
		ezlog.Debug(prefix)
		ezlog.DebugP(MustToJsonStrP(d.Pages))
	}

	ezlog.Trace(prefix + ": End")
	return d
}

func HttpGetJson[T any](urlStr string, jsonObjP *T, timeout int) (err error) {
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

	ezlog.Trace(prefix)
	ezlog.TraceP(MustToJsonStrP(jsonObjP))

	return err
}

func MustToJsonStrP(obj any) *string {
	prefix := "MustToJsonStrP"
	var str string
	b, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		str = string(b)
	} else {
		ezlog.Trace(prefix + ": Err: " + err.Error())
		return nil
	}
	return &str
}
