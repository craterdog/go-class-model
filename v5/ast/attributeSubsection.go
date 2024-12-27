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

func AttributeSubsectionClass() AttributeSubsectionClassLike {
	return attributeSubsectionClass()
}

// Constructor Methods

func (c *attributeSubsectionClass_) AttributeSubsection(
	attributeMethods abs.Sequential[AttributeMethodLike],
) AttributeSubsectionLike {
	if uti.IsUndefined(attributeMethods) {
		panic("The \"attributeMethods\" attribute is required by this class.")
	}
	var instance = &attributeSubsection_{
		// Initialize the instance attributes.
		attributeMethods_: attributeMethods,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *attributeSubsection_) GetClass() AttributeSubsectionClassLike {
	return attributeSubsectionClass()
}

// Attribute Methods

func (v *attributeSubsection_) GetAttributeMethods() abs.Sequential[AttributeMethodLike] {
	return v.attributeMethods_
}

// PROTECTED INTERFACE

// Instance Structure

type attributeSubsection_ struct {
	// Declare the instance attributes.
	attributeMethods_ abs.Sequential[AttributeMethodLike]
}

// Class Structure

type attributeSubsectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func attributeSubsectionClass() *attributeSubsectionClass_ {
	return attributeSubsectionClassReference_
}

var attributeSubsectionClassReference_ = &attributeSubsectionClass_{
	// Initialize the class constants.
}
