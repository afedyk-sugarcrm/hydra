// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apis/discovery/v1/discovery.proto

/*
Package discovery is a generated protocol buffer package.

It is generated from these files:
	apis/discovery/v1/discovery.proto

It has these top-level messages:
	Kinds
	Kind
	Regions
	Region
	Services
	Service
	ServiceEndpoint
*/
package discovery

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Kinds struct {
	Partition string  `protobuf:"bytes,1,opt,name=partition" json:"partition,omitempty"`
	Kinds     []*Kind `protobuf:"bytes,2,rep,name=kinds" json:"kinds,omitempty"`
}

func (m *Kinds) Reset()                    { *m = Kinds{} }
func (m *Kinds) String() string            { return proto.CompactTextString(m) }
func (*Kinds) ProtoMessage()               {}
func (*Kinds) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Kinds) GetPartition() string {
	if m != nil {
		return m.Partition
	}
	return ""
}

func (m *Kinds) GetKinds() []*Kind {
	if m != nil {
		return m.Kinds
	}
	return nil
}

type Kind struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Kind) Reset()                    { *m = Kind{} }
func (m *Kind) String() string            { return proto.CompactTextString(m) }
func (*Kind) ProtoMessage()               {}
func (*Kind) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Kind) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Kind) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Regions struct {
	Regions []*Region `protobuf:"bytes,1,rep,name=regions" json:"regions,omitempty"`
}

func (m *Regions) Reset()                    { *m = Regions{} }
func (m *Regions) String() string            { return proto.CompactTextString(m) }
func (*Regions) ProtoMessage()               {}
func (*Regions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Regions) GetRegions() []*Region {
	if m != nil {
		return m.Regions
	}
	return nil
}

