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

func ResultClass() ResultClassLike {
	return resultClass()
}

// Constructor Methods

func (c *resultClass_) Result(
	any_ any,
) ResultLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &result_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *result_) GetClass() ResultClassLike {
	return resultClass()
}

// Attribute Methods

func (v *result_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type result_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type resultClass_ struct {
	// Declare the class constants.
}

// Class Reference

func resultClass() *resultClass_ {
	return resultClassReference_
}

var resultClassReference_ = &resultClass_{
	// Initialize the class constants.
}
