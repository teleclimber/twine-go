package twine

import "testing"

func TestIncrementNextID(t *testing.T) {
	r := &messageRegistry{
		firstMsgID: 128,
		lastMsgID:  255,
		nextID:     250,
	}

	r.incrementNextID()
	if r.nextID != 251 {
		t.Errorf("expected next id to be 251, got %v", r.nextID)
	}

	for i := 0; i < 5; i++ {
		r.incrementNextID()
	}
	if r.nextID != 128 {
		t.Errorf("expected next id to be 128, got %v", r.nextID)
	}
}
