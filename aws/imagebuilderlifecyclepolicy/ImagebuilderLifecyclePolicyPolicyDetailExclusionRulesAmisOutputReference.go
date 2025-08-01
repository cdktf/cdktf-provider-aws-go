// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderlifecyclepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/imagebuilderlifecyclepolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference interface {
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
	IsPublic() interface{}
	SetIsPublic(val interface{})
	IsPublicInput() interface{}
	LastLaunched() ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisLastLaunchedList
	LastLaunchedInput() interface{}
	Regions() *[]*string
	SetRegions(val *[]*string)
	RegionsInput() *[]*string
	SharedAccounts() *[]*string
	SetSharedAccounts(val *[]*string)
	SharedAccountsInput() *[]*string
	TagMap() *map[string]*string
	SetTagMap(val *map[string]*string)
	TagMapInput() *map[string]*string
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
	PutLastLaunched(value interface{})
	ResetIsPublic()
	ResetLastLaunched()
	ResetRegions()
	ResetSharedAccounts()
	ResetTagMap()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference
type jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) IsPublic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) IsPublicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPublicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) LastLaunched() ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisLastLaunchedList {
	var returns ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisLastLaunchedList
	_jsii_.Get(
		j,
		"lastLaunched",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) LastLaunchedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastLaunchedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) SharedAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) SharedAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) TagMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) TagMapInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference {
	_init_.Initialize()

	if err := validateNewImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.imagebuilderLifecyclePolicy.ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference_Override(i ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.imagebuilderLifecyclePolicy.ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference)SetIsPublic(val interface{}) {
	if err := j.validateSetIsPublicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPublic",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference)SetSharedAccounts(val *[]*string) {
	if err := j.validateSetSharedAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedAccounts",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference)SetTagMap(val *map[string]*string) {
	if err := j.validateSetTagMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagMap",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) PutLastLaunched(value interface{}) {
	if err := i.validatePutLastLaunchedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLastLaunched",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) ResetIsPublic() {
	_jsii_.InvokeVoid(
		i,
		"resetIsPublic",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) ResetLastLaunched() {
	_jsii_.InvokeVoid(
		i,
		"resetLastLaunched",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		i,
		"resetRegions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) ResetSharedAccounts() {
	_jsii_.InvokeVoid(
		i,
		"resetSharedAccounts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) ResetTagMap() {
	_jsii_.InvokeVoid(
		i,
		"resetTagMap",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesAmisOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

