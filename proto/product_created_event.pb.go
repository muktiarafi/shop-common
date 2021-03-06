// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: product_created_event.proto

package proto

import (
	encoding_binary "encoding/binary"
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

type ProductCreatedEvent struct {
	ProductID   int64   `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Title       string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Price       float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl    string  `protobuf:"bytes,5,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	UserID      int64   `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Version     int64   `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *ProductCreatedEvent) Reset()         { *m = ProductCreatedEvent{} }
func (m *ProductCreatedEvent) String() string { return proto.CompactTextString(m) }
func (*ProductCreatedEvent) ProtoMessage()    {}
func (*ProductCreatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_d93e78b7237ea830, []int{0}
}
func (m *ProductCreatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProductCreatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProductCreatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProductCreatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductCreatedEvent.Merge(m, src)
}
func (m *ProductCreatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *ProductCreatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductCreatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ProductCreatedEvent proto.InternalMessageInfo

func (m *ProductCreatedEvent) GetProductID() int64 {
	if m != nil {
		return m.ProductID
	}
	return 0
}

func (m *ProductCreatedEvent) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ProductCreatedEvent) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ProductCreatedEvent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProductCreatedEvent) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *ProductCreatedEvent) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *ProductCreatedEvent) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*ProductCreatedEvent)(nil), "proto.ProductCreatedEvent")
}

func init() { proto.RegisterFile("product_created_event.proto", fileDescriptor_d93e78b7237ea830) }

var fileDescriptor_d93e78b7237ea830 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x6b, 0x4a, 0x12, 0xf2, 0x10, 0x8b, 0x61, 0xb0, 0x8a, 0xe4, 0x46, 0xb0, 0x64, 0x80,
	0x76, 0xe0, 0x0f, 0x4a, 0x19, 0xb2, 0xa1, 0x48, 0x9d, 0xab, 0x26, 0x7e, 0x04, 0x4b, 0x6d, 0x1d,
	0x39, 0x4e, 0xbf, 0x83, 0xcf, 0x62, 0xec, 0xc8, 0x54, 0xa1, 0x64, 0xe5, 0x23, 0x90, 0xed, 0x16,
	0x31, 0xd9, 0xe7, 0xbe, 0xfb, 0xae, 0xaf, 0x0c, 0xb7, 0xb5, 0x56, 0xa2, 0x2d, 0xcd, 0xb2, 0xd4,
	0xb8, 0x32, 0x28, 0x96, 0xb8, 0xc3, 0xad, 0x99, 0xd4, 0x5a, 0x19, 0x45, 0x03, 0x77, 0x8c, 0x1e,
	0x2b, 0x69, 0xde, 0xdb, 0x62, 0x52, 0xaa, 0xcd, 0xb4, 0x52, 0x95, 0x9a, 0x3a, 0xb9, 0x68, 0xdf,
	0x1c, 0x39, 0x70, 0x37, 0xbf, 0x75, 0xf7, 0x43, 0xe0, 0xfa, 0xd5, 0xa7, 0x3e, 0xfb, 0xd0, 0x17,
	0x9b, 0x49, 0x1f, 0x00, 0x4e, 0x8f, 0x49, 0xc1, 0x48, 0x42, 0xd2, 0xe1, 0xec, 0xaa, 0x3b, 0x8c,
	0xe3, 0xa3, 0x39, 0x9b, 0xe7, 0xf1, 0xd1, 0x90, 0x09, 0x7a, 0x03, 0x81, 0x91, 0x66, 0x8d, 0xec,
	0x2c, 0x21, 0x69, 0x9c, 0x7b, 0xb0, 0x6a, 0xad, 0x65, 0x89, 0x6c, 0x98, 0x90, 0x94, 0xe4, 0x1e,
	0x68, 0x02, 0x97, 0x02, 0x9b, 0x52, 0xcb, 0xda, 0x48, 0xb5, 0x65, 0xe7, 0x6e, 0xe3, 0xbf, 0x44,
	0x47, 0x70, 0x21, 0x37, 0xab, 0x0a, 0x17, 0x7a, 0xcd, 0x02, 0x37, 0xfe, 0x63, 0x7a, 0x0f, 0x51,
	0xdb, 0xa0, 0xb6, 0xa5, 0x42, 0x57, 0x0a, 0xba, 0xc3, 0x38, 0x5c, 0x34, 0xa8, 0xb3, 0x79, 0x1e,
	0xda, 0x51, 0x26, 0x28, 0x83, 0x68, 0x87, 0xba, 0xb1, 0xf1, 0x91, 0x35, 0xe5, 0x27, 0x9c, 0xb1,
	0xcf, 0x8e, 0x93, 0x7d, 0xc7, 0xc9, 0x77, 0xc7, 0xc9, 0x47, 0xcf, 0x07, 0xfb, 0x9e, 0x0f, 0xbe,
	0x7a, 0x3e, 0x28, 0x42, 0xf7, 0x1f, 0x4f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x09, 0x1b,
	0x6c, 0x64, 0x01, 0x00, 0x00,
}

func (m *ProductCreatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProductCreatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProductCreatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintProductCreatedEvent(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x38
	}
	if m.UserID != 0 {
		i = encodeVarintProductCreatedEvent(dAtA, i, uint64(m.UserID))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ImageUrl) > 0 {
		i -= len(m.ImageUrl)
		copy(dAtA[i:], m.ImageUrl)
		i = encodeVarintProductCreatedEvent(dAtA, i, uint64(len(m.ImageUrl)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProductCreatedEvent(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if m.Price != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Price))))
		i--
		dAtA[i] = 0x19
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProductCreatedEvent(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProductID != 0 {
		i = encodeVarintProductCreatedEvent(dAtA, i, uint64(m.ProductID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProductCreatedEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovProductCreatedEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProductCreatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProductID != 0 {
		n += 1 + sovProductCreatedEvent(uint64(m.ProductID))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProductCreatedEvent(uint64(l))
	}
	if m.Price != 0 {
		n += 9
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProductCreatedEvent(uint64(l))
	}
	l = len(m.ImageUrl)
	if l > 0 {
		n += 1 + l + sovProductCreatedEvent(uint64(l))
	}
	if m.UserID != 0 {
		n += 1 + sovProductCreatedEvent(uint64(m.UserID))
	}
	if m.Version != 0 {
		n += 1 + sovProductCreatedEvent(uint64(m.Version))
	}
	return n
}

func sovProductCreatedEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProductCreatedEvent(x uint64) (n int) {
	return sovProductCreatedEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProductCreatedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProductCreatedEvent
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
			return fmt.Errorf("proto: ProductCreatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProductCreatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProductID", wireType)
			}
			m.ProductID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProductCreatedEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProductID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProductCreatedEvent
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
				return ErrInvalidLengthProductCreatedEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProductCreatedEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Price = float64(math.Float64frombits(v))
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProductCreatedEvent
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
				return ErrInvalidLengthProductCreatedEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProductCreatedEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProductCreatedEvent
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
				return ErrInvalidLengthProductCreatedEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProductCreatedEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImageUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProductCreatedEvent
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
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProductCreatedEvent
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
			skippy, err := skipProductCreatedEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProductCreatedEvent
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
func skipProductCreatedEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProductCreatedEvent
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
					return 0, ErrIntOverflowProductCreatedEvent
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
					return 0, ErrIntOverflowProductCreatedEvent
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
				return 0, ErrInvalidLengthProductCreatedEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProductCreatedEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProductCreatedEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProductCreatedEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProductCreatedEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProductCreatedEvent = fmt.Errorf("proto: unexpected end of group")
)
