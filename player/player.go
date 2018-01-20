package player

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func Play(data []byte) error {
	dt := bytes.NewReader(data)
	f := ioutil.NopCloser(dt)

	d, err := mp3.NewDecoder(f)
	if err != nil {
		log.Println("player: " + err.Error())
		return err
	}
	defer d.Close()

	p, err := oto.NewPlayer(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}
