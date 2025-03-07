// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/codebuildproject/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/codebuild_project aws_codebuild_project}.
type CodebuildProject interface {
	cdktf.TerraformResource
	Arn() *string
	Artifacts() CodebuildProjectArtifactsOutputReference
	ArtifactsInput() *CodebuildProjectArtifacts
	BadgeEnabled() interface{}
	SetBadgeEnabled(val interface{})
	BadgeEnabledInput() interface{}
	BadgeUrl() *string
	BuildBatchConfig() CodebuildProjectBuildBatchConfigOutputReference
	BuildBatchConfigInput() *CodebuildProjectBuildBatchConfig
	BuildTimeout() *float64
	SetBuildTimeout(val *float64)
	BuildTimeoutInput() *float64
	Cache() CodebuildProjectCacheOutputReference
	CacheInput() *CodebuildProjectCache
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConcurrentBuildLimit() *float64
	SetConcurrentBuildLimit(val *float64)
	ConcurrentBuildLimitInput() *float64
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EncryptionKey() *string
	SetEncryptionKey(val *string)
	EncryptionKeyInput() *string
	Environment() CodebuildProjectEnvironmentOutputReference
	EnvironmentInput() *CodebuildProjectEnvironment
	FileSystemLocations() CodebuildProjectFileSystemLocationsList
	FileSystemLocationsInput() interface{}
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
	LogsConfig() CodebuildProjectLogsConfigOutputReference
	LogsConfigInput() *CodebuildProjectLogsConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProjectVisibility() *string
	SetProjectVisibility(val *string)
	ProjectVisibilityInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicProjectAlias() *string
	QueuedTimeout() *float64
	SetQueuedTimeout(val *float64)
	QueuedTimeoutInput() *float64
	// Experimental.
	RawOverrides() interface{}
	ResourceAccessRole() *string
	SetResourceAccessRole(val *string)
	ResourceAccessRoleInput() *string
	SecondaryArtifacts() CodebuildProjectSecondaryArtifactsList
	SecondaryArtifactsInput() interface{}
	SecondarySources() CodebuildProjectSecondarySourcesList
	SecondarySourcesInput() interface{}
	SecondarySourceVersion() CodebuildProjectSecondarySourceVersionList
	SecondarySourceVersionInput() interface{}
	ServiceRole() *string
	SetServiceRole(val *string)
	ServiceRoleInput() *string
	Source() CodebuildProjectSourceOutputReference
	SourceInput() *CodebuildProjectSource
	SourceVersion() *string
	SetSourceVersion(val *string)
	SourceVersionInput() *string
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
	VpcConfig() CodebuildProjectVpcConfigOutputReference
	VpcConfigInput() *CodebuildProjectVpcConfig
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
	PutArtifacts(value *CodebuildProjectArtifacts)
	PutBuildBatchConfig(value *CodebuildProjectBuildBatchConfig)
	PutCache(value *CodebuildProjectCache)
	PutEnvironment(value *CodebuildProjectEnvironment)
	PutFileSystemLocations(value interface{})
	PutLogsConfig(value *CodebuildProjectLogsConfig)
	PutSecondaryArtifacts(value interface{})
	PutSecondarySources(value interface{})
	PutSecondarySourceVersion(value interface{})
	PutSource(value *CodebuildProjectSource)
	PutVpcConfig(value *CodebuildProjectVpcConfig)
	ResetBadgeEnabled()
	ResetBuildBatchConfig()
	ResetBuildTimeout()
	ResetCache()
	ResetConcurrentBuildLimit()
	ResetDescription()
	ResetEncryptionKey()
	ResetFileSystemLocations()
	ResetId()
	ResetLogsConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjectVisibility()
	ResetQueuedTimeout()
	ResetResourceAccessRole()
	ResetSecondaryArtifacts()
	ResetSecondarySources()
	ResetSecondarySourceVersion()
	ResetSourceVersion()
	ResetTags()
	ResetTagsAll()
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

// The jsii proxy struct for CodebuildProject
type jsiiProxy_CodebuildProject struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CodebuildProject) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Artifacts() CodebuildProjectArtifactsOutputReference {
	var returns CodebuildProjectArtifactsOutputReference
	_jsii_.Get(
		j,
		"artifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ArtifactsInput() *CodebuildProjectArtifacts {
	var returns *CodebuildProjectArtifacts
	_jsii_.Get(
		j,
		"artifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) BadgeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"badgeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) BadgeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"badgeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) BadgeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"badgeUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) BuildBatchConfig() CodebuildProjectBuildBatchConfigOutputReference {
	var returns CodebuildProjectBuildBatchConfigOutputReference
	_jsii_.Get(
		j,
		"buildBatchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) BuildBatchConfigInput() *CodebuildProjectBuildBatchConfig {
	var returns *CodebuildProjectBuildBatchConfig
	_jsii_.Get(
		j,
		"buildBatchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) BuildTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"buildTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) BuildTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"buildTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Cache() CodebuildProjectCacheOutputReference {
	var returns CodebuildProjectCacheOutputReference
	_jsii_.Get(
		j,
		"cache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) CacheInput() *CodebuildProjectCache {
	var returns *CodebuildProjectCache
	_jsii_.Get(
		j,
		"cacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ConcurrentBuildLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"concurrentBuildLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ConcurrentBuildLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"concurrentBuildLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) EncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) EncryptionKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Environment() CodebuildProjectEnvironmentOutputReference {
	var returns CodebuildProjectEnvironmentOutputReference
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) EnvironmentInput() *CodebuildProjectEnvironment {
	var returns *CodebuildProjectEnvironment
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) FileSystemLocations() CodebuildProjectFileSystemLocationsList {
	var returns CodebuildProjectFileSystemLocationsList
	_jsii_.Get(
		j,
		"fileSystemLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) FileSystemLocationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileSystemLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) LogsConfig() CodebuildProjectLogsConfigOutputReference {
	var returns CodebuildProjectLogsConfigOutputReference
	_jsii_.Get(
		j,
		"logsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) LogsConfigInput() *CodebuildProjectLogsConfig {
	var returns *CodebuildProjectLogsConfig
	_jsii_.Get(
		j,
		"logsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ProjectVisibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ProjectVisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) PublicProjectAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicProjectAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) QueuedTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queuedTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) QueuedTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queuedTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ResourceAccessRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceAccessRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ResourceAccessRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceAccessRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SecondaryArtifacts() CodebuildProjectSecondaryArtifactsList {
	var returns CodebuildProjectSecondaryArtifactsList
	_jsii_.Get(
		j,
		"secondaryArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SecondaryArtifactsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SecondarySources() CodebuildProjectSecondarySourcesList {
	var returns CodebuildProjectSecondarySourcesList
	_jsii_.Get(
		j,
		"secondarySources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SecondarySourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondarySourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SecondarySourceVersion() CodebuildProjectSecondarySourceVersionList {
	var returns CodebuildProjectSecondarySourceVersionList
	_jsii_.Get(
		j,
		"secondarySourceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SecondarySourceVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondarySourceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Source() CodebuildProjectSourceOutputReference {
	var returns CodebuildProjectSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SourceInput() *CodebuildProjectSource {
	var returns *CodebuildProjectSource
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SourceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SourceVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) VpcConfig() CodebuildProjectVpcConfigOutputReference {
	var returns CodebuildProjectVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) VpcConfigInput() *CodebuildProjectVpcConfig {
	var returns *CodebuildProjectVpcConfig
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/codebuild_project aws_codebuild_project} Resource.
func NewCodebuildProject(scope constructs.Construct, id *string, config *CodebuildProjectConfig) CodebuildProject {
	_init_.Initialize()

	if err := validateNewCodebuildProjectParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodebuildProject{}

	_jsii_.Create(
		"@cdktf/provider-aws.codebuildProject.CodebuildProject",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/codebuild_project aws_codebuild_project} Resource.
func NewCodebuildProject_Override(c CodebuildProject, scope constructs.Construct, id *string, config *CodebuildProjectConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.codebuildProject.CodebuildProject",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CodebuildProject)SetBadgeEnabled(val interface{}) {
	if err := j.validateSetBadgeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"badgeEnabled",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetBuildTimeout(val *float64) {
	if err := j.validateSetBuildTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildTimeout",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetConcurrentBuildLimit(val *float64) {
	if err := j.validateSetConcurrentBuildLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"concurrentBuildLimit",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetEncryptionKey(val *string) {
	if err := j.validateSetEncryptionKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKey",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetProjectVisibility(val *string) {
	if err := j.validateSetProjectVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectVisibility",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetQueuedTimeout(val *float64) {
	if err := j.validateSetQueuedTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queuedTimeout",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetResourceAccessRole(val *string) {
	if err := j.validateSetResourceAccessRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceAccessRole",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetServiceRole(val *string) {
	if err := j.validateSetServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetSourceVersion(val *string) {
	if err := j.validateSetSourceVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceVersion",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a CodebuildProject resource upon running "cdktf plan <stack-name>".
func CodebuildProject_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCodebuildProject_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.codebuildProject.CodebuildProject",
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
func CodebuildProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCodebuildProject_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.codebuildProject.CodebuildProject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CodebuildProject_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCodebuildProject_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.codebuildProject.CodebuildProject",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CodebuildProject_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCodebuildProject_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.codebuildProject.CodebuildProject",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CodebuildProject_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.codebuildProject.CodebuildProject",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CodebuildProject) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CodebuildProject) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CodebuildProject) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CodebuildProject) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodebuildProject) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CodebuildProject) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CodebuildProject) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CodebuildProject) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CodebuildProject) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CodebuildProject) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CodebuildProject) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CodebuildProject) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProject) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CodebuildProject) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodebuildProject) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CodebuildProject) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CodebuildProject) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CodebuildProject) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CodebuildProject) PutArtifacts(value *CodebuildProjectArtifacts) {
	if err := c.validatePutArtifactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putArtifacts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutBuildBatchConfig(value *CodebuildProjectBuildBatchConfig) {
	if err := c.validatePutBuildBatchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBuildBatchConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutCache(value *CodebuildProjectCache) {
	if err := c.validatePutCacheParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCache",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutEnvironment(value *CodebuildProjectEnvironment) {
	if err := c.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutFileSystemLocations(value interface{}) {
	if err := c.validatePutFileSystemLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFileSystemLocations",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutLogsConfig(value *CodebuildProjectLogsConfig) {
	if err := c.validatePutLogsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLogsConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutSecondaryArtifacts(value interface{}) {
	if err := c.validatePutSecondaryArtifactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecondaryArtifacts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutSecondarySources(value interface{}) {
	if err := c.validatePutSecondarySourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecondarySources",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutSecondarySourceVersion(value interface{}) {
	if err := c.validatePutSecondarySourceVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecondarySourceVersion",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutSource(value *CodebuildProjectSource) {
	if err := c.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutVpcConfig(value *CodebuildProjectVpcConfig) {
	if err := c.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) ResetBadgeEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetBadgeEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetBuildBatchConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildBatchConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetBuildTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetCache() {
	_jsii_.InvokeVoid(
		c,
		"resetCache",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetConcurrentBuildLimit() {
	_jsii_.InvokeVoid(
		c,
		"resetConcurrentBuildLimit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetEncryptionKey() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryptionKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetFileSystemLocations() {
	_jsii_.InvokeVoid(
		c,
		"resetFileSystemLocations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetLogsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLogsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetProjectVisibility() {
	_jsii_.InvokeVoid(
		c,
		"resetProjectVisibility",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetQueuedTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetQueuedTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetResourceAccessRole() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceAccessRole",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetSecondaryArtifacts() {
	_jsii_.InvokeVoid(
		c,
		"resetSecondaryArtifacts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetSecondarySources() {
	_jsii_.InvokeVoid(
		c,
		"resetSecondarySources",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetSecondarySourceVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetSecondarySourceVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetSourceVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProject) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProject) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProject) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProject) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

