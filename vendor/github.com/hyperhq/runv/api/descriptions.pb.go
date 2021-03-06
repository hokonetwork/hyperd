// Code generated by protoc-gen-go. DO NOT EDIT.
// source: descriptions.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	descriptions.proto

It has these top-level messages:
	SandboxConfig
	ContainerDescription
	VolumeDescription
	InterfaceDescription
	PortDescription
	NeighborNetworks
	VolumeReference
	VolumeMount
	VolumeOption
	UserGroupInfo
	Rlimit
	Process
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SandboxConfig struct {
	Hostname   string            `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Dns        []string          `protobuf:"bytes,2,rep,name=dns" json:"dns,omitempty"`
	Neighbors  *NeighborNetworks `protobuf:"bytes,3,opt,name=neighbors" json:"neighbors,omitempty"`
	DnsOptions []string          `protobuf:"bytes,4,rep,name=dnsOptions" json:"dnsOptions,omitempty"`
	DnsSearch  []string          `protobuf:"bytes,5,rep,name=dnsSearch" json:"dnsSearch,omitempty"`
}

func (m *SandboxConfig) Reset()                    { *m = SandboxConfig{} }
func (m *SandboxConfig) String() string            { return proto.CompactTextString(m) }
func (*SandboxConfig) ProtoMessage()               {}
func (*SandboxConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SandboxConfig) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *SandboxConfig) GetDns() []string {
	if m != nil {
		return m.Dns
	}
	return nil
}

func (m *SandboxConfig) GetNeighbors() *NeighborNetworks {
	if m != nil {
		return m.Neighbors
	}
	return nil
}

func (m *SandboxConfig) GetDnsOptions() []string {
	if m != nil {
		return m.DnsOptions
	}
	return nil
}

func (m *SandboxConfig) GetDnsSearch() []string {
	if m != nil {
		return m.DnsSearch
	}
	return nil
}

type ContainerDescription struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Static Info, got from client input
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Image string `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	// User content or user specified behavior
	Labels     map[string]string `protobuf:"bytes,4,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Tty        bool              `protobuf:"varint,5,opt,name=tty" json:"tty,omitempty"`
	StopSignal string            `protobuf:"bytes,6,opt,name=stopSignal" json:"stopSignal,omitempty"`
	// Creation Info, got during creation
	RootVolume *VolumeDescription `protobuf:"bytes,7,opt,name=rootVolume" json:"rootVolume,omitempty"`
	MountId    string             `protobuf:"bytes,8,opt,name=mountId" json:"mountId,omitempty"`
	RootPath   string             `protobuf:"bytes,9,opt,name=rootPath" json:"rootPath,omitempty"`
	// runtime info, combined during creation
	UGI        *UserGroupInfo              `protobuf:"bytes,10,opt,name=UGI" json:"UGI,omitempty"`
	Envs       map[string]string           `protobuf:"bytes,11,rep,name=envs" json:"envs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Workdir    string                      `protobuf:"bytes,12,opt,name=workdir" json:"workdir,omitempty"`
	Path       string                      `protobuf:"bytes,13,opt,name=path" json:"path,omitempty"`
	Args       []string                    `protobuf:"bytes,14,rep,name=args" json:"args,omitempty"`
	Rlimits    []*Rlimit                   `protobuf:"bytes,15,rep,name=rlimits" json:"rlimits,omitempty"`
	Sysctl     map[string]string           `protobuf:"bytes,16,rep,name=sysctl" json:"sysctl,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Volumes    map[string]*VolumeReference `protobuf:"bytes,17,rep,name=volumes" json:"volumes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Initialize bool                        `protobuf:"varint,24,opt,name=initialize" json:"initialize,omitempty"`
}

func (m *ContainerDescription) Reset()                    { *m = ContainerDescription{} }
func (m *ContainerDescription) String() string            { return proto.CompactTextString(m) }
func (*ContainerDescription) ProtoMessage()               {}
func (*ContainerDescription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ContainerDescription) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ContainerDescription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerDescription) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ContainerDescription) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ContainerDescription) GetTty() bool {
	if m != nil {
		return m.Tty
	}
	return false
}

