package bktree

import (
  "testing"
  "github.com/steakknife/hamming"
)

func TestTree(t *testing.T) {
  tree := Tree{Fn: hamming.Uint64}

  tree.Add(8992787323816807617)
  tree.Add(3545795011398387613)


  result := tree.Query(8992787323816807618, 10)

  if result[8992787323816807617] != 2 {
    t.Fatalf("unexpected value: %+v", result)
  }
}