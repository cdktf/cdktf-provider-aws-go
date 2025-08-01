// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/alb/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb aws_alb}.
type Alb interface {
	cdktf.TerraformResource
	AccessLogs() AlbAccessLogsOutputReference
	AccessLogsInput() *AlbAccessLogs
	Arn() *string
	ArnSuffix() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientKeepAlive() *float64
	SetClientKeepAlive(val *float64)
	ClientKeepAliveInput() *float64
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionLogs() AlbConnectionLogsOutputReference
	ConnectionLogsInput() *AlbConnectionLogs
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomerOwnedIpv4Pool() *string
	SetCustomerOwnedIpv4Pool(val *string)
	CustomerOwnedIpv4PoolInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesyncMitigationMode() *string
	SetDesyncMitigationMode(val *string)
	DesyncMitigationModeInput() *string
	DnsName() *string
	DnsRecordClientRoutingPolicy() *string
	SetDnsRecordClientRoutingPolicy(val *string)
	DnsRecordClientRoutingPolicyInput() *string
	DropInvalidHeaderFields() interface{}
	SetDropInvalidHeaderFields(val interface{})
	DropInvalidHeaderFieldsInput() interface{}
	EnableCrossZoneLoadBalancing() interface{}
	SetEnableCrossZoneLoadBalancing(val interface{})
	EnableCrossZoneLoadBalancingInput() interface{}
	EnableDeletionProtection() interface{}
	SetEnableDeletionProtection(val interface{})
	EnableDeletionProtectionInput() interface{}
	EnableHttp2() interface{}
	SetEnableHttp2(val interface{})
	EnableHttp2Input() interface{}
	EnableTlsVersionAndCipherSuiteHeaders() interface{}
	SetEnableTlsVersionAndCipherSuiteHeaders(val interface{})
	EnableTlsVersionAndCipherSuiteHeadersInput() interface{}
	EnableWafFailOpen() interface{}
	SetEnableWafFailOpen(val interface{})
	EnableWafFailOpenInput() interface{}
	EnableXffClientPort() interface{}
	SetEnableXffClientPort(val interface{})
	EnableXffClientPortInput() interface{}
	EnableZonalShift() interface{}
	SetEnableZonalShift(val interface{})
	EnableZonalShiftInput() interface{}
	EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic() *string
	SetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic(val *string)
	EnforceSecurityGroupInboundRulesOnPrivateLinkTrafficInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IdleTimeout() *float64
	SetIdleTimeout(val *float64)
	IdleTimeoutInput() *float64
	Internal() interface{}
	SetInternal(val interface{})
	InternalInput() interface{}
	IpAddressType() *string
	SetIpAddressType(val *string)
	IpAddressTypeInput() *string
	IpamPools() AlbIpamPoolsOutputReference
	IpamPoolsInput() *AlbIpamPools
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancerType() *string
	SetLoadBalancerType(val *string)
	LoadBalancerTypeInput() *string
	MinimumLoadBalancerCapacity() AlbMinimumLoadBalancerCapacityOutputReference
	MinimumLoadBalancerCapacityInput() *AlbMinimumLoadBalancerCapacity
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	// The tree node.
	Node() constructs.Node
	PreserveHostHeader() interface{}
	SetPreserveHostHeader(val interface{})
	PreserveHostHeaderInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	SubnetMapping() AlbSubnetMappingList
	SubnetMappingInput() interface{}
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	SubnetsInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AlbTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcId() *string
	XffHeaderProcessingMode() *string
	SetXffHeaderProcessingMode(val *string)
	XffHeaderProcessingModeInput() *string
	ZoneId() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAccessLogs(value *AlbAccessLogs)
	PutConnectionLogs(value *AlbConnectionLogs)
	PutIpamPools(value *AlbIpamPools)
	PutMinimumLoadBalancerCapacity(value *AlbMinimumLoadBalancerCapacity)
	PutSubnetMapping(value interface{})
	PutTimeouts(value *AlbTimeouts)
	ResetAccessLogs()
	ResetClientKeepAlive()
	ResetConnectionLogs()
	ResetCustomerOwnedIpv4Pool()
	ResetDesyncMitigationMode()
	ResetDnsRecordClientRoutingPolicy()
	ResetDropInvalidHeaderFields()
	ResetEnableCrossZoneLoadBalancing()
	ResetEnableDeletionProtection()
	ResetEnableHttp2()
	ResetEnableTlsVersionAndCipherSuiteHeaders()
	ResetEnableWafFailOpen()
	ResetEnableXffClientPort()
	ResetEnableZonalShift()
	ResetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic()
	ResetId()
	ResetIdleTimeout()
	ResetInternal()
	ResetIpAddressType()
	ResetIpamPools()
	ResetLoadBalancerType()
	ResetMinimumLoadBalancerCapacity()
	ResetName()
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreserveHostHeader()
	ResetRegion()
	ResetSecurityGroups()
	ResetSubnetMapping()
	ResetSubnets()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetXffHeaderProcessingMode()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Alb
type jsiiProxy_Alb struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Alb) AccessLogs() AlbAccessLogsOutputReference {
	var returns AlbAccessLogsOutputReference
	_jsii_.Get(
		j,
		"accessLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) AccessLogsInput() *AlbAccessLogs {
	var returns *AlbAccessLogs
	_jsii_.Get(
		j,
		"accessLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) ArnSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) ClientKeepAlive() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) ClientKeepAliveInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) ConnectionLogs() AlbConnectionLogsOutputReference {
	var returns AlbConnectionLogsOutputReference
	_jsii_.Get(
		j,
		"connectionLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) ConnectionLogsInput() *AlbConnectionLogs {
	var returns *AlbConnectionLogs
	_jsii_.Get(
		j,
		"connectionLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) CustomerOwnedIpv4Pool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerOwnedIpv4Pool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) CustomerOwnedIpv4PoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerOwnedIpv4PoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) DesyncMitigationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desyncMitigationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) DesyncMitigationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desyncMitigationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) DnsRecordClientRoutingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsRecordClientRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) DnsRecordClientRoutingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsRecordClientRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) DropInvalidHeaderFields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropInvalidHeaderFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) DropInvalidHeaderFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropInvalidHeaderFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnableCrossZoneLoadBalancing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCrossZoneLoadBalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnableCrossZoneLoadBalancingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCrossZoneLoadBalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnableDeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDeletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnableDeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDeletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnableHttp2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttp2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnableHttp2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttp2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnableTlsVersionAndCipherSuiteHeaders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTlsVersionAndCipherSuiteHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnableTlsVersionAndCipherSuiteHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTlsVersionAndCipherSuiteHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnableWafFailOpen() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWafFailOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnableWafFailOpenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWafFailOpenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnableXffClientPort() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableXffClientPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnableXffClientPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableXffClientPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnableZonalShift() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableZonalShift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnableZonalShiftInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableZonalShiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) EnforceSecurityGroupInboundRulesOnPrivateLinkTrafficInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceSecurityGroupInboundRulesOnPrivateLinkTrafficInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) IdleTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) IdleTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) Internal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) InternalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) IpamPools() AlbIpamPoolsOutputReference {
	var returns AlbIpamPoolsOutputReference
	_jsii_.Get(
		j,
		"ipamPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) IpamPoolsInput() *AlbIpamPools {
	var returns *AlbIpamPools
	_jsii_.Get(
		j,
		"ipamPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) LoadBalancerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) LoadBalancerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) MinimumLoadBalancerCapacity() AlbMinimumLoadBalancerCapacityOutputReference {
	var returns AlbMinimumLoadBalancerCapacityOutputReference
	_jsii_.Get(
		j,
		"minimumLoadBalancerCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) MinimumLoadBalancerCapacityInput() *AlbMinimumLoadBalancerCapacity {
	var returns *AlbMinimumLoadBalancerCapacity
	_jsii_.Get(
		j,
		"minimumLoadBalancerCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) PreserveHostHeader() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) PreserveHostHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveHostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) SubnetMapping() AlbSubnetMappingList {
	var returns AlbSubnetMappingList
	_jsii_.Get(
		j,
		"subnetMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) SubnetMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) SubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) Timeouts() AlbTimeoutsOutputReference {
	var returns AlbTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) XffHeaderProcessingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xffHeaderProcessingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) XffHeaderProcessingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xffHeaderProcessingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alb) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb aws_alb} Resource.
func NewAlb(scope constructs.Construct, id *string, config *AlbConfig) Alb {
	_init_.Initialize()

	if err := validateNewAlbParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Alb{}

	_jsii_.Create(
		"@cdktf/provider-aws.alb.Alb",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb aws_alb} Resource.
func NewAlb_Override(a Alb, scope constructs.Construct, id *string, config *AlbConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.alb.Alb",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Alb)SetClientKeepAlive(val *float64) {
	if err := j.validateSetClientKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientKeepAlive",
		val,
	)
}

func (j *jsiiProxy_Alb)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Alb)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Alb)SetCustomerOwnedIpv4Pool(val *string) {
	if err := j.validateSetCustomerOwnedIpv4PoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerOwnedIpv4Pool",
		val,
	)
}

func (j *jsiiProxy_Alb)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Alb)SetDesyncMitigationMode(val *string) {
	if err := j.validateSetDesyncMitigationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desyncMitigationMode",
		val,
	)
}

func (j *jsiiProxy_Alb)SetDnsRecordClientRoutingPolicy(val *string) {
	if err := j.validateSetDnsRecordClientRoutingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsRecordClientRoutingPolicy",
		val,
	)
}

