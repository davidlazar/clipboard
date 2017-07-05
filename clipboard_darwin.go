// +build darwin

package clipboard

var clipboardGetCmd = []string{"pbpaste"}
var clipboardSetCmd = []string{"pbcopy"}
