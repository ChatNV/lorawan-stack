// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/applicationserver.proto

package ttnpb // import "go.thethings.network/lorawan-stack/pkg/ttnpb"

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import context "context"
import grpc "google.golang.org/grpc"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ApplicationLink struct {
	NetworkServerAddress string   `protobuf:"bytes,1,opt,name=network_server_address,json=networkServerAddress,proto3" json:"network_server_address,omitempty"`
	APIKey               string   `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplicationLink) Reset()      { *m = ApplicationLink{} }
func (*ApplicationLink) ProtoMessage() {}
func (*ApplicationLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_applicationserver_845b29a6795362f3, []int{0}
}
func (m *ApplicationLink) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApplicationLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApplicationLink.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ApplicationLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationLink.Merge(dst, src)
}
func (m *ApplicationLink) XXX_Size() int {
	return m.Size()
}
func (m *ApplicationLink) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationLink.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationLink proto.InternalMessageInfo

func (m *ApplicationLink) GetNetworkServerAddress() string {
	if m != nil {
		return m.NetworkServerAddress
	}
	return ""
}

func (m *ApplicationLink) GetAPIKey() string {
	if m != nil {
		return m.APIKey
	}
	return ""
}

type SetApplicationLinkRequest struct {
	ApplicationIdentifiers `protobuf:"bytes,1,opt,name=application_ids,json=applicationIds,embedded=application_ids" json:"application_ids"`
	ApplicationLink        `protobuf:"bytes,2,opt,name=link,embedded=link" json:"link"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *SetApplicationLinkRequest) Reset()      { *m = SetApplicationLinkRequest{} }
func (*SetApplicationLinkRequest) ProtoMessage() {}
func (*SetApplicationLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_applicationserver_845b29a6795362f3, []int{1}
}
func (m *SetApplicationLinkRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetApplicationLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetApplicationLinkRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SetApplicationLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetApplicationLinkRequest.Merge(dst, src)
}
func (m *SetApplicationLinkRequest) XXX_Size() int {
	return m.Size()
}
func (m *SetApplicationLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetApplicationLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetApplicationLinkRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ApplicationLink)(nil), "ttn.lorawan.v3.ApplicationLink")
	golang_proto.RegisterType((*ApplicationLink)(nil), "ttn.lorawan.v3.ApplicationLink")
	proto.RegisterType((*SetApplicationLinkRequest)(nil), "ttn.lorawan.v3.SetApplicationLinkRequest")
	golang_proto.RegisterType((*SetApplicationLinkRequest)(nil), "ttn.lorawan.v3.SetApplicationLinkRequest")
}
func (this *ApplicationLink) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ApplicationLink)
	if !ok {
		that2, ok := that.(ApplicationLink)
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
	if this.NetworkServerAddress != that1.NetworkServerAddress {
		return false
	}
	if this.APIKey != that1.APIKey {
		return false
	}
	return true
}
func (this *SetApplicationLinkRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SetApplicationLinkRequest)
	if !ok {
		that2, ok := that.(SetApplicationLinkRequest)
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
	if !this.ApplicationIdentifiers.Equal(&that1.ApplicationIdentifiers) {
		return false
	}
	if !this.ApplicationLink.Equal(&that1.ApplicationLink) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for As service

type AsClient interface {
	GetLink(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*ApplicationLink, error)
	SetLink(ctx context.Context, in *SetApplicationLinkRequest, opts ...grpc.CallOption) (*types.Empty, error)
	DeleteLink(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
	Subscribe(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (As_SubscribeClient, error)
}

type asClient struct {
	cc *grpc.ClientConn
}

func NewAsClient(cc *grpc.ClientConn) AsClient {
	return &asClient{cc}
}

func (c *asClient) GetLink(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*ApplicationLink, error) {
	out := new(ApplicationLink)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.As/GetLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asClient) SetLink(ctx context.Context, in *SetApplicationLinkRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.As/SetLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asClient) DeleteLink(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.As/DeleteLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asClient) Subscribe(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (As_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_As_serviceDesc.Streams[0], "/ttn.lorawan.v3.As/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &asSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type As_SubscribeClient interface {
	Recv() (*ApplicationUp, error)
	grpc.ClientStream
}

type asSubscribeClient struct {
	grpc.ClientStream
}

func (x *asSubscribeClient) Recv() (*ApplicationUp, error) {
	m := new(ApplicationUp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for As service

type AsServer interface {
	GetLink(context.Context, *ApplicationIdentifiers) (*ApplicationLink, error)
	SetLink(context.Context, *SetApplicationLinkRequest) (*types.Empty, error)
	DeleteLink(context.Context, *ApplicationIdentifiers) (*types.Empty, error)
	Subscribe(*ApplicationIdentifiers, As_SubscribeServer) error
}

func RegisterAsServer(s *grpc.Server, srv AsServer) {
	s.RegisterService(&_As_serviceDesc, srv)
}

func _As_GetLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsServer).GetLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.As/GetLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsServer).GetLink(ctx, req.(*ApplicationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _As_SetLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetApplicationLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsServer).SetLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.As/SetLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsServer).SetLink(ctx, req.(*SetApplicationLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _As_DeleteLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsServer).DeleteLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.As/DeleteLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsServer).DeleteLink(ctx, req.(*ApplicationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _As_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ApplicationIdentifiers)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AsServer).Subscribe(m, &asSubscribeServer{stream})
}

type As_SubscribeServer interface {
	Send(*ApplicationUp) error
	grpc.ServerStream
}

type asSubscribeServer struct {
	grpc.ServerStream
}

func (x *asSubscribeServer) Send(m *ApplicationUp) error {
	return x.ServerStream.SendMsg(m)
}

var _As_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.As",
	HandlerType: (*AsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLink",
			Handler:    _As_GetLink_Handler,
		},
		{
			MethodName: "SetLink",
			Handler:    _As_SetLink_Handler,
		},
		{
			MethodName: "DeleteLink",
			Handler:    _As_DeleteLink_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _As_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lorawan-stack/api/applicationserver.proto",
}

func (m *ApplicationLink) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApplicationLink) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NetworkServerAddress) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApplicationserver(dAtA, i, uint64(len(m.NetworkServerAddress)))
		i += copy(dAtA[i:], m.NetworkServerAddress)
	}
	if len(m.APIKey) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApplicationserver(dAtA, i, uint64(len(m.APIKey)))
		i += copy(dAtA[i:], m.APIKey)
	}
	return i, nil
}

