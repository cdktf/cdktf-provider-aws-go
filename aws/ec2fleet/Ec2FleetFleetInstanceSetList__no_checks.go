// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ec2fleet

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2FleetFleetInstanceSetList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_Ec2FleetFleetInstanceSetList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2FleetFleetInstanceSetList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2FleetFleetInstanceSetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2FleetFleetInstanceSetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2FleetFleetInstanceSetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2FleetFleetInstanceSetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2FleetFleetInstanceSetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

