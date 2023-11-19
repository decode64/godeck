package godeck

import (
	"encoding/binary"

	"github.com/sstallion/go-hid"
)

type StreamdeckOriginalV2 struct {
	device      *hid.Device
	rowCount    int
	columnCount int
}

func NewStreamdeckOriginalV2(device *hid.Device) Streamdeck {
	streamdeck := StreamdeckOriginalV2{device: device, rowCount: 3, columnCount: 5}
	return &streamdeck
}

// https://den.dev/blog/reverse-engineering-stream-deck/
func (streamdeck *StreamdeckOriginalV2) SetButtonImage(buttonIndex int, image []byte) error {
	imagePacketSize := 1024
	imagePacketHeaderSize := 8
	imagePacketPayloadSize := imagePacketSize - imagePacketHeaderSize
	imageBytesToSend := len(image)
	page := 0

	for imageBytesToSend > 0 {
		payloadSize := imagePacketPayloadSize
		if imageBytesToSend < imagePacketPayloadSize {
			payloadSize = imageBytesToSend
		}

		isFinalPacket := byte(0)
		if imageBytesToSend == payloadSize {
			isFinalPacket = 1
		}

		payloadSizeBytes := make([]byte, 2)
		binary.LittleEndian.PutUint16(payloadSizeBytes, uint16(payloadSize))

		pageBytes := make([]byte, 2)
		binary.LittleEndian.PutUint16(pageBytes, uint16(page))

		header := []byte{
			0x02,
			0x07,
			byte(buttonIndex),
			isFinalPacket,
			payloadSizeBytes[0],
			payloadSizeBytes[1],
			pageBytes[0],
			pageBytes[1]}

		packet := make([]byte, imagePacketSize)
		copy(packet, header)

		imageBytesOffset := page * imagePacketPayloadSize
		copy(packet[len(header):], image[imageBytesOffset:])

		if _, err := streamdeck.device.Write(packet); err != nil {
			return err
		}

		page += 1
		imageBytesToSend -= payloadSize
	}

	return nil
}

func (streamdeck *StreamdeckOriginalV2) SetBrightness(value int) error {
	packet := []byte{0x03, 0x08, byte(value)}

	if _, err := streamdeck.device.SendFeatureReport(packet); err != nil {
		return err
	}

	return nil
}

func (streamdeck *StreamdeckOriginalV2) Reset() error {
	packet := []byte{0x03, 0x02}

	if _, err := streamdeck.device.SendFeatureReport(packet); err != nil {
		return err
	}

	return nil
}

func (streamdeck *StreamdeckOriginalV2) ReadButtonState() []byte {
	state := make([]byte, streamdeck.GetButtonCount()+4)
	streamdeck.device.Read(state)

	return state[4:]
}

func (streamdeck *StreamdeckOriginalV2) GetButtonCount() int {
	return streamdeck.rowCount * streamdeck.columnCount
}

func (streamdeck *StreamdeckOriginalV2) GetRowCount() int {
	return streamdeck.rowCount
}

func (streamdeck *StreamdeckOriginalV2) GetColumnCount() int {
	return streamdeck.columnCount
}
