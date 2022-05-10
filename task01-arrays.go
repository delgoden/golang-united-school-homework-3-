package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32
	for _, num := range input {
		sum += num
	}
	result = sum / float32(len(input))
	return
}
