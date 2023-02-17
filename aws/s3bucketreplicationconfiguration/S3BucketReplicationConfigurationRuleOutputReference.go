package s3bucketreplicationconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/s3bucketreplicationconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3BucketReplicationConfigurationRuleOutputReference interface {
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
	DeleteMarkerReplication() S3BucketReplicationConfigurationRuleDeleteMarkerReplicationOutputReference
	DeleteMarkerReplicationInput() *S3BucketReplicationConfigurationRuleDeleteMarkerReplication
	Destination() S3BucketReplicationConfigurationRuleDestinationOutputReference
	DestinationInput() *S3BucketReplicationConfigurationRuleDestination
	ExistingObjectReplication() S3BucketReplicationConfigurationRuleExistingObjectReplicationOutputReference
	ExistingObjectReplicationInput() *S3BucketReplicationConfigurationRuleExistingObjectReplication
	Filter() S3BucketReplicationConfigurationRuleFilterOutputReference
	FilterInput() *S3BucketReplicationConfigurationRuleFilter
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	SourceSelectionCriteria() S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference
	SourceSelectionCriteriaInput() *S3BucketReplicationConfigurationRuleSourceSelectionCriteria
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
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
	PutDeleteMarkerReplication(value *S3BucketReplicationConfigurationRuleDeleteMarkerReplication)
	PutDestination(value *S3BucketReplicationConfigurationRuleDestination)
	PutExistingObjectReplication(value *S3BucketReplicationConfigurationRuleExistingObjectReplication)
	PutFilter(value *S3BucketReplicationConfigurationRuleFilter)
	PutSourceSelectionCriteria(value *S3BucketReplicationConfigurationRuleSourceSelectionCriteria)
	ResetDeleteMarkerReplication()
	ResetExistingObjectReplication()
	ResetFilter()
	ResetId()
	ResetPrefix()
	ResetPriority()
	ResetSourceSelectionCriteria()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3BucketReplicationConfigurationRuleOutputReference
type jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) DeleteMarkerReplication() S3BucketReplicationConfigurationRuleDeleteMarkerReplicationOutputReference {
	var returns S3BucketReplicationConfigurationRuleDeleteMarkerReplicationOutputReference
	_jsii_.Get(
		j,
		"deleteMarkerReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) DeleteMarkerReplicationInput() *S3BucketReplicationConfigurationRuleDeleteMarkerReplication {
	var returns *S3BucketReplicationConfigurationRuleDeleteMarkerReplication
	_jsii_.Get(
		j,
		"deleteMarkerReplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) Destination() S3BucketReplicationConfigurationRuleDestinationOutputReference {
	var returns S3BucketReplicationConfigurationRuleDestinationOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) DestinationInput() *S3BucketReplicationConfigurationRuleDestination {
	var returns *S3BucketReplicationConfigurationRuleDestination
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) ExistingObjectReplication() S3BucketReplicationConfigurationRuleExistingObjectReplicationOutputReference {
	var returns S3BucketReplicationConfigurationRuleExistingObjectReplicationOutputReference
	_jsii_.Get(
		j,
		"existingObjectReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) ExistingObjectReplicationInput() *S3BucketReplicationConfigurationRuleExistingObjectReplication {
	var returns *S3BucketReplicationConfigurationRuleExistingObjectReplication
	_jsii_.Get(
		j,
		"existingObjectReplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) Filter() S3BucketReplicationConfigurationRuleFilterOutputReference {
	var returns S3BucketReplicationConfigurationRuleFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) FilterInput() *S3BucketReplicationConfigurationRuleFilter {
	var returns *S3BucketReplicationConfigurationRuleFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) SourceSelectionCriteria() S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference {
	var returns S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference
	_jsii_.Get(
		j,
		"sourceSelectionCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) SourceSelectionCriteriaInput() *S3BucketReplicationConfigurationRuleSourceSelectionCriteria {
	var returns *S3BucketReplicationConfigurationRuleSourceSelectionCriteria
	_jsii_.Get(
		j,
		"sourceSelectionCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3BucketReplicationConfigurationRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) S3BucketReplicationConfigurationRuleOutputReference {
	_init_.Initialize()

	if err := validateNewS3BucketReplicationConfigurationRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.s3BucketReplicationConfiguration.S3BucketReplicationConfigurationRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewS3BucketReplicationConfigurationRuleOutputReference_Override(s S3BucketReplicationConfigurationRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.s3BucketReplicationConfiguration.S3BucketReplicationConfigurationRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) PutDeleteMarkerReplication(value *S3BucketReplicationConfigurationRuleDeleteMarkerReplication) {
	if err := s.validatePutDeleteMarkerReplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDeleteMarkerReplication",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) PutDestination(value *S3BucketReplicationConfigurationRuleDestination) {
	if err := s.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDestination",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) PutExistingObjectReplication(value *S3BucketReplicationConfigurationRuleExistingObjectReplication) {
	if err := s.validatePutExistingObjectReplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExistingObjectReplication",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) PutFilter(value *S3BucketReplicationConfigurationRuleFilter) {
	if err := s.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFilter",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) PutSourceSelectionCriteria(value *S3BucketReplicationConfigurationRuleSourceSelectionCriteria) {
	if err := s.validatePutSourceSelectionCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSourceSelectionCriteria",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) ResetDeleteMarkerReplication() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteMarkerReplication",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) ResetExistingObjectReplication() {
	_jsii_.InvokeVoid(
		s,
		"resetExistingObjectReplication",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		s,
		"resetFilter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		s,
		"resetPriority",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) ResetSourceSelectionCriteria() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceSelectionCriteria",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