func (j *jsiiProxy_Alb)SetDropInvalidHeaderFields(val interface{}) {
	if err := j.validateSetDropInvalidHeaderFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dropInvalidHeaderFields",
		val,
	)
}

func (j *jsiiProxy_Alb)SetEnableCrossZoneLoadBalancing(val interface{}) {
	if err := j.validateSetEnableCrossZoneLoadBalancingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCrossZoneLoadBalancing",
		val,
	)
}

func (j *jsiiProxy_Alb)SetEnableDeletionProtection(val interface{}) {
	if err := j.validateSetEnableDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDeletionProtection",
		val,
	)
}

func (j *jsiiProxy_Alb)SetEnableHttp2(val interface{}) {
	if err := j.validateSetEnableHttp2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHttp2",
		val,
	)
}

func (j *jsiiProxy_Alb)SetEnableTlsVersionAndCipherSuiteHeaders(val interface{}) {
	if err := j.validateSetEnableTlsVersionAndCipherSuiteHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTlsVersionAndCipherSuiteHeaders",
		val,
	)
}

func (j *jsiiProxy_Alb)SetEnableWafFailOpen(val interface{}) {
	if err := j.validateSetEnableWafFailOpenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableWafFailOpen",
		val,
	)
}

func (j *jsiiProxy_Alb)SetEnableXffClientPort(val interface{}) {
	if err := j.validateSetEnableXffClientPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableXffClientPort",
		val,
	)
}