func (m *ContainerDescription) GetStopSignal() string {
	if m != nil {
		return m.StopSignal
	}
	return ""
}

func (m *ContainerDescription) GetRootVolume() *VolumeDescription {
	if m != nil {
		return m.RootVolume
	}
	return nil
}

func (m *ContainerDescription) GetMountId() string {
	if m != nil {
		return m.MountId
	}
	return ""
}

func (m *ContainerDescription) GetRootPath() string {
	if m != nil {
		return m.RootPath
	}
	return ""
}

func (m *ContainerDescription) GetUGI() *UserGroupInfo {
	if m != nil {
		return m.UGI
	}
	return nil
}

func (m *ContainerDescription) GetEnvs() map[string]string {
	if m != nil {
		return m.Envs
	}
	return nil
}

func (m *ContainerDescription) GetWorkdir() string {
	if m != nil {
		return m.Workdir
	}
	return ""
}

func (m *ContainerDescription) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ContainerDescription) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *ContainerDescription) GetRlimits() []*Rlimit {
	if m != nil {
		return m.Rlimits
	}
	return nil
}

func (m *ContainerDescription) GetSysctl() map[string]string {
	if m != nil {
		return m.Sysctl
	}
	return nil
}

func (m *ContainerDescription) GetVolumes() map[string]*VolumeReference {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *ContainerDescription) GetInitialize() bool {
	if m != nil {
		return m.Initialize
	}
	return false
}

