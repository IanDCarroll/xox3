package ai

import (
  "testing"
  "reflect"
  "bitbucket.org/IanDCarroll/xox3/rules"
  "bitbucket.org/IanDCarroll/xox3/board"
)

func TestAiIdentityDefaultsToPlayingFirst(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  //When
  actualID := subject.identity
  actualOP := subject.opponent
  //Then
  expectedID := 1
  expectedOP := 2
  assertIntEqual(t, actualID, expectedID)
  assertIntEqual(t, actualOP, expectedOP)
}

func TestAiIdentityCanBeSetToPlayingSecond(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  //When
  subject.SetID(2)
  //Then
  actualID := subject.identity
  actualOP := subject.opponent
  expectedID := 2
  expectedOP := 1
  assertIntEqual(t, actualID, expectedID)
  assertIntEqual(t, actualOP, expectedOP)
}

func TestAiIdentityIsUnaffectedBySubesequentSetIDCalls(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  //When
  subject.SetID(2)
  subject.SetID(2)
  //Then
  actualID := subject.identity
  actualOP := subject.opponent
  expectedID := 2
  expectedOP := 1
  assertIntEqual(t, actualID, expectedID)
  assertIntEqual(t, actualOP, expectedOP)
}

func TestAiIdentityIsChangedIfSubesequentSetIDCallsAreDifferent(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  //When
  subject.SetID(2)
  subject.SetID(1)
  //Then
  actualID := subject.identity
  actualOP := subject.opponent
  expectedID := 1
  expectedOP := 2
  assertIntEqual(t, actualID, expectedID)
  assertIntEqual(t, actualOP, expectedOP)
}

func TestAiIdentityIsUnaffectedBySetIDCallsWithDefaultValues(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  //When
  subject.SetID(1)
  //Then
  actualID := subject.identity
  actualOP := subject.opponent
  expectedID := 1
  expectedOP := 2
  assertIntEqual(t, actualID, expectedID)
  assertIntEqual(t, actualOP, expectedOP)
}

func TestAppropritatePlayerReturns1IfTheBoardIsEmpty(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  //When
  actual := subject.appropriatePlayer(board)
  //Then
  expected := 1
  assertIntEqual(t, actual, expected)
}

func TestAppropritatePlayerReturns2IfTheBoardHasBeenMarkedOnce(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  board.Mark(0, 1)
  //When
  actual := subject.appropriatePlayer(board)
  //Then
  expected := 2
  assertIntEqual(t, actual, expected)
}

func TestAppropritatePlayerReturns1IfTheBoardHasBeenMarkedTwice(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  board.Mark(0, 1)
  board.Mark(1, 2)
  //When
  actual := subject.appropriatePlayer(board)
  //Then
  expected := 1
  assertIntEqual(t, actual, expected)
}

func TestAppropritatePlayerReturns2IfTheBoardHasBeenMarkedThrice(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  board.Mark(0, 1)
  board.Mark(1, 2)
  board.Mark(2, 1)
  //When
  actual := subject.appropriatePlayer(board)
  //Then
  expected := 2
  assertIntEqual(t, actual, expected)
}


func TestMaxScoreReturnsTheMaxValueFromASlice(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  scores := map[int]int{ 1:0, 3:-1, 5:-1, 6:-1, 7:-1, 8:-1 }
  //When
  actual := subject.maxScore(scores)
  //Then
  expected := 0
  assertIntEqual(t, actual, expected)
}

func TestMinScoreReturnsTheMinValueFromASlice(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  scores := map[int]int{ 1:0, 3:-1, 5:-1, 6:-1, 7:-1, 8:-1 }
  //When
  actual := subject.minScore(scores)
  //Then
  expected := -1
  assertIntEqual(t, actual, expected)
}

func TestScoreTheSpaceReturnsTheMaxValueIftheDepthIsOdd(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  scores := map[int]int{ 1:0, 3:-1, 5:-1, 6:-1, 7:-1, 8:-1 }
  depth := 2
  //When
  actual := subject.scoreTheSpace(scores, depth)
  //Then
  expected := 0
  assertIntEqual(t, actual, expected)
}

