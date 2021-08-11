// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/enterprise/networking/v1beta1/virtual_host.proto

package v1beta1

import (
	reflect "reflect"
	sync "sync"

	_ "cuelang.org/go/encoding/protobuf/cue"
	proto "github.com/golang/protobuf/proto"
	v11 "github.com/solo-io/gloo-mesh/pkg/api/common.mesh.gloo.solo.io/v1"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v12 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//
//A `VirtualHost` is used to configure routes. It is selected by a `VirtualGateway`, and may be attached
//to more than one gateway. The `VirtualHost` contains the top-level configuration and route options, such
//as domains to match against, and any options to be shared by its routes. Routes can send traffic directly
//to a service, or can delegate to a `RouteTable` to perform further routing decisions.
type VirtualHostSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of domains (i.e.: matching the `Host` header of a request) that belong to this virtual host.
	// Note that the wildcard will not match the empty string. e.g. “*-bar.foo.com” will match “baz-bar.foo.com”
	// but not “-bar.foo.com”. Additionally, a special entry “*” is allowed which will match any host/authority header.
	// Only a single virtual host on a gateway can match on “*”. A domain must be unique across all
	// virtual hosts on a gateway or the config will be invalidated by Gloo
	// Domains on virtual hosts obey the same rules as [Envoy Virtual Hosts](https://github.com/envoyproxy/envoy/blob/master/api/envoy/api/v2/route/route.proto)
	Domains []string `protobuf:"bytes,1,rep,name=domains,proto3" json:"domains,omitempty"`
	// The list of HTTP routes define routing actions to be taken for incoming HTTP requests whose host header matches
	// this virtual host. If the request matches more than one route in the list, the first route matched will be selected.
	// If the list of routes is empty, the virtual host will be ignored by Gloo.
	Routes []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	// Route table options contain additional configuration to be applied to all traffic served by the route table.
	// Some configuration here can be overridden by Route Options.
	// OutlierDetection and TrafficShift isn't supported on the route level.
	Options *VirtualHostOptions `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *VirtualHostSpec) Reset() {
	*x = VirtualHostSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualHostSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualHostSpec) ProtoMessage() {}

func (x *VirtualHostSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualHostSpec.ProtoReflect.Descriptor instead.
func (*VirtualHostSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_rawDescGZIP(), []int{0}
}

func (x *VirtualHostSpec) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *VirtualHostSpec) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *VirtualHostSpec) GetOptions() *VirtualHostOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type VirtualHostOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrafficPolicy *v1.TrafficPolicySpec_Policy `protobuf:"bytes,1,opt,name=traffic_policy,json=trafficPolicy,proto3" json:"traffic_policy,omitempty"`
}

func (x *VirtualHostOptions) Reset() {
	*x = VirtualHostOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualHostOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualHostOptions) ProtoMessage() {}

func (x *VirtualHostOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualHostOptions.ProtoReflect.Descriptor instead.
func (*VirtualHostOptions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_rawDescGZIP(), []int{1}
}

func (x *VirtualHostOptions) GetTrafficPolicy() *v1.TrafficPolicySpec_Policy {
	if x != nil {
		return x.TrafficPolicy
	}
	return nil
}

type VirtualHostStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The most recent generation observed in the the VirtualHost metadata.
	// If the `observedGeneration` does not match `metadata.generation`,
	// Gloo Mesh has not processed the most recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The state of the overall resource.
	State v11.ApprovalState `protobuf:"varint,7,opt,name=state,proto3,enum=common.mesh.gloo.solo.io.ApprovalState" json:"state,omitempty"`
	// Any errors found while processing this generation of the resource.
	Errors []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	// Any warnings found while processing this generation of the resource.
	Warnings []string `protobuf:"bytes,3,rep,name=warnings,proto3" json:"warnings,omitempty"`
	// List of each VirtualGateway which has selected this VirtualHost
	AttachedVirtualGateways []*v12.ObjectRef `protobuf:"bytes,4,rep,name=attached_virtual_gateways,json=attachedVirtualGateways,proto3" json:"attached_virtual_gateways,omitempty"`
	// List of RouteTables that this Route table delegates to
	SelectedRouteTables []*v12.ObjectRef `protobuf:"bytes,5,rep,name=selected_route_tables,json=selectedRouteTables,proto3" json:"selected_route_tables,omitempty"`
	// The set of TrafficPolicies that have been applied to this Destination.
	// {{/* Note: validation of this field disabled because it slows down cue tremendously*/}}
	AppliedTrafficPolicies []*v1.AppliedTrafficPolicy `protobuf:"bytes,6,rep,name=applied_traffic_policies,json=appliedTrafficPolicies,proto3" json:"applied_traffic_policies,omitempty"`
}

func (x *VirtualHostStatus) Reset() {
	*x = VirtualHostStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualHostStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualHostStatus) ProtoMessage() {}

func (x *VirtualHostStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualHostStatus.ProtoReflect.Descriptor instead.
func (*VirtualHostStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_rawDescGZIP(), []int{2}
}

func (x *VirtualHostStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *VirtualHostStatus) GetState() v11.ApprovalState {
	if x != nil {
		return x.State
	}
	return v11.ApprovalState_PENDING
}

func (x *VirtualHostStatus) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *VirtualHostStatus) GetWarnings() []string {
	if x != nil {
		return x.Warnings
	}
	return nil
}

func (x *VirtualHostStatus) GetAttachedVirtualGateways() []*v12.ObjectRef {
	if x != nil {
		return x.AttachedVirtualGateways
	}
	return nil
}

func (x *VirtualHostStatus) GetSelectedRouteTables() []*v12.ObjectRef {
	if x != nil {
		return x.SelectedRouteTables
	}
	return nil
}

func (x *VirtualHostStatus) GetAppliedTrafficPolicies() []*v1.AppliedTrafficPolicy {
	if x != nil {
		return x.AppliedTrafficPolicies
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_rawDesc = []byte{
	0x0a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x27, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d,
	0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x73, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d,
	0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65,
	0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x75, 0x65, 0x2f, 0x63, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x0f, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f,
	0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73,
	0x12, 0x46, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x73, 0x0a, 0x12, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5d, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x22, 0xd8, 0x03, 0x0a, 0x11, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x48, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x58,
	0x0a, 0x19, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52,
	0x17, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x12, 0x50, 0x0a, 0x15, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73,
	0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x13, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x73, 0x0a, 0x18, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x64, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x42, 0x05, 0xea, 0x42, 0x02, 0x10, 0x01, 0x52, 0x16, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x42,
	0x5a, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_goTypes = []interface{}{
	(*VirtualHostSpec)(nil),             // 0: networking.enterprise.mesh.gloo.solo.io.VirtualHostSpec
	(*VirtualHostOptions)(nil),          // 1: networking.enterprise.mesh.gloo.solo.io.VirtualHostOptions
	(*VirtualHostStatus)(nil),           // 2: networking.enterprise.mesh.gloo.solo.io.VirtualHostStatus
	(*Route)(nil),                       // 3: networking.enterprise.mesh.gloo.solo.io.Route
	(*v1.TrafficPolicySpec_Policy)(nil), // 4: networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy
	(v11.ApprovalState)(0),              // 5: common.mesh.gloo.solo.io.ApprovalState
	(*v12.ObjectRef)(nil),               // 6: core.skv2.solo.io.ObjectRef
	(*v1.AppliedTrafficPolicy)(nil),     // 7: networking.mesh.gloo.solo.io.AppliedTrafficPolicy
}
var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_depIdxs = []int32{
	3, // 0: networking.enterprise.mesh.gloo.solo.io.VirtualHostSpec.routes:type_name -> networking.enterprise.mesh.gloo.solo.io.Route
	1, // 1: networking.enterprise.mesh.gloo.solo.io.VirtualHostSpec.options:type_name -> networking.enterprise.mesh.gloo.solo.io.VirtualHostOptions
	4, // 2: networking.enterprise.mesh.gloo.solo.io.VirtualHostOptions.traffic_policy:type_name -> networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy
	5, // 3: networking.enterprise.mesh.gloo.solo.io.VirtualHostStatus.state:type_name -> common.mesh.gloo.solo.io.ApprovalState
	6, // 4: networking.enterprise.mesh.gloo.solo.io.VirtualHostStatus.attached_virtual_gateways:type_name -> core.skv2.solo.io.ObjectRef
	6, // 5: networking.enterprise.mesh.gloo.solo.io.VirtualHostStatus.selected_route_tables:type_name -> core.skv2.solo.io.ObjectRef
	7, // 6: networking.enterprise.mesh.gloo.solo.io.VirtualHostStatus.applied_traffic_policies:type_name -> networking.mesh.gloo.solo.io.AppliedTrafficPolicy
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_init()
}
func file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualHostSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualHostOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualHostStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_virtual_host_proto_depIdxs = nil
}