type VolumeDescription struct {
	Name         string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Source       string        `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	Format       string        `protobuf:"bytes,3,opt,name=format" json:"format,omitempty"`
	Fstype       string        `protobuf:"bytes,4,opt,name=fstype" json:"fstype,omitempty"`
	Cache        string        `protobuf:"bytes,5,opt,name=cache" json:"cache,omitempty"`
	Options      *VolumeOption `protobuf:"bytes,8,opt,name=options" json:"options,omitempty"`
	DockerVolume bool          `protobuf:"varint,9,opt,name=dockerVolume" json:"dockerVolume,omitempty"`
	ReadOnly     bool          `protobuf:"varint,10,opt,name=readOnly" json:"readOnly,omitempty"`
}

func (m *VolumeDescription) Reset()                    { *m = VolumeDescription{} }
func (m *VolumeDescription) String() string            { return proto.CompactTextString(m) }
func (*VolumeDescription) ProtoMessage()               {}
func (*VolumeDescription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VolumeDescription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VolumeDescription) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *VolumeDescription) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *VolumeDescription) GetFstype() string {
	if m != nil {
		return m.Fstype
	}
	return ""
}

func (m *VolumeDescription) GetCache() string {
	if m != nil {
		return m.Cache
	}
	return ""
}

func (m *VolumeDescription) GetOptions() *VolumeOption {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *VolumeDescription) GetDockerVolume() bool {
	if m != nil {
		return m.DockerVolume
	}
	return false
}

func (m *VolumeDescription) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

type InterfaceDescription struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Lo      bool   `protobuf:"varint,3,opt,name=lo" json:"lo,omitempty"`
	Bridge  string `protobuf:"bytes,4,opt,name=bridge" json:"bridge,omitempty"`
	Ip      string `protobuf:"bytes,5,opt,name=ip" json:"ip,omitempty"`
	Mac     string `protobuf:"bytes,6,opt,name=mac" json:"mac,omitempty"`
	Mtu     uint64 `protobuf:"varint,7,opt,name=mtu" json:"mtu,omitempty"`
	Gw      string `protobuf:"bytes,8,opt,name=gw" json:"gw,omitempty"`
	TapName string `protobuf:"bytes,9,opt,name=tapName" json:"tapName,omitempty"`
	Options string `protobuf:"bytes,10,opt,name=options" json:"options,omitempty"`
}

func (m *InterfaceDescription) Reset()                    { *m = InterfaceDescription{} }
func (m *InterfaceDescription) String() string            { return proto.CompactTextString(m) }
func (*InterfaceDescription) ProtoMessage()               {}
func (*InterfaceDescription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *InterfaceDescription) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InterfaceDescription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InterfaceDescription) GetLo() bool {
	if m != nil {
		return m.Lo
	}
	return false
}

func (m *InterfaceDescription) GetBridge() string {
	if m != nil {
		return m.Bridge
	}
	return ""
}

func (m *InterfaceDescription) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *InterfaceDescription) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *InterfaceDescription) GetMtu() uint64 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *InterfaceDescription) GetGw() string {
	if m != nil {
		return m.Gw
	}
	return ""
}

func (m *InterfaceDescription) GetTapName() string {
	if m != nil {
		return m.TapName
	}
	return ""
}

func (m *InterfaceDescription) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

type PortDescription struct {
	HostPort      int32  `protobuf:"varint,1,opt,name=hostPort" json:"hostPort,omitempty"`
	ContainerPort int32  `protobuf:"varint,2,opt,name=containerPort" json:"containerPort,omitempty"`
	Protocol      string `protobuf:"bytes,3,opt,name=protocol" json:"protocol,omitempty"`
}

func (m *PortDescription) Reset()                    { *m = PortDescription{} }
func (m *PortDescription) String() string            { return proto.CompactTextString(m) }
func (*PortDescription) ProtoMessage()               {}
func (*PortDescription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PortDescription) GetHostPort() int32 {
	if m != nil {
		return m.HostPort
	}
	return 0
}

func (m *PortDescription) GetContainerPort() int32 {
	if m != nil {
		return m.ContainerPort
	}
	return 0
}

func (m *PortDescription) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type NeighborNetworks struct {
	InternalNetworks []string `protobuf:"bytes,1,rep,name=internalNetworks" json:"internalNetworks,omitempty"`
	ExternalNetworks []string `protobuf:"bytes,2,rep,name=externalNetworks" json:"externalNetworks,omitempty"`
}

func (m *NeighborNetworks) Reset()                    { *m = NeighborNetworks{} }
func (m *NeighborNetworks) String() string            { return proto.CompactTextString(m) }
func (*NeighborNetworks) ProtoMessage()               {}
func (*NeighborNetworks) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *NeighborNetworks) GetInternalNetworks() []string {
	if m != nil {
		return m.InternalNetworks
	}
	return nil
}

func (m *NeighborNetworks) GetExternalNetworks() []string {
	if m != nil {
		return m.ExternalNetworks
	}
	return nil
}

type VolumeReference struct {
	Name        string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	MountPoints []*VolumeMount `protobuf:"bytes,2,rep,name=mountPoints" json:"mountPoints,omitempty"`
}

func (m *VolumeReference) Reset()                    { *m = VolumeReference{} }
func (m *VolumeReference) String() string            { return proto.CompactTextString(m) }
func (*VolumeReference) ProtoMessage()               {}
func (*VolumeReference) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *VolumeReference) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VolumeReference) GetMountPoints() []*VolumeMount {
	if m != nil {
		return m.MountPoints
	}
	return nil
}

type VolumeMount struct {
	Path     string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	ReadOnly bool   `protobuf:"varint,2,opt,name=readOnly" json:"readOnly,omitempty"`
}

func (m *VolumeMount) Reset()                    { *m = VolumeMount{} }
func (m *VolumeMount) String() string            { return proto.CompactTextString(m) }
func (*VolumeMount) ProtoMessage()               {}
func (*VolumeMount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *VolumeMount) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *VolumeMount) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

type VolumeOption struct {
	User        string   `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Monitors    []string `protobuf:"bytes,2,rep,name=monitors" json:"monitors,omitempty"`
	Keyring     string   `protobuf:"bytes,3,opt,name=keyring" json:"keyring,omitempty"`
	BytesPerSec int32    `protobuf:"varint,4,opt,name=bytesPerSec" json:"bytesPerSec,omitempty"`
	Iops        int32    `protobuf:"varint,5,opt,name=iops" json:"iops,omitempty"`
}

