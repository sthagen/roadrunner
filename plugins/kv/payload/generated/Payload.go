// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package generated

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Payload struct {
	_tab flatbuffers.Table
}

func GetRootAsPayload(buf []byte, offset flatbuffers.UOffsetT) *Payload {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Payload{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Payload) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Payload) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Payload) Storage() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Payload) Items(obj *Item, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Payload) ItemsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func PayloadStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PayloadAddStorage(builder *flatbuffers.Builder, Storage flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Storage), 0)
}
func PayloadAddItems(builder *flatbuffers.Builder, Items flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(Items), 0)
}
func PayloadStartItemsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func PayloadEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
