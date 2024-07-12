// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/lb/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/lb aws_lb}.
type Lb interface {
	cdktf.TerraformResource
	AccessLogs() LbAccessLogsOutputReference
	AccessLogsInput() *LbAccessLogs
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
	ConnectionLogs() LbConnectionLogsOutputReference
	ConnectionLogsInput() *LbConnectionLogs
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancerType() *string
	SetLoadBalancerType(val *string)
	LoadBalancerTypeInput() *string
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
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	SubnetMapping() LbSubnetMappingList
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
	Timeouts() LbTimeoutsOutputReference
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
	PutAccessLogs(value *LbAccessLogs)
	PutConnectionLogs(value *LbConnectionLogs)
	PutSubnetMapping(value interface{})
	PutTimeouts(value *LbTimeouts)
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
	ResetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic()
	ResetId()
	ResetIdleTimeout()
	ResetInternal()
	ResetIpAddressType()
	ResetLoadBalancerType()
	ResetName()
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreserveHostHeader()
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

// The jsii proxy struct for Lb
type jsiiProxy_Lb struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Lb) AccessLogs() LbAccessLogsOutputReference {
	var returns LbAccessLogsOutputReference
	_jsii_.Get(
		j,
		"accessLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) AccessLogsInput() *LbAccessLogs {
	var returns *LbAccessLogs
	_jsii_.Get(
		j,
		"accessLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) ArnSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) ClientKeepAlive() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) ClientKeepAliveInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) ConnectionLogs() LbConnectionLogsOutputReference {
	var returns LbConnectionLogsOutputReference
	_jsii_.Get(
		j,
		"connectionLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) ConnectionLogsInput() *LbConnectionLogs {
	var returns *LbConnectionLogs
	_jsii_.Get(
		j,
		"connectionLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) CustomerOwnedIpv4Pool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerOwnedIpv4Pool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) CustomerOwnedIpv4PoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerOwnedIpv4PoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) DesyncMitigationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desyncMitigationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) DesyncMitigationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desyncMitigationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) DnsRecordClientRoutingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsRecordClientRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) DnsRecordClientRoutingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsRecordClientRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) DropInvalidHeaderFields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropInvalidHeaderFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) DropInvalidHeaderFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropInvalidHeaderFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) EnableCrossZoneLoadBalancing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCrossZoneLoadBalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) EnableCrossZoneLoadBalancingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCrossZoneLoadBalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) EnableDeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDeletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) EnableDeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDeletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) EnableHttp2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttp2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) EnableHttp2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttp2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) EnableTlsVersionAndCipherSuiteHeaders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTlsVersionAndCipherSuiteHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) EnableTlsVersionAndCipherSuiteHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTlsVersionAndCipherSuiteHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) EnableWafFailOpen() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWafFailOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) EnableWafFailOpenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWafFailOpenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) EnableXffClientPort() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableXffClientPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) EnableXffClientPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableXffClientPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) EnforceSecurityGroupInboundRulesOnPrivateLinkTrafficInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceSecurityGroupInboundRulesOnPrivateLinkTrafficInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) IdleTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) IdleTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) Internal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) InternalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) LoadBalancerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) LoadBalancerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) PreserveHostHeader() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) PreserveHostHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveHostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) SubnetMapping() LbSubnetMappingList {
	var returns LbSubnetMappingList
	_jsii_.Get(
		j,
		"subnetMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) SubnetMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) SubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) Timeouts() LbTimeoutsOutputReference {
	var returns LbTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) XffHeaderProcessingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xffHeaderProcessingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) XffHeaderProcessingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xffHeaderProcessingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lb) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/lb aws_lb} Resource.
func NewLb(scope constructs.Construct, id *string, config *LbConfig) Lb {
	_init_.Initialize()

	if err := validateNewLbParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lb{}

	_jsii_.Create(
		"@cdktf/provider-aws.lb.Lb",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/lb aws_lb} Resource.
func NewLb_Override(l Lb, scope constructs.Construct, id *string, config *LbConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lb.Lb",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_Lb)SetClientKeepAlive(val *float64) {
	if err := j.validateSetClientKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientKeepAlive",
		val,
	)
}

func (j *jsiiProxy_Lb)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Lb)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Lb)SetCustomerOwnedIpv4Pool(val *string) {
	if err := j.validateSetCustomerOwnedIpv4PoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerOwnedIpv4Pool",
		val,
	)
}

func (j *jsiiProxy_Lb)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Lb)SetDesyncMitigationMode(val *string) {
	if err := j.validateSetDesyncMitigationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desyncMitigationMode",
		val,
	)
}

func (j *jsiiProxy_Lb)SetDnsRecordClientRoutingPolicy(val *string) {
	if err := j.validateSetDnsRecordClientRoutingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsRecordClientRoutingPolicy",
		val,
	)
}

func (j *jsiiProxy_Lb)SetDropInvalidHeaderFields(val interface{}) {
	if err := j.validateSetDropInvalidHeaderFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dropInvalidHeaderFields",
		val,
	)
}

func (j *jsiiProxy_Lb)SetEnableCrossZoneLoadBalancing(val interface{}) {
	if err := j.validateSetEnableCrossZoneLoadBalancingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCrossZoneLoadBalancing",
		val,
	)
}

