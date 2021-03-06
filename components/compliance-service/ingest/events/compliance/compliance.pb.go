// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/compliance-service/ingest/events/compliance/compliance.proto

package compliance // import "github.com/chef/automate/components/compliance-service/ingest/events/compliance"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import inspec "github.com/chef/automate/components/compliance-service/ingest/events/inspec"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Report struct {
	// inspec full json report fields
	Version     string             `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Platform    *inspec.Platform   `protobuf:"bytes,16,opt,name=platform,proto3" json:"platform,omitempty"`
	Statistics  *inspec.Statistics `protobuf:"bytes,17,opt,name=statistics,proto3" json:"statistics,omitempty"`
	Profiles    []*inspec.Profile  `protobuf:"bytes,18,rep,name=profiles,proto3" json:"profiles,omitempty"`
	OtherChecks []string           `protobuf:"bytes,19,rep,name=other_checks,json=otherChecks,proto3" json:"other_checks,omitempty"`
	// extra report fields added by the audit cookbook
	ReportUuid           string   `protobuf:"bytes,20,opt,name=report_uuid,json=reportUuid,proto3" json:"report_uuid,omitempty"`
	NodeUuid             string   `protobuf:"bytes,21,opt,name=node_uuid,json=nodeUuid,proto3" json:"node_uuid,omitempty"`
	JobUuid              string   `protobuf:"bytes,22,opt,name=job_uuid,json=jobUuid,proto3" json:"job_uuid,omitempty"`
	NodeName             string   `protobuf:"bytes,23,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Environment          string   `protobuf:"bytes,24,opt,name=environment,proto3" json:"environment,omitempty"`
	Roles                []string `protobuf:"bytes,25,rep,name=roles,proto3" json:"roles,omitempty"`
	Recipes              []string `protobuf:"bytes,26,rep,name=recipes,proto3" json:"recipes,omitempty"`
	EndTime              string   `protobuf:"bytes,27,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Type                 string   `protobuf:"bytes,28,opt,name=type,proto3" json:"type,omitempty"`
	SourceId             string   `protobuf:"bytes,29,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	SourceRegion         string   `protobuf:"bytes,30,opt,name=source_region,json=sourceRegion,proto3" json:"source_region,omitempty"`
	SourceAccountId      string   `protobuf:"bytes,31,opt,name=source_account_id,json=sourceAccountId,proto3" json:"source_account_id,omitempty"`
	PolicyName           string   `protobuf:"bytes,32,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	PolicyGroup          string   `protobuf:"bytes,33,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	OrganizationName     string   `protobuf:"bytes,34,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty"`
	SourceFqdn           string   `protobuf:"bytes,35,opt,name=source_fqdn,json=sourceFqdn,proto3" json:"source_fqdn,omitempty"`
	ChefTags             []string `protobuf:"bytes,36,rep,name=chef_tags,json=chefTags,proto3" json:"chef_tags,omitempty"`
	Ipaddress            string   `protobuf:"bytes,37,opt,name=ipaddress,proto3" json:"ipaddress,omitempty"`
	Fqdn                 string   `protobuf:"bytes,38,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_compliance_dcbc78db4e09f334, []int{0}
}
func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (dst *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(dst, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Report) GetPlatform() *inspec.Platform {
	if m != nil {
		return m.Platform
	}
	return nil
}

func (m *Report) GetStatistics() *inspec.Statistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *Report) GetProfiles() []*inspec.Profile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func (m *Report) GetOtherChecks() []string {
	if m != nil {
		return m.OtherChecks
	}
	return nil
}

func (m *Report) GetReportUuid() string {
	if m != nil {
		return m.ReportUuid
	}
	return ""
}

func (m *Report) GetNodeUuid() string {
	if m != nil {
		return m.NodeUuid
	}
	return ""
}

func (m *Report) GetJobUuid() string {
	if m != nil {
		return m.JobUuid
	}
	return ""
}

func (m *Report) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Report) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *Report) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Report) GetRecipes() []string {
	if m != nil {
		return m.Recipes
	}
	return nil
}

func (m *Report) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *Report) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Report) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *Report) GetSourceRegion() string {
	if m != nil {
		return m.SourceRegion
	}
	return ""
}

func (m *Report) GetSourceAccountId() string {
	if m != nil {
		return m.SourceAccountId
	}
	return ""
}

func (m *Report) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *Report) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

func (m *Report) GetOrganizationName() string {
	if m != nil {
		return m.OrganizationName
	}
	return ""
}

func (m *Report) GetSourceFqdn() string {
	if m != nil {
		return m.SourceFqdn
	}
	return ""
}

func (m *Report) GetChefTags() []string {
	if m != nil {
		return m.ChefTags
	}
	return nil
}

func (m *Report) GetIpaddress() string {
	if m != nil {
		return m.Ipaddress
	}
	return ""
}

func (m *Report) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func init() {
	proto.RegisterType((*Report)(nil), "chef.automate.domain.compliance.ingest.events.compliance.Report")
}

func init() {
	proto.RegisterFile("components/compliance-service/ingest/events/compliance/compliance.proto", fileDescriptor_compliance_dcbc78db4e09f334)
}

var fileDescriptor_compliance_dcbc78db4e09f334 = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdb, 0x4f, 0x14, 0x31,
	0x14, 0xc6, 0x43, 0x50, 0xd8, 0xed, 0x60, 0x84, 0x8a, 0x5a, 0x2e, 0xca, 0x00, 0x6a, 0x36, 0x1a,
	0x67, 0x12, 0x7c, 0xf1, 0xc5, 0x7b, 0x22, 0xe1, 0xc5, 0xcb, 0x88, 0x0f, 0xf2, 0xb2, 0x76, 0xdb,
	0xb3, 0xb3, 0xc5, 0x9d, 0x76, 0x68, 0x3b, 0x9b, 0xe0, 0xdf, 0xe3, 0x1f, 0x6a, 0xda, 0x33, 0xc0,
	0xe8, 0x1b, 0xfb, 0xb4, 0xed, 0xef, 0x9c, 0x7e, 0xfd, 0xf6, 0x9b, 0x93, 0x92, 0x43, 0x61, 0xaa,
	0xda, 0x68, 0xd0, 0xde, 0xe5, 0x61, 0x39, 0x55, 0x5c, 0x0b, 0x78, 0xee, 0xc0, 0xce, 0x94, 0x80,
	0x5c, 0xe9, 0x12, 0x9c, 0xcf, 0x61, 0xf6, 0x5f, 0x43, 0x67, 0x99, 0xd5, 0xd6, 0x78, 0x43, 0x5f,
	0x8a, 0x09, 0x8c, 0x33, 0xde, 0x78, 0x53, 0x71, 0x0f, 0x99, 0x34, 0x15, 0x57, 0x3a, 0xeb, 0xb4,
	0xa1, 0x54, 0x86, 0x52, 0x9d, 0xc2, 0xe6, 0x9b, 0xeb, 0x58, 0x50, 0xda, 0xd5, 0x20, 0xda, 0x1f,
	0xbc, 0x7a, 0xef, 0xcf, 0x32, 0x59, 0x2a, 0xa0, 0x36, 0xd6, 0x53, 0x46, 0x96, 0x67, 0x60, 0x9d,
	0x32, 0x9a, 0x2d, 0xa4, 0x0b, 0x83, 0x7e, 0x71, 0xb1, 0xa5, 0x27, 0xa4, 0x57, 0x4f, 0xb9, 0x1f,
	0x1b, 0x5b, 0xb1, 0xd5, 0x74, 0x61, 0x90, 0x1c, 0xbc, 0xce, 0xae, 0x67, 0xb9, 0xbd, 0xf3, 0x4b,
	0xab, 0x52, 0x5c, 0xea, 0xd1, 0x9f, 0x84, 0x38, 0xcf, 0xbd, 0x72, 0x5e, 0x09, 0xc7, 0xd6, 0xa2,
	0xfa, 0xdb, 0xf9, 0xd4, 0xbf, 0x5d, 0xea, 0x14, 0x1d, 0x4d, 0xfa, 0x83, 0xf4, 0x6a, 0x6b, 0xc6,
	0x6a, 0x0a, 0x8e, 0xd1, 0x74, 0x71, 0x90, 0x1c, 0xbc, 0x9a, 0xd3, 0x3d, 0xaa, 0x14, 0x97, 0x72,
	0x74, 0x97, 0xac, 0x18, 0x3f, 0x01, 0x3b, 0x14, 0x13, 0x10, 0xbf, 0x1c, 0xbb, 0x93, 0x2e, 0x0e,
	0xfa, 0x45, 0x12, 0xd9, 0x87, 0x88, 0xe8, 0x0e, 0x49, 0x6c, 0xcc, 0x77, 0xd8, 0x34, 0x4a, 0xb2,
	0xf5, 0x98, 0x2c, 0x41, 0xf4, 0xbd, 0x51, 0x92, 0x6e, 0x91, 0xbe, 0x36, 0x12, 0xb0, 0x7c, 0x37,
	0x96, 0x7b, 0x01, 0xc4, 0xe2, 0x06, 0xe9, 0x9d, 0x9a, 0x11, 0xd6, 0xee, 0xe1, 0x47, 0x39, 0x35,
	0xa3, 0x7f, 0xce, 0x69, 0x5e, 0x01, 0xbb, 0x7f, 0x75, 0xee, 0x13, 0xaf, 0x80, 0xa6, 0x24, 0x01,
	0x3d, 0x53, 0xd6, 0xe8, 0x0a, 0xb4, 0x67, 0x2c, 0x96, 0xbb, 0x88, 0xae, 0x93, 0x9b, 0xd6, 0x84,
	0x48, 0x36, 0xa2, 0x67, 0xdc, 0x84, 0x19, 0xb0, 0x20, 0x54, 0x0d, 0x8e, 0x6d, 0x46, 0x7e, 0xb1,
	0x0d, 0x4e, 0x40, 0xcb, 0xa1, 0x57, 0x15, 0xb0, 0x2d, 0x74, 0x02, 0x5a, 0x1e, 0xab, 0x0a, 0x28,
	0x25, 0x37, 0xfc, 0x79, 0x0d, 0x6c, 0x3b, 0xe2, 0xb8, 0x0e, 0xee, 0x9c, 0x69, 0xac, 0x80, 0xa1,
	0x92, 0xec, 0x01, 0xba, 0x43, 0x70, 0x24, 0xe9, 0x3e, 0xb9, 0xd5, 0x16, 0x2d, 0x94, 0x61, 0xde,
	0x1e, 0xc6, 0x86, 0x15, 0x84, 0x45, 0x64, 0xf4, 0x29, 0x59, 0x6b, 0x9b, 0xb8, 0x10, 0xa6, 0xd1,
	0x3e, 0x28, 0xed, 0xc4, 0xc6, 0xdb, 0x58, 0x78, 0x87, 0xfc, 0x48, 0x86, 0x90, 0x6b, 0x33, 0x55,
	0xe2, 0x1c, 0xd3, 0x48, 0x31, 0x64, 0x44, 0x31, 0x8f, 0x5d, 0xb2, 0xd2, 0x36, 0x94, 0xd6, 0x34,
	0x35, 0xdb, 0xc5, 0x40, 0x90, 0x1d, 0x06, 0x44, 0x9f, 0x91, 0x35, 0x63, 0x4b, 0xae, 0xd5, 0x6f,
	0xee, 0x95, 0xd1, 0xa8, 0xb4, 0x17, 0xfb, 0x56, 0xbb, 0x85, 0xa8, 0xb7, 0x43, 0x92, 0xd6, 0xdc,
	0xf8, 0x4c, 0x6a, 0xb6, 0x8f, 0x17, 0x22, 0xfa, 0x78, 0x26, 0x75, 0xf8, 0xff, 0x61, 0xc6, 0x86,
	0x9e, 0x97, 0x8e, 0x3d, 0x8a, 0x51, 0xf6, 0x02, 0x38, 0xe6, 0xa5, 0xa3, 0xdb, 0xa4, 0xaf, 0x6a,
	0x2e, 0xa5, 0x05, 0xe7, 0xd8, 0xe3, 0x78, 0xf6, 0x0a, 0x84, 0x38, 0xa3, 0xe8, 0x13, 0x8c, 0x33,
	0xac, 0xdf, 0x7f, 0x3d, 0xf9, 0x5c, 0x2a, 0x3f, 0x69, 0x46, 0x61, 0x48, 0xf3, 0x20, 0x94, 0x5f,
	0x4c, 0x6f, 0x3e, 0xdf, 0x2b, 0x34, 0x5a, 0x8a, 0x0f, 0xc0, 0x8b, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xc5, 0xf1, 0x74, 0x7f, 0xc6, 0x04, 0x00, 0x00,
}
