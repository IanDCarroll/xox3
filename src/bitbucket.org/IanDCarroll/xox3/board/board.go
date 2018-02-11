package board

type Board interface {
  View() []int
  Mark(int, int)
  AvailableSpaces() []int
  Size() int
  Width() int
  Height() int
}

type board struct {
  spaces []int
  width int
  height int
}

func NewMNKX(mnkValue int) board {
  width := mnkValue
  height := mnkValue
  size := height * width
  spaces := make([]int, size)
  return board{spaces, width, height}
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

func (b board) Size() int { return len(b.spaces) }

func (b board) Width() int { return b.width }

func (b board) Height() int { return b.height }
