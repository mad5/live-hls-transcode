package main

import (
	"io/ioutil"
	"log"
	"os"
)

// Stored Information about a Stream
type StreamInfo struct {
	tempDir string
	handle  TranscoderHandle
}

type StreamStatusManager struct {
	tempDir    string
	streamInfo map[string]StreamInfo
	transcoder Transcoder
}

func NewStreamStatusManager(tempDir string) StreamStatusManager {
	err := os.MkdirAll(tempDir, 0770)
	if err != nil {
		log.Fatal("Cannot create Temp-Dir %s", tempDir)
	}

	return StreamStatusManager{
		tempDir,
		make(map[string]StreamInfo),
		NewTranscoder(),
	}
}

type StreamStatus int

const (
	NoStream = iota
	StreamTranscodingFailed
	StreamInPreparation
	StreamReady
)

func (manager StreamStatusManager) GetStreamStatus(calculatedPath string) StreamStatus {
	info, hasInfo := manager.streamInfo[calculatedPath]

	if hasInfo {
		if info.handle.IsReady() {
			return StreamReady
		} else if info.handle.IsRunning() {
			return StreamInPreparation
		} else {
			return StreamTranscodingFailed
		}
	} else {
		return NoStream
	}
}

func (manager StreamStatusManager) StartStream(calculatedPath string) {
	_, hasInfo := manager.streamInfo[calculatedPath]
	if hasInfo {
		log.Printf("%s: Stream already active", calculatedPath)
		return
	}

	tempDir, err := ioutil.TempDir(manager.tempDir, "hls")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s: Starting Stream-Transcoder into %s", calculatedPath, tempDir)
	handle := manager.transcoder.StartTranscoder(calculatedPath, tempDir)

	manager.streamInfo[calculatedPath] = StreamInfo{
		tempDir,
		handle,
	}
}

func (manager StreamStatusManager) StopStream(calculatedPath string) {
	info, hasInfo := manager.streamInfo[calculatedPath]
	if hasInfo {
		if info.handle.IsFinished() {
			log.Printf("%s: Stream-Transcoder already finished, ignoring Stop-Command", calculatedPath)
		} else {
			log.Printf("%s: Stopping unfinished Stream-Transcoder", calculatedPath)
			info.handle.Stop()

			delete(manager.streamInfo, calculatedPath)
		}
	}
}
