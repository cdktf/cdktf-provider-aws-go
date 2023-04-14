package comprehendentityrecognizer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v13/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v13/comprehendentityrecognizer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference interface {
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

// The jsii proxy struct for ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference
type jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) AnnotationDataS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationDataS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) AnnotationDataS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationDataS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) AttributeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) AttributeNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) DocumentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) DocumentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) S3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) S3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) SourceDocumentsS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDocumentsS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) SourceDocumentsS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDocumentsS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) Split() *string {
	var returns *string
	_jsii_.Get(
		j,
		"split",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) SplitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference {
	_init_.Initialize()

	if err := validateNewComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.comprehendEntityRecognizer.ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference_Override(c ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.comprehendEntityRecognizer.ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference)SetAnnotationDataS3Uri(val *string) {
	if err := j.validateSetAnnotationDataS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotationDataS3Uri",
		val,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference)SetAttributeNames(val *[]*string) {
	if err := j.validateSetAttributeNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeNames",
		val,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference)SetDocumentType(val *string) {
	if err := j.validateSetDocumentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentType",
		val,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference)SetS3Uri(val *string) {
	if err := j.validateSetS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Uri",
		val,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference)SetSourceDocumentsS3Uri(val *string) {
	if err := j.validateSetSourceDocumentsS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDocumentsS3Uri",
		val,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference)SetSplit(val *string) {
	if err := j.validateSetSplitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"split",
		val,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) ResetAnnotationDataS3Uri() {
	_jsii_.InvokeVoid(
		c,
		"resetAnnotationDataS3Uri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) ResetDocumentType() {
	_jsii_.InvokeVoid(
		c,
		"resetDocumentType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) ResetSourceDocumentsS3Uri() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceDocumentsS3Uri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) ResetSplit() {
	_jsii_.InvokeVoid(
		c,
		"resetSplit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComprehendEntityRecognizerInputDataConfigAugmentedManifestsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

