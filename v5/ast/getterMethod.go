/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func GetterMethodClass() GetterMethodClassLike {
	return getterMethodClass()
}

// Constructor Methods

func (c *getterMethodClass_) GetterMethod(
	name string,
	abstraction AbstractionLike,
) GetterMethodLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(abstraction) {
		panic("The \"abstraction\" attribute is required by this class.")
	}
	var instance = &getterMethod_{
		// Initialize the instance attributes.
		name_:        name,
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

func (v *getterMethod_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// PROTECTED INTERFACE

// Instance Structure

type getterMethod_ struct {
	// Declare the instance attributes.
	name_        string
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
