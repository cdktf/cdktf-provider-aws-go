package s3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/s3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3BucketLifecycleRuleOutputReference interface {
	cdktf.ComplexObject
	AbortIncompleteMultipartUploadDays() *float64
	SetAbortIncompleteMultipartUploadDays(val *float64)
	AbortIncompleteMultipartUploadDaysInput() *float64
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Expiration() S3BucketLifecycleRuleExpirationOutputReference
	ExpirationInput() *S3BucketLifecycleRuleExpiration
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NoncurrentVersionExpiration() S3BucketLifecycleRuleNoncurrentVersionExpirationOutputReference
	NoncurrentVersionExpirationInput() *S3BucketLifecycleRuleNoncurrentVersionExpiration
	NoncurrentVersionTransition() S3BucketLifecycleRuleNoncurrentVersionTransitionList
	NoncurrentVersionTransitionInput() interface{}
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Transition() S3BucketLifecycleRuleTransitionList
	TransitionInput() interface{}
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
	PutExpiration(value *S3BucketLifecycleRuleExpiration)
	PutNoncurrentVersionExpiration(value *S3BucketLifecycleRuleNoncurrentVersionExpiration)
	PutNoncurrentVersionTransition(value interface{})
	PutTransition(value interface{})
	ResetAbortIncompleteMultipartUploadDays()
	ResetExpiration()
	ResetId()
	ResetNoncurrentVersionExpiration()
	ResetNoncurrentVersionTransition()
	ResetPrefix()
	ResetTags()
	ResetTransition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3BucketLifecycleRuleOutputReference
type jsiiProxy_S3BucketLifecycleRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) AbortIncompleteMultipartUploadDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUploadDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) AbortIncompleteMultipartUploadDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUploadDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) Expiration() S3BucketLifecycleRuleExpirationOutputReference {
	var returns S3BucketLifecycleRuleExpirationOutputReference
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) ExpirationInput() *S3BucketLifecycleRuleExpiration {
	var returns *S3BucketLifecycleRuleExpiration
	_jsii_.Get(
		j,
		"expirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) NoncurrentVersionExpiration() S3BucketLifecycleRuleNoncurrentVersionExpirationOutputReference {
	var returns S3BucketLifecycleRuleNoncurrentVersionExpirationOutputReference
	_jsii_.Get(
		j,
		"noncurrentVersionExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) NoncurrentVersionExpirationInput() *S3BucketLifecycleRuleNoncurrentVersionExpiration {
	var returns *S3BucketLifecycleRuleNoncurrentVersionExpiration
	_jsii_.Get(
		j,
		"noncurrentVersionExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) NoncurrentVersionTransition() S3BucketLifecycleRuleNoncurrentVersionTransitionList {
	var returns S3BucketLifecycleRuleNoncurrentVersionTransitionList
	_jsii_.Get(
		j,
		"noncurrentVersionTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) NoncurrentVersionTransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noncurrentVersionTransitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) Transition() S3BucketLifecycleRuleTransitionList {
	var returns S3BucketLifecycleRuleTransitionList
	_jsii_.Get(
		j,
		"transition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference) TransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitionInput",
		&returns,
	)
	return returns
}


func NewS3BucketLifecycleRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) S3BucketLifecycleRuleOutputReference {
	_init_.Initialize()

	if err := validateNewS3BucketLifecycleRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3BucketLifecycleRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.s3.S3BucketLifecycleRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewS3BucketLifecycleRuleOutputReference_Override(s S3BucketLifecycleRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.s3.S3BucketLifecycleRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference)SetAbortIncompleteMultipartUploadDays(val *float64) {
	if err := j.validateSetAbortIncompleteMultipartUploadDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"abortIncompleteMultipartUploadDays",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) PutExpiration(value *S3BucketLifecycleRuleExpiration) {
	if err := s.validatePutExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExpiration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) PutNoncurrentVersionExpiration(value *S3BucketLifecycleRuleNoncurrentVersionExpiration) {
	if err := s.validatePutNoncurrentVersionExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNoncurrentVersionExpiration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) PutNoncurrentVersionTransition(value interface{}) {
	if err := s.validatePutNoncurrentVersionTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNoncurrentVersionTransition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) PutTransition(value interface{}) {
	if err := s.validatePutTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTransition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) ResetAbortIncompleteMultipartUploadDays() {
	_jsii_.InvokeVoid(
		s,
		"resetAbortIncompleteMultipartUploadDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) ResetExpiration() {
	_jsii_.InvokeVoid(
		s,
		"resetExpiration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) ResetNoncurrentVersionExpiration() {
	_jsii_.InvokeVoid(
		s,
		"resetNoncurrentVersionExpiration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) ResetNoncurrentVersionTransition() {
	_jsii_.InvokeVoid(
		s,
		"resetNoncurrentVersionTransition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) ResetTransition() {
	_jsii_.InvokeVoid(
		s,
		"resetTransition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3BucketLifecycleRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

