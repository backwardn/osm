// Code generated by protoc-gen-go.
// source: osm.proto
// DO NOT EDIT!

/*
Package osmpb is a generated protocol buffer package.

It is generated from these files:
	osm.proto

It has these top-level messages:
	Changeset
	Bounds
	Change
	OSM
	Node
	Info
	DenseNodes
	DenseInfo
	Way
	Relation
*/
package osmpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Relation_MemberType int32

const (
	Relation_NODE     Relation_MemberType = 0
	Relation_WAY      Relation_MemberType = 1
	Relation_RELATION Relation_MemberType = 2
)

var Relation_MemberType_name = map[int32]string{
	0: "NODE",
	1: "WAY",
	2: "RELATION",
}
var Relation_MemberType_value = map[string]int32{
	"NODE":     0,
	"WAY":      1,
	"RELATION": 2,
}

func (x Relation_MemberType) Enum() *Relation_MemberType {
	p := new(Relation_MemberType)
	*p = x
	return p
}
func (x Relation_MemberType) String() string {
	return proto.EnumName(Relation_MemberType_name, int32(x))
}
func (x *Relation_MemberType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Relation_MemberType_value, data, "Relation_MemberType")
	if err != nil {
		return err
	}
	*x = Relation_MemberType(value)
	return nil
}
func (Relation_MemberType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{9, 0} }

type Changeset struct {
	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys      []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals      []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	UserId    *int32   `protobuf:"varint,5,opt,name=user_id" json:"user_id,omitempty"`
	UserSid   *uint32  `protobuf:"varint,6,opt,name=user_sid" json:"user_sid,omitempty"`
	CreatedAt *int64   `protobuf:"varint,7,opt,name=created_at" json:"created_at,omitempty"`
	ClosedAt  *int64   `protobuf:"varint,8,opt,name=closed_at" json:"closed_at,omitempty"`
	Open      *bool    `protobuf:"varint,9,opt,name=open" json:"open,omitempty"`
	Bounds    *Bounds  `protobuf:"bytes,10,opt,name=bounds" json:"bounds,omitempty"`
	Change    *Change  `protobuf:"bytes,11,opt,name=change" json:"change,omitempty"`
	// contains the tag strings for everything
	// in this entire changeset.
	Strings          []string `protobuf:"bytes,20,rep,name=strings" json:"strings,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Changeset) Reset()                    { *m = Changeset{} }
func (m *Changeset) String() string            { return proto.CompactTextString(m) }
func (*Changeset) ProtoMessage()               {}
func (*Changeset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Changeset) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Changeset) GetKeys() []uint32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Changeset) GetVals() []uint32 {
	if m != nil {
		return m.Vals
	}
	return nil
}

func (m *Changeset) GetUserId() int32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Changeset) GetUserSid() uint32 {
	if m != nil && m.UserSid != nil {
		return *m.UserSid
	}
	return 0
}

func (m *Changeset) GetCreatedAt() int64 {
	if m != nil && m.CreatedAt != nil {
		return *m.CreatedAt
	}
	return 0
}

func (m *Changeset) GetClosedAt() int64 {
	if m != nil && m.ClosedAt != nil {
		return *m.ClosedAt
	}
	return 0
}

func (m *Changeset) GetOpen() bool {
	if m != nil && m.Open != nil {
		return *m.Open
	}
	return false
}

func (m *Changeset) GetBounds() *Bounds {
	if m != nil {
		return m.Bounds
	}
	return nil
}

func (m *Changeset) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

func (m *Changeset) GetStrings() []string {
	if m != nil {
		return m.Strings
	}
	return nil
}

type Bounds struct {
	MinLng           *int64 `protobuf:"zigzag64,1,req,name=min_lng" json:"min_lng,omitempty"`
	MaxLng           *int64 `protobuf:"zigzag64,2,req,name=max_lng" json:"max_lng,omitempty"`
	MinLat           *int64 `protobuf:"zigzag64,3,req,name=min_lat" json:"min_lat,omitempty"`
	MaxLat           *int64 `protobuf:"zigzag64,4,req,name=max_lat" json:"max_lat,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Bounds) Reset()                    { *m = Bounds{} }
func (m *Bounds) String() string            { return proto.CompactTextString(m) }
func (*Bounds) ProtoMessage()               {}
func (*Bounds) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Bounds) GetMinLng() int64 {
	if m != nil && m.MinLng != nil {
		return *m.MinLng
	}
	return 0
}

