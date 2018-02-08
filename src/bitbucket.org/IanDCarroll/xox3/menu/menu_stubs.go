package menu

import (
  "bitbucket.org/IanDCarroll/xox3/ui/shell"
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/ui/selector"
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/ui/display/rec"
)

type shellStub struct { shellOutput string }
func (s shellStub) Read() interface{} { return "2" }
func (s shellStub) Write(message interface{}) { s.shellOutput = message.(string) }

func buildEnglishMenuRec() rec.MenuRec {
  return rec.NewEnglishMenu()
}

func buildMenuUI(shellOut shell.ShellOut, shellIn shell.ShellIn) ui.Ui {
  display := buildMenuDisplay(shellOut)
  selector := buildMenuSelector(shellIn)
  return ui.New(display, selector)
}

func buildMenuDisplay(shell shell.ShellOut) display.Display {
  return display.NewMenu(shell)
}

func buildMenuSelector(shell shell.ShellIn) selector.Selector {
  return selector.NewMenu(shell)
}
