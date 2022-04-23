package models

import (
	"encoding/json"
)

type Data struct {
	Programs []Programs `json:"programs"`
}

type Programs struct {
	ProgramName string `json:"programName"`
	Nodes       []Node `json:"nodes,omitempty"`
	Uid         string `json:"uid,omitempty"`
}

type Node struct {
	Id       int                `json:"id"`
	Name     string             `json:"name"`
	Data     map[string]string  `json:"data"`
	Class    string             `json:"class"`
	Html     string             `json:"html"`
	Typenode bool               `json:"typenode"`
	Inputs   map[string]Inputs  `json:"inputs"`
	Outputs  map[string]Outputs `json:"outputs"`
	PosX     float32            `json:"pos_x"`
	PosY     float32            `json:"pos_y"`
}

type Inputs struct {
	Connections []Connection2 `json:"connections"`
}

type Outputs struct {
	Connections []Connection1 `json:"connections"`
}

type Connection1 struct {
	Node   string `json:"node,omitempty"`
	Output string `json:"output,omitempty"`
}

type Connection2 struct {
	Order int    `json:"order"`
	Node  string `json:"node,omitempty"`
	Input string `json:"input,omitempty"`
}

//UnmarshalJSON to customize the generated json
func (i *Node) UnmarshalJSON(data []byte) error {
	type tmp Node
	if err := json.Unmarshal(data, (*tmp)(i)); err != nil {
		return err
	}

	if i.Data == nil {
		i.Data = map[string]string{}
	}

	if i.Outputs == nil {
		if i.Name == "if" ||
			i.Name == "if-condition" ||
			i.Name == "if-body" ||
			i.Name == "else-body" ||
			i.Name == "for" ||
			i.Name == "for-range" ||
			i.Name == "for-body" ||
			i.Name == "assign" ||
			i.Name == "number" ||
			i.Name == "add" ||
			i.Name == "sub" ||
			i.Name == "mul" ||
			i.Name == "div" ||
			i.Name == "print" {
			tempOutputs1 := make(map[string]Outputs, 1)
			tempOutputs1["output_1"] = Outputs{Connections: make([]Connection1, 0)}
			i.Outputs = tempOutputs1
		} else {
			temp := make(map[string]Outputs, 0)
			i.Outputs = temp
		}
	}

	if i.Inputs == nil {
		tempInputs1 := make(map[string]Inputs, 1)
		tempInputs1["input_1"] = Inputs{Connections: make([]Connection2, 0)}

		tempInputs2 := make(map[string]Inputs, 3)
		tempInputs2["input_1"] = Inputs{Connections: make([]Connection2, 0)}
		tempInputs2["input_2"] = Inputs{Connections: make([]Connection2, 0)}

		tempInputs3 := make(map[string]Inputs, 3)
		tempInputs3["input_1"] = Inputs{Connections: make([]Connection2, 0)}
		tempInputs3["input_2"] = Inputs{Connections: make([]Connection2, 0)}
		tempInputs3["input_3"] = Inputs{Connections: make([]Connection2, 0)}

		if i.Name == "if-condition" ||
			i.Name == "number" ||
			i.Name == "print" {
			temp := make(map[string]Inputs, 0)
			i.Inputs = temp
		}

		if i.Name == "root" ||
			i.Name == "if-body" ||
			i.Name == "else-body" ||
			i.Name == "for-body" ||
			i.Name == "assign" {
			i.Inputs = tempInputs1
		}

		if i.Name == "for" ||
			i.Name == "range" ||
			i.Name == "add" ||
			i.Name == "sub" ||
			i.Name == "mul" ||
			i.Name == "div" {
			i.Inputs = tempInputs2
		}

		if i.Name == "if" {
			i.Inputs = tempInputs3
		}

	} else {

		if i.Name == "if" {
			if _, ok := i.Inputs["input_1"]; !ok {
				i.Inputs["input_1"] = Inputs{Connections: make([]Connection2, 0)}
			}

			if _, ok := i.Inputs["input_2"]; !ok {
				i.Inputs["input_2"] = Inputs{Connections: make([]Connection2, 0)}
			}

			if _, ok := i.Inputs["input_3"]; !ok {
				i.Inputs["input_3"] = Inputs{Connections: make([]Connection2, 0)}
			}
		}

		if i.Name == "for" ||
			i.Name == "range" ||
			i.Name == "add" ||
			i.Name == "sub" ||
			i.Name == "mul" ||
			i.Name == "div" {
			if _, ok := i.Inputs["input_1"]; !ok {
				i.Inputs["input_1"] = Inputs{Connections: make([]Connection2, 0)}
			}

			if _, ok := i.Inputs["input_2"]; !ok {
				i.Inputs["input_2"] = Inputs{Connections: make([]Connection2, 0)}
			}
		}

	}

	return nil
}

//MarshalJSON to customize the json to save.
func (i *Node) MarshalJSON() ([]byte, error) {
	type temp Node

	if i.Name == "root" ||
		i.Name == "if-body" ||
		i.Name == "else-body" ||
		i.Name == "for-body" {
		connections := i.Inputs["input_1"].Connections
		for index := range connections {
			connections[index].Order = index + 1
		}
	}

	return json.Marshal((*temp)(i))
}