func (j *jsiiProxy_Lb)SetEnableDeletionProtection(val interface{}) {
	if err := j.validateSetEnableDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDeletionProtection",
		val,
	)
}

func (j *jsiiProxy_Lb)SetEnableHttp2(val interface{}) {
	if err := j.validateSetEnableHttp2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHttp2",
		val,
	)
}

func (j *jsiiProxy_Lb)SetEnableTlsVersionAndCipherSuiteHeaders(val interface{}) {
	if err := j.validateSetEnableTlsVersionAndCipherSuiteHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTlsVersionAndCipherSuiteHeaders",
		val,
	)
}

func (j *jsiiProxy_Lb)SetEnableWafFailOpen(val interface{}) {
	if err := j.validateSetEnableWafFailOpenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableWafFailOpen",
		val,
	)
}

func (j *jsiiProxy_Lb)SetEnableXffClientPort(val interface{}) {
	if err := j.validateSetEnableXffClientPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableXffClientPort",
		val,
	)
}

func (j *jsiiProxy_Lb)SetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic(val *string) {
	if err := j.validateSetEnforceSecurityGroupInboundRulesOnPrivateLinkTrafficParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic",
		val,
	)
}

func (j *jsiiProxy_Lb)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Lb)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Lb)SetIdleTimeout(val *float64) {
	if err := j.validateSetIdleTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeout",
		val,
	)
}

func (j *jsiiProxy_Lb)SetInternal(val interface{}) {
	if err := j.validateSetInternalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internal",
		val,
	)
}

func (j *jsiiProxy_Lb)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_Lb)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Lb)SetLoadBalancerType(val *string) {
	if err := j.validateSetLoadBalancerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerType",
		val,
	)
}

func (j *jsiiProxy_Lb)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Lb)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_Lb)SetPreserveHostHeader(val interface{}) {
	if err := j.validateSetPreserveHostHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveHostHeader",
		val,
	)
}

func (j *jsiiProxy_Lb)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Lb)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Lb)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_Lb)SetSubnets(val *[]*string) {
	if err := j.validateSetSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_Lb)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Lb)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_Lb)SetXffHeaderProcessingMode(val *string) {
	if err := j.validateSetXffHeaderProcessingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xffHeaderProcessingMode",
		val,
	)
}

// Generates CDKTF code for importing a Lb resource upon running "cdktf plan <stack-name>".
func Lb_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLb_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lb.Lb",
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
func Lb_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLb_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lb.Lb",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Lb_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLb_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lb.Lb",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Lb_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLb_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lb.Lb",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Lb_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.lb.Lb",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_Lb) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_Lb) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_Lb) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_Lb) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_Lb) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_Lb) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_Lb) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_Lb) PutAccessLogs(value *LbAccessLogs) {
	if err := l.validatePutAccessLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAccessLogs",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lb) PutConnectionLogs(value *LbConnectionLogs) {
	if err := l.validatePutConnectionLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putConnectionLogs",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lb) PutSubnetMapping(value interface{}) {
	if err := l.validatePutSubnetMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSubnetMapping",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lb) PutTimeouts(value *LbTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lb) ResetAccessLogs() {
	_jsii_.InvokeVoid(
		l,
		"resetAccessLogs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetClientKeepAlive() {
	_jsii_.InvokeVoid(
		l,
		"resetClientKeepAlive",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetConnectionLogs() {
	_jsii_.InvokeVoid(
		l,
		"resetConnectionLogs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetCustomerOwnedIpv4Pool() {
	_jsii_.InvokeVoid(
		l,
		"resetCustomerOwnedIpv4Pool",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetDesyncMitigationMode() {
	_jsii_.InvokeVoid(
		l,
		"resetDesyncMitigationMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetDnsRecordClientRoutingPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetDnsRecordClientRoutingPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetDropInvalidHeaderFields() {
	_jsii_.InvokeVoid(
		l,
		"resetDropInvalidHeaderFields",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetEnableCrossZoneLoadBalancing() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableCrossZoneLoadBalancing",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetEnableDeletionProtection() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableDeletionProtection",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetEnableHttp2() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableHttp2",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetEnableTlsVersionAndCipherSuiteHeaders() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableTlsVersionAndCipherSuiteHeaders",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetEnableWafFailOpen() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableWafFailOpen",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetEnableXffClientPort() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableXffClientPort",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic() {
	_jsii_.InvokeVoid(
		l,
		"resetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetIdleTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetIdleTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetInternal() {
	_jsii_.InvokeVoid(
		l,
		"resetInternal",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		l,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetLoadBalancerType() {
	_jsii_.InvokeVoid(
		l,
		"resetLoadBalancerType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetName() {
	_jsii_.InvokeVoid(
		l,
		"resetName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		l,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetPreserveHostHeader() {
	_jsii_.InvokeVoid(
		l,
		"resetPreserveHostHeader",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		l,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetSubnetMapping() {
	_jsii_.InvokeVoid(
		l,
		"resetSubnetMapping",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetSubnets() {
	_jsii_.InvokeVoid(
		l,
		"resetSubnets",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetTagsAll() {
	_jsii_.InvokeVoid(
		l,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) ResetXffHeaderProcessingMode() {
	_jsii_.InvokeVoid(
		l,
		"resetXffHeaderProcessingMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lb) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lb) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

