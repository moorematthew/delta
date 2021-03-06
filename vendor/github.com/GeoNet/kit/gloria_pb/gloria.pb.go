// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gloria.proto

/*
Package gloria_pb is a generated protocol buffer package.

It is generated from these files:
	gloria.proto

It has these top-level messages:
	Marks
	Mark
	Point
	Span
	Receiver
	Firmware
	Antenna
	Radome
	Offset
	InstalledAntenna
	DeployedReceiver
	InstalledRadome
	Download
	Distribution
*/
package gloria_pb

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

// All GNSS Marks.  Use Mark.Code for the map key.
type Marks struct {
	Marks map[string]*Mark `protobuf:"bytes,1,rep,name=marks" json:"marks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Marks) Reset()                    { *m = Marks{} }
func (m *Marks) String() string            { return proto.CompactTextString(m) }
func (*Marks) ProtoMessage()               {}
func (*Marks) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Marks) GetMarks() map[string]*Mark {
	if m != nil {
		return m.Marks
	}
	return nil
}

// A GNSS Mark.
type Mark struct {
	// Code used to uniquely identify the GNSS Mark e.g. TAUP
	// teqc param: -O.mo
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	// ITRF DOMES number e.g., 50217M001
	// teqc param: -O.mn
	DomesNumber      string              `protobuf:"bytes,2,opt,name=domes_number,json=domesNumber" json:"domes_number,omitempty"`
	Point            *Point              `protobuf:"bytes,3,opt,name=point" json:"point,omitempty"`
	DeployedReceiver []*DeployedReceiver `protobuf:"bytes,4,rep,name=deployed_receiver,json=deployedReceiver" json:"deployed_receiver,omitempty"`
	InstalledAntenna []*InstalledAntenna `protobuf:"bytes,5,rep,name=installed_antenna,json=installedAntenna" json:"installed_antenna,omitempty"`
	InstalledRadome  []*InstalledRadome  `protobuf:"bytes,6,rep,name=installed_radome,json=installedRadome" json:"installed_radome,omitempty"`
	Comment          string              `protobuf:"bytes,7,opt,name=comment" json:"comment,omitempty"`
	Download         *Download           `protobuf:"bytes,8,opt,name=download" json:"download,omitempty"`
	Distribution     *Distribution       `protobuf:"bytes,9,opt,name=distribution" json:"distribution,omitempty"`
	Span             *Span               `protobuf:"bytes,10,opt,name=span" json:"span,omitempty"`
}

func (m *Mark) Reset()                    { *m = Mark{} }
func (m *Mark) String() string            { return proto.CompactTextString(m) }
func (*Mark) ProtoMessage()               {}
func (*Mark) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Mark) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Mark) GetDomesNumber() string {
	if m != nil {
		return m.DomesNumber
	}
	return ""
}

func (m *Mark) GetPoint() *Point {
	if m != nil {
		return m.Point
	}
	return nil
}

func (m *Mark) GetDeployedReceiver() []*DeployedReceiver {
	if m != nil {
		return m.DeployedReceiver
	}
	return nil
}

func (m *Mark) GetInstalledAntenna() []*InstalledAntenna {
	if m != nil {
		return m.InstalledAntenna
	}
	return nil
}

func (m *Mark) GetInstalledRadome() []*InstalledRadome {
	if m != nil {
		return m.InstalledRadome
	}
	return nil
}

func (m *Mark) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Mark) GetDownload() *Download {
	if m != nil {
		return m.Download
	}
	return nil
}

func (m *Mark) GetDistribution() *Distribution {
	if m != nil {
		return m.Distribution
	}
	return nil
}

func (m *Mark) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

// A geographical point on NZGD2000
type Point struct {
	// Latitude - geographical latitude of the point for the given datum.
	// teqc param: -O.pg[1]
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	// Longitude - geographical longitude of the point for the given datum.
	// teqc param: -O.pg[2]
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
	// Elevation - Height in meters of the point for the given datum.
	// teqc param: -O.pg[3]
	Elevation float64 `protobuf:"fixed64,3,opt,name=elevation" json:"elevation,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Point) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Point) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Point) GetElevation() float64 {
	if m != nil {
		return m.Elevation
	}
	return 0
}

// A time span that has a start and and end.
type Span struct {
	// Start - time in Unix seconds.
	Start int64 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	// End - time in Unix seconds.  A future date of 9999-01-01T00:00:00Z is used to indicate still open.
	End int64 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
}

