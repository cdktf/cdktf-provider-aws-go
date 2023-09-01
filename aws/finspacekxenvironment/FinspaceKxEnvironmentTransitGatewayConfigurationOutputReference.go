// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package finspacekxenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/finspacekxenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference interface {
	cdktf.ComplexObject
	AttachmentNetworkAclConfiguration() FinspaceKxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationList
	AttachmentNetworkAclConfigurationInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *FinspaceKxEnvironmentTransitGatewayConfiguration
	SetInternalValue(val *FinspaceKxEnvironmentTransitGatewayConfiguration)
	RoutableCidrSpace() *string
	SetRoutableCidrSpace(val *string)
	RoutableCidrSpaceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransitGatewayId() *string
	SetTransitGatewayId(val *string)
	TransitGatewayIdInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAttachmentNetworkAclConfiguration(value interface{})
	ResetAttachmentNetworkAclConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference
type jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) AttachmentNetworkAclConfiguration() FinspaceKxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationList {
	var returns FinspaceKxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationList
	_jsii_.Get(
		j,
		"attachmentNetworkAclConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) AttachmentNetworkAclConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachmentNetworkAclConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) InternalValue() *FinspaceKxEnvironmentTransitGatewayConfiguration {
	var returns *FinspaceKxEnvironmentTransitGatewayConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) RoutableCidrSpace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routableCidrSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) RoutableCidrSpaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routableCidrSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) TransitGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayIdInput",
		&returns,
	)
	return returns
}


func NewFinspaceKxEnvironmentTransitGatewayConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewFinspaceKxEnvironmentTransitGatewayConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.finspaceKxEnvironment.FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFinspaceKxEnvironmentTransitGatewayConfigurationOutputReference_Override(f FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.finspaceKxEnvironment.FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference)SetInternalValue(val *FinspaceKxEnvironmentTransitGatewayConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference)SetRoutableCidrSpace(val *string) {
	if err := j.validateSetRoutableCidrSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routableCidrSpace",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference)SetTransitGatewayId(val *string) {
	if err := j.validateSetTransitGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayId",
		val,
	)
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) PutAttachmentNetworkAclConfiguration(value interface{}) {
	if err := f.validatePutAttachmentNetworkAclConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAttachmentNetworkAclConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) ResetAttachmentNetworkAclConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetAttachmentNetworkAclConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceKxEnvironmentTransitGatewayConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

