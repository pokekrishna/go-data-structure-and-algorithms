package singlylinkedlist_test

import (
	"github.com/pokekrishna/dsa/data-structures/linear/linkedlist/singlylinkedlist"
	"testing"
)

func Test_create(t *testing.T) {
	tests := []struct {
		name string
	}{
		{""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			singlylinkedlist.CreateDummy()
		})
	}
}