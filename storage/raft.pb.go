// Code generated by protoc-gen-gogo.
// source: cockroach/storage/raft.proto
// DO NOT EDIT!

/*
	Package storage is a generated protocol buffer package.

	It is generated from these files:
		cockroach/storage/raft.proto

	It has these top-level messages:
		RaftMessageRequest
		RaftMessageResponse
		ConfChangeContext
*/
package storage

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_roachpb "github.com/cockroachdb/cockroach/roachpb"
import raftpb "github.com/coreos/etcd/raft/raftpb"

// skipping weak import gogoproto "github.com/cockroachdb/gogoproto"

import github_com_cockroachdb_cockroach_roachpb "github.com/cockroachdb/cockroach/roachpb"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// RaftMessageRequest is the request used to send raft messages using our
// protobuf-based RPC codec.
type RaftMessageRequest struct {
	GroupID     github_com_cockroachdb_cockroach_roachpb.RangeID `protobuf:"varint,1,opt,name=group_id,json=groupId,casttype=github.com/cockroachdb/cockroach/roachpb.RangeID" json:"group_id"`
	FromReplica cockroach_roachpb.ReplicaDescriptor              `protobuf:"bytes,2,opt,name=from_replica,json=fromReplica" json:"from_replica"`
	ToReplica   cockroach_roachpb.ReplicaDescriptor              `protobuf:"bytes,3,opt,name=to_replica,json=toReplica" json:"to_replica"`
	Message     raftpb.Message                                   `protobuf:"bytes,4,opt,name=message" json:"message"`
}

func (m *RaftMessageRequest) Reset()                    { *m = RaftMessageRequest{} }
func (m *RaftMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*RaftMessageRequest) ProtoMessage()               {}
func (*RaftMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptorRaft, []int{0} }

// RaftMessageResponse is an empty message returned by raft RPCs. If a
// response is needed it will be sent as a separate message.
type RaftMessageResponse struct {
}

func (m *RaftMessageResponse) Reset()                    { *m = RaftMessageResponse{} }
func (m *RaftMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*RaftMessageResponse) ProtoMessage()               {}
func (*RaftMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptorRaft, []int{1} }

