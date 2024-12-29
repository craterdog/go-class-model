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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func PrimitiveDeclarationsClass() PrimitiveDeclarationsClassLike {
	return primitiveDeclarationsClass()
}

// Constructor Methods

func (c *primitiveDeclarationsClass_) PrimitiveDeclarations(
	typeSection TypeSectionLike,
	functionalSection FunctionalSectionLike,
) PrimitiveDeclarationsLike {
	if uti.IsUndefined(typeSection) {
		panic("The \"typeSection\" attribute is required by this class.")
	}
	if uti.IsUndefined(functionalSection) {
		panic("The \"functionalSection\" attribute is required by this class.")
	}
	var instance = &primitiveDeclarations_{
		// Initialize the instance attributes.
		typeSection_:       typeSection,
		functionalSection_: functionalSection,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *primitiveDeclarations_) GetClass() PrimitiveDeclarationsClassLike {
	return primitiveDeclarationsClass()
}

// Attribute Methods

func (v *primitiveDeclarations_) GetTypeSection() TypeSectionLike {
	return v.typeSection_
}

func (v *primitiveDeclarations_) GetFunctionalSection() FunctionalSectionLike {
	return v.functionalSection_
}

// PROTECTED INTERFACE

// Instance Structure

type primitiveDeclarations_ struct {
	// Declare the instance attributes.
	typeSection_       TypeSectionLike
	functionalSection_ FunctionalSectionLike
}

// Class Structure

type primitiveDeclarationsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func primitiveDeclarationsClass() *primitiveDeclarationsClass_ {
	return primitiveDeclarationsClassReference_
}

var primitiveDeclarationsClassReference_ = &primitiveDeclarationsClass_{
	// Initialize the class constants.
}
