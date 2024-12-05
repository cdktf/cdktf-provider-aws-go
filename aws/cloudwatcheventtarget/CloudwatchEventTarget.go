// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventtarget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/cloudwatcheventtarget/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/cloudwatch_event_target aws_cloudwatch_event_target}.
type CloudwatchEventTarget interface {
	cdktf.TerraformResource
	AppsyncTarget() CloudwatchEventTargetAppsyncTargetOutputReference
	AppsyncTargetInput() *CloudwatchEventTargetAppsyncTarget
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
	BatchTarget() CloudwatchEventTargetBatchTargetOutputReference
	BatchTargetInput() *CloudwatchEventTargetBatchTarget
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
	DeadLetterConfig() CloudwatchEventTargetDeadLetterConfigOutputReference
	DeadLetterConfigInput() *CloudwatchEventTargetDeadLetterConfig
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EcsTarget() CloudwatchEventTargetEcsTargetOutputReference
	EcsTargetInput() *CloudwatchEventTargetEcsTarget
	EventBusName() *string
	SetEventBusName(val *string)
	EventBusNameInput() *string
	ForceDestroy() interface{}
	SetForceDestroy(val interface{})
	ForceDestroyInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpTarget() CloudwatchEventTargetHttpTargetOutputReference
	HttpTargetInput() *CloudwatchEventTargetHttpTarget
	Id() *string
	SetId(val *string)
	IdInput() *string
	Input() *string
	SetInput(val *string)
	InputInput() *string
	InputPath() *string
	SetInputPath(val *string)
	InputPathInput() *string
	InputTransformer() CloudwatchEventTargetInputTransformerOutputReference
	InputTransformerInput() *CloudwatchEventTargetInputTransformer
	KinesisTarget() CloudwatchEventTargetKinesisTargetOutputReference
	KinesisTargetInput() *CloudwatchEventTargetKinesisTarget
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	RedshiftTarget() CloudwatchEventTargetRedshiftTargetOutputReference
	RedshiftTargetInput() *CloudwatchEventTargetRedshiftTarget
	RetryPolicy() CloudwatchEventTargetRetryPolicyOutputReference
	RetryPolicyInput() *CloudwatchEventTargetRetryPolicy
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Rule() *string
	SetRule(val *string)
	RuleInput() *string
	RunCommandTargets() CloudwatchEventTargetRunCommandTargetsList
	RunCommandTargetsInput() interface{}
	SagemakerPipelineTarget() CloudwatchEventTargetSagemakerPipelineTargetOutputReference
	SagemakerPipelineTargetInput() *CloudwatchEventTargetSagemakerPipelineTarget
	SqsTarget() CloudwatchEventTargetSqsTargetOutputReference
	SqsTargetInput() *CloudwatchEventTargetSqsTarget
	TargetId() *string
	SetTargetId(val *string)
	TargetIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutAppsyncTarget(value *CloudwatchEventTargetAppsyncTarget)
	PutBatchTarget(value *CloudwatchEventTargetBatchTarget)
	PutDeadLetterConfig(value *CloudwatchEventTargetDeadLetterConfig)
	PutEcsTarget(value *CloudwatchEventTargetEcsTarget)
	PutHttpTarget(value *CloudwatchEventTargetHttpTarget)
	PutInputTransformer(value *CloudwatchEventTargetInputTransformer)
	PutKinesisTarget(value *CloudwatchEventTargetKinesisTarget)
	PutRedshiftTarget(value *CloudwatchEventTargetRedshiftTarget)
	PutRetryPolicy(value *CloudwatchEventTargetRetryPolicy)
	PutRunCommandTargets(value interface{})
	PutSagemakerPipelineTarget(value *CloudwatchEventTargetSagemakerPipelineTarget)
	PutSqsTarget(value *CloudwatchEventTargetSqsTarget)
	ResetAppsyncTarget()
	ResetBatchTarget()
	ResetDeadLetterConfig()
	ResetEcsTarget()
	ResetEventBusName()
	ResetForceDestroy()
	ResetHttpTarget()
	ResetId()
	ResetInput()
	ResetInputPath()
	ResetInputTransformer()
	ResetKinesisTarget()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRedshiftTarget()
	ResetRetryPolicy()
	ResetRoleArn()
	ResetRunCommandTargets()
	ResetSagemakerPipelineTarget()
	ResetSqsTarget()
	ResetTargetId()
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

// The jsii proxy struct for CloudwatchEventTarget
type jsiiProxy_CloudwatchEventTarget struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudwatchEventTarget) AppsyncTarget() CloudwatchEventTargetAppsyncTargetOutputReference {
	var returns CloudwatchEventTargetAppsyncTargetOutputReference
	_jsii_.Get(
		j,
		"appsyncTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) AppsyncTargetInput() *CloudwatchEventTargetAppsyncTarget {
	var returns *CloudwatchEventTargetAppsyncTarget
	_jsii_.Get(
		j,
		"appsyncTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) BatchTarget() CloudwatchEventTargetBatchTargetOutputReference {
	var returns CloudwatchEventTargetBatchTargetOutputReference
	_jsii_.Get(
		j,
		"batchTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) BatchTargetInput() *CloudwatchEventTargetBatchTarget {
	var returns *CloudwatchEventTargetBatchTarget
	_jsii_.Get(
		j,
		"batchTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) DeadLetterConfig() CloudwatchEventTargetDeadLetterConfigOutputReference {
	var returns CloudwatchEventTargetDeadLetterConfigOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) DeadLetterConfigInput() *CloudwatchEventTargetDeadLetterConfig {
	var returns *CloudwatchEventTargetDeadLetterConfig
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) EcsTarget() CloudwatchEventTargetEcsTargetOutputReference {
	var returns CloudwatchEventTargetEcsTargetOutputReference
	_jsii_.Get(
		j,
		"ecsTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) EcsTargetInput() *CloudwatchEventTargetEcsTarget {
	var returns *CloudwatchEventTargetEcsTarget
	_jsii_.Get(
		j,
		"ecsTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) EventBusName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) EventBusNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) HttpTarget() CloudwatchEventTargetHttpTargetOutputReference {
	var returns CloudwatchEventTargetHttpTargetOutputReference
	_jsii_.Get(
		j,
		"httpTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) HttpTargetInput() *CloudwatchEventTargetHttpTarget {
	var returns *CloudwatchEventTargetHttpTarget
	_jsii_.Get(
		j,
		"httpTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) InputPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) InputTransformer() CloudwatchEventTargetInputTransformerOutputReference {
	var returns CloudwatchEventTargetInputTransformerOutputReference
	_jsii_.Get(
		j,
		"inputTransformer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) InputTransformerInput() *CloudwatchEventTargetInputTransformer {
	var returns *CloudwatchEventTargetInputTransformer
	_jsii_.Get(
		j,
		"inputTransformerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) KinesisTarget() CloudwatchEventTargetKinesisTargetOutputReference {
	var returns CloudwatchEventTargetKinesisTargetOutputReference
	_jsii_.Get(
		j,
		"kinesisTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) KinesisTargetInput() *CloudwatchEventTargetKinesisTarget {
	var returns *CloudwatchEventTargetKinesisTarget
	_jsii_.Get(
		j,
		"kinesisTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) RedshiftTarget() CloudwatchEventTargetRedshiftTargetOutputReference {
	var returns CloudwatchEventTargetRedshiftTargetOutputReference
	_jsii_.Get(
		j,
		"redshiftTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) RedshiftTargetInput() *CloudwatchEventTargetRedshiftTarget {
	var returns *CloudwatchEventTargetRedshiftTarget
	_jsii_.Get(
		j,
		"redshiftTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) RetryPolicy() CloudwatchEventTargetRetryPolicyOutputReference {
	var returns CloudwatchEventTargetRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) RetryPolicyInput() *CloudwatchEventTargetRetryPolicy {
	var returns *CloudwatchEventTargetRetryPolicy
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) Rule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) RuleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) RunCommandTargets() CloudwatchEventTargetRunCommandTargetsList {
	var returns CloudwatchEventTargetRunCommandTargetsList
	_jsii_.Get(
		j,
		"runCommandTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) RunCommandTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runCommandTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) SagemakerPipelineTarget() CloudwatchEventTargetSagemakerPipelineTargetOutputReference {
	var returns CloudwatchEventTargetSagemakerPipelineTargetOutputReference
	_jsii_.Get(
		j,
		"sagemakerPipelineTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) SagemakerPipelineTargetInput() *CloudwatchEventTargetSagemakerPipelineTarget {
	var returns *CloudwatchEventTargetSagemakerPipelineTarget
	_jsii_.Get(
		j,
		"sagemakerPipelineTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) SqsTarget() CloudwatchEventTargetSqsTargetOutputReference {
	var returns CloudwatchEventTargetSqsTargetOutputReference
	_jsii_.Get(
		j,
		"sqsTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) SqsTargetInput() *CloudwatchEventTargetSqsTarget {
	var returns *CloudwatchEventTargetSqsTarget
	_jsii_.Get(
		j,
		"sqsTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) TargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) TargetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTarget) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/cloudwatch_event_target aws_cloudwatch_event_target} Resource.
func NewCloudwatchEventTarget(scope constructs.Construct, id *string, config *CloudwatchEventTargetConfig) CloudwatchEventTarget {
	_init_.Initialize()

	if err := validateNewCloudwatchEventTargetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudwatchEventTarget{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudwatchEventTarget.CloudwatchEventTarget",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/cloudwatch_event_target aws_cloudwatch_event_target} Resource.
func NewCloudwatchEventTarget_Override(c CloudwatchEventTarget, scope constructs.Construct, id *string, config *CloudwatchEventTargetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudwatchEventTarget.CloudwatchEventTarget",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetEventBusName(val *string) {
	if err := j.validateSetEventBusNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventBusName",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetInput(val *string) {
	if err := j.validateSetInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetInputPath(val *string) {
	if err := j.validateSetInputPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputPath",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetRule(val *string) {
	if err := j.validateSetRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rule",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTarget)SetTargetId(val *string) {
	if err := j.validateSetTargetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetId",
		val,
	)
}

// Generates CDKTF code for importing a CloudwatchEventTarget resource upon running "cdktf plan <stack-name>".
func CloudwatchEventTarget_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudwatchEventTarget_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cloudwatchEventTarget.CloudwatchEventTarget",
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
func CloudwatchEventTarget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudwatchEventTarget_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cloudwatchEventTarget.CloudwatchEventTarget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudwatchEventTarget_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudwatchEventTarget_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cloudwatchEventTarget.CloudwatchEventTarget",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudwatchEventTarget_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudwatchEventTarget_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cloudwatchEventTarget.CloudwatchEventTarget",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudwatchEventTarget_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.cloudwatchEventTarget.CloudwatchEventTarget",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) PutAppsyncTarget(value *CloudwatchEventTargetAppsyncTarget) {
	if err := c.validatePutAppsyncTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAppsyncTarget",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) PutBatchTarget(value *CloudwatchEventTargetBatchTarget) {
	if err := c.validatePutBatchTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBatchTarget",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) PutDeadLetterConfig(value *CloudwatchEventTargetDeadLetterConfig) {
	if err := c.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) PutEcsTarget(value *CloudwatchEventTargetEcsTarget) {
	if err := c.validatePutEcsTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEcsTarget",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) PutHttpTarget(value *CloudwatchEventTargetHttpTarget) {
	if err := c.validatePutHttpTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpTarget",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) PutInputTransformer(value *CloudwatchEventTargetInputTransformer) {
	if err := c.validatePutInputTransformerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInputTransformer",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) PutKinesisTarget(value *CloudwatchEventTargetKinesisTarget) {
	if err := c.validatePutKinesisTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKinesisTarget",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) PutRedshiftTarget(value *CloudwatchEventTargetRedshiftTarget) {
	if err := c.validatePutRedshiftTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRedshiftTarget",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) PutRetryPolicy(value *CloudwatchEventTargetRetryPolicy) {
	if err := c.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) PutRunCommandTargets(value interface{}) {
	if err := c.validatePutRunCommandTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRunCommandTargets",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) PutSagemakerPipelineTarget(value *CloudwatchEventTargetSagemakerPipelineTarget) {
	if err := c.validatePutSagemakerPipelineTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSagemakerPipelineTarget",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) PutSqsTarget(value *CloudwatchEventTargetSqsTarget) {
	if err := c.validatePutSqsTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSqsTarget",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetAppsyncTarget() {
	_jsii_.InvokeVoid(
		c,
		"resetAppsyncTarget",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetBatchTarget() {
	_jsii_.InvokeVoid(
		c,
		"resetBatchTarget",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetEcsTarget() {
	_jsii_.InvokeVoid(
		c,
		"resetEcsTarget",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetEventBusName() {
	_jsii_.InvokeVoid(
		c,
		"resetEventBusName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		c,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetHttpTarget() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpTarget",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetInput() {
	_jsii_.InvokeVoid(
		c,
		"resetInput",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetInputPath() {
	_jsii_.InvokeVoid(
		c,
		"resetInputPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetInputTransformer() {
	_jsii_.InvokeVoid(
		c,
		"resetInputTransformer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetKinesisTarget() {
	_jsii_.InvokeVoid(
		c,
		"resetKinesisTarget",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetRedshiftTarget() {
	_jsii_.InvokeVoid(
		c,
		"resetRedshiftTarget",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetRoleArn() {
	_jsii_.InvokeVoid(
		c,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetRunCommandTargets() {
	_jsii_.InvokeVoid(
		c,
		"resetRunCommandTargets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetSagemakerPipelineTarget() {
	_jsii_.InvokeVoid(
		c,
		"resetSagemakerPipelineTarget",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetSqsTarget() {
	_jsii_.InvokeVoid(
		c,
		"resetSqsTarget",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) ResetTargetId() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTarget) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTarget) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

