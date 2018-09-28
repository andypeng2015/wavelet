// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tree/node.proto

package iavl

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Tree struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Height               uint32   `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Size_                uint64   `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Left                 *Tree    `protobuf:"bytes,5,opt,name=left" json:"left,omitempty"`
	Right                *Tree    `protobuf:"bytes,6,opt,name=right" json:"right,omitempty"`
	StorageKey           []byte   `protobuf:"bytes,7,opt,name=storage_key,json=storageKey,proto3" json:"storage_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tree) Reset()         { *m = Tree{} }
func (m *Tree) String() string { return proto.CompactTextString(m) }
func (*Tree) ProtoMessage()    {}
func (*Tree) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_e2b386da01505e9b, []int{0}
}
func (m *Tree) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tree.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Tree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tree.Merge(dst, src)
}
func (m *Tree) XXX_Size() int {
	return m.Size()
}
func (m *Tree) XXX_DiscardUnknown() {
	xxx_messageInfo_Tree.DiscardUnknown(m)
}

var xxx_messageInfo_Tree proto.InternalMessageInfo

func (m *Tree) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Tree) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Tree) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Tree) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *Tree) GetLeft() *Tree {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *Tree) GetRight() *Tree {
	if m != nil {
		return m.Right
	}
	return nil
}

func (m *Tree) GetStorageKey() []byte {
	if m != nil {
		return m.StorageKey
	}
	return nil
}

func init() {
	proto.RegisterType((*Tree)(nil), "tree.Tree")
}
func (m *Tree) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tree) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNode(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNode(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if m.Height != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.Height))
	}
	if m.Size_ != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.Size_))
	}
	if m.Left != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.Left.Size()))
		n1, err := m.Left.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Right != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.Right.Size()))
		n2, err := m.Right.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.StorageKey) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintNode(dAtA, i, uint64(len(m.StorageKey)))
		i += copy(dAtA[i:], m.StorageKey)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintNode(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Tree) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovNode(uint64(m.Height))
	}
	if m.Size_ != 0 {
		n += 1 + sovNode(uint64(m.Size_))
	}
	if m.Left != nil {
		l = m.Left.Size()
		n += 1 + l + sovNode(uint64(l))
	}
	if m.Right != nil {
		l = m.Right.Size()
		n += 1 + l + sovNode(uint64(l))
	}
	l = len(m.StorageKey)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNode(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNode(x uint64) (n int) {
	return sovNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Tree) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Tree: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tree: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Left", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Left == nil {
				m.Left = &Tree{}
			}
			if err := m.Left.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Right", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Right == nil {
				m.Right = &Tree{}
			}
			if err := m.Right.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StorageKey = append(m.StorageKey[:0], dAtA[iNdEx:postIndex]...)
			if m.StorageKey == nil {
				m.StorageKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthNode
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNode
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
				next, err := skipNode(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthNode = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNode   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("tree/node.proto", fileDescriptor_node_e2b386da01505e9b) }

var fileDescriptor_node_e2b386da01505e9b = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x29, 0x4a, 0x4d,
	0xd5, 0xcf, 0xcb, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x09, 0x28,
	0xed, 0x65, 0xe4, 0x62, 0x09, 0x29, 0x4a, 0x4d, 0x15, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a,
	0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x82, 0x20, 0x1c, 0x21, 0x31, 0x2e, 0xb6, 0x8c, 0xd4,
	0xcc, 0xf4, 0x8c, 0x12, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xde, 0x20, 0x28, 0x4f, 0x48, 0x88, 0x8b,
	0xa5, 0x38, 0xb3, 0x2a, 0x55, 0x82, 0x45, 0x81, 0x51, 0x83, 0x25, 0x08, 0xcc, 0x16, 0x92, 0xe3,
	0x62, 0xc9, 0x49, 0x4d, 0x2b, 0x91, 0x60, 0x55, 0x60, 0xd4, 0xe0, 0x36, 0xe2, 0xd2, 0x03, 0xd9,
	0xa8, 0x07, 0xb2, 0x2d, 0x08, 0x2c, 0x2e, 0xa4, 0xc0, 0xc5, 0x5a, 0x04, 0x36, 0x8a, 0x0d, 0x43,
	0x01, 0x44, 0x42, 0x48, 0x9e, 0x8b, 0xbb, 0xb8, 0x24, 0xbf, 0x28, 0x31, 0x3d, 0x35, 0x1e, 0xe4,
	0x3a, 0x76, 0xb0, 0x4b, 0xb8, 0xa0, 0x42, 0xde, 0xa9, 0x95, 0x4e, 0x3c, 0x27, 0x1e, 0xc9, 0x31,
	0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x63, 0x12, 0x1b, 0xd8, 0x6b, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x7c, 0xd1, 0x86, 0x8f, 0xed, 0x00, 0x00, 0x00,
}
