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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func FunctionalSection() FunctionalSectionClassLike {
	return functionalSectionReference()
}

// Constructor Methods

func (c *functionalSectionClass_) Make(
	functionalDeclarations abs.Sequential[FunctionalDeclarationLike],
) FunctionalSectionLike {
	if uti.IsUndefined(functionalDeclarations) {
		panic("The \"functionalDeclarations\" attribute is required by this class.")
	}
	var instance = &functionalSection_{
		// Initialize the instance attributes.
		functionalDeclarations_: functionalDeclarations,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *functionalSection_) GetClass() FunctionalSectionClassLike {
	return functionalSectionReference()
}

// Attribute Methods

func (v *functionalSection_) GetFunctionalDeclarations() abs.Sequential[FunctionalDeclarationLike] {
	return v.functionalDeclarations_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type functionalSection_ struct {
	// Declare the instance attributes.
	functionalDeclarations_ abs.Sequential[FunctionalDeclarationLike]
}

// Class Structure

type functionalSectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func functionalSectionReference() *functionalSectionClass_ {
	return functionalSectionReference_
}

var functionalSectionReference_ = &functionalSectionClass_{
	// Initialize the class constants.
}
