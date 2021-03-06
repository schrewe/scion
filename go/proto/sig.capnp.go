// Code generated by capnpc-go. DO NOT EDIT.

package proto

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type HostInfo struct{ capnp.Struct }
type HostInfo_addrs HostInfo

// HostInfo_TypeID is the unique identifier for the type HostInfo.
const HostInfo_TypeID = 0xe02e04db98f2686a

func NewHostInfo(s *capnp.Segment) (HostInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return HostInfo{st}, err
}

func NewRootHostInfo(s *capnp.Segment) (HostInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return HostInfo{st}, err
}

func ReadRootHostInfo(msg *capnp.Message) (HostInfo, error) {
	root, err := msg.RootPtr()
	return HostInfo{root.Struct()}, err
}

func (s HostInfo) String() string {
	str, _ := text.Marshal(0xe02e04db98f2686a, s.Struct)
	return str
}

func (s HostInfo) Port() uint16 {
	return s.Struct.Uint16(0)
}

func (s HostInfo) SetPort(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s HostInfo) Addrs() HostInfo_addrs { return HostInfo_addrs(s) }

func (s HostInfo_addrs) Ipv4() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s HostInfo_addrs) HasIpv4() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HostInfo_addrs) SetIpv4(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s HostInfo_addrs) Ipv6() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s HostInfo_addrs) HasIpv6() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HostInfo_addrs) SetIpv6(v []byte) error {
	return s.Struct.SetData(1, v)
}

// HostInfo_List is a list of HostInfo.
type HostInfo_List struct{ capnp.List }

// NewHostInfo creates a new list of HostInfo.
func NewHostInfo_List(s *capnp.Segment, sz int32) (HostInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return HostInfo_List{l}, err
}

func (s HostInfo_List) At(i int) HostInfo { return HostInfo{s.List.Struct(i)} }

func (s HostInfo_List) Set(i int, v HostInfo) error { return s.List.SetStruct(i, v.Struct) }

func (s HostInfo_List) String() string {
	str, _ := text.MarshalList(0xe02e04db98f2686a, s.List)
	return str
}

// HostInfo_Promise is a wrapper for a HostInfo promised by a client call.
type HostInfo_Promise struct{ *capnp.Pipeline }

func (p HostInfo_Promise) Struct() (HostInfo, error) {
	s, err := p.Pipeline.Struct()
	return HostInfo{s}, err
}

func (p HostInfo_Promise) Addrs() HostInfo_addrs_Promise { return HostInfo_addrs_Promise{p.Pipeline} }

// HostInfo_addrs_Promise is a wrapper for a HostInfo_addrs promised by a client call.
type HostInfo_addrs_Promise struct{ *capnp.Pipeline }

func (p HostInfo_addrs_Promise) Struct() (HostInfo_addrs, error) {
	s, err := p.Pipeline.Struct()
	return HostInfo_addrs{s}, err
}

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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SIGAddr{st}, err
}

