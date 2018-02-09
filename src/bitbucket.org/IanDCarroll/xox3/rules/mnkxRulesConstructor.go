package rules

func constructMNKXRules(size int) [][]int {
  waysToWin := make([][]int, 0)
  waysToWin = append(waysToWin, constructRows(size)...)
  waysToWin = append(waysToWin, constructCols(size)...)
  waysToWin = append(waysToWin, constructDiags(size)...)
  return waysToWin
}

func constructRows(size int) [][]int {
  prime := primeRowSlice(size)
  return constructLines(size, prime, 1)
}

func constructCols(size int) [][]int {
  prime := primeColSlice(size)
  return constructLines(size, prime, size)
}

func constructLines(size int, prime []int, inc int) [][]int {
  lines := make([][]int, size)
  for i := 0; i < size; i++ {
    lines[i] = make([]int,size)
    lines[i][0] = prime[i]
    for j := 1; j < size; j++ {
      lines[i][j] = lines[i][j-1] + inc
    }
  }
  return lines
}

func primeRowSlice(size int) []int {
  return primeSlice(size*size, size)
}

func primeColSlice(size int) []int {
  return primeSlice(size, 1)
}

func primeSlice(loopRange, increment int) []int {
  prime := make([]int, 0)
  for i := 0; i < loopRange; i = i+increment {
    prime = append(prime, i)
  }
  return prime
}

func constructDiags(size int) [][]int {
  diags := make([][]int, 2)
  diags[0] = constructDexter(size)
  diags[1] = constructSinister(size)
  return diags
}

func constructDexter(size int) []int {
  return constructDiag(size, 0, 1)
}

func constructSinister(size int) []int {
  return constructDiag(size, size-1, -1)
}

func constructDiag(size, firstValue, inc int) []int {
  diag := make([]int, size)
  diag[0] = firstValue
  for i := 1; i < size; i++ {
    diag[i] =  diag[i-1] + size + inc
  }
  return diag
}