type Region struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Region) Reset()                    { *m = Region{} }
func (m *Region) String() string            { return proto.CompactTextString(m) }
func (*Region) ProtoMessage()               {}
func (*Region) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Region) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Region) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Services struct {
	Services []*Service `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
}

func (m *Services) Reset()                    { *m = Services{} }
func (m *Services) String() string            { return proto.CompactTextString(m) }
func (*Services) ProtoMessage()               {}
func (*Services) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Services) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type Service struct {
	Name        string             `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string             `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Type        string             `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Endpoints   []*ServiceEndpoint `protobuf:"bytes,4,rep,name=endpoints" json:"endpoints,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Service) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Service) GetEndpoints() []*ServiceEndpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type ServiceEndpoint struct {
	Url    string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Region string `protobuf:"bytes,2,opt,name=region" json:"region,omitempty"`
}

func (m *ServiceEndpoint) Reset()                    { *m = ServiceEndpoint{} }
func (m *ServiceEndpoint) String() string            { return proto.CompactTextString(m) }
func (*ServiceEndpoint) ProtoMessage()               {}
func (*ServiceEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ServiceEndpoint) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ServiceEndpoint) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func init() {
	proto.RegisterType((*Kinds)(nil), "sugarcrm.apis.discovery.v1.Kinds")
	proto.RegisterType((*Kind)(nil), "sugarcrm.apis.discovery.v1.Kind")
	proto.RegisterType((*Regions)(nil), "sugarcrm.apis.discovery.v1.Regions")
	proto.RegisterType((*Region)(nil), "sugarcrm.apis.discovery.v1.Region")
	proto.RegisterType((*Services)(nil), "sugarcrm.apis.discovery.v1.Services")
	proto.RegisterType((*Service)(nil), "sugarcrm.apis.discovery.v1.Service")
	proto.RegisterType((*ServiceEndpoint)(nil), "sugarcrm.apis.discovery.v1.ServiceEndpoint")
}

func init() { proto.RegisterFile("apis/discovery/v1/discovery.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0xc7, 0xe9, 0xde, 0xf7, 0xec, 0xa0, 0xe4, 0x20, 0x45, 0x3c, 0xd4, 0x78, 0x19, 0x08, 0x2d,
	0x53, 0xf0, 0xe0, 0x86, 0x82, 0x20, 0x22, 0xbb, 0xd5, 0x9b, 0xe0, 0xa1, 0x6b, 0x43, 0x0d, 0xae,
	0x49, 0x48, 0xd2, 0xc2, 0xbe, 0x8c, 0x9f, 0x55, 0xf2, 0xb2, 0x55, 0x14, 0x87, 0xec, 0xf6, 0x7f,
	0x5e, 0x7e, 0x4f, 0x7f, 0x94, 0xc0, 0x79, 0x26, 0xa8, 0x4a, 0x0a, 0xaa, 0x72, 0xde, 0x10, 0xb9,
	0x49, 0x9a, 0x59, 0x5b, 0xc4, 0x42, 0x72, 0xcd, 0xd1, 0xa9, 0xaa, 0xcb, 0x4c, 0xe6, 0xb2, 0x8a,
	0xcd, 0x6e, 0xdc, 0x8e, 0x9b, 0x19, 0x7e, 0x83, 0xfe, 0x92, 0xb2, 0x42, 0xa1, 0x33, 0x18, 0x8b,
	0x4c, 0x6a, 0xaa, 0x29, 0x67, 0x61, 0x10, 0x05, 0xd3, 0x71, 0xda, 0x36, 0xd0, 0x0d, 0xf4, 0x3f,
	0xcc, 0x5a, 0xd8, 0x89, 0xba, 0xd3, 0xc9, 0x55, 0x14, 0xff, 0x7d, 0x32, 0x36, 0xf7, 0x52, 0xb7,
	0x8e, 0x17, 0xd0, 0x33, 0x25, 0x42, 0xd0, 0x63, 0x59, 0x45, 0xfc, 0x61, 0x9b, 0x51, 0x04, 0x93,
	0x82, 0xa8, 0x5c, 0x52, 0x61, 0xbf, 0xd9, 0xb1, 0xa3, 0xef, 0x2d, 0xfc, 0x04, 0xc3, 0x94, 0x94,
	0x94, 0x33, 0x85, 0x16, 0x30, 0x94, 0x2e, 0x86, 0x81, 0x55, 0xc0, 0xfb, 0x14, 0x1c, 0x95, 0x6e,
	0x11, 0x7c, 0x07, 0x03, 0xd7, 0x3a, 0x50, 0x64, 0x09, 0xa3, 0x17, 0x22, 0x1b, 0x9a, 0x13, 0x85,
	0xee, 0x61, 0xa4, 0x7c, 0xf6, 0x2a, 0x17, 0xfb, 0x54, 0x3c, 0x97, 0xee, 0x20, 0xfc, 0x19, 0xc0,
	0xd0, 0x77, 0x0f, 0xd3, 0x31, 0x94, 0xde, 0x08, 0x12, 0x76, 0x1d, 0x65, 0x32, 0x7a, 0x86, 0x31,
	0x61, 0x85, 0xe0, 0x94, 0x69, 0x15, 0xf6, 0xac, 0xd7, 0xe5, 0x3f, 0xbc, 0x1e, 0x3d, 0x93, 0xb6,
	0x34, 0x9e, 0xc3, 0xd1, 0x8f, 0x29, 0x3a, 0x86, 0x6e, 0x2d, 0xd7, 0x5e, 0xd3, 0x44, 0x74, 0x02,
	0x03, 0xf7, 0x77, 0xbd, 0xa0, 0xaf, 0x1e, 0x16, 0xaf, 0xb7, 0x25, 0xd5, 0xef, 0xf5, 0x2a, 0xce,
	0x79, 0x95, 0x6c, 0x05, 0x92, 0xaa, 0x5e, 0x6b, 0xda, 0x10, 0xa9, 0x48, 0xf2, 0xeb, 0xc1, 0xce,
	0x77, 0xc5, 0x6a, 0x60, 0x5f, 0xec, 0xf5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0xec, 0xf2,
	0xe6, 0xd6, 0x02, 0x00, 0x00,
}
