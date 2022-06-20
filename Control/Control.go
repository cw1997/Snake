package Control

import (
	"log"

	"github.com/mattn/go-tty"
)

type TTYHolder struct {
	TtyHolder *tty.TTY
	err error
	Key rune
}

func (ttyholder *TTYHolder)Start(){
	ttyholder.TtyHolder, ttyholder.err = tty.Open()
	if ttyholder.err != nil{
		log .Fatal(ttyholder.err)
	}
}

func (ttyholder *TTYHolder)GetKey(Key chan string){
	ttyholder.Key ,ttyholder.err = ttyholder.TtyHolder.ReadRune()
	if ttyholder.err != nil{
		log.Fatal(ttyholder.err)
	}
	Key <- string(ttyholder.Key)
}