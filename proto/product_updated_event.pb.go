// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: product_updated_event.proto

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

type ProductUpdatedEvent struct {
	ProductID   int64   `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Title       string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Price       float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl    string  `protobuf:"bytes,5,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	UserID      int64   `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Version     int64   `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *ProductUpdatedEvent) Reset()         { *m = ProductUpdatedEvent{} }
func (m *ProductUpdatedEvent) String() string { return proto.CompactTextString(m) }
func (*ProductUpdatedEvent) ProtoMessage()    {}
func (*ProductUpdatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e87ff72754c0cdf, []int{0}
}
func (m *ProductUpdatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProductUpdatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProductUpdatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProductUpdatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductUpdatedEvent.Merge(m, src)
}
func (m *ProductUpdatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *ProductUpdatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductUpdatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ProductUpdatedEvent proto.InternalMessageInfo

func (m *ProductUpdatedEvent) GetProductID() int64 {
	if m != nil {
		return m.ProductID
	}
	return 0
}

func (m *ProductUpdatedEvent) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ProductUpdatedEvent) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ProductUpdatedEvent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProductUpdatedEvent) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *ProductUpdatedEvent) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *ProductUpdatedEvent) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*ProductUpdatedEvent)(nil), "proto.ProductUpdatedEvent")
}

func init() { proto.RegisterFile("product_updated_event.proto", fileDescriptor_0e87ff72754c0cdf) }

var fileDescriptor_0e87ff72754c0cdf = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x63, 0x4a, 0x12, 0xf2, 0x10, 0x8b, 0x61, 0xb0, 0x8a, 0xe4, 0x46, 0xb0, 0x64, 0x80,
	0x76, 0xe0, 0x0f, 0xaa, 0x32, 0x64, 0x43, 0x91, 0x32, 0x57, 0x4d, 0xfc, 0x08, 0x96, 0xda, 0x3a,
	0x72, 0x9c, 0x7e, 0x07, 0x9f, 0xc5, 0xd8, 0x91, 0xa9, 0x42, 0xc9, 0xca, 0x47, 0x20, 0x3b, 0x2d,
	0xea, 0x64, 0x9f, 0xfb, 0xee, 0xbb, 0xbe, 0x32, 0xdc, 0xd7, 0x5a, 0x89, 0xb6, 0x34, 0xcb, 0xb6,
	0x16, 0x2b, 0x83, 0x62, 0x89, 0x3b, 0xdc, 0x9a, 0x69, 0xad, 0x95, 0x51, 0xd4, 0x77, 0xc7, 0xf8,
	0xb9, 0x92, 0xe6, 0xa3, 0x2d, 0xa6, 0xa5, 0xda, 0xcc, 0x2a, 0x55, 0xa9, 0x99, 0x93, 0x8b, 0xf6,
	0xdd, 0x91, 0x03, 0x77, 0x1b, 0xb6, 0x1e, 0x7e, 0x09, 0xdc, 0xbe, 0x0d, 0xa9, 0xf9, 0x10, 0xfa,
	0x6a, 0x33, 0xe9, 0x13, 0xc0, 0xe9, 0x31, 0x29, 0x18, 0x89, 0x49, 0x32, 0x9a, 0xdf, 0x74, 0x87,
	0x49, 0x74, 0x34, 0xa7, 0x8b, 0x2c, 0x3a, 0x1a, 0x52, 0x41, 0xef, 0xc0, 0x37, 0xd2, 0xac, 0x91,
	0x5d, 0xc4, 0x24, 0x89, 0xb2, 0x01, 0xac, 0x5a, 0x6b, 0x59, 0x22, 0x1b, 0xc5, 0x24, 0x21, 0xd9,
	0x00, 0x34, 0x86, 0x6b, 0x81, 0x4d, 0xa9, 0x65, 0x6d, 0xa4, 0xda, 0xb2, 0x4b, 0xb7, 0x71, 0x2e,
	0xd1, 0x31, 0x5c, 0xc9, 0xcd, 0xaa, 0xc2, 0x5c, 0xaf, 0x99, 0xef, 0xc6, 0xff, 0x4c, 0x1f, 0x21,
	0x6c, 0x1b, 0xd4, 0xb6, 0x54, 0xe0, 0x4a, 0x41, 0x77, 0x98, 0x04, 0x79, 0x83, 0x3a, 0x5d, 0x64,
	0x81, 0x1d, 0xa5, 0x82, 0x32, 0x08, 0x77, 0xa8, 0x1b, 0x1b, 0x1f, 0x5a, 0x53, 0x76, 0xc2, 0x39,
	0xfb, 0xea, 0x38, 0xd9, 0x77, 0x9c, 0xfc, 0x74, 0x9c, 0x7c, 0xf6, 0xdc, 0xdb, 0xf7, 0xdc, 0xfb,
	0xee, 0xb9, 0x57, 0x04, 0xee, 0x3f, 0x5e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x27, 0x67,
	0xf0, 0x64, 0x01, 0x00, 0x00,
}

