// +build linux

package clipboard

var clipboardGetCmd = []string{"xclip", "-out"}
var clipboardSetCmd = []string{"xclip"}
