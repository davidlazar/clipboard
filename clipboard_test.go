package clipboard

import (
	"bytes"
	"testing"
	"time"
)

func TestGetClipboard(t *testing.T) {
	expected, err := GetClipboard()
	ensure(t, err)
	actual, err := GetClipboard()
	ensure(t, err)

	if !bytes.Equal(expected, actual) {
		t.Fatalf("expected=%q  actual=%q", expected, actual)
	}
}

func TestSetClipboard(t *testing.T) {
	expected := []byte("test set clipboard\n")
	ensure(t, SetClipboard(expected))
	actual, err := GetClipboard()
	ensure(t, err)

	if !bytes.Equal(expected, actual) {
		t.Fatalf("expected=%q  actual=%q", expected, actual)
	}
}

func TestSetClipboardTemporarily(t *testing.T) {
	prev := []byte("nothing")
	ensure(t, SetClipboard(prev))

	tmp := []byte("temporary")
	go func() {
		ensure(t, SetClipboardTemporarily(tmp, 2*time.Second))
		cur, err := GetClipboard()
		ensure(t, err)
		if !bytes.Equal(cur, prev) {
			t.Fatalf("0: expected=%q  actual=%q", prev, cur)
		}
	}()

	time.Sleep(10 * time.Millisecond)
	cur, err := GetClipboard()
	ensure(t, err)
	if !bytes.Equal(cur, tmp) {
		t.Fatalf("1: expected=%q  actual=%q", tmp, cur)
	}

	time.Sleep(1 * time.Second)
	cur, err = GetClipboard()
	ensure(t, err)
	if !bytes.Equal(cur, tmp) {
		t.Fatalf("2: expected=%q  actual=%q", tmp, cur)
	}

	time.Sleep(1200 * time.Millisecond)
	cur, err = GetClipboard()
	ensure(t, err)
	if !bytes.Equal(cur, prev) {
		t.Fatalf("3: expected=%q  actual=%q", prev, cur)
	}
}

func ensure(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("error: %q", err)
	}
}