func (m *Bounds) GetMaxLng() int64 {
	if m != nil && m.MaxLng != nil {
		return *m.MaxLng
	}
	return 0
}

func (m *Bounds) GetMinLat() int64 {
	if m != nil && m.MinLat != nil {
		return *m.MinLat
	}
	return 0
}

func (m *Bounds) GetMaxLat() int64 {
	if m != nil && m.MaxLat != nil {
		return *m.MaxLat
	}
	return 0
}

type Change struct {
	Create *OSM `protobuf:"bytes,1,opt,name=create" json:"create,omitempty"`
	Modify *OSM `protobuf:"bytes,2,opt,name=modify" json:"modify,omitempty"`
	Delete *OSM `protobuf:"bytes,3,opt,name=delete" json:"delete,omitempty"`
	// contains the tag strings if this is the root of the data.
	Strings          []string `protobuf:"bytes,20,rep,name=strings" json:"strings,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Change) Reset()                    { *m = Change{} }
func (m *Change) String() string            { return proto.CompactTextString(m) }
func (*Change) ProtoMessage()               {}
func (*Change) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Change) GetCreate() *OSM {
	if m != nil {
		return m.Create
	}
	return nil
}

func (m *Change) GetModify() *OSM {
	if m != nil {
		return m.Modify
	}
	return nil
}

func (m *Change) GetDelete() *OSM {
	if m != nil {
		return m.Delete
	}
	return nil
}

func (m *Change) GetStrings() []string {
	if m != nil {
		return m.Strings
	}
	return nil
}

type OSM struct {
	Bounds *Bounds `protobuf:"bytes,1,opt,name=bounds" json:"bounds,omitempty"`
	// an encoded should have either nodes or a dense_nodes, but not both.
	Nodes      []*Node     `protobuf:"bytes,2,rep,name=nodes" json:"nodes,omitempty"`
	DenseNodes *DenseNodes `protobuf:"bytes,3,opt,name=dense_nodes" json:"dense_nodes,omitempty"`
	Ways       []*Way      `protobuf:"bytes,4,rep,name=ways" json:"ways,omitempty"`
	Relations  []*Relation `protobuf:"bytes,5,rep,name=relations" json:"relations,omitempty"`
	// contains the tag strings if this is the root of the data.
	Strings          []string `protobuf:"bytes,20,rep,name=strings" json:"strings,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *OSM) Reset()                    { *m = OSM{} }
func (m *OSM) String() string            { return proto.CompactTextString(m) }
func (*OSM) ProtoMessage()               {}
func (*OSM) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OSM) GetBounds() *Bounds {
	if m != nil {
		return m.Bounds
	}
	return nil
}

func (m *OSM) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *OSM) GetDenseNodes() *DenseNodes {
	if m != nil {
		return m.DenseNodes
	}
	return nil
}

func (m *OSM) GetWays() []*Way {
	if m != nil {
		return m.Ways
	}
	return nil
}

func (m *OSM) GetRelations() []*Relation {
	if m != nil {
		return m.Relations
	}
	return nil
}

func (m *OSM) GetStrings() []string {
	if m != nil {
		return m.Strings
	}
	return nil
}

