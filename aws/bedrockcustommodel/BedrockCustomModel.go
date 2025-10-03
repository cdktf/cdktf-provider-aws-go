// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockcustommodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/bedrockcustommodel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/bedrock_custom_model aws_bedrock_custom_model}.
type BedrockCustomModel interface {
	cdktf.TerraformResource
	BaseModelIdentifier() *string
	SetBaseModelIdentifier(val *string)
	BaseModelIdentifierInput() *string
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
	CustomizationType() *string
	SetCustomizationType(val *string)
	CustomizationTypeInput() *string
	CustomModelArn() *string
	CustomModelKmsKeyId() *string
	SetCustomModelKmsKeyId(val *string)
	CustomModelKmsKeyIdInput() *string
	CustomModelName() *string
	SetCustomModelName(val *string)
	CustomModelNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hyperparameters() *map[string]*string
	SetHyperparameters(val *map[string]*string)
	HyperparametersInput() *map[string]*string
	Id() *string
	JobArn() *string
	JobName() *string
	SetJobName(val *string)
	JobNameInput() *string
	JobStatus() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OutputDataConfig() BedrockCustomModelOutputDataConfigList
	OutputDataConfigInput() interface{}
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
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktf.StringMap
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BedrockCustomModelTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrainingDataConfig() BedrockCustomModelTrainingDataConfigList
	TrainingDataConfigInput() interface{}
	TrainingMetrics() BedrockCustomModelTrainingMetricsList
	ValidationDataConfig() BedrockCustomModelValidationDataConfigList
	ValidationDataConfigInput() interface{}
	ValidationMetrics() BedrockCustomModelValidationMetricsList
	VpcConfig() BedrockCustomModelVpcConfigList
	VpcConfigInput() interface{}
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
	PutOutputDataConfig(value interface{})
	PutTimeouts(value *BedrockCustomModelTimeouts)
	PutTrainingDataConfig(value interface{})
	PutValidationDataConfig(value interface{})
	PutVpcConfig(value interface{})
	ResetCustomizationType()
	ResetCustomModelKmsKeyId()
	ResetOutputDataConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetTags()
	ResetTimeouts()
	ResetTrainingDataConfig()
	ResetValidationDataConfig()
	ResetVpcConfig()
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

// The jsii proxy struct for BedrockCustomModel
type jsiiProxy_BedrockCustomModel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BedrockCustomModel) BaseModelIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseModelIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) BaseModelIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseModelIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) CustomizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) CustomizationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customizationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) CustomModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customModelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) CustomModelKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customModelKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) CustomModelKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customModelKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) CustomModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) CustomModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customModelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) Hyperparameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hyperparameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) HyperparametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hyperparametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) JobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) JobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) JobStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) OutputDataConfig() BedrockCustomModelOutputDataConfigList {
	var returns BedrockCustomModelOutputDataConfigList
	_jsii_.Get(
		j,
		"outputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) OutputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) TagsAll() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) Timeouts() BedrockCustomModelTimeoutsOutputReference {
	var returns BedrockCustomModelTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) TrainingDataConfig() BedrockCustomModelTrainingDataConfigList {
	var returns BedrockCustomModelTrainingDataConfigList
	_jsii_.Get(
		j,
		"trainingDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) TrainingDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trainingDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) TrainingMetrics() BedrockCustomModelTrainingMetricsList {
	var returns BedrockCustomModelTrainingMetricsList
	_jsii_.Get(
		j,
		"trainingMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) ValidationDataConfig() BedrockCustomModelValidationDataConfigList {
	var returns BedrockCustomModelValidationDataConfigList
	_jsii_.Get(
		j,
		"validationDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) ValidationDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) ValidationMetrics() BedrockCustomModelValidationMetricsList {
	var returns BedrockCustomModelValidationMetricsList
	_jsii_.Get(
		j,
		"validationMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) VpcConfig() BedrockCustomModelVpcConfigList {
	var returns BedrockCustomModelVpcConfigList
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCustomModel) VpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/bedrock_custom_model aws_bedrock_custom_model} Resource.
func NewBedrockCustomModel(scope constructs.Construct, id *string, config *BedrockCustomModelConfig) BedrockCustomModel {
	_init_.Initialize()

	if err := validateNewBedrockCustomModelParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockCustomModel{}

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockCustomModel.BedrockCustomModel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/bedrock_custom_model aws_bedrock_custom_model} Resource.
func NewBedrockCustomModel_Override(b BedrockCustomModel, scope constructs.Construct, id *string, config *BedrockCustomModelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockCustomModel.BedrockCustomModel",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetBaseModelIdentifier(val *string) {
	if err := j.validateSetBaseModelIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseModelIdentifier",
		val,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetCustomizationType(val *string) {
	if err := j.validateSetCustomizationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customizationType",
		val,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetCustomModelKmsKeyId(val *string) {
	if err := j.validateSetCustomModelKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customModelKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetCustomModelName(val *string) {
	if err := j.validateSetCustomModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customModelName",
		val,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetHyperparameters(val *map[string]*string) {
	if err := j.validateSetHyperparametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hyperparameters",
		val,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetJobName(val *string) {
	if err := j.validateSetJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobName",
		val,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_BedrockCustomModel)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a BedrockCustomModel resource upon running "cdktf plan <stack-name>".
func BedrockCustomModel_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBedrockCustomModel_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockCustomModel.BedrockCustomModel",
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
func BedrockCustomModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockCustomModel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockCustomModel.BedrockCustomModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockCustomModel_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockCustomModel_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockCustomModel.BedrockCustomModel",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockCustomModel_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockCustomModel_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockCustomModel.BedrockCustomModel",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BedrockCustomModel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.bedrockCustomModel.BedrockCustomModel",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BedrockCustomModel) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BedrockCustomModel) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BedrockCustomModel) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockCustomModel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BedrockCustomModel) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockCustomModel) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockCustomModel) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockCustomModel) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockCustomModel) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockCustomModel) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockCustomModel) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockCustomModel) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCustomModel) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BedrockCustomModel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BedrockCustomModel) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockCustomModel) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BedrockCustomModel) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockCustomModel) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BedrockCustomModel) PutOutputDataConfig(value interface{}) {
	if err := b.validatePutOutputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putOutputDataConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockCustomModel) PutTimeouts(value *BedrockCustomModelTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockCustomModel) PutTrainingDataConfig(value interface{}) {
	if err := b.validatePutTrainingDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTrainingDataConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockCustomModel) PutValidationDataConfig(value interface{}) {
	if err := b.validatePutValidationDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putValidationDataConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockCustomModel) PutVpcConfig(value interface{}) {
	if err := b.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockCustomModel) ResetCustomizationType() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomizationType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockCustomModel) ResetCustomModelKmsKeyId() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomModelKmsKeyId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockCustomModel) ResetOutputDataConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetOutputDataConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockCustomModel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockCustomModel) ResetRegion() {
	_jsii_.InvokeVoid(
		b,
		"resetRegion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockCustomModel) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockCustomModel) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockCustomModel) ResetTrainingDataConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetTrainingDataConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockCustomModel) ResetValidationDataConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetValidationDataConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockCustomModel) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockCustomModel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCustomModel) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCustomModel) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCustomModel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCustomModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCustomModel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

