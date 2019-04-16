// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: api/external/applications/applications.proto

package applications

import policyv2 "github.com/chef/automate/components/automate-gateway/authz/policy_v2"

func init() {
	policyv2.MapMethodTo("/chef.automate.api.applications.ApplicationsService/GetServiceGroups", "infra:nodes", "infra:nodes:list", "GET", "/beta/applications/service-groups", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.applications.ApplicationsService/GetServiceGroupsHealthCounts", "infra:nodes", "infra:nodes:list", "GET", "/beta/applications/service_groups_health_counts", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.applications.ApplicationsService/GetServices", "infra:nodes", "infra:nodes:list", "GET", "/beta/applications/services", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.applications.ApplicationsService/GetServicesBySG", "infra:nodes", "infra:nodes:list", "GET", "/beta/applications/service-groups/{service_group_id}", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.applications.ApplicationsService/GetVersion", "system:service:version", "system:serviceVersion:get", "GET", "/beta/applications/version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
