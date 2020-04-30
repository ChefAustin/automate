// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/teams.proto

package v2

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2/request"
	response "github.com/chef/automate/components/automate-gateway/api/iam/v2/response"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/teams.proto", fileDescriptor_16a78463b445c3d8)
}

var fileDescriptor_16a78463b445c3d8 = []byte{
	// 924 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xd6, 0x5a, 0x6d, 0xaa, 0x0e, 0x22, 0x24, 0x43, 0xd4, 0x86, 0x55, 0x51, 0xdd, 0x51, 0x69,
	0x8b, 0xc9, 0xee, 0x92, 0x8d, 0xed, 0x8d, 0xcd, 0x05, 0x18, 0x10, 0x11, 0x88, 0x52, 0x68, 0xd2,
	0x1b, 0x2a, 0x84, 0xc6, 0xbb, 0xc7, 0x9b, 0x2d, 0xde, 0x9d, 0xc9, 0xce, 0x38, 0x21, 0x8a, 0x22,
	0x21, 0x7e, 0x8a, 0xe8, 0x0d, 0x52, 0xb8, 0xe5, 0x16, 0xc4, 0x03, 0x50, 0xc1, 0x0b, 0xf0, 0x04,
	0x08, 0x78, 0x01, 0x2e, 0xb9, 0xe2, 0x01, 0x10, 0x9a, 0x75, 0xb2, 0x3f, 0xd8, 0x1b, 0x6f, 0x6e,
	0xb9, 0xb2, 0x66, 0xe6, 0x3b, 0x67, 0xbe, 0xf3, 0x9d, 0xef, 0x78, 0x07, 0xd9, 0x2e, 0x0b, 0x39,
	0x8b, 0x20, 0x92, 0xc2, 0xa2, 0x23, 0xc9, 0x42, 0x2a, 0xc1, 0xf0, 0xa9, 0x84, 0x3d, 0xba, 0x6f,
	0x51, 0x1e, 0x58, 0x01, 0x0d, 0xad, 0x5d, 0xdb, 0x92, 0x40, 0x43, 0x61, 0xf2, 0x98, 0x49, 0x86,
	0x97, 0xdd, 0x6d, 0x18, 0x98, 0x27, 0x68, 0x93, 0xf2, 0xc0, 0x0c, 0x68, 0x68, 0xee, 0xda, 0xfa,
	0x15, 0x9f, 0x31, 0x7f, 0x08, 0x49, 0x20, 0x8d, 0x22, 0x26, 0xa9, 0x0c, 0x58, 0x74, 0x1c, 0xa7,
	0xaf, 0x24, 0x3f, 0xae, 0xe1, 0x43, 0x64, 0x88, 0x3d, 0xea, 0xfb, 0x10, 0x5b, 0x8c, 0x27, 0x88,
	0x29, 0xe8, 0x6e, 0x45, 0x66, 0x31, 0xec, 0x8c, 0x40, 0xc8, 0x3c, 0x43, 0xfd, 0xa5, 0xca, 0xb1,
	0x82, 0xb3, 0x48, 0x40, 0x21, 0xf8, 0x95, 0xa9, 0xc1, 0x31, 0x77, 0xad, 0x1c, 0x7f, 0xce, 0x86,
	0x81, 0xbb, 0x9f, 0xe4, 0x99, 0xa0, 0x6e, 0xff, 0x82, 0xd1, 0xf9, 0x2d, 0x95, 0x11, 0xff, 0x54,
	0x43, 0xe8, 0xb5, 0x18, 0xa8, 0x04, 0xb5, 0xc6, 0x37, 0xcd, 0x32, 0xe9, 0xcc, 0x0c, 0x75, 0x17,
	0x76, 0xf4, 0x5b, 0xd5, 0x80, 0x82, 0x93, 0xdf, 0xb4, 0xa3, 0xde, 0xd7, 0x1a, 0x3a, 0x9f, 0x54,
	0xf0, 0xe0, 0x0b, 0x0d, 0xcd, 0x7f, 0x6c, 0xb8, 0xcc, 0x03, 0x43, 0xd0, 0x90, 0x0f, 0x41, 0x60,
	0x61, 0xef, 0x20, 0xd6, 0x08, 0xd1, 0x3c, 0x3a, 0x37, 0xa4, 0x91, 0x8f, 0xe7, 0xf4, 0x73, 0x6f,
	0x6d, 0xde, 0x79, 0x07, 0xdd, 0x47, 0x73, 0x82, 0x8d, 0x62, 0x17, 0xf0, 0x7b, 0xfa, 0x9d, 0x03,
	0x12, 0x78, 0xa4, 0x5b, 0x27, 0x12, 0x84, 0x34, 0x02, 0x8f, 0xac, 0xd4, 0x49, 0x44, 0x43, 0x50,
	0x5b, 0xb7, 0xf7, 0xeb, 0x5b, 0x20, 0x64, 0x5d, 0xdd, 0xa8, 0xf6, 0x79, 0xcc, 0x1e, 0x80, 0x2b,
	0x05, 0xe9, 0xd6, 0xef, 0x9f, 0x2c, 0x56, 0x73, 0x07, 0x36, 0xf9, 0xe0, 0xf0, 0xd1, 0xe3, 0xe5,
	0x27, 0xd0, 0xc5, 0x80, 0x86, 0xdd, 0x84, 0xd6, 0xa3, 0xc7, 0xcb, 0x18, 0x2f, 0xa4, 0xcb, 0xae,
	0x9b, 0x94, 0xf0, 0xe9, 0xaf, 0x7f, 0x7e, 0x53, 0xbb, 0x4c, 0xb0, 0x6a, 0x88, 0x28, 0xf8, 0xac,
	0xab, 0x35, 0xf0, 0x77, 0x1a, 0xba, 0xf8, 0x76, 0x20, 0xe4, 0x58, 0xc7, 0x1b, 0xe5, 0x7a, 0xa4,
	0x20, 0xa5, 0xdb, 0xcd, 0x4a, 0x38, 0xc1, 0xc9, 0xc6, 0x51, 0xef, 0xc2, 0xb1, 0x68, 0x93, 0x64,
	0x17, 0xf0, 0x7c, 0x46, 0x76, 0x18, 0x08, 0x99, 0x50, 0x5d, 0xc2, 0x53, 0xa8, 0xe2, 0x1f, 0x34,
	0x74, 0x61, 0x03, 0x92, 0xcc, 0xf8, 0x7a, 0xf9, 0xed, 0xc7, 0x10, 0xc5, 0xf1, 0xb9, 0x0a, 0x28,
	0xc1, 0xc9, 0xbb, 0x05, 0x86, 0x0b, 0x28, 0x47, 0xe9, 0x20, 0xf0, 0x94, 0xc4, 0x4f, 0xe1, 0x27,
	0xb3, 0x3d, 0x1f, 0xc6, 0x2c, 0x9f, 0xc1, 0x97, 0x27, 0x59, 0x5a, 0x2a, 0x02, 0xff, 0x5c, 0x43,
	0xe8, 0x1e, 0xf7, 0x2a, 0x98, 0x31, 0x43, 0xcd, 0x30, 0x63, 0x1e, 0x28, 0x38, 0xf9, 0x5d, 0x3b,
	0xea, 0x3d, 0x4c, 0xcd, 0x78, 0x38, 0xe1, 0xc5, 0x8f, 0xec, 0x00, 0xf9, 0x0d, 0x98, 0xf0, 0xe2,
	0x66, 0xea, 0xc5, 0x37, 0xf5, 0x8d, 0x83, 0xbc, 0xf5, 0xc6, 0x17, 0x9c, 0xcd, 0x81, 0x11, 0xec,
	0x8d, 0x3d, 0x38, 0x4d, 0xb4, 0x82, 0x11, 0x47, 0x49, 0xf6, 0x44, 0xb7, 0x2b, 0x7a, 0x99, 0x6e,
	0xca, 0x8d, 0x3f, 0x6a, 0x08, 0xbd, 0x0e, 0x43, 0x98, 0x2d, 0x5d, 0x86, 0x9a, 0x21, 0x5d, 0x1e,
	0x28, 0x38, 0xd9, 0x9c, 0xd9, 0xee, 0x02, 0x73, 0x2f, 0x89, 0x1e, 0x77, 0xbc, 0x51, 0xda, 0xf1,
	0x3f, 0x34, 0xb4, 0x78, 0xec, 0xa9, 0xdb, 0x10, 0xf6, 0x21, 0x16, 0xdb, 0x01, 0xc7, 0xe6, 0x4c,
	0x03, 0x66, 0x60, 0x55, 0x84, 0x75, 0x26, 0xbc, 0xe0, 0x84, 0x16, 0x6a, 0xb9, 0x84, 0x96, 0x8a,
	0xb5, 0x74, 0x47, 0x02, 0x62, 0x75, 0xb2, 0x84, 0xf1, 0xc9, 0xc9, 0x3d, 0xb5, 0x95, 0xcd, 0xda,
	0x55, 0xfc, 0x6c, 0x49, 0x4d, 0x56, 0x12, 0x8c, 0xff, 0xae, 0xa1, 0xf9, 0x9e, 0xe7, 0xe5, 0x2e,
	0xc7, 0x2f, 0x94, 0xd3, 0x2c, 0x22, 0x55, 0x4d, 0x2b, 0xd5, 0xc1, 0x82, 0x93, 0xcf, 0x6b, 0x47,
	0xbd, 0xef, 0x53, 0x5f, 0x7f, 0x3b, 0xf9, 0x27, 0xfb, 0x95, 0x66, 0x3f, 0xd4, 0xd0, 0x67, 0x5a,
	0xe3, 0x13, 0x6d, 0xc2, 0xdc, 0x51, 0x6a, 0x6e, 0x4f, 0xef, 0x1f, 0x84, 0xa9, 0x54, 0x1f, 0x06,
	0x9e, 0x50, 0x0e, 0x6e, 0xd9, 0x0e, 0x78, 0x9d, 0xf6, 0xc0, 0xb0, 0xc1, 0xed, 0x1b, 0xcd, 0xc1,
	0xfa, 0xc0, 0xa0, 0x7d, 0xcf, 0x31, 0x5e, 0xec, 0x0f, 0xda, 0xad, 0xd5, 0xd5, 0x66, 0xab, 0x43,
	0x5d, 0xe5, 0xee, 0xce, 0xba, 0xe3, 0xae, 0x37, 0x9d, 0x96, 0xd1, 0x72, 0x9a, 0x8e, 0xd1, 0x1c,
	0x74, 0xfa, 0x06, 0x75, 0xda, 0x6d, 0xc3, 0x5d, 0x5b, 0x73, 0x06, 0xce, 0x5a, 0xa7, 0xdd, 0xa2,
	0x50, 0xee, 0xfb, 0x4b, 0x78, 0xa9, 0xa8, 0x75, 0xee, 0x4f, 0xf8, 0x06, 0xb9, 0x76, 0xaa, 0xda,
	0x5d, 0xea, 0x79, 0x6a, 0x0a, 0xfe, 0xa9, 0xa1, 0xc5, 0xbb, 0x10, 0xb2, 0x5d, 0xc8, 0xeb, 0x7e,
	0x8a, 0x9d, 0x26, 0xc0, 0x33, 0xec, 0x34, 0x05, 0x2f, 0x38, 0xf9, 0xf2, 0x7f, 0xa8, 0x7e, 0x6e,
	0x7e, 0x9f, 0x27, 0xd7, 0x4f, 0x57, 0x3f, 0x4e, 0x54, 0x51, 0x0d, 0xf8, 0x2b, 0x9b, 0x67, 0xf1,
	0x06, 0x8b, 0xc7, 0x22, 0x55, 0x98, 0xe7, 0x0c, 0x5c, 0x6d, 0x9e, 0xf3, 0x78, 0xc1, 0xc9, 0x6e,
	0x61, 0x9e, 0xaf, 0xa1, 0xab, 0xaa, 0x96, 0x31, 0xb7, 0xa2, 0xa2, 0x87, 0xe9, 0x27, 0xf4, 0x69,
	0xbc, 0x78, 0x02, 0xda, 0x2a, 0x7c, 0x9f, 0x1a, 0xf8, 0x56, 0xa1, 0xda, 0x24, 0x8d, 0xf5, 0x9f,
	0x34, 0x63, 0x0d, 0x5e, 0xed, 0xbd, 0xff, 0xb2, 0x1f, 0xc8, 0xed, 0x51, 0xdf, 0x74, 0x59, 0x68,
	0x29, 0xd2, 0xe9, 0x83, 0xcc, 0xaa, 0xf6, 0xc2, 0xeb, 0xcf, 0x25, 0x2f, 0xb2, 0xb5, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xaf, 0x23, 0x68, 0x4a, 0xe8, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TeamsClient is the client API for Teams service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TeamsClient interface {
	//
	//Creates a local team
	//
	//Creates a local team that is used to group local users as members of IAM policies.
	//
	//Authorization Action:
	//```
	//iam:teams:create
	//```
	CreateTeam(ctx context.Context, in *request.CreateTeamReq, opts ...grpc.CallOption) (*response.CreateTeamResp, error)
	//
	//Lists all local teams
	//
	//Lists all local teams.
	//
	//Authorization Action:
	//```
	//iam:teams:list
	//```
	ListTeams(ctx context.Context, in *request.ListTeamsReq, opts ...grpc.CallOption) (*response.ListTeamsResp, error)
	//
	//Get a team
	//
	//Returns the details for a team.
	//
	//Authorization Action:
	//```
	//iam:teams:get
	//```
	GetTeam(ctx context.Context, in *request.GetTeamReq, opts ...grpc.CallOption) (*response.GetTeamResp, error)
	//
	//Updates a local team
	//
	//Updates a local team.
	//
	//Authorization Action:
	//```
	//iam:teams:update
	//```
	UpdateTeam(ctx context.Context, in *request.UpdateTeamReq, opts ...grpc.CallOption) (*response.UpdateTeamResp, error)
	//
	//Deletes a local team
	//
	//Deletes a local team and removes it from any policies.
	//
	//Authorization Action:
	//```
	//iam:teams:delete
	//```
	DeleteTeam(ctx context.Context, in *request.DeleteTeamReq, opts ...grpc.CallOption) (*response.DeleteTeamResp, error)
	//
	//Gets local team membership
	//
	//Lists all users of a local team. Users are listed by their membership_id.
	//
	//Authorization Action:
	//```
	//iam:teamUsers:list
	//```
	GetTeamMembership(ctx context.Context, in *request.GetTeamMembershipReq, opts ...grpc.CallOption) (*response.GetTeamMembershipResp, error)
	//
	//Adds local team membership
	//
	//Adds a list of users to a local team. Users are added by their membership_id.
	//The request currently does not validate that membership_id maps to a real user.
	//
	//The membership_id for users can be found via GET /apis/apis/iam/v2/users/<user_id>.
	//
	//Authorization Action:
	//```
	//iam:teamUsers:create
	//```
	AddTeamMembers(ctx context.Context, in *request.AddTeamMembersReq, opts ...grpc.CallOption) (*response.AddTeamMembersResp, error)
	//
	//Removes local team membership
	//
	//Removes a list of users from a local team. Users are removed by their membership_id.
	//The request currently does not validate that membership_id maps to a real user.
	//
	//The membership_id for users can be found via GET /apis/apis/iam/v2/users/<user_id>.
	//
	//Authorization Action:
	//```
	//iam:teamUsers:delete
	//```
	RemoveTeamMembers(ctx context.Context, in *request.RemoveTeamMembersReq, opts ...grpc.CallOption) (*response.RemoveTeamMembersResp, error)
	//
	//Gets team membership for a user
	//
	//Lists all local teams for a specific user. You must use their membership_id in the request URL.
	//
	//Authorization Action:
	//```
	//iam:userTeams:get
	//```
	GetTeamsForMember(ctx context.Context, in *request.GetTeamsForMemberReq, opts ...grpc.CallOption) (*response.GetTeamsForMemberResp, error)
}

type teamsClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamsClient(cc grpc.ClientConnInterface) TeamsClient {
	return &teamsClient{cc}
}

func (c *teamsClient) CreateTeam(ctx context.Context, in *request.CreateTeamReq, opts ...grpc.CallOption) (*response.CreateTeamResp, error) {
	out := new(response.CreateTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/CreateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) ListTeams(ctx context.Context, in *request.ListTeamsReq, opts ...grpc.CallOption) (*response.ListTeamsResp, error) {
	out := new(response.ListTeamsResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/ListTeams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) GetTeam(ctx context.Context, in *request.GetTeamReq, opts ...grpc.CallOption) (*response.GetTeamResp, error) {
	out := new(response.GetTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/GetTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) UpdateTeam(ctx context.Context, in *request.UpdateTeamReq, opts ...grpc.CallOption) (*response.UpdateTeamResp, error) {
	out := new(response.UpdateTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/UpdateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) DeleteTeam(ctx context.Context, in *request.DeleteTeamReq, opts ...grpc.CallOption) (*response.DeleteTeamResp, error) {
	out := new(response.DeleteTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/DeleteTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) GetTeamMembership(ctx context.Context, in *request.GetTeamMembershipReq, opts ...grpc.CallOption) (*response.GetTeamMembershipResp, error) {
	out := new(response.GetTeamMembershipResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/GetTeamMembership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) AddTeamMembers(ctx context.Context, in *request.AddTeamMembersReq, opts ...grpc.CallOption) (*response.AddTeamMembersResp, error) {
	out := new(response.AddTeamMembersResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/AddTeamMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) RemoveTeamMembers(ctx context.Context, in *request.RemoveTeamMembersReq, opts ...grpc.CallOption) (*response.RemoveTeamMembersResp, error) {
	out := new(response.RemoveTeamMembersResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/RemoveTeamMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) GetTeamsForMember(ctx context.Context, in *request.GetTeamsForMemberReq, opts ...grpc.CallOption) (*response.GetTeamsForMemberResp, error) {
	out := new(response.GetTeamsForMemberResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/GetTeamsForMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamsServer is the server API for Teams service.
type TeamsServer interface {
	//
	//Creates a local team
	//
	//Creates a local team that is used to group local users as members of IAM policies.
	//
	//Authorization Action:
	//```
	//iam:teams:create
	//```
	CreateTeam(context.Context, *request.CreateTeamReq) (*response.CreateTeamResp, error)
	//
	//Lists all local teams
	//
	//Lists all local teams.
	//
	//Authorization Action:
	//```
	//iam:teams:list
	//```
	ListTeams(context.Context, *request.ListTeamsReq) (*response.ListTeamsResp, error)
	//
	//Get a team
	//
	//Returns the details for a team.
	//
	//Authorization Action:
	//```
	//iam:teams:get
	//```
	GetTeam(context.Context, *request.GetTeamReq) (*response.GetTeamResp, error)
	//
	//Updates a local team
	//
	//Updates a local team.
	//
	//Authorization Action:
	//```
	//iam:teams:update
	//```
	UpdateTeam(context.Context, *request.UpdateTeamReq) (*response.UpdateTeamResp, error)
	//
	//Deletes a local team
	//
	//Deletes a local team and removes it from any policies.
	//
	//Authorization Action:
	//```
	//iam:teams:delete
	//```
	DeleteTeam(context.Context, *request.DeleteTeamReq) (*response.DeleteTeamResp, error)
	//
	//Gets local team membership
	//
	//Lists all users of a local team. Users are listed by their membership_id.
	//
	//Authorization Action:
	//```
	//iam:teamUsers:list
	//```
	GetTeamMembership(context.Context, *request.GetTeamMembershipReq) (*response.GetTeamMembershipResp, error)
	//
	//Adds local team membership
	//
	//Adds a list of users to a local team. Users are added by their membership_id.
	//The request currently does not validate that membership_id maps to a real user.
	//
	//The membership_id for users can be found via GET /apis/apis/iam/v2/users/<user_id>.
	//
	//Authorization Action:
	//```
	//iam:teamUsers:create
	//```
	AddTeamMembers(context.Context, *request.AddTeamMembersReq) (*response.AddTeamMembersResp, error)
	//
	//Removes local team membership
	//
	//Removes a list of users from a local team. Users are removed by their membership_id.
	//The request currently does not validate that membership_id maps to a real user.
	//
	//The membership_id for users can be found via GET /apis/apis/iam/v2/users/<user_id>.
	//
	//Authorization Action:
	//```
	//iam:teamUsers:delete
	//```
	RemoveTeamMembers(context.Context, *request.RemoveTeamMembersReq) (*response.RemoveTeamMembersResp, error)
	//
	//Gets team membership for a user
	//
	//Lists all local teams for a specific user. You must use their membership_id in the request URL.
	//
	//Authorization Action:
	//```
	//iam:userTeams:get
	//```
	GetTeamsForMember(context.Context, *request.GetTeamsForMemberReq) (*response.GetTeamsForMemberResp, error)
}

// UnimplementedTeamsServer can be embedded to have forward compatible implementations.
type UnimplementedTeamsServer struct {
}

func (*UnimplementedTeamsServer) CreateTeam(ctx context.Context, req *request.CreateTeamReq) (*response.CreateTeamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeam not implemented")
}
func (*UnimplementedTeamsServer) ListTeams(ctx context.Context, req *request.ListTeamsReq) (*response.ListTeamsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTeams not implemented")
}
func (*UnimplementedTeamsServer) GetTeam(ctx context.Context, req *request.GetTeamReq) (*response.GetTeamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeam not implemented")
}
func (*UnimplementedTeamsServer) UpdateTeam(ctx context.Context, req *request.UpdateTeamReq) (*response.UpdateTeamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeam not implemented")
}
func (*UnimplementedTeamsServer) DeleteTeam(ctx context.Context, req *request.DeleteTeamReq) (*response.DeleteTeamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeam not implemented")
}
func (*UnimplementedTeamsServer) GetTeamMembership(ctx context.Context, req *request.GetTeamMembershipReq) (*response.GetTeamMembershipResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamMembership not implemented")
}
func (*UnimplementedTeamsServer) AddTeamMembers(ctx context.Context, req *request.AddTeamMembersReq) (*response.AddTeamMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTeamMembers not implemented")
}
func (*UnimplementedTeamsServer) RemoveTeamMembers(ctx context.Context, req *request.RemoveTeamMembersReq) (*response.RemoveTeamMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTeamMembers not implemented")
}
func (*UnimplementedTeamsServer) GetTeamsForMember(ctx context.Context, req *request.GetTeamsForMemberReq) (*response.GetTeamsForMemberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamsForMember not implemented")
}

func RegisterTeamsServer(s *grpc.Server, srv TeamsServer) {
	s.RegisterService(&_Teams_serviceDesc, srv)
}

func _Teams_CreateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).CreateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/CreateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).CreateTeam(ctx, req.(*request.CreateTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_ListTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListTeamsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).ListTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/ListTeams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).ListTeams(ctx, req.(*request.ListTeamsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_GetTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/GetTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetTeam(ctx, req.(*request.GetTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_UpdateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).UpdateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/UpdateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).UpdateTeam(ctx, req.(*request.UpdateTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_DeleteTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).DeleteTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/DeleteTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).DeleteTeam(ctx, req.(*request.DeleteTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_GetTeamMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetTeamMembershipReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetTeamMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/GetTeamMembership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetTeamMembership(ctx, req.(*request.GetTeamMembershipReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_AddTeamMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.AddTeamMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).AddTeamMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/AddTeamMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).AddTeamMembers(ctx, req.(*request.AddTeamMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_RemoveTeamMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.RemoveTeamMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).RemoveTeamMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/RemoveTeamMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).RemoveTeamMembers(ctx, req.(*request.RemoveTeamMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_GetTeamsForMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetTeamsForMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetTeamsForMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/GetTeamsForMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetTeamsForMember(ctx, req.(*request.GetTeamsForMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Teams_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.iam.v2.Teams",
	HandlerType: (*TeamsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTeam",
			Handler:    _Teams_CreateTeam_Handler,
		},
		{
			MethodName: "ListTeams",
			Handler:    _Teams_ListTeams_Handler,
		},
		{
			MethodName: "GetTeam",
			Handler:    _Teams_GetTeam_Handler,
		},
		{
			MethodName: "UpdateTeam",
			Handler:    _Teams_UpdateTeam_Handler,
		},
		{
			MethodName: "DeleteTeam",
			Handler:    _Teams_DeleteTeam_Handler,
		},
		{
			MethodName: "GetTeamMembership",
			Handler:    _Teams_GetTeamMembership_Handler,
		},
		{
			MethodName: "AddTeamMembers",
			Handler:    _Teams_AddTeamMembers_Handler,
		},
		{
			MethodName: "RemoveTeamMembers",
			Handler:    _Teams_RemoveTeamMembers_Handler,
		},
		{
			MethodName: "GetTeamsForMember",
			Handler:    _Teams_GetTeamsForMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/iam/v2/teams.proto",
}
