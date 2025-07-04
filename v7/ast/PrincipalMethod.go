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
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func PrincipalMethodClass() PrincipalMethodClassLike {
	return principalMethodClass()
}

// Constructor Methods

func (c *principalMethodClass_) PrincipalMethod(
	method MethodLike,
) PrincipalMethodLike {
	if uti.IsUndefined(method) {
		panic("The \"method\" attribute is required by this class.")
	}
	var instance = &principalMethod_{
		// Initialize the instance attributes.
		method_: method,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *principalMethod_) GetClass() PrincipalMethodClassLike {
	return principalMethodClass()
}

// Attribute Methods

func (v *principalMethod_) GetMethod() MethodLike {
	return v.method_
}

// PROTECTED INTERFACE

// Instance Structure

type principalMethod_ struct {
	// Declare the instance attributes.
	method_ MethodLike
}

// Class Structure

type principalMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func principalMethodClass() *principalMethodClass_ {
	return principalMethodClassReference_
}

var principalMethodClassReference_ = &principalMethodClass_{
	// Initialize the class constants.
}
