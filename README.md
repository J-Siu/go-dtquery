# dq - Devtools Query

Query devtools version, pages.

- [Install](#install)
- [Usage](#usage)
  - [DQ package](#dq-package)
  - [DT\_Url](#dt_url)
  - [Command Line Demo](#command-line-demo)
- [License](#license)

<!--more-->

### Install

Go install

```sh
go install github.com/J-Siu/go-dtquery@latest
```

Download

- https://github.com/J-Siu/go-dtquery/releases

### Usage

#### DQ package

```sh
go get github.com/J-Siu/go-dtquery/dq
```

```go
import "github.com/J-Siu/go-dtquery/dq"
```

See [root.go](/cmd/root.go) for code sample.

#### DT_Url

As Chrome/Chromium name(key) of the devtool url in http//localhost:9222/json/version may change overtime, starting dq v1.2.0, AFTER calling `dq.Get(host, port)`, the devtool url can be access directly with `dq.DT_Url`:

```go
devtools := dq.Get(host, port)
devtool_url := dq.DT_Url()
# # or
# devtool_url := devtools.DT_Url
```

Current support devtools url json name:
- WsUrl (Chrome < 140)
- WebSocketDebuggerUrl (Chrome >= 140)

#### Command Line Demo

1. Close running Chrome. Start Chrome with following option:

    ```sh
    chrome --remote-debugging-port=9222
    ```

    Or Chromium with the same option.

2. Run

    ```sh
    go run main.go [-r <hostname/ip>] [-p <port>]
    ```

### License

The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
