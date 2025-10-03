// Copyright 2020-2025 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func ExampleTagID_Valid() {
	// Valid string TagID
	validStringTagID := TagID{val: "com.acme.product-v1.0.0"}
	if err := validStringTagID.Valid(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("String TagID is valid")
	}

	// Valid UUID TagID
	validUUIDTagID := TagID{val: uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")}
	if err := validUUIDTagID.Valid(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("UUID TagID is valid")
	}

	// Invalid empty string TagID
	emptyStringTagID := TagID{val: ""}
	if err := emptyStringTagID.Valid(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Invalid nil TagID
	nilTagID := TagID{val: nil}
	if err := nilTagID.Valid(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Output:
	// String TagID is valid
	// UUID TagID is valid
	// Error: tag-id string value is empty
	// Error: tag-id value is nil
}

func ExampleEvidence_Valid() {
	// Valid Evidence
	validEvidence := Evidence{
		DeviceID: "device-001",
		Date:     time.Now(),
	}
	if err := validEvidence.Valid(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Evidence is valid")
	}

	// Invalid Evidence - empty DeviceID
	invalidEvidence1 := Evidence{
		DeviceID: "",
		Date:     time.Now(),
	}
	if err := invalidEvidence1.Valid(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Invalid Evidence - zero Date
	invalidEvidence2 := Evidence{
		DeviceID: "device-001",
		Date:     time.Time{},
	}
	if err := invalidEvidence2.Valid(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Output:
	// Evidence is valid
	// Error: evidence device-id is empty
	// Error: evidence date is zero
}
