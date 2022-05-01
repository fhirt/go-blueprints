package trace

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	msg := "Hello trace package."
	want := "Hello trace package.\n"
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		descr := fmt.Sprintf("tracer.Trace(%q)", msg)
		tracer.Trace(msg)

		if got := buf.String(); got != want {
			t.Errorf("%s = %q, want %q", descr, got, want)
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("something")
}