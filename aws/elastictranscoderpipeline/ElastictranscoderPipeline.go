// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastictranscoderpipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/elastictranscoderpipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/elastictranscoder_pipeline aws_elastictranscoder_pipeline}.
type ElastictranscoderPipeline interface {
	cdktf.TerraformResource
	Arn() *string
	AwsKmsKeyArn() *string
	SetAwsKmsKeyArn(val *string)
	AwsKmsKeyArnInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentConfig() ElastictranscoderPipelineContentConfigOutputReference
	ContentConfigInput() *ElastictranscoderPipelineContentConfig
	ContentConfigPermissions() ElastictranscoderPipelineContentConfigPermissionsList
	ContentConfigPermissionsInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	InputBucket() *string
	SetInputBucket(val *string)
	InputBucketInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Notifications() ElastictranscoderPipelineNotificationsOutputReference
	NotificationsInput() *ElastictranscoderPipelineNotifications
	OutputBucket() *string
	SetOutputBucket(val *string)
	OutputBucketInput() *string
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
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThumbnailConfig() ElastictranscoderPipelineThumbnailConfigOutputReference
	ThumbnailConfigInput() *ElastictranscoderPipelineThumbnailConfig
	ThumbnailConfigPermissions() ElastictranscoderPipelineThumbnailConfigPermissionsList
	ThumbnailConfigPermissionsInput() interface{}
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
	PutContentConfig(value *ElastictranscoderPipelineContentConfig)
	PutContentConfigPermissions(value interface{})
	PutNotifications(value *ElastictranscoderPipelineNotifications)
	PutThumbnailConfig(value *ElastictranscoderPipelineThumbnailConfig)
	PutThumbnailConfigPermissions(value interface{})
	ResetAwsKmsKeyArn()
	ResetContentConfig()
	ResetContentConfigPermissions()
	ResetId()
	ResetName()
	ResetNotifications()
	ResetOutputBucket()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetThumbnailConfig()
	ResetThumbnailConfigPermissions()
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

// The jsii proxy struct for ElastictranscoderPipeline
type jsiiProxy_ElastictranscoderPipeline struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElastictranscoderPipeline) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) AwsKmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsKmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) AwsKmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsKmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ContentConfig() ElastictranscoderPipelineContentConfigOutputReference {
	var returns ElastictranscoderPipelineContentConfigOutputReference
	_jsii_.Get(
		j,
		"contentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ContentConfigInput() *ElastictranscoderPipelineContentConfig {
	var returns *ElastictranscoderPipelineContentConfig
	_jsii_.Get(
		j,
		"contentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ContentConfigPermissions() ElastictranscoderPipelineContentConfigPermissionsList {
	var returns ElastictranscoderPipelineContentConfigPermissionsList
	_jsii_.Get(
		j,
		"contentConfigPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ContentConfigPermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentConfigPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) InputBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) InputBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Notifications() ElastictranscoderPipelineNotificationsOutputReference {
	var returns ElastictranscoderPipelineNotificationsOutputReference
	_jsii_.Get(
		j,
		"notifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) NotificationsInput() *ElastictranscoderPipelineNotifications {
	var returns *ElastictranscoderPipelineNotifications
	_jsii_.Get(
		j,
		"notificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) OutputBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) OutputBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ThumbnailConfig() ElastictranscoderPipelineThumbnailConfigOutputReference {
	var returns ElastictranscoderPipelineThumbnailConfigOutputReference
	_jsii_.Get(
		j,
		"thumbnailConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ThumbnailConfigInput() *ElastictranscoderPipelineThumbnailConfig {
	var returns *ElastictranscoderPipelineThumbnailConfig
	_jsii_.Get(
		j,
		"thumbnailConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ThumbnailConfigPermissions() ElastictranscoderPipelineThumbnailConfigPermissionsList {
	var returns ElastictranscoderPipelineThumbnailConfigPermissionsList
	_jsii_.Get(
		j,
		"thumbnailConfigPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ThumbnailConfigPermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thumbnailConfigPermissionsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/elastictranscoder_pipeline aws_elastictranscoder_pipeline} Resource.
func NewElastictranscoderPipeline(scope constructs.Construct, id *string, config *ElastictranscoderPipelineConfig) ElastictranscoderPipeline {
	_init_.Initialize()

	if err := validateNewElastictranscoderPipelineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastictranscoderPipeline{}

	_jsii_.Create(
		"@cdktf/provider-aws.elastictranscoderPipeline.ElastictranscoderPipeline",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/elastictranscoder_pipeline aws_elastictranscoder_pipeline} Resource.
func NewElastictranscoderPipeline_Override(e ElastictranscoderPipeline, scope constructs.Construct, id *string, config *ElastictranscoderPipelineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.elastictranscoderPipeline.ElastictranscoderPipeline",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline)SetAwsKmsKeyArn(val *string) {
	if err := j.validateSetAwsKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsKmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline)SetInputBucket(val *string) {
	if err := j.validateSetInputBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputBucket",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline)SetOutputBucket(val *string) {
	if err := j.validateSetOutputBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputBucket",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

// Generates CDKTF code for importing a ElastictranscoderPipeline resource upon running "cdktf plan <stack-name>".
func ElastictranscoderPipeline_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateElastictranscoderPipeline_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.elastictranscoderPipeline.ElastictranscoderPipeline",
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
func ElastictranscoderPipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastictranscoderPipeline_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.elastictranscoderPipeline.ElastictranscoderPipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElastictranscoderPipeline_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastictranscoderPipeline_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.elastictranscoderPipeline.ElastictranscoderPipeline",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElastictranscoderPipeline_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElastictranscoderPipeline_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.elastictranscoderPipeline.ElastictranscoderPipeline",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElastictranscoderPipeline_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.elastictranscoderPipeline.ElastictranscoderPipeline",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) PutContentConfig(value *ElastictranscoderPipelineContentConfig) {
	if err := e.validatePutContentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putContentConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) PutContentConfigPermissions(value interface{}) {
	if err := e.validatePutContentConfigPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putContentConfigPermissions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) PutNotifications(value *ElastictranscoderPipelineNotifications) {
	if err := e.validatePutNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNotifications",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) PutThumbnailConfig(value *ElastictranscoderPipelineThumbnailConfig) {
	if err := e.validatePutThumbnailConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putThumbnailConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) PutThumbnailConfigPermissions(value interface{}) {
	if err := e.validatePutThumbnailConfigPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putThumbnailConfigPermissions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetAwsKmsKeyArn() {
	_jsii_.InvokeVoid(
		e,
		"resetAwsKmsKeyArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetContentConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetContentConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetContentConfigPermissions() {
	_jsii_.InvokeVoid(
		e,
		"resetContentConfigPermissions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetNotifications() {
	_jsii_.InvokeVoid(
		e,
		"resetNotifications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetOutputBucket() {
	_jsii_.InvokeVoid(
		e,
		"resetOutputBucket",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetThumbnailConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetThumbnailConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetThumbnailConfigPermissions() {
	_jsii_.InvokeVoid(
		e,
		"resetThumbnailConfigPermissions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipeline) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

