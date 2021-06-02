package pbar

import (
	"bytes"
	"io"
	"testing"
)

func TestPBar(t *testing.T) {
	t.Parallel()

	p, err := New(3)
	if err != nil {
		t.Fatal(err)
	}

	var b bytes.Buffer
	p.SetOutput(&b)

	p.Start()
	for i := 0; i < 3; i++ {
		p.Increment(1)
	}
	p.End()

	got := b.String()
	expected := "Progress 0 / 3: |--------------------------------------------------| 0.00% Complete\rProgress 1 / 3: |████████████████----------------------------------| 33.33% Complete\rProgress 2 / 3: |█████████████████████████████████-----------------| 66.67% Complete\rProgress 3 / 3: |██████████████████████████████████████████████████| 100.00% Complete\r\n"

	if expected != got {
		t.Fatalf("expected'%s'\n, got'%s'\n", expected, got)
	}
}

func TestTotalNotGreaterThanZero(t *testing.T) {
	t.Parallel()

	_, err := New(0)
	if err == nil {
		t.Fatal("expected an error when total count <= 0, but did not get one")
	}
}

func TestIncrementGreaterThanTotal(t *testing.T) {
	t.Parallel()

	p, err := New(1)
	if err != nil {
		t.Fatal(err)
	}

	p.SetOutput(io.Discard)

	p.Start()

	p.Increment(10)

	expected := 1
	got := p.incrementedCount

	if expected != got {
		t.Fatalf("expected '%d', got '%d'", expected, got)
	}
}

func TestIncrementTotalNotGreaterThanZero(t *testing.T) {
	t.Parallel()

	p, err := New(3)
	if err != nil {
		t.Fatal(err)
	}

	p.SetOutput(io.Discard)

	p.Start()

	p.Increment(1)
	p.Increment(-100)

	expected := 0
	got := p.incrementedCount

	if expected != got {
		t.Fatalf("expected '%d', got '%d'", expected, got)
	}
}
