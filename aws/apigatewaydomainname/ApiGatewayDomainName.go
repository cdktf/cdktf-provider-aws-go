// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewaydomainname

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/apigatewaydomainname/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/api_gateway_domain_name aws_api_gateway_domain_name}.
type ApiGatewayDomainName interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateArn() *string
	SetCertificateArn(val *string)
	CertificateArnInput() *string
	CertificateBody() *string
	SetCertificateBody(val *string)
	CertificateBodyInput() *string
	CertificateChain() *string
	SetCertificateChain(val *string)
	CertificateChainInput() *string
	CertificateName() *string
	SetCertificateName(val *string)
	CertificateNameInput() *string
	CertificatePrivateKey() *string
	SetCertificatePrivateKey(val *string)
	CertificatePrivateKeyInput() *string
	CertificateUploadDate() *string
	CloudfrontDomainName() *string
	CloudfrontZoneId() *string
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
	DomainName() *string
	SetDomainName(val *string)
	DomainNameId() *string
	DomainNameInput() *string
	EndpointConfiguration() ApiGatewayDomainNameEndpointConfigurationOutputReference
	EndpointConfigurationInput() *ApiGatewayDomainNameEndpointConfiguration
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
	MutualTlsAuthentication() ApiGatewayDomainNameMutualTlsAuthenticationOutputReference
	MutualTlsAuthenticationInput() *ApiGatewayDomainNameMutualTlsAuthentication
	// The tree node.
	Node() constructs.Node
	OwnershipVerificationCertificateArn() *string
	SetOwnershipVerificationCertificateArn(val *string)
	OwnershipVerificationCertificateArnInput() *string
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
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
	RegionalCertificateArn() *string
	SetRegionalCertificateArn(val *string)
	RegionalCertificateArnInput() *string
	RegionalCertificateName() *string
	SetRegionalCertificateName(val *string)
	RegionalCertificateNameInput() *string
	RegionalDomainName() *string
	RegionalZoneId() *string
	RegionInput() *string
	SecurityPolicy() *string
	SetSecurityPolicy(val *string)
	SecurityPolicyInput() *string
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
	PutEndpointConfiguration(value *ApiGatewayDomainNameEndpointConfiguration)
	PutMutualTlsAuthentication(value *ApiGatewayDomainNameMutualTlsAuthentication)
	ResetCertificateArn()
	ResetCertificateBody()
	ResetCertificateChain()
	ResetCertificateName()
	ResetCertificatePrivateKey()
	ResetEndpointConfiguration()
	ResetId()
	ResetMutualTlsAuthentication()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwnershipVerificationCertificateArn()
	ResetPolicy()
	ResetRegion()
	ResetRegionalCertificateArn()
	ResetRegionalCertificateName()
	ResetSecurityPolicy()
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