func (m *Span) Reset()                    { *m = Span{} }
func (m *Span) String() string            { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()               {}
func (*Span) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Span) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Span) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

// A GNSS receiver.
type Receiver struct {
	// The receiver model e.g., TRIMBLE NETR9
	// teqc param: -O.rt
	Model string `protobuf:"bytes,1,opt,name=model" json:"model,omitempty"`
	// The receiver serial number e.g., 5033K69574
	// teqc param: -O.rn
	SerialNumber string      `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
	Firmware     []*Firmware `protobuf:"bytes,3,rep,name=firmware" json:"firmware,omitempty"`
}

func (m *Receiver) Reset()                    { *m = Receiver{} }
func (m *Receiver) String() string            { return proto.CompactTextString(m) }
func (*Receiver) ProtoMessage()               {}
func (*Receiver) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Receiver) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Receiver) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *Receiver) GetFirmware() []*Firmware {
	if m != nil {
		return m.Firmware
	}
	return nil
}

// Firmware versions
type Firmware struct {
	// The firmware version e.g., 5.15
	// teqc param: -O.rv
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Span    *Span  `protobuf:"bytes,2,opt,name=span" json:"span,omitempty"`
}

func (m *Firmware) Reset()                    { *m = Firmware{} }
func (m *Firmware) String() string            { return proto.CompactTextString(m) }
func (*Firmware) ProtoMessage()               {}
func (*Firmware) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Firmware) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Firmware) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

// A GNSS antenna
type Antenna struct {
	// The antenna model TRM57971.00
	// teqc param: -O.at[1]
	Model string `protobuf:"bytes,1,opt,name=model" json:"model,omitempty"`
	// The antenna serial number e.g., 1441031450
	// teqc param: -O.an
	SerialNumber string `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
}

func (m *Antenna) Reset()                    { *m = Antenna{} }
func (m *Antenna) String() string            { return proto.CompactTextString(m) }
func (*Antenna) ProtoMessage()               {}
func (*Antenna) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Antenna) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Antenna) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

// An antenna radome
type Radome struct {
	// The radome model
	// teqc param: -O.at[2]
	Model string `protobuf:"bytes,1,opt,name=model" json:"model,omitempty"`
}

func (m *Radome) Reset()                    { *m = Radome{} }
func (m *Radome) String() string            { return proto.CompactTextString(m) }
func (*Radome) ProtoMessage()               {}
func (*Radome) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Radome) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

// Offset for an installed antenna
type Offset struct {
	// The vertical offset in m e.g., 0.0550
	// teqc param: -O.pe[1]
	Vertical float64 `protobuf:"fixed64,1,opt,name=vertical" json:"vertical,omitempty"`
	// The offset north in m e.g., 0.0
	// teqc param: -O.pe[2]
	North float64 `protobuf:"fixed64,2,opt,name=north" json:"north,omitempty"`
	// The offset east in m e.g., 0.0
	// teqc param: -O.pe[3]
	East float64 `protobuf:"fixed64,3,opt,name=east" json:"east,omitempty"`
}

func (m *Offset) Reset()                    { *m = Offset{} }
func (m *Offset) String() string            { return proto.CompactTextString(m) }
func (*Offset) ProtoMessage()               {}
func (*Offset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Offset) GetVertical() float64 {
	if m != nil {
		return m.Vertical
	}
	return 0
}

func (m *Offset) GetNorth() float64 {
	if m != nil {
		return m.North
	}
	return 0
}

func (m *Offset) GetEast() float64 {
	if m != nil {
		return m.East
	}
	return 0
}

type InstalledAntenna struct {
	Antenna *Antenna `protobuf:"bytes,1,opt,name=antenna" json:"antenna,omitempty"`
	Offset  *Offset  `protobuf:"bytes,2,opt,name=offset" json:"offset,omitempty"`
	Span    *Span    `protobuf:"bytes,3,opt,name=span" json:"span,omitempty"`
}

func (m *InstalledAntenna) Reset()                    { *m = InstalledAntenna{} }
func (m *InstalledAntenna) String() string            { return proto.CompactTextString(m) }
func (*InstalledAntenna) ProtoMessage()               {}
func (*InstalledAntenna) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *InstalledAntenna) GetAntenna() *Antenna {
	if m != nil {
		return m.Antenna
	}
	return nil
}

func (m *InstalledAntenna) GetOffset() *Offset {
	if m != nil {
		return m.Offset
	}
	return nil
}

func (m *InstalledAntenna) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

type DeployedReceiver struct {
	Receiver *Receiver `protobuf:"bytes,1,opt,name=receiver" json:"receiver,omitempty"`
	Span     *Span     `protobuf:"bytes,2,opt,name=span" json:"span,omitempty"`
}

func (m *DeployedReceiver) Reset()                    { *m = DeployedReceiver{} }
func (m *DeployedReceiver) String() string            { return proto.CompactTextString(m) }
func (*DeployedReceiver) ProtoMessage()               {}
func (*DeployedReceiver) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeployedReceiver) GetReceiver() *Receiver {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func (m *DeployedReceiver) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

type InstalledRadome struct {
	Radome *Radome `protobuf:"bytes,1,opt,name=radome" json:"radome,omitempty"`
	Span   *Span   `protobuf:"bytes,2,opt,name=span" json:"span,omitempty"`
}

func (m *InstalledRadome) Reset()                    { *m = InstalledRadome{} }
func (m *InstalledRadome) String() string            { return proto.CompactTextString(m) }
func (*InstalledRadome) ProtoMessage()               {}
func (*InstalledRadome) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *InstalledRadome) GetRadome() *Radome {
	if m != nil {
		return m.Radome
	}
	return nil
}

func (m *InstalledRadome) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

type Download struct {
	// The download priority, 0 for lowest, higher numbers for higher priority.
	Priority int64 `protobuf:"varint,1,opt,name=priority" json:"priority,omitempty"`
	// Download rate limit in KBytes/sec. 0 = no limit.
	Rate int64 `protobuf:"varint,2,opt,name=rate" json:"rate,omitempty"`
}

func (m *Download) Reset()                    { *m = Download{} }
func (m *Download) String() string            { return proto.CompactTextString(m) }
func (*Download) ProtoMessage()               {}
func (*Download) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Download) GetPriority() int64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Download) GetRate() int64 {
	if m != nil {
		return m.Rate
	}
	return 0
}

type Distribution struct {
	// Set true if data should be sent to the IGS
	Igs bool `protobuf:"varint,1,opt,name=igs" json:"igs,omitempty"`
	// Set true if this site belongs to LINZ
	Linz bool `protobuf:"varint,2,opt,name=linz" json:"linz,omitempty"`
}

func (m *Distribution) Reset()                    { *m = Distribution{} }
func (m *Distribution) String() string            { return proto.CompactTextString(m) }
func (*Distribution) ProtoMessage()               {}
func (*Distribution) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Distribution) GetIgs() bool {
	if m != nil {
		return m.Igs
	}
	return false
}

func (m *Distribution) GetLinz() bool {
	if m != nil {
		return m.Linz
	}
	return false
}

func init() {
	proto.RegisterType((*Marks)(nil), "gloria.Marks")
	proto.RegisterType((*Mark)(nil), "gloria.Mark")
	proto.RegisterType((*Point)(nil), "gloria.Point")
	proto.RegisterType((*Span)(nil), "gloria.Span")
	proto.RegisterType((*Receiver)(nil), "gloria.Receiver")
	proto.RegisterType((*Firmware)(nil), "gloria.Firmware")
	proto.RegisterType((*Antenna)(nil), "gloria.Antenna")
	proto.RegisterType((*Radome)(nil), "gloria.Radome")
	proto.RegisterType((*Offset)(nil), "gloria.Offset")
	proto.RegisterType((*InstalledAntenna)(nil), "gloria.InstalledAntenna")
	proto.RegisterType((*DeployedReceiver)(nil), "gloria.DeployedReceiver")
	proto.RegisterType((*InstalledRadome)(nil), "gloria.InstalledRadome")
	proto.RegisterType((*Download)(nil), "gloria.Download")
	proto.RegisterType((*Distribution)(nil), "gloria.Distribution")
}

