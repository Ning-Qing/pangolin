package pangolin

import (
	"fmt"
	"testing"
)


func TestServer(t *testing.T){
	s,err := NewServer(":8001")
	if err != nil {
		fmt.Println(err)
	}
	s.Run()
}