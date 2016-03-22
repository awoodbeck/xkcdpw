// +build dev

package assets

import "net/http"

// Assets implements file system access to the "files" directory.
var Assets http.FileSystem = http.Dir("files")
