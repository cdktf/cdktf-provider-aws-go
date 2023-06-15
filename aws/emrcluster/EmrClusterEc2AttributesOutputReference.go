package emrcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v15/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v15/emrcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrClusterEc2AttributesOutputReference interface {
	cdktf.ComplexObject
	AdditionalMasterSecurityGroups() *string
	SetAdditionalMasterSecurityGroups(val *string)
	AdditionalMasterSecurityGroupsInput() *string
	AdditionalSlaveSecurityGroups() *string
	SetAdditionalSlaveSecurityGroups(val *string)
	AdditionalSlaveSecurityGroupsInput() *string
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
	EmrManagedMasterSecurityGroup() *string
	SetEmrManagedMasterSecurityGroup(val *string)
	EmrManagedMasterSecurityGroupInput() *string
	EmrManagedSlaveSecurityGroup() *string
	SetEmrManagedSlaveSecurityGroup(val *string)
	EmrManagedSlaveSecurityGroupInput() *string
	// Experimental.
	Fqn() *string
	InstanceProfile() *string
	SetInstanceProfile(val *string)
	InstanceProfileInput() *string
	InternalValue() *EmrClusterEc2Attributes
	SetInternalValue(val *EmrClusterEc2Attributes)
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	ServiceAccessSecurityGroup() *string
	SetServiceAccessSecurityGroup(val *string)
	ServiceAccessSecurityGroupInput() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
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
	ResetAdditionalMasterSecurityGroups()
	ResetAdditionalSlaveSecurityGroups()
	ResetEmrManagedMasterSecurityGroup()
	ResetEmrManagedSlaveSecurityGroup()
	ResetKeyName()
	ResetServiceAccessSecurityGroup()
	ResetSubnetId()
	ResetSubnetIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmrClusterEc2AttributesOutputReference
type jsiiProxy_EmrClusterEc2AttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) AdditionalMasterSecurityGroups() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalMasterSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) AdditionalMasterSecurityGroupsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalMasterSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) AdditionalSlaveSecurityGroups() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalSlaveSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) AdditionalSlaveSecurityGroupsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalSlaveSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) EmrManagedMasterSecurityGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emrManagedMasterSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) EmrManagedMasterSecurityGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emrManagedMasterSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) EmrManagedSlaveSecurityGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emrManagedSlaveSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) EmrManagedSlaveSecurityGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emrManagedSlaveSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) InstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) InstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) InternalValue() *EmrClusterEc2Attributes {
	var returns *EmrClusterEc2Attributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) ServiceAccessSecurityGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) ServiceAccessSecurityGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEmrClusterEc2AttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EmrClusterEc2AttributesOutputReference {
	_init_.Initialize()

	if err := validateNewEmrClusterEc2AttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrClusterEc2AttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrCluster.EmrClusterEc2AttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEmrClusterEc2AttributesOutputReference_Override(e EmrClusterEc2AttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrCluster.EmrClusterEc2AttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference)SetAdditionalMasterSecurityGroups(val *string) {
	if err := j.validateSetAdditionalMasterSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalMasterSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference)SetAdditionalSlaveSecurityGroups(val *string) {
	if err := j.validateSetAdditionalSlaveSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalSlaveSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference)SetEmrManagedMasterSecurityGroup(val *string) {
	if err := j.validateSetEmrManagedMasterSecurityGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emrManagedMasterSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference)SetEmrManagedSlaveSecurityGroup(val *string) {
	if err := j.validateSetEmrManagedSlaveSecurityGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emrManagedSlaveSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference)SetInstanceProfile(val *string) {
	if err := j.validateSetInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfile",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference)SetInternalValue(val *EmrClusterEc2Attributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference)SetServiceAccessSecurityGroup(val *string) {
	if err := j.validateSetServiceAccessSecurityGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetAdditionalMasterSecurityGroups() {
	_jsii_.InvokeVoid(
		e,
		"resetAdditionalMasterSecurityGroups",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetAdditionalSlaveSecurityGroups() {
	_jsii_.InvokeVoid(
		e,
		"resetAdditionalSlaveSecurityGroups",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetEmrManagedMasterSecurityGroup() {
	_jsii_.InvokeVoid(
		e,
		"resetEmrManagedMasterSecurityGroup",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetEmrManagedSlaveSecurityGroup() {
	_jsii_.InvokeVoid(
		e,
		"resetEmrManagedSlaveSecurityGroup",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetKeyName() {
	_jsii_.InvokeVoid(
		e,
		"resetKeyName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetServiceAccessSecurityGroup() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceAccessSecurityGroup",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		e,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

