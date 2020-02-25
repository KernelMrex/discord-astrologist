package zodiac

import (
	"encoding/json"
	"errors"
	"os"
)

type Associate struct {
	signs map[string]Sign `json:"signs"`
}

type Sign struct {
	Pseudo []string `json:"pseudo"`
	Emoji  string   `json:"emoji"`
}

func LoadAssociateFromJson(filepath string) (*Associate, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return &Associate{}, errors.New("couldn't open file '" + filepath + "': " + err.Error())
	}
	defer file.Close()

	signs := make(map[string]Sign)
	if err := json.NewDecoder(file).Decode(&signs); err != nil {
		return &Associate{}, errors.New("error while parsing file '" + filepath + "': " + err.Error())
	}

	return &Associate{signs:signs}, nil
}
