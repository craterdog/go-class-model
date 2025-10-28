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

func SetterMethodClass() SetterMethodClassLike {
	return setterMethodClass()
}

// Constructor Methods

func (c *setterMethodClass_) SetterMethod(
	name string,
	delimiter1 string,
	parameter ParameterLike,
	delimiter2 string,
) SetterMethodLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(parameter) {
		panic("The \"parameter\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &setterMethod_{
		// Initialize the instance attributes.
		name_:       name,
		delimiter1_: delimiter1,
		parameter_:  parameter,
		delimiter2_: delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *setterMethod_) GetClass() SetterMethodClassLike {
	return setterMethodClass()
}

// Attribute Methods

func (v *setterMethod_) GetName() string {
	return v.name_
}

func (v *setterMethod_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *setterMethod_) GetParameter() ParameterLike {
	return v.parameter_
}

func (v *setterMethod_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type setterMethod_ struct {
	// Declare the instance attributes.
	name_       string
	delimiter1_ string
	parameter_  ParameterLike
	delimiter2_ string
}

// Class Structure

type setterMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func setterMethodClass() *setterMethodClass_ {
	return setterMethodClassReference_
}

var setterMethodClassReference_ = &setterMethodClass_{
	// Initialize the class constants.
}
