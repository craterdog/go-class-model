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

func FunctionMethodClass() FunctionMethodClassLike {
	return functionMethodClass()
}

// Constructor Methods

func (c *functionMethodClass_) FunctionMethod(
	name string,
	parameters col.Sequential[ParameterLike],
	result ResultLike,
) FunctionMethodLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(parameters) {
		panic("The \"parameters\" attribute is required by this class.")
	}
	if uti.IsUndefined(result) {
		panic("The \"result\" attribute is required by this class.")
	}
	var instance = &functionMethod_{
		// Initialize the instance attributes.
		name_:       name,
		parameters_: parameters,
		result_:     result,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *functionMethod_) GetClass() FunctionMethodClassLike {
	return functionMethodClass()
}

// Attribute Methods

func (v *functionMethod_) GetName() string {
	return v.name_
}

func (v *functionMethod_) GetParameters() col.Sequential[ParameterLike] {
	return v.parameters_
}

func (v *functionMethod_) GetResult() ResultLike {
	return v.result_
}

// PROTECTED INTERFACE

// Instance Structure

type functionMethod_ struct {
	// Declare the instance attributes.
	name_       string
	parameters_ col.Sequential[ParameterLike]
	result_     ResultLike
}

// Class Structure

type functionMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func functionMethodClass() *functionMethodClass_ {
	return functionMethodClassReference_
}

var functionMethodClassReference_ = &functionMethodClass_{
	// Initialize the class constants.
}
