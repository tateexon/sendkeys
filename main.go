package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/micmonay/keybd_event"
)

type SendKeys struct {
	KeyBindings   map[rune]KeyMapping
	StartDelay    time.Duration
	InterKeyDelay time.Duration
	kb            keybd_event.KeyBonding
}

type KeyMapping struct {
	KeyCode int
	Shift   bool
}

func (sk *SendKeys) Init() error {
	if sk.KeyBindings == nil {
		sk.KeyBindings = sk.GetWindowsKeyBindings()
	}
	var err error
	sk.kb, err = keybd_event.NewKeyBonding()
	if err != nil {
		return err
	}
	return nil
}

func (sk SendKeys) GetWindowsKeyBindings() map[rune]KeyMapping {
	return map[rune]KeyMapping{
		'a':  {keybd_event.VK_A, false},
		'b':  {keybd_event.VK_B, false},
		'c':  {keybd_event.VK_C, false},
		'd':  {keybd_event.VK_D, false},
		'e':  {keybd_event.VK_E, false},
		'f':  {keybd_event.VK_F, false},
		'g':  {keybd_event.VK_G, false},
		'h':  {keybd_event.VK_H, false},
		'i':  {keybd_event.VK_I, false},
		'j':  {keybd_event.VK_J, false},
		'k':  {keybd_event.VK_K, false},
		'l':  {keybd_event.VK_L, false},
		'm':  {keybd_event.VK_M, false},
		'n':  {keybd_event.VK_N, false},
		'o':  {keybd_event.VK_O, false},
		'p':  {keybd_event.VK_P, false},
		'q':  {keybd_event.VK_Q, false},
		'r':  {keybd_event.VK_R, false},
		's':  {keybd_event.VK_S, false},
		't':  {keybd_event.VK_T, false},
		'u':  {keybd_event.VK_U, false},
		'v':  {keybd_event.VK_V, false},
		'w':  {keybd_event.VK_W, false},
		'x':  {keybd_event.VK_X, false},
		'y':  {keybd_event.VK_Y, false},
		'z':  {keybd_event.VK_Z, false},
		'A':  {keybd_event.VK_A, true},
		'B':  {keybd_event.VK_B, true},
		'C':  {keybd_event.VK_C, true},
		'D':  {keybd_event.VK_D, true},
		'E':  {keybd_event.VK_E, true},
		'F':  {keybd_event.VK_F, true},
		'G':  {keybd_event.VK_G, true},
		'H':  {keybd_event.VK_H, true},
		'I':  {keybd_event.VK_I, true},
		'J':  {keybd_event.VK_J, true},
		'K':  {keybd_event.VK_K, true},
		'L':  {keybd_event.VK_L, true},
		'M':  {keybd_event.VK_M, true},
		'N':  {keybd_event.VK_N, true},
		'O':  {keybd_event.VK_O, true},
		'P':  {keybd_event.VK_P, true},
		'Q':  {keybd_event.VK_Q, true},
		'R':  {keybd_event.VK_R, true},
		'S':  {keybd_event.VK_S, true},
		'T':  {keybd_event.VK_T, true},
		'U':  {keybd_event.VK_U, true},
		'V':  {keybd_event.VK_V, true},
		'W':  {keybd_event.VK_W, true},
		'X':  {keybd_event.VK_X, true},
		'Y':  {keybd_event.VK_Y, true},
		'Z':  {keybd_event.VK_Z, true},
		'0':  {keybd_event.VK_0, false},
		'1':  {keybd_event.VK_1, false},
		'2':  {keybd_event.VK_2, false},
		'3':  {keybd_event.VK_3, false},
		'4':  {keybd_event.VK_4, false},
		'5':  {keybd_event.VK_5, false},
		'6':  {keybd_event.VK_6, false},
		'7':  {keybd_event.VK_7, false},
		'8':  {keybd_event.VK_8, false},
		'9':  {keybd_event.VK_9, false},
		' ':  {keybd_event.VK_SPACE, false},
		'.':  {keybd_event.VK_DOT, false},
		',':  {keybd_event.VK_COMMA, false},
		'!':  {keybd_event.VK_1, true},
		'?':  {keybd_event.VK_SLASH, true},
		'@':  {keybd_event.VK_2, true},
		'#':  {keybd_event.VK_3, true},
		'$':  {keybd_event.VK_4, true},
		'%':  {keybd_event.VK_5, true},
		'^':  {keybd_event.VK_6, true},
		'&':  {keybd_event.VK_7, true},
		'*':  {keybd_event.VK_8, true},
		'(':  {keybd_event.VK_9, true},
		')':  {keybd_event.VK_0, true},
		'-':  {keybd_event.VK_MINUS, false},
		'_':  {keybd_event.VK_MINUS, true},
		'=':  {keybd_event.VK_EQUAL, false},
		'+':  {keybd_event.VK_EQUAL, true},
		'[':  {keybd_event.VK_LEFTBRACE, false},
		']':  {keybd_event.VK_RIGHTBRACE, false},
		'{':  {keybd_event.VK_LEFTBRACE, true},
		'}':  {keybd_event.VK_RIGHTBRACE, true},
		'\\': {keybd_event.VK_BACKSLASH, false},
		'|':  {keybd_event.VK_BACKSLASH, true},
		';':  {keybd_event.VK_SEMICOLON, false},
		':':  {keybd_event.VK_SEMICOLON, true},
		'\'': {keybd_event.VK_APOSTROPHE, false},
		'"':  {keybd_event.VK_APOSTROPHE, true},
		'`':  {keybd_event.VK_GRAVE, false},
		'~':  {keybd_event.VK_GRAVE, true},
		'/':  {keybd_event.VK_SLASH, false},
		'<':  {keybd_event.VK_COMMA, true},
		'>':  {keybd_event.VK_DOT, true},
		'\n': {keybd_event.VK_ENTER, false},
	}
}

func (sk SendKeys) TypeString(s string) error {
	for _, char := range s {
		if keyMapping, ok := sk.KeyBindings[char]; ok {
			if keyMapping.Shift {
				sk.kb.HasSHIFT(true)
			} else {
				sk.kb.HasSHIFT(false)
			}
			sk.kb.SetKeys(keyMapping.KeyCode)
			err := sk.kb.Launching()
			if err != nil {
				fmt.Println("Error typing key:", err)
				return err
			}
			time.Sleep(sk.InterKeyDelay) // Adjust delay as needed
		} else {
			return fmt.Errorf("unsupported character: %s", string(char))
		}
	}
	return nil
}

func main() {
	// Define command-line arguments
	t := flag.Int("t", 5, "Number of seconds to wait before starting")
	i := flag.Int("i", 85, "Interval in milliseconds between each key press")
	s := flag.String("s", "", "String to type")

	// Parse the command-line arguments
	flag.Parse()

	if *s == "" {
		fmt.Println("No string to type")
		return
	}

	sk := &SendKeys{
		StartDelay:    time.Duration(*t) * time.Second,
		InterKeyDelay: time.Duration(*i) * time.Millisecond,
	}
	err := sk.Init()
	if err != nil {
		fmt.Println("Error initializing keybd_event:", err)
		return
	}

	// Wait a few seconds to switch to the target application
	fmt.Println("You have 5 seconds to switch to the target application...")
	time.Sleep(sk.StartDelay)

	// type the text
	err = sk.TypeString(*s)
	if err != nil {
		fmt.Println("Error typing string:", err)
		return
	}

	fmt.Println("Key presses sent successfully.")
}
