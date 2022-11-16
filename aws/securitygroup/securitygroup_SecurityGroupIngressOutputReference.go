package securitygroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/securitygroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityGroupIngressOutputReference interface {
	cdktf.ComplexObject
	CidrBlocks() *[]*string
	SetCidrBlocks(val *[]*string)
	CidrBlocksInput() *[]*string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	FromPort() *float64
	SetFromPort(val *float64)
	FromPortInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ipv6CidrBlocks() *[]*string
	SetIpv6CidrBlocks(val *[]*string)
	Ipv6CidrBlocksInput() *[]*string
	PrefixListIds() *[]*string
	SetPrefixListIds(val *[]*string)
	PrefixListIdsInput() *[]*string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	SelfAttribute() interface{}
	SetSelfAttribute(val interface{})
	SelfAttributeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ToPort() *float64
	SetToPort(val *float64)
	ToPortInput() *float64
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
	ResetCidrBlocks()
	ResetDescription()
	ResetFromPort()
	ResetIpv6CidrBlocks()
	ResetPrefixListIds()
	ResetProtocol()
	ResetSecurityGroups()
	ResetSelfAttribute()
	ResetToPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityGroupIngressOutputReference
type jsiiProxy_SecurityGroupIngressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) CidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) FromPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) FromPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) Ipv6CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) Ipv6CidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6CidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) PrefixListIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prefixListIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) PrefixListIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prefixListIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) SelfAttribute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) SelfAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) ToPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference) ToPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toPortInput",
		&returns,
	)
	return returns
}


func NewSecurityGroupIngressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SecurityGroupIngressOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityGroupIngressOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityGroupIngressOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.securityGroup.SecurityGroupIngressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSecurityGroupIngressOutputReference_Override(s SecurityGroupIngressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.securityGroup.SecurityGroupIngressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference)SetCidrBlocks(val *[]*string) {
	if err := j.validateSetCidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrBlocks",
		val,
	)
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference)SetFromPort(val *float64) {
	if err := j.validateSetFromPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fromPort",
		val,
	)
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference)SetIpv6CidrBlocks(val *[]*string) {
	if err := j.validateSetIpv6CidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6CidrBlocks",
		val,
	)
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference)SetPrefixListIds(val *[]*string) {
	if err := j.validateSetPrefixListIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixListIds",
		val,
	)
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference)SetSelfAttribute(val interface{}) {
	if err := j.validateSetSelfAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityGroupIngressOutputReference)SetToPort(val *float64) {
	if err := j.validateSetToPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toPort",
		val,
	)
}

func (s *jsiiProxy_SecurityGroupIngressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityGroupIngressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityGroupIngressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityGroupIngressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityGroupIngressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityGroupIngressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityGroupIngressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityGroupIngressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityGroupIngressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityGroupIngressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityGroupIngressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityGroupIngressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityGroupIngressOutputReference) ResetCidrBlocks() {
	_jsii_.InvokeVoid(
		s,
		"resetCidrBlocks",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityGroupIngressOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityGroupIngressOutputReference) ResetFromPort() {
	_jsii_.InvokeVoid(
		s,
		"resetFromPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityGroupIngressOutputReference) ResetIpv6CidrBlocks() {
	_jsii_.InvokeVoid(
		s,
		"resetIpv6CidrBlocks",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityGroupIngressOutputReference) ResetPrefixListIds() {
	_jsii_.InvokeVoid(
		s,
		"resetPrefixListIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityGroupIngressOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		s,
		"resetProtocol",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityGroupIngressOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		s,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityGroupIngressOutputReference) ResetSelfAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetSelfAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityGroupIngressOutputReference) ResetToPort() {
	_jsii_.InvokeVoid(
		s,
		"resetToPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityGroupIngressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecurityGroupIngressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

