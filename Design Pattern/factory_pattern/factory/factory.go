package factory

// TEST for test
const TEST = "TEST"

type (
	//Car for car
	Car struct{}

	//Train for train
	Train struct{}

	//Engine for engine
	Engine interface {
		Accelerate(name string) (kecepatan int)
	}
)

func (c Car) Accelerate(name string) int {
	return 100
}

func (t Train) Accelerate(str string) int {
	return 200
}

func GetEngine(engine string) Engine {
	switch engine {
	case "car":
		return Car{}
	case "train":
		return Train{}
	default:
		return nil
	}

}
