package main

import "os/exec"

func processVideoForFastStart(filePath string) (string, error) {
	tempPath := filePath + ".processing"
	cmd := exec.Command("ffmpeg", "-i", filePath, "-c", "copy", "-movflags", "faststart", "-f", "mp4", tempPath)
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return tempPath, nil
}
