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
	abs "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func FunctionalSectionClass() FunctionalSectionClassLike {
	return functionalSectionClass()
}

// Constructor Methods

func (c *functionalSectionClass_) FunctionalSection(
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

// INSTANCE INTERFACE

// Principal Methods

func (v *functionalSection_) GetClass() FunctionalSectionClassLike {
	return functionalSectionClass()
}

// Attribute Methods

func (v *functionalSection_) GetFunctionalDeclarations() abs.Sequential[FunctionalDeclarationLike] {
	return v.functionalDeclarations_
}

// PROTECTED INTERFACE

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

func functionalSectionClass() *functionalSectionClass_ {
	return functionalSectionClassReference_
}

var functionalSectionClassReference_ = &functionalSectionClass_{
	// Initialize the class constants.
}
