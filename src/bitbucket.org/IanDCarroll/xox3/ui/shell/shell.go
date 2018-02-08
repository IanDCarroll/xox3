package shell

type ShellOut interface {
  Write(interface{})
}

type ShellIn interface {
  Read() interface{}
}
