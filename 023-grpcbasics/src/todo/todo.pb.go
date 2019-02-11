// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

/*
Package todo is a generated protocol buffer package.

It is generated from these files:
	todo.proto

It has these top-level messages:
	Task
	TaskList
	Text
	Void
*/
package todo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Task struct {
	Todo string `protobuf:"bytes,1,opt,name=todo" json:"todo,omitempty"`
	Done bool   `protobuf:"varint,2,opt,name=done" json:"done,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Task) GetTodo() string {
	if m != nil {
		return m.Todo
	}
	return ""
}

func (m *Task) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

type TaskList struct {
	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *TaskList) Reset()                    { *m = TaskList{} }
func (m *TaskList) String() string            { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()               {}
func (*TaskList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TaskList) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type Text struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *Text) Reset()                    { *m = Text{} }
func (m *Text) String() string            { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()               {}
func (*Text) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Text) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Task)(nil), "todo.Task")
	proto.RegisterType((*TaskList)(nil), "todo.TaskList")
	proto.RegisterType((*Text)(nil), "todo.Text")
	proto.RegisterType((*Void)(nil), "todo.Void")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Tasks service

type TasksClient interface {
	List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TaskList, error)
	Add(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Task, error)
}

type tasksClient struct {
	cc *grpc.ClientConn
}

func NewTasksClient(cc *grpc.ClientConn) TasksClient {
	return &tasksClient{cc}
}

func (c *tasksClient) List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TaskList, error) {
	out := new(TaskList)
	err := grpc.Invoke(ctx, "/todo.Tasks/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) Add(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/todo.Tasks/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tasks service

type TasksServer interface {
	List(context.Context, *Void) (*TaskList, error)
	Add(context.Context, *Text) (*Task, error)
}

func RegisterTasksServer(s *grpc.Server, srv TasksServer) {
	s.RegisterService(&_Tasks_serviceDesc, srv)
}

func _Tasks_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Tasks/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).List(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Tasks/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).Add(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tasks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.Tasks",
	HandlerType: (*TasksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Tasks_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Tasks_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xbd, 0xaa, 0xc3, 0x30,
	0x0c, 0x85, 0xe3, 0x1b, 0x27, 0xe4, 0xaa, 0xd0, 0xc1, 0x53, 0xc8, 0x52, 0x63, 0x3a, 0x64, 0x28,
	0x19, 0xd2, 0x27, 0xe8, 0x5e, 0x3a, 0x84, 0xd0, 0x3d, 0xc5, 0x1e, 0x42, 0xa0, 0x2a, 0xb5, 0x06,
	0x3f, 0x7e, 0x91, 0xfa, 0xbb, 0x1d, 0xfb, 0x3b, 0xfa, 0x84, 0x00, 0x08, 0x3d, 0x76, 0xb7, 0x3b,
	0x12, 0x1a, 0xcd, 0xd9, 0x75, 0xa0, 0xc7, 0x29, 0x2e, 0xc6, 0x80, 0xbc, 0x6b, 0x65, 0x55, 0xfb,
	0x3f, 0x48, 0xe6, 0x3f, 0x8f, 0xd7, 0x50, 0xff, 0x59, 0xd5, 0x56, 0x83, 0x64, 0xb7, 0x83, 0x8a,
	0xfb, 0xc7, 0x39, 0x92, 0xb1, 0x50, 0xd0, 0x14, 0x97, 0x58, 0x2b, 0x9b, 0xb7, 0xab, 0x1e, 0x3a,
	0xb1, 0x33, 0x1e, 0x9e, 0xc0, 0x35, 0xa0, 0xc7, 0x90, 0x48, 0xec, 0x21, 0xd1, 0xc7, 0x1e, 0x12,
	0xb9, 0x12, 0xf4, 0x19, 0x67, 0xdf, 0x9f, 0xa0, 0xe0, 0x91, 0x68, 0xb6, 0xa0, 0x45, 0xfb, 0xf2,
	0x30, 0x6c, 0xd6, 0x5f, 0x27, 0x33, 0x97, 0x99, 0x0d, 0xe4, 0x07, 0xef, 0xdf, 0x25, 0xb6, 0x37,
	0x3f, 0x8b, 0x5d, 0x76, 0x29, 0xe5, 0xbc, 0xfd, 0x23, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xd0, 0xa6,
	0x74, 0xec, 0x00, 0x00, 0x00,
}