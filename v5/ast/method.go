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
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func MethodClass() MethodClassLike {
	return methodClass()
}

// Constructor Methods

func (c *methodClass_) Method(
	name string,
	parameters col.Sequential[ParameterLike],
	optionalResult ResultLike,
) MethodLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(parameters) {
		panic("The \"parameters\" attribute is required by this class.")
	}
	var instance = &method_{
		// Initialize the instance attributes.
		name_:           name,
		parameters_:     parameters,
		optionalResult_: optionalResult,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *method_) GetClass() MethodClassLike {
	return methodClass()
}

// Attribute Methods

func (v *method_) GetName() string {
	return v.name_
}

func (v *method_) GetParameters() col.Sequential[ParameterLike] {
	return v.parameters_
}

func (v *method_) GetOptionalResult() ResultLike {
	return v.optionalResult_
}

// PROTECTED INTERFACE

// Instance Structure

type method_ struct {
	// Declare the instance attributes.
	name_           string
	parameters_     col.Sequential[ParameterLike]
	optionalResult_ ResultLike
}

// Class Structure

type methodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func methodClass() *methodClass_ {
	return methodClassReference_
}

var methodClassReference_ = &methodClass_{
	// Initialize the class constants.
}
