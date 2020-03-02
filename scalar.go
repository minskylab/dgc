package dgc

import "encoding/binary"

type scalarV1 struct {
	t          [8]byte
	addr       *scalarV1 // Virtual [in hardware must be a real address of memory]
	parentAddr *scalarV1 // Virtual
	size       [8]byte
	data  	   [8]byte
}

type scalarv2 [40]byte


func uint64ToFrac(number uint64) [8]byte {
	bl := make([]byte, 8)
	binary.BigEndian.PutUint64(bl, number)
	bx := [8]byte{}
	copy(bx[:], bl)
	return bx
}

func newStringv1(value string) *scalarV1 {
	l := uint64ToFrac(uint64(len(value)))

	var head *scalarV1
	for _ ,v := range value {
		head = &scalarV1{
			t:          [8]byte{0,0,0,0,0,0,0,0},
			addr:       nil, // because in this example (virtual)
							 // the address is the address of the
							 // instance of the struct scalarV1
			parentAddr: head,
			size:       l,
			data:       uint64ToFrac(uint64(v)),
		}
	}
	return head
}