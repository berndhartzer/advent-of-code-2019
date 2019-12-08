package eight_space_image_format

type Layer struct {
	zeros, ones, twos int
}

func SpaceImageFormatPartOne(image string, width, height int) int {
	layers := []*Layer{}

	for i := 0; i < len(image); i += width * height {
		layer := &Layer{0, 0, 0}

		for _, pixel := range image[i : i+(width*height)] {
			switch string(pixel) {
			case "0":
				layer.zeros += 1
			case "1":
				layer.ones += 1
			case "2":
				layer.twos += 1
			}
		}

		layers = append(layers, layer)
	}

	fewestZeros := 0
	output := 0
	for _, layer := range layers {
		if layer.zeros < fewestZeros || fewestZeros == 0 {
			fewestZeros = layer.zeros
			output = layer.ones * layer.twos
		}
	}

	return output
}