func TestScoreTheSpaceReturnsTheMinValueIftheDepthIsEven(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  scores := map[int]int{ 1:0, 3:-1, 5:-1, 6:-1, 7:-1, 8:-1 }
  depth := 3
  //When
  actual := subject.scoreTheSpace(scores, depth)
  //Then
  expected := -1
  assertIntEqual(t, actual, expected)
}

func TestOptimumChoiceWhenTheBoardIsAlreadyDrawn(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  b := board.NewMNKX(3)
  subject := NewMNKAi(rules, b)
  subject.SetID(2)
  b.Mark(0,2)
  b.Mark(1,1)
  b.Mark(2,2)
  b.Mark(3,2)
  b.Mark(4,1)
  b.Mark(5,1)
  b.Mark(6,1)
  b.Mark(7,2)
  b.Mark(8,1)
  boardCopy := board.Board(b)
  depth := 0
  choices := map[int]int{}
  //When
  actualChoice := subject.optimumChoice(boardCopy, depth, choices)
  //Then
  expectedChoice := 0
  assertIntEqual(t, actualChoice, expectedChoice)
}

func TestOptimumChoiceWhenTheBoardIsLostAtTheLastMoveAndAiIsP2(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  b := board.NewMNKX(3)
  subject := NewMNKAi(rules, b)
  subject.SetID(2)
  b.Mark(0,1)
  b.Mark(1,1)
  b.Mark(2,1)
  b.Mark(3,1)
  b.Mark(4,2)
  b.Mark(5,2)
  b.Mark(6,1)
  b.Mark(7,2)
  b.Mark(8,2)
  boardCopy := board.Board(b)
  depth := 0
  choices := map[int]int{}
  //When
  actualChoice := subject.optimumChoice(boardCopy, depth, choices)
  //Then
  expectedChoice := -1
  assertIntEqual(t, actualChoice, expectedChoice)
}

func TestOptimumChoiceWhenTheBoardIsWonAtTheLastMoveAndAiIsP1(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  b := board.NewMNKX(3)
  subject := NewMNKAi(rules, b)
  subject.SetID(1)
  b.Mark(0,1)
  b.Mark(1,1)
  b.Mark(2,1)
  b.Mark(3,1)
  b.Mark(4,2)
  b.Mark(5,2)
  b.Mark(6,1)
  b.Mark(7,2)
  b.Mark(8,2)
  boardCopy := board.Board(b)
  depth := 0
  choices := map[int]int{}
  //When
  actualChoice := subject.optimumChoice(boardCopy, depth, choices)
  //Then
  expectedChoice := 100
  assertIntEqual(t, actualChoice, expectedChoice)
}

func TestOptimumChoiceWhenTheBoardIsLostEarlyAndAiIsP2(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  b := board.NewMNKX(3)
  subject := NewMNKAi(rules, b)
  subject.SetID(2)
  b.Mark(0,1)
  b.Mark(1,1)
  b.Mark(2,1)
  //3,0
  //4,0
  //5,0
  //6,0
  b.Mark(7,2)
  b.Mark(8,2)
  boardCopy := board.Board(b)
  depth := 0
  choices := map[int]int{}
  //When
  actualChoice := subject.optimumChoice(boardCopy, depth, choices)
  //Then
  expectedChoice := -1
  assertIntEqual(t, actualChoice, expectedChoice)
}

func TestOptimumChoiceWhenTheBoardIsWonEarlyAndAiIsP2(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  b := board.NewMNKX(3)
  subject := NewMNKAi(rules, b)
  subject.SetID(2)
  b.Mark(0,1)
  b.Mark(1,1)
  //2,0
  //3,0
  b.Mark(4,1)
  //5,0
  b.Mark(6,2)
  b.Mark(7,2)
  b.Mark(8,2)
  boardCopy := board.Board(b)
  depth := 0
  choices := map[int]int{}
  //When
  actualChoice := subject.optimumChoice(boardCopy, depth, choices)
  //Then
  expectedChoice := 100
  assertIntEqual(t, actualChoice, expectedChoice)
}

