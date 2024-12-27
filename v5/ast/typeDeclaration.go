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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func TypeDeclarationClass() TypeDeclarationClassLike {
	return typeDeclarationClass()
}

// Constructor Methods

func (c *typeDeclarationClass_) TypeDeclaration(
	declaration DeclarationLike,
	abstraction AbstractionLike,
	optionalEnumeration EnumerationLike,
) TypeDeclarationLike {
	if uti.IsUndefined(declaration) {
		panic("The \"declaration\" attribute is required by this class.")
	}
	if uti.IsUndefined(abstraction) {
		panic("The \"abstraction\" attribute is required by this class.")
	}
	var instance = &typeDeclaration_{
		// Initialize the instance attributes.
		declaration_:         declaration,
		abstraction_:         abstraction,
		optionalEnumeration_: optionalEnumeration,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *typeDeclaration_) GetClass() TypeDeclarationClassLike {
	return typeDeclarationClass()
}

// Attribute Methods

func (v *typeDeclaration_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *typeDeclaration_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

func (v *typeDeclaration_) GetOptionalEnumeration() EnumerationLike {
	return v.optionalEnumeration_
}

// PROTECTED INTERFACE

// Instance Structure

type typeDeclaration_ struct {
	// Declare the instance attributes.
	declaration_         DeclarationLike
	abstraction_         AbstractionLike
	optionalEnumeration_ EnumerationLike
}

// Class Structure

type typeDeclarationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func typeDeclarationClass() *typeDeclarationClass_ {
	return typeDeclarationClassReference_
}

var typeDeclarationClassReference_ = &typeDeclarationClass_{
	// Initialize the class constants.
}
