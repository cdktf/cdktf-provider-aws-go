// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluesecurityconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/gluesecurityconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference interface {
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
	InternalValue() *GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryption
	SetInternalValue(val *GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryption)
	JobBookmarksEncryptionMode() *string
	SetJobBookmarksEncryptionMode(val *string)
	JobBookmarksEncryptionModeInput() *string
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetJobBookmarksEncryptionMode()
	ResetKmsKeyArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference
type jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) InternalValue() *GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryption {
	var returns *GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryption
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) JobBookmarksEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobBookmarksEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) JobBookmarksEncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobBookmarksEncryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference {
	_init_.Initialize()

	if err := validateNewGlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.glueSecurityConfiguration.GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference_Override(g GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.glueSecurityConfiguration.GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference)SetInternalValue(val *GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryption) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference)SetJobBookmarksEncryptionMode(val *string) {
	if err := j.validateSetJobBookmarksEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobBookmarksEncryptionMode",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) ResetJobBookmarksEncryptionMode() {
	_jsii_.InvokeVoid(
		g,
		"resetJobBookmarksEncryptionMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

