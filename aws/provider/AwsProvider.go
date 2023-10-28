// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs aws}.
type AwsProvider interface {
	cdktf.TerraformProvider
	AccessKey() *string
	SetAccessKey(val *string)
	AccessKeyInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AllowedAccountIds() *[]*string
	SetAllowedAccountIds(val *[]*string)
	AllowedAccountIdsInput() *[]*string
	AssumeRole() interface{}
	SetAssumeRole(val interface{})
	AssumeRoleInput() interface{}
	AssumeRoleWithWebIdentity() interface{}
	SetAssumeRoleWithWebIdentity(val interface{})
	AssumeRoleWithWebIdentityInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CustomCaBundle() *string
	SetCustomCaBundle(val *string)
	CustomCaBundleInput() *string
	DefaultTags() interface{}
	SetDefaultTags(val interface{})
	DefaultTagsInput() interface{}
	Ec2MetadataServiceEndpoint() *string
	SetEc2MetadataServiceEndpoint(val *string)
	Ec2MetadataServiceEndpointInput() *string
	Ec2MetadataServiceEndpointMode() *string
	SetEc2MetadataServiceEndpointMode(val *string)
	Ec2MetadataServiceEndpointModeInput() *string
	Endpoints() interface{}
	SetEndpoints(val interface{})
	EndpointsInput() interface{}
	ForbiddenAccountIds() *[]*string
	SetForbiddenAccountIds(val *[]*string)
	ForbiddenAccountIdsInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpProxy() *string
	SetHttpProxy(val *string)
	HttpProxyInput() *string
	IgnoreTags() interface{}
	SetIgnoreTags(val interface{})
	IgnoreTagsInput() interface{}
	Insecure() interface{}
	SetInsecure(val interface{})
	InsecureInput() interface{}
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RetryMode() *string
	SetRetryMode(val *string)
	RetryModeInput() *string
	S3UsEast1RegionalEndpoint() *string
	SetS3UsEast1RegionalEndpoint(val *string)
	S3UsEast1RegionalEndpointInput() *string
	S3UsePathStyle() interface{}
	SetS3UsePathStyle(val interface{})
	S3UsePathStyleInput() interface{}
	SecretKey() *string
	SetSecretKey(val *string)
	SecretKeyInput() *string
	SharedConfigFiles() *[]*string
	SetSharedConfigFiles(val *[]*string)
	SharedConfigFilesInput() *[]*string
	SharedCredentialsFiles() *[]*string
	SetSharedCredentialsFiles(val *[]*string)
	SharedCredentialsFilesInput() *[]*string
	SkipCredentialsValidation() interface{}
	SetSkipCredentialsValidation(val interface{})
	SkipCredentialsValidationInput() interface{}
	SkipMetadataApiCheck() *string
	SetSkipMetadataApiCheck(val *string)
	SkipMetadataApiCheckInput() *string
	SkipRegionValidation() interface{}
	SetSkipRegionValidation(val interface{})
	SkipRegionValidationInput() interface{}
	SkipRequestingAccountId() interface{}
	SetSkipRequestingAccountId(val interface{})
	SkipRequestingAccountIdInput() interface{}
	StsRegion() *string
	SetStsRegion(val *string)
	StsRegionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	UseDualstackEndpoint() interface{}
	SetUseDualstackEndpoint(val interface{})
	UseDualstackEndpointInput() interface{}
	UseFipsEndpoint() interface{}
	SetUseFipsEndpoint(val interface{})
	UseFipsEndpointInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccessKey()
	ResetAlias()
	ResetAllowedAccountIds()
	ResetAssumeRole()
	ResetAssumeRoleWithWebIdentity()
	ResetCustomCaBundle()
	ResetDefaultTags()
	ResetEc2MetadataServiceEndpoint()
	ResetEc2MetadataServiceEndpointMode()
	ResetEndpoints()
	ResetForbiddenAccountIds()
	ResetHttpProxy()
	ResetIgnoreTags()
	ResetInsecure()
	ResetMaxRetries()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProfile()
	ResetRegion()
	ResetRetryMode()
	ResetS3UsEast1RegionalEndpoint()
	ResetS3UsePathStyle()
	ResetSecretKey()
	ResetSharedConfigFiles()
	ResetSharedCredentialsFiles()
	ResetSkipCredentialsValidation()
	ResetSkipMetadataApiCheck()
	ResetSkipRegionValidation()
	ResetSkipRequestingAccountId()
	ResetStsRegion()
	ResetToken()
	ResetUseDualstackEndpoint()
	ResetUseFipsEndpoint()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AwsProvider
type jsiiProxy_AwsProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_AwsProvider) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AllowedAccountIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAccountIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AllowedAccountIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAccountIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AssumeRole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assumeRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AssumeRoleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assumeRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AssumeRoleWithWebIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assumeRoleWithWebIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AssumeRoleWithWebIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assumeRoleWithWebIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) CustomCaBundle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCaBundle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) CustomCaBundleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCaBundleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) DefaultTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) DefaultTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Ec2MetadataServiceEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2MetadataServiceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Ec2MetadataServiceEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2MetadataServiceEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Ec2MetadataServiceEndpointMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2MetadataServiceEndpointMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Ec2MetadataServiceEndpointModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2MetadataServiceEndpointModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Endpoints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) EndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) ForbiddenAccountIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"forbiddenAccountIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) ForbiddenAccountIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"forbiddenAccountIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) HttpProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) HttpProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) IgnoreTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) IgnoreTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Insecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) InsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) RetryMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) RetryModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) S3UsEast1RegionalEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3UsEast1RegionalEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) S3UsEast1RegionalEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3UsEast1RegionalEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) S3UsePathStyle() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3UsePathStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) S3UsePathStyleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3UsePathStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SharedConfigFiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedConfigFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SharedConfigFilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedConfigFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SharedCredentialsFiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedCredentialsFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SharedCredentialsFilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedCredentialsFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipCredentialsValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipCredentialsValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipCredentialsValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipCredentialsValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipMetadataApiCheck() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipMetadataApiCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipMetadataApiCheckInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipMetadataApiCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipRegionValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRegionValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipRegionValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRegionValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipRequestingAccountId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRequestingAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipRequestingAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRequestingAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) StsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) StsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) UseDualstackEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDualstackEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) UseDualstackEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDualstackEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) UseFipsEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useFipsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) UseFipsEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useFipsEndpointInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs aws} Resource.
func NewAwsProvider(scope constructs.Construct, id *string, config *AwsProviderConfig) AwsProvider {
	_init_.Initialize()

	if err := validateNewAwsProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsProvider{}

	_jsii_.Create(
		"@cdktf/provider-aws.provider.AwsProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs aws} Resource.
func NewAwsProvider_Override(a AwsProvider, scope constructs.Construct, id *string, config *AwsProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.provider.AwsProvider",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsProvider)SetAccessKey(val *string) {
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetAllowedAccountIds(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedAccountIds",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetAssumeRole(val interface{}) {
	if err := j.validateSetAssumeRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assumeRole",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetAssumeRoleWithWebIdentity(val interface{}) {
	if err := j.validateSetAssumeRoleWithWebIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assumeRoleWithWebIdentity",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetCustomCaBundle(val *string) {
	_jsii_.Set(
		j,
		"customCaBundle",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetDefaultTags(val interface{}) {
	if err := j.validateSetDefaultTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTags",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetEc2MetadataServiceEndpoint(val *string) {
	_jsii_.Set(
		j,
		"ec2MetadataServiceEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetEc2MetadataServiceEndpointMode(val *string) {
	_jsii_.Set(
		j,
		"ec2MetadataServiceEndpointMode",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetEndpoints(val interface{}) {
	if err := j.validateSetEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoints",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetForbiddenAccountIds(val *[]*string) {
	_jsii_.Set(
		j,
		"forbiddenAccountIds",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetHttpProxy(val *string) {
	_jsii_.Set(
		j,
		"httpProxy",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetIgnoreTags(val interface{}) {
	if err := j.validateSetIgnoreTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreTags",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetInsecure(val interface{}) {
	if err := j.validateSetInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecure",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetProfile(val *string) {
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetRetryMode(val *string) {
	_jsii_.Set(
		j,
		"retryMode",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetS3UsEast1RegionalEndpoint(val *string) {
	_jsii_.Set(
		j,
		"s3UsEast1RegionalEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetS3UsePathStyle(val interface{}) {
	if err := j.validateSetS3UsePathStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3UsePathStyle",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetSecretKey(val *string) {
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetSharedConfigFiles(val *[]*string) {
	_jsii_.Set(
		j,
		"sharedConfigFiles",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetSharedCredentialsFiles(val *[]*string) {
	_jsii_.Set(
		j,
		"sharedCredentialsFiles",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetSkipCredentialsValidation(val interface{}) {
	if err := j.validateSetSkipCredentialsValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipCredentialsValidation",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetSkipMetadataApiCheck(val *string) {
	_jsii_.Set(
		j,
		"skipMetadataApiCheck",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetSkipRegionValidation(val interface{}) {
	if err := j.validateSetSkipRegionValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipRegionValidation",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetSkipRequestingAccountId(val interface{}) {
	if err := j.validateSetSkipRequestingAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipRequestingAccountId",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetStsRegion(val *string) {
	_jsii_.Set(
		j,
		"stsRegion",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetUseDualstackEndpoint(val interface{}) {
	if err := j.validateSetUseDualstackEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDualstackEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetUseFipsEndpoint(val interface{}) {
	if err := j.validateSetUseFipsEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useFipsEndpoint",
		val,
	)
}

// Generates CDKTF code for importing a AwsProvider resource upon running "cdktf plan <stack-name>".
func AwsProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAwsProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.provider.AwsProvider",
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
func AwsProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.provider.AwsProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.provider.AwsProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.provider.AwsProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.provider.AwsProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsProvider) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsProvider) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsProvider) ResetAccessKey() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetAllowedAccountIds() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedAccountIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetAssumeRole() {
	_jsii_.InvokeVoid(
		a,
		"resetAssumeRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetAssumeRoleWithWebIdentity() {
	_jsii_.InvokeVoid(
		a,
		"resetAssumeRoleWithWebIdentity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetCustomCaBundle() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomCaBundle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetDefaultTags() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetEc2MetadataServiceEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2MetadataServiceEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetEc2MetadataServiceEndpointMode() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2MetadataServiceEndpointMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetEndpoints() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpoints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetForbiddenAccountIds() {
	_jsii_.InvokeVoid(
		a,
		"resetForbiddenAccountIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetHttpProxy() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpProxy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetIgnoreTags() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetInsecure() {
	_jsii_.InvokeVoid(
		a,
		"resetInsecure",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetProfile() {
	_jsii_.InvokeVoid(
		a,
		"resetProfile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetRetryMode() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetS3UsEast1RegionalEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetS3UsEast1RegionalEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetS3UsePathStyle() {
	_jsii_.InvokeVoid(
		a,
		"resetS3UsePathStyle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSecretKey() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSharedConfigFiles() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedConfigFiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSharedCredentialsFiles() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedCredentialsFiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSkipCredentialsValidation() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipCredentialsValidation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSkipMetadataApiCheck() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipMetadataApiCheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSkipRegionValidation() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipRegionValidation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSkipRequestingAccountId() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipRequestingAccountId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetStsRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetStsRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetToken() {
	_jsii_.InvokeVoid(
		a,
		"resetToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetUseDualstackEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetUseDualstackEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetUseFipsEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetUseFipsEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

