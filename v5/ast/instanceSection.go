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
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func InstanceSectionClass() InstanceSectionClassLike {
	return instanceSectionClass()
}

// Constructor Methods

func (c *instanceSectionClass_) InstanceSection(
	instanceDeclarations col.Sequential[InstanceDeclarationLike],
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

// INSTANCE INTERFACE

// Principal Methods

func (v *instanceSection_) GetClass() InstanceSectionClassLike {
	return instanceSectionClass()
}

// Attribute Methods

func (v *instanceSection_) GetInstanceDeclarations() col.Sequential[InstanceDeclarationLike] {
	return v.instanceDeclarations_
}

// PROTECTED INTERFACE

// Instance Structure

type instanceSection_ struct {
	// Declare the instance attributes.
	instanceDeclarations_ col.Sequential[InstanceDeclarationLike]
}

// Class Structure

type instanceSectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func instanceSectionClass() *instanceSectionClass_ {
	return instanceSectionClassReference_
}

var instanceSectionClassReference_ = &instanceSectionClass_{
	// Initialize the class constants.
}