func (j *jsiiProxy_Alb)SetEnableZonalShift(val interface{}) {
	if err := j.validateSetEnableZonalShiftParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableZonalShift",
		val,
	)
}

func (j *jsiiProxy_Alb)SetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic(val *string) {
	if err := j.validateSetEnforceSecurityGroupInboundRulesOnPrivateLinkTrafficParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic",
		val,
	)
}

func (j *jsiiProxy_Alb)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Alb)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Alb)SetIdleTimeout(val *float64) {
	if err := j.validateSetIdleTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeout",
		val,
	)
}

func (j *jsiiProxy_Alb)SetInternal(val interface{}) {
	if err := j.validateSetInternalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internal",
		val,
	)
}

func (j *jsiiProxy_Alb)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_Alb)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Alb)SetLoadBalancerType(val *string) {
	if err := j.validateSetLoadBalancerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerType",
		val,
	)
}

func (j *jsiiProxy_Alb)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Alb)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_Alb)SetPreserveHostHeader(val interface{}) {
	if err := j.validateSetPreserveHostHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveHostHeader",
		val,
	)
}

func (j *jsiiProxy_Alb)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Alb)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Alb)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_Alb)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_Alb)SetSubnets(val *[]*string) {
	if err := j.validateSetSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_Alb)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Alb)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_Alb)SetXffHeaderProcessingMode(val *string) {
	if err := j.validateSetXffHeaderProcessingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xffHeaderProcessingMode",
		val,
	)
}

// Generates CDKTF code for importing a Alb resource upon running "cdktf plan <stack-name>".
func Alb_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAlb_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.alb.Alb",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Alb_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlb_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.alb.Alb",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Alb_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlb_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.alb.Alb",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Alb_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlb_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.alb.Alb",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Alb_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.alb.Alb",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_Alb) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_Alb) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_Alb) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_Alb) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_Alb) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_Alb) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_Alb) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Alb) PutAccessLogs(value *AlbAccessLogs) {
	if err := a.validatePutAccessLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Alb) PutConnectionLogs(value *AlbConnectionLogs) {
	if err := a.validatePutConnectionLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectionLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Alb) PutIpamPools(value *AlbIpamPools) {
	if err := a.validatePutIpamPoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIpamPools",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Alb) PutMinimumLoadBalancerCapacity(value *AlbMinimumLoadBalancerCapacity) {
	if err := a.validatePutMinimumLoadBalancerCapacityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMinimumLoadBalancerCapacity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Alb) PutSubnetMapping(value interface{}) {
	if err := a.validatePutSubnetMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSubnetMapping",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Alb) PutTimeouts(value *AlbTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Alb) ResetAccessLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetClientKeepAlive() {
	_jsii_.InvokeVoid(
		a,
		"resetClientKeepAlive",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetConnectionLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetCustomerOwnedIpv4Pool() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerOwnedIpv4Pool",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetDesyncMitigationMode() {
	_jsii_.InvokeVoid(
		a,
		"resetDesyncMitigationMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetDnsRecordClientRoutingPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetDnsRecordClientRoutingPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetDropInvalidHeaderFields() {
	_jsii_.InvokeVoid(
		a,
		"resetDropInvalidHeaderFields",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetEnableCrossZoneLoadBalancing() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableCrossZoneLoadBalancing",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetEnableDeletionProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableDeletionProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetEnableHttp2() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableHttp2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetEnableTlsVersionAndCipherSuiteHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableTlsVersionAndCipherSuiteHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetEnableWafFailOpen() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableWafFailOpen",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetEnableXffClientPort() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableXffClientPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetEnableZonalShift() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableZonalShift",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic() {
	_jsii_.InvokeVoid(
		a,
		"resetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetIdleTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetIdleTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetInternal() {
	_jsii_.InvokeVoid(
		a,
		"resetInternal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		a,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetIpamPools() {
	_jsii_.InvokeVoid(
		a,
		"resetIpamPools",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetLoadBalancerType() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancerType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetMinimumLoadBalancerCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetMinimumLoadBalancerCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetPreserveHostHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetPreserveHostHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetSubnetMapping() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetMapping",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetSubnets() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) ResetXffHeaderProcessingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetXffHeaderProcessingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Alb) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alb) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

