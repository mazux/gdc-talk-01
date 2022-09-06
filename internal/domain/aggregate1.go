package domain

type Aggregate1 struct {
	f1 string
	f2 int
}

func MutateAggregate1(f1 string, f2 int, domainBus chan<- interface{}) (*Aggregate1, error) {
	a := &Aggregate1{f1, f2}

	domainBus <- Event1{f1}
	return a, nil
}

type Event1 struct {
	f1 string
}
