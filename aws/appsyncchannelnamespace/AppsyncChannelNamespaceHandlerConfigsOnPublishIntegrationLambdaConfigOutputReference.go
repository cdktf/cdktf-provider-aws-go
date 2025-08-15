// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncchannelnamespace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/appsyncchannelnamespace/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	InvokeType() *string
	SetInvokeType(val *string)
	InvokeTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetInvokeType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference
type jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) InvokeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) InvokeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncChannelNamespace.AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference_Override(a AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncChannelNamespace.AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference)SetInvokeType(val *string) {
	if err := j.validateSetInvokeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invokeType",
		val,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) ResetInvokeType() {
	_jsii_.InvokeVoid(
		a,
		"resetInvokeType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

