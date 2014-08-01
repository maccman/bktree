package tree

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

  pow := make([]int, threshold)

  for i := range pow {
    child := t.children[score + i]

    if child != nil {
      child.Query(term, threshold, collected)
    }
  }
}

func (t *Node) Distance(term uint64) int {
  return t.fn(term, t.term)
}

type Tree struct {
  root *Node
  fn Distancer
}

func (t *Tree) Add(term uint64) {
  if t.root != nil {
    t.root.Add(term)
  } else {
    t.root = &Node{term: term, fn: t.fn}
  }
}

func (t *Tree) Query(term uint64, threshold int) map[uint64]int {
  collected := make(map[uint64]int)
  t.root.Query(term, threshold, collected)
  return collected
}