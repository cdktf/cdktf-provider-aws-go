package rdscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/rdscluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsClusterS3ImportOutputReference interface {
	cdktf.ComplexObject
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
	BucketPrefix() *string
	SetBucketPrefix(val *string)
	BucketPrefixInput() *string
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
	IngestionRole() *string
	SetIngestionRole(val *string)
	IngestionRoleInput() *string
	InternalValue() *RdsClusterS3Import
	SetInternalValue(val *RdsClusterS3Import)
	SourceEngine() *string
	SetSourceEngine(val *string)
	SourceEngineInput() *string
	SourceEngineVersion() *string
	SetSourceEngineVersion(val *string)
	SourceEngineVersionInput() *string
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
	ResetBucketPrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RdsClusterS3ImportOutputReference
type jsiiProxy_RdsClusterS3ImportOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) BucketPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) BucketPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) IngestionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingestionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) IngestionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingestionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) InternalValue() *RdsClusterS3Import {
	var returns *RdsClusterS3Import
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) SourceEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) SourceEngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) SourceEngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEngineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) SourceEngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEngineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRdsClusterS3ImportOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RdsClusterS3ImportOutputReference {
	_init_.Initialize()

	if err := validateNewRdsClusterS3ImportOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsClusterS3ImportOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.rdsCluster.RdsClusterS3ImportOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRdsClusterS3ImportOutputReference_Override(r RdsClusterS3ImportOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.rdsCluster.RdsClusterS3ImportOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference)SetBucketPrefix(val *string) {
	if err := j.validateSetBucketPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketPrefix",
		val,
	)
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference)SetIngestionRole(val *string) {
	if err := j.validateSetIngestionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingestionRole",
		val,
	)
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference)SetInternalValue(val *RdsClusterS3Import) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference)SetSourceEngine(val *string) {
	if err := j.validateSetSourceEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceEngine",
		val,
	)
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference)SetSourceEngineVersion(val *string) {
	if err := j.validateSetSourceEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceEngineVersion",
		val,
	)
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RdsClusterS3ImportOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RdsClusterS3ImportOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterS3ImportOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterS3ImportOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterS3ImportOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterS3ImportOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterS3ImportOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterS3ImportOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterS3ImportOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterS3ImportOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterS3ImportOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterS3ImportOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterS3ImportOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterS3ImportOutputReference) ResetBucketPrefix() {
	_jsii_.InvokeVoid(
		r,
		"resetBucketPrefix",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsClusterS3ImportOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterS3ImportOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