func (m *VolumeOption) Reset()                    { *m = VolumeOption{} }
func (m *VolumeOption) String() string            { return proto.CompactTextString(m) }
func (*VolumeOption) ProtoMessage()               {}
func (*VolumeOption) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *VolumeOption) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *VolumeOption) GetMonitors() []string {
	if m != nil {
		return m.Monitors
	}
	return nil
}

func (m *VolumeOption) GetKeyring() string {
	if m != nil {
		return m.Keyring
	}
	return ""
}

func (m *VolumeOption) GetBytesPerSec() int32 {
	if m != nil {
		return m.BytesPerSec
	}
	return 0
}

func (m *VolumeOption) GetIops() int32 {
	if m != nil {
		return m.Iops
	}
	return 0
}

type UserGroupInfo struct {
	User             string   `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Group            string   `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
	AdditionalGroups []string `protobuf:"bytes,3,rep,name=additionalGroups" json:"additionalGroups,omitempty"`
}

func (m *UserGroupInfo) Reset()                    { *m = UserGroupInfo{} }
func (m *UserGroupInfo) String() string            { return proto.CompactTextString(m) }
func (*UserGroupInfo) ProtoMessage()               {}
func (*UserGroupInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UserGroupInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserGroupInfo) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *UserGroupInfo) GetAdditionalGroups() []string {
	if m != nil {
		return m.AdditionalGroups
	}
	return nil
}

type Rlimit struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Hard uint64 `protobuf:"varint,2,opt,name=hard" json:"hard,omitempty"`
	Soft uint64 `protobuf:"varint,3,opt,name=soft" json:"soft,omitempty"`
}

func (m *Rlimit) Reset()                    { *m = Rlimit{} }
func (m *Rlimit) String() string            { return proto.CompactTextString(m) }
func (*Rlimit) ProtoMessage()               {}
func (*Rlimit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Rlimit) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Rlimit) GetHard() uint64 {
	if m != nil {
		return m.Hard
	}
	return 0
}

func (m *Rlimit) GetSoft() uint64 {
	if m != nil {
		return m.Soft
	}
	return 0
}

type Process struct {
	Container       string   `protobuf:"bytes,1,opt,name=Container" json:"Container,omitempty"`
	Id              string   `protobuf:"bytes,2,opt,name=Id" json:"Id,omitempty"`
	User            string   `protobuf:"bytes,3,opt,name=User" json:"User,omitempty"`
	Group           string   `protobuf:"bytes,4,opt,name=Group" json:"Group,omitempty"`
	AdditionalGroup []string `protobuf:"bytes,5,rep,name=AdditionalGroup" json:"AdditionalGroup,omitempty"`
	Terminal        bool     `protobuf:"varint,6,opt,name=Terminal" json:"Terminal,omitempty"`
	Args            []string `protobuf:"bytes,7,rep,name=Args" json:"Args,omitempty"`
	Envs            []string `protobuf:"bytes,8,rep,name=Envs" json:"Envs,omitempty"`
	Workdir         string   `protobuf:"bytes,9,opt,name=Workdir" json:"Workdir,omitempty"`
}

func (m *Process) Reset()                    { *m = Process{} }
func (m *Process) String() string            { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()               {}
func (*Process) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Process) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *Process) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Process) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Process) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Process) GetAdditionalGroup() []string {
	if m != nil {
		return m.AdditionalGroup
	}
	return nil
}

func (m *Process) GetTerminal() bool {
	if m != nil {
		return m.Terminal
	}
	return false
}

func (m *Process) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Process) GetEnvs() []string {
	if m != nil {
		return m.Envs
	}
	return nil
}

func (m *Process) GetWorkdir() string {
	if m != nil {
		return m.Workdir
	}
	return ""
}

