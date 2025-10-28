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
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func ArrayClass() ArrayClassLike {
	return arrayClass()
}

// Constructor Methods

func (c *arrayClass_) Array(
	delimiter1 string,
	delimiter2 string,
) ArrayLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &array_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		delimiter2_: delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *array_) GetClass() ArrayClassLike {
	return arrayClass()
}

// Attribute Methods

func (v *array_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *array_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type array_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	delimiter2_ string
}

// Class Structure

type arrayClass_ struct {
	// Declare the class constants.
}

// Class Reference

func arrayClass() *arrayClass_ {
	return arrayClassReference_
}

var arrayClassReference_ = &arrayClass_{
	// Initialize the class constants.
}
