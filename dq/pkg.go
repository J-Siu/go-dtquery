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
	"github.com/J-Siu/go-helper/v2/ezlog"
)

var (
	devTools = new(DevTools)
)

// Return a DevTools object with `Ver`, `Tabs` and `Pages` populated
func Get(host string, port int) *DevTools {
	prefix := "dq.Get"
	d := devTools.New(host, port).GetVer().GetTabs()
	if d.Err == nil {
		ezlog.Debug().N(prefix).Nn("Pages").M(d).Out()
	}
	return d
}

// This should be called after `Get`
func DT_Url() (dt_url string) {
	return devTools.DT_Url
}
