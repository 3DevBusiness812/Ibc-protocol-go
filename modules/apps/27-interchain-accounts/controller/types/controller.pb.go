// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/interchain_accounts/controller/v1/controller.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	// controller_enabled enables or disables the controller submodule.
	ControllerEnabled bool `protobuf:"varint,1,opt,name=controller_enabled,json=controllerEnabled,proto3" json:"controller_enabled,omitempty" yaml:"controller_enabled"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_177fd0fec5eb3400, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetControllerEnabled() bool {
	if m != nil {
		return m.ControllerEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*Params)(nil), "ibc.applications.interchain_accounts.controller.v1.Params")
}

func init() {
	proto.RegisterFile("ibc/applications/interchain_accounts/controller/v1/controller.proto", fileDescriptor_177fd0fec5eb3400)
}

var fileDescriptor_177fd0fec5eb3400 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0xce, 0x4c, 0x4a, 0xd6,
	0x4f, 0x2c, 0x28, 0xc8, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0xcf, 0xcc, 0x2b,
	0x49, 0x2d, 0x4a, 0xce, 0x48, 0xcc, 0xcc, 0x8b, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x29,
	0xd6, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0xca, 0xcf, 0xc9, 0x49, 0x2d, 0xd2, 0x2f, 0x33, 0x44, 0xe2,
	0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x19, 0x65, 0x26, 0x25, 0xeb, 0x21, 0x1b, 0xa2, 0x87,
	0xc5, 0x10, 0x3d, 0x24, 0x6d, 0x65, 0x86, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xed, 0xfa,
	0x20, 0x16, 0xc4, 0x24, 0xa5, 0x30, 0x2e, 0xb6, 0x80, 0xc4, 0xa2, 0xc4, 0xdc, 0x62, 0x21, 0x1f,
	0x2e, 0x21, 0x84, 0x86, 0xf8, 0xd4, 0xbc, 0xc4, 0xa4, 0x9c, 0xd4, 0x14, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0x0e, 0x27, 0xd9, 0x4f, 0xf7, 0xe4, 0x25, 0x2b, 0x13, 0x73, 0x73, 0xac, 0x94, 0x30, 0xd5,
	0x28, 0x05, 0x09, 0x22, 0x04, 0x5d, 0x21, 0x62, 0x4e, 0x59, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78,
	0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc,
	0x78, 0x2c, 0xc7, 0x10, 0x15, 0x90, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab,
	0x9f, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0xac, 0x9f, 0x99, 0x94, 0xac, 0x9b, 0x9e, 0xaf, 0x5f, 0x66,
	0xac, 0x9f, 0x9b, 0x9f, 0x52, 0x9a, 0x93, 0x5a, 0x0c, 0x0a, 0xa0, 0x62, 0x7d, 0x23, 0x73, 0x5d,
	0x84, 0xb7, 0x74, 0xb1, 0x85, 0x4d, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x2b, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x79, 0xfc, 0x19, 0x5b, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ControllerEnabled {
		i--
		if m.ControllerEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintController(dAtA []byte, offset int, v uint64) int {
	offset -= sovController(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ControllerEnabled {
		n += 2
	}
	return n
}

func sovController(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozController(x uint64) (n int) {
	return sovController(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowController
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ControllerEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowController
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ControllerEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipController(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthController
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
func skipController(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowController
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
					return 0, ErrIntOverflowController
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowController
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
				return 0, ErrInvalidLengthController
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupController
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthController
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthController        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowController          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupController = fmt.Errorf("proto: unexpected end of group")
)
