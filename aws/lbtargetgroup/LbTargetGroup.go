// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbtargetgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/lbtargetgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/lb_target_group aws_lb_target_group}.
type LbTargetGroup interface {
	cdktf.TerraformResource
	Arn() *string
	ArnSuffix() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionTermination() interface{}
	SetConnectionTermination(val interface{})
	ConnectionTerminationInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeregistrationDelay() *string
	SetDeregistrationDelay(val *string)
	DeregistrationDelayInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthCheck() LbTargetGroupHealthCheckOutputReference
	HealthCheckInput() *LbTargetGroupHealthCheck
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpAddressType() *string
	SetIpAddressType(val *string)
	IpAddressTypeInput() *string
	LambdaMultiValueHeadersEnabled() interface{}
	SetLambdaMultiValueHeadersEnabled(val interface{})
	LambdaMultiValueHeadersEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancerArns() *[]*string
	LoadBalancingAlgorithmType() *string
	SetLoadBalancingAlgorithmType(val *string)
	LoadBalancingAlgorithmTypeInput() *string
	LoadBalancingAnomalyMitigation() *string
	SetLoadBalancingAnomalyMitigation(val *string)
	LoadBalancingAnomalyMitigationInput() *string
	LoadBalancingCrossZoneEnabled() *string
	SetLoadBalancingCrossZoneEnabled(val *string)
	LoadBalancingCrossZoneEnabledInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	// The tree node.
	Node() constructs.Node
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PreserveClientIp() *string
	SetPreserveClientIp(val *string)
	PreserveClientIpInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	ProtocolVersion() *string
	SetProtocolVersion(val *string)
	ProtocolVersionInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProxyProtocolV2() interface{}
	SetProxyProtocolV2(val interface{})
	ProxyProtocolV2Input() interface{}
	// Experimental.
	RawOverrides() interface{}
	SlowStart() *float64
	SetSlowStart(val *float64)
	SlowStartInput() *float64
	Stickiness() LbTargetGroupStickinessOutputReference
	StickinessInput() *LbTargetGroupStickiness
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TargetFailover() LbTargetGroupTargetFailoverList
	TargetFailoverInput() interface{}
	TargetGroupHealth() LbTargetGroupTargetGroupHealthOutputReference
	TargetGroupHealthInput() *LbTargetGroupTargetGroupHealth
	TargetHealthState() LbTargetGroupTargetHealthStateList
	TargetHealthStateInput() interface{}
	TargetType() *string
	SetTargetType(val *string)
	TargetTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
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
	PutHealthCheck(value *LbTargetGroupHealthCheck)
	PutStickiness(value *LbTargetGroupStickiness)
	PutTargetFailover(value interface{})
	PutTargetGroupHealth(value *LbTargetGroupTargetGroupHealth)
	PutTargetHealthState(value interface{})
	ResetConnectionTermination()
	ResetDeregistrationDelay()
	ResetHealthCheck()
	ResetId()
	ResetIpAddressType()
	ResetLambdaMultiValueHeadersEnabled()
	ResetLoadBalancingAlgorithmType()
	ResetLoadBalancingAnomalyMitigation()
	ResetLoadBalancingCrossZoneEnabled()
	ResetName()
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetPreserveClientIp()
	ResetProtocol()
	ResetProtocolVersion()
	ResetProxyProtocolV2()
	ResetSlowStart()
	ResetStickiness()
	ResetTags()
	ResetTagsAll()
	ResetTargetFailover()
	ResetTargetGroupHealth()
	ResetTargetHealthState()
	ResetTargetType()
	ResetVpcId()
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

// The jsii proxy struct for LbTargetGroup
type jsiiProxy_LbTargetGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LbTargetGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) ArnSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) ConnectionTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) ConnectionTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) DeregistrationDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deregistrationDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) DeregistrationDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deregistrationDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) HealthCheck() LbTargetGroupHealthCheckOutputReference {
	var returns LbTargetGroupHealthCheckOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) HealthCheckInput() *LbTargetGroupHealthCheck {
	var returns *LbTargetGroupHealthCheck
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) LambdaMultiValueHeadersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaMultiValueHeadersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) LambdaMultiValueHeadersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaMultiValueHeadersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) LoadBalancerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) LoadBalancingAlgorithmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingAlgorithmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) LoadBalancingAlgorithmTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingAlgorithmTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) LoadBalancingAnomalyMitigation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingAnomalyMitigation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) LoadBalancingAnomalyMitigationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingAnomalyMitigationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) LoadBalancingCrossZoneEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingCrossZoneEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) LoadBalancingCrossZoneEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingCrossZoneEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) PreserveClientIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveClientIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) PreserveClientIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveClientIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) ProtocolVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) ProtocolVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) ProxyProtocolV2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyProtocolV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) ProxyProtocolV2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyProtocolV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) SlowStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) SlowStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slowStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) Stickiness() LbTargetGroupStickinessOutputReference {
	var returns LbTargetGroupStickinessOutputReference
	_jsii_.Get(
		j,
		"stickiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) StickinessInput() *LbTargetGroupStickiness {
	var returns *LbTargetGroupStickiness
	_jsii_.Get(
		j,
		"stickinessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) TargetFailover() LbTargetGroupTargetFailoverList {
	var returns LbTargetGroupTargetFailoverList
	_jsii_.Get(
		j,
		"targetFailover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) TargetFailoverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetFailoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) TargetGroupHealth() LbTargetGroupTargetGroupHealthOutputReference {
	var returns LbTargetGroupTargetGroupHealthOutputReference
	_jsii_.Get(
		j,
		"targetGroupHealth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) TargetGroupHealthInput() *LbTargetGroupTargetGroupHealth {
	var returns *LbTargetGroupTargetGroupHealth
	_jsii_.Get(
		j,
		"targetGroupHealthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) TargetHealthState() LbTargetGroupTargetHealthStateList {
	var returns LbTargetGroupTargetHealthStateList
	_jsii_.Get(
		j,
		"targetHealthState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) TargetHealthStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetHealthStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) TargetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) TargetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroup) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/lb_target_group aws_lb_target_group} Resource.
func NewLbTargetGroup(scope constructs.Construct, id *string, config *LbTargetGroupConfig) LbTargetGroup {
	_init_.Initialize()

	if err := validateNewLbTargetGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbTargetGroup{}

	_jsii_.Create(
		"@cdktf/provider-aws.lbTargetGroup.LbTargetGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/lb_target_group aws_lb_target_group} Resource.
func NewLbTargetGroup_Override(l LbTargetGroup, scope constructs.Construct, id *string, config *LbTargetGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lbTargetGroup.LbTargetGroup",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetConnectionTermination(val interface{}) {
	if err := j.validateSetConnectionTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionTermination",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetDeregistrationDelay(val *string) {
	if err := j.validateSetDeregistrationDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deregistrationDelay",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetLambdaMultiValueHeadersEnabled(val interface{}) {
	if err := j.validateSetLambdaMultiValueHeadersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaMultiValueHeadersEnabled",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetLoadBalancingAlgorithmType(val *string) {
	if err := j.validateSetLoadBalancingAlgorithmTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingAlgorithmType",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetLoadBalancingAnomalyMitigation(val *string) {
	if err := j.validateSetLoadBalancingAnomalyMitigationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingAnomalyMitigation",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetLoadBalancingCrossZoneEnabled(val *string) {
	if err := j.validateSetLoadBalancingCrossZoneEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingCrossZoneEnabled",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetPreserveClientIp(val *string) {
	if err := j.validateSetPreserveClientIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveClientIp",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetProtocolVersion(val *string) {
	if err := j.validateSetProtocolVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolVersion",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetProxyProtocolV2(val interface{}) {
	if err := j.validateSetProxyProtocolV2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyProtocolV2",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetSlowStart(val *float64) {
	if err := j.validateSetSlowStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slowStart",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetTargetType(val *string) {
	if err := j.validateSetTargetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroup)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a LbTargetGroup resource upon running "cdktf plan <stack-name>".
func LbTargetGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLbTargetGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lbTargetGroup.LbTargetGroup",
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
func LbTargetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbTargetGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lbTargetGroup.LbTargetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbTargetGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbTargetGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lbTargetGroup.LbTargetGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbTargetGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbTargetGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lbTargetGroup.LbTargetGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LbTargetGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.lbTargetGroup.LbTargetGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LbTargetGroup) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LbTargetGroup) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LbTargetGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LbTargetGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbTargetGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LbTargetGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LbTargetGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LbTargetGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LbTargetGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LbTargetGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LbTargetGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LbTargetGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LbTargetGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbTargetGroup) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LbTargetGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LbTargetGroup) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LbTargetGroup) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LbTargetGroup) PutHealthCheck(value *LbTargetGroupHealthCheck) {
	if err := l.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbTargetGroup) PutStickiness(value *LbTargetGroupStickiness) {
	if err := l.validatePutStickinessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putStickiness",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbTargetGroup) PutTargetFailover(value interface{}) {
	if err := l.validatePutTargetFailoverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTargetFailover",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbTargetGroup) PutTargetGroupHealth(value *LbTargetGroupTargetGroupHealth) {
	if err := l.validatePutTargetGroupHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTargetGroupHealth",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbTargetGroup) PutTargetHealthState(value interface{}) {
	if err := l.validatePutTargetHealthStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTargetHealthState",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetConnectionTermination() {
	_jsii_.InvokeVoid(
		l,
		"resetConnectionTermination",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetDeregistrationDelay() {
	_jsii_.InvokeVoid(
		l,
		"resetDeregistrationDelay",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		l,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		l,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetLambdaMultiValueHeadersEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetLambdaMultiValueHeadersEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetLoadBalancingAlgorithmType() {
	_jsii_.InvokeVoid(
		l,
		"resetLoadBalancingAlgorithmType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetLoadBalancingAnomalyMitigation() {
	_jsii_.InvokeVoid(
		l,
		"resetLoadBalancingAnomalyMitigation",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetLoadBalancingCrossZoneEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetLoadBalancingCrossZoneEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetName() {
	_jsii_.InvokeVoid(
		l,
		"resetName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		l,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetPort() {
	_jsii_.InvokeVoid(
		l,
		"resetPort",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetPreserveClientIp() {
	_jsii_.InvokeVoid(
		l,
		"resetPreserveClientIp",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetProtocol() {
	_jsii_.InvokeVoid(
		l,
		"resetProtocol",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetProtocolVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetProtocolVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetProxyProtocolV2() {
	_jsii_.InvokeVoid(
		l,
		"resetProxyProtocolV2",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetSlowStart() {
	_jsii_.InvokeVoid(
		l,
		"resetSlowStart",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetStickiness() {
	_jsii_.InvokeVoid(
		l,
		"resetStickiness",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		l,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetTargetFailover() {
	_jsii_.InvokeVoid(
		l,
		"resetTargetFailover",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetTargetGroupHealth() {
	_jsii_.InvokeVoid(
		l,
		"resetTargetGroupHealth",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetTargetHealthState() {
	_jsii_.InvokeVoid(
		l,
		"resetTargetHealthState",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetTargetType() {
	_jsii_.InvokeVoid(
		l,
		"resetTargetType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) ResetVpcId() {
	_jsii_.InvokeVoid(
		l,
		"resetVpcId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

