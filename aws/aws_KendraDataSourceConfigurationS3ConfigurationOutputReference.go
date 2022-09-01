// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendraDataSourceConfigurationS3ConfigurationOutputReference interface {
	cdktf.ComplexObject
	AccessControlListConfiguration() KendraDataSourceConfigurationS3ConfigurationAccessControlListConfigurationOutputReference
	AccessControlListConfigurationInput() *KendraDataSourceConfigurationS3ConfigurationAccessControlListConfiguration
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
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
	DocumentsMetadataConfiguration() KendraDataSourceConfigurationS3ConfigurationDocumentsMetadataConfigurationOutputReference
	DocumentsMetadataConfigurationInput() *KendraDataSourceConfigurationS3ConfigurationDocumentsMetadataConfiguration
	ExclusionPatterns() *[]*string
	SetExclusionPatterns(val *[]*string)
	ExclusionPatternsInput() *[]*string
	// Experimental.
	Fqn() *string
	InclusionPatterns() *[]*string
	SetInclusionPatterns(val *[]*string)
	InclusionPatternsInput() *[]*string
	InclusionPrefixes() *[]*string
	SetInclusionPrefixes(val *[]*string)
	InclusionPrefixesInput() *[]*string
	InternalValue() *KendraDataSourceConfigurationS3Configuration
	SetInternalValue(val *KendraDataSourceConfigurationS3Configuration)
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
	PutAccessControlListConfiguration(value *KendraDataSourceConfigurationS3ConfigurationAccessControlListConfiguration)
	PutDocumentsMetadataConfiguration(value *KendraDataSourceConfigurationS3ConfigurationDocumentsMetadataConfiguration)
	ResetAccessControlListConfiguration()
	ResetDocumentsMetadataConfiguration()
	ResetExclusionPatterns()
	ResetInclusionPatterns()
	ResetInclusionPrefixes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceConfigurationS3ConfigurationOutputReference
type jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) AccessControlListConfiguration() KendraDataSourceConfigurationS3ConfigurationAccessControlListConfigurationOutputReference {
	var returns KendraDataSourceConfigurationS3ConfigurationAccessControlListConfigurationOutputReference
	_jsii_.Get(
		j,
		"accessControlListConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) AccessControlListConfigurationInput() *KendraDataSourceConfigurationS3ConfigurationAccessControlListConfiguration {
	var returns *KendraDataSourceConfigurationS3ConfigurationAccessControlListConfiguration
	_jsii_.Get(
		j,
		"accessControlListConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) DocumentsMetadataConfiguration() KendraDataSourceConfigurationS3ConfigurationDocumentsMetadataConfigurationOutputReference {
	var returns KendraDataSourceConfigurationS3ConfigurationDocumentsMetadataConfigurationOutputReference
	_jsii_.Get(
		j,
		"documentsMetadataConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) DocumentsMetadataConfigurationInput() *KendraDataSourceConfigurationS3ConfigurationDocumentsMetadataConfiguration {
	var returns *KendraDataSourceConfigurationS3ConfigurationDocumentsMetadataConfiguration
	_jsii_.Get(
		j,
		"documentsMetadataConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) ExclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) ExclusionPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusionPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) InclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) InclusionPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) InclusionPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) InclusionPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) InternalValue() *KendraDataSourceConfigurationS3Configuration {
	var returns *KendraDataSourceConfigurationS3Configuration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKendraDataSourceConfigurationS3ConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KendraDataSourceConfigurationS3ConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceConfigurationS3ConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.KendraDataSourceConfigurationS3ConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraDataSourceConfigurationS3ConfigurationOutputReference_Override(k KendraDataSourceConfigurationS3ConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.KendraDataSourceConfigurationS3ConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference)SetExclusionPatterns(val *[]*string) {
	if err := j.validateSetExclusionPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusionPatterns",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference)SetInclusionPatterns(val *[]*string) {
	if err := j.validateSetInclusionPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inclusionPatterns",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference)SetInclusionPrefixes(val *[]*string) {
	if err := j.validateSetInclusionPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inclusionPrefixes",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference)SetInternalValue(val *KendraDataSourceConfigurationS3Configuration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) PutAccessControlListConfiguration(value *KendraDataSourceConfigurationS3ConfigurationAccessControlListConfiguration) {
	if err := k.validatePutAccessControlListConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAccessControlListConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) PutDocumentsMetadataConfiguration(value *KendraDataSourceConfigurationS3ConfigurationDocumentsMetadataConfiguration) {
	if err := k.validatePutDocumentsMetadataConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putDocumentsMetadataConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) ResetAccessControlListConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetAccessControlListConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) ResetDocumentsMetadataConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetDocumentsMetadataConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) ResetExclusionPatterns() {
	_jsii_.InvokeVoid(
		k,
		"resetExclusionPatterns",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) ResetInclusionPatterns() {
	_jsii_.InvokeVoid(
		k,
		"resetInclusionPatterns",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) ResetInclusionPrefixes() {
	_jsii_.InvokeVoid(
		k,
		"resetInclusionPrefixes",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationS3ConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

