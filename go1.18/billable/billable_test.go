package billable_test

import (
	"github.com/gocuntian/training/go1.18/billable"
	"github.com/google/go-cmp/cmp"
	"sync"
	"testing"
)

func TestSendReceive(t *testing.T) {
	t.Parallel()
	c := billable.NewChannel[int](1)
	want := 99
	c.Send(want)
	got := c.Receive()
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestChannel_Sends(t *testing.T) {
	t.Parallel()
	c := billable.NewChannel[float64](3)
	c.Send(1.0)
	c.Send(2.0)
	c.Send(3.0)
	want := 3
	got := c.Sends()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestChannel_Receives(t *testing.T) {
	t.Parallel()
	c := billable.NewChannel[struct{}](1)
	c.Send(struct{}{})
	_ = c.Receive()
	c.Send(struct{}{})
	_ = c.Receive()
	want := 2
	got := c.Receives()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestConcurrencySafety(t *testing.T) {
	t.Parallel()
	c := billable.NewChannel[string](10)
	want := 100
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < want; i++ {
			c.Send("hello")
			_ = c.Receives()
		}
		wg.Done()
	}()
	for i := 0; i < want; i++ {
		_ = c.Receive()
		_ = c.Sends()
	}
	wg.Wait()
	got := c.Sends()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
	got = c.Receives()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