func (m *ProductUpdatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProductUpdatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProductUpdatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintProductUpdatedEvent(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x38
	}
	if m.UserID != 0 {
		i = encodeVarintProductUpdatedEvent(dAtA, i, uint64(m.UserID))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ImageUrl) > 0 {
		i -= len(m.ImageUrl)
		copy(dAtA[i:], m.ImageUrl)
		i = encodeVarintProductUpdatedEvent(dAtA, i, uint64(len(m.ImageUrl)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProductUpdatedEvent(dAtA, i, uint64(len(m.Description)))
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
		i = encodeVarintProductUpdatedEvent(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProductID != 0 {
		i = encodeVarintProductUpdatedEvent(dAtA, i, uint64(m.ProductID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProductUpdatedEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovProductUpdatedEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProductUpdatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProductID != 0 {
		n += 1 + sovProductUpdatedEvent(uint64(m.ProductID))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProductUpdatedEvent(uint64(l))
	}
	if m.Price != 0 {
		n += 9
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProductUpdatedEvent(uint64(l))
	}
	l = len(m.ImageUrl)
	if l > 0 {
		n += 1 + l + sovProductUpdatedEvent(uint64(l))
	}
	if m.UserID != 0 {
		n += 1 + sovProductUpdatedEvent(uint64(m.UserID))
	}
	if m.Version != 0 {
		n += 1 + sovProductUpdatedEvent(uint64(m.Version))
	}
	return n
}

func sovProductUpdatedEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProductUpdatedEvent(x uint64) (n int) {
	return sovProductUpdatedEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProductUpdatedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProductUpdatedEvent
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
			return fmt.Errorf("proto: ProductUpdatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProductUpdatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProductID", wireType)
			}
			m.ProductID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProductUpdatedEvent
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
					return ErrIntOverflowProductUpdatedEvent
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
				return ErrInvalidLengthProductUpdatedEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProductUpdatedEvent
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
					return ErrIntOverflowProductUpdatedEvent
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
				return ErrInvalidLengthProductUpdatedEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProductUpdatedEvent
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
					return ErrIntOverflowProductUpdatedEvent
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
				return ErrInvalidLengthProductUpdatedEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProductUpdatedEvent
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
					return ErrIntOverflowProductUpdatedEvent
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
					return ErrIntOverflowProductUpdatedEvent
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
			skippy, err := skipProductUpdatedEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProductUpdatedEvent
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
func skipProductUpdatedEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProductUpdatedEvent
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
					return 0, ErrIntOverflowProductUpdatedEvent
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
					return 0, ErrIntOverflowProductUpdatedEvent
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
				return 0, ErrInvalidLengthProductUpdatedEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProductUpdatedEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProductUpdatedEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProductUpdatedEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProductUpdatedEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProductUpdatedEvent = fmt.Errorf("proto: unexpected end of group")
)
