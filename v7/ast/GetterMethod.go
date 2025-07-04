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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func GetterMethodClass() GetterMethodClassLike {
	return getterMethodClass()
}

// Constructor Methods

func (c *getterMethodClass_) GetterMethod(
	name string,
	delimiter1 string,
	delimiter2 string,
	abstraction AbstractionLike,
) GetterMethodLike {
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
	var instance = &getterMethod_{
		// Initialize the instance attributes.
		name_:        name,
		delimiter1_:  delimiter1,
		delimiter2_:  delimiter2,
		abstraction_: abstraction,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *getterMethod_) GetClass() GetterMethodClassLike {
	return getterMethodClass()
}

// Attribute Methods

func (v *getterMethod_) GetName() string {
	return v.name_
}

func (v *getterMethod_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *getterMethod_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *getterMethod_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// PROTECTED INTERFACE

// Instance Structure

type getterMethod_ struct {
	// Declare the instance attributes.
	name_        string
	delimiter1_  string
	delimiter2_  string
	abstraction_ AbstractionLike
}

// Class Structure

type getterMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func getterMethodClass() *getterMethodClass_ {
	return getterMethodClassReference_
}

var getterMethodClassReference_ = &getterMethodClass_{
	// Initialize the class constants.
}
