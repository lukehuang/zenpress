package qimport

import (
	"qlang.io/lib/bufio"
	"qlang.io/lib/bytes"
	"qlang.io/lib/crypto/md5"
	"qlang.io/lib/encoding/hex"
	"qlang.io/lib/encoding/json"
	"qlang.io/lib/eqlang"
	"qlang.io/lib/errors"
	"qlang.io/lib/io"
	"qlang.io/lib/io/ioutil"
	"qlang.io/lib/math"
	"qlang.io/lib/meta"
	"qlang.io/lib/net/http"
	"qlang.io/lib/os"
	"qlang.io/lib/path"
	"qlang.io/lib/reflect"
	"qlang.io/lib/runtime"
	"qlang.io/lib/strconv"
	"qlang.io/lib/sync"
	"qlang.io/lib/terminal"
	"qlang.io/lib/tpl/extractor"
	"qlang.io/lib/version"
	qlang "qlang.io/spec"

	// qlang builtin modules
	_ "qlang.io/lib/builtin"
	_ "qlang.io/lib/chan"

	extFmt "github.com/insionng/zenpress/extension/fmt"
	extStrings "github.com/insionng/zenpress/extension/strings"

	"github.com/insionng/zenpress/extension/github.com/insionng/makross"
	"github.com/insionng/zenpress/extension/github.com/insionng/makross/cache"
	"github.com/insionng/zenpress/extension/github.com/insionng/makross/captcha"
	"github.com/insionng/zenpress/extension/github.com/insionng/makross/logger"
	"github.com/insionng/zenpress/extension/github.com/insionng/makross/pongor"
	"github.com/insionng/zenpress/extension/github.com/insionng/makross/session"
	"github.com/insionng/zenpress/extension/github.com/insionng/makross/static"
	"github.com/insionng/zenpress/extension/github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/extension/github.com/insionng/zenpress/model"
	"github.com/insionng/zenpress/extension/github.com/insionng/zenpress/module/hook"
	"github.com/insionng/zenpress/extension/github.com/insionng/zenpress/module/switchr"
)

// -----------------------------------------------------------------------------

// Copyright prints qlang copyright information.
//
func Copyright() {
	version.Copyright()
}

// InitSafe inits qlang and imports modules.
//
func InitSafe(safeMode bool) {

	qlang.SafeMode = safeMode

	qlang.Import("", math.Exports) // import math as builtin package
	qlang.Import("", meta.Exports) // import meta package
	qlang.Import("bufio", bufio.Exports)
	qlang.Import("bytes", bytes.Exports)
	qlang.Import("md5", md5.Exports)
	qlang.Import("io", io.Exports)
	qlang.Import("ioutil", ioutil.Exports)
	qlang.Import("hex", hex.Exports)
	qlang.Import("json", json.Exports)
	qlang.Import("errors", errors.Exports)
	qlang.Import("eqlang", eqlang.Exports)
	qlang.Import("math", math.Exports)
	qlang.Import("os", os.Exports)
	qlang.Import("", os.InlineExports)
	qlang.Import("path", path.Exports)
	qlang.Import("http", http.Exports)
	qlang.Import("reflect", reflect.Exports)
	qlang.Import("runtime", runtime.Exports)
	qlang.Import("strconv", strconv.Exports)
	//qlang.Import("strings", strings.Exports)
	qlang.Import("sync", sync.Exports)
	qlang.Import("terminal", terminal.Exports)
	qlang.Import("extractor", extractor.Exports)
	/*---------------------------------------------*/
	qlang.Import("makross", makross.Exports)
	qlang.Import("logger", logger.Exports)
	qlang.Import("switchr", switchr.Exports)
	qlang.Import("pongor", pongor.Exports)
	qlang.Import("static", static.Exports)
	qlang.Import("session", session.Exports)
	qlang.Import("captcha", captcha.Exports)
	qlang.Import("cache", cache.Exports)
	qlang.Import("model", model.Exports)
	qlang.Import("helper", helper.Exports)

	qlang.Import("hook", hook.Exports)
	qlang.Import("fmt", extFmt.Exports)
	qlang.Import("strings", extStrings.Exports)
}

// -----------------------------------------------------------------------------
