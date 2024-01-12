// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package prometheusscraper

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrometheusScraperDestinationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PrometheusScraperDestinationList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrometheusScraperDestinationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperDestinationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperDestinationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperDestinationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrometheusScraperDestinationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrometheusScraperDestinationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

