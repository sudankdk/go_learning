package compositions

type engine struct{
	hp int
}

func (e engine) HP () int{
	return e.hp
}

type wheel struct{
	wheelDimension int
}


//struct inside another struct is called composition
type Car struct{
	CarName string
	Wheel wheel
	engine
}

func NewCar (carName string, w, e int) Car{
	return Car{
		CarName: carName,
		Wheel:wheel{
			wheelDimension: w,
	
		},
		engine: engine{
			hp: e,
		},
	}
}