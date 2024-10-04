// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/syntheticscanary/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/synthetics_canary aws_synthetics_canary}.
type SyntheticsCanary interface {
	cdktf.TerraformResource
	Arn() *string
	ArtifactConfig() SyntheticsCanaryArtifactConfigOutputReference
	ArtifactConfigInput() *SyntheticsCanaryArtifactConfig
	ArtifactS3Location() *string
	SetArtifactS3Location(val *string)
	ArtifactS3LocationInput() *string
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
	DeleteLambda() interface{}
	SetDeleteLambda(val interface{})
	DeleteLambdaInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EngineArn() *string
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	ExecutionRoleArnInput() *string
	FailureRetentionPeriod() *float64
	SetFailureRetentionPeriod(val *float64)
	FailureRetentionPeriodInput() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Handler() *string
	SetHandler(val *string)
	HandlerInput() *string
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
	RunConfig() SyntheticsCanaryRunConfigOutputReference
	RunConfigInput() *SyntheticsCanaryRunConfig
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	RuntimeVersionInput() *string
	S3Bucket() *string
	SetS3Bucket(val *string)
	S3BucketInput() *string
	S3Key() *string
	SetS3Key(val *string)
	S3KeyInput() *string
	S3Version() *string
	SetS3Version(val *string)
	S3VersionInput() *string
	Schedule() SyntheticsCanaryScheduleOutputReference
	ScheduleInput() *SyntheticsCanarySchedule
	SourceLocationArn() *string
	StartCanary() interface{}
	SetStartCanary(val interface{})
	StartCanaryInput() interface{}
	Status() *string
	SuccessRetentionPeriod() *float64
	SetSuccessRetentionPeriod(val *float64)
	SuccessRetentionPeriodInput() *float64
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
	Timeline() SyntheticsCanaryTimelineList
	VpcConfig() SyntheticsCanaryVpcConfigOutputReference
	VpcConfigInput() *SyntheticsCanaryVpcConfig
	ZipFile() *string
	SetZipFile(val *string)
	ZipFileInput() *string
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
	PutArtifactConfig(value *SyntheticsCanaryArtifactConfig)
	PutRunConfig(value *SyntheticsCanaryRunConfig)
	PutSchedule(value *SyntheticsCanarySchedule)
	PutVpcConfig(value *SyntheticsCanaryVpcConfig)
	ResetArtifactConfig()
	ResetDeleteLambda()
	ResetFailureRetentionPeriod()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRunConfig()
	ResetS3Bucket()
	ResetS3Key()
	ResetS3Version()
	ResetStartCanary()
	ResetSuccessRetentionPeriod()
	ResetTags()
	ResetTagsAll()
	ResetVpcConfig()
	ResetZipFile()
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

// The jsii proxy struct for SyntheticsCanary
type jsiiProxy_SyntheticsCanary struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SyntheticsCanary) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ArtifactConfig() SyntheticsCanaryArtifactConfigOutputReference {
	var returns SyntheticsCanaryArtifactConfigOutputReference
	_jsii_.Get(
		j,
		"artifactConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ArtifactConfigInput() *SyntheticsCanaryArtifactConfig {
	var returns *SyntheticsCanaryArtifactConfig
	_jsii_.Get(
		j,
		"artifactConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ArtifactS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ArtifactS3LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactS3LocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) DeleteLambda() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteLambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) DeleteLambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteLambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) EngineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) FailureRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) FailureRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) HandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) RunConfig() SyntheticsCanaryRunConfigOutputReference {
	var returns SyntheticsCanaryRunConfigOutputReference
	_jsii_.Get(
		j,
		"runConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) RunConfigInput() *SyntheticsCanaryRunConfig {
	var returns *SyntheticsCanaryRunConfig
	_jsii_.Get(
		j,
		"runConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) S3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) S3Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) S3KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) S3Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) S3VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3VersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Schedule() SyntheticsCanaryScheduleOutputReference {
	var returns SyntheticsCanaryScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ScheduleInput() *SyntheticsCanarySchedule {
	var returns *SyntheticsCanarySchedule
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) SourceLocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceLocationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) StartCanary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startCanary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) StartCanaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startCanaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) SuccessRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) SuccessRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Timeline() SyntheticsCanaryTimelineList {
	var returns SyntheticsCanaryTimelineList
	_jsii_.Get(
		j,
		"timeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) VpcConfig() SyntheticsCanaryVpcConfigOutputReference {
	var returns SyntheticsCanaryVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) VpcConfigInput() *SyntheticsCanaryVpcConfig {
	var returns *SyntheticsCanaryVpcConfig
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ZipFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ZipFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipFileInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/synthetics_canary aws_synthetics_canary} Resource.
func NewSyntheticsCanary(scope constructs.Construct, id *string, config *SyntheticsCanaryConfig) SyntheticsCanary {
	_init_.Initialize()

	if err := validateNewSyntheticsCanaryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsCanary{}

	_jsii_.Create(
		"@cdktf/provider-aws.syntheticsCanary.SyntheticsCanary",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/synthetics_canary aws_synthetics_canary} Resource.
func NewSyntheticsCanary_Override(s SyntheticsCanary, scope constructs.Construct, id *string, config *SyntheticsCanaryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.syntheticsCanary.SyntheticsCanary",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetArtifactS3Location(val *string) {
	if err := j.validateSetArtifactS3LocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"artifactS3Location",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetDeleteLambda(val interface{}) {
	if err := j.validateSetDeleteLambdaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteLambda",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetFailureRetentionPeriod(val *float64) {
	if err := j.validateSetFailureRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetHandler(val *string) {
	if err := j.validateSetHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetS3Bucket(val *string) {
	if err := j.validateSetS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Bucket",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetS3Key(val *string) {
	if err := j.validateSetS3KeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Key",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetS3Version(val *string) {
	if err := j.validateSetS3VersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Version",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetStartCanary(val interface{}) {
	if err := j.validateSetStartCanaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startCanary",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetSuccessRetentionPeriod(val *float64) {
	if err := j.validateSetSuccessRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetZipFile(val *string) {
	if err := j.validateSetZipFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zipFile",
		val,
	)
}

// Generates CDKTF code for importing a SyntheticsCanary resource upon running "cdktf plan <stack-name>".
func SyntheticsCanary_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSyntheticsCanary_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.syntheticsCanary.SyntheticsCanary",
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
func SyntheticsCanary_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsCanary_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.syntheticsCanary.SyntheticsCanary",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsCanary_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsCanary_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.syntheticsCanary.SyntheticsCanary",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsCanary_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsCanary_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.syntheticsCanary.SyntheticsCanary",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SyntheticsCanary_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.syntheticsCanary.SyntheticsCanary",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SyntheticsCanary) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SyntheticsCanary) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SyntheticsCanary) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsCanary) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SyntheticsCanary) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsCanary) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutArtifactConfig(value *SyntheticsCanaryArtifactConfig) {
	if err := s.validatePutArtifactConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putArtifactConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutRunConfig(value *SyntheticsCanaryRunConfig) {
	if err := s.validatePutRunConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRunConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutSchedule(value *SyntheticsCanarySchedule) {
	if err := s.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSchedule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutVpcConfig(value *SyntheticsCanaryVpcConfig) {
	if err := s.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetArtifactConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetArtifactConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetDeleteLambda() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteLambda",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetFailureRetentionPeriod() {
	_jsii_.InvokeVoid(
		s,
		"resetFailureRetentionPeriod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetRunConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetRunConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetS3Bucket() {
	_jsii_.InvokeVoid(
		s,
		"resetS3Bucket",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetS3Key() {
	_jsii_.InvokeVoid(
		s,
		"resetS3Key",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetS3Version() {
	_jsii_.InvokeVoid(
		s,
		"resetS3Version",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetStartCanary() {
	_jsii_.InvokeVoid(
		s,
		"resetStartCanary",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetSuccessRetentionPeriod() {
	_jsii_.InvokeVoid(
		s,
		"resetSuccessRetentionPeriod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetZipFile() {
	_jsii_.InvokeVoid(
		s,
		"resetZipFile",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

