package mqbroker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/mqbroker/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MqBrokerLdapServerMetadataOutputReference interface {
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
	Hosts() *[]*string
	SetHosts(val *[]*string)
	HostsInput() *[]*string
	InternalValue() *MqBrokerLdapServerMetadata
	SetInternalValue(val *MqBrokerLdapServerMetadata)
	RoleBase() *string
	SetRoleBase(val *string)
	RoleBaseInput() *string
	RoleName() *string
	SetRoleName(val *string)
	RoleNameInput() *string
	RoleSearchMatching() *string
	SetRoleSearchMatching(val *string)
	RoleSearchMatchingInput() *string
	RoleSearchSubtree() interface{}
	SetRoleSearchSubtree(val interface{})
	RoleSearchSubtreeInput() interface{}
	ServiceAccountPassword() *string
	SetServiceAccountPassword(val *string)
	ServiceAccountPasswordInput() *string
	ServiceAccountUsername() *string
	SetServiceAccountUsername(val *string)
	ServiceAccountUsernameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserBase() *string
	SetUserBase(val *string)
	UserBaseInput() *string
	UserRoleName() *string
	SetUserRoleName(val *string)
	UserRoleNameInput() *string
	UserSearchMatching() *string
	SetUserSearchMatching(val *string)
	UserSearchMatchingInput() *string
	UserSearchSubtree() interface{}
	SetUserSearchSubtree(val interface{})
	UserSearchSubtreeInput() interface{}
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
	ResetHosts()
	ResetRoleBase()
	ResetRoleName()
	ResetRoleSearchMatching()
	ResetRoleSearchSubtree()
	ResetServiceAccountPassword()
	ResetServiceAccountUsername()
	ResetUserBase()
	ResetUserRoleName()
	ResetUserSearchMatching()
	ResetUserSearchSubtree()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MqBrokerLdapServerMetadataOutputReference
type jsiiProxy_MqBrokerLdapServerMetadataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) Hosts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) HostsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) InternalValue() *MqBrokerLdapServerMetadata {
	var returns *MqBrokerLdapServerMetadata
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleSearchMatching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleSearchMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleSearchMatchingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleSearchMatchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleSearchSubtree() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roleSearchSubtree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) RoleSearchSubtreeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roleSearchSubtreeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ServiceAccountPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ServiceAccountPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ServiceAccountUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ServiceAccountUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserSearchMatching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSearchMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserSearchMatchingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSearchMatchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserSearchSubtree() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userSearchSubtree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) UserSearchSubtreeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userSearchSubtreeInput",
		&returns,
	)
	return returns
}


func NewMqBrokerLdapServerMetadataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MqBrokerLdapServerMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewMqBrokerLdapServerMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MqBrokerLdapServerMetadataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.mqBroker.MqBrokerLdapServerMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMqBrokerLdapServerMetadataOutputReference_Override(m MqBrokerLdapServerMetadataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.mqBroker.MqBrokerLdapServerMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetHosts(val *[]*string) {
	if err := j.validateSetHostsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hosts",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetInternalValue(val *MqBrokerLdapServerMetadata) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetRoleBase(val *string) {
	if err := j.validateSetRoleBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleBase",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetRoleName(val *string) {
	if err := j.validateSetRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleName",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetRoleSearchMatching(val *string) {
	if err := j.validateSetRoleSearchMatchingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleSearchMatching",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetRoleSearchSubtree(val interface{}) {
	if err := j.validateSetRoleSearchSubtreeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleSearchSubtree",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetServiceAccountPassword(val *string) {
	if err := j.validateSetServiceAccountPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountPassword",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetServiceAccountUsername(val *string) {
	if err := j.validateSetServiceAccountUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountUsername",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetUserBase(val *string) {
	if err := j.validateSetUserBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userBase",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetUserRoleName(val *string) {
	if err := j.validateSetUserRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userRoleName",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetUserSearchMatching(val *string) {
	if err := j.validateSetUserSearchMatchingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userSearchMatching",
		val,
	)
}

func (j *jsiiProxy_MqBrokerLdapServerMetadataOutputReference)SetUserSearchSubtree(val interface{}) {
	if err := j.validateSetUserSearchSubtreeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userSearchSubtree",
		val,
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetHosts() {
	_jsii_.InvokeVoid(
		m,
		"resetHosts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetRoleBase() {
	_jsii_.InvokeVoid(
		m,
		"resetRoleBase",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetRoleName() {
	_jsii_.InvokeVoid(
		m,
		"resetRoleName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetRoleSearchMatching() {
	_jsii_.InvokeVoid(
		m,
		"resetRoleSearchMatching",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetRoleSearchSubtree() {
	_jsii_.InvokeVoid(
		m,
		"resetRoleSearchSubtree",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetServiceAccountPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceAccountPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetServiceAccountUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceAccountUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetUserBase() {
	_jsii_.InvokeVoid(
		m,
		"resetUserBase",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetUserRoleName() {
	_jsii_.InvokeVoid(
		m,
		"resetUserRoleName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetUserSearchMatching() {
	_jsii_.InvokeVoid(
		m,
		"resetUserSearchMatching",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ResetUserSearchSubtree() {
	_jsii_.InvokeVoid(
		m,
		"resetUserSearchSubtree",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MqBrokerLdapServerMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

