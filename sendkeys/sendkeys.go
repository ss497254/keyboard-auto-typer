package sendkeys

import (
	"errors"
	"time"

	kbd "github.com/micmonay/keybd_event"
)

// KBWrap is a wrapper for the keybd_event library for convenience
type KBWrap struct {
	d        kbd.KeyBonding
	delay    time.Duration
	errors   []error
	stubborn bool
	noisy    bool
	nodelay  bool
}

// NewKBWrapWithOptions creates a new keyboard wrapper with the given options.
// As of writing, those options include: Stubborn Noisy NoDely.
// The defaults are all false.
func NewKBWrapWithOptions(delay time.Duration, opts ...KBOpt) (kbw *KBWrap, err error) {
	kbw = &KBWrap{errors: []error{},
		delay:    delay,
		stubborn: false,
		noisy:    false,
		nodelay:  false,
	}

	kbw.d, err = kbd.NewKeyBonding()
	if err != nil {
		return nil, err
	}
	kbw.processOptions(opts...)
	kbw.linuxDelay()
	return
}

func (kb *KBWrap) down() {
	if !kb.check() {
		return
	}
	kb.handle(kb.d.Press())
}
func (kb *KBWrap) up() {
	if !kb.check() {
		return
	}
	kb.handle(kb.d.Release())
}

func (kb *KBWrap) press() {
	kb.down()
	time.Sleep(kb.delay)
	kb.up()
	time.Sleep(kb.delay)
}

func (kb *KBWrap) set(keys ...int) {
	kb.d.SetKeys(keys...)
}

func (kb *KBWrap) clr() {
	kb.d.Clear()
}

func (kb *KBWrap) only(k int) {
	kb.clr()
	kb.set(k)
	kb.press()
	kb.clr()
}

func (kb *KBWrap) Escape() {
	kb.only(kbd.VK_ESC)
}

func (kb *KBWrap) Tab() {
	kb.only(kbd.VK_TAB)
}

func (kb *KBWrap) Enter() {
	kb.only(kbd.VK_ENTER)
}

func (kb *KBWrap) BackSpace() {
	kb.only(kbd.VK_ENTER)
}

func (kb *KBWrap) Type(s string) error {
	keys := kb.strToKeys(s)
	for _, key := range keys {
		if !kb.check() {
			return errors.New(compoundErr(kb.errors))
		}
		if key < 0 {
			kb.d.HasSHIFT(true)
			key = abs(key)
		}

		kb.set(key)
		kb.press()
		kb.clr()
	}
	return nil
}

func SendStrings(text []string, delay time.Duration) error {
	k, err := NewKBWrapWithOptions(delay, Noisy)
	if err != nil {
		return err
	}

	for idx, line := range text {
		k.Type(line)

		if idx != len(text)-1 {
			k.Enter()
			time.Sleep(delay)
		}
	}

	return nil
}
