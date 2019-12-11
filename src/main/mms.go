// @title mms
// @version 0.0.1
// @description 商品中心
// @contact.name richard sun
// @contact.email cugriver@163.com
// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:13147
// @BasePath /v1
// @schemes http https
package main

import (
    "mms/src/config"
)

var (
    commit  = "000000"
    version = "v0.0.1"
    buildTime = "2006-01-03 16:05:06"
)

const (
    ErrorLoadArgs = "error: load args failed"
    ErrorLoadConfigure = "error: load configure file failed"
)

func main() {
    var err error
    err = config.LoadArgs(commit, version, buildTime)
    if err != nil {
        config.ExitWithInfo(ErrorLoadArgs)
    }
    err = config.LoadConfigure()
    if err != nil {
        config.ExitWithInfo(ErrorLoadConfigure)
    }
}

