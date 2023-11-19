package main

import (
	"bufio"
	"log"
	"os"
	"time"

	"github.com/decode64/godeck"
	"github.com/sstallion/go-hid"
)

const (
	streamdeckVid = 0x0fd9
	streamdeckPid = 0x006d
)

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	size := stats.Size()
	image := make([]byte, size)

	reader := bufio.NewReader(file)
	_, err = reader.Read(image)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func newStreamdeck() (godeck.Streamdeck, error) {
	if err := hid.Init(); err != nil {
		return nil, err
	}
	defer hid.Exit()

	device, err := hid.OpenFirst(streamdeckVid, streamdeckPid)
	if err != nil {
		return nil, err
	}

	streamdeck := godeck.NewStreamdeckOriginalV2(device)
	streamdeck.Reset()
	return streamdeck, nil
}

func calculateNewDirection(state []byte) direction {
	if state[2] == 1 {
		return UP
	}

	if state[6] == 1 {
		return LEFT
	}

	if state[8] == 1 {
		return RIGHT
	}

	if state[12] == 1 {
		return DOWN
	}

	return 0
}

func main() {
	whiteSquareImage, err := readFile("./white-square.jpg")
	if err != nil {
		log.Fatal(err)
		return
	}

	blackSquareImage, err := readFile("./black-square.jpg")
	if err != nil {
		log.Fatal(err)
		return
	}

	streamdeck, err := newStreamdeck()
	if err != nil {
		log.Fatal(err)
		return
	}

	for i := 0; i < streamdeck.GetButtonCount(); i++ {
		streamdeck.SetButtonImage(i, blackSquareImage)
	}

	remove := func(i int, j int) {
		streamdeck.SetButtonImage(i*5+j, blackSquareImage)
	}

	add := func(i int, j int) {
		streamdeck.SetButtonImage(i*5+j, whiteSquareImage)
	}

	game := newGame()

	go func() {
		for {
			buttonState := streamdeck.ReadButtonState()
			newDirection := calculateNewDirection(buttonState)

			if newDirection != 0 {
				game.ChangeDirection(newDirection)
			}
		}
	}()

	for {
		game.Step(add, remove)
		time.Sleep(250 * time.Millisecond)
	}
}
