package stw

type ctx interface {
	StopTheWorld()
	StartTheWorld()
}

func NewSTWCtx() ctx {
	return newSTWCtx()
}
