package db

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z Comment) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "i"
	o = append(o, 0x83, 0xa1, 0x69)
	o = msgp.AppendInt64(o, z.Id)
	// string "u"
	o = append(o, 0xa1, 0x75)
	o = msgp.AppendInt64(o, z.UserId)
	// string "c"
	o = append(o, 0xa1, 0x63)
	o = msgp.AppendString(o, z.Content)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Comment) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "i":
			z.Id, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Id")
				return
			}
		case "u":
			z.UserId, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "UserId")
				return
			}
		case "c":
			z.Content, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Content")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Comment) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.StringPrefixSize + len(z.Content)
	return
}
