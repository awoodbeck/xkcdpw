// +build dev

package server

import "net/http"

// Assets implements file system access to the "server/assets" directory.
var Assets http.FileSystem = http.Dir("server/assets")
