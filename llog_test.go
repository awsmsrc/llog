package llog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	stdo := os.Stdout
	r, w, _ := os.Pipe()
	expected := new(bytes.Buffer)
	os.Stdout = w

	_, _ = fmt.Fprintf(
		expected,
		"%v[%v] %v%v\n",
		"\033[36m",
		time.Now().Format("2006/01/02-15:04:05"),
		"boo",
		"\033[0m",
	)

	Info("boo")

	outC := make(chan []byte)
	// copy the output in a separate goroutine so printing can't block
	// indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.Bytes()
	}()

	// back to normal state
	w.Close()
	os.Stdout = stdo // restoring  the real stdout
	out := <-outC

	if !(bytes.Equal(expected.Bytes(), out)) {
		t.Errorf("\nexp %+v\n got %+v\n", expected.Bytes(), out)
	}
}

func TestLoggerInfo(t *testing.T) {
	output := new(bytes.Buffer)
	expected := new(bytes.Buffer)

	_, _ = fmt.Fprintf(
		expected,
		"%v[%v] %v%v\n",
		"\033[36m",
		time.Now().Format("2006/01/02-15:04:05"),
		"boo",
		"\033[0m",
	)

	logger := &Logger{
		output,
	}
	logger.Info("boo")

	if !(bytes.Equal(expected.Bytes(), output.Bytes())) {
		t.Errorf("\nexp %+v\n got %+v\n", expected.Bytes(), output.Bytes())
	}
}

func TestLoggerInfof(t *testing.T) {
	output := new(bytes.Buffer)
	expected := new(bytes.Buffer)

	_, _ = fmt.Fprintf(
		expected,
		"%v[%v] %v%v\n",
		"\033[36m",
		time.Now().Format("2006/01/02-15:04:05"),
		"boo 1 2",
		"\033[0m",
	)

	logger := &Logger{
		output,
	}
	logger.Infof("boo %d %d", 1, 2)

	if !(bytes.Equal(expected.Bytes(), output.Bytes())) {
		t.Errorf("\nexp %+v\n got %+v\n", expected.String(), output.String())
	}
}
