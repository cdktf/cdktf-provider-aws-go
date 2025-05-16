// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package amplifyapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/amplifyapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/amplify_app aws_amplify_app}.
type AmplifyApp interface {
	cdktf.TerraformResource
	AccessToken() *string
	SetAccessToken(val *string)
	AccessTokenInput() *string
	Arn() *string
	AutoBranchCreationConfig() AmplifyAppAutoBranchCreationConfigOutputReference
	AutoBranchCreationConfigInput() *AmplifyAppAutoBranchCreationConfig
	AutoBranchCreationPatterns() *[]*string
	SetAutoBranchCreationPatterns(val *[]*string)
	AutoBranchCreationPatternsInput() *[]*string
	BasicAuthCredentials() *string
	SetBasicAuthCredentials(val *string)
	BasicAuthCredentialsInput() *string
	BuildSpec() *string
	SetBuildSpec(val *string)
	BuildSpecInput() *string
	CacheConfig() AmplifyAppCacheConfigOutputReference
	CacheConfigInput() *AmplifyAppCacheConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComputeRoleArn() *string
	SetComputeRoleArn(val *string)
	ComputeRoleArnInput() *string
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
	CustomHeaders() *string
	SetCustomHeaders(val *string)
	CustomHeadersInput() *string
	CustomRule() AmplifyAppCustomRuleList
	CustomRuleInput() interface{}
	DefaultDomain() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnableAutoBranchCreation() interface{}
	SetEnableAutoBranchCreation(val interface{})
	EnableAutoBranchCreationInput() interface{}
	EnableBasicAuth() interface{}
	SetEnableBasicAuth(val interface{})
	EnableBasicAuthInput() interface{}
	EnableBranchAutoBuild() interface{}
	SetEnableBranchAutoBuild(val interface{})
	EnableBranchAutoBuildInput() interface{}
	EnableBranchAutoDeletion() interface{}
	SetEnableBranchAutoDeletion(val interface{})
	EnableBranchAutoDeletionInput() interface{}
	EnvironmentVariables() *map[string]*string
	SetEnvironmentVariables(val *map[string]*string)
	EnvironmentVariablesInput() *map[string]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IamServiceRoleArn() *string
	SetIamServiceRoleArn(val *string)
	IamServiceRoleArnInput() *string
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
	OauthToken() *string
	SetOauthToken(val *string)
	OauthTokenInput() *string
	Platform() *string
	SetPlatform(val *string)
	PlatformInput() *string
	ProductionBranch() AmplifyAppProductionBranchList
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
	Repository() *string
	SetRepository(val *string)
	RepositoryInput() *string
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
	PutAutoBranchCreationConfig(value *AmplifyAppAutoBranchCreationConfig)
	PutCacheConfig(value *AmplifyAppCacheConfig)
	PutCustomRule(value interface{})
	ResetAccessToken()
	ResetAutoBranchCreationConfig()
	ResetAutoBranchCreationPatterns()
	ResetBasicAuthCredentials()
	ResetBuildSpec()
	ResetCacheConfig()
	ResetComputeRoleArn()
	ResetCustomHeaders()
	ResetCustomRule()
	ResetDescription()
	ResetEnableAutoBranchCreation()
	ResetEnableBasicAuth()
	ResetEnableBranchAutoBuild()
	ResetEnableBranchAutoDeletion()
	ResetEnvironmentVariables()
	ResetIamServiceRoleArn()
	ResetId()
	ResetOauthToken()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlatform()
	ResetRepository()
	ResetTags()
	ResetTagsAll()
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

// The jsii proxy struct for AmplifyApp
type jsiiProxy_AmplifyApp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AmplifyApp) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) AutoBranchCreationConfig() AmplifyAppAutoBranchCreationConfigOutputReference {
	var returns AmplifyAppAutoBranchCreationConfigOutputReference
	_jsii_.Get(
		j,
		"autoBranchCreationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) AutoBranchCreationConfigInput() *AmplifyAppAutoBranchCreationConfig {
	var returns *AmplifyAppAutoBranchCreationConfig
	_jsii_.Get(
		j,
		"autoBranchCreationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) AutoBranchCreationPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoBranchCreationPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) AutoBranchCreationPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoBranchCreationPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) BasicAuthCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuthCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) BasicAuthCredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuthCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) BuildSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) BuildSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) CacheConfig() AmplifyAppCacheConfigOutputReference {
	var returns AmplifyAppCacheConfigOutputReference
	_jsii_.Get(
		j,
		"cacheConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) CacheConfigInput() *AmplifyAppCacheConfig {
	var returns *AmplifyAppCacheConfig
	_jsii_.Get(
		j,
		"cacheConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) ComputeRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) ComputeRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) CustomHeaders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) CustomHeadersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) CustomRule() AmplifyAppCustomRuleList {
	var returns AmplifyAppCustomRuleList
	_jsii_.Get(
		j,
		"customRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) CustomRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) DefaultDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableAutoBranchCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoBranchCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableAutoBranchCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoBranchCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableBasicAuth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBasicAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableBasicAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBasicAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableBranchAutoBuild() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBranchAutoBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableBranchAutoBuildInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBranchAutoBuildInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableBranchAutoDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBranchAutoDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnableBranchAutoDeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBranchAutoDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) IamServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamServiceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) IamServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamServiceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) OauthToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) OauthTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) PlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) ProductionBranch() AmplifyAppProductionBranchList {
	var returns AmplifyAppProductionBranchList
	_jsii_.Get(
		j,
		"productionBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Repository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) RepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/amplify_app aws_amplify_app} Resource.
func NewAmplifyApp(scope constructs.Construct, id *string, config *AmplifyAppConfig) AmplifyApp {
	_init_.Initialize()

	if err := validateNewAmplifyAppParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AmplifyApp{}

	_jsii_.Create(
		"@cdktf/provider-aws.amplifyApp.AmplifyApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/amplify_app aws_amplify_app} Resource.
func NewAmplifyApp_Override(a AmplifyApp, scope constructs.Construct, id *string, config *AmplifyAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.amplifyApp.AmplifyApp",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AmplifyApp)SetAccessToken(val *string) {
	if err := j.validateSetAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetAutoBranchCreationPatterns(val *[]*string) {
	if err := j.validateSetAutoBranchCreationPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoBranchCreationPatterns",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetBasicAuthCredentials(val *string) {
	if err := j.validateSetBasicAuthCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"basicAuthCredentials",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetBuildSpec(val *string) {
	if err := j.validateSetBuildSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildSpec",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetComputeRoleArn(val *string) {
	if err := j.validateSetComputeRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeRoleArn",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetCustomHeaders(val *string) {
	if err := j.validateSetCustomHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customHeaders",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetEnableAutoBranchCreation(val interface{}) {
	if err := j.validateSetEnableAutoBranchCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutoBranchCreation",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetEnableBasicAuth(val interface{}) {
	if err := j.validateSetEnableBasicAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBasicAuth",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetEnableBranchAutoBuild(val interface{}) {
	if err := j.validateSetEnableBranchAutoBuildParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBranchAutoBuild",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetEnableBranchAutoDeletion(val interface{}) {
	if err := j.validateSetEnableBranchAutoDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBranchAutoDeletion",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetIamServiceRoleArn(val *string) {
	if err := j.validateSetIamServiceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamServiceRoleArn",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetOauthToken(val *string) {
	if err := j.validateSetOauthTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthToken",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetPlatform(val *string) {
	if err := j.validateSetPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platform",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetRepository(val *string) {
	if err := j.validateSetRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repository",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AmplifyApp)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a AmplifyApp resource upon running "cdktf plan <stack-name>".
func AmplifyApp_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAmplifyApp_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.amplifyApp.AmplifyApp",
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
func AmplifyApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAmplifyApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.amplifyApp.AmplifyApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AmplifyApp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAmplifyApp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.amplifyApp.AmplifyApp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AmplifyApp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAmplifyApp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.amplifyApp.AmplifyApp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AmplifyApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.amplifyApp.AmplifyApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AmplifyApp) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AmplifyApp) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AmplifyApp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AmplifyApp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AmplifyApp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AmplifyApp) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AmplifyApp) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AmplifyApp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AmplifyApp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AmplifyApp) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AmplifyApp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AmplifyApp) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyApp) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AmplifyApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AmplifyApp) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AmplifyApp) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AmplifyApp) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AmplifyApp) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AmplifyApp) PutAutoBranchCreationConfig(value *AmplifyAppAutoBranchCreationConfig) {
	if err := a.validatePutAutoBranchCreationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutoBranchCreationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AmplifyApp) PutCacheConfig(value *AmplifyAppCacheConfig) {
	if err := a.validatePutCacheConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCacheConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AmplifyApp) PutCustomRule(value interface{}) {
	if err := a.validatePutCustomRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AmplifyApp) ResetAccessToken() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetAutoBranchCreationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoBranchCreationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetAutoBranchCreationPatterns() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoBranchCreationPatterns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetBasicAuthCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetBasicAuthCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetBuildSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetBuildSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetCacheConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetComputeRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetComputeRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetCustomHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetCustomRule() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomRule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetEnableAutoBranchCreation() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableAutoBranchCreation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetEnableBasicAuth() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBasicAuth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetEnableBranchAutoBuild() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBranchAutoBuild",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetEnableBranchAutoDeletion() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBranchAutoDeletion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetIamServiceRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetIamServiceRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetOauthToken() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetPlatform() {
	_jsii_.InvokeVoid(
		a,
		"resetPlatform",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyApp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyApp) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyApp) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyApp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

