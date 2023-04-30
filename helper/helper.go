package helper

/*
Pokeapi uses hectograms for weight and decimetres for height.

lbs = hectograms / 4.536
inches = decimetres * 3.937
*/

func calculateLbs(value float32) float32 {
	return value / 4.536
}

func calculateInches(value float32) float32 {
	return value * 3.937
}