type Node struct {
	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys             []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals             []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	Info             *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	Lat              *int64   `protobuf:"zigzag64,8,req,name=lat" json:"lat,omitempty"`
	Lng              *int64   `protobuf:"zigzag64,9,req,name=lng" json:"lng,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Node) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Node) GetKeys() []uint32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Node) GetVals() []uint32 {
	if m != nil {
		return m.Vals
	}
	return nil
}

func (m *Node) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Node) GetLat() int64 {
	if m != nil && m.Lat != nil {
		return *m.Lat
	}
	return 0
}

func (m *Node) GetLng() int64 {
	if m != nil && m.Lng != nil {
		return *m.Lng
	}
	return 0
}

type Info struct {
	Version   *int32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Timestamp *int64 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	// these will be omitted if the object represents one changeset
	// since they will be all the same.
	ChangesetId *int64  `protobuf:"varint,3,opt,name=changeset_id" json:"changeset_id,omitempty"`
	UserId      *int32  `protobuf:"varint,4,opt,name=user_id" json:"user_id,omitempty"`
	UserSid     *uint32 `protobuf:"varint,5,opt,name=user_sid" json:"user_sid,omitempty"`
	// The visible flag is used to store history information. It indicates that
	// the current object version has been created by a delete operation on the
	// OSM API. This info may be omitted if it can be inferred from its group
	// ie. create, modify, delete.
	Visible          *bool  `protobuf:"varint,6,opt,name=visible" json:"visible,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Info) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Info) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Info) GetChangesetId() int64 {
	if m != nil && m.ChangesetId != nil {
		return *m.ChangesetId
	}
	return 0
}

func (m *Info) GetUserId() int32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Info) GetUserSid() uint32 {
	if m != nil && m.UserSid != nil {
		return *m.UserSid
	}
	return 0
}

func (m *Info) GetVisible() bool {
	if m != nil && m.Visible != nil {
		return *m.Visible
	}
	return false
}

type DenseNodes struct {
	Id        []int64    `protobuf:"zigzag64,1,rep,packed,name=id" json:"id,omitempty"`
	DenseInfo *DenseInfo `protobuf:"bytes,5,opt,name=dense_info" json:"dense_info,omitempty"`
	Lat       []int64    `protobuf:"zigzag64,8,rep,packed,name=lat" json:"lat,omitempty"`
	Lng       []int64    `protobuf:"zigzag64,9,rep,packed,name=lng" json:"lng,omitempty"`
	// Special packing of keys and vals into one array. We use a single stringid
	// of 0 to delimit when the tags of a node ends and the tags of the next node
	// begin. The storage pattern is: ((<keyid> <valid>)* '0' )* As an exception,
	// if no node in the current block has any key/value pairs, this array does
	// not contain any delimiters, but is simply empty.
	KeysVals         []uint32 `protobuf:"varint,10,rep,packed,name=keys_vals" json:"keys_vals,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DenseNodes) Reset()                    { *m = DenseNodes{} }
func (m *DenseNodes) String() string            { return proto.CompactTextString(m) }
func (*DenseNodes) ProtoMessage()               {}
func (*DenseNodes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DenseNodes) GetId() []int64 {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DenseNodes) GetDenseInfo() *DenseInfo {
	if m != nil {
		return m.DenseInfo
	}
	return nil
}

func (m *DenseNodes) GetLat() []int64 {
	if m != nil {
		return m.Lat
	}
	return nil
}

func (m *DenseNodes) GetLng() []int64 {
	if m != nil {
		return m.Lng
	}
	return nil
}

func (m *DenseNodes) GetKeysVals() []uint32 {
	if m != nil {
		return m.KeysVals
	}
	return nil
}

type DenseInfo struct {
	Version   []int32 `protobuf:"varint,1,rep,packed,name=version" json:"version,omitempty"`
	Timestamp []int64 `protobuf:"zigzag64,2,rep,packed,name=timestamp" json:"timestamp,omitempty"`
	// these will be omitted if the object represents one changeset
	// and these will be all the same.
	ChangesetId []int64 `protobuf:"zigzag64,3,rep,packed,name=changeset_id" json:"changeset_id,omitempty"`
	UserId      []int32 `protobuf:"zigzag32,4,rep,packed,name=user_id" json:"user_id,omitempty"`
	UserSid     []int32 `protobuf:"zigzag32,5,rep,packed,name=user_sid" json:"user_sid,omitempty"`
	// The visible flag is used to store history information. It indicates that
	// the current object version has been created by a delete operation on the
	// OSM API. This info may be omitted if it can be inferred from it's group
	// ie. create, modify, delete.
	Visible          []bool `protobuf:"varint,6,rep,packed,name=visible" json:"visible,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DenseInfo) Reset()                    { *m = DenseInfo{} }
func (m *DenseInfo) String() string            { return proto.CompactTextString(m) }
func (*DenseInfo) ProtoMessage()               {}
func (*DenseInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DenseInfo) GetVersion() []int32 {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *DenseInfo) GetTimestamp() []int64 {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *DenseInfo) GetChangesetId() []int64 {
	if m != nil {
		return m.ChangesetId
	}
	return nil
}

func (m *DenseInfo) GetUserId() []int32 {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *DenseInfo) GetUserSid() []int32 {
	if m != nil {
		return m.UserSid
	}
	return nil
}

func (m *DenseInfo) GetVisible() []bool {
	if m != nil {
		return m.Visible
	}
	return nil
}

type Way struct {
	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys             []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals             []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	Info             *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	Refs             []int64  `protobuf:"zigzag64,8,rep,packed,name=refs" json:"refs,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Way) Reset()                    { *m = Way{} }
func (m *Way) String() string            { return proto.CompactTextString(m) }
func (*Way) ProtoMessage()               {}
func (*Way) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Way) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Way) GetKeys() []uint32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Way) GetVals() []uint32 {
	if m != nil {
		return m.Vals
	}
	return nil
}

func (m *Way) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Way) GetRefs() []int64 {
	if m != nil {
		return m.Refs
	}
	return nil
}

type Relation struct {
	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	Info *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	// Parallel arrays
	// Roles has been changed int32 -> uint32 form the osm proto,
	// this is for consistency and backwards compatible.
	Roles            []uint32              `protobuf:"varint,8,rep,packed,name=roles" json:"roles,omitempty"`
	Refs             []int64               `protobuf:"zigzag64,9,rep,packed,name=refs" json:"refs,omitempty"`
	Types            []Relation_MemberType `protobuf:"varint,10,rep,packed,name=types,enum=osm.Relation_MemberType" json:"types,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Relation) Reset()                    { *m = Relation{} }
func (m *Relation) String() string            { return proto.CompactTextString(m) }
func (*Relation) ProtoMessage()               {}
func (*Relation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Relation) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Relation) GetKeys() []uint32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Relation) GetVals() []uint32 {
	if m != nil {
		return m.Vals
	}
	return nil
}