func NewRootSIGAddr(s *capnp.Segment) (SIGAddr, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
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

func (s SIGAddr) Data() (HostInfo, error) {
	p, err := s.Struct.Ptr(1)
	return HostInfo{Struct: p.Struct()}, err
}

func (s SIGAddr) HasData() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SIGAddr) SetData(v HostInfo) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewData sets the data field to a newly
// allocated HostInfo struct, preferring placement in s's segment.
func (s SIGAddr) NewData() (HostInfo, error) {
	ss, err := NewHostInfo(s.Struct.Segment())
	if err != nil {
		return HostInfo{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// SIGAddr_List is a list of SIGAddr.
type SIGAddr_List struct{ capnp.List }

// NewSIGAddr creates a new list of SIGAddr.
func NewSIGAddr_List(s *capnp.Segment, sz int32) (SIGAddr_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
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

func (p SIGAddr_Promise) Data() HostInfo_Promise {
	return HostInfo_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

const schema_8273379c3e06a721 = "x\xda|\x92\xc1k\x13A\x14\xc6\xbf\xefM\x92m\x0b" +
	"!Y6^\x14)B\x85\xa4\x90\xd0\x86\xaa\x10P\xac" +
	"\"\xda[G\x0f^D\\\xba\xa9\x89\xac\xd95\xbbZ" +
	"<\x15\x04\x0f\x1e\xbc\x04\x04\xa9\x0a\x1e\x04\xff\x0cO\x1e" +
	"\x14<{\x10\xc1\"\"\x1e<\xe8\xd9\xba2\xbb\xb6Y" +
	"R\xf44\xcc\x9bo\xbe\x8f\xdf{o\x81<-\x8b\xc5" +
	"W\x04t\xadXJ\x96\xdd\x17\xc7\xa4\xf3~\x0bz\x86" +
	"L\x8e\xbc,\x9dzz\"\xba\x87\"-\xc0~=\xb2" +
	"\xdf\x99\xf3\xcd\x06\x98|~~\xf7\xf2\xa1G\xf5g\xd0" +
	"\x0e\x99\xdc\xe8\xfd|\xfc\xa1\xd0\xfa\x84\x03b\x11p\x8e" +
	"\xf2-\xe84\xb8\x01\xee\xd4\x9f\x1c\xde\xfe\xf5\xe3\xa3=" +
	"\x937\x14\x0bp\xees\xe4<4\xd6\xce\x03\xa3\x1c\xbb" +
	"L\xa4\xa7\xe2/\xdcr\xbe\xa7\xe2o\xa9x\xead;" +
	"j\xcc]\xdd6b\x19\x8b\xcf\xd1R,8\x0d\x199" +
	"\x8b\xe9\xbf\xa6|\x05\x93\xa8\x7f\xbd\xb5\xe6\x86\x03\x86\x9d" +
	"K+\xe7W\x03\x9f\xfe*\xa9\xa7T\x01(\x10\xb0\x1b" +
	"\xf3\x80\x9eS\xd4\x0bB\xb2FSk\x9e\x01t]Q" +
	"/\x09+\xae\xe7\x0dY\xdd\xa5\x01Y\x057\xa3n\x14" +
	"\xf5\x83\x01K\x10\x96r1\x12v.\x04Q\xbc2X" +
	"\x0fZ\xe6c\x04dY\x99q>\xcc6ib\xd2\xe6" +
	"si\xfd\xf0\xce\x12\xcb\x10\x96\x91^\x8e\xef^&H" +
	"\x96=\x8f\xc3\x7f\x93\xd8{(y\xf3\xb5x\xe8\xb3:" +
	"nw\x06S\xf1\xdc\xd8\xdd_\xce\x07fLj=\xf8" +
	"_\xef\x98[\x0e\xbb\xd9\x86T\xc2`\x18\xd3\x82\xd0\x02" +
	"g\xd3vLP\x9c\x8d\x87\xd9<\xaa{\x9e\xeeA@" +
	"_Q\xd4=a\x99I\x92at\xdb\x80\xbe\xa6\xa8}" +
	"aY~'Y\xe7\xfafN\x9e\xa2\x0e\x85e\xb5\x93" +
	"\xd4\xa8\x00\xfb\xa6\xa9\xf6\x14u,T}\x8f\xd3\x10N" +
	"\x83\xb3\xb7\x07Q7Fi3\x0c|\xffb\xf7\x16\xab" +
	"\xe3\xad\xff;\xd5\xec%\xdc\xff\xf2'\x00\x00\xff\xff-" +
	"a\xbb\x1d"

func init() {
	schemas.Register(schema_8273379c3e06a721,
		0x9ad73a0235a46141,
		0x9d28951b5779a0e3,
		0xddf1fce11d9b0028,
		0xe02e04db98f2686a,
		0xe15e242973323d08)
}
