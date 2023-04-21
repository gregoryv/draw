package goviz

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestListener(t *testing.T) {
	h := func(w http.ResponseWriter, r *http.Request) {
		SaveFrameSeq("testdata/http.svg")
	}

	go http.ListenAndServe(":9999", http.HandlerFunc(h))
	<-time.After(10 * time.Millisecond)
	http.Get("http://localhost:9999")
}

// ----------------------------------------
func TestHouse(t *testing.T) {
	h := NewHouse()
	ctx, cancel := context.WithCancel(context.Background())
	go h.Run(ctx)
	<-time.After(100 * time.Millisecond)
	cancel()
}

// NewHouse returns a single room house
func NewHouse() *House {
	return &House{
		floors: []*Floor{
			&Floor{
				rooms: []*Room{
					&Room{},
				},
			},
		},
	}
}

type House struct {
	floors []*Floor
}

func (h *House) Run(ctx context.Context) error {
	for {
		h.check()
		select {
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (h *House) check() error {
	for _, f := range h.floors {
		if !f.allRoomsLocked() {
			return fmt.Errorf("not all rooms locked")
		}
	}
	return nil
}

type Floor struct {
	rooms []*Room
}

func (f *Floor) allRoomsLocked() bool {
	for _, r := range f.rooms {
		if !r.locked() {
			return false
		}
	}
	return true
}

type Room struct{}

func (r *Room) locked() bool {
	d, _ := FrameSequence()
	d.SaveAs("testdata/house.svg")
	return true
}
