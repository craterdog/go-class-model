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

func ConstraintClass() ConstraintClassLike {
	return constraintClass()
}

// Constructor Methods

func (c *constraintClass_) Constraint(
	name string,
	abstraction AbstractionLike,
) ConstraintLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(abstraction) {
		panic("The \"abstraction\" attribute is required by this class.")
	}
	var instance = &constraint_{
		// Initialize the instance attributes.
		name_:        name,
		abstraction_: abstraction,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *constraint_) GetClass() ConstraintClassLike {
	return constraintClass()
}

// Attribute Methods

func (v *constraint_) GetName() string {
	return v.name_
}

func (v *constraint_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// PROTECTED INTERFACE

// Instance Structure

type constraint_ struct {
	// Declare the instance attributes.
	name_        string
	abstraction_ AbstractionLike
}

// Class Structure

type constraintClass_ struct {
	// Declare the class constants.
}

// Class Reference

func constraintClass() *constraintClass_ {
	return constraintClassReference_
}

var constraintClassReference_ = &constraintClass_{
	// Initialize the class constants.
}
