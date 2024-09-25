// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedaccessendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/verifiedaccessendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/verifiedaccess_endpoint aws_verifiedaccess_endpoint}.
type VerifiedaccessEndpoint interface {
	cdktf.TerraformResource
	ApplicationDomain() *string
	SetApplicationDomain(val *string)
	ApplicationDomainInput() *string
	AttachmentType() *string
	SetAttachmentType(val *string)
	AttachmentTypeInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DeviceValidationDomain() *string
	DomainCertificateArn() *string
	SetDomainCertificateArn(val *string)
	DomainCertificateArnInput() *string
	EndpointDomain() *string
	EndpointDomainPrefix() *string
	SetEndpointDomainPrefix(val *string)
	EndpointDomainPrefixInput() *string
	EndpointType() *string
	SetEndpointType(val *string)
	EndpointTypeInput() *string
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
	LoadBalancerOptions() VerifiedaccessEndpointLoadBalancerOptionsOutputReference
	LoadBalancerOptionsInput() *VerifiedaccessEndpointLoadBalancerOptions
	NetworkInterfaceOptions() VerifiedaccessEndpointNetworkInterfaceOptionsOutputReference
	NetworkInterfaceOptionsInput() *VerifiedaccessEndpointNetworkInterfaceOptions
	// The tree node.
	Node() constructs.Node
	PolicyDocument() *string
	SetPolicyDocument(val *string)
	PolicyDocumentInput() *string
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
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SseSpecification() VerifiedaccessEndpointSseSpecificationOutputReference
	SseSpecificationInput() *VerifiedaccessEndpointSseSpecification
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
	Timeouts() VerifiedaccessEndpointTimeoutsOutputReference
	TimeoutsInput() interface{}
	VerifiedAccessGroupId() *string
	SetVerifiedAccessGroupId(val *string)
	VerifiedAccessGroupIdInput() *string
	VerifiedAccessInstanceId() *string
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
	PutLoadBalancerOptions(value *VerifiedaccessEndpointLoadBalancerOptions)
	PutNetworkInterfaceOptions(value *VerifiedaccessEndpointNetworkInterfaceOptions)
	PutSseSpecification(value *VerifiedaccessEndpointSseSpecification)
	PutTimeouts(value *VerifiedaccessEndpointTimeouts)
	ResetDescription()
	ResetId()
	ResetLoadBalancerOptions()
	ResetNetworkInterfaceOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyDocument()
	ResetSecurityGroupIds()
	ResetSseSpecification()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
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

// The jsii proxy struct for VerifiedaccessEndpoint
type jsiiProxy_VerifiedaccessEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VerifiedaccessEndpoint) ApplicationDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) ApplicationDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) AttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) AttachmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) DeviceValidationDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceValidationDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) DomainCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) DomainCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) EndpointDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) EndpointDomainPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointDomainPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) EndpointDomainPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointDomainPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) LoadBalancerOptions() VerifiedaccessEndpointLoadBalancerOptionsOutputReference {
	var returns VerifiedaccessEndpointLoadBalancerOptionsOutputReference
	_jsii_.Get(
		j,
		"loadBalancerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) LoadBalancerOptionsInput() *VerifiedaccessEndpointLoadBalancerOptions {
	var returns *VerifiedaccessEndpointLoadBalancerOptions
	_jsii_.Get(
		j,
		"loadBalancerOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) NetworkInterfaceOptions() VerifiedaccessEndpointNetworkInterfaceOptionsOutputReference {
	var returns VerifiedaccessEndpointNetworkInterfaceOptionsOutputReference
	_jsii_.Get(
		j,
		"networkInterfaceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) NetworkInterfaceOptionsInput() *VerifiedaccessEndpointNetworkInterfaceOptions {
	var returns *VerifiedaccessEndpointNetworkInterfaceOptions
	_jsii_.Get(
		j,
		"networkInterfaceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) PolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) PolicyDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) SseSpecification() VerifiedaccessEndpointSseSpecificationOutputReference {
	var returns VerifiedaccessEndpointSseSpecificationOutputReference
	_jsii_.Get(
		j,
		"sseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) SseSpecificationInput() *VerifiedaccessEndpointSseSpecification {
	var returns *VerifiedaccessEndpointSseSpecification
	_jsii_.Get(
		j,
		"sseSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) Timeouts() VerifiedaccessEndpointTimeoutsOutputReference {
	var returns VerifiedaccessEndpointTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) VerifiedAccessGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedAccessGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) VerifiedAccessGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedAccessGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpoint) VerifiedAccessInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedAccessInstanceId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/verifiedaccess_endpoint aws_verifiedaccess_endpoint} Resource.
func NewVerifiedaccessEndpoint(scope constructs.Construct, id *string, config *VerifiedaccessEndpointConfig) VerifiedaccessEndpoint {
	_init_.Initialize()

	if err := validateNewVerifiedaccessEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VerifiedaccessEndpoint{}

	_jsii_.Create(
		"@cdktf/provider-aws.verifiedaccessEndpoint.VerifiedaccessEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/verifiedaccess_endpoint aws_verifiedaccess_endpoint} Resource.
func NewVerifiedaccessEndpoint_Override(v VerifiedaccessEndpoint, scope constructs.Construct, id *string, config *VerifiedaccessEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.verifiedaccessEndpoint.VerifiedaccessEndpoint",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetApplicationDomain(val *string) {
	if err := j.validateSetApplicationDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationDomain",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetAttachmentType(val *string) {
	if err := j.validateSetAttachmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachmentType",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetDomainCertificateArn(val *string) {
	if err := j.validateSetDomainCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainCertificateArn",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetEndpointDomainPrefix(val *string) {
	if err := j.validateSetEndpointDomainPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointDomainPrefix",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetEndpointType(val *string) {
	if err := j.validateSetEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetPolicyDocument(val *string) {
	if err := j.validateSetPolicyDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDocument",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpoint)SetVerifiedAccessGroupId(val *string) {
	if err := j.validateSetVerifiedAccessGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifiedAccessGroupId",
		val,
	)
}

// Generates CDKTF code for importing a VerifiedaccessEndpoint resource upon running "cdktf plan <stack-name>".
func VerifiedaccessEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVerifiedaccessEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.verifiedaccessEndpoint.VerifiedaccessEndpoint",
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
func VerifiedaccessEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVerifiedaccessEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.verifiedaccessEndpoint.VerifiedaccessEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VerifiedaccessEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVerifiedaccessEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.verifiedaccessEndpoint.VerifiedaccessEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VerifiedaccessEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVerifiedaccessEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.verifiedaccessEndpoint.VerifiedaccessEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VerifiedaccessEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.verifiedaccessEndpoint.VerifiedaccessEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) PutLoadBalancerOptions(value *VerifiedaccessEndpointLoadBalancerOptions) {
	if err := v.validatePutLoadBalancerOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putLoadBalancerOptions",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) PutNetworkInterfaceOptions(value *VerifiedaccessEndpointNetworkInterfaceOptions) {
	if err := v.validatePutNetworkInterfaceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putNetworkInterfaceOptions",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) PutSseSpecification(value *VerifiedaccessEndpointSseSpecification) {
	if err := v.validatePutSseSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putSseSpecification",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) PutTimeouts(value *VerifiedaccessEndpointTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ResetLoadBalancerOptions() {
	_jsii_.InvokeVoid(
		v,
		"resetLoadBalancerOptions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ResetNetworkInterfaceOptions() {
	_jsii_.InvokeVoid(
		v,
		"resetNetworkInterfaceOptions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ResetPolicyDocument() {
	_jsii_.InvokeVoid(
		v,
		"resetPolicyDocument",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		v,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ResetSseSpecification() {
	_jsii_.InvokeVoid(
		v,
		"resetSseSpecification",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ResetTagsAll() {
	_jsii_.InvokeVoid(
		v,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

