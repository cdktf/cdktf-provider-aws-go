// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/batchjobdefinition/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/batch_job_definition aws_batch_job_definition}.
type BatchJobDefinition interface {
	cdktf.TerraformResource
	Arn() *string
	ArnPrefix() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerProperties() *string
	SetContainerProperties(val *string)
	ContainerPropertiesInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeregisterOnNewRevision() interface{}
	SetDeregisterOnNewRevision(val interface{})
	DeregisterOnNewRevisionInput() interface{}
	EcsProperties() *string
	SetEcsProperties(val *string)
	EcsPropertiesInput() *string
	EksProperties() BatchJobDefinitionEksPropertiesOutputReference
	EksPropertiesInput() *BatchJobDefinitionEksProperties
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeProperties() *string
	SetNodeProperties(val *string)
	NodePropertiesInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	PlatformCapabilities() *[]*string
	SetPlatformCapabilities(val *[]*string)
	PlatformCapabilitiesInput() *[]*string
	PropagateTags() interface{}
	SetPropagateTags(val interface{})
	PropagateTagsInput() interface{}
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
	RetryStrategy() BatchJobDefinitionRetryStrategyOutputReference
	RetryStrategyInput() *BatchJobDefinitionRetryStrategy
	Revision() *float64
	SchedulingPriority() *float64
	SetSchedulingPriority(val *float64)
	SchedulingPriorityInput() *float64
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
	Timeout() BatchJobDefinitionTimeoutOutputReference
	TimeoutInput() *BatchJobDefinitionTimeout
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
	PutEksProperties(value *BatchJobDefinitionEksProperties)
	PutRetryStrategy(value *BatchJobDefinitionRetryStrategy)
	PutTimeout(value *BatchJobDefinitionTimeout)
	ResetContainerProperties()
	ResetDeregisterOnNewRevision()
	ResetEcsProperties()
	ResetEksProperties()
	ResetId()
	ResetNodeProperties()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetPlatformCapabilities()
	ResetPropagateTags()
	ResetRegion()
	ResetRetryStrategy()
	ResetSchedulingPriority()
	ResetTags()
	ResetTagsAll()
	ResetTimeout()
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

// The jsii proxy struct for BatchJobDefinition
type jsiiProxy_BatchJobDefinition struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BatchJobDefinition) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) ArnPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) ContainerProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) ContainerPropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) DeregisterOnNewRevision() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deregisterOnNewRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) DeregisterOnNewRevisionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deregisterOnNewRevisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) EcsProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecsProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) EcsPropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecsPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) EksProperties() BatchJobDefinitionEksPropertiesOutputReference {
	var returns BatchJobDefinitionEksPropertiesOutputReference
	_jsii_.Get(
		j,
		"eksProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) EksPropertiesInput() *BatchJobDefinitionEksProperties {
	var returns *BatchJobDefinitionEksProperties
	_jsii_.Get(
		j,
		"eksPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) NodeProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) NodePropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) PlatformCapabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"platformCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) PlatformCapabilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"platformCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) PropagateTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) PropagateTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) RetryStrategy() BatchJobDefinitionRetryStrategyOutputReference {
	var returns BatchJobDefinitionRetryStrategyOutputReference
	_jsii_.Get(
		j,
		"retryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) RetryStrategyInput() *BatchJobDefinitionRetryStrategy {
	var returns *BatchJobDefinitionRetryStrategy
	_jsii_.Get(
		j,
		"retryStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Revision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) SchedulingPriority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulingPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) SchedulingPriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulingPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Timeout() BatchJobDefinitionTimeoutOutputReference {
	var returns BatchJobDefinitionTimeoutOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TimeoutInput() *BatchJobDefinitionTimeout {
	var returns *BatchJobDefinitionTimeout
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/batch_job_definition aws_batch_job_definition} Resource.
func NewBatchJobDefinition(scope constructs.Construct, id *string, config *BatchJobDefinitionConfig) BatchJobDefinition {
	_init_.Initialize()

	if err := validateNewBatchJobDefinitionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchJobDefinition{}

	_jsii_.Create(
		"@cdktf/provider-aws.batchJobDefinition.BatchJobDefinition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/batch_job_definition aws_batch_job_definition} Resource.
func NewBatchJobDefinition_Override(b BatchJobDefinition, scope constructs.Construct, id *string, config *BatchJobDefinitionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.batchJobDefinition.BatchJobDefinition",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetContainerProperties(val *string) {
	if err := j.validateSetContainerPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerProperties",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetDeregisterOnNewRevision(val interface{}) {
	if err := j.validateSetDeregisterOnNewRevisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deregisterOnNewRevision",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetEcsProperties(val *string) {
	if err := j.validateSetEcsPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecsProperties",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetNodeProperties(val *string) {
	if err := j.validateSetNodePropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeProperties",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetPlatformCapabilities(val *[]*string) {
	if err := j.validateSetPlatformCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformCapabilities",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetPropagateTags(val interface{}) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetSchedulingPriority(val *float64) {
	if err := j.validateSetSchedulingPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulingPriority",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a BatchJobDefinition resource upon running "cdktf plan <stack-name>".
func BatchJobDefinition_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBatchJobDefinition_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.batchJobDefinition.BatchJobDefinition",
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
func BatchJobDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBatchJobDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.batchJobDefinition.BatchJobDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BatchJobDefinition_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBatchJobDefinition_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.batchJobDefinition.BatchJobDefinition",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BatchJobDefinition_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBatchJobDefinition_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.batchJobDefinition.BatchJobDefinition",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BatchJobDefinition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.batchJobDefinition.BatchJobDefinition",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BatchJobDefinition) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BatchJobDefinition) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BatchJobDefinition) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BatchJobDefinition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BatchJobDefinition) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BatchJobDefinition) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BatchJobDefinition) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BatchJobDefinition) PutEksProperties(value *BatchJobDefinitionEksProperties) {
	if err := b.validatePutEksPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEksProperties",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinition) PutRetryStrategy(value *BatchJobDefinitionRetryStrategy) {
	if err := b.validatePutRetryStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRetryStrategy",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinition) PutTimeout(value *BatchJobDefinitionTimeout) {
	if err := b.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeout",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetContainerProperties() {
	_jsii_.InvokeVoid(
		b,
		"resetContainerProperties",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetDeregisterOnNewRevision() {
	_jsii_.InvokeVoid(
		b,
		"resetDeregisterOnNewRevision",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetEcsProperties() {
	_jsii_.InvokeVoid(
		b,
		"resetEcsProperties",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetEksProperties() {
	_jsii_.InvokeVoid(
		b,
		"resetEksProperties",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetNodeProperties() {
	_jsii_.InvokeVoid(
		b,
		"resetNodeProperties",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetParameters() {
	_jsii_.InvokeVoid(
		b,
		"resetParameters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetPlatformCapabilities() {
	_jsii_.InvokeVoid(
		b,
		"resetPlatformCapabilities",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		b,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetRegion() {
	_jsii_.InvokeVoid(
		b,
		"resetRegion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetRetryStrategy() {
	_jsii_.InvokeVoid(
		b,
		"resetRetryStrategy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetSchedulingPriority() {
	_jsii_.InvokeVoid(
		b,
		"resetSchedulingPriority",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetTagsAll() {
	_jsii_.InvokeVoid(
		b,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetTimeout() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeout",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

