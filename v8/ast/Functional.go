/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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

func FunctionalClass() FunctionalClassLike {
	return functionalClass()
}

// Constructor Methods

func (c *functionalClass_) Functional(
	delimiter1 string,
	delimiter2 string,
	optionalParameterList ParameterListLike,
	delimiter3 string,
	optionalResult ResultLike,
) FunctionalLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter3) {
		panic("The \"delimiter3\" attribute is required by this class.")
	}
	var instance = &functional_{
		// Initialize the instance attributes.
		delimiter1_:            delimiter1,
		delimiter2_:            delimiter2,
		optionalParameterList_: optionalParameterList,
		delimiter3_:            delimiter3,
		optionalResult_:        optionalResult,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *functional_) GetClass() FunctionalClassLike {
	return functionalClass()
}

// Attribute Methods

func (v *functional_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *functional_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *functional_) GetOptionalParameterList() ParameterListLike {
	return v.optionalParameterList_
}

func (v *functional_) GetDelimiter3() string {
	return v.delimiter3_
}

func (v *functional_) GetOptionalResult() ResultLike {
	return v.optionalResult_
}

// PROTECTED INTERFACE

// Instance Structure

type functional_ struct {
	// Declare the instance attributes.
	delimiter1_            string
	delimiter2_            string
	optionalParameterList_ ParameterListLike
	delimiter3_            string
	optionalResult_        ResultLike
}

// Class Structure

type functionalClass_ struct {
	// Declare the class constants.
}

// Class Reference

func functionalClass() *functionalClass_ {
	return functionalClassReference_
}

var functionalClassReference_ = &functionalClass_{
	// Initialize the class constants.
}
