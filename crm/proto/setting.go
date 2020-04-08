package geiqin_srv_crm

import "github.com/gogo/protobuf/jsonpb"

var pbMarshaler jsonpb.Marshaler

func init() {
	pbMarshaler = jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
		EnumsAsInts:  true,
	}
}
