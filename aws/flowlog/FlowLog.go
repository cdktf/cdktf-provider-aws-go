// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package flowlog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/flowlog/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/flow_log aws_flow_log}.
type FlowLog interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DeliverCrossAccountRole() *string
	SetDeliverCrossAccountRole(val *string)
	DeliverCrossAccountRoleInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DestinationOptions() FlowLogDestinationOptionsOutputReference
	DestinationOptionsInput() *FlowLogDestinationOptions
	EniId() *string
	SetEniId(val *string)
	EniIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IamRoleArn() *string
	SetIamRoleArn(val *string)
	IamRoleArnInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogDestination() *string
	SetLogDestination(val *string)
	LogDestinationInput() *string
	LogDestinationType() *string
	SetLogDestinationType(val *string)
	LogDestinationTypeInput() *string
	LogFormat() *string
	SetLogFormat(val *string)
	LogFormatInput() *string
	LogGroupName() *string
	SetLogGroupName(val *string)
	LogGroupNameInput() *string
	MaxAggregationInterval() *float64
	SetMaxAggregationInterval(val *float64)
	MaxAggregationIntervalInput() *float64
	// The tree node.
	Node() constructs.Node
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
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
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
	TrafficType() *string
	SetTrafficType(val *string)
	TrafficTypeInput() *string
	TransitGatewayAttachmentId() *string
	SetTransitGatewayAttachmentId(val *string)
	TransitGatewayAttachmentIdInput() *string
	TransitGatewayId() *string
	SetTransitGatewayId(val *string)
	TransitGatewayIdInput() *string
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
	PutDestinationOptions(value *FlowLogDestinationOptions)
	ResetDeliverCrossAccountRole()
	ResetDestinationOptions()
	ResetEniId()
	ResetIamRoleArn()
	ResetId()
	ResetLogDestination()
	ResetLogDestinationType()
	ResetLogFormat()
	ResetLogGroupName()
	ResetMaxAggregationInterval()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSubnetId()
	ResetTags()
	ResetTagsAll()
	ResetTrafficType()
	ResetTransitGatewayAttachmentId()
	ResetTransitGatewayId()
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

// The jsii proxy struct for FlowLog
type jsiiProxy_FlowLog struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FlowLog) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) DeliverCrossAccountRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliverCrossAccountRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) DeliverCrossAccountRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliverCrossAccountRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) DestinationOptions() FlowLogDestinationOptionsOutputReference {
	var returns FlowLogDestinationOptionsOutputReference
	_jsii_.Get(
		j,
		"destinationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) DestinationOptionsInput() *FlowLogDestinationOptions {
	var returns *FlowLogDestinationOptions
	_jsii_.Get(
		j,
		"destinationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) EniId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eniId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) EniIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eniIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) IamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) IamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) LogDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) LogDestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) LogDestinationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logDestinationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) LogDestinationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logDestinationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) LogFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) LogFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) LogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) MaxAggregationInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAggregationInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) MaxAggregationIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAggregationIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) TrafficType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) TrafficTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) TransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) TransitGatewayAttachmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAttachmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) TransitGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/flow_log aws_flow_log} Resource.
func NewFlowLog(scope constructs.Construct, id *string, config *FlowLogConfig) FlowLog {
	_init_.Initialize()

	if err := validateNewFlowLogParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FlowLog{}

	_jsii_.Create(
		"@cdktf/provider-aws.flowLog.FlowLog",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/flow_log aws_flow_log} Resource.
func NewFlowLog_Override(f FlowLog, scope constructs.Construct, id *string, config *FlowLogConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.flowLog.FlowLog",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FlowLog)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetDeliverCrossAccountRole(val *string) {
	if err := j.validateSetDeliverCrossAccountRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliverCrossAccountRole",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetEniId(val *string) {
	if err := j.validateSetEniIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eniId",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetIamRoleArn(val *string) {
	if err := j.validateSetIamRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoleArn",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetLogDestination(val *string) {
	if err := j.validateSetLogDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logDestination",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetLogDestinationType(val *string) {
	if err := j.validateSetLogDestinationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logDestinationType",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetLogFormat(val *string) {
	if err := j.validateSetLogFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logFormat",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetLogGroupName(val *string) {
	if err := j.validateSetLogGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetMaxAggregationInterval(val *float64) {
	if err := j.validateSetMaxAggregationIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAggregationInterval",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetTrafficType(val *string) {
	if err := j.validateSetTrafficTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficType",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetTransitGatewayAttachmentId(val *string) {
	if err := j.validateSetTransitGatewayAttachmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayAttachmentId",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetTransitGatewayId(val *string) {
	if err := j.validateSetTransitGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayId",
		val,
	)
}

func (j *jsiiProxy_FlowLog)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a FlowLog resource upon running "cdktf plan <stack-name>".
func FlowLog_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFlowLog_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.flowLog.FlowLog",
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
func FlowLog_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFlowLog_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.flowLog.FlowLog",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FlowLog_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFlowLog_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.flowLog.FlowLog",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FlowLog_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFlowLog_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.flowLog.FlowLog",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FlowLog_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.flowLog.FlowLog",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FlowLog) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FlowLog) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FlowLog) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FlowLog) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FlowLog) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FlowLog) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FlowLog) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FlowLog) PutDestinationOptions(value *FlowLogDestinationOptions) {
	if err := f.validatePutDestinationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putDestinationOptions",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FlowLog) ResetDeliverCrossAccountRole() {
	_jsii_.InvokeVoid(
		f,
		"resetDeliverCrossAccountRole",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetDestinationOptions() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationOptions",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetEniId() {
	_jsii_.InvokeVoid(
		f,
		"resetEniId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetIamRoleArn() {
	_jsii_.InvokeVoid(
		f,
		"resetIamRoleArn",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetLogDestination() {
	_jsii_.InvokeVoid(
		f,
		"resetLogDestination",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetLogDestinationType() {
	_jsii_.InvokeVoid(
		f,
		"resetLogDestinationType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetLogFormat() {
	_jsii_.InvokeVoid(
		f,
		"resetLogFormat",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetLogGroupName() {
	_jsii_.InvokeVoid(
		f,
		"resetLogGroupName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetMaxAggregationInterval() {
	_jsii_.InvokeVoid(
		f,
		"resetMaxAggregationInterval",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetSubnetId() {
	_jsii_.InvokeVoid(
		f,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetTagsAll() {
	_jsii_.InvokeVoid(
		f,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetTrafficType() {
	_jsii_.InvokeVoid(
		f,
		"resetTrafficType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetTransitGatewayAttachmentId() {
	_jsii_.InvokeVoid(
		f,
		"resetTransitGatewayAttachmentId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetTransitGatewayId() {
	_jsii_.InvokeVoid(
		f,
		"resetTransitGatewayId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) ResetVpcId() {
	_jsii_.InvokeVoid(
		f,
		"resetVpcId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLog) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

