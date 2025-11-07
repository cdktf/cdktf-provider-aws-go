// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsearchdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/elasticsearchdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticsearchDomainDomainEndpointOptionsOutputReference interface {
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
	CustomEndpoint() *string
	SetCustomEndpoint(val *string)
	CustomEndpointCertificateArn() *string
	SetCustomEndpointCertificateArn(val *string)
	CustomEndpointCertificateArnInput() *string
	CustomEndpointEnabled() interface{}
	SetCustomEndpointEnabled(val interface{})
	CustomEndpointEnabledInput() interface{}
	CustomEndpointInput() *string
	EnforceHttps() interface{}
	SetEnforceHttps(val interface{})
	EnforceHttpsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ElasticsearchDomainDomainEndpointOptions
	SetInternalValue(val *ElasticsearchDomainDomainEndpointOptions)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsSecurityPolicy() *string
	SetTlsSecurityPolicy(val *string)
	TlsSecurityPolicyInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetCustomEndpoint()
	ResetCustomEndpointCertificateArn()
	ResetCustomEndpointEnabled()
	ResetEnforceHttps()
	ResetTlsSecurityPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElasticsearchDomainDomainEndpointOptionsOutputReference
type jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) CustomEndpointCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEndpointCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) CustomEndpointCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEndpointCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) CustomEndpointEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customEndpointEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) CustomEndpointEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customEndpointEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) EnforceHttps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) EnforceHttpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceHttpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) InternalValue() *ElasticsearchDomainDomainEndpointOptions {
	var returns *ElasticsearchDomainDomainEndpointOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) TlsSecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) TlsSecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsSecurityPolicyInput",
		&returns,
	)
	return returns
}


func NewElasticsearchDomainDomainEndpointOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElasticsearchDomainDomainEndpointOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewElasticsearchDomainDomainEndpointOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainDomainEndpointOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElasticsearchDomainDomainEndpointOptionsOutputReference_Override(e ElasticsearchDomainDomainEndpointOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainDomainEndpointOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference)SetCustomEndpoint(val *string) {
	if err := j.validateSetCustomEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customEndpoint",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference)SetCustomEndpointCertificateArn(val *string) {
	if err := j.validateSetCustomEndpointCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customEndpointCertificateArn",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference)SetCustomEndpointEnabled(val interface{}) {
	if err := j.validateSetCustomEndpointEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customEndpointEnabled",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference)SetEnforceHttps(val interface{}) {
	if err := j.validateSetEnforceHttpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceHttps",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference)SetInternalValue(val *ElasticsearchDomainDomainEndpointOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference)SetTlsSecurityPolicy(val *string) {
	if err := j.validateSetTlsSecurityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsSecurityPolicy",
		val,
	)
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) ResetCustomEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) ResetCustomEndpointCertificateArn() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomEndpointCertificateArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) ResetCustomEndpointEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomEndpointEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) ResetEnforceHttps() {
	_jsii_.InvokeVoid(
		e,
		"resetEnforceHttps",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) ResetTlsSecurityPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetTlsSecurityPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

