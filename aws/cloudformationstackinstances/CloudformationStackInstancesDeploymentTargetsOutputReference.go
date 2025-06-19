// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudformationstackinstances

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/cloudformationstackinstances/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudformationStackInstancesDeploymentTargetsOutputReference interface {
	cdktf.ComplexObject
	AccountFilterType() *string
	SetAccountFilterType(val *string)
	AccountFilterTypeInput() *string
	Accounts() *[]*string
	SetAccounts(val *[]*string)
	AccountsInput() *[]*string
	AccountsUrl() *string
	SetAccountsUrl(val *string)
	AccountsUrlInput() *string
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
	InternalValue() *CloudformationStackInstancesDeploymentTargets
	SetInternalValue(val *CloudformationStackInstancesDeploymentTargets)
	OrganizationalUnitIds() *[]*string
	SetOrganizationalUnitIds(val *[]*string)
	OrganizationalUnitIdsInput() *[]*string
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
	ResetAccountFilterType()
	ResetAccounts()
	ResetAccountsUrl()
	ResetOrganizationalUnitIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudformationStackInstancesDeploymentTargetsOutputReference
type jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) AccountFilterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountFilterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) AccountFilterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountFilterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) Accounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) AccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) AccountsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) AccountsUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountsUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) InternalValue() *CloudformationStackInstancesDeploymentTargets {
	var returns *CloudformationStackInstancesDeploymentTargets
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) OrganizationalUnitIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnitIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) OrganizationalUnitIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnitIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudformationStackInstancesDeploymentTargetsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudformationStackInstancesDeploymentTargetsOutputReference {
	_init_.Initialize()

	if err := validateNewCloudformationStackInstancesDeploymentTargetsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudformationStackInstances.CloudformationStackInstancesDeploymentTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudformationStackInstancesDeploymentTargetsOutputReference_Override(c CloudformationStackInstancesDeploymentTargetsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudformationStackInstances.CloudformationStackInstancesDeploymentTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference)SetAccountFilterType(val *string) {
	if err := j.validateSetAccountFilterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountFilterType",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference)SetAccounts(val *[]*string) {
	if err := j.validateSetAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accounts",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference)SetAccountsUrl(val *string) {
	if err := j.validateSetAccountsUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountsUrl",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference)SetInternalValue(val *CloudformationStackInstancesDeploymentTargets) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference)SetOrganizationalUnitIds(val *[]*string) {
	if err := j.validateSetOrganizationalUnitIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationalUnitIds",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) ResetAccountFilterType() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountFilterType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) ResetAccounts() {
	_jsii_.InvokeVoid(
		c,
		"resetAccounts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) ResetAccountsUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountsUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) ResetOrganizationalUnitIds() {
	_jsii_.InvokeVoid(
		c,
		"resetOrganizationalUnitIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackInstancesDeploymentTargetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

