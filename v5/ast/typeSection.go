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

func TypeSectionClass() TypeSectionClassLike {
	return typeSectionClass()
}

// Constructor Methods

func (c *typeSectionClass_) TypeSection(
	typeDeclarations abs.Sequential[TypeDeclarationLike],
) TypeSectionLike {
	if uti.IsUndefined(typeDeclarations) {
		panic("The \"typeDeclarations\" attribute is required by this class.")
	}
	var instance = &typeSection_{
		// Initialize the instance attributes.
		typeDeclarations_: typeDeclarations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *typeSection_) GetClass() TypeSectionClassLike {
	return typeSectionClass()
}

// Attribute Methods

func (v *typeSection_) GetTypeDeclarations() abs.Sequential[TypeDeclarationLike] {
	return v.typeDeclarations_
}

// PROTECTED INTERFACE

// Instance Structure

type typeSection_ struct {
	// Declare the instance attributes.
	typeDeclarations_ abs.Sequential[TypeDeclarationLike]
}

// Class Structure

type typeSectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func typeSectionClass() *typeSectionClass_ {
	return typeSectionClassReference_
}

var typeSectionClassReference_ = &typeSectionClass_{
	// Initialize the class constants.
}