func TestOptimumChoiceWhenTheBoardWillBeDrawn(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  b := board.NewMNKX(3)
  subject := NewMNKAi(rules, b)
  subject.SetID(1)
  b.Mark(0,2)
  b.Mark(1,1)
  b.Mark(2,2)
  b.Mark(3,2)
  b.Mark(4,1)
  b.Mark(5,1)
  //6,0
  b.Mark(7,2)
  b.Mark(8,1)
  boardCopy := board.Board(b)
  depth := 0
  choices := map[int]int{}
  //When
  actualChoice := subject.optimumChoice(boardCopy, depth, choices)
  //Then
  expectedChoice := 6
  assertIntEqual(t, actualChoice, expectedChoice)
}

func TestOptimumChoiceWhenTheBoardWillBeWonAtTheLastMove(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  b := board.NewMNKX(3)
  subject := NewMNKAi(rules, b)
  subject.SetID(1)
  //0,0
  b.Mark(1,1)
  b.Mark(2,1)
  b.Mark(3,1)
  b.Mark(4,2)
  b.Mark(5,2)
  b.Mark(6,1)
  b.Mark(7,2)
  b.Mark(8,2)
  boardCopy := board.Board(b)
  depth := 0
  choices := map[int]int{}
  //When
  actualChoice := subject.optimumChoice(boardCopy, depth, choices)
  //Then
  expectedChoice := 0
  assertIntEqual(t, actualChoice, expectedChoice)
}

func TestOptimumChoiceWhenAiIsP2AndThereAre2MovesLeftOneDrawsAndOneLoses(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  b := board.NewMNKX(3)
  subject := NewMNKAi(rules, b)
  subject.SetID(2)
  //0,0
  b.Mark(1,1)
  b.Mark(2,2)
  //3,0
  b.Mark(4,1)
  b.Mark(5,1)
  b.Mark(6,1)
  b.Mark(7,2)
  b.Mark(8,2)
  boardCopy := board.Board(b)
  depth := 0
  choices := map[int]int{}
  //When
  actualChoice := subject.optimumChoice(boardCopy, depth, choices)
  //then
  expectedChoice := 3
  assertIntEqual(t, actualChoice, expectedChoice)
}

func TestOptimumChoiceWhenAiIsP2AndThereAre2MovesLeftOneWinsAndOneDraws(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  b := board.NewMNKX(3)
  subject := NewMNKAi(rules, b)
  subject.SetID(2)
  b.Mark(0,2)
  b.Mark(1,1)
  b.Mark(2,1)
  b.Mark(3,1)
  b.Mark(4,2)
  b.Mark(5,2)
  b.Mark(6,1)
  //7,0
  //8,0
  boardCopy := board.Board(b)
  depth := 0
  choices := map[int]int{}
  //When
  actualChoice := subject.optimumChoice(boardCopy, depth, choices)
  //Then
  expectedChoice := 8
  assertIntEqual(t, actualChoice, expectedChoice)
}

func TestOptimumChoiceWhenAiIsP2AndThereAre2MovesLeftOneWinsOneLoses(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  b := board.NewMNKX(3)
  subject := NewMNKAi(rules, b)
  subject.SetID(2)
  b.Mark(0,1)
  b.Mark(1,1)
  //2,0
  b.Mark(3,1)
  b.Mark(4,2)
  b.Mark(5,1)
  //6,0
  b.Mark(7,2)
  b.Mark(8,2)
  boardCopy := board.Board(b)
  depth := 0
  choices := map[int]int{}
  //When
  actualChoice := subject.optimumChoice(boardCopy, depth, choices)
  //then
  expectedChoice := 6
  assertIntEqual(t, actualChoice, expectedChoice)
}

