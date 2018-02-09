# xox3 #
An Unbeatable Noughts and Crosses Game

### Dependencies ###

- Go 1.9.3

### Setup ###

- clone this repo into a directory within your [Go workspace](https://golang.org/doc/code.html#Workspaces).
- from project root `go install ./...`
- `export PATH=$HOME/xox3/bin:$PATH`

### Run ###

- `xox3`

### Test ###

- from project root `go test ./...`

### Common Snags ###

if go can't find any of the projects packages, you might need to [set the GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)

### Domain Naming ###

You will see in the codebase certain names that will contain `MNK` `MNKX` or `MNK3`.
These name are referring to the mathematical abstract of the kind of game noughts and crosses belongs to. Noughts and crosses is an [m,n,k-game](https://en.wikipedia.org/wiki/M,n,k-game), specifically a 3,3,3-game.
So if you were to see `MNK` in a name, it means the implementation is valid for any arbitrary m,n,k-game.

`MNKX` means the implementation will work for any arbitrary m,n,k-game where all three variable are equal.

`MNK3` means the implementation will work for only a 3X3 game of classic noughts and crosses where 3-in-a-row wins.

These identifiers are meant to give information on what would need to be replaced to extend the game's versatility.