func init() {
	proto.RegisterType((*SandboxConfig)(nil), "api.SandboxConfig")
	proto.RegisterType((*ContainerDescription)(nil), "api.ContainerDescription")
	proto.RegisterType((*VolumeDescription)(nil), "api.VolumeDescription")
	proto.RegisterType((*InterfaceDescription)(nil), "api.InterfaceDescription")
	proto.RegisterType((*PortDescription)(nil), "api.PortDescription")
	proto.RegisterType((*NeighborNetworks)(nil), "api.NeighborNetworks")
	proto.RegisterType((*VolumeReference)(nil), "api.VolumeReference")
	proto.RegisterType((*VolumeMount)(nil), "api.VolumeMount")
	proto.RegisterType((*VolumeOption)(nil), "api.VolumeOption")
	proto.RegisterType((*UserGroupInfo)(nil), "api.UserGroupInfo")
	proto.RegisterType((*Rlimit)(nil), "api.Rlimit")
	proto.RegisterType((*Process)(nil), "api.Process")
}

func init() { proto.RegisterFile("descriptions.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1066 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdd, 0x6a, 0x1b, 0x47,
	0x14, 0x66, 0xd7, 0x92, 0x25, 0x1d, 0xf9, 0x77, 0x70, 0xc3, 0x62, 0x4a, 0x31, 0xdb, 0xa4, 0x98,
	0x14, 0x7c, 0xe1, 0x40, 0xd3, 0x14, 0x02, 0x0d, 0x89, 0x09, 0x82, 0xd6, 0x16, 0xe3, 0xa6, 0xa5,
	0x97, 0xa3, 0xdd, 0x91, 0x34, 0xf5, 0x6a, 0x66, 0x99, 0x19, 0xd9, 0x51, 0x1f, 0xa2, 0xf7, 0x7d,
	0x8b, 0x3e, 0x4f, 0xdf, 0xa0, 0xb7, 0x7d, 0x82, 0x72, 0xce, 0xce, 0xae, 0xd7, 0xb2, 0x43, 0xf1,
	0xdd, 0xf9, 0xbe, 0x3d, 0x73, 0xe6, 0xfc, 0xcf, 0x02, 0xcb, 0xa5, 0xcb, 0xac, 0x2a, 0xbd, 0x32,
	0xda, 0x9d, 0x94, 0xd6, 0x78, 0xc3, 0x36, 0x44, 0xa9, 0xd2, 0xbf, 0x22, 0xd8, 0xbe, 0x14, 0x3a,
	0x9f, 0x98, 0x8f, 0x6f, 0x8d, 0x9e, 0xaa, 0x19, 0x3b, 0x84, 0xfe, 0xdc, 0x38, 0xaf, 0xc5, 0x42,
	0x26, 0xd1, 0x51, 0x74, 0x3c, 0xe0, 0x0d, 0x66, 0x7b, 0xb0, 0x91, 0x6b, 0x97, 0xc4, 0x47, 0x1b,
	0xc7, 0x03, 0x8e, 0x22, 0x7b, 0x01, 0x03, 0x2d, 0xd5, 0x6c, 0x3e, 0x31, 0xd6, 0x25, 0x1b, 0x47,
	0xd1, 0xf1, 0xf0, 0xf4, 0xb3, 0x13, 0x51, 0xaa, 0x93, 0xf3, 0xc0, 0x9e, 0x4b, 0x7f, 0x63, 0xec,
	0x95, 0xe3, 0xb7, 0x7a, 0xec, 0x0b, 0x80, 0x5c, 0xbb, 0x8b, 0xca, 0x9b, 0xa4, 0x43, 0xd6, 0x5a,
	0x0c, 0xfb, 0x1c, 0x06, 0xb9, 0x76, 0x97, 0x52, 0xd8, 0x6c, 0x9e, 0x74, 0xe9, 0xf3, 0x2d, 0x91,
	0xfe, 0xd9, 0x83, 0x83, 0xb7, 0x46, 0x7b, 0xa1, 0xb4, 0xb4, 0xef, 0x6e, 0xe3, 0x62, 0x3b, 0x10,
	0xab, 0x3c, 0xf8, 0x1c, 0xab, 0x9c, 0x31, 0xe8, 0x50, 0x14, 0x31, 0x31, 0x24, 0xb3, 0x03, 0xe8,
	0xaa, 0x85, 0x98, 0x49, 0xf2, 0x75, 0xc0, 0x2b, 0xc0, 0x5e, 0xc3, 0x66, 0x21, 0x26, 0xb2, 0xa8,
	0x9c, 0x19, 0x9e, 0x3e, 0xa3, 0x10, 0x1e, 0xba, 0xe4, 0xe4, 0x07, 0xd2, 0x3b, 0xd3, 0xde, 0xae,
	0x78, 0x38, 0x84, 0x69, 0xf1, 0x7e, 0x95, 0x74, 0x8f, 0xa2, 0xe3, 0x3e, 0x47, 0x11, 0x23, 0x74,
	0xde, 0x94, 0x97, 0x6a, 0xa6, 0x45, 0x91, 0x6c, 0xd2, 0x5d, 0x2d, 0x86, 0x7d, 0x03, 0x60, 0x8d,
	0xf1, 0x3f, 0x9b, 0x62, 0xb9, 0x90, 0x49, 0x8f, 0xf2, 0xf6, 0x84, 0x2e, 0xad, 0xa8, 0xd6, 0x8d,
	0xbc, 0xa5, 0xc9, 0x12, 0xe8, 0x2d, 0xcc, 0x52, 0xfb, 0x51, 0x9e, 0xf4, 0xc9, 0x68, 0x0d, 0xb1,
	0x6c, 0xa8, 0x37, 0x16, 0x7e, 0x9e, 0x0c, 0xaa, 0xb2, 0xd5, 0x98, 0x3d, 0x85, 0x8d, 0x0f, 0xef,
	0x47, 0x09, 0xd0, 0x35, 0x8c, 0xae, 0xf9, 0xe0, 0xa4, 0x7d, 0x6f, 0xcd, 0xb2, 0x1c, 0xe9, 0xa9,
	0xe1, 0xf8, 0x99, 0xbd, 0x84, 0x8e, 0xd4, 0xd7, 0x2e, 0x19, 0x52, 0x0a, 0xbe, 0xfc, 0x74, 0x0a,
	0xce, 0xf4, 0x75, 0x48, 0x00, 0x1d, 0x40, 0xa7, 0xb0, 0xc4, 0xb9, 0xb2, 0xc9, 0x56, 0xe5, 0x54,
	0x80, 0x58, 0x81, 0x12, 0x1d, 0xda, 0xae, 0x2a, 0x80, 0x32, 0x72, 0xc2, 0xce, 0x5c, 0xb2, 0x43,
	0x75, 0x25, 0x99, 0x3d, 0x83, 0x9e, 0x2d, 0xd4, 0x42, 0x79, 0x97, 0xec, 0xd2, 0xed, 0x43, 0xba,
	0x9d, 0x13, 0xc7, 0xeb, 0x6f, 0x58, 0x26, 0xb7, 0x72, 0x99, 0x2f, 0x92, 0xbd, 0xff, 0x2b, 0xd3,
	0x25, 0xe9, 0x85, 0x32, 0x55, 0x87, 0xd8, 0xf7, 0xd0, 0xbb, 0xa6, 0x34, 0xba, 0x64, 0x9f, 0xce,
	0x7f, 0xf5, 0xe9, 0xf3, 0x55, 0xbe, 0x43, 0x98, 0xf5, 0x31, 0x2c, 0xab, 0xd2, 0xca, 0x2b, 0x51,
	0xa8, 0xdf, 0x65, 0x92, 0x50, 0xbd, 0x5b, 0xcc, 0xe1, 0x2b, 0x18, 0xb6, 0xfa, 0x03, 0xfb, 0xe2,
	0x4a, 0xae, 0x42, 0x47, 0xa2, 0x88, 0xed, 0x77, 0x2d, 0x8a, 0x65, 0xdd, 0x93, 0x15, 0xf8, 0x2e,
	0xfe, 0x36, 0x3a, 0x7c, 0x09, 0x83, 0x26, 0xaf, 0x8f, 0x3a, 0xf8, 0x0a, 0x86, 0xad, 0x60, 0x1f,
	0x75, 0x74, 0x0c, 0x5b, 0xed, 0x38, 0x1f, 0x38, 0xfb, 0xbc, 0x7d, 0x76, 0x78, 0x7a, 0xd0, 0x6a,
	0x51, 0x2e, 0xa7, 0xd2, 0x4a, 0x9d, 0xc9, 0x96, 0xc5, 0xf4, 0xdf, 0x08, 0xf6, 0xef, 0x75, 0x70,
	0x33, 0x88, 0x51, 0x6b, 0x10, 0x9f, 0xc0, 0xa6, 0x33, 0x4b, 0x9b, 0xd5, 0x6e, 0x05, 0x84, 0xfc,
	0xd4, 0xd8, 0x85, 0xf0, 0x61, 0x42, 0x03, 0x22, 0xde, 0xf9, 0x55, 0x29, 0x93, 0x4e, 0xe0, 0x09,
	0x61, 0x74, 0x99, 0xc8, 0xe6, 0x92, 0xa6, 0x6f, 0xc0, 0x2b, 0xc0, 0xbe, 0x86, 0x9e, 0x09, 0xeb,
	0xa5, 0x4f, 0x9e, 0xef, 0xb7, 0x3c, 0xaf, 0xd6, 0x0c, 0xaf, 0x35, 0x58, 0x0a, 0x5b, 0xb9, 0xc9,
	0xae, 0xa4, 0x0d, 0xe3, 0x38, 0xa0, 0xba, 0xde, 0xe1, 0x68, 0xbc, 0xa4, 0xc8, 0x2f, 0x74, 0xb1,
	0xa2, 0x39, 0xea, 0xf3, 0x06, 0xa7, 0x7f, 0x47, 0x70, 0x30, 0xd2, 0x5e, 0xda, 0xa9, 0xc8, 0xe4,
	0x63, 0x17, 0xd2, 0x0e, 0xc4, 0x85, 0xa1, 0x58, 0xfb, 0x3c, 0x2e, 0x0c, 0xc6, 0x39, 0xb1, 0x2a,
	0x9f, 0x35, 0x71, 0x56, 0x88, 0x6c, 0x95, 0x21, 0xc8, 0x58, 0x95, 0x58, 0xab, 0x85, 0xc8, 0xc2,
	0x6a, 0x41, 0x91, 0x18, 0xbf, 0xa4, 0x65, 0xd2, 0xe1, 0x28, 0xe2, 0x99, 0xd9, 0x4d, 0x58, 0x14,
	0xf1, 0xec, 0x06, 0x07, 0xd5, 0x8b, 0xf2, 0x5c, 0x84, 0x18, 0x07, 0xbc, 0x86, 0xf8, 0xa5, 0xce,
	0x17, 0x54, 0x5f, 0x02, 0x4c, 0x0d, 0xec, 0x8e, 0x8d, 0xf5, 0xed, 0xb0, 0xc2, 0x0b, 0x81, 0x34,
	0x05, 0xd7, 0xe5, 0x0d, 0x66, 0x4f, 0x61, 0x3b, 0xab, 0xe7, 0x89, 0x14, 0x62, 0x52, 0xb8, 0x4b,
	0xa2, 0x05, 0x7a, 0x83, 0x32, 0x53, 0x84, 0x32, 0x37, 0x38, 0xfd, 0x0d, 0xf6, 0xd6, 0xdf, 0x0e,
	0xf6, 0x1c, 0xf6, 0x14, 0x26, 0x58, 0x8b, 0xa2, 0xe6, 0x92, 0x88, 0xf6, 0xc7, 0x3d, 0x1e, 0x75,
	0xe5, 0xc7, 0x35, 0xdd, 0xea, 0xc1, 0xba, 0xc7, 0xa7, 0xbf, 0xc2, 0xee, 0x5a, 0x33, 0x3f, 0xd8,
	0xab, 0xa7, 0x30, 0xa4, 0x35, 0x3b, 0x36, 0x4a, 0xfb, 0xca, 0xda, 0xf0, 0x74, 0xaf, 0xd5, 0x51,
	0x3f, 0xe2, 0x57, 0xde, 0x56, 0x4a, 0x5f, 0xc3, 0xb0, 0xf5, 0xad, 0xd9, 0x84, 0x51, 0x6b, 0x13,
	0xb6, 0x7b, 0x2a, 0x5e, 0xeb, 0xa9, 0x3f, 0xa2, 0x7a, 0x36, 0x2f, 0x9a, 0x19, 0x5a, 0x3a, 0x69,
	0x6b, 0x03, 0x28, 0xa3, 0x81, 0x85, 0xd1, 0xca, 0xe3, 0xdb, 0x5b, 0x85, 0xd8, 0x60, 0xac, 0xe8,
	0x95, 0x5c, 0x59, 0xa5, 0x67, 0x21, 0xc3, 0x35, 0x64, 0x47, 0x30, 0x9c, 0xac, 0xbc, 0x74, 0x63,
	0x69, 0x2f, 0x65, 0x46, 0x6d, 0xd6, 0xe5, 0x6d, 0x0a, 0xef, 0x52, 0xa6, 0x74, 0xd4, 0x6d, 0x5d,
	0x4e, 0x72, 0x2a, 0x61, 0xfb, 0xce, 0x9b, 0xf1, 0xa0, 0x43, 0x07, 0xd0, 0x9d, 0xa1, 0x42, 0xbd,
	0x6a, 0x08, 0x60, 0x45, 0x44, 0x9e, 0x2b, 0x0c, 0x43, 0x14, 0x64, 0x00, 0x7f, 0x15, 0xa8, 0x22,
	0xeb, 0x7c, 0xfa, 0x0e, 0x36, 0xab, 0xad, 0x8f, 0xf6, 0x69, 0xdc, 0x83, 0x7d, 0x1a, 0x76, 0x06,
	0x9d, 0xb9, 0xb0, 0x39, 0x99, 0xef, 0x70, 0x92, 0x91, 0x73, 0x66, 0x5a, 0xad, 0x8b, 0x0e, 0x27,
	0x39, 0xfd, 0x27, 0x82, 0xde, 0xd8, 0x9a, 0x4c, 0x3a, 0xfa, 0x99, 0x68, 0x36, 0x7c, 0x30, 0x76,
	0x4b, 0xe0, 0x88, 0x8c, 0xf2, 0xe0, 0x6e, 0x3c, 0x22, 0x6b, 0x18, 0x66, 0xc8, 0x19, 0xc9, 0x18,
	0x15, 0x79, 0x17, 0x26, 0xb2, 0x02, 0xec, 0x18, 0x76, 0xdf, 0xdc, 0xf5, 0x3e, 0xfc, 0xaa, 0xac,
	0xd3, 0x58, 0xa6, 0x9f, 0xa4, 0x5d, 0xa8, 0xfa, 0x57, 0xa0, 0xcf, 0x1b, 0x8c, 0xf7, 0xbd, 0xc1,
	0xd7, 0xb0, 0x57, 0xbd, 0x86, 0x28, 0x23, 0x87, 0x4f, 0x41, 0xd2, 0xaf, 0xb8, 0xb3, 0xf0, 0xc6,
	0xfe, 0x12, 0xde, 0xd8, 0x30, 0xba, 0x01, 0x4e, 0x36, 0x69, 0x72, 0x5e, 0xfc, 0x17, 0x00, 0x00,
	0xff, 0xff, 0xa8, 0x9c, 0xe1, 0xeb, 0xe3, 0x09, 0x00, 0x00,
}
