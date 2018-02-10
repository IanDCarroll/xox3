package board

type Board interface {
  View() []int
  Mark(int, int)
  AvailableSpaces() []int
  Size() int
}

type board struct {
  spaces []int
}

func NewMNKX(mnkValue int) board {
  size := mnkValue * mnkValue
  spaces := make([]int, size)
  return board{spaces}
}

func (b board) View() []int { return b.spaces }

func (b board) Mark(space, player int) {
  b.spaces[space] = player
}

func (b board) AvailableSpaces() []int {
  available := make([]int, 0)
  for i := 0; i < b.Size(); i++ {
    if b.spaces[i] == 0 { available = append(available, i) }
  }
  return available
}

func (b board) Size() int {
  return len(b.spaces)
}