// The jsii proxy struct for ApiGatewayDomainName
type jsiiProxy_ApiGatewayDomainName struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiGatewayDomainName) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateChainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificatePrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificatePrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePrivateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CertificateUploadDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateUploadDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CloudfrontDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfrontDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) CloudfrontZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfrontZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) DomainNameId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) EndpointConfiguration() ApiGatewayDomainNameEndpointConfigurationOutputReference {
	var returns ApiGatewayDomainNameEndpointConfigurationOutputReference
	_jsii_.Get(
		j,
		"endpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) EndpointConfigurationInput() *ApiGatewayDomainNameEndpointConfiguration {
	var returns *ApiGatewayDomainNameEndpointConfiguration
	_jsii_.Get(
		j,
		"endpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) MutualTlsAuthentication() ApiGatewayDomainNameMutualTlsAuthenticationOutputReference {
	var returns ApiGatewayDomainNameMutualTlsAuthenticationOutputReference
	_jsii_.Get(
		j,
		"mutualTlsAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) MutualTlsAuthenticationInput() *ApiGatewayDomainNameMutualTlsAuthentication {
	var returns *ApiGatewayDomainNameMutualTlsAuthentication
	_jsii_.Get(
		j,
		"mutualTlsAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) OwnershipVerificationCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownershipVerificationCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) OwnershipVerificationCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownershipVerificationCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) RegionalCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) RegionalCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) RegionalCertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalCertificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) RegionalCertificateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalCertificateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) RegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) RegionalZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) SecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) SecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayDomainName) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/api_gateway_domain_name aws_api_gateway_domain_name} Resource.
func NewApiGatewayDomainName(scope constructs.Construct, id *string, config *ApiGatewayDomainNameConfig) ApiGatewayDomainName {
	_init_.Initialize()

	if err := validateNewApiGatewayDomainNameParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGatewayDomainName{}

	_jsii_.Create(
		"@cdktf/provider-aws.apiGatewayDomainName.ApiGatewayDomainName",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/api_gateway_domain_name aws_api_gateway_domain_name} Resource.
func NewApiGatewayDomainName_Override(a ApiGatewayDomainName, scope constructs.Construct, id *string, config *ApiGatewayDomainNameConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.apiGatewayDomainName.ApiGatewayDomainName",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetCertificateArn(val *string) {
	if err := j.validateSetCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetCertificateBody(val *string) {
	if err := j.validateSetCertificateBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateBody",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetCertificateChain(val *string) {
	if err := j.validateSetCertificateChainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateChain",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetCertificateName(val *string) {
	if err := j.validateSetCertificateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateName",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetCertificatePrivateKey(val *string) {
	if err := j.validateSetCertificatePrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificatePrivateKey",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetOwnershipVerificationCertificateArn(val *string) {
	if err := j.validateSetOwnershipVerificationCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownershipVerificationCertificateArn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetPolicy(val *string) {
	if err := j.validateSetPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetRegionalCertificateArn(val *string) {
	if err := j.validateSetRegionalCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionalCertificateArn",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetRegionalCertificateName(val *string) {
	if err := j.validateSetRegionalCertificateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionalCertificateName",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetSecurityPolicy(val *string) {
	if err := j.validateSetSecurityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityPolicy",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayDomainName)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a ApiGatewayDomainName resource upon running "cdktf plan <stack-name>".
func ApiGatewayDomainName_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApiGatewayDomainName_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.apiGatewayDomainName.ApiGatewayDomainName",
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
func ApiGatewayDomainName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiGatewayDomainName_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.apiGatewayDomainName.ApiGatewayDomainName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiGatewayDomainName_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiGatewayDomainName_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.apiGatewayDomainName.ApiGatewayDomainName",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiGatewayDomainName_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiGatewayDomainName_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.apiGatewayDomainName.ApiGatewayDomainName",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiGatewayDomainName_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.apiGatewayDomainName.ApiGatewayDomainName",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApiGatewayDomainName) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiGatewayDomainName) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiGatewayDomainName) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiGatewayDomainName) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiGatewayDomainName) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiGatewayDomainName) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiGatewayDomainName) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiGatewayDomainName) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiGatewayDomainName) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiGatewayDomainName) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayDomainName) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiGatewayDomainName) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) PutEndpointConfiguration(value *ApiGatewayDomainNameEndpointConfiguration) {
	if err := a.validatePutEndpointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEndpointConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) PutMutualTlsAuthentication(value *ApiGatewayDomainNameMutualTlsAuthentication) {
	if err := a.validatePutMutualTlsAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMutualTlsAuthentication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetCertificateBody() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateBody",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetCertificateChain() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateChain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetCertificateName() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetCertificatePrivateKey() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificatePrivateKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetEndpointConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetMutualTlsAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetMutualTlsAuthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetOwnershipVerificationCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetOwnershipVerificationCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetRegionalCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRegionalCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetRegionalCertificateName() {
	_jsii_.InvokeVoid(
		a,
		"resetRegionalCertificateName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetSecurityPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayDomainName) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayDomainName) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayDomainName) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayDomainName) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayDomainName) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayDomainName) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

