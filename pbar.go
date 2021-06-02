// Package pbar implements a simple terminal progress bar.
package pbar

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

const (
	fill       = "█"
	fillLength = 50
	empty      = "-"
)

// ProgressBar tracks the progress of tasks, displaying the # completed / count
// as well as the percentage. Output defaults to Stderr, but is configurable.
type ProgressBar struct {
	incrementedCount int
	totalCount       int
	out              io.Writer
	mu               *sync.Mutex
}

// New creates a ProgressBar, erroring if the count is not greater than 0.
func New(count int) (*ProgressBar, error) {
	if count <= 0 {
		return nil, errors.New("count must be greater than 0")
	}

	return &ProgressBar{
		totalCount: count,
		out:        os.Stderr,
		mu:         &sync.Mutex{},
	}, nil
}

// SetOutput determines where to write progress.
func (p *ProgressBar) SetOutput(w io.Writer) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.out = w
}

// Start writes the initial bar.
func (p *ProgressBar) Start() {
	p.Increment(0)
}

// End writes a new line to end the bar.
func (p *ProgressBar) End() {
	fmt.Fprintln(p.out)
}

// Increment refreshes the bar with updated progress.
func (p *ProgressBar) Increment(count int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	switch {
	case count > 0 && p.incrementedCount > p.totalCount-count:
		p.incrementedCount = p.totalCount
	case count < 0 && p.incrementedCount < -count:
		p.incrementedCount = 0
	default:
		p.incrementedCount = p.incrementedCount + count
	}

	var (
		completedLength = (fillLength * p.incrementedCount) / p.totalCount
		remainingLength = fillLength - completedLength
		bar             = make([]string, completedLength+remainingLength)
	)

	for i := 0; i < completedLength; i++ {
		bar[i] = fill
	}

	for i := completedLength; i < completedLength+remainingLength; i++ {
		bar[i] = empty
	}

	percent := 100 * (float64(p.incrementedCount) / float64(p.totalCount))

	fmt.Fprintf(
		p.out,
		"Progress %d / %d: |%s| %.2f%% Complete\r",
		p.incrementedCount,
		p.totalCount,
		strings.Join(bar, ""),
		percent,
	)
}
