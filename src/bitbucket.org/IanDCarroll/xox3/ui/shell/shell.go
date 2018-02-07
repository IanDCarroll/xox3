package shell

type Shell interface {
  Write(string)
  Read() int
}
