package mskcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/mskcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference interface {
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
	InternalValue() *MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfo
	SetInternalValue(val *MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfo)
	ProvisionedThroughput() MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputOutputReference
	ProvisionedThroughputInput() *MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughput
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VolumeSize() *float64
	SetVolumeSize(val *float64)
	VolumeSizeInput() *float64
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
	PutProvisionedThroughput(value *MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughput)
	ResetProvisionedThroughput()
	ResetVolumeSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference
type jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) InternalValue() *MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfo {
	var returns *MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) ProvisionedThroughput() MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputOutputReference {
	var returns MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputOutputReference
	_jsii_.Get(
		j,
		"provisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) ProvisionedThroughputInput() *MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughput {
	var returns *MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughput
	_jsii_.Get(
		j,
		"provisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) VolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) VolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInput",
		&returns,
	)
	return returns
}


func NewMskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference {
	_init_.Initialize()

	if err := validateNewMskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.mskCluster.MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference_Override(m MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.mskCluster.MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference)SetInternalValue(val *MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference)SetVolumeSize(val *float64) {
	if err := j.validateSetVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSize",
		val,
	)
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) PutProvisionedThroughput(value *MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughput) {
	if err := m.validatePutProvisionedThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putProvisionedThroughput",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) ResetProvisionedThroughput() {
	_jsii_.InvokeVoid(
		m,
		"resetProvisionedThroughput",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) ResetVolumeSize() {
	_jsii_.InvokeVoid(
		m,
		"resetVolumeSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

