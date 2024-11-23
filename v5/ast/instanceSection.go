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

func InstanceSection() InstanceSectionClassLike {
	return instanceSectionReference()
}

// Constructor Methods

func (c *instanceSectionClass_) Make(
	instanceDeclarations abs.Sequential[InstanceDeclarationLike],
) InstanceSectionLike {
	if uti.IsUndefined(instanceDeclarations) {
		panic("The \"instanceDeclarations\" attribute is required by this class.")
	}
	var instance = &instanceSection_{
		// Initialize the instance attributes.
		instanceDeclarations_: instanceDeclarations,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *instanceSection_) GetClass() InstanceSectionClassLike {
	return instanceSectionReference()
}

// Attribute Methods

func (v *instanceSection_) GetInstanceDeclarations() abs.Sequential[InstanceDeclarationLike] {
	return v.instanceDeclarations_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type instanceSection_ struct {
	// Declare the instance attributes.
	instanceDeclarations_ abs.Sequential[InstanceDeclarationLike]
}

// Class Structure

type instanceSectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func instanceSectionReference() *instanceSectionClass_ {
	return instanceSectionReference_
}

var instanceSectionReference_ = &instanceSectionClass_{
	// Initialize the class constants.
}
