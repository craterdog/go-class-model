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
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func ParameterClass() ParameterClassLike {
	return parameterClass()
}

// Constructor Methods

func (c *parameterClass_) Parameter(
	name string,
	abstraction AbstractionLike,
	delimiter string,
) ParameterLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(abstraction) {
		panic("The \"abstraction\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	var instance = &parameter_{
		// Initialize the instance attributes.
		name_:        name,
		abstraction_: abstraction,
		delimiter_:   delimiter,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parameter_) GetClass() ParameterClassLike {
	return parameterClass()
}

// Attribute Methods

func (v *parameter_) GetName() string {
	return v.name_
}

func (v *parameter_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

func (v *parameter_) GetDelimiter() string {
	return v.delimiter_
}

// PROTECTED INTERFACE

// Instance Structure

type parameter_ struct {
	// Declare the instance attributes.
	name_        string
	abstraction_ AbstractionLike
	delimiter_   string
}

// Class Structure

type parameterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func parameterClass() *parameterClass_ {
	return parameterClassReference_
}

var parameterClassReference_ = &parameterClass_{
	// Initialize the class constants.
}
