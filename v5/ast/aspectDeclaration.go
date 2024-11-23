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

func AspectDeclaration() AspectDeclarationClassLike {
	return aspectDeclarationReference()
}

// Constructor Methods

func (c *aspectDeclarationClass_) Make(
	declaration DeclarationLike,
	aspectMethods abs.Sequential[AspectMethodLike],
) AspectDeclarationLike {
	if uti.IsUndefined(declaration) {
		panic("The \"declaration\" attribute is required by this class.")
	}
	if uti.IsUndefined(aspectMethods) {
		panic("The \"aspectMethods\" attribute is required by this class.")
	}
	var instance = &aspectDeclaration_{
		// Initialize the instance attributes.
		declaration_:   declaration,
		aspectMethods_: aspectMethods,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *aspectDeclaration_) GetClass() AspectDeclarationClassLike {
	return aspectDeclarationReference()
}

// Attribute Methods

func (v *aspectDeclaration_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *aspectDeclaration_) GetAspectMethods() abs.Sequential[AspectMethodLike] {
	return v.aspectMethods_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type aspectDeclaration_ struct {
	// Declare the instance attributes.
	declaration_   DeclarationLike
	aspectMethods_ abs.Sequential[AspectMethodLike]
}

// Class Structure

type aspectDeclarationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func aspectDeclarationReference() *aspectDeclarationClass_ {
	return aspectDeclarationReference_
}

var aspectDeclarationReference_ = &aspectDeclarationClass_{
	// Initialize the class constants.
}
