// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package albtargetgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/albtargetgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/alb_target_group aws_alb_target_group}.
type AlbTargetGroup interface {
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
	HealthCheck() AlbTargetGroupHealthCheckOutputReference
	HealthCheckInput() *AlbTargetGroupHealthCheck
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
	LoadBalancingAlgorithmType() *string
	SetLoadBalancingAlgorithmType(val *string)
	LoadBalancingAlgorithmTypeInput() *string
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
	Stickiness() AlbTargetGroupStickinessOutputReference
	StickinessInput() *AlbTargetGroupStickiness
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TargetFailover() AlbTargetGroupTargetFailoverList
	TargetFailoverInput() interface{}
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutHealthCheck(value *AlbTargetGroupHealthCheck)
	PutStickiness(value *AlbTargetGroupStickiness)
	PutTargetFailover(value interface{})
	ResetConnectionTermination()
	ResetDeregistrationDelay()
	ResetHealthCheck()
	ResetId()
	ResetIpAddressType()
	ResetLambdaMultiValueHeadersEnabled()
	ResetLoadBalancingAlgorithmType()
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
	ResetTargetType()
	ResetVpcId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AlbTargetGroup
type jsiiProxy_AlbTargetGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AlbTargetGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) ArnSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) ConnectionTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) ConnectionTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) DeregistrationDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deregistrationDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) DeregistrationDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deregistrationDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) HealthCheck() AlbTargetGroupHealthCheckOutputReference {
	var returns AlbTargetGroupHealthCheckOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) HealthCheckInput() *AlbTargetGroupHealthCheck {
	var returns *AlbTargetGroupHealthCheck
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) LambdaMultiValueHeadersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaMultiValueHeadersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) LambdaMultiValueHeadersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaMultiValueHeadersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) LoadBalancingAlgorithmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingAlgorithmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) LoadBalancingAlgorithmTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingAlgorithmTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) LoadBalancingCrossZoneEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingCrossZoneEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) LoadBalancingCrossZoneEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingCrossZoneEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) PreserveClientIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveClientIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) PreserveClientIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveClientIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) ProtocolVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) ProtocolVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) ProxyProtocolV2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyProtocolV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) ProxyProtocolV2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyProtocolV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) SlowStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) SlowStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slowStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) Stickiness() AlbTargetGroupStickinessOutputReference {
	var returns AlbTargetGroupStickinessOutputReference
	_jsii_.Get(
		j,
		"stickiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) StickinessInput() *AlbTargetGroupStickiness {
	var returns *AlbTargetGroupStickiness
	_jsii_.Get(
		j,
		"stickinessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) TargetFailover() AlbTargetGroupTargetFailoverList {
	var returns AlbTargetGroupTargetFailoverList
	_jsii_.Get(
		j,
		"targetFailover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) TargetFailoverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetFailoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) TargetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) TargetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbTargetGroup) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/alb_target_group aws_alb_target_group} Resource.
func NewAlbTargetGroup(scope constructs.Construct, id *string, config *AlbTargetGroupConfig) AlbTargetGroup {
	_init_.Initialize()

	if err := validateNewAlbTargetGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlbTargetGroup{}

	_jsii_.Create(
		"@cdktf/provider-aws.albTargetGroup.AlbTargetGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/alb_target_group aws_alb_target_group} Resource.
func NewAlbTargetGroup_Override(a AlbTargetGroup, scope constructs.Construct, id *string, config *AlbTargetGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.albTargetGroup.AlbTargetGroup",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetConnectionTermination(val interface{}) {
	if err := j.validateSetConnectionTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionTermination",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetDeregistrationDelay(val *string) {
	if err := j.validateSetDeregistrationDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deregistrationDelay",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetLambdaMultiValueHeadersEnabled(val interface{}) {
	if err := j.validateSetLambdaMultiValueHeadersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaMultiValueHeadersEnabled",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetLoadBalancingAlgorithmType(val *string) {
	if err := j.validateSetLoadBalancingAlgorithmTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingAlgorithmType",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetLoadBalancingCrossZoneEnabled(val *string) {
	if err := j.validateSetLoadBalancingCrossZoneEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingCrossZoneEnabled",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetPreserveClientIp(val *string) {
	if err := j.validateSetPreserveClientIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveClientIp",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetProtocolVersion(val *string) {
	if err := j.validateSetProtocolVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolVersion",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetProxyProtocolV2(val interface{}) {
	if err := j.validateSetProxyProtocolV2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyProtocolV2",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetSlowStart(val *float64) {
	if err := j.validateSetSlowStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slowStart",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetTargetType(val *string) {
	if err := j.validateSetTargetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

func (j *jsiiProxy_AlbTargetGroup)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a AlbTargetGroup resource upon running "cdktf plan <stack-name>".
func AlbTargetGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAlbTargetGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.albTargetGroup.AlbTargetGroup",
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
func AlbTargetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlbTargetGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.albTargetGroup.AlbTargetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AlbTargetGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlbTargetGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.albTargetGroup.AlbTargetGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AlbTargetGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlbTargetGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.albTargetGroup.AlbTargetGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AlbTargetGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.albTargetGroup.AlbTargetGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AlbTargetGroup) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AlbTargetGroup) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AlbTargetGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AlbTargetGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlbTargetGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AlbTargetGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AlbTargetGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AlbTargetGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AlbTargetGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AlbTargetGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AlbTargetGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AlbTargetGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AlbTargetGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlbTargetGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AlbTargetGroup) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AlbTargetGroup) PutHealthCheck(value *AlbTargetGroupHealthCheck) {
	if err := a.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbTargetGroup) PutStickiness(value *AlbTargetGroupStickiness) {
	if err := a.validatePutStickinessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStickiness",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbTargetGroup) PutTargetFailover(value interface{}) {
	if err := a.validatePutTargetFailoverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTargetFailover",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetConnectionTermination() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionTermination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetDeregistrationDelay() {
	_jsii_.InvokeVoid(
		a,
		"resetDeregistrationDelay",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		a,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetLambdaMultiValueHeadersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaMultiValueHeadersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetLoadBalancingAlgorithmType() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancingAlgorithmType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetLoadBalancingCrossZoneEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancingCrossZoneEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetPreserveClientIp() {
	_jsii_.InvokeVoid(
		a,
		"resetPreserveClientIp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetProtocolVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocolVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetProxyProtocolV2() {
	_jsii_.InvokeVoid(
		a,
		"resetProxyProtocolV2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetSlowStart() {
	_jsii_.InvokeVoid(
		a,
		"resetSlowStart",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetStickiness() {
	_jsii_.InvokeVoid(
		a,
		"resetStickiness",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetTargetFailover() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetFailover",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetTargetType() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) ResetVpcId() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbTargetGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbTargetGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbTargetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbTargetGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

