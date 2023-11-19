package godeck

type Streamdeck interface {
	SetButtonImage(buttonIndex int, image []byte) error
	SetBrightness(value int) error
	Reset() error
	ReadButtonState() []byte
	GetButtonCount() int
	GetRowCount() int
	GetColumnCount() int
}
