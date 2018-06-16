package fontmap

import (
	log "github.com/sirupsen/logrus"
)

type Font struct {
	Name     string       `json:"name"`
	Metadata MetadataType `json:"metadata"`
	Charmap  CharmapType  `json:"charmap"`
}

type MetadataType struct {
	TallerCharacters map[string]int `json:"tallerCharacters"`
	AverageHeight    int            `json:"averageHeight"`
	AverageWidth     int            `json:"averageWidth"`
	MaxHeight        int            `json:"maxHeight"`
}

type CharmapType map[string]Letter
type Letter []Row
type Row []int

// The String representation of a letter uses emojis for quick debug rendering
func (letter Letter) String() string {
	var out string
	for _, row := range letter {
		for _, cell := range row {
			if cell == 1 {
				out = out + "⚫️"
			} else {
				out = out + "⚪️"
			}
		}
		out += "\n"
	}
	return out
}

// GenerateSpace will a letter given width and height.
func GenerateSpace(width int, height int, fill int) Letter {
	var space Letter
	for j := 0; j < height; j++ {
		var row Row
		for i := 0; i < width; i++ {
			row = append(row, fill)
		}

		space = append(space, row)
	}

	return space
}

// AddKerning will add trailing whitespace to the end of the letter
func AddKerning(letter Letter, amountOfKerning int) Letter {
	if amountOfKerning < 0 {
		log.Error("AddKerning does not support negative amounts yet")
		return letter
	}

	for rowIndex := range letter {
			for j := 0; j < amountOfKerning; j++ {
				letter[rowIndex] = append(letter[rowIndex], 0)
			}
	}

	return letter
}
