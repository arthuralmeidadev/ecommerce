package orders

type Sub struct{}

func (s *Sub) TestSub() {
	println("SUB TEST")
}

type OrdersProvider struct {
	Sub Sub
}

func (p *OrdersProvider) Hello() {
	println("HELLO ORDERS CONTROLLER")
}
