// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/emrcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrClusterKerberosAttributesOutputReference interface {
	cdktf.ComplexObject
	AdDomainJoinPassword() *string
	SetAdDomainJoinPassword(val *string)
	AdDomainJoinPasswordInput() *string
	AdDomainJoinUser() *string
	SetAdDomainJoinUser(val *string)
	AdDomainJoinUserInput() *string
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
	CrossRealmTrustPrincipalPassword() *string
	SetCrossRealmTrustPrincipalPassword(val *string)
	CrossRealmTrustPrincipalPasswordInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *EmrClusterKerberosAttributes
	SetInternalValue(val *EmrClusterKerberosAttributes)
	KdcAdminPassword() *string
	SetKdcAdminPassword(val *string)
	KdcAdminPasswordInput() *string
	Realm() *string
	SetRealm(val *string)
	RealmInput() *string
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
	ResetAdDomainJoinPassword()
	ResetAdDomainJoinUser()
	ResetCrossRealmTrustPrincipalPassword()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmrClusterKerberosAttributesOutputReference
type jsiiProxy_EmrClusterKerberosAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) AdDomainJoinPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) AdDomainJoinPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) AdDomainJoinUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) AdDomainJoinUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) CrossRealmTrustPrincipalPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustPrincipalPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) CrossRealmTrustPrincipalPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustPrincipalPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) InternalValue() *EmrClusterKerberosAttributes {
	var returns *EmrClusterKerberosAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) KdcAdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) KdcAdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcAdminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) Realm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) RealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEmrClusterKerberosAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EmrClusterKerberosAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewEmrClusterKerberosAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrClusterKerberosAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrCluster.EmrClusterKerberosAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEmrClusterKerberosAttributesOutputReference_Override(e EmrClusterKerberosAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrCluster.EmrClusterKerberosAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference)SetAdDomainJoinPassword(val *string) {
	if err := j.validateSetAdDomainJoinPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adDomainJoinPassword",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference)SetAdDomainJoinUser(val *string) {
	if err := j.validateSetAdDomainJoinUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adDomainJoinUser",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference)SetCrossRealmTrustPrincipalPassword(val *string) {
	if err := j.validateSetCrossRealmTrustPrincipalPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustPrincipalPassword",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference)SetInternalValue(val *EmrClusterKerberosAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference)SetKdcAdminPassword(val *string) {
	if err := j.validateSetKdcAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kdcAdminPassword",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference)SetRealm(val *string) {
	if err := j.validateSetRealmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realm",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) ResetAdDomainJoinPassword() {
	_jsii_.InvokeVoid(
		e,
		"resetAdDomainJoinPassword",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) ResetAdDomainJoinUser() {
	_jsii_.InvokeVoid(
		e,
		"resetAdDomainJoinUser",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) ResetCrossRealmTrustPrincipalPassword() {
	_jsii_.InvokeVoid(
		e,
		"resetCrossRealmTrustPrincipalPassword",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

