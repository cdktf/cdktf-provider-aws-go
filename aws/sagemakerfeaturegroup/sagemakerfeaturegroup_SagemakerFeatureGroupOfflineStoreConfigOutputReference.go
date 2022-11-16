package sagemakerfeaturegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/sagemakerfeaturegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerFeatureGroupOfflineStoreConfigOutputReference interface {
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
	DataCatalogConfig() SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfigOutputReference
	DataCatalogConfigInput() *SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfig
	DisableGlueTableCreation() interface{}
	SetDisableGlueTableCreation(val interface{})
	DisableGlueTableCreationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerFeatureGroupOfflineStoreConfig
	SetInternalValue(val *SagemakerFeatureGroupOfflineStoreConfig)
	S3StorageConfig() SagemakerFeatureGroupOfflineStoreConfigS3StorageConfigOutputReference
	S3StorageConfigInput() *SagemakerFeatureGroupOfflineStoreConfigS3StorageConfig
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
	PutDataCatalogConfig(value *SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfig)
	PutS3StorageConfig(value *SagemakerFeatureGroupOfflineStoreConfigS3StorageConfig)
	ResetDataCatalogConfig()
	ResetDisableGlueTableCreation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerFeatureGroupOfflineStoreConfigOutputReference
type jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) DataCatalogConfig() SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfigOutputReference {
	var returns SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfigOutputReference
	_jsii_.Get(
		j,
		"dataCatalogConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) DataCatalogConfigInput() *SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfig {
	var returns *SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfig
	_jsii_.Get(
		j,
		"dataCatalogConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) DisableGlueTableCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableGlueTableCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) DisableGlueTableCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableGlueTableCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) InternalValue() *SagemakerFeatureGroupOfflineStoreConfig {
	var returns *SagemakerFeatureGroupOfflineStoreConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) S3StorageConfig() SagemakerFeatureGroupOfflineStoreConfigS3StorageConfigOutputReference {
	var returns SagemakerFeatureGroupOfflineStoreConfigS3StorageConfigOutputReference
	_jsii_.Get(
		j,
		"s3StorageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) S3StorageConfigInput() *SagemakerFeatureGroupOfflineStoreConfigS3StorageConfig {
	var returns *SagemakerFeatureGroupOfflineStoreConfigS3StorageConfig
	_jsii_.Get(
		j,
		"s3StorageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerFeatureGroupOfflineStoreConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerFeatureGroupOfflineStoreConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerFeatureGroupOfflineStoreConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerFeatureGroup.SagemakerFeatureGroupOfflineStoreConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerFeatureGroupOfflineStoreConfigOutputReference_Override(s SagemakerFeatureGroupOfflineStoreConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerFeatureGroup.SagemakerFeatureGroupOfflineStoreConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference)SetDisableGlueTableCreation(val interface{}) {
	if err := j.validateSetDisableGlueTableCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableGlueTableCreation",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference)SetInternalValue(val *SagemakerFeatureGroupOfflineStoreConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) PutDataCatalogConfig(value *SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfig) {
	if err := s.validatePutDataCatalogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDataCatalogConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) PutS3StorageConfig(value *SagemakerFeatureGroupOfflineStoreConfigS3StorageConfig) {
	if err := s.validatePutS3StorageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putS3StorageConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) ResetDataCatalogConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetDataCatalogConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) ResetDisableGlueTableCreation() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableGlueTableCreation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

