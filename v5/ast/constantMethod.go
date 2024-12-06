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

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func ConstantMethodClass() ConstantMethodClassLike {
	return constantMethodClassReference()
}

// Constructor Methods

func (c *constantMethodClass_) Make(
	name string,
	abstraction AbstractionLike,
) ConstantMethodLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(abstraction) {
		panic("The \"abstraction\" attribute is required by this class.")
	}
	var instance = &constantMethod_{
		// Initialize the instance attributes.
		name_:        name,
		abstraction_: abstraction,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *constantMethod_) GetClass() ConstantMethodClassLike {
	return constantMethodClassReference()
}

// Attribute Methods

func (v *constantMethod_) GetName() string {
	return v.name_
}

func (v *constantMethod_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type constantMethod_ struct {
	// Declare the instance attributes.
	name_        string
	abstraction_ AbstractionLike
}

// Class Structure

type constantMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func constantMethodClassReference() *constantMethodClass_ {
	return constantMethodClassReference_
}

var constantMethodClassReference_ = &constantMethodClass_{
	// Initialize the class constants.
}
