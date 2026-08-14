package batcher_test

import (
	"testing"

	"github.com/LYH2263/go-event-batch/internal/batcher"
	"github.com/LYH2263/go-event-batch/internal/event"
)

func TestSplitPreservesEachBatchContents(t *testing.T) {
	in := []event.Event{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}}
	batches := batcher.Split(in, 2)
	if len(batches) != 2 {
		t.Fatalf("batches=%d", len(batches))
	}
	if batches[0].Events[0].ID != 1 || batches[0].Events[1].ID != 2 {
		t.Fatalf("batch0 polluted: %+v", batches[0].Events)
	}
	if batches[1].Events[0].ID != 3 || batches[1].Events[1].ID != 4 {
		t.Fatalf("batch1: %+v", batches[1].Events)
	}
}
