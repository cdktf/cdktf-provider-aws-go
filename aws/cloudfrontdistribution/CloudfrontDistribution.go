// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/cloudfrontdistribution/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/resources/cloudfront_distribution aws_cloudfront_distribution}.
type CloudfrontDistribution interface {
	cdktf.TerraformResource
	Aliases() *[]*string
	SetAliases(val *[]*string)
	AliasesInput() *[]*string
	Arn() *string
	CallerReference() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContinuousDeploymentPolicyId() *string
	SetContinuousDeploymentPolicyId(val *string)
	ContinuousDeploymentPolicyIdInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomErrorResponse() CloudfrontDistributionCustomErrorResponseList
	CustomErrorResponseInput() interface{}
	DefaultCacheBehavior() CloudfrontDistributionDefaultCacheBehaviorOutputReference
	DefaultCacheBehaviorInput() *CloudfrontDistributionDefaultCacheBehavior
	DefaultRootObject() *string
	SetDefaultRootObject(val *string)
	DefaultRootObjectInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DomainName() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Etag() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostedZoneId() *string
	HttpVersion() *string
	SetHttpVersion(val *string)
	HttpVersionInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InProgressValidationBatches() *float64
	IsIpv6Enabled() interface{}
	SetIsIpv6Enabled(val interface{})
	IsIpv6EnabledInput() interface{}
	LastModifiedTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoggingConfig() CloudfrontDistributionLoggingConfigOutputReference
	LoggingConfigInput() *CloudfrontDistributionLoggingConfig
	// The tree node.
	Node() constructs.Node
	OrderedCacheBehavior() CloudfrontDistributionOrderedCacheBehaviorList
	OrderedCacheBehaviorInput() interface{}
	Origin() CloudfrontDistributionOriginList
	OriginGroup() CloudfrontDistributionOriginGroupList
	OriginGroupInput() interface{}
	OriginInput() interface{}
	PriceClass() *string
	SetPriceClass(val *string)
	PriceClassInput() *string
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
	Restrictions() CloudfrontDistributionRestrictionsOutputReference
	RestrictionsInput() *CloudfrontDistributionRestrictions
	RetainOnDelete() interface{}
	SetRetainOnDelete(val interface{})
	RetainOnDeleteInput() interface{}
	Staging() interface{}
	SetStaging(val interface{})
	StagingInput() interface{}
	Status() *string
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
	TrustedKeyGroups() CloudfrontDistributionTrustedKeyGroupsList
	TrustedSigners() CloudfrontDistributionTrustedSignersList
	ViewerCertificate() CloudfrontDistributionViewerCertificateOutputReference
	ViewerCertificateInput() *CloudfrontDistributionViewerCertificate
	WaitForDeployment() interface{}
	SetWaitForDeployment(val interface{})
	WaitForDeploymentInput() interface{}
	WebAclId() *string
	SetWebAclId(val *string)
	WebAclIdInput() *string
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
	PutCustomErrorResponse(value interface{})
	PutDefaultCacheBehavior(value *CloudfrontDistributionDefaultCacheBehavior)
	PutLoggingConfig(value *CloudfrontDistributionLoggingConfig)
	PutOrderedCacheBehavior(value interface{})
	PutOrigin(value interface{})
	PutOriginGroup(value interface{})
	PutRestrictions(value *CloudfrontDistributionRestrictions)
	PutViewerCertificate(value *CloudfrontDistributionViewerCertificate)
	ResetAliases()
	ResetComment()
	ResetContinuousDeploymentPolicyId()
	ResetCustomErrorResponse()
	ResetDefaultRootObject()
	ResetHttpVersion()
	ResetId()
	ResetIsIpv6Enabled()
	ResetLoggingConfig()
	ResetOrderedCacheBehavior()
	ResetOriginGroup()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPriceClass()
	ResetRetainOnDelete()
	ResetStaging()
	ResetTags()
	ResetTagsAll()
	ResetWaitForDeployment()
	ResetWebAclId()
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

// The jsii proxy struct for CloudfrontDistribution
type jsiiProxy_CloudfrontDistribution struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudfrontDistribution) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) AliasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) CallerReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) ContinuousDeploymentPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"continuousDeploymentPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) ContinuousDeploymentPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"continuousDeploymentPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) CustomErrorResponse() CloudfrontDistributionCustomErrorResponseList {
	var returns CloudfrontDistributionCustomErrorResponseList
	_jsii_.Get(
		j,
		"customErrorResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) CustomErrorResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customErrorResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) DefaultCacheBehavior() CloudfrontDistributionDefaultCacheBehaviorOutputReference {
	var returns CloudfrontDistributionDefaultCacheBehaviorOutputReference
	_jsii_.Get(
		j,
		"defaultCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) DefaultCacheBehaviorInput() *CloudfrontDistributionDefaultCacheBehavior {
	var returns *CloudfrontDistributionDefaultCacheBehavior
	_jsii_.Get(
		j,
		"defaultCacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) DefaultRootObject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) DefaultRootObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) HttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) HttpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) InProgressValidationBatches() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inProgressValidationBatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) IsIpv6Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIpv6Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) IsIpv6EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIpv6EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) LoggingConfig() CloudfrontDistributionLoggingConfigOutputReference {
	var returns CloudfrontDistributionLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) LoggingConfigInput() *CloudfrontDistributionLoggingConfig {
	var returns *CloudfrontDistributionLoggingConfig
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) OrderedCacheBehavior() CloudfrontDistributionOrderedCacheBehaviorList {
	var returns CloudfrontDistributionOrderedCacheBehaviorList
	_jsii_.Get(
		j,
		"orderedCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) OrderedCacheBehaviorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orderedCacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Origin() CloudfrontDistributionOriginList {
	var returns CloudfrontDistributionOriginList
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) OriginGroup() CloudfrontDistributionOriginGroupList {
	var returns CloudfrontDistributionOriginGroupList
	_jsii_.Get(
		j,
		"originGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) OriginGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) PriceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) PriceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Restrictions() CloudfrontDistributionRestrictionsOutputReference {
	var returns CloudfrontDistributionRestrictionsOutputReference
	_jsii_.Get(
		j,
		"restrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) RestrictionsInput() *CloudfrontDistributionRestrictions {
	var returns *CloudfrontDistributionRestrictions
	_jsii_.Get(
		j,
		"restrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) RetainOnDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) RetainOnDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainOnDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Staging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) StagingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stagingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) TrustedKeyGroups() CloudfrontDistributionTrustedKeyGroupsList {
	var returns CloudfrontDistributionTrustedKeyGroupsList
	_jsii_.Get(
		j,
		"trustedKeyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) TrustedSigners() CloudfrontDistributionTrustedSignersList {
	var returns CloudfrontDistributionTrustedSignersList
	_jsii_.Get(
		j,
		"trustedSigners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) ViewerCertificate() CloudfrontDistributionViewerCertificateOutputReference {
	var returns CloudfrontDistributionViewerCertificateOutputReference
	_jsii_.Get(
		j,
		"viewerCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) ViewerCertificateInput() *CloudfrontDistributionViewerCertificate {
	var returns *CloudfrontDistributionViewerCertificate
	_jsii_.Get(
		j,
		"viewerCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) WaitForDeployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) WaitForDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) WebAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) WebAclIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/resources/cloudfront_distribution aws_cloudfront_distribution} Resource.
func NewCloudfrontDistribution(scope constructs.Construct, id *string, config *CloudfrontDistributionConfig) CloudfrontDistribution {
	_init_.Initialize()

	if err := validateNewCloudfrontDistributionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontDistribution{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistribution",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/resources/cloudfront_distribution aws_cloudfront_distribution} Resource.
func NewCloudfrontDistribution_Override(c CloudfrontDistribution, scope constructs.Construct, id *string, config *CloudfrontDistributionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistribution",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetAliases(val *[]*string) {
	if err := j.validateSetAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliases",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetContinuousDeploymentPolicyId(val *string) {
	if err := j.validateSetContinuousDeploymentPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"continuousDeploymentPolicyId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetDefaultRootObject(val *string) {
	if err := j.validateSetDefaultRootObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRootObject",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetHttpVersion(val *string) {
	if err := j.validateSetHttpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpVersion",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetIsIpv6Enabled(val interface{}) {
	if err := j.validateSetIsIpv6EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isIpv6Enabled",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetPriceClass(val *string) {
	if err := j.validateSetPriceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priceClass",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetRetainOnDelete(val interface{}) {
	if err := j.validateSetRetainOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retainOnDelete",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetStaging(val interface{}) {
	if err := j.validateSetStagingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staging",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetWaitForDeployment(val interface{}) {
	if err := j.validateSetWaitForDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForDeployment",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution)SetWebAclId(val *string) {
	if err := j.validateSetWebAclIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webAclId",
		val,
	)
}

// Generates CDKTF code for importing a CloudfrontDistribution resource upon running "cdktf plan <stack-name>".
func CloudfrontDistribution_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudfrontDistribution_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistribution",
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
func CloudfrontDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudfrontDistribution_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudfrontDistribution_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudfrontDistribution_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistribution",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudfrontDistribution_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudfrontDistribution_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistribution",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudfrontDistribution_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistribution",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudfrontDistribution) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontDistribution) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontDistribution) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontDistribution) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontDistribution) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontDistribution) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontDistribution) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontDistribution) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontDistribution) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontDistribution) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistribution) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontDistribution) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) PutCustomErrorResponse(value interface{}) {
	if err := c.validatePutCustomErrorResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomErrorResponse",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) PutDefaultCacheBehavior(value *CloudfrontDistributionDefaultCacheBehavior) {
	if err := c.validatePutDefaultCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDefaultCacheBehavior",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) PutLoggingConfig(value *CloudfrontDistributionLoggingConfig) {
	if err := c.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) PutOrderedCacheBehavior(value interface{}) {
	if err := c.validatePutOrderedCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOrderedCacheBehavior",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) PutOrigin(value interface{}) {
	if err := c.validatePutOriginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOrigin",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) PutOriginGroup(value interface{}) {
	if err := c.validatePutOriginGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOriginGroup",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) PutRestrictions(value *CloudfrontDistributionRestrictions) {
	if err := c.validatePutRestrictionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRestrictions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) PutViewerCertificate(value *CloudfrontDistributionViewerCertificate) {
	if err := c.validatePutViewerCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putViewerCertificate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetAliases() {
	_jsii_.InvokeVoid(
		c,
		"resetAliases",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetContinuousDeploymentPolicyId() {
	_jsii_.InvokeVoid(
		c,
		"resetContinuousDeploymentPolicyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetCustomErrorResponse() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomErrorResponse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetDefaultRootObject() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultRootObject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetHttpVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetIsIpv6Enabled() {
	_jsii_.InvokeVoid(
		c,
		"resetIsIpv6Enabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetOrderedCacheBehavior() {
	_jsii_.InvokeVoid(
		c,
		"resetOrderedCacheBehavior",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetOriginGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetPriceClass() {
	_jsii_.InvokeVoid(
		c,
		"resetPriceClass",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetRetainOnDelete() {
	_jsii_.InvokeVoid(
		c,
		"resetRetainOnDelete",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetStaging() {
	_jsii_.InvokeVoid(
		c,
		"resetStaging",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetWaitForDeployment() {
	_jsii_.InvokeVoid(
		c,
		"resetWaitForDeployment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetWebAclId() {
	_jsii_.InvokeVoid(
		c,
		"resetWebAclId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistribution) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistribution) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistribution) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistribution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistribution) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

