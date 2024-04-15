package inspector2

type T struct {
	L1 *L1
}

type L1 struct {
	L2 *L2
}

type L2 struct {
	L3 *L3
}

type L3 struct {
	S string
	I int64
	F float64
}
