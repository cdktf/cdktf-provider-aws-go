// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsearchdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/elasticsearchdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference interface {
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
	InternalValue() *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions
	SetInternalValue(val *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions)
	MasterUserArn() *string
	SetMasterUserArn(val *string)
	MasterUserArnInput() *string
	MasterUserName() *string
	SetMasterUserName(val *string)
	MasterUserNameInput() *string
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	MasterUserPasswordInput() *string
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
	ResetMasterUserArn()
	ResetMasterUserName()
	ResetMasterUserPassword()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference
type jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) InternalValue() *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions {
	var returns *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference_Override(e ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetInternalValue(val *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetMasterUserArn(val *string) {
	if err := j.validateSetMasterUserArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserArn",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetMasterUserName(val *string) {
	if err := j.validateSetMasterUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserName",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetMasterUserPassword(val *string) {
	if err := j.validateSetMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ResetMasterUserArn() {
	_jsii_.InvokeVoid(
		e,
		"resetMasterUserArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ResetMasterUserName() {
	_jsii_.InvokeVoid(
		e,
		"resetMasterUserName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ResetMasterUserPassword() {
	_jsii_.InvokeVoid(
		e,
		"resetMasterUserPassword",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

