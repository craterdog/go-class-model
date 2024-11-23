/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│   Updates to any section other than the Private Methods may be overwritten.  │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func PrincipalSubsection() PrincipalSubsectionClassLike {
	return principalSubsectionReference()
}

// Constructor Methods

func (c *principalSubsectionClass_) Make(
	principalMethods abs.Sequential[PrincipalMethodLike],
) PrincipalSubsectionLike {
	if uti.IsUndefined(principalMethods) {
		panic("The \"principalMethods\" attribute is required by this class.")
	}
	var instance = &principalSubsection_{
		// Initialize the instance attributes.
		principalMethods_: principalMethods,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *principalSubsection_) GetClass() PrincipalSubsectionClassLike {
	return principalSubsectionReference()
}

// Attribute Methods

func (v *principalSubsection_) GetPrincipalMethods() abs.Sequential[PrincipalMethodLike] {
	return v.principalMethods_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type principalSubsection_ struct {
	// Declare the instance attributes.
	principalMethods_ abs.Sequential[PrincipalMethodLike]
}

// Class Structure

type principalSubsectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func principalSubsectionReference() *principalSubsectionClass_ {
	return principalSubsectionReference_
}

var principalSubsectionReference_ = &principalSubsectionClass_{
	// Initialize the class constants.
}