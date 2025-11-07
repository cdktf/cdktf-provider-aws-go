// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package defaultsecuritygroup

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DefaultSecurityGroupEgressList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DefaultSecurityGroupEgressList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DefaultSecurityGroupEgressList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DefaultSecurityGroupEgressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DefaultSecurityGroupEgressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DefaultSecurityGroupEgressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DefaultSecurityGroupEgressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDefaultSecurityGroupEgressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

