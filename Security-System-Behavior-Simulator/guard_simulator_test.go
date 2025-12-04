package main

import (
	"bytes"
	"io"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

func TestRunSimulation_Output(t *testing.T) {
	rand.Seed(1)

	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w

	done := make(chan struct{})
	go func() {
		runSimulation()
		_ = w.Close()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(3 * time.Second):
		_ = w.Close()
		os.Stdout = old
		t.Fatal("runSimulation timed out")
	}

	os.Stdout = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	out := buf.String()

	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 9 {
		t.Fatalf("expected 9 lines of output, got %d; output:\n%s", len(lines), out)
	}
	for _, name := range []string{"Амир", "Жалгас", "Мурад"} {
		if !strings.Contains(out, name) {
			t.Fatalf("output does not contain guard name %q", name)
		}
	}
	for _, l := range lines {
		if !strings.Contains(l, "Угроза:") {
			t.Fatalf("line missing 'Угроза:': %q", l)
		}
	}
}
