// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdafunction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/lambdafunction/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/lambda_function aws_lambda_function}.
type LambdaFunction interface {
	cdktf.TerraformResource
	Architectures() *[]*string
	SetArchitectures(val *[]*string)
	ArchitecturesInput() *[]*string
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CodeSha256() *string
	CodeSigningConfigArn() *string
	SetCodeSigningConfigArn(val *string)
	CodeSigningConfigArnInput() *string
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
	DeadLetterConfig() LambdaFunctionDeadLetterConfigOutputReference
	DeadLetterConfigInput() *LambdaFunctionDeadLetterConfig
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Environment() LambdaFunctionEnvironmentOutputReference
	EnvironmentInput() *LambdaFunctionEnvironment
	EphemeralStorage() LambdaFunctionEphemeralStorageOutputReference
	EphemeralStorageInput() *LambdaFunctionEphemeralStorage
	Filename() *string
	SetFilename(val *string)
	FilenameInput() *string
	FileSystemConfig() LambdaFunctionFileSystemConfigOutputReference
	FileSystemConfigInput() *LambdaFunctionFileSystemConfig
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionNameInput() *string
	Handler() *string
	SetHandler(val *string)
	HandlerInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageConfig() LambdaFunctionImageConfigOutputReference
	ImageConfigInput() *LambdaFunctionImageConfig
	ImageUri() *string
	SetImageUri(val *string)
	ImageUriInput() *string
	InvokeArn() *string
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	LastModified() *string
	Layers() *[]*string
	SetLayers(val *[]*string)
	LayersInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoggingConfig() LambdaFunctionLoggingConfigOutputReference
	LoggingConfigInput() *LambdaFunctionLoggingConfig
	MemorySize() *float64
	SetMemorySize(val *float64)
	MemorySizeInput() *float64
	// The tree node.
	Node() constructs.Node
	PackageType() *string
	SetPackageType(val *string)
	PackageTypeInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Publish() interface{}
	SetPublish(val interface{})
	PublishInput() interface{}
	QualifiedArn() *string
	QualifiedInvokeArn() *string
	// Experimental.
	RawOverrides() interface{}
	ReplacementSecurityGroupIds() *[]*string
	SetReplacementSecurityGroupIds(val *[]*string)
	ReplacementSecurityGroupIdsInput() *[]*string
	ReplaceSecurityGroupsOnDestroy() interface{}
	SetReplaceSecurityGroupsOnDestroy(val interface{})
	ReplaceSecurityGroupsOnDestroyInput() interface{}
	ReservedConcurrentExecutions() *float64
	SetReservedConcurrentExecutions(val *float64)
	ReservedConcurrentExecutionsInput() *float64
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	S3Bucket() *string
	SetS3Bucket(val *string)
	S3BucketInput() *string
	S3Key() *string
	SetS3Key(val *string)
	S3KeyInput() *string
	S3ObjectVersion() *string
	SetS3ObjectVersion(val *string)
	S3ObjectVersionInput() *string
	SigningJobArn() *string
	SigningProfileVersionArn() *string
	SkipDestroy() interface{}
	SetSkipDestroy(val interface{})
	SkipDestroyInput() interface{}
	SnapStart() LambdaFunctionSnapStartOutputReference
	SnapStartInput() *LambdaFunctionSnapStart
	SourceCodeHash() *string
	SetSourceCodeHash(val *string)
	SourceCodeHashInput() *string
	SourceCodeSize() *float64
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
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	Timeouts() LambdaFunctionTimeoutsOutputReference
	TimeoutsInput() interface{}
	TracingConfig() LambdaFunctionTracingConfigOutputReference
	TracingConfigInput() *LambdaFunctionTracingConfig
	Version() *string
	VpcConfig() LambdaFunctionVpcConfigOutputReference
	VpcConfigInput() *LambdaFunctionVpcConfig
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
	PutDeadLetterConfig(value *LambdaFunctionDeadLetterConfig)
	PutEnvironment(value *LambdaFunctionEnvironment)
	PutEphemeralStorage(value *LambdaFunctionEphemeralStorage)
	PutFileSystemConfig(value *LambdaFunctionFileSystemConfig)
	PutImageConfig(value *LambdaFunctionImageConfig)
	PutLoggingConfig(value *LambdaFunctionLoggingConfig)
	PutSnapStart(value *LambdaFunctionSnapStart)
	PutTimeouts(value *LambdaFunctionTimeouts)
	PutTracingConfig(value *LambdaFunctionTracingConfig)
	PutVpcConfig(value *LambdaFunctionVpcConfig)
	ResetArchitectures()
	ResetCodeSigningConfigArn()
	ResetDeadLetterConfig()
	ResetDescription()
	ResetEnvironment()
	ResetEphemeralStorage()
	ResetFilename()
	ResetFileSystemConfig()
	ResetHandler()
	ResetId()
	ResetImageConfig()
	ResetImageUri()
	ResetKmsKeyArn()
	ResetLayers()
	ResetLoggingConfig()
	ResetMemorySize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPackageType()
	ResetPublish()
	ResetReplacementSecurityGroupIds()
	ResetReplaceSecurityGroupsOnDestroy()
	ResetReservedConcurrentExecutions()
	ResetRuntime()
	ResetS3Bucket()
	ResetS3Key()
	ResetS3ObjectVersion()
	ResetSkipDestroy()
	ResetSnapStart()
	ResetSourceCodeHash()
	ResetTags()
	ResetTagsAll()
	ResetTimeout()
	ResetTimeouts()
	ResetTracingConfig()
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

// The jsii proxy struct for LambdaFunction
type jsiiProxy_LambdaFunction struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LambdaFunction) Architectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ArchitecturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architecturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CodeSha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CodeSigningConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CodeSigningConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DeadLetterConfig() LambdaFunctionDeadLetterConfigOutputReference {
	var returns LambdaFunctionDeadLetterConfigOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DeadLetterConfigInput() *LambdaFunctionDeadLetterConfig {
	var returns *LambdaFunctionDeadLetterConfig
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Environment() LambdaFunctionEnvironmentOutputReference {
	var returns LambdaFunctionEnvironmentOutputReference
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) EnvironmentInput() *LambdaFunctionEnvironment {
	var returns *LambdaFunctionEnvironment
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) EphemeralStorage() LambdaFunctionEphemeralStorageOutputReference {
	var returns LambdaFunctionEphemeralStorageOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) EphemeralStorageInput() *LambdaFunctionEphemeralStorage {
	var returns *LambdaFunctionEphemeralStorage
	_jsii_.Get(
		j,
		"ephemeralStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Filename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FileSystemConfig() LambdaFunctionFileSystemConfigOutputReference {
	var returns LambdaFunctionFileSystemConfigOutputReference
	_jsii_.Get(
		j,
		"fileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FileSystemConfigInput() *LambdaFunctionFileSystemConfig {
	var returns *LambdaFunctionFileSystemConfig
	_jsii_.Get(
		j,
		"fileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) HandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ImageConfig() LambdaFunctionImageConfigOutputReference {
	var returns LambdaFunctionImageConfigOutputReference
	_jsii_.Get(
		j,
		"imageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ImageConfigInput() *LambdaFunctionImageConfig {
	var returns *LambdaFunctionImageConfig
	_jsii_.Get(
		j,
		"imageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ImageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) InvokeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Layers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) LayersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) LoggingConfig() LambdaFunctionLoggingConfigOutputReference {
	var returns LambdaFunctionLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) LoggingConfigInput() *LambdaFunctionLoggingConfig {
	var returns *LambdaFunctionLoggingConfig
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) MemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) MemorySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) PackageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) PackageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Publish() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) PublishInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) QualifiedArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifiedArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) QualifiedInvokeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifiedInvokeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ReplacementSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replacementSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ReplacementSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replacementSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ReplaceSecurityGroupsOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceSecurityGroupsOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ReplaceSecurityGroupsOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceSecurityGroupsOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ReservedConcurrentExecutions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ReservedConcurrentExecutionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) S3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) S3Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) S3KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) S3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) S3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SigningJobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingJobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SigningProfileVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingProfileVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SkipDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SkipDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SnapStart() LambdaFunctionSnapStartOutputReference {
	var returns LambdaFunctionSnapStartOutputReference
	_jsii_.Get(
		j,
		"snapStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SnapStartInput() *LambdaFunctionSnapStart {
	var returns *LambdaFunctionSnapStart
	_jsii_.Get(
		j,
		"snapStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SourceCodeHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodeHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SourceCodeHashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodeHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SourceCodeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceCodeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Timeouts() LambdaFunctionTimeoutsOutputReference {
	var returns LambdaFunctionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TracingConfig() LambdaFunctionTracingConfigOutputReference {
	var returns LambdaFunctionTracingConfigOutputReference
	_jsii_.Get(
		j,
		"tracingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TracingConfigInput() *LambdaFunctionTracingConfig {
	var returns *LambdaFunctionTracingConfig
	_jsii_.Get(
		j,
		"tracingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) VpcConfig() LambdaFunctionVpcConfigOutputReference {
	var returns LambdaFunctionVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) VpcConfigInput() *LambdaFunctionVpcConfig {
	var returns *LambdaFunctionVpcConfig
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/lambda_function aws_lambda_function} Resource.
func NewLambdaFunction(scope constructs.Construct, id *string, config *LambdaFunctionConfig) LambdaFunction {
	_init_.Initialize()

	if err := validateNewLambdaFunctionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaFunction{}

	_jsii_.Create(
		"@cdktf/provider-aws.lambdaFunction.LambdaFunction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/lambda_function aws_lambda_function} Resource.
func NewLambdaFunction_Override(l LambdaFunction, scope constructs.Construct, id *string, config *LambdaFunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lambdaFunction.LambdaFunction",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LambdaFunction)SetArchitectures(val *[]*string) {
	if err := j.validateSetArchitecturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architectures",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetCodeSigningConfigArn(val *string) {
	if err := j.validateSetCodeSigningConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeSigningConfigArn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetFilename(val *string) {
	if err := j.validateSetFilenameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filename",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetFunctionName(val *string) {
	if err := j.validateSetFunctionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetHandler(val *string) {
	if err := j.validateSetHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetImageUri(val *string) {
	if err := j.validateSetImageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetLayers(val *[]*string) {
	if err := j.validateSetLayersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"layers",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetMemorySize(val *float64) {
	if err := j.validateSetMemorySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySize",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetPackageType(val *string) {
	if err := j.validateSetPackageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageType",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetPublish(val interface{}) {
	if err := j.validateSetPublishParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publish",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetReplacementSecurityGroupIds(val *[]*string) {
	if err := j.validateSetReplacementSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replacementSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetReplaceSecurityGroupsOnDestroy(val interface{}) {
	if err := j.validateSetReplaceSecurityGroupsOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceSecurityGroupsOnDestroy",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetReservedConcurrentExecutions(val *float64) {
	if err := j.validateSetReservedConcurrentExecutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedConcurrentExecutions",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetS3Bucket(val *string) {
	if err := j.validateSetS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Bucket",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetS3Key(val *string) {
	if err := j.validateSetS3KeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Key",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetS3ObjectVersion(val *string) {
	if err := j.validateSetS3ObjectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetSkipDestroy(val interface{}) {
	if err := j.validateSetSkipDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipDestroy",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetSourceCodeHash(val *string) {
	if err := j.validateSetSourceCodeHashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCodeHash",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

// Generates CDKTF code for importing a LambdaFunction resource upon running "cdktf plan <stack-name>".
func LambdaFunction_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLambdaFunction_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lambdaFunction.LambdaFunction",
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
func LambdaFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lambdaFunction.LambdaFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LambdaFunction_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaFunction_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lambdaFunction.LambdaFunction",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LambdaFunction_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaFunction_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lambdaFunction.LambdaFunction",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaFunction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.lambdaFunction.LambdaFunction",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LambdaFunction) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LambdaFunction) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LambdaFunction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LambdaFunction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LambdaFunction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LambdaFunction) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LambdaFunction) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LambdaFunction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LambdaFunction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LambdaFunction) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LambdaFunction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LambdaFunction) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LambdaFunction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LambdaFunction) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LambdaFunction) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LambdaFunction) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LambdaFunction) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LambdaFunction) PutDeadLetterConfig(value *LambdaFunctionDeadLetterConfig) {
	if err := l.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutEnvironment(value *LambdaFunctionEnvironment) {
	if err := l.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutEphemeralStorage(value *LambdaFunctionEphemeralStorage) {
	if err := l.validatePutEphemeralStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putEphemeralStorage",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutFileSystemConfig(value *LambdaFunctionFileSystemConfig) {
	if err := l.validatePutFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFileSystemConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutImageConfig(value *LambdaFunctionImageConfig) {
	if err := l.validatePutImageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putImageConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutLoggingConfig(value *LambdaFunctionLoggingConfig) {
	if err := l.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutSnapStart(value *LambdaFunctionSnapStart) {
	if err := l.validatePutSnapStartParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSnapStart",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutTimeouts(value *LambdaFunctionTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutTracingConfig(value *LambdaFunctionTracingConfig) {
	if err := l.validatePutTracingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTracingConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutVpcConfig(value *LambdaFunctionVpcConfig) {
	if err := l.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) ResetArchitectures() {
	_jsii_.InvokeVoid(
		l,
		"resetArchitectures",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetCodeSigningConfigArn() {
	_jsii_.InvokeVoid(
		l,
		"resetCodeSigningConfigArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetEnvironment() {
	_jsii_.InvokeVoid(
		l,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetEphemeralStorage() {
	_jsii_.InvokeVoid(
		l,
		"resetEphemeralStorage",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetFilename() {
	_jsii_.InvokeVoid(
		l,
		"resetFilename",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetFileSystemConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetFileSystemConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetHandler() {
	_jsii_.InvokeVoid(
		l,
		"resetHandler",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetImageConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetImageConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetImageUri() {
	_jsii_.InvokeVoid(
		l,
		"resetImageUri",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		l,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetLayers() {
	_jsii_.InvokeVoid(
		l,
		"resetLayers",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetMemorySize() {
	_jsii_.InvokeVoid(
		l,
		"resetMemorySize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetPackageType() {
	_jsii_.InvokeVoid(
		l,
		"resetPackageType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetPublish() {
	_jsii_.InvokeVoid(
		l,
		"resetPublish",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetReplacementSecurityGroupIds() {
	_jsii_.InvokeVoid(
		l,
		"resetReplacementSecurityGroupIds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetReplaceSecurityGroupsOnDestroy() {
	_jsii_.InvokeVoid(
		l,
		"resetReplaceSecurityGroupsOnDestroy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetReservedConcurrentExecutions() {
	_jsii_.InvokeVoid(
		l,
		"resetReservedConcurrentExecutions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetRuntime() {
	_jsii_.InvokeVoid(
		l,
		"resetRuntime",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetS3Bucket() {
	_jsii_.InvokeVoid(
		l,
		"resetS3Bucket",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetS3Key() {
	_jsii_.InvokeVoid(
		l,
		"resetS3Key",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetS3ObjectVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetS3ObjectVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetSkipDestroy() {
	_jsii_.InvokeVoid(
		l,
		"resetSkipDestroy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetSnapStart() {
	_jsii_.InvokeVoid(
		l,
		"resetSnapStart",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetSourceCodeHash() {
	_jsii_.InvokeVoid(
		l,
		"resetSourceCodeHash",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTagsAll() {
	_jsii_.InvokeVoid(
		l,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTracingConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetTracingConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