func TestOptimumChoiceWithTestBoardNumber1(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  b := board.NewMNKX(3)
  subject := NewMNKAi(rules, b)
  subject.SetID(1)
  b.Mark(0,1)
  b.Mark(1,2)
  b.Mark(2,1)
  b.Mark(3,2)
  b.Mark(4,1)
  b.Mark(5,2)
  //6,0
  b.Mark(7,1)
  b.Mark(8,2)
  boardCopy := board.Board(b)
  depth := 0
  choices := map[int]int{}
  //When
  actualChoice := subject.optimumChoice(boardCopy, depth, choices)
  //Then
  expectedChoice := 6
  assertIntEqual(t, actualChoice, expectedChoice)
}

func TestOptimumChoiceWithTestBoardNumber2(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  b := board.NewMNKX(3)
  subject := NewMNKAi(rules, b)
  subject.SetID(1)
  b.Mark(0,1)
  b.Mark(1,2)
  b.Mark(2,1)
  b.Mark(3,2)
  b.Mark(4,1)
  //5,0
  //6,0
  //7,0
  b.Mark(8,2)
  boardCopy := board.Board(b)
  depth := 0
  choices := map[int]int{}
  //When
  actualChoice := subject.optimumChoice(boardCopy, depth, choices)
  //Then
  expectedChoice := 6
  assertIntEqual(t, actualChoice, expectedChoice)
}

func TestOptimumChoiceWithTestBoardNumber3(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  b := board.NewMNKX(3)
  subject := NewMNKAi(rules, b)
  subject.SetID(2)
  b.Mark(0,1)
  b.Mark(1,2)
  b.Mark(2,1)
  b.Mark(3,2)
  b.Mark(4,1)
  b.Mark(5,1)
  //6,0
  //7,0
  b.Mark(8,2)
  boardCopy := board.Board(b)
  depth := 0
  choices := map[int]int{}
  //When
  actualChoice := subject.optimumChoice(boardCopy, depth, choices)
  //Then
  expectedChoice := 6
  assertIntEqual(t, actualChoice, expectedChoice)
}

func TestOptimumChoiceWhenThereIsAnEarlyChanceToToWinOrLoseOnRows(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  b := board.NewMNKX(3)
  subject := NewMNKAi(rules, b)
  subject.SetID(2)
  b.Mark(0,1)
  //1,0
  b.Mark(2,1)
  //3,0
  b.Mark(4,1)
  //4,0
  b.Mark(6,2)
  //7,0
  b.Mark(8,2)
  boardCopy := board.Board(b)
  depth := 0
  choices := map[int]int{}
  //When
  actualChoice := subject.optimumChoice(boardCopy, depth, choices)
  //Then
  expectedChoice := 7
  assertIntEqual(t, actualChoice, expectedChoice)
}

func TestOptimumChoiceWhenThereIsAnEarlyChanceToToWinOrLoseOn2(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  b := board.NewMNKX(3)
  subject := NewMNKAi(rules, b)
  subject.SetID(2)
  b.Mark(0,1)
  b.Mark(1,1)
  //2,0
  //3,0
  //4,0
  //4,0
  //5,0
  //6,0
  //7,0
  b.Mark(8,2)
  boardCopy := board.Board(b)
  depth := 0
  choices := map[int]int{}
  //When
  actualChoice := subject.optimumChoice(boardCopy, depth, choices)
  //Then
  expectedChoice := 2
  assertIntEqual(t, actualChoice, expectedChoice)
}

func assertIntEqual(t *testing.T, actual, expected int) {
  if actual != expected {
    t.Errorf("\n%d\nshould have been\n%d", actual, expected)
  }
}

func assertChoicesEqual(t *testing.T, actual, expected []map[int]int) {
  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("\n%d\nshould be\n%d", actual, expected)
  }
}

func assertChoiceEqual(t *testing.T, actual, expected map[int]int) {
  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("\n%d\nshould be\n%d", actual, expected)
  }
}

func assertSlicesEqual(t *testing.T, actual, expected []int) {
  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("\n%d\nshould be\n%d", actual, expected)
  }
}

func assertBoardsEqual(t *testing.T, actual, expected [][]int) {
  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("\n%d\nshould be\n%d", actual, expected)
  }
}
