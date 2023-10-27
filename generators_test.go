package generators

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	done := make(chan struct{})

	repeatStream := Repeat(done, 1, 2, 3)
	expectedValues := []int{1, 2, 3, 1, 2, 3, 1, 2, 3}

	for _, v := range expectedValues {
		if k, ok := <-repeatStream; k != v || !ok {
			t.Error("Repeat() did not repeat correctly")
		}
	}

	close(done)
	if k, ok := <-repeatStream; k != 0 || ok {
		t.Error("Repeat() did not close correctly")
	}
}

func TestRepeatEmpty(t *testing.T) {
	done := make(chan struct{})
	repeatStream := Repeat(done, []int{}...)
	if k, ok := <-repeatStream; k != 0 || ok {
		t.Error("Repeat() from empty slice did not close")
	}
	close(done)
	if k, ok := <-repeatStream; k != 0 || ok {
		t.Error("Repeat() did not close correctly")
	}
}

func TestTake(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	repeatStream := Repeat(done, 1, 2, 3)
	takeStream := Take(done, repeatStream, 5)
	expectedValues := []int{1, 2, 3, 1, 2}

	for _, v := range expectedValues {
		if k, ok := <-takeStream; k != v || !ok {
			t.Error("Take() did not take correctly")
		}
	}
	if k, ok := <-takeStream; k != 0 || ok {
		t.Error("Take() did not close correctly")
	}
}

func TestRepeatFn(t *testing.T) {
	done := make(chan struct{})

	repeatStream := RepeatFn(done, func() int { return 1 })
	expectedValues := []int{1, 1, 1, 1, 1}

	for _, v := range expectedValues {
		if k, ok := <-repeatStream; k != v || !ok {
			t.Error("RepeatFn() did not repeat correctly")
		}
	}

	close(done)
	if k, ok := <-repeatStream; k != 0 || ok {
		t.Error("RepeatFn() did not close correctly")
	}
}
