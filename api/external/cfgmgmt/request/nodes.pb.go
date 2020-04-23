// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/cfgmgmt/request/nodes.proto

package request

import (
	fmt "fmt"
	query "github.com/chef/automate/api/external/common/query"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Nodes struct {
	// Filters to apply to the request for nodes list.
	Filter []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	// Pagination parameters to apply to the returned node list.
	Pagination *query.Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	// Sorting parameters to apply to the returned node list.
	Sorting *query.Sorting `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	// Earliest most recent check-in node information to return.
	Start string `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	// Latest most recent check-in node information to return.
	End                  string   `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nodes) Reset()         { *m = Nodes{} }
func (m *Nodes) String() string { return proto.CompactTextString(m) }
func (*Nodes) ProtoMessage()    {}
func (*Nodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{0}
}

func (m *Nodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nodes.Unmarshal(m, b)
}
func (m *Nodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nodes.Marshal(b, m, deterministic)
}
func (m *Nodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nodes.Merge(m, src)
}
func (m *Nodes) XXX_Size() int {
	return xxx_messageInfo_Nodes.Size(m)
}
func (m *Nodes) XXX_DiscardUnknown() {
	xxx_messageInfo_Nodes.DiscardUnknown(m)
}

var xxx_messageInfo_Nodes proto.InternalMessageInfo

func (m *Nodes) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *Nodes) GetPagination() *query.Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *Nodes) GetSorting() *query.Sorting {
	if m != nil {
		return m.Sorting
	}
	return nil
}

func (m *Nodes) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *Nodes) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

type Node struct {
	// Chef guid for the requested node.
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{1}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type NodeRun struct {
	// Chef guid for the requested node.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Run id for the node.
	RunId string `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	// End time on the node's run.
	EndTime              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NodeRun) Reset()         { *m = NodeRun{} }
func (m *NodeRun) String() string { return proto.CompactTextString(m) }
func (*NodeRun) ProtoMessage()    {}
func (*NodeRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{2}
}

func (m *NodeRun) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRun.Unmarshal(m, b)
}
func (m *NodeRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRun.Marshal(b, m, deterministic)
}
func (m *NodeRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRun.Merge(m, src)
}
func (m *NodeRun) XXX_Size() int {
	return xxx_messageInfo_NodeRun.Size(m)
}
func (m *NodeRun) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRun.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRun proto.InternalMessageInfo

func (m *NodeRun) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *NodeRun) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *NodeRun) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type Runs struct {
	// Chef guid for the node.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Filters to apply to the request for runs list.
	Filter []string `protobuf:"bytes,2,rep,name=filter,proto3" json:"filter,omitempty"`
	// Pagination parameters to apply to the returned runs list.
	Pagination *query.Pagination `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
	// Earliest (in history) run information to return for the runs list.
	Start string `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	// Latest (in history) run information to return for the runs list.
	End                  string   `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Runs) Reset()         { *m = Runs{} }
func (m *Runs) String() string { return proto.CompactTextString(m) }
func (*Runs) ProtoMessage()    {}
func (*Runs) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{3}
}

func (m *Runs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Runs.Unmarshal(m, b)
}
func (m *Runs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Runs.Marshal(b, m, deterministic)
}
func (m *Runs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Runs.Merge(m, src)
}
func (m *Runs) XXX_Size() int {
	return xxx_messageInfo_Runs.Size(m)
}
func (m *Runs) XXX_DiscardUnknown() {
	xxx_messageInfo_Runs.DiscardUnknown(m)
}

var xxx_messageInfo_Runs proto.InternalMessageInfo

func (m *Runs) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *Runs) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *Runs) GetPagination() *query.Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *Runs) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *Runs) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

type NodeExport struct {
	// Filters to apply to the request for nodes list.
	Filter []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	// Sorting parameters to apply to the returned node list.
	Sorting *query.Sorting `protobuf:"bytes,2,opt,name=sorting,proto3" json:"sorting,omitempty"`
	// File type, either JSON or CSV.
	OutputType           string   `protobuf:"bytes,3,opt,name=output_type,json=outputType,proto3" json:"output_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExport) Reset()         { *m = NodeExport{} }
func (m *NodeExport) String() string { return proto.CompactTextString(m) }
func (*NodeExport) ProtoMessage()    {}
func (*NodeExport) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{4}
}

