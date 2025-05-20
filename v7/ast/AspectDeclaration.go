/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
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
	col "github.com/craterdog/go-collection-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func AspectDeclarationClass() AspectDeclarationClassLike {
	return aspectDeclarationClass()
}

// Constructor Methods

func (c *aspectDeclarationClass_) AspectDeclaration(
	declaration DeclarationLike,
	delimiter1 string,
	delimiter2 string,
	aspectMethods col.Sequential[AspectMethodLike],
	delimiter3 string,
) AspectDeclarationLike {
	if uti.IsUndefined(declaration) {
		panic("The \"declaration\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(aspectMethods) {
		panic("The \"aspectMethods\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter3) {
		panic("The \"delimiter3\" attribute is required by this class.")
	}
	var instance = &aspectDeclaration_{
		// Initialize the instance attributes.
		declaration_:   declaration,
		delimiter1_:    delimiter1,
		delimiter2_:    delimiter2,
		aspectMethods_: aspectMethods,
		delimiter3_:    delimiter3,
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

func (v *aspectDeclaration_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *aspectDeclaration_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *aspectDeclaration_) GetAspectMethods() col.Sequential[AspectMethodLike] {
	return v.aspectMethods_
}

func (v *aspectDeclaration_) GetDelimiter3() string {
	return v.delimiter3_
}

// PROTECTED INTERFACE

// Instance Structure

type aspectDeclaration_ struct {
	// Declare the instance attributes.
	declaration_   DeclarationLike
	delimiter1_    string
	delimiter2_    string
	aspectMethods_ col.Sequential[AspectMethodLike]
	delimiter3_    string
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
