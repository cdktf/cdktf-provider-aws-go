// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/pipespipe/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference interface {
	cdktf.ComplexObject
	BasicAuth() *string
	SetBasicAuth(val *string)
	BasicAuthInput() *string
	ClientCertificateTlsAuth() *string
	SetClientCertificateTlsAuth(val *string)
	ClientCertificateTlsAuthInput() *string
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
	InternalValue() *PipesPipeSourceParametersSelfManagedKafkaParametersCredentials
	SetInternalValue(val *PipesPipeSourceParametersSelfManagedKafkaParametersCredentials)
	SaslScram256Auth() *string
	SetSaslScram256Auth(val *string)
	SaslScram256AuthInput() *string
	SaslScram512Auth() *string
	SetSaslScram512Auth(val *string)
	SaslScram512AuthInput() *string
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
	ResetClientCertificateTlsAuth()
	ResetSaslScram256Auth()
	ResetSaslScram512Auth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference
type jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) BasicAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) BasicAuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) ClientCertificateTlsAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateTlsAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) ClientCertificateTlsAuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateTlsAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) InternalValue() *PipesPipeSourceParametersSelfManagedKafkaParametersCredentials {
	var returns *PipesPipeSourceParametersSelfManagedKafkaParametersCredentials
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) SaslScram256Auth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslScram256Auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) SaslScram256AuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslScram256AuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) SaslScram512Auth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslScram512Auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) SaslScram512AuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslScram512AuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference {
	_init_.Initialize()

	if err := validateNewPipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference_Override(p PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference)SetBasicAuth(val *string) {
	if err := j.validateSetBasicAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"basicAuth",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference)SetClientCertificateTlsAuth(val *string) {
	if err := j.validateSetClientCertificateTlsAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateTlsAuth",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference)SetInternalValue(val *PipesPipeSourceParametersSelfManagedKafkaParametersCredentials) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference)SetSaslScram256Auth(val *string) {
	if err := j.validateSetSaslScram256AuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslScram256Auth",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference)SetSaslScram512Auth(val *string) {
	if err := j.validateSetSaslScram512AuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslScram512Auth",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) ResetClientCertificateTlsAuth() {
	_jsii_.InvokeVoid(
		p,
		"resetClientCertificateTlsAuth",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) ResetSaslScram256Auth() {
	_jsii_.InvokeVoid(
		p,
		"resetSaslScram256Auth",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) ResetSaslScram512Auth() {
	_jsii_.InvokeVoid(
		p,
		"resetSaslScram512Auth",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersSelfManagedKafkaParametersCredentialsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

