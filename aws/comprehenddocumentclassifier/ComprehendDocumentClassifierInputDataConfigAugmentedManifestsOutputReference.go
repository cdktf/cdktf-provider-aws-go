// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package comprehenddocumentclassifier

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/comprehenddocumentclassifier/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference interface {
	cdktf.ComplexObject
	AnnotationDataS3Uri() *string
	SetAnnotationDataS3Uri(val *string)
	AnnotationDataS3UriInput() *string
	AttributeNames() *[]*string
	SetAttributeNames(val *[]*string)
	AttributeNamesInput() *[]*string
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
	DocumentType() *string
	SetDocumentType(val *string)
	DocumentTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3Uri() *string
	SetS3Uri(val *string)
	S3UriInput() *string
	SourceDocumentsS3Uri() *string
	SetSourceDocumentsS3Uri(val *string)
	SourceDocumentsS3UriInput() *string
	Split() *string
	SetSplit(val *string)
	SplitInput() *string
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
	ResetAnnotationDataS3Uri()
	ResetDocumentType()
	ResetSourceDocumentsS3Uri()
	ResetSplit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference
type jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) AnnotationDataS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationDataS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) AnnotationDataS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationDataS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) AttributeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) AttributeNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) DocumentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) DocumentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) S3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) S3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) SourceDocumentsS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDocumentsS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) SourceDocumentsS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDocumentsS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) Split() *string {
	var returns *string
	_jsii_.Get(
		j,
		"split",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) SplitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference {
	_init_.Initialize()

	if err := validateNewComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.comprehendDocumentClassifier.ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference_Override(c ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.comprehendDocumentClassifier.ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference)SetAnnotationDataS3Uri(val *string) {
	if err := j.validateSetAnnotationDataS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotationDataS3Uri",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference)SetAttributeNames(val *[]*string) {
	if err := j.validateSetAttributeNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeNames",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference)SetDocumentType(val *string) {
	if err := j.validateSetDocumentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentType",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference)SetS3Uri(val *string) {
	if err := j.validateSetS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Uri",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference)SetSourceDocumentsS3Uri(val *string) {
	if err := j.validateSetSourceDocumentsS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDocumentsS3Uri",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference)SetSplit(val *string) {
	if err := j.validateSetSplitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"split",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) ResetAnnotationDataS3Uri() {
	_jsii_.InvokeVoid(
		c,
		"resetAnnotationDataS3Uri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) ResetDocumentType() {
	_jsii_.InvokeVoid(
		c,
		"resetDocumentType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) ResetSourceDocumentsS3Uri() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceDocumentsS3Uri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) ResetSplit() {
	_jsii_.InvokeVoid(
		c,
		"resetSplit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigAugmentedManifestsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

