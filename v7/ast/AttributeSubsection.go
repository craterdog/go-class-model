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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func AttributeSubsectionClass() AttributeSubsectionClassLike {
	return attributeSubsectionClass()
}

// Constructor Methods

func (c *attributeSubsectionClass_) AttributeSubsection(
	delimiter string,
	attributeMethods fra.ListLike[AttributeMethodLike],
) AttributeSubsectionLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(attributeMethods) {
		panic("The \"attributeMethods\" attribute is required by this class.")
	}
	var instance = &attributeSubsection_{
		// Initialize the instance attributes.
		delimiter_:        delimiter,
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

func (v *attributeSubsection_) GetDelimiter() string {
	return v.delimiter_
}

func (v *attributeSubsection_) GetAttributeMethods() fra.ListLike[AttributeMethodLike] {
	return v.attributeMethods_
}

// PROTECTED INTERFACE

// Instance Structure

type attributeSubsection_ struct {
	// Declare the instance attributes.
	delimiter_        string
	attributeMethods_ fra.ListLike[AttributeMethodLike]
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