func (m *NodeExport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExport.Unmarshal(m, b)
}
func (m *NodeExport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExport.Marshal(b, m, deterministic)
}
func (m *NodeExport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExport.Merge(m, src)
}
func (m *NodeExport) XXX_Size() int {
	return xxx_messageInfo_NodeExport.Size(m)
}
func (m *NodeExport) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExport.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExport proto.InternalMessageInfo

func (m *NodeExport) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *NodeExport) GetSorting() *query.Sorting {
	if m != nil {
		return m.Sorting
	}
	return nil
}

func (m *NodeExport) GetOutputType() string {
	if m != nil {
		return m.OutputType
	}
	return ""
}

type ReportExport struct {
	// Filters to apply to the request for node report list.
	Filter []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	// File type, either JSON or CSV.
	OutputType string `protobuf:"bytes,2,opt,name=output_type,json=outputType,proto3" json:"output_type,omitempty"`
	// The node ID of the reports
	NodeId string `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Earliest (in history) run information to return for the runs list.
	Start *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	// Latest (in history) run information to return for the runs list.
	End                  *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ReportExport) Reset()         { *m = ReportExport{} }
func (m *ReportExport) String() string { return proto.CompactTextString(m) }
func (*ReportExport) ProtoMessage()    {}
func (*ReportExport) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{5}
}

func (m *ReportExport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportExport.Unmarshal(m, b)
}
func (m *ReportExport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportExport.Marshal(b, m, deterministic)
}
func (m *ReportExport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportExport.Merge(m, src)
}
func (m *ReportExport) XXX_Size() int {
	return xxx_messageInfo_ReportExport.Size(m)
}
func (m *ReportExport) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportExport.DiscardUnknown(m)
}

var xxx_messageInfo_ReportExport proto.InternalMessageInfo

func (m *ReportExport) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ReportExport) GetOutputType() string {
	if m != nil {
		return m.OutputType
	}
	return ""
}

func (m *ReportExport) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *ReportExport) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *ReportExport) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

type MissingNodeDurationCounts struct {
	// A valid duration is any number zero or greater with one of these characters 'h', 'd', 'w', or 'M'.
	// 'h' is hours
	// 'd' is days
	// 'w' is weeks
	// 'M' is months
	// Will contain one or many.
	Durations            []string `protobuf:"bytes,1,rep,name=durations,proto3" json:"durations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MissingNodeDurationCounts) Reset()         { *m = MissingNodeDurationCounts{} }
func (m *MissingNodeDurationCounts) String() string { return proto.CompactTextString(m) }
func (*MissingNodeDurationCounts) ProtoMessage()    {}
func (*MissingNodeDurationCounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{6}
}

func (m *MissingNodeDurationCounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MissingNodeDurationCounts.Unmarshal(m, b)
}
func (m *MissingNodeDurationCounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MissingNodeDurationCounts.Marshal(b, m, deterministic)
}
func (m *MissingNodeDurationCounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissingNodeDurationCounts.Merge(m, src)
}
func (m *MissingNodeDurationCounts) XXX_Size() int {
	return xxx_messageInfo_MissingNodeDurationCounts.Size(m)
}
func (m *MissingNodeDurationCounts) XXX_DiscardUnknown() {
	xxx_messageInfo_MissingNodeDurationCounts.DiscardUnknown(m)
}

var xxx_messageInfo_MissingNodeDurationCounts proto.InternalMessageInfo

func (m *MissingNodeDurationCounts) GetDurations() []string {
	if m != nil {
		return m.Durations
	}
	return nil
}

type NodeMetadataCounts struct {
	// Types of node fields to collect value counts for
	Type []string `protobuf:"bytes,1,rep,name=type,proto3" json:"type,omitempty"`
	// Filters to apply to the counts returned
	Filter []string `protobuf:"bytes,2,rep,name=filter,proto3" json:"filter,omitempty"`
	// Earliest most recent check-in node information to return.
	Start string `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	// Latest most recent check-in node information to return.
	End                  string   `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeMetadataCounts) Reset()         { *m = NodeMetadataCounts{} }
func (m *NodeMetadataCounts) String() string { return proto.CompactTextString(m) }
func (*NodeMetadataCounts) ProtoMessage()    {}
func (*NodeMetadataCounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{7}
}

func (m *NodeMetadataCounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeMetadataCounts.Unmarshal(m, b)
}
func (m *NodeMetadataCounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeMetadataCounts.Marshal(b, m, deterministic)
}
func (m *NodeMetadataCounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMetadataCounts.Merge(m, src)
}
func (m *NodeMetadataCounts) XXX_Size() int {
	return xxx_messageInfo_NodeMetadataCounts.Size(m)
}
func (m *NodeMetadataCounts) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMetadataCounts.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMetadataCounts proto.InternalMessageInfo

func (m *NodeMetadataCounts) GetType() []string {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *NodeMetadataCounts) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *NodeMetadataCounts) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *NodeMetadataCounts) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func init() {
	proto.RegisterType((*Nodes)(nil), "chef.automate.api.cfgmgmt.request.Nodes")
	proto.RegisterType((*Node)(nil), "chef.automate.api.cfgmgmt.request.Node")
	proto.RegisterType((*NodeRun)(nil), "chef.automate.api.cfgmgmt.request.NodeRun")
	proto.RegisterType((*Runs)(nil), "chef.automate.api.cfgmgmt.request.Runs")
	proto.RegisterType((*NodeExport)(nil), "chef.automate.api.cfgmgmt.request.NodeExport")
	proto.RegisterType((*ReportExport)(nil), "chef.automate.api.cfgmgmt.request.ReportExport")
	proto.RegisterType((*MissingNodeDurationCounts)(nil), "chef.automate.api.cfgmgmt.request.MissingNodeDurationCounts")
	proto.RegisterType((*NodeMetadataCounts)(nil), "chef.automate.api.cfgmgmt.request.NodeMetadataCounts")
}

