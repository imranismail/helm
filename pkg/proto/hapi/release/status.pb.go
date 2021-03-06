// Code generated by protoc-gen-go.
// source: hapi/release/status.proto
// DO NOT EDIT!

package release

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Status_Code int32

const (
	// Status_UNKNOWN indicates that a release is in an uncertain state.
	Status_UNKNOWN Status_Code = 0
	// Status_DEPLOYED indicates that the release has been pushed to Kubernetes.
	Status_DEPLOYED Status_Code = 1
	// Status_DELETED indicates that a release has been deleted from Kubermetes.
	Status_DELETED Status_Code = 2
	// Status_SUPERSEDED indicates that this release object is outdated and a newer one exists.
	Status_SUPERSEDED Status_Code = 3
	// Status_FAILED indicates that the release was not successfully deployed.
	Status_FAILED Status_Code = 4
	// Status_DELETING indicates that a delete operation is underway.
	Status_DELETING Status_Code = 5
)

var Status_Code_name = map[int32]string{
	0: "UNKNOWN",
	1: "DEPLOYED",
	2: "DELETED",
	3: "SUPERSEDED",
	4: "FAILED",
	5: "DELETING",
}
var Status_Code_value = map[string]int32{
	"UNKNOWN":    0,
	"DEPLOYED":   1,
	"DELETED":    2,
	"SUPERSEDED": 3,
	"FAILED":     4,
	"DELETING":   5,
}

func (x Status_Code) String() string {
	return proto.EnumName(Status_Code_name, int32(x))
}
func (Status_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

// Status defines the status of a release.
type Status struct {
	Code Status_Code `protobuf:"varint,1,opt,name=code,enum=hapi.release.Status_Code" json:"code,omitempty"`
	// Cluster resources as kubectl would print them.
	Resources string `protobuf:"bytes,3,opt,name=resources" json:"resources,omitempty"`
	// Contains the rendered templates/NOTES.txt if available
	Notes string `protobuf:"bytes,4,opt,name=notes" json:"notes,omitempty"`
	// LastTestSuiteRun provides results on the last test run on a release
	LastTestSuiteRun *TestSuite `protobuf:"bytes,5,opt,name=last_test_suite_run,json=lastTestSuiteRun" json:"last_test_suite_run,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Status) GetLastTestSuiteRun() *TestSuite {
	if m != nil {
		return m.LastTestSuiteRun
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "hapi.release.Status")
	proto.RegisterEnum("hapi.release.Status_Code", Status_Code_name, Status_Code_value)
}

func init() { proto.RegisterFile("hapi/release/status.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xdf, 0x6a, 0xc2, 0x30,
	0x14, 0xc6, 0x57, 0xad, 0x3a, 0x8f, 0x22, 0x21, 0x1b, 0xac, 0xca, 0x06, 0xc5, 0xab, 0xde, 0xac,
	0x05, 0xf7, 0x04, 0xdb, 0x12, 0x87, 0xac, 0x54, 0x69, 0x2b, 0xfb, 0x73, 0x53, 0xaa, 0x9e, 0x39,
	0xa1, 0x34, 0xd2, 0x24, 0x17, 0x7b, 0x88, 0xbd, 0xf3, 0x68, 0x2b, 0x74, 0x5e, 0x7e, 0xf9, 0xfd,
	0x4e, 0xce, 0xc7, 0x81, 0xf1, 0x77, 0x7a, 0x3c, 0x78, 0x05, 0x66, 0x98, 0x4a, 0xf4, 0xa4, 0x4a,
	0x95, 0x96, 0xee, 0xb1, 0x10, 0x4a, 0xd0, 0x61, 0x89, 0xdc, 0x13, 0x9a, 0xdc, 0x9d, 0x89, 0x0a,
	0xa5, 0x4a, 0xa4, 0x3e, 0x28, 0xac, 0xe5, 0xc9, 0x78, 0x2f, 0xc4, 0x3e, 0x43, 0xaf, 0x4a, 0x1b,
	0xfd, 0xe5, 0xa5, 0xf9, 0x4f, 0x8d, 0xa6, 0xbf, 0x2d, 0xe8, 0x46, 0xd5, 0xc7, 0xf4, 0x1e, 0xcc,
	0xad, 0xd8, 0xa1, 0x65, 0xd8, 0x86, 0x33, 0x9a, 0x8d, 0xdd, 0xff, 0x1b, 0xdc, 0xda, 0x71, 0x9f,
	0xc5, 0x0e, 0xc3, 0x4a, 0xa3, 0xb7, 0xd0, 0x2f, 0x50, 0x0a, 0x5d, 0x6c, 0x51, 0x5a, 0x6d, 0xdb,
	0x70, 0xfa, 0x61, 0xf3, 0x40, 0xaf, 0xa1, 0x93, 0x0b, 0x85, 0xd2, 0x32, 0x2b, 0x52, 0x07, 0x3a,
	0x87, 0xab, 0x2c, 0x95, 0x2a, 0x69, 0x1a, 0x26, 0x85, 0xce, 0xad, 0x8e, 0x6d, 0x38, 0x83, 0xd9,
	0xcd, 0xf9, 0xc6, 0x18, 0xa5, 0x8a, 0x4a, 0x25, 0x24, 0xe5, 0x4c, 0x13, 0x75, 0x3e, 0x7d, 0x07,
	0xb3, 0x6c, 0x42, 0x07, 0xd0, 0x5b, 0x07, 0xaf, 0xc1, 0xf2, 0x2d, 0x20, 0x17, 0x74, 0x08, 0x97,
	0x8c, 0xaf, 0xfc, 0xe5, 0x07, 0x67, 0xc4, 0x28, 0x11, 0xe3, 0x3e, 0x8f, 0x39, 0x23, 0x2d, 0x3a,
	0x02, 0x88, 0xd6, 0x2b, 0x1e, 0x46, 0x9c, 0x71, 0x46, 0xda, 0x14, 0xa0, 0x3b, 0x7f, 0x5c, 0xf8,
	0x9c, 0x11, 0xb3, 0x1e, 0xf3, 0x79, 0xbc, 0x08, 0x5e, 0x48, 0xe7, 0xa9, 0xff, 0xd9, 0x3b, 0x15,
	0xd8, 0x74, 0xab, 0x0b, 0x3d, 0xfc, 0x05, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x11, 0x21, 0x30, 0x86,
	0x01, 0x00, 0x00,
}
