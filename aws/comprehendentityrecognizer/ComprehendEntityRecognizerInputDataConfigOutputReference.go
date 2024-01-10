// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package comprehendentityrecognizer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/comprehendentityrecognizer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComprehendEntityRecognizerInputDataConfigOutputReference interface {
	cdktf.ComplexObject
	Annotations() ComprehendEntityRecognizerInputDataConfigAnnotationsOutputReference
	AnnotationsInput() *ComprehendEntityRecognizerInputDataConfigAnnotations
	AugmentedManifests() ComprehendEntityRecognizerInputDataConfigAugmentedManifestsList
	AugmentedManifestsInput() interface{}
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
	DataFormat() *string
	SetDataFormat(val *string)
	DataFormatInput() *string
	Documents() ComprehendEntityRecognizerInputDataConfigDocumentsOutputReference
	DocumentsInput() *ComprehendEntityRecognizerInputDataConfigDocuments
	EntityList() ComprehendEntityRecognizerInputDataConfigEntityListStructOutputReference
	EntityListInput() *ComprehendEntityRecognizerInputDataConfigEntityListStruct
	EntityTypes() ComprehendEntityRecognizerInputDataConfigEntityTypesList
	EntityTypesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ComprehendEntityRecognizerInputDataConfig
	SetInternalValue(val *ComprehendEntityRecognizerInputDataConfig)
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
	PutAnnotations(value *ComprehendEntityRecognizerInputDataConfigAnnotations)
	PutAugmentedManifests(value interface{})
	PutDocuments(value *ComprehendEntityRecognizerInputDataConfigDocuments)
	PutEntityList(value *ComprehendEntityRecognizerInputDataConfigEntityListStruct)
	PutEntityTypes(value interface{})
	ResetAnnotations()
	ResetAugmentedManifests()
	ResetDataFormat()
	ResetDocuments()
	ResetEntityList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComprehendEntityRecognizerInputDataConfigOutputReference
type jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) Annotations() ComprehendEntityRecognizerInputDataConfigAnnotationsOutputReference {
	var returns ComprehendEntityRecognizerInputDataConfigAnnotationsOutputReference
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) AnnotationsInput() *ComprehendEntityRecognizerInputDataConfigAnnotations {
	var returns *ComprehendEntityRecognizerInputDataConfigAnnotations
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) AugmentedManifests() ComprehendEntityRecognizerInputDataConfigAugmentedManifestsList {
	var returns ComprehendEntityRecognizerInputDataConfigAugmentedManifestsList
	_jsii_.Get(
		j,
		"augmentedManifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) AugmentedManifestsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"augmentedManifestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) DataFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) Documents() ComprehendEntityRecognizerInputDataConfigDocumentsOutputReference {
	var returns ComprehendEntityRecognizerInputDataConfigDocumentsOutputReference
	_jsii_.Get(
		j,
		"documents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) DocumentsInput() *ComprehendEntityRecognizerInputDataConfigDocuments {
	var returns *ComprehendEntityRecognizerInputDataConfigDocuments
	_jsii_.Get(
		j,
		"documentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) EntityList() ComprehendEntityRecognizerInputDataConfigEntityListStructOutputReference {
	var returns ComprehendEntityRecognizerInputDataConfigEntityListStructOutputReference
	_jsii_.Get(
		j,
		"entityList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) EntityListInput() *ComprehendEntityRecognizerInputDataConfigEntityListStruct {
	var returns *ComprehendEntityRecognizerInputDataConfigEntityListStruct
	_jsii_.Get(
		j,
		"entityListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) EntityTypes() ComprehendEntityRecognizerInputDataConfigEntityTypesList {
	var returns ComprehendEntityRecognizerInputDataConfigEntityTypesList
	_jsii_.Get(
		j,
		"entityTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) EntityTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entityTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) InternalValue() *ComprehendEntityRecognizerInputDataConfig {
	var returns *ComprehendEntityRecognizerInputDataConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComprehendEntityRecognizerInputDataConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComprehendEntityRecognizerInputDataConfigOutputReference {
	_init_.Initialize()

	if err := validateNewComprehendEntityRecognizerInputDataConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.comprehendEntityRecognizer.ComprehendEntityRecognizerInputDataConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComprehendEntityRecognizerInputDataConfigOutputReference_Override(c ComprehendEntityRecognizerInputDataConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.comprehendEntityRecognizer.ComprehendEntityRecognizerInputDataConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference)SetDataFormat(val *string) {
	if err := j.validateSetDataFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFormat",
		val,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference)SetInternalValue(val *ComprehendEntityRecognizerInputDataConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) PutAnnotations(value *ComprehendEntityRecognizerInputDataConfigAnnotations) {
	if err := c.validatePutAnnotationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAnnotations",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) PutAugmentedManifests(value interface{}) {
	if err := c.validatePutAugmentedManifestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAugmentedManifests",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) PutDocuments(value *ComprehendEntityRecognizerInputDataConfigDocuments) {
	if err := c.validatePutDocumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDocuments",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) PutEntityList(value *ComprehendEntityRecognizerInputDataConfigEntityListStruct) {
	if err := c.validatePutEntityListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEntityList",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) PutEntityTypes(value interface{}) {
	if err := c.validatePutEntityTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEntityTypes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) ResetAnnotations() {
	_jsii_.InvokeVoid(
		c,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) ResetAugmentedManifests() {
	_jsii_.InvokeVoid(
		c,
		"resetAugmentedManifests",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) ResetDataFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetDataFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) ResetDocuments() {
	_jsii_.InvokeVoid(
		c,
		"resetDocuments",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) ResetEntityList() {
	_jsii_.InvokeVoid(
		c,
		"resetEntityList",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

