//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package vpnconnection

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpnConnectionRoutesList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VpnConnectionRoutesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VpnConnectionRoutesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VpnConnectionRoutesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VpnConnectionRoutesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVpnConnectionRoutesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

