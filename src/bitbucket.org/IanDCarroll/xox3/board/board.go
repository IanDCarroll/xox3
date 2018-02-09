package board

type Board interface {
  View() []int
  Mark(int, int)
}

type board struct {
  spaces []int
}

func NewMNKX(size int) board {
  squaredSize := size * size
  spaces := make([]int, squaredSize)
  return board{spaces}
}

func (b board) View() []int { return b.spaces }

func (b board) Mark(space, player int) {
  b.spaces[space] = player
}
