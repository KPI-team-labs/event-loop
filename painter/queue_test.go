package painter

import (
	"testing"
)

func TestMessageQueueEmpty(t *testing.T) {
	mq := &messageQueue{}
	if !mq.empty() {
		t.Errorf("Expected empty message queue, got non-empty.")
	}
}

func TestMessageQueuePush(t *testing.T) {
	mq := &messageQueue{}
	op1 := new(MockOperation)
	mq.push(op1)
	if mq.empty() {
		t.Errorf("Expected non-empty message queue, got empty.")
	}
}

func TestMessageQueuePull(t *testing.T) {
	mq := &messageQueue{}
	op1 := new(MockOperation)
	mq.push(op1)
	op2 := mq.pull()
	if op1 != op2 {
		t.Errorf("Expected %v, got %v.", op1, op2)
	}
	if !mq.empty() {
		t.Errorf("Expected empty message queue after pull, got non-empty.")
	}
}
