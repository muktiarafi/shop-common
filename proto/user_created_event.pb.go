// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user_created_event.proto

package proto

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type UserCreatedEvent struct {
	UserID          int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProfileImageURL string `protobuf:"bytes,3,opt,name=profile_image_url,json=profileImageUrl,proto3" json:"profile_image_url,omitempty"`
	Version         int64  `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *UserCreatedEvent) Reset()         { *m = UserCreatedEvent{} }
func (m *UserCreatedEvent) String() string { return proto.CompactTextString(m) }
func (*UserCreatedEvent) ProtoMessage()    {}
func (*UserCreatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_dedfea96a1eedd87, []int{0}
}
func (m *UserCreatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserCreatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserCreatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserCreatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreatedEvent.Merge(m, src)
}
func (m *UserCreatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *UserCreatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreatedEvent proto.InternalMessageInfo

func (m *UserCreatedEvent) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *UserCreatedEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserCreatedEvent) GetProfileImageURL() string {
	if m != nil {
		return m.ProfileImageURL
	}
	return ""
}

func (m *UserCreatedEvent) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*UserCreatedEvent)(nil), "proto.UserCreatedEvent")
}

func init() { proto.RegisterFile("user_created_event.proto", fileDescriptor_dedfea96a1eedd87) }

var fileDescriptor_dedfea96a1eedd87 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x2d, 0x4e, 0x2d,
	0x8a, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x4d, 0x89, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49,
	0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0xfa, 0x60, 0xe1, 0xa4, 0xd2, 0x34, 0x30,
	0x0f, 0xcc, 0x01, 0xb3, 0x20, 0xba, 0x94, 0x96, 0x31, 0x72, 0x09, 0x84, 0x16, 0xa7, 0x16, 0x39,
	0x43, 0x4c, 0x74, 0x05, 0x19, 0x28, 0xa4, 0xcc, 0xc5, 0x0e, 0xb6, 0x26, 0x33, 0x45, 0x82, 0x51,
	0x81, 0x51, 0x83, 0xd9, 0x89, 0xeb, 0xd1, 0x3d, 0x79, 0x36, 0x90, 0x32, 0x4f, 0x97, 0x20, 0x36,
	0x90, 0x94, 0x67, 0x8a, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x93, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x98, 0x2d, 0x64, 0xcf, 0x25, 0x58, 0x50, 0x94, 0x9f, 0x96, 0x99, 0x93, 0x1a,
	0x9f, 0x99, 0x9b, 0x98, 0x9e, 0x1a, 0x5f, 0x5a, 0x94, 0x23, 0xc1, 0x0c, 0x52, 0xe0, 0x24, 0xfc,
	0xe8, 0x9e, 0x3c, 0x7f, 0x00, 0x44, 0xd2, 0x13, 0x24, 0x17, 0x1a, 0xe4, 0x13, 0xc4, 0x5f, 0x80,
	0x2c, 0x50, 0x94, 0x23, 0x24, 0xc1, 0xc5, 0x5e, 0x96, 0x5a, 0x54, 0x9c, 0x99, 0x9f, 0x27, 0xc1,
	0x02, 0xb2, 0x39, 0x08, 0xc6, 0x75, 0x92, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x86, 0x24, 0x36, 0xb0, 0x4f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x48, 0x02, 0xd0, 0x0a,
	0x1b, 0x01, 0x00, 0x00,
}

func (m *UserCreatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserCreatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserCreatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintUserCreatedEvent(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ProfileImageURL) > 0 {
		i -= len(m.ProfileImageURL)
		copy(dAtA[i:], m.ProfileImageURL)
		i = encodeVarintUserCreatedEvent(dAtA, i, uint64(len(m.ProfileImageURL)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUserCreatedEvent(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.UserID != 0 {
		i = encodeVarintUserCreatedEvent(dAtA, i, uint64(m.UserID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintUserCreatedEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovUserCreatedEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserCreatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserID != 0 {
		n += 1 + sovUserCreatedEvent(uint64(m.UserID))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUserCreatedEvent(uint64(l))
	}
	l = len(m.ProfileImageURL)
	if l > 0 {
		n += 1 + l + sovUserCreatedEvent(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovUserCreatedEvent(uint64(m.Version))
	}
	return n
}

func sovUserCreatedEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUserCreatedEvent(x uint64) (n int) {
	return sovUserCreatedEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserCreatedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserCreatedEvent
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
			return fmt.Errorf("proto: UserCreatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserCreatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCreatedEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCreatedEvent
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
				return ErrInvalidLengthUserCreatedEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserCreatedEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileImageURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCreatedEvent
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
				return ErrInvalidLengthUserCreatedEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserCreatedEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfileImageURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCreatedEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUserCreatedEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserCreatedEvent
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
func skipUserCreatedEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUserCreatedEvent
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
					return 0, ErrIntOverflowUserCreatedEvent
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
					return 0, ErrIntOverflowUserCreatedEvent
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
				return 0, ErrInvalidLengthUserCreatedEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUserCreatedEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUserCreatedEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUserCreatedEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUserCreatedEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUserCreatedEvent = fmt.Errorf("proto: unexpected end of group")
)