func init() { proto.RegisterFile("gloria.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6b, 0x1b, 0x3f,
	0x10, 0x65, 0xbd, 0x5e, 0x7b, 0x3d, 0x76, 0x7e, 0xf1, 0x4f, 0x04, 0x2a, 0x42, 0x29, 0xee, 0x06,
	0x4a, 0x0a, 0xc1, 0x87, 0xb4, 0x87, 0x90, 0x5b, 0x43, 0x12, 0xe8, 0xa1, 0x69, 0x51, 0x6f, 0xed,
	0xc1, 0xc8, 0x5e, 0x25, 0x15, 0x59, 0x4b, 0x8b, 0x56, 0x76, 0x70, 0x2f, 0x3d, 0xf6, 0x2b, 0xf6,
	0xe3, 0x14, 0xfd, 0x4d, 0xd6, 0x6d, 0x20, 0xf4, 0x62, 0x66, 0xde, 0x1b, 0x8d, 0x34, 0xf3, 0x66,
	0xd6, 0x30, 0xba, 0xa9, 0xa4, 0xe2, 0x74, 0x5a, 0x2b, 0xa9, 0x25, 0xea, 0x39, 0xaf, 0xf8, 0x01,
	0xd9, 0x07, 0xaa, 0x6e, 0x1b, 0x34, 0x85, 0x6c, 0x69, 0x0c, 0x9c, 0x4c, 0xd2, 0xc3, 0xe1, 0x31,
	0x9e, 0xfa, 0x70, 0xcb, 0xba, 0xdf, 0x0b, 0xa1, 0xd5, 0x86, 0xb8, 0xb0, 0xfd, 0x4b, 0x80, 0x7b,
	0x10, 0x8d, 0x21, 0xbd, 0x65, 0x1b, 0x9c, 0x4c, 0x92, 0xc3, 0x01, 0x31, 0x26, 0x2a, 0x20, 0x5b,
	0xd3, 0x6a, 0xc5, 0x70, 0x67, 0x92, 0x1c, 0x0e, 0x8f, 0x47, 0x0f, 0xf3, 0x11, 0x47, 0x9d, 0x76,
	0x4e, 0x92, 0xe2, 0x57, 0x0a, 0x5d, 0x83, 0x21, 0x04, 0xdd, 0x85, 0x2c, 0x99, 0xcf, 0x61, 0x6d,
	0xf4, 0x12, 0x46, 0xa5, 0x5c, 0xb2, 0x66, 0x26, 0x56, 0xcb, 0x39, 0x53, 0x36, 0xd7, 0x80, 0x0c,
	0x2d, 0x76, 0x65, 0x21, 0x74, 0x00, 0x59, 0x2d, 0xb9, 0xd0, 0x38, 0xb5, 0xf7, 0xec, 0x84, 0x7b,
	0x3e, 0x19, 0x90, 0x38, 0x0e, 0x5d, 0xc0, 0xff, 0x25, 0xab, 0x2b, 0xb9, 0x61, 0xe5, 0x4c, 0xb1,
	0x05, 0xe3, 0x6b, 0xa6, 0x70, 0xb7, 0x5d, 0xe8, 0xb9, 0x0f, 0x20, 0x9e, 0x27, 0xe3, 0x72, 0x0b,
	0x31, 0x69, 0xb8, 0x68, 0x34, 0xad, 0x2a, 0x56, 0xce, 0xa8, 0xd0, 0x4c, 0x08, 0x8a, 0xb3, 0x76,
	0x9a, 0xf7, 0x21, 0xe0, 0x9d, 0xe3, 0xc9, 0x98, 0x6f, 0x21, 0xe8, 0x0c, 0xee, 0xb1, 0x99, 0xa2,
	0xa6, 0x1a, 0xdc, 0xb3, 0x59, 0x9e, 0xfd, 0x91, 0x85, 0x58, 0x9a, 0xec, 0xf2, 0x36, 0x80, 0x30,
	0xf4, 0x17, 0x72, 0xb9, 0x64, 0x42, 0xe3, 0xbe, 0x6d, 0x4a, 0x70, 0xd1, 0x11, 0xe4, 0xa5, 0xbc,
	0x13, 0x95, 0xa4, 0x25, 0xce, 0x6d, 0x4f, 0xc6, 0xb1, 0x44, 0x8f, 0x93, 0x18, 0x81, 0x4e, 0x60,
	0x54, 0xf2, 0x46, 0x2b, 0x3e, 0x5f, 0x69, 0x2e, 0x05, 0x1e, 0xd8, 0x13, 0x7b, 0xf1, 0xc4, 0x03,
	0x8e, 0xb4, 0x22, 0xd1, 0x04, 0xba, 0x4d, 0x4d, 0x05, 0x86, 0xb6, 0xbe, 0x9f, 0x6b, 0x2a, 0x88,
	0x65, 0x8a, 0x19, 0x64, 0x56, 0x05, 0xb4, 0x0f, 0x79, 0x45, 0x35, 0xd7, 0x2b, 0x2f, 0x6f, 0x42,
	0xa2, 0x8f, 0x9e, 0xc3, 0xa0, 0x92, 0xe2, 0xc6, 0x91, 0x1d, 0x4b, 0xde, 0x03, 0x86, 0x65, 0x15,
	0x5b, 0x53, 0xfb, 0xb6, 0xd4, 0xb1, 0x11, 0x28, 0xa6, 0xd0, 0x35, 0xd7, 0xa1, 0x3d, 0xc8, 0x1a,
	0x4d, 0x95, 0xb6, 0xc9, 0x53, 0xe2, 0x1c, 0x33, 0x93, 0x4c, 0x94, 0x36, 0x67, 0x4a, 0x8c, 0x59,
	0x34, 0x90, 0x47, 0x2d, 0xf7, 0x20, 0x5b, 0xca, 0x92, 0x55, 0x7e, 0xde, 0x9c, 0x83, 0x0e, 0x60,
	0xa7, 0x61, 0x8a, 0xd3, 0xaa, 0x3d, 0x71, 0x23, 0x07, 0xfa, 0x91, 0x3b, 0x82, 0xfc, 0x9a, 0xab,
	0xe5, 0x1d, 0x55, 0x0c, 0xa7, 0x56, 0xb7, 0xd8, 0xe1, 0x4b, 0x8f, 0x93, 0x18, 0x51, 0x5c, 0x42,
	0x1e, 0x50, 0xa3, 0xda, 0x9a, 0xa9, 0xc6, 0x14, 0xe3, 0xae, 0x0d, 0x6e, 0xec, 0x66, 0xe7, 0xd1,
	0x6e, 0x9e, 0x43, 0x3f, 0x0c, 0xd0, 0xbf, 0xbf, 0xbd, 0x78, 0x01, 0x3d, 0x3f, 0x41, 0x7f, 0x4d,
	0x52, 0x5c, 0x41, 0xef, 0xe3, 0xf5, 0x75, 0xc3, 0xac, 0x68, 0x6b, 0xa6, 0x34, 0x5f, 0xd0, 0x2a,
	0x88, 0x16, 0x7c, 0x73, 0x56, 0x48, 0xa5, 0xbf, 0x79, 0xc1, 0x9c, 0x63, 0x36, 0x98, 0xd1, 0x46,
	0x7b, 0x9d, 0xac, 0x5d, 0xfc, 0x4c, 0x60, 0xbc, 0xbd, 0x12, 0xe8, 0x35, 0xf4, 0xc3, 0xf6, 0x24,
	0xb6, 0xde, 0xdd, 0x50, 0x6f, 0x58, 0x9a, 0xc0, 0xa3, 0x57, 0xd0, 0x93, 0xf6, 0x3d, 0xbe, 0x33,
	0xff, 0x85, 0x48, 0xf7, 0x4a, 0xe2, 0xd9, 0xd8, 0xbf, 0xf4, 0xd1, 0xfe, 0xcd, 0x61, 0xbc, 0xbd,
	0xe2, 0x46, 0xc9, 0xf8, 0x39, 0x48, 0xda, 0xbb, 0x12, 0x3f, 0x03, 0x31, 0xe2, 0x09, 0x1a, 0x7d,
	0x85, 0xdd, 0xad, 0xcd, 0x35, 0x05, 0xf8, 0x15, 0x4f, 0xda, 0x05, 0xf8, 0xcd, 0xf6, 0xec, 0x13,
	0x92, 0x9f, 0x42, 0x1e, 0x16, 0xd8, 0x88, 0x53, 0x2b, 0x2e, 0x15, 0xd7, 0x1b, 0x3f, 0xf4, 0xd1,
	0x37, 0x32, 0x28, 0xaa, 0x99, 0x1f, 0x7c, 0x6b, 0x17, 0x6f, 0x61, 0xf4, 0x70, 0x95, 0xcd, 0x6e,
	0xf0, 0x9b, 0xc6, 0x1e, 0xcd, 0x89, 0x31, 0xcd, 0xa9, 0x8a, 0x8b, 0xef, 0xf6, 0x54, 0x4e, 0xac,
	0x7d, 0x36, 0xfc, 0x32, 0x70, 0xcf, 0x98, 0xd5, 0xf3, 0x79, 0xcf, 0xfe, 0x71, 0xbc, 0xf9, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x51, 0xf3, 0x2e, 0x8b, 0x48, 0x06, 0x00, 0x00,
}
