// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53healthcheck

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/route53healthcheck/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/route53_health_check aws_route53_health_check}.
type Route53HealthCheck interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChildHealthchecks() *[]*string
	SetChildHealthchecks(val *[]*string)
	ChildHealthchecksInput() *[]*string
	ChildHealthThreshold() *float64
	SetChildHealthThreshold(val *float64)
	ChildHealthThresholdInput() *float64
	CloudwatchAlarmName() *string
	SetCloudwatchAlarmName(val *string)
	CloudwatchAlarmNameInput() *string
	CloudwatchAlarmRegion() *string
	SetCloudwatchAlarmRegion(val *string)
	CloudwatchAlarmRegionInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
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
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	EnableSni() interface{}
	SetEnableSni(val interface{})
	EnableSniInput() interface{}
	FailureThreshold() *float64
	SetFailureThreshold(val *float64)
	FailureThresholdInput() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Fqdn() *string
	SetFqdn(val *string)
	FqdnInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InsufficientDataHealthStatus() *string
	SetInsufficientDataHealthStatus(val *string)
	InsufficientDataHealthStatusInput() *string
	InvertHealthcheck() interface{}
	SetInvertHealthcheck(val interface{})
	InvertHealthcheckInput() interface{}
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MeasureLatency() interface{}
	SetMeasureLatency(val interface{})
	MeasureLatencyInput() interface{}
	// The tree node.
	Node() constructs.Node
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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
	ReferenceName() *string
	SetReferenceName(val *string)
	ReferenceNameInput() *string
	Regions() *[]*string
	SetRegions(val *[]*string)
	RegionsInput() *[]*string
	RequestInterval() *float64
	SetRequestInterval(val *float64)
	RequestIntervalInput() *float64
	ResourcePath() *string
	SetResourcePath(val *string)
	ResourcePathInput() *string
	RoutingControlArn() *string
	SetRoutingControlArn(val *string)
	RoutingControlArnInput() *string
	SearchString() *string
	SetSearchString(val *string)
	SearchStringInput() *string
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
	Triggers() *map[string]*string
	SetTriggers(val *map[string]*string)
	TriggersInput() *map[string]*string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetChildHealthchecks()
	ResetChildHealthThreshold()
	ResetCloudwatchAlarmName()
	ResetCloudwatchAlarmRegion()
	ResetDisabled()
	ResetEnableSni()
	ResetFailureThreshold()
	ResetFqdn()
	ResetId()
	ResetInsufficientDataHealthStatus()
	ResetInvertHealthcheck()
	ResetIpAddress()
	ResetMeasureLatency()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetReferenceName()
	ResetRegions()
	ResetRequestInterval()
	ResetResourcePath()
	ResetRoutingControlArn()
	ResetSearchString()
	ResetTags()
	ResetTagsAll()
	ResetTriggers()
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

