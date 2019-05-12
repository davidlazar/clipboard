// Package clipboard manipulates the system clipboard.
package clipboard

import (
	"bytes"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func GetClipboard() ([]byte, error) {
	return exec.Command(clipboardGetCmd[0], clipboardGetCmd[1:]...).Output()
}

func SetClipboard(data []byte) error {
	cmd := exec.Command(clipboardSetCmd[0], clipboardSetCmd[1:]...)
	cmd.Stdin = bytes.NewReader(data)
	return cmd.Run()
}

func SetClipboardTemporarily(data []byte, d time.Duration) (chan error, error) {
	prev, err := GetClipboard()
	if err != nil {
		return nil, err
	}

	donechan := make(chan error, 1)
	sigchan := make(chan os.Signal, 1)
	go func() {
		select {
		case <-sigchan:
			donechan <- SetClipboard(prev)
			os.Exit(0)
		case <-time.After(d):
			donechan <- SetClipboard(prev)
		}
	}()
	signal.Notify(sigchan, os.Interrupt, os.Kill)

	if err := SetClipboard(data); err != nil {
		return nil, err
	}
	return donechan, nil
}