func init() {
	proto.RegisterFile("api/external/cfgmgmt/request/nodes.proto", fileDescriptor_eeeb0a9b2be1cdd5)
}

var fileDescriptor_eeeb0a9b2be1cdd5 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xc7, 0xb5, 0xd9, 0x34, 0x21, 0xd3, 0x0a, 0x81, 0xc5, 0xc7, 0x12, 0x21, 0x25, 0xec, 0x85,
	0xa8, 0x6a, 0xd7, 0x08, 0x54, 0x21, 0x71, 0x2b, 0x1f, 0x87, 0x22, 0x8a, 0xd0, 0xd2, 0x13, 0x97,
	0xc8, 0xc9, 0x3a, 0x9b, 0x95, 0xb2, 0xb6, 0x6b, 0x8f, 0x45, 0xf3, 0x06, 0x9c, 0x79, 0x0d, 0x5e,
	0x82, 0x67, 0xe0, 0xc0, 0xf3, 0x20, 0x7b, 0x37, 0x1f, 0xad, 0x9a, 0x14, 0xc1, 0x29, 0xf6, 0x64,
	0xe6, 0xef, 0x99, 0xff, 0x6f, 0xb4, 0x30, 0x60, 0xaa, 0xa0, 0xfc, 0x02, 0xb9, 0x16, 0x6c, 0x46,
	0xc7, 0x93, 0xbc, 0xcc, 0x4b, 0xa4, 0x9a, 0x9f, 0x5b, 0x6e, 0x90, 0x0a, 0x99, 0x71, 0x93, 0x28,
	0x2d, 0x51, 0x92, 0x27, 0xe3, 0x29, 0x9f, 0x24, 0xcc, 0xa2, 0x2c, 0x19, 0xf2, 0x84, 0xa9, 0x22,
	0xa9, 0xd3, 0x93, 0x3a, 0xbd, 0xdb, 0xcb, 0xa5, 0xcc, 0x67, 0x9c, 0xfa, 0x82, 0x91, 0x9d, 0x50,
	0x2c, 0x4a, 0x6e, 0x90, 0x95, 0xaa, 0xd2, 0xe8, 0xee, 0x5f, 0x7e, 0x4d, 0x96, 0xa5, 0x14, 0xf4,
	0xdc, 0x72, 0x3d, 0xa7, 0x8a, 0x69, 0x56, 0x72, 0xe4, 0xba, 0x7e, 0xaf, 0x7b, 0xe0, 0x7f, 0xc6,
	0x87, 0x39, 0x17, 0x87, 0xe6, 0x2b, 0xcb, 0x73, 0xae, 0xa9, 0x54, 0x58, 0x48, 0x61, 0x28, 0x13,
	0x42, 0x22, 0xf3, 0xe7, 0x2a, 0x3b, 0xfe, 0x1d, 0xc0, 0xce, 0x47, 0xd7, 0x2d, 0x79, 0x00, 0xad,
	0x49, 0x31, 0x43, 0xae, 0xa3, 0xa0, 0x1f, 0x0e, 0x3a, 0x69, 0x7d, 0x23, 0xef, 0x01, 0x14, 0xcb,
	0x0b, 0xe1, 0xcb, 0xa2, 0x46, 0x3f, 0x18, 0xec, 0x3e, 0xdf, 0x4f, 0xae, 0x19, 0xca, 0x77, 0x95,
	0xf8, 0xae, 0x92, 0x4f, 0xcb, 0x8a, 0x74, 0xad, 0x9a, 0x1c, 0x43, 0xdb, 0x48, 0x8d, 0x85, 0xc8,
	0xa3, 0xd0, 0x0b, 0x3d, 0xbd, 0x49, 0xe8, 0x73, 0x95, 0x9e, 0x2e, 0xea, 0xc8, 0x3d, 0xd8, 0x31,
	0xc8, 0x34, 0x46, 0xcd, 0x7e, 0x30, 0xe8, 0xa4, 0xd5, 0x85, 0xdc, 0x81, 0x90, 0x8b, 0x2c, 0xda,
	0xf1, 0x31, 0x77, 0x8c, 0x7b, 0xd0, 0x74, 0x73, 0x91, 0x87, 0xd0, 0x76, 0x34, 0x86, 0x45, 0x16,
	0x05, 0xfe, 0xdf, 0x96, 0xbb, 0x9e, 0x64, 0xb1, 0x82, 0xb6, 0x4b, 0x48, 0xad, 0xd8, 0x98, 0x43,
	0xee, 0x43, 0x4b, 0x5b, 0xe1, 0xe2, 0x8d, 0xea, 0x35, 0x6d, 0xc5, 0x49, 0x46, 0x8e, 0xe0, 0x16,
	0x17, 0xd9, 0xd0, 0x51, 0xaa, 0xe7, 0xe8, 0x26, 0x15, 0xc2, 0x64, 0x81, 0x30, 0x39, 0x5b, 0x20,
	0x4c, 0xdb, 0x5c, 0x64, 0xee, 0x16, 0xff, 0x08, 0xa0, 0x99, 0x5a, 0x61, 0x36, 0xbf, 0xb7, 0x62,
	0xd0, 0xd8, 0xc2, 0x20, 0xfc, 0x2f, 0x06, 0x7f, 0x6b, 0xe0, 0xb7, 0x00, 0xc0, 0x19, 0xf4, 0xee,
	0x42, 0x49, 0x8d, 0x1b, 0xd7, 0x63, 0x0d, 0x69, 0xe3, 0x1f, 0x91, 0xf6, 0x60, 0x57, 0x5a, 0x54,
	0x16, 0x87, 0x38, 0x57, 0x95, 0xa3, 0x9d, 0x14, 0xaa, 0xd0, 0xd9, 0x5c, 0xf1, 0xf8, 0x67, 0x00,
	0x7b, 0x29, 0x77, 0x6d, 0xdc, 0xd0, 0xcc, 0x15, 0xa5, 0xc6, 0x55, 0xa5, 0x75, 0xe7, 0xc3, 0x4b,
	0xce, 0x3f, 0x5b, 0x77, 0x65, 0x3b, 0xcf, 0xda, 0xb1, 0x83, 0x95, 0x63, 0xdb, 0xf3, 0xbd, 0x9b,
	0x1f, 0xe0, 0xd1, 0x69, 0x61, 0x4c, 0x21, 0x72, 0xe7, 0xe9, 0x5b, 0xab, 0x3d, 0x8c, 0x37, 0xd2,
	0x0a, 0x34, 0xe4, 0x31, 0x74, 0xb2, 0x3a, 0x62, 0xea, 0x89, 0x56, 0x81, 0x57, 0x77, 0xbf, 0x1f,
	0xdf, 0x86, 0xbd, 0x5f, 0xc1, 0x2a, 0x14, 0x4f, 0x81, 0x38, 0x99, 0x53, 0x8e, 0x2c, 0x63, 0xc8,
	0x6a, 0x19, 0x02, 0x4d, 0x3f, 0x76, 0xa5, 0xe0, 0xcf, 0x1b, 0x37, 0x6a, 0xb9, 0x05, 0xe1, 0x35,
	0x5b, 0xd0, 0x5c, 0x6e, 0xc1, 0xeb, 0x97, 0x5f, 0x8e, 0xf2, 0x02, 0xa7, 0x76, 0xe4, 0x18, 0x52,
	0x47, 0x96, 0x2e, 0xc8, 0xd2, 0x6d, 0x9f, 0xc0, 0x51, 0xcb, 0x3b, 0xf1, 0xe2, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xea, 0x1f, 0x24, 0xb8, 0x29, 0x05, 0x00, 0x00,
}
