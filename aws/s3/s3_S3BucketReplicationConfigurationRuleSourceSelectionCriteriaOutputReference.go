package s3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/s3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference interface {
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
	InternalValue() *S3BucketReplicationConfigurationRuleSourceSelectionCriteria
	SetInternalValue(val *S3BucketReplicationConfigurationRuleSourceSelectionCriteria)
	ReplicaModifications() S3BucketReplicationConfigurationRuleSourceSelectionCriteriaReplicaModificationsOutputReference
	ReplicaModificationsInput() *S3BucketReplicationConfigurationRuleSourceSelectionCriteriaReplicaModifications
	SseKmsEncryptedObjects() S3BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutputReference
	SseKmsEncryptedObjectsInput() *S3BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects
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
	PutReplicaModifications(value *S3BucketReplicationConfigurationRuleSourceSelectionCriteriaReplicaModifications)
	PutSseKmsEncryptedObjects(value *S3BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects)
	ResetReplicaModifications()
	ResetSseKmsEncryptedObjects()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference
type jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) InternalValue() *S3BucketReplicationConfigurationRuleSourceSelectionCriteria {
	var returns *S3BucketReplicationConfigurationRuleSourceSelectionCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) ReplicaModifications() S3BucketReplicationConfigurationRuleSourceSelectionCriteriaReplicaModificationsOutputReference {
	var returns S3BucketReplicationConfigurationRuleSourceSelectionCriteriaReplicaModificationsOutputReference
	_jsii_.Get(
		j,
		"replicaModifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) ReplicaModificationsInput() *S3BucketReplicationConfigurationRuleSourceSelectionCriteriaReplicaModifications {
	var returns *S3BucketReplicationConfigurationRuleSourceSelectionCriteriaReplicaModifications
	_jsii_.Get(
		j,
		"replicaModificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) SseKmsEncryptedObjects() S3BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutputReference {
	var returns S3BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutputReference
	_jsii_.Get(
		j,
		"sseKmsEncryptedObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) SseKmsEncryptedObjectsInput() *S3BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects {
	var returns *S3BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects
	_jsii_.Get(
		j,
		"sseKmsEncryptedObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewS3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.s3.S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference_Override(s S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.s3.S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference)SetInternalValue(val *S3BucketReplicationConfigurationRuleSourceSelectionCriteria) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) PutReplicaModifications(value *S3BucketReplicationConfigurationRuleSourceSelectionCriteriaReplicaModifications) {
	if err := s.validatePutReplicaModificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putReplicaModifications",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) PutSseKmsEncryptedObjects(value *S3BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects) {
	if err := s.validatePutSseKmsEncryptedObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSseKmsEncryptedObjects",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) ResetReplicaModifications() {
	_jsii_.InvokeVoid(
		s,
		"resetReplicaModifications",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) ResetSseKmsEncryptedObjects() {
	_jsii_.InvokeVoid(
		s,
		"resetSseKmsEncryptedObjects",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
