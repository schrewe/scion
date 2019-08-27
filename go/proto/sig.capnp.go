// Code generated by capnpc-go. DO NOT EDIT.

package proto

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type SIGCtrl struct{ capnp.Struct }
type SIGCtrl_Which uint16

const (
	SIGCtrl_Which_unset   SIGCtrl_Which = 0
	SIGCtrl_Which_pollReq SIGCtrl_Which = 1
	SIGCtrl_Which_pollRep SIGCtrl_Which = 2
)

func (w SIGCtrl_Which) String() string {
	const s = "unsetpollReqpollRep"
	switch w {
	case SIGCtrl_Which_unset:
		return s[0:5]
	case SIGCtrl_Which_pollReq:
		return s[5:12]
	case SIGCtrl_Which_pollRep:
		return s[12:19]

	}
	return "SIGCtrl_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// SIGCtrl_TypeID is the unique identifier for the type SIGCtrl.
const SIGCtrl_TypeID = 0xe15e242973323d08

func NewSIGCtrl(s *capnp.Segment) (SIGCtrl, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return SIGCtrl{st}, err
}

func NewRootSIGCtrl(s *capnp.Segment) (SIGCtrl, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return SIGCtrl{st}, err
}

func ReadRootSIGCtrl(msg *capnp.Message) (SIGCtrl, error) {
	root, err := msg.RootPtr()
	return SIGCtrl{root.Struct()}, err
}

func (s SIGCtrl) String() string {
	str, _ := text.Marshal(0xe15e242973323d08, s.Struct)
	return str
}

func (s SIGCtrl) Which() SIGCtrl_Which {
	return SIGCtrl_Which(s.Struct.Uint16(8))
}
func (s SIGCtrl) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s SIGCtrl) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s SIGCtrl) SetUnset() {
	s.Struct.SetUint16(8, 0)

}

func (s SIGCtrl) PollReq() (SIGPoll, error) {
	if s.Struct.Uint16(8) != 1 {
		panic("Which() != pollReq")
	}
	p, err := s.Struct.Ptr(0)
	return SIGPoll{Struct: p.Struct()}, err
}

