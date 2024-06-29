// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcendpointservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/vpcendpointservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/vpc_endpoint_service aws_vpc_endpoint_service}.
type VpcEndpointService interface {
	cdktf.TerraformResource
	AcceptanceRequired() interface{}
	SetAcceptanceRequired(val interface{})
	AcceptanceRequiredInput() interface{}
	AllowedPrincipals() *[]*string
	SetAllowedPrincipals(val *[]*string)
	AllowedPrincipalsInput() *[]*string
	Arn() *string
	AvailabilityZones() *[]*string
	BaseEndpointDnsNames() *[]*string
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayLoadBalancerArns() *[]*string
	SetGatewayLoadBalancerArns(val *[]*string)
	GatewayLoadBalancerArnsInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManagesVpcEndpoints() cdktf.IResolvable
	NetworkLoadBalancerArns() *[]*string
	SetNetworkLoadBalancerArns(val *[]*string)
	NetworkLoadBalancerArnsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	PrivateDnsName() *string
	SetPrivateDnsName(val *string)
	PrivateDnsNameConfiguration() VpcEndpointServicePrivateDnsNameConfigurationList
	PrivateDnsNameInput() *string
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
	ServiceName() *string
	ServiceType() *string
	State() *string
	SupportedIpAddressTypes() *[]*string
	SetSupportedIpAddressTypes(val *[]*string)
	SupportedIpAddressTypesInput() *[]*string
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
	Timeouts() VpcEndpointServiceTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutTimeouts(value *VpcEndpointServiceTimeouts)
	ResetAllowedPrincipals()
	ResetGatewayLoadBalancerArns()
	ResetId()
	ResetNetworkLoadBalancerArns()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateDnsName()
	ResetSupportedIpAddressTypes()
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

// The jsii proxy struct for VpcEndpointService
type jsiiProxy_VpcEndpointService struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VpcEndpointService) AcceptanceRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptanceRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) AcceptanceRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptanceRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) AllowedPrincipals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) AllowedPrincipalsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedPrincipalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) BaseEndpointDnsNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"baseEndpointDnsNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) GatewayLoadBalancerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gatewayLoadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) GatewayLoadBalancerArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gatewayLoadBalancerArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) ManagesVpcEndpoints() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"managesVpcEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) NetworkLoadBalancerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkLoadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) NetworkLoadBalancerArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkLoadBalancerArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) PrivateDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) PrivateDnsNameConfiguration() VpcEndpointServicePrivateDnsNameConfigurationList {
	var returns VpcEndpointServicePrivateDnsNameConfigurationList
	_jsii_.Get(
		j,
		"privateDnsNameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) PrivateDnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) ServiceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) SupportedIpAddressTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedIpAddressTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) SupportedIpAddressTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedIpAddressTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) Timeouts() VpcEndpointServiceTimeoutsOutputReference {
	var returns VpcEndpointServiceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointService) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/vpc_endpoint_service aws_vpc_endpoint_service} Resource.
func NewVpcEndpointService(scope constructs.Construct, id *string, config *VpcEndpointServiceConfig) VpcEndpointService {
	_init_.Initialize()

	if err := validateNewVpcEndpointServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpcEndpointService{}

	_jsii_.Create(
		"@cdktf/provider-aws.vpcEndpointService.VpcEndpointService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/vpc_endpoint_service aws_vpc_endpoint_service} Resource.
func NewVpcEndpointService_Override(v VpcEndpointService, scope constructs.Construct, id *string, config *VpcEndpointServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.vpcEndpointService.VpcEndpointService",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetAcceptanceRequired(val interface{}) {
	if err := j.validateSetAcceptanceRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptanceRequired",
		val,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetAllowedPrincipals(val *[]*string) {
	if err := j.validateSetAllowedPrincipalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedPrincipals",
		val,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetGatewayLoadBalancerArns(val *[]*string) {
	if err := j.validateSetGatewayLoadBalancerArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayLoadBalancerArns",
		val,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetNetworkLoadBalancerArns(val *[]*string) {
	if err := j.validateSetNetworkLoadBalancerArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkLoadBalancerArns",
		val,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetPrivateDnsName(val *string) {
	if err := j.validateSetPrivateDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsName",
		val,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetSupportedIpAddressTypes(val *[]*string) {
	if err := j.validateSetSupportedIpAddressTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedIpAddressTypes",
		val,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VpcEndpointService)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a VpcEndpointService resource upon running "cdktf plan <stack-name>".
func VpcEndpointService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVpcEndpointService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcEndpointService.VpcEndpointService",
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
func VpcEndpointService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcEndpointService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcEndpointService.VpcEndpointService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcEndpointService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcEndpointService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcEndpointService.VpcEndpointService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcEndpointService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcEndpointService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcEndpointService.VpcEndpointService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpcEndpointService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.vpcEndpointService.VpcEndpointService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpcEndpointService) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VpcEndpointService) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpcEndpointService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpcEndpointService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpcEndpointService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpcEndpointService) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpcEndpointService) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpcEndpointService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpcEndpointService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpcEndpointService) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpcEndpointService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpcEndpointService) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEndpointService) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VpcEndpointService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpcEndpointService) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpcEndpointService) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VpcEndpointService) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpcEndpointService) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpcEndpointService) PutTimeouts(value *VpcEndpointServiceTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpcEndpointService) ResetAllowedPrincipals() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowedPrincipals",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpointService) ResetGatewayLoadBalancerArns() {
	_jsii_.InvokeVoid(
		v,
		"resetGatewayLoadBalancerArns",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpointService) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpointService) ResetNetworkLoadBalancerArns() {
	_jsii_.InvokeVoid(
		v,
		"resetNetworkLoadBalancerArns",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpointService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpointService) ResetPrivateDnsName() {
	_jsii_.InvokeVoid(
		v,
		"resetPrivateDnsName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpointService) ResetSupportedIpAddressTypes() {
	_jsii_.InvokeVoid(
		v,
		"resetSupportedIpAddressTypes",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpointService) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpointService) ResetTagsAll() {
	_jsii_.InvokeVoid(
		v,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpointService) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpointService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEndpointService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEndpointService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEndpointService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEndpointService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEndpointService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

