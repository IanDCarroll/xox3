package rules

type Rules interface {
  Gameover([]int) (bool, int)
}

type rules struct {
  waysToWin [][]int
}

func NewMNKX(size int) rules {
  return rules{constructMNKXRules(size)}
}

func (r rules) Gameover(board []int) (bool, int) {
  winningPlayer := r.idAnyWinner(board)
  finishedGame := r.thereIsA(winningPlayer) || r.noMovesLeft(board)
  return finishedGame, winningPlayer
}

func (r rules) idAnyWinner(board []int) int {
  noWinner := 0
  for i := 0; i < len(r.waysToWin); i++ {
    sliceOfGame := r.getThisSliceOfGame(r.waysToWin[i], board)
    if r.allValuesEqual(sliceOfGame) { return sliceOfGame[0] }
  }
  return noWinner
}

func (r rules) getThisSliceOfGame(wayToWin, board []int) []int {
  sliceOfGame := make([]int, 0)
  for i := 0; i < len(wayToWin); i++ {
    sliceOfGame = append(sliceOfGame, board[wayToWin[i]])
  }
  return sliceOfGame
}

func (r rules) allValuesEqual(sliceOfGame []int) bool {
  for i := 0; i < len(sliceOfGame); i++ {
    if r.notValidAndEqual(sliceOfGame[i], sliceOfGame[0]) { return false }
  }
  return true
}

func (r rules) notValidAndEqual(i, comparison int) bool {
    return i == 0 || i != comparison
}

func (r rules) thereIsA(winningPlayer int) bool {
  return winningPlayer != 0
}

func (r rules) noMovesLeft(board []int) bool {
  noMoves := true
  for i := 0; i < len(board); i++ {
    if board[i] == 0 { noMoves = false }
  }
  return noMoves
}
