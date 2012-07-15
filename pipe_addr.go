package gotesthelpers

type pipeAddr int

func (pipeAddr) Network() string {
	return "pipe"
}

func (pipeAddr) String() string {
	return "pipe"
}
