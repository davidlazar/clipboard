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

func SetClipboardTemporarily(data []byte, d time.Duration) error {
	prev, err := GetClipboard()
	if err != nil {
		return err
	}

	sigchan := make(chan os.Signal, 1)
	go func() {
		for _ = range sigchan {
			SetClipboard(prev)
			os.Exit(0)
		}
	}()
	signal.Notify(sigchan, os.Interrupt, os.Kill)

	if err := SetClipboard(data); err != nil {
		return err
	}
	time.Sleep(d)
	if err := SetClipboard(prev); err != nil {
		return err
	}
	return nil
}
