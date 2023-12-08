// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customerprofilesdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/customerprofilesdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference interface {
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
	InternalValue() *CustomerprofilesDomainRuleBasedMatchingExportingConfigS3Exporting
	SetInternalValue(val *CustomerprofilesDomainRuleBasedMatchingExportingConfigS3Exporting)
	S3BucketName() *string
	SetS3BucketName(val *string)
	S3BucketNameInput() *string
	S3KeyName() *string
	SetS3KeyName(val *string)
	S3KeyNameInput() *string
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
	ResetS3KeyName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference
type jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) InternalValue() *CustomerprofilesDomainRuleBasedMatchingExportingConfigS3Exporting {
	var returns *CustomerprofilesDomainRuleBasedMatchingExportingConfigS3Exporting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) S3BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) S3KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) S3KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference {
	_init_.Initialize()

	if err := validateNewCustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.customerprofilesDomain.CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference_Override(c CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.customerprofilesDomain.CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference)SetInternalValue(val *CustomerprofilesDomainRuleBasedMatchingExportingConfigS3Exporting) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference)SetS3BucketName(val *string) {
	if err := j.validateSetS3BucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference)SetS3KeyName(val *string) {
	if err := j.validateSetS3KeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3KeyName",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) ResetS3KeyName() {
	_jsii_.InvokeVoid(
		c,
		"resetS3KeyName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingExportingConfigS3ExportingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

