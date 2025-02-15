package main

import (
	"bytes"
	"image"
	"syscall/js"

	"github.com/disintegration/imaging"
)

func Resize(this js.Value, args []js.Value) any {
	// retrieve args
	w, h := args[0].Int(), args[1].Int()
	imgBytesLen := args[2].Get("length").Int()
	imgBytes := make([]byte, imgBytesLen)
	js.CopyBytesToGo(imgBytes, args[2])

	// process
	src, _, _ := image.Decode(bytes.NewReader(imgBytes))
	newImg := imaging.Resize(src, w, h, imaging.Lanczos)

	// output
	buffer := new(bytes.Buffer)
	imaging.Encode(buffer, newImg, imaging.JPEG)
	jsArray := js.Global().Get("Uint8Array").New(args[2].Get("length").Int())
	js.CopyBytesToJS(jsArray, buffer.Bytes())
	return jsArray
}

func main() {
	js.Global().Set("Resize", js.FuncOf(Resize))
	select {}
}
