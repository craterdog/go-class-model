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
	fra "github.com/craterdog/go-collection-framework/v8"
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func ConstraintsClass() ConstraintsClassLike {
	return constraintsClass()
}

// Constructor Methods

func (c *constraintsClass_) Constraints(
	delimiter1 string,
	constraint ConstraintLike,
	additionalConstraints fra.Sequential[AdditionalConstraintLike],
	delimiter2 string,
) ConstraintsLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(constraint) {
		panic("The \"constraint\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalConstraints) {
		panic("The \"additionalConstraints\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &constraints_{
		// Initialize the instance attributes.
		delimiter1_:            delimiter1,
		constraint_:            constraint,
		additionalConstraints_: additionalConstraints,
		delimiter2_:            delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *constraints_) GetClass() ConstraintsClassLike {
	return constraintsClass()
}

// Attribute Methods

func (v *constraints_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *constraints_) GetConstraint() ConstraintLike {
	return v.constraint_
}

func (v *constraints_) GetAdditionalConstraints() fra.Sequential[AdditionalConstraintLike] {
	return v.additionalConstraints_
}

func (v *constraints_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type constraints_ struct {
	// Declare the instance attributes.
	delimiter1_            string
	constraint_            ConstraintLike
	additionalConstraints_ fra.Sequential[AdditionalConstraintLike]
	delimiter2_            string
}

// Class Structure

type constraintsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func constraintsClass() *constraintsClass_ {
	return constraintsClassReference_
}

var constraintsClassReference_ = &constraintsClass_{
	// Initialize the class constants.
}
