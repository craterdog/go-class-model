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

package ast

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func TypeSection() TypeSectionClassLike {
	return typeSectionReference()
}

// Constructor Methods

func (c *typeSectionClass_) Make(
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

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *typeSection_) GetClass() TypeSectionClassLike {
	return typeSectionReference()
}

// Attribute Methods

func (v *typeSection_) GetTypeDeclarations() abs.Sequential[TypeDeclarationLike] {
	return v.typeDeclarations_
}

// PROTECTED INTERFACE

// Private Methods

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

func typeSectionReference() *typeSectionClass_ {
	return typeSectionReference_
}

var typeSectionReference_ = &typeSectionClass_{
	// Initialize the class constants.
}