// ConfChangeContext is encoded in the raftpb.ConfChange.Context field.
type ConfChangeContext struct {
	CommandID string `protobuf:"bytes,1,opt,name=command_id,json=commandId" json:"command_id"`
	// Payload is the application-level command (i.e. an encoded
	// roachpb.EndTransactionRequest).
	Payload []byte `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	// Replica contains full details about the replica being added or removed.
	Replica cockroach_roachpb.ReplicaDescriptor `protobuf:"bytes,3,opt,name=replica" json:"replica"`
}

func (m *ConfChangeContext) Reset()                    { *m = ConfChangeContext{} }
func (m *ConfChangeContext) String() string            { return proto.CompactTextString(m) }
func (*ConfChangeContext) ProtoMessage()               {}
func (*ConfChangeContext) Descriptor() ([]byte, []int) { return fileDescriptorRaft, []int{2} }

func init() {
	proto.RegisterType((*RaftMessageRequest)(nil), "cockroach.storage.RaftMessageRequest")
	proto.RegisterType((*RaftMessageResponse)(nil), "cockroach.storage.RaftMessageResponse")
	proto.RegisterType((*ConfChangeContext)(nil), "cockroach.storage.ConfChangeContext")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for MultiRaft service

type MultiRaftClient interface {
	RaftMessage(ctx context.Context, opts ...grpc.CallOption) (MultiRaft_RaftMessageClient, error)
}

type multiRaftClient struct {
	cc *grpc.ClientConn
}

func NewMultiRaftClient(cc *grpc.ClientConn) MultiRaftClient {
	return &multiRaftClient{cc}
}

func (c *multiRaftClient) RaftMessage(ctx context.Context, opts ...grpc.CallOption) (MultiRaft_RaftMessageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MultiRaft_serviceDesc.Streams[0], c.cc, "/cockroach.storage.MultiRaft/RaftMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &multiRaftRaftMessageClient{stream}
	return x, nil
}

type MultiRaft_RaftMessageClient interface {
	Send(*RaftMessageRequest) error
	CloseAndRecv() (*RaftMessageResponse, error)
	grpc.ClientStream
}

type multiRaftRaftMessageClient struct {
	grpc.ClientStream
}

func (x *multiRaftRaftMessageClient) Send(m *RaftMessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *multiRaftRaftMessageClient) CloseAndRecv() (*RaftMessageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RaftMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MultiRaft service

type MultiRaftServer interface {
	RaftMessage(MultiRaft_RaftMessageServer) error
}

func RegisterMultiRaftServer(s *grpc.Server, srv MultiRaftServer) {
	s.RegisterService(&_MultiRaft_serviceDesc, srv)
}

func _MultiRaft_RaftMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MultiRaftServer).RaftMessage(&multiRaftRaftMessageServer{stream})
}

type MultiRaft_RaftMessageServer interface {
	SendAndClose(*RaftMessageResponse) error
	Recv() (*RaftMessageRequest, error)
	grpc.ServerStream
}

type multiRaftRaftMessageServer struct {
	grpc.ServerStream
}

func (x *multiRaftRaftMessageServer) SendAndClose(m *RaftMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *multiRaftRaftMessageServer) Recv() (*RaftMessageRequest, error) {
	m := new(RaftMessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MultiRaft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.storage.MultiRaft",
	HandlerType: (*MultiRaftServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RaftMessage",
			Handler:       _MultiRaft_RaftMessage_Handler,
			ClientStreams: true,
		},
	},
}

func (m *RaftMessageRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftMessageRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintRaft(data, i, uint64(m.GroupID))
	data[i] = 0x12
	i++
	i = encodeVarintRaft(data, i, uint64(m.FromReplica.Size()))
	n1, err := m.FromReplica.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x1a
	i++
	i = encodeVarintRaft(data, i, uint64(m.ToReplica.Size()))
	n2, err := m.ToReplica.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x22
	i++
	i = encodeVarintRaft(data, i, uint64(m.Message.Size()))
	n3, err := m.Message.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *RaftMessageResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftMessageResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ConfChangeContext) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ConfChangeContext) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintRaft(data, i, uint64(len(m.CommandID)))
	i += copy(data[i:], m.CommandID)
	if m.Payload != nil {
		data[i] = 0x12
		i++
		i = encodeVarintRaft(data, i, uint64(len(m.Payload)))
		i += copy(data[i:], m.Payload)
	}
	data[i] = 0x1a
	i++
	i = encodeVarintRaft(data, i, uint64(m.Replica.Size()))
	n4, err := m.Replica.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func encodeFixed64Raft(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Raft(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRaft(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *RaftMessageRequest) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovRaft(uint64(m.GroupID))
	l = m.FromReplica.Size()
	n += 1 + l + sovRaft(uint64(l))
	l = m.ToReplica.Size()
	n += 1 + l + sovRaft(uint64(l))
	l = m.Message.Size()
	n += 1 + l + sovRaft(uint64(l))
	return n
}

func (m *RaftMessageResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ConfChangeContext) Size() (n int) {
	var l int
	_ = l
	l = len(m.CommandID)
	n += 1 + l + sovRaft(uint64(l))
	if m.Payload != nil {
		l = len(m.Payload)
		n += 1 + l + sovRaft(uint64(l))
	}
	l = m.Replica.Size()
	n += 1 + l + sovRaft(uint64(l))
	return n
}

func sovRaft(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRaft(x uint64) (n int) {
	return sovRaft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaftMessageRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RaftMessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftMessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			m.GroupID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.GroupID |= (github_com_cockroachdb_cockroach_roachpb.RangeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromReplica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FromReplica.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToReplica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ToReplica.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Message.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func (m *RaftMessageResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RaftMessageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftMessageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func (m *ConfChangeContext) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfChangeContext: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfChangeContext: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommandID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommandID = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], data[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Replica.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func skipRaft(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRaft
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRaft
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipRaft(data[start:])
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
	ErrInvalidLengthRaft = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRaft   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorRaft = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x6e, 0x9b, 0x30,
	0x18, 0xc7, 0x21, 0x8b, 0xc4, 0x70, 0x22, 0x4d, 0xf1, 0x36, 0x09, 0xb1, 0x29, 0xc9, 0xa2, 0x6d,
	0xca, 0xc9, 0x44, 0x79, 0x04, 0x40, 0x9a, 0x38, 0xe4, 0xc2, 0x69, 0xda, 0x61, 0x99, 0x03, 0x0e,
	0x41, 0x0b, 0x98, 0x19, 0x23, 0xad, 0x6f, 0xd1, 0x97, 0xe8, 0x73, 0xf4, 0x9a, 0x63, 0x8f, 0x3d,
	0x45, 0x6d, 0xfa, 0x16, 0x3d, 0xd5, 0x80, 0x21, 0xad, 0x72, 0xa8, 0xda, 0x83, 0x91, 0xf1, 0xf7,
	0xff, 0xff, 0xfd, 0x7d, 0x3f, 0x19, 0x7c, 0x0e, 0x68, 0xf0, 0x97, 0x51, 0x1c, 0x6c, 0xac, 0x9c,
	0x53, 0x86, 0x23, 0x62, 0x31, 0xbc, 0xe6, 0x28, 0x63, 0x94, 0x53, 0x38, 0x68, 0xab, 0x48, 0x56,
	0xcd, 0xf1, 0xd1, 0x50, 0x7d, 0xb3, 0x95, 0x95, 0x10, 0x8e, 0x43, 0xcc, 0x71, 0x6d, 0x32, 0x3f,
	0x11, 0x1e, 0x84, 0x55, 0x4a, 0xf5, 0x11, 0x82, 0x63, 0xa2, 0xf9, 0x21, 0xa2, 0x11, 0xad, 0xb6,
	0x56, 0xb9, 0xab, 0x4f, 0x27, 0x97, 0x1d, 0x00, 0x7d, 0x21, 0x5a, 0x90, 0x3c, 0x17, 0x97, 0xf8,
	0xe4, 0x5f, 0x41, 0x72, 0x0e, 0x7f, 0x83, 0xb7, 0x11, 0xa3, 0x45, 0xb6, 0x8c, 0x43, 0x43, 0x1d,
	0xab, 0xd3, 0xae, 0xed, 0xec, 0xf6, 0x23, 0xe5, 0xb0, 0x1f, 0x69, 0x3f, 0xca, 0x73, 0xcf, 0xbd,
	0xdf, 0x8f, 0x66, 0x51, 0xcc, 0x37, 0xc5, 0x0a, 0x05, 0x34, 0xb1, 0xda, 0xde, 0xc2, 0x95, 0x75,
	0xd2, 0x27, 0xf2, 0x71, 0x1a, 0x11, 0xcf, 0xf5, 0xb5, 0x2a, 0xd4, 0x0b, 0xe1, 0x02, 0xf4, 0xd7,
	0x8c, 0x26, 0x4b, 0x46, 0xb2, 0x6d, 0x1c, 0x60, 0xa3, 0x23, 0xee, 0xe8, 0xcd, 0xbf, 0xa2, 0xe3,
	0xd4, 0xad, 0xb5, 0x56, 0xb8, 0x24, 0x0f, 0x58, 0x9c, 0x09, 0x14, 0x76, 0xb7, 0xec, 0xc4, 0xef,
	0x95, 0x7e, 0x59, 0x84, 0x1e, 0x00, 0x9c, 0xb6, 0x61, 0x6f, 0x5e, 0x1c, 0xa6, 0x73, 0xda, 0x44,
	0x59, 0x40, 0x4b, 0x6a, 0x16, 0x46, 0xb7, 0xca, 0x79, 0x87, 0x6a, 0x96, 0x48, 0x22, 0x92, 0x96,
	0x46, 0x35, 0xf9, 0x08, 0xde, 0x3f, 0x01, 0x98, 0x67, 0x34, 0xcd, 0xc9, 0xe4, 0x42, 0x05, 0x03,
	0x87, 0xa6, 0x6b, 0x67, 0x53, 0xce, 0x2e, 0x76, 0x9c, 0xfc, 0xe7, 0x70, 0x06, 0x80, 0xa0, 0x95,
	0xe0, 0x34, 0x6c, 0xc8, 0xea, 0xf6, 0x40, 0x92, 0xd5, 0x9d, 0xba, 0x22, 0x38, 0xe9, 0x52, 0x24,
	0x48, 0x19, 0x40, 0xcb, 0xf0, 0xd9, 0x96, 0xe2, 0xb0, 0x82, 0xd4, 0xf7, 0x9b, 0x5f, 0xe8, 0x02,
	0xed, 0xf5, 0x13, 0x37, 0xd6, 0x79, 0x02, 0xf4, 0x45, 0xb1, 0xe5, 0x71, 0x39, 0x03, 0xfc, 0x03,
	0x7a, 0x8f, 0x66, 0x81, 0xdf, 0xd0, 0xc9, 0x2b, 0x44, 0xa7, 0x8f, 0xc5, 0xfc, 0xfe, 0x9c, 0x4c,
	0x22, 0x51, 0xa6, 0xaa, 0xfd, 0x65, 0x77, 0x3b, 0x54, 0x76, 0x87, 0xa1, 0x7a, 0x25, 0xd6, 0xb5,
	0x58, 0x37, 0x62, 0x9d, 0xdf, 0x0d, 0x95, 0x5f, 0x9a, 0xb4, 0xfe, 0xec, 0x3c, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xae, 0x2e, 0xa5, 0xf6, 0x1b, 0x03, 0x00, 0x00,
}
