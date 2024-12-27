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

func AspectDeclarationClass() AspectDeclarationClassLike {
	return aspectDeclarationClass()
}

// Constructor Methods

func (c *aspectDeclarationClass_) AspectDeclaration(
	declaration DeclarationLike,
	aspectMethods col.Sequential[AspectMethodLike],
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

// INSTANCE INTERFACE

// Principal Methods

func (v *aspectDeclaration_) GetClass() AspectDeclarationClassLike {
	return aspectDeclarationClass()
}

// Attribute Methods

func (v *aspectDeclaration_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *aspectDeclaration_) GetAspectMethods() col.Sequential[AspectMethodLike] {
	return v.aspectMethods_
}

// PROTECTED INTERFACE

// Instance Structure

type aspectDeclaration_ struct {
	// Declare the instance attributes.
	declaration_   DeclarationLike
	aspectMethods_ col.Sequential[AspectMethodLike]
}

// Class Structure

type aspectDeclarationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func aspectDeclarationClass() *aspectDeclarationClass_ {
	return aspectDeclarationClassReference_
}

var aspectDeclarationClassReference_ = &aspectDeclarationClass_{
	// Initialize the class constants.
}
