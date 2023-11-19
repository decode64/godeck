# GoDeck

GoDeck provides a simple and convenient way to interact with Elgato Stream Deck devices in Go applications. It abstracts the communication with the device, allowing you to focus on developing applications that leverage the Stream Deck's capabilities.

## Supported devices

Currently, only the original StreamDeck V2 hardware is supported.


## Getting Started

To use this library in your Go project, you can simply use the go get command:


```bash
go get github.com/decode64/godeck
```

## Usage

To use the GoDeck Library, you need to create an `Streamdeck` instance. The library provides a helper function called `StreamdeckOriginalV2` to instantiate a `StreamDeck` from a `hid.Device`:

```go
import (
	"github.com/decode64/godeck"
	"github.com/sstallion/go-hid"
)

device, err := hid.OpenFirst(0x0fd9, 0x006d)
if err != nil {
  log.Fatal(err)
}
defer device.Close()

streamdeck := godeck.NewStreamdeckOriginalV2(device)
```

## Streamdeck Interface

The Streamdeck interface defines the following methods:

### SetButtonImage

```go
SetButtonImage(buttonIndex int, image []byte) error
```

This method sets the image for a specific button on the Streamdeck. `buttonIndex` represents the index of the button (zero-based), and `image` is a byte array containing the image data. The expected image format is a 72x72 JPEG.

### SetBrightness

```go
SetBrightness(value int) error
```

SetBrightness allows you to adjust the brightness of the Streamdeck display. The value parameter should be an integer between 0 and 255, representing the brightness level.

### Reset

```go
Reset() error
```

The Reset method resets the Streamdeck device, restoring it to its default state.

### ReadButtonState

```go
ReadButtonState() []byte
```

ReadButtonState retrieves the current state of all buttons on the Streamdeck as a byte array. Each byte represents the state of a button, allowing you to determine whether a button is pressed or released.

### GetButtonCount

```go
GetButtonCount() int
```

GetButtonCount returns the total number of buttons on the Streamdeck.

### GetRowCount

```go
GetRowCount() int
```

GetRowCount returns the number of rows on the Streamdeck.

### GetColumnCount

```go
GetColumnCount() int
```

GetColumnCount returns the number of columns on the Streamdeck.

## Contribution

Contributions to this library are welcome. If you encounter any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

## License

This library is licensed under the MIT License - see the LICENSE file for details.
