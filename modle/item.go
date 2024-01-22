package modle

type Item struct {
	Name        string
	Count       int
	Description string
	State       int
	Instruction string //使用方法
}

var Items map[string][]Item

func Add(uuid string, i Item) error {
	Items[uuid] = append(Items[uuid], i)
	return nil
}