func (m *SetApplicationLinkRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetApplicationLinkRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintApplicationserver(dAtA, i, uint64(m.ApplicationIdentifiers.Size()))
	n1, err := m.ApplicationIdentifiers.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintApplicationserver(dAtA, i, uint64(m.ApplicationLink.Size()))
	n2, err := m.ApplicationLink.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func encodeVarintApplicationserver(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedApplicationLink(r randyApplicationserver, easy bool) *ApplicationLink {
	this := &ApplicationLink{}
	this.NetworkServerAddress = randStringApplicationserver(r)
	this.APIKey = randStringApplicationserver(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedSetApplicationLinkRequest(r randyApplicationserver, easy bool) *SetApplicationLinkRequest {
	this := &SetApplicationLinkRequest{}
	v1 := NewPopulatedApplicationIdentifiers(r, easy)
	this.ApplicationIdentifiers = *v1
	v2 := NewPopulatedApplicationLink(r, easy)
	this.ApplicationLink = *v2
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyApplicationserver interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneApplicationserver(r randyApplicationserver) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringApplicationserver(r randyApplicationserver) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneApplicationserver(r)
	}
	return string(tmps)
}
func randUnrecognizedApplicationserver(r randyApplicationserver, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldApplicationserver(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldApplicationserver(dAtA []byte, r randyApplicationserver, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateApplicationserver(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateApplicationserver(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateApplicationserver(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateApplicationserver(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateApplicationserver(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateApplicationserver(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateApplicationserver(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ApplicationLink) Size() (n int) {
	var l int
	_ = l
	l = len(m.NetworkServerAddress)
	if l > 0 {
		n += 1 + l + sovApplicationserver(uint64(l))
	}
	l = len(m.APIKey)
	if l > 0 {
		n += 1 + l + sovApplicationserver(uint64(l))
	}
	return n
}

func (m *SetApplicationLinkRequest) Size() (n int) {
	var l int
	_ = l
	l = m.ApplicationIdentifiers.Size()
	n += 1 + l + sovApplicationserver(uint64(l))
	l = m.ApplicationLink.Size()
	n += 1 + l + sovApplicationserver(uint64(l))
	return n
}

func sovApplicationserver(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApplicationserver(x uint64) (n int) {
	return sovApplicationserver((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *ApplicationLink) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ApplicationLink{`,
		`NetworkServerAddress:` + fmt.Sprintf("%v", this.NetworkServerAddress) + `,`,
		`APIKey:` + fmt.Sprintf("%v", this.APIKey) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SetApplicationLinkRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SetApplicationLinkRequest{`,
		`ApplicationIdentifiers:` + strings.Replace(strings.Replace(this.ApplicationIdentifiers.String(), "ApplicationIdentifiers", "ApplicationIdentifiers", 1), `&`, ``, 1) + `,`,
		`ApplicationLink:` + strings.Replace(strings.Replace(this.ApplicationLink.String(), "ApplicationLink", "ApplicationLink", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringApplicationserver(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ApplicationLink) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplicationserver
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
			return fmt.Errorf("proto: ApplicationLink: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApplicationLink: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkServerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserver
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
				return ErrInvalidLengthApplicationserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NetworkServerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field APIKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserver
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
				return ErrInvalidLengthApplicationserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.APIKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplicationserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApplicationserver
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
func (m *SetApplicationLinkRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplicationserver
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
			return fmt.Errorf("proto: SetApplicationLinkRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetApplicationLinkRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationIdentifiers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserver
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
				return ErrInvalidLengthApplicationserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ApplicationIdentifiers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationLink", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserver
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
				return ErrInvalidLengthApplicationserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ApplicationLink.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplicationserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApplicationserver
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
func skipApplicationserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApplicationserver
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
					return 0, ErrIntOverflowApplicationserver
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
					return 0, ErrIntOverflowApplicationserver
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
				return 0, ErrInvalidLengthApplicationserver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApplicationserver
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
				next, err := skipApplicationserver(dAtA[start:])
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
	ErrInvalidLengthApplicationserver = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApplicationserver   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("lorawan-stack/api/applicationserver.proto", fileDescriptor_applicationserver_845b29a6795362f3)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/applicationserver.proto", fileDescriptor_applicationserver_845b29a6795362f3)
}

var fileDescriptor_applicationserver_845b29a6795362f3 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x3d, 0x4c, 0x14, 0x41,
	0x18, 0x9d, 0x41, 0x03, 0x32, 0x24, 0x90, 0x6c, 0x0c, 0xc1, 0x53, 0xbf, 0x23, 0x47, 0x42, 0x80,
	0xc8, 0xac, 0x01, 0x63, 0xa1, 0xb1, 0x38, 0xa2, 0x31, 0x04, 0x0b, 0x73, 0xa7, 0x85, 0x36, 0x97,
	0xb9, 0xbb, 0x61, 0x99, 0xec, 0xb2, 0xb3, 0xee, 0xcc, 0x41, 0x2e, 0x6a, 0x42, 0xa8, 0xe8, 0x34,
	0xb1, 0xb1, 0x34, 0x56, 0x94, 0x34, 0x26, 0x94, 0x94, 0x94, 0x24, 0x36, 0x54, 0x84, 0x9d, 0xb5,
	0xa0, 0x93, 0x92, 0xd2, 0xdc, 0xee, 0x22, 0x77, 0x20, 0x3f, 0x76, 0x33, 0x79, 0x6f, 0xde, 0x7b,
	0xdf, 0xcb, 0x37, 0x64, 0xdc, 0x93, 0x21, 0x5b, 0x66, 0xfe, 0xa4, 0xd2, 0xac, 0xe6, 0xda, 0x2c,
	0x10, 0x36, 0x0b, 0x02, 0x4f, 0xd4, 0x98, 0x16, 0xd2, 0x57, 0x3c, 0x5c, 0xe2, 0x21, 0x0d, 0x42,
	0xa9, 0xa5, 0xd5, 0xaf, 0xb5, 0x4f, 0x33, 0x3a, 0x5d, 0x9a, 0xce, 0x4d, 0x3a, 0x42, 0x2f, 0x34,
	0xaa, 0xb4, 0x26, 0x17, 0x6d, 0x47, 0x3a, 0xd2, 0x4e, 0x68, 0xd5, 0xc6, 0x7c, 0x72, 0x4b, 0x2e,
	0xc9, 0x29, 0x7d, 0x9e, 0xbb, 0xe3, 0x48, 0xe9, 0x78, 0x3c, 0xb5, 0xf0, 0x7d, 0xa9, 0x53, 0x87,
	0x0c, 0xbd, 0x9d, 0xa1, 0x7f, 0x35, 0xf8, 0x62, 0xa0, 0x9b, 0x19, 0x38, 0x72, 0x36, 0xa4, 0xa8,
	0x73, 0x5f, 0x8b, 0x79, 0xc1, 0xc3, 0x63, 0x85, 0xe1, 0xb3, 0xa4, 0x45, 0xae, 0x14, 0x73, 0x78,
	0xc6, 0x28, 0x78, 0x64, 0xa0, 0x78, 0x32, 0xdb, 0x0b, 0xe1, 0xbb, 0xd6, 0x03, 0x32, 0xe8, 0x73,
	0xbd, 0x2c, 0x43, 0xb7, 0x92, 0xce, 0x5a, 0x61, 0xf5, 0x7a, 0xc8, 0x95, 0x1a, 0xc2, 0xc3, 0x78,
	0xac, 0xb7, 0x74, 0x33, 0x43, 0xcb, 0x09, 0x58, 0x4c, 0x31, 0x6b, 0x84, 0xf4, 0xb0, 0x40, 0x54,
	0x5c, 0xde, 0x1c, 0xea, 0x6a, 0xd1, 0x66, 0x88, 0xd9, 0xcb, 0x77, 0x17, 0x5f, 0xce, 0xce, 0xf1,
	0x66, 0xa9, 0x9b, 0x05, 0x62, 0x8e, 0x37, 0x0b, 0x3f, 0x30, 0xb9, 0x55, 0xe6, 0xfa, 0x94, 0x63,
	0x89, 0xbf, 0x6b, 0x70, 0xa5, 0xad, 0x37, 0x64, 0xa0, 0xad, 0xe7, 0x8a, 0xa8, 0xa7, 0x8e, 0x7d,
	0x53, 0xa3, 0xb4, 0xb3, 0x66, 0xda, 0x26, 0x30, 0x7b, 0x32, 0xf4, 0xcc, 0x8d, 0xed, 0xbd, 0x3c,
	0xda, 0xd9, 0xcb, 0xe3, 0x52, 0x3f, 0x6b, 0x67, 0x28, 0xeb, 0x09, 0xb9, 0xee, 0x09, 0xdf, 0x4d,
	0xa2, 0xf5, 0x4d, 0xe5, 0x2f, 0xd0, 0x6b, 0x05, 0x6a, 0x13, 0x4a, 0x9e, 0x4d, 0xfd, 0xbe, 0x46,
	0xba, 0x8a, 0xca, 0x5a, 0xc5, 0xa4, 0xe7, 0x39, 0xd7, 0x49, 0x4b, 0x57, 0xcc, 0x94, 0xbb, 0xcc,
	0xab, 0x40, 0x57, 0x7f, 0xfe, 0xfa, 0xd2, 0x35, 0x66, 0x8d, 0xda, 0x4c, 0x75, 0xec, 0x99, 0xfd,
	0xbe, 0xb3, 0x8d, 0x8f, 0x76, 0x2b, 0x8b, 0xf5, 0x09, 0x93, 0x9e, 0x72, 0x16, 0x62, 0xfc, 0xb4,
	0xf8, 0xb9, 0xe5, 0xe6, 0x06, 0x69, 0xba, 0x4d, 0xf4, 0x78, 0x9b, 0xe8, 0xb3, 0xd6, 0x36, 0x15,
	0x8a, 0x89, 0xfd, 0xe3, 0xdc, 0xc3, 0xcb, 0xec, 0x15, 0xfd, 0x57, 0x9c, 0x47, 0x78, 0xc2, 0xfa,
	0x40, 0xc8, 0x53, 0xee, 0x71, 0xcd, 0xff, 0xab, 0x98, 0xf3, 0x02, 0x65, 0x7d, 0x4c, 0x5c, 0xb5,
	0x8f, 0x57, 0xa4, 0xb7, 0xdc, 0xa8, 0xaa, 0x5a, 0x28, 0xaa, 0xfc, 0xca, 0xe6, 0x77, 0x2f, 0xe0,
	0xbd, 0x0e, 0xee, 0xe3, 0x99, 0xef, 0x78, 0x3b, 0x02, 0xbc, 0x13, 0x01, 0xde, 0x8d, 0x00, 0xed,
	0x47, 0x80, 0x0e, 0x22, 0x40, 0x87, 0x11, 0xa0, 0xa3, 0x08, 0xf0, 0x8a, 0x01, 0xbc, 0x66, 0x00,
	0xad, 0x1b, 0xc0, 0x1b, 0x06, 0xd0, 0xa6, 0x01, 0xb4, 0x65, 0x00, 0x6d, 0x1b, 0xc0, 0x3b, 0x06,
	0xf0, 0xae, 0x01, 0xb4, 0x6f, 0x00, 0x1f, 0x18, 0x40, 0x87, 0x06, 0xf0, 0x91, 0x01, 0xb4, 0x12,
	0x03, 0x5a, 0x8b, 0x01, 0x7f, 0x8e, 0x01, 0x7d, 0x8d, 0x01, 0x7f, 0x8b, 0x01, 0xad, 0xc7, 0x80,
	0x36, 0x62, 0xc0, 0x9b, 0x31, 0xe0, 0xad, 0x18, 0xf0, 0xdb, 0x7b, 0x8e, 0xa4, 0x7a, 0x81, 0xeb,
	0x05, 0xe1, 0x3b, 0x8a, 0x66, 0x1f, 0xcd, 0xee, 0xfc, 0xc5, 0x81, 0xeb, 0xd8, 0x5a, 0xfb, 0x41,
	0xb5, 0xda, 0x9d, 0x54, 0x37, 0xfd, 0x27, 0x00, 0x00, 0xff, 0xff, 0x99, 0x65, 0xf8, 0xfe, 0xb1,
	0x04, 0x00, 0x00,
}