func (s SIGCtrl) HasPollReq() bool {
	if s.Struct.Uint16(8) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SIGCtrl) SetPollReq(v SIGPoll) error {
	s.Struct.SetUint16(8, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPollReq sets the pollReq field to a newly
// allocated SIGPoll struct, preferring placement in s's segment.
func (s SIGCtrl) NewPollReq() (SIGPoll, error) {
	s.Struct.SetUint16(8, 1)
	ss, err := NewSIGPoll(s.Struct.Segment())
	if err != nil {
		return SIGPoll{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s SIGCtrl) PollRep() (SIGPoll, error) {
	if s.Struct.Uint16(8) != 2 {
		panic("Which() != pollRep")
	}
	p, err := s.Struct.Ptr(0)
	return SIGPoll{Struct: p.Struct()}, err
}

func (s SIGCtrl) HasPollRep() bool {
	if s.Struct.Uint16(8) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SIGCtrl) SetPollRep(v SIGPoll) error {
	s.Struct.SetUint16(8, 2)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPollRep sets the pollRep field to a newly
// allocated SIGPoll struct, preferring placement in s's segment.
func (s SIGCtrl) NewPollRep() (SIGPoll, error) {
	s.Struct.SetUint16(8, 2)
	ss, err := NewSIGPoll(s.Struct.Segment())
	if err != nil {
		return SIGPoll{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// SIGCtrl_List is a list of SIGCtrl.
type SIGCtrl_List struct{ capnp.List }

// NewSIGCtrl creates a new list of SIGCtrl.
func NewSIGCtrl_List(s *capnp.Segment, sz int32) (SIGCtrl_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return SIGCtrl_List{l}, err
}

func (s SIGCtrl_List) At(i int) SIGCtrl { return SIGCtrl{s.List.Struct(i)} }

func (s SIGCtrl_List) Set(i int, v SIGCtrl) error { return s.List.SetStruct(i, v.Struct) }

func (s SIGCtrl_List) String() string {
	str, _ := text.MarshalList(0xe15e242973323d08, s.List)
	return str
}

// SIGCtrl_Promise is a wrapper for a SIGCtrl promised by a client call.
type SIGCtrl_Promise struct{ *capnp.Pipeline }

func (p SIGCtrl_Promise) Struct() (SIGCtrl, error) {
	s, err := p.Pipeline.Struct()
	return SIGCtrl{s}, err
}

func (p SIGCtrl_Promise) PollReq() SIGPoll_Promise {
	return SIGPoll_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p SIGCtrl_Promise) PollRep() SIGPoll_Promise {
	return SIGPoll_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type SIGPoll struct{ capnp.Struct }

// SIGPoll_TypeID is the unique identifier for the type SIGPoll.
const SIGPoll_TypeID = 0x9ad73a0235a46141

func NewSIGPoll(s *capnp.Segment) (SIGPoll, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return SIGPoll{st}, err
}

func NewRootSIGPoll(s *capnp.Segment) (SIGPoll, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return SIGPoll{st}, err
}

func ReadRootSIGPoll(msg *capnp.Message) (SIGPoll, error) {
	root, err := msg.RootPtr()
	return SIGPoll{root.Struct()}, err
}

func (s SIGPoll) String() string {
	str, _ := text.Marshal(0x9ad73a0235a46141, s.Struct)
	return str
}

func (s SIGPoll) Addr() (SIGAddr, error) {
	p, err := s.Struct.Ptr(0)
	return SIGAddr{Struct: p.Struct()}, err
}

func (s SIGPoll) HasAddr() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SIGPoll) SetAddr(v SIGAddr) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewAddr sets the addr field to a newly
// allocated SIGAddr struct, preferring placement in s's segment.
func (s SIGPoll) NewAddr() (SIGAddr, error) {
	ss, err := NewSIGAddr(s.Struct.Segment())
	if err != nil {
		return SIGAddr{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s SIGPoll) Session() uint8 {
	return s.Struct.Uint8(0)
}

func (s SIGPoll) SetSession(v uint8) {
	s.Struct.SetUint8(0, v)
}

// SIGPoll_List is a list of SIGPoll.
type SIGPoll_List struct{ capnp.List }

// NewSIGPoll creates a new list of SIGPoll.
func NewSIGPoll_List(s *capnp.Segment, sz int32) (SIGPoll_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return SIGPoll_List{l}, err
}

func (s SIGPoll_List) At(i int) SIGPoll { return SIGPoll{s.List.Struct(i)} }

func (s SIGPoll_List) Set(i int, v SIGPoll) error { return s.List.SetStruct(i, v.Struct) }

func (s SIGPoll_List) String() string {
	str, _ := text.MarshalList(0x9ad73a0235a46141, s.List)
	return str
}

// SIGPoll_Promise is a wrapper for a SIGPoll promised by a client call.
type SIGPoll_Promise struct{ *capnp.Pipeline }

func (p SIGPoll_Promise) Struct() (SIGPoll, error) {
	s, err := p.Pipeline.Struct()
	return SIGPoll{s}, err
}

func (p SIGPoll_Promise) Addr() SIGAddr_Promise {
	return SIGAddr_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type SIGAddr struct{ capnp.Struct }

// SIGAddr_TypeID is the unique identifier for the type SIGAddr.
const SIGAddr_TypeID = 0xddf1fce11d9b0028

func NewSIGAddr(s *capnp.Segment) (SIGAddr, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return SIGAddr{st}, err
}

func NewRootSIGAddr(s *capnp.Segment) (SIGAddr, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return SIGAddr{st}, err
}

func ReadRootSIGAddr(msg *capnp.Message) (SIGAddr, error) {
	root, err := msg.RootPtr()
	return SIGAddr{root.Struct()}, err
}

func (s SIGAddr) String() string {
	str, _ := text.Marshal(0xddf1fce11d9b0028, s.Struct)
	return str
}

func (s SIGAddr) Ctrl() (HostInfo, error) {
	p, err := s.Struct.Ptr(0)
	return HostInfo{Struct: p.Struct()}, err
}

func (s SIGAddr) HasCtrl() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SIGAddr) SetCtrl(v HostInfo) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewCtrl sets the ctrl field to a newly
// allocated HostInfo struct, preferring placement in s's segment.
func (s SIGAddr) NewCtrl() (HostInfo, error) {
	ss, err := NewHostInfo(s.Struct.Segment())
	if err != nil {
		return HostInfo{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s SIGAddr) EncapPort() uint16 {
	return s.Struct.Uint16(0)
}

func (s SIGAddr) SetEncapPort(v uint16) {
	s.Struct.SetUint16(0, v)
}

// SIGAddr_List is a list of SIGAddr.
type SIGAddr_List struct{ capnp.List }

// NewSIGAddr creates a new list of SIGAddr.
func NewSIGAddr_List(s *capnp.Segment, sz int32) (SIGAddr_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return SIGAddr_List{l}, err
}

func (s SIGAddr_List) At(i int) SIGAddr { return SIGAddr{s.List.Struct(i)} }

func (s SIGAddr_List) Set(i int, v SIGAddr) error { return s.List.SetStruct(i, v.Struct) }

func (s SIGAddr_List) String() string {
	str, _ := text.MarshalList(0xddf1fce11d9b0028, s.List)
	return str
}

// SIGAddr_Promise is a wrapper for a SIGAddr promised by a client call.
type SIGAddr_Promise struct{ *capnp.Pipeline }

func (p SIGAddr_Promise) Struct() (SIGAddr, error) {
	s, err := p.Pipeline.Struct()
	return SIGAddr{s}, err
}

func (p SIGAddr_Promise) Ctrl() HostInfo_Promise {
	return HostInfo_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

const schema_8273379c3e06a721 = "x\xda\x8c\x91\xb1\xeb\xd3@\x1c\xc5\xdf\xbbk\x92\xdfO" +
	",Mh\xa7\x82\xa8P\xb1\x15\x15[,jA\xb1\x8a" +
	"\x88[\xaf\xae\"\x86&\xd8@L\xd2$\xa5c\xc1?" +
	"\xa1n\xea(8\xb9\xf8\x7f88\x88\x93C'gw" +
	"\xeb\xc9\x11\xb5C\x1d~\xd3\x17\xbe\xef\xdd}\xee\xdds" +
	"?\xdc\x15}\xabA@\x9d\xb6l=\xf6\xdf\x0d\xc5\xe8" +
	"\xebk\xa8S\xa4>\xff\xde\xbe\xf3\xf6F\xf1\x12\x16\x1d" +
	"\xc0{\xb1\xf1\x96f.V\xe0\xae\xfb\xe6\xcc\xf6\xe7\x8f" +
	"o\xff\xf3}\xdax_\xcc\xfc\xbc\x02\xf5\xd1\xedA\xd1" +
	"\xeb<\xdd\x9a\x1b\xc5\xde\xf9\x80\x8ed\xad\xd9\xe7\xa6y" +
	"\xcb\x1cj\x0e\xf9\x1d\xd4E\xf4\xfc\xea\xcc\xcf\x12f\xa3" +
	"\xc7\x8f\x1eN\xd2\x98\xf1\x84TG\xb2\x06\xd4\x08x\xbd" +
	"K\x80\xeaH\xaak\x82d\x8bfw\xe5\x1e\xa0\xba\x92" +
	"\xea\xba`\xc3\x0f\x82\x9c\xee\xdf\xe7\x81t\xc1u\x11\x16" +
	"E\x94&\xb4!h\x1f`\xc6A\xc0\xfc\x04\x98)\xa0" +
	".K\xaa\x9b\x82\x8dY\x99\xc7t\xf5\xb9\x0b\xafV\xd6" +
	"\xc5\xf6GT \x1d&3?\x9b\xa49X\xd2\x81\xa0" +
	"s\x00\xbb_\xe6U&\xf7\x1f\xcco\x03\xea\x89\xa4\x9a" +
	"\x0b\xd6\xa9u\x85\x0b\x07\x80z&\xa9b\xc1\xba\xf8\xa5" +
	"[\x14\x80\x17\x99\xac\x81\xa4\xca\x04\xebr\xa7[\x94\xa6" +
	"\x19\xb3\x9dK\xaaRPF\x01\x8f!x\x0c\x9e]&" +
	"EX\xc2^gi\x1cO\xc3\x05\xdd}\xc1\x7f~\xa6" +
	"R\xb2C\xe5w\x00\x00\x00\xff\xfflyx\xa7"

func init() {
	schemas.Register(schema_8273379c3e06a721,
		0x9ad73a0235a46141,
		0xddf1fce11d9b0028,
		0xe15e242973323d08)
}