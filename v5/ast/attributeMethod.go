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

func AttributeMethodClass() AttributeMethodClassLike {
	return attributeMethodClassReference()
}

// Constructor Methods

func (c *attributeMethodClass_) AttributeMethod(
	any_ any,
) AttributeMethodLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &attributeMethod_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *attributeMethod_) GetClass() AttributeMethodClassLike {
	return attributeMethodClassReference()
}

// Attribute Methods

func (v *attributeMethod_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type attributeMethod_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type attributeMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func attributeMethodClassReference() *attributeMethodClass_ {
	return attributeMethodClassReference_
}

var attributeMethodClassReference_ = &attributeMethodClass_{
	// Initialize the class constants.
}
