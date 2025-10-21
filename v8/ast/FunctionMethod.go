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

func FunctionMethodClass() FunctionMethodClassLike {
	return functionMethodClass()
}

// Constructor Methods

func (c *functionMethodClass_) FunctionMethod(
	name string,
	delimiter1 string,
	optionalParameterList ParameterListLike,
	delimiter2 string,
	result ResultLike,
) FunctionMethodLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(result) {
		panic("The \"result\" attribute is required by this class.")
	}
	var instance = &functionMethod_{
		// Initialize the instance attributes.
		name_:                  name,
		delimiter1_:            delimiter1,
		optionalParameterList_: optionalParameterList,
		delimiter2_:            delimiter2,
		result_:                result,
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

func (v *functionMethod_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *functionMethod_) GetOptionalParameterList() ParameterListLike {
	return v.optionalParameterList_
}

func (v *functionMethod_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *functionMethod_) GetResult() ResultLike {
	return v.result_
}

// PROTECTED INTERFACE

// Instance Structure

type functionMethod_ struct {
	// Declare the instance attributes.
	name_                  string
	delimiter1_            string
	optionalParameterList_ ParameterListLike
	delimiter2_            string
	result_                ResultLike
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
