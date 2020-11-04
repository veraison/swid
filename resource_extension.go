// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

var (
	// ResourceTypePSAMeasuredSoftwareComponent is the resource type to use for
	// a PSA measured software component
	ResourceTypePSAMeasuredSoftwareComponent = "arm.com-PSAMeasuredSoftwareComponent"
)

// ResourceExtension is a placeholder for $$resource-extension
type ResourceExtension struct {
	// PSA endorsements extensions
	PSAMeasuredSoftwareComponent
}

// PSAMeasuredSoftwareComponent describes a PSA measured software component
// See Section 3.4.1 of draft-tschofenig-rats-psa-token-05
type PSAMeasuredSoftwareComponent struct {
	MeasurementValue HashEntry `cbor:"arm.com-PSAMeasurementValue" json:"arm.com-PSAMeasurementValue" xml:"measurementValue,attr"`
	SignerID         HashEntry `cbor:"arm.com-PSASignerId" json:"arm.com-PSASignerId" xml:"signerId,attr"`
}

// NewPSAMeasuredSoftwareComponentResource returns a Resource of type
// PSAMeasuredSoftwareComponent initialized according to the supplied
// measurement value and signer ID
func NewPSAMeasuredSoftwareComponentResource(
	measurementValue HashEntry, signerID HashEntry,
) (*Resource, error) {
	return &Resource{
		Type: ResourceTypePSAMeasuredSoftwareComponent,
		ResourceExtension: ResourceExtension{
			PSAMeasuredSoftwareComponent{
				MeasurementValue: measurementValue,
				SignerID:         signerID,
			},
		},
	}, nil
}
