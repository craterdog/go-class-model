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

func AspectMethodClass() AspectMethodClassLike {
	return aspectMethodClassReference()
}

// Constructor Methods

func (c *aspectMethodClass_) AspectMethod(
	method MethodLike,
) AspectMethodLike {
	if uti.IsUndefined(method) {
		panic("The \"method\" attribute is required by this class.")
	}
	var instance = &aspectMethod_{
		// Initialize the instance attributes.
		method_: method,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *aspectMethod_) GetClass() AspectMethodClassLike {
	return aspectMethodClassReference()
}

// Attribute Methods

func (v *aspectMethod_) GetMethod() MethodLike {
	return v.method_
}

// PROTECTED INTERFACE

// Instance Structure

type aspectMethod_ struct {
	// Declare the instance attributes.
	method_ MethodLike
}

// Class Structure

type aspectMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func aspectMethodClassReference() *aspectMethodClass_ {
	return aspectMethodClassReference_
}

var aspectMethodClassReference_ = &aspectMethodClass_{
	// Initialize the class constants.
}
