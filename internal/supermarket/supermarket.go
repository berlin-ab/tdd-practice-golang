package supermarket

type register struct {
	items []string
}

func (r *register) Scan(item string) {
	r.items = append(r.items, item)
}

func (r *register) Total() int {
	total := 0

	for _, item := range(r.items) {
		switch item {
		case "apple":
			total += 3
		case "bananas":
			total += 1
		case "pie":
			total += 10
		case "orange":
			total += 2
		}
	}

	return total
}

func NewRegister() register {
	return register{}
}