package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type AspectRatio struct {
	Streams []struct {
		Width  int `json:"width,omitempty"`
		Height int `json:"height,omitempty"`
	} `json:"streams"`
}

func getVideoAspectRatio(filePath string) (string, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-print_format", "json", "-show_streams", filePath)
	var b bytes.Buffer
	cmd.Stdout = &b
	if err := cmd.Run(); err != nil {
		fmt.Printf("ffprobe error: %v\n", err)
		return "", err
	}

	var Output AspectRatio
	if err := json.Unmarshal(b.Bytes(), &Output); err != nil {
		return "", err
	}

	width := float64(Output.Streams[0].Width)
	height := float64(Output.Streams[0].Height)
	ratio := width / height
	if ratio >= (16.0/9.0)-0.1 && ratio <= (16.0/9.0)+0.1 {
		return "16:9", nil
	} else if ratio >= (9.0/16.0)-0.1 && ratio <= (9.0/16.0)+0.1 {
		return "9:16", nil
	} else {
		return "other", nil
	}
}