// The jsii proxy struct for Route53HealthCheck
type jsiiProxy_Route53HealthCheck struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Route53HealthCheck) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) ChildHealthchecks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"childHealthchecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) ChildHealthchecksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"childHealthchecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) ChildHealthThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"childHealthThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) ChildHealthThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"childHealthThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) CloudwatchAlarmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchAlarmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) CloudwatchAlarmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchAlarmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) CloudwatchAlarmRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchAlarmRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) CloudwatchAlarmRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchAlarmRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) EnableSni() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSni",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) EnableSniInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSniInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) FailureThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) FailureThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) FqdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) InsufficientDataHealthStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insufficientDataHealthStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) InsufficientDataHealthStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insufficientDataHealthStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) InvertHealthcheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invertHealthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) InvertHealthcheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invertHealthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) MeasureLatency() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"measureLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) MeasureLatencyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"measureLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) ReferenceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) ReferenceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) RequestInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) RequestIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) ResourcePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) ResourcePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) RoutingControlArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingControlArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) RoutingControlArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingControlArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) SearchString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) SearchStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) Triggers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) TriggersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheck) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/route53_health_check aws_route53_health_check} Resource.
func NewRoute53HealthCheck(scope constructs.Construct, id *string, config *Route53HealthCheckConfig) Route53HealthCheck {
	_init_.Initialize()

	if err := validateNewRoute53HealthCheckParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53HealthCheck{}

	_jsii_.Create(
		"@cdktf/provider-aws.route53HealthCheck.Route53HealthCheck",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/route53_health_check aws_route53_health_check} Resource.
func NewRoute53HealthCheck_Override(r Route53HealthCheck, scope constructs.Construct, id *string, config *Route53HealthCheckConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.route53HealthCheck.Route53HealthCheck",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetChildHealthchecks(val *[]*string) {
	if err := j.validateSetChildHealthchecksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"childHealthchecks",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetChildHealthThreshold(val *float64) {
	if err := j.validateSetChildHealthThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"childHealthThreshold",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetCloudwatchAlarmName(val *string) {
	if err := j.validateSetCloudwatchAlarmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchAlarmName",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetCloudwatchAlarmRegion(val *string) {
	if err := j.validateSetCloudwatchAlarmRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchAlarmRegion",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetEnableSni(val interface{}) {
	if err := j.validateSetEnableSniParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSni",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetFailureThreshold(val *float64) {
	if err := j.validateSetFailureThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureThreshold",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetFqdn(val *string) {
	if err := j.validateSetFqdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fqdn",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetInsufficientDataHealthStatus(val *string) {
	if err := j.validateSetInsufficientDataHealthStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insufficientDataHealthStatus",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetInvertHealthcheck(val interface{}) {
	if err := j.validateSetInvertHealthcheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invertHealthcheck",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetMeasureLatency(val interface{}) {
	if err := j.validateSetMeasureLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measureLatency",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetReferenceName(val *string) {
	if err := j.validateSetReferenceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceName",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetRequestInterval(val *float64) {
	if err := j.validateSetRequestIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestInterval",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetResourcePath(val *string) {
	if err := j.validateSetResourcePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePath",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetRoutingControlArn(val *string) {
	if err := j.validateSetRoutingControlArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingControlArn",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetSearchString(val *string) {
	if err := j.validateSetSearchStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchString",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetTriggers(val *map[string]*string) {
	if err := j.validateSetTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggers",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheck)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a Route53HealthCheck resource upon running "cdktf plan <stack-name>".
func Route53HealthCheck_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRoute53HealthCheck_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53HealthCheck.Route53HealthCheck",
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
func Route53HealthCheck_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53HealthCheck_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53HealthCheck.Route53HealthCheck",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Route53HealthCheck_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53HealthCheck_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53HealthCheck.Route53HealthCheck",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Route53HealthCheck_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53HealthCheck_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53HealthCheck.Route53HealthCheck",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Route53HealthCheck_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.route53HealthCheck.Route53HealthCheck",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Route53HealthCheck) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_Route53HealthCheck) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_Route53HealthCheck) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_Route53HealthCheck) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Route53HealthCheck) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_Route53HealthCheck) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Route53HealthCheck) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetChildHealthchecks() {
	_jsii_.InvokeVoid(
		r,
		"resetChildHealthchecks",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetChildHealthThreshold() {
	_jsii_.InvokeVoid(
		r,
		"resetChildHealthThreshold",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetCloudwatchAlarmName() {
	_jsii_.InvokeVoid(
		r,
		"resetCloudwatchAlarmName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetCloudwatchAlarmRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetCloudwatchAlarmRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetDisabled() {
	_jsii_.InvokeVoid(
		r,
		"resetDisabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetEnableSni() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableSni",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetFailureThreshold() {
	_jsii_.InvokeVoid(
		r,
		"resetFailureThreshold",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetFqdn() {
	_jsii_.InvokeVoid(
		r,
		"resetFqdn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetInsufficientDataHealthStatus() {
	_jsii_.InvokeVoid(
		r,
		"resetInsufficientDataHealthStatus",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetInvertHealthcheck() {
	_jsii_.InvokeVoid(
		r,
		"resetInvertHealthcheck",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetIpAddress() {
	_jsii_.InvokeVoid(
		r,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetMeasureLatency() {
	_jsii_.InvokeVoid(
		r,
		"resetMeasureLatency",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetPort() {
	_jsii_.InvokeVoid(
		r,
		"resetPort",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetReferenceName() {
	_jsii_.InvokeVoid(
		r,
		"resetReferenceName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetRegions() {
	_jsii_.InvokeVoid(
		r,
		"resetRegions",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetRequestInterval() {
	_jsii_.InvokeVoid(
		r,
		"resetRequestInterval",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetResourcePath() {
	_jsii_.InvokeVoid(
		r,
		"resetResourcePath",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetRoutingControlArn() {
	_jsii_.InvokeVoid(
		r,
		"resetRoutingControlArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetSearchString() {
	_jsii_.InvokeVoid(
		r,
		"resetSearchString",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetTagsAll() {
	_jsii_.InvokeVoid(
		r,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) ResetTriggers() {
	_jsii_.InvokeVoid(
		r,
		"resetTriggers",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheck) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheck) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

