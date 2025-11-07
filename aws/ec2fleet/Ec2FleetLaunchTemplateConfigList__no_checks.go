// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ec2fleet

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2FleetLaunchTemplateConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

