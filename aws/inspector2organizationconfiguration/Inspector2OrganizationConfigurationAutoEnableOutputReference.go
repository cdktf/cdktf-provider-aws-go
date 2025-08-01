// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inspector2organizationconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/inspector2organizationconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Inspector2OrganizationConfigurationAutoEnableOutputReference interface {
	cdktf.ComplexObject
	CodeRepository() interface{}
	SetCodeRepository(val interface{})
	CodeRepositoryInput() interface{}
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
	Ec2() interface{}
	SetEc2(val interface{})
	Ec2Input() interface{}
	Ecr() interface{}
	SetEcr(val interface{})
	EcrInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *Inspector2OrganizationConfigurationAutoEnable
	SetInternalValue(val *Inspector2OrganizationConfigurationAutoEnable)
	Lambda() interface{}
	SetLambda(val interface{})
	LambdaCode() interface{}
	SetLambdaCode(val interface{})
	LambdaCodeInput() interface{}
	LambdaInput() interface{}
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
	ResetCodeRepository()
	ResetLambda()
	ResetLambdaCode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Inspector2OrganizationConfigurationAutoEnableOutputReference
type jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) CodeRepository() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) CodeRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) Ec2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) Ec2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) Ecr() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) EcrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) InternalValue() *Inspector2OrganizationConfigurationAutoEnable {
	var returns *Inspector2OrganizationConfigurationAutoEnable
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) Lambda() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) LambdaCode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) LambdaCodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewInspector2OrganizationConfigurationAutoEnableOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Inspector2OrganizationConfigurationAutoEnableOutputReference {
	_init_.Initialize()

	if err := validateNewInspector2OrganizationConfigurationAutoEnableOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.inspector2OrganizationConfiguration.Inspector2OrganizationConfigurationAutoEnableOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewInspector2OrganizationConfigurationAutoEnableOutputReference_Override(i Inspector2OrganizationConfigurationAutoEnableOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.inspector2OrganizationConfiguration.Inspector2OrganizationConfigurationAutoEnableOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference)SetCodeRepository(val interface{}) {
	if err := j.validateSetCodeRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeRepository",
		val,
	)
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference)SetEc2(val interface{}) {
	if err := j.validateSetEc2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2",
		val,
	)
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference)SetEcr(val interface{}) {
	if err := j.validateSetEcrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecr",
		val,
	)
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference)SetInternalValue(val *Inspector2OrganizationConfigurationAutoEnable) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference)SetLambda(val interface{}) {
	if err := j.validateSetLambdaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambda",
		val,
	)
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference)SetLambdaCode(val interface{}) {
	if err := j.validateSetLambdaCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaCode",
		val,
	)
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) ResetCodeRepository() {
	_jsii_.InvokeVoid(
		i,
		"resetCodeRepository",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		i,
		"resetLambda",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) ResetLambdaCode() {
	_jsii_.InvokeVoid(
		i,
		"resetLambdaCode",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_Inspector2OrganizationConfigurationAutoEnableOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