func (m *Relation) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Relation) GetRoles() []uint32 {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Relation) GetRefs() []int64 {
	if m != nil {
		return m.Refs
	}
	return nil
}

func (m *Relation) GetTypes() []Relation_MemberType {
	if m != nil {
		return m.Types
	}
	return nil
}

func init() {
	proto.RegisterType((*Changeset)(nil), "osm.Changeset")
	proto.RegisterType((*Bounds)(nil), "osm.Bounds")
	proto.RegisterType((*Change)(nil), "osm.Change")
	proto.RegisterType((*OSM)(nil), "osm.OSM")
	proto.RegisterType((*Node)(nil), "osm.Node")
	proto.RegisterType((*Info)(nil), "osm.Info")
	proto.RegisterType((*DenseNodes)(nil), "osm.DenseNodes")
	proto.RegisterType((*DenseInfo)(nil), "osm.DenseInfo")
	proto.RegisterType((*Way)(nil), "osm.Way")
	proto.RegisterType((*Relation)(nil), "osm.Relation")
	proto.RegisterEnum("osm.Relation_MemberType", Relation_MemberType_name, Relation_MemberType_value)
}

var fileDescriptor0 = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x5e, 0x3b, 0xb1, 0xc7, 0xfd, 0x49, 0xb7, 0x05, 0x56, 0xe2, 0x12, 0x59, 0x1c, 0x2c,
	0x21, 0x7a, 0xc8, 0x1b, 0x34, 0xb4, 0x87, 0x4a, 0x6d, 0x23, 0x41, 0x25, 0x04, 0x97, 0xe0, 0xc4,
	0x9b, 0x60, 0xe1, 0x9f, 0xc8, 0xeb, 0x16, 0xf2, 0x18, 0x3c, 0x08, 0x6f, 0xc3, 0x99, 0x67, 0x61,
	0x76, 0x6c, 0xc7, 0x0e, 0x3e, 0x21, 0xf5, 0x16, 0x7f, 0xdf, 0xcc, 0xec, 0xf7, 0x7d, 0x33, 0x01,
	0x37, 0x57, 0xe9, 0xf9, 0xa6, 0xc8, 0xcb, 0x9c, 0x33, 0xfc, 0xe9, 0xff, 0x31, 0xc0, 0x7d, 0xf7,
	0x35, 0xcc, 0xd6, 0x52, 0xc9, 0x92, 0x03, 0x98, 0x71, 0x24, 0x8c, 0xb1, 0x19, 0x30, 0x3e, 0x02,
	0xeb, 0x9b, 0xdc, 0x2a, 0x61, 0x8e, 0x59, 0x70, 0x38, 0x35, 0x47, 0x86, 0x46, 0x1e, 0xc3, 0x44,
	0x09, 0xb6, 0x43, 0x8e, 0x61, 0xf8, 0xa0, 0x64, 0x31, 0xc7, 0x26, 0x7b, 0x6c, 0x04, 0x36, 0x96,
	0x38, 0x04, 0x28, 0x44, 0x06, 0x88, 0x1c, 0x72, 0x9c, 0xb9, 0x2c, 0x64, 0x58, 0xca, 0x68, 0x1e,
	0x96, 0x62, 0x88, 0x18, 0xe3, 0x27, 0xe0, 0x2e, 0x93, 0x5c, 0x55, 0x90, 0x43, 0xd0, 0x01, 0x58,
	0xf9, 0x46, 0x66, 0xc2, 0xc5, 0x2f, 0x87, 0xbf, 0x82, 0xc1, 0x22, 0x7f, 0xc8, 0x22, 0x25, 0x00,
	0xbf, 0xbd, 0x89, 0x77, 0xae, 0x65, 0x4f, 0x09, 0xd2, 0xe4, 0x92, 0x14, 0x0b, 0xaf, 0x43, 0x56,
	0x26, 0xb4, 0x22, 0x55, 0x16, 0x71, 0xb6, 0x56, 0xe2, 0x0c, 0x65, 0xba, 0xfe, 0x0d, 0x0c, 0xea,
	0x3e, 0xa4, 0xd2, 0x38, 0x9b, 0x27, 0xd9, 0x9a, 0x1c, 0x72, 0x02, 0xc2, 0x1f, 0x04, 0x98, 0x3b,
	0x40, 0x57, 0xa0, 0x2a, 0xb6, 0x57, 0x81, 0x80, 0xa5, 0x01, 0x7f, 0x0d, 0x83, 0xfa, 0x21, 0x81,
	0x2a, 0xc8, 0x17, 0x0e, 0xd3, 0x2a, 0x1c, 0x52, 0x31, 0xfb, 0x70, 0xab, 0x99, 0x34, 0x8f, 0xe2,
	0xd5, 0x16, 0xa7, 0xf6, 0x98, 0x48, 0x26, 0x12, 0x7b, 0xd8, 0x3f, 0x4c, 0x4f, 0xf6, 0x2f, 0x03,
	0x98, 0x26, 0xda, 0x24, 0x8c, 0x7e, 0x12, 0x02, 0xec, 0x2c, 0x8f, 0x64, 0xb5, 0x23, 0x6f, 0xe2,
	0x12, 0x77, 0x87, 0x08, 0x7f, 0x0d, 0x5e, 0x24, 0x33, 0x25, 0xe7, 0x15, 0x5f, 0x3d, 0x77, 0x4c,
	0xfc, 0xa5, 0xc6, 0x75, 0x91, 0xe2, 0x2f, 0xc0, 0xfa, 0x1e, 0xe2, 0x8a, 0x2d, 0x6a, 0xaf, 0xd4,
	0x7c, 0x0c, 0xb7, 0x7c, 0x0c, 0x6e, 0x21, 0xd1, 0x74, 0x9c, 0x67, 0x0a, 0x17, 0xab, 0xc9, 0x43,
	0x22, 0xdf, 0xd7, 0x68, 0x5f, 0x6f, 0x0c, 0x16, 0x3d, 0xfc, 0xbf, 0x17, 0xf4, 0x12, 0xac, 0x38,
	0x5b, 0xe5, 0x28, 0xc1, 0xd8, 0x39, 0xb8, 0x46, 0x80, 0x7b, 0xc0, 0x12, 0xba, 0x0e, 0xbd, 0x07,
	0xfd, 0x81, 0x5b, 0x72, 0x69, 0x07, 0x25, 0x58, 0x54, 0x81, 0x1a, 0x1e, 0x65, 0xa1, 0x50, 0x0e,
	0x65, 0x63, 0xeb, 0xb3, 0x2a, 0xe3, 0x54, 0xaa, 0x32, 0x4c, 0x37, 0x94, 0x3d, 0xe3, 0x67, 0x70,
	0xb0, 0x6c, 0xae, 0x5b, 0x5f, 0x29, 0x23, 0xb4, 0x73, 0xb6, 0x56, 0xef, 0x6c, 0x6d, 0x3a, 0x5b,
	0x3d, 0x3c, 0x56, 0xf1, 0x22, 0x91, 0x74, 0xc7, 0x8e, 0xbf, 0x05, 0xe8, 0x24, 0x77, 0x54, 0xdb,
	0x64, 0x01, 0x27, 0x1b, 0x3e, 0x40, 0x95, 0x37, 0x99, 0xb1, 0xc9, 0xcc, 0x51, 0x1b, 0x77, 0xad,
	0xb7, 0x76, 0xd4, 0x34, 0x1d, 0x37, 0xae, 0x1a, 0xe0, 0x39, 0xb8, 0x3a, 0xb0, 0x39, 0x65, 0x04,
	0x4d, 0x46, 0xfe, 0x4f, 0xfc, 0x8f, 0xb6, 0x63, 0x4e, 0xbb, 0xb6, 0x59, 0x60, 0x37, 0x9d, 0x5d,
	0xeb, 0xcd, 0x40, 0xd1, 0xb3, 0xdf, 0x30, 0xa7, 0xdd, 0x08, 0x58, 0x70, 0x42, 0xe0, 0xd9, 0x5e,
	0x0c, 0x0d, 0x7a, 0xda, 0x8d, 0x82, 0x05, 0x0e, 0x69, 0xfa, 0x02, 0x4c, 0x5f, 0xca, 0x93, 0xad,
	0x1b, 0x4b, 0x0b, 0xb9, 0x52, 0x6d, 0x3a, 0xfe, 0x6f, 0x03, 0x9c, 0xdd, 0xbd, 0x3d, 0xd9, 0x3b,
	0x27, 0x60, 0x17, 0x79, 0x22, 0xab, 0x87, 0x76, 0xdd, 0xf4, 0x74, 0xbb, 0x87, 0x37, 0x60, 0x97,
	0xdb, 0x8d, 0xac, 0x76, 0x70, 0x34, 0x11, 0x7b, 0xb7, 0x7f, 0x7e, 0x2b, 0xd3, 0x85, 0x2c, 0xee,
	0xb1, 0x80, 0x74, 0xbe, 0x05, 0x68, 0x11, 0xee, 0xe0, 0xff, 0x60, 0x76, 0x79, 0x35, 0x7a, 0xc6,
	0x87, 0x98, 0xd0, 0xc5, 0x27, 0x9c, 0x76, 0x80, 0x3e, 0xae, 0x6e, 0x2e, 0xee, 0xaf, 0x67, 0x77,
	0x23, 0x73, 0x3a, 0xfc, 0x6c, 0xe3, 0xb4, 0xcd, 0xe2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a,
	0x57, 0x1e, 0x33, 0x8a, 0x05, 0x00, 0x00,
}
