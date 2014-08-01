package bktree

import(
  "fmt"
  "strconv"
)

type Distancer func(uint64, uint64) int

type Node struct {
  children map[int]*Node
  term uint64
  fn Distancer
}

func (t *Node) Add(term uint64) {
  score := t.Distance(term)

  if t.children == nil {
    t.children = make(map[int]*Node)
  }

  if _, ok := t.children[score]; !ok {
    t.children[score] = &Node{term: term, fn: t.fn}
  } else {
    t.children[score].Add(term)
  }
}

func (t *Node) Query(term uint64, threshold int, collected map[uint64]int) {
  score := t.Distance(term)

  if score <= threshold {
    collected[t.term] = score
  }

  for i := -threshold; i < threshold + 1; i++ {
    child := t.children[score + i]

    if child != nil {
      child.Query(term, threshold, collected)
    }
  }
}

func (t *Node) Distance(term uint64) int {
  return t.fn(term, t.term)
}

func (t *Node) print(depth int, key string) {
  for i := 0; i < depth; i++ {
    fmt.Print(" ")
  }

  fmt.Printf("%s: %s\n", key, strconv.FormatUint(t.term, 10))

  for k, v := range t.children {
    v.print(depth+1, strconv.Itoa(k))
  }
}

type Tree struct {
  root *Node
  Fn Distancer
}

func (t *Tree) Add(term uint64) {
  if t.root != nil {
    t.root.Add(term)
  } else {
    t.root = &Node{term: term, fn: t.Fn}
  }
}

func (t *Tree) Query(term uint64, threshold int) map[uint64]int {
  collected := make(map[uint64]int)
  t.root.Query(term, threshold, collected)
  return collected
}

func (t *Tree) Print() {
  t.root.print(0, "*")
}