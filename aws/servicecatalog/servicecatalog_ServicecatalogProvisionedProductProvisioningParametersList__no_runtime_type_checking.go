//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package servicecatalog

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServicecatalogProvisionedProductProvisioningParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServicecatalogProvisionedProductProvisioningParametersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServicecatalogProvisionedProductProvisioningParametersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServicecatalogProvisionedProductProvisioningParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServicecatalogProvisionedProductProvisioningParametersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServicecatalogProvisionedProductProvisioningParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServicecatalogProvisionedProductProvisioningParametersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

