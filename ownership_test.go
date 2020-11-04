// Copyright 2020 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package swid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOwnership_String(t *testing.T) {
	var o Ownership

	o = Ownership{OwnershipShared}
	assert.Equal(t, "shared", o.String())

	o = Ownership{OwnershipPrivate}
	assert.Equal(t, "private", o.String())

	o = Ownership{OwnershipAbandon}
	assert.Equal(t, "abandon", o.String())

	o = Ownership{uint64(100)}
	assert.Equal(t, "ownership(100)", o.String())

	o = Ownership{"ozymandias"}
	assert.Equal(t, "ozymandias", o.String())

	o = Ownership{map[string]string{"hey": "duggie"}}
	assert.Equal(t, "", o.String())
}

func TestOwnership_Check(t *testing.T) {
	var o Ownership

	o = Ownership{OwnershipShared}
	assert.Nil(t, o.Check())

	o = Ownership{OwnershipPrivate}
	assert.Nil(t, o.Check())

	o = Ownership{OwnershipAbandon}
	assert.Nil(t, o.Check())

	o = Ownership{uint64(100)}
	assert.Nil(t, o.Check())

	o = Ownership{"ozymandias"}
	assert.Nil(t, o.Check())

	o = Ownership{map[string]string{"hey": "duggie"}}
	assert.EqualError(t, o.Check(), "ownership MUST be uint64 or string; got map[string]string")
}
