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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ValueClass() ValueClassLike {
	return valueClass()
}

// Constructor Methods

func (c *valueClass_) Value(
	name string,
	abstraction AbstractionLike,
	delimiter1 string,
	delimiter2 string,
) ValueLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(abstraction) {
		panic("The \"abstraction\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &value_{
		// Initialize the instance attributes.
		name_:        name,
		abstraction_: abstraction,
		delimiter1_:  delimiter1,
		delimiter2_:  delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *value_) GetClass() ValueClassLike {
	return valueClass()
}

// Attribute Methods

func (v *value_) GetName() string {
	return v.name_
}

func (v *value_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

func (v *value_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *value_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type value_ struct {
	// Declare the instance attributes.
	name_        string
	abstraction_ AbstractionLike
	delimiter1_  string
	delimiter2_  string
}

// Class Structure

type valueClass_ struct {
	// Declare the class constants.
}

// Class Reference

func valueClass() *valueClass_ {
	return valueClassReference_
}

var valueClassReference_ = &valueClass_{
	// Initialize the class constants.
}
