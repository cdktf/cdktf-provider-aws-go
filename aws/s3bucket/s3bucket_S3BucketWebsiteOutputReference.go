package s3bucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/s3bucket/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3BucketWebsiteOutputReference interface {
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
	ErrorDocument() *string
	SetErrorDocument(val *string)
	ErrorDocumentInput() *string
	// Experimental.
	Fqn() *string
	IndexDocument() *string
	SetIndexDocument(val *string)
	IndexDocumentInput() *string
	InternalValue() *S3BucketWebsite
	SetInternalValue(val *S3BucketWebsite)
	RedirectAllRequestsTo() *string
	SetRedirectAllRequestsTo(val *string)
	RedirectAllRequestsToInput() *string
	RoutingRules() *string
	SetRoutingRules(val *string)
	RoutingRulesInput() *string
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
	ResetErrorDocument()
	ResetIndexDocument()
	ResetRedirectAllRequestsTo()
	ResetRoutingRules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3BucketWebsiteOutputReference
type jsiiProxy_S3BucketWebsiteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference) ErrorDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference) ErrorDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference) IndexDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference) IndexDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference) InternalValue() *S3BucketWebsite {
	var returns *S3BucketWebsite
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference) RedirectAllRequestsTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectAllRequestsTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference) RedirectAllRequestsToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectAllRequestsToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference) RoutingRules() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference) RoutingRulesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3BucketWebsiteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) S3BucketWebsiteOutputReference {
	_init_.Initialize()

	if err := validateNewS3BucketWebsiteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3BucketWebsiteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.s3Bucket.S3BucketWebsiteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3BucketWebsiteOutputReference_Override(s S3BucketWebsiteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.s3Bucket.S3BucketWebsiteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference)SetErrorDocument(val *string) {
	if err := j.validateSetErrorDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorDocument",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference)SetIndexDocument(val *string) {
	if err := j.validateSetIndexDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexDocument",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference)SetInternalValue(val *S3BucketWebsite) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference)SetRedirectAllRequestsTo(val *string) {
	if err := j.validateSetRedirectAllRequestsToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectAllRequestsTo",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference)SetRoutingRules(val *string) {
	if err := j.validateSetRoutingRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingRules",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) ResetErrorDocument() {
	_jsii_.InvokeVoid(
		s,
		"resetErrorDocument",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) ResetIndexDocument() {
	_jsii_.InvokeVoid(
		s,
		"resetIndexDocument",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) ResetRedirectAllRequestsTo() {
	_jsii_.InvokeVoid(
		s,
		"resetRedirectAllRequestsTo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) ResetRoutingRules() {
	_jsii_.InvokeVoid(
		s,
		"resetRoutingRules",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

