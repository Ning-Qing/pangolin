package pangolin

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	r := NewRouter()
	r.Add("test", func(ctx *PangolinCtx) {
		fmt.Println("ok")
	})
	_, err := NewClient(":7001", ":8001", ":9001",r)
	if err != nil {
		fmt.Println(err)
	}
}
