package dontpanic

import (
	"errors"
	"sync"
	"testing"
)

func TestGo_NormalExecution(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	Go(func() {
		defer wg.Done()
		// Normal execution
	})

	wg.Wait()
	// Test passes if no panic occurs
}

func TestGo_PanicRecovery(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan any, 1)

	Go(func() {
		defer wg.Done()
		panic("test panic")
	}, WithRecover(func(r any) {
		ch <- r
	}))

	wg.Wait()

	select {
	case r := <-ch:
		if r != "test panic" {
			t.Errorf("Expected 'test panic', got %v", r)
		}
	default:
		t.Error("Recovery function was not called")
	}
}

func TestGo_CustomRecover(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	testErr := errors.New("test error")
	ch := make(chan any, 1)

	Go(func() {
		defer wg.Done()
		panic(testErr)
	}, WithRecover(func(r any) {
		ch <- r
	}))

	wg.Wait()

	if r := <-ch; r != testErr {
		t.Errorf("Expected test error, got %v", r)
	}
}
