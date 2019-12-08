package eight_space_image_format

type LayerCounter struct {
	zeros, ones, twos int
}

func SpaceImageFormatPartOne(image string, width, height int) int {
	layers := []*LayerCounter{}

	for i := 0; i < len(image); i += width * height {
		layer := &LayerCounter{0, 0, 0}

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

func SpaceImageFormatPartTwo(image string, width, height int) [][]string {
	paint := map[string]string{
		"0": "_",
		"1": "*",
	}

	layer := make([][]string, height)
	for i := range layer {
		layer[i] = make([]string, width)
	}

	for i := 0; i < len(image); i += width * height {
		fullLayer := image[i : i+(width*height)]

		pixel := 0
		for j := 0; j < height; j += 1 {
			for k := 0; k < width; k += 1 {

				if string(fullLayer[pixel]) == "0" || string(fullLayer[pixel]) == "1" {
					if layer[j][k] == "" {
						layer[j][k] = paint[string(fullLayer[pixel])]
					}
				}

				pixel += 1
			}
		}
	}

	return layer
}
