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

func ConstructorMethodClass() ConstructorMethodClassLike {
	return constructorMethodClass()
}

// Constructor Methods

func (c *constructorMethodClass_) ConstructorMethod(
	name string,
	delimiter1 string,
	optionalParameterList ParameterListLike,
	delimiter2 string,
	abstraction AbstractionLike,
) ConstructorMethodLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(abstraction) {
		panic("The \"abstraction\" attribute is required by this class.")
	}
	var instance = &constructorMethod_{
		// Initialize the instance attributes.
		name_:                  name,
		delimiter1_:            delimiter1,
		optionalParameterList_: optionalParameterList,
		delimiter2_:            delimiter2,
		abstraction_:           abstraction,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *constructorMethod_) GetClass() ConstructorMethodClassLike {
	return constructorMethodClass()
}

// Attribute Methods

func (v *constructorMethod_) GetName() string {
	return v.name_
}

func (v *constructorMethod_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *constructorMethod_) GetOptionalParameterList() ParameterListLike {
	return v.optionalParameterList_
}

func (v *constructorMethod_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *constructorMethod_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// PROTECTED INTERFACE

// Instance Structure

type constructorMethod_ struct {
	// Declare the instance attributes.
	name_                  string
	delimiter1_            string
	optionalParameterList_ ParameterListLike
	delimiter2_            string
	abstraction_           AbstractionLike
}

// Class Structure

type constructorMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func constructorMethodClass() *constructorMethodClass_ {
	return constructorMethodClassReference_
}

var constructorMethodClassReference_ = &constructorMethodClass_{
	// Initialize the class constants.
}
