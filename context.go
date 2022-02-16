package pangolin

import (
	"context"
	"encoding/gob"
	"io"
)

type PangolinCtx struct{
	ctx context.Context
	header Header
	body Body
	w io.Writer
}

func NewPangolinCtx()*PangolinCtx{
	return nil
}

func (ctx *PangolinCtx)Write(m Message) error{
	encoder :=gob.NewEncoder(ctx.w)
	return encoder.Encode(m)
}

func (ctx *PangolinCtx)IsNewConnection()bool {
	return ctx.header.IsNewConnection()
}