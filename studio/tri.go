package studio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//"unsafe"
)

// triangles
type Tri struct {
	TriCmd int16 // minus:TRIANGLE_FUN, 0:END, plus:TRIANGLE_STRIP
	Points [][4]int16
}

func (t *Tri) Read(r *bytes.Reader) int {

	// read tricmd
	if err := binary.Read(r, binary.LittleEndian, &t.TriCmd); err != nil {
		fmt.Print(err)
		return int(0)
	}

	var num int
	if num = int(t.TriCmd); num < 0 {
		num = -num
	} else if num == 0 {
		return int(0)
	}

	// read points
	t.Points = make([][4]int16, num)
	if err := binary.Read(r, binary.LittleEndian, t.Points); err != nil {
		fmt.Print(err)
		return int(0)
	}

	return int(t.TriCmd)
}

func NewTris(buf []byte) []Tri {
	var tris []Tri
	r := bytes.NewReader(buf)

	for {
		t := new(Tri)

		if cmd := t.Read(r); cmd == 0 {
			break
		}

		tris = append(tris, *t)
	}

	return tris
}
