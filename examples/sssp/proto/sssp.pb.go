// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sssp.proto

package sssp

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SSSPMessage struct {
	FromVertexId string `protobuf:"bytes,1,opt,name=from_vertex_id,json=fromVertexId,proto3" json:"from_vertex_id,omitempty"`
	Value        uint32 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *SSSPMessage) Reset()      { *m = SSSPMessage{} }
func (*SSSPMessage) ProtoMessage() {}
func (*SSSPMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d6ce279ff146973, []int{0}
}
func (m *SSSPMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SSSPMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SSSPMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SSSPMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSSPMessage.Merge(m, src)
}
func (m *SSSPMessage) XXX_Size() int {
	return m.Size()
}
func (m *SSSPMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SSSPMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SSSPMessage proto.InternalMessageInfo

func (m *SSSPMessage) GetFromVertexId() string {
	if m != nil {
		return m.FromVertexId
	}
	return ""
}

func (m *SSSPMessage) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*SSSPMessage)(nil), "SSSPMessage")
}

func init() { proto.RegisterFile("sssp.proto", fileDescriptor_7d6ce279ff146973) }

var fileDescriptor_7d6ce279ff146973 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2e, 0x2e, 0x2e,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xf2, 0xe4, 0xe2, 0x0e, 0x0e, 0x0e, 0x0e, 0xf0, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x52, 0xe1, 0xe2, 0x4b, 0x2b, 0xca, 0xcf, 0x8d, 0x2f, 0x4b,
	0x2d, 0x2a, 0x49, 0xad, 0x88, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x01,
	0x89, 0x86, 0x81, 0x05, 0x3d, 0x53, 0x84, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25,
	0x98, 0x14, 0x18, 0x35, 0x78, 0x83, 0x20, 0x1c, 0x27, 0x93, 0x0b, 0x0f, 0xe5, 0x18, 0x6e, 0x3c,
	0x94, 0x63, 0xf8, 0xf0, 0x50, 0x8e, 0xb1, 0xe1, 0x91, 0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x27,
	0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8b, 0x47, 0x72, 0x0c,
	0x1f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72,
	0x0c, 0x49, 0x6c, 0x60, 0x77, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x20, 0x95, 0xb5, 0x55,
	0x95, 0x00, 0x00, 0x00,
}

func (this *SSSPMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SSSPMessage)
	if !ok {
		that2, ok := that.(SSSPMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.FromVertexId != that1.FromVertexId {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *SSSPMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&sssp.SSSPMessage{")
	s = append(s, "FromVertexId: "+fmt.Sprintf("%#v", this.FromVertexId)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSssp(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SSSPMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SSSPMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FromVertexId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSssp(dAtA, i, uint64(len(m.FromVertexId)))
		i += copy(dAtA[i:], m.FromVertexId)
	}
	if m.Value != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSssp(dAtA, i, uint64(m.Value))
	}
	return i, nil
}

func encodeVarintSssp(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SSSPMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromVertexId)
	if l > 0 {
		n += 1 + l + sovSssp(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovSssp(uint64(m.Value))
	}
	return n
}

func sovSssp(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSssp(x uint64) (n int) {
	return sovSssp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SSSPMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SSSPMessage{`,
		`FromVertexId:` + fmt.Sprintf("%v", this.FromVertexId) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSssp(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SSSPMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSssp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SSSPMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SSSPMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromVertexId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSssp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSssp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSssp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromVertexId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSssp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSssp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSssp
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSssp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSssp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSssp
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSssp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSssp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSssp
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSssp
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSssp
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSssp(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSssp
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSssp = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSssp   = fmt.Errorf("proto: integer overflow")
)
