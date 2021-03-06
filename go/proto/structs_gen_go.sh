#!/bin/bash

ROOTTYPES="ASEntry CtrlPld PathSegment PathSegmentSignedData RevInfo SignedBlob SignedCtrlPld SVCResolutionReply ColibriRequestPayload"

cat <<EOF
// Code generated by go/proto/structs_gen_go.sh; DO NOT EDIT.

package proto

import (
	"zombiezen.com/go/capnproto2"

	"github.com/scionproto/scion/go/lib/common"
)

// NewRootStruct calls the appropriate NewRoot<x> function corresponding to the capnp proto type ID,
// and returns the inner capnp.Struct that it receives. This allows the helper
// functions in cereal.go to support generic capnp root struct types.
func NewRootStruct(id ProtoIdType, seg *capnp.Segment) (capnp.Struct, error) {
	var blank capnp.Struct
	switch id {
EOF

for roottype in $ROOTTYPES
do
    cat <<EOF
	case ${roottype}_TypeID:
		v, err := NewRoot$roottype(seg)
		if err != nil {
			return blank, common.NewBasicError("Error creating new $roottype capnp struct", err)
		}
		return v.Struct, nil
EOF
done

cat <<EOF
	}
	return blank, common.NewBasicError(
		"Unsupported capnp struct type (i.e. not listed in go/proto/structs_gen_go.sh:ROOTTYPES)",
		nil,
		"id", id,
	)
}

EOF

for roottype in $ROOTTYPES
do
    cat <<EOF
func (s $roottype) GetStruct() capnp.Struct {
	return s.Struct
}
EOF
done
