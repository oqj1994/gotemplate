package modle

import "github.com/google/uuid"

type Box struct {
	UUID  uuid.UUID
	Title string
	State int
	Count int
}

var Boxs = []Box{
	Box{
		UUID:  uuid.New(),
		Title: "闲置的玩具",
		Count: 3,
		State: 0,
	},
	Box{
		UUID:  uuid.New(),
		Title: "闲置的书",
		Count: 12,
		State: 0,
	},
	Box{
		UUID:  uuid.New(),
		Count: 22,
		Title: "闲置的人",
		State: 2,
	},
}
