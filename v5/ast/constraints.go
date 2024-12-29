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

func ConstraintsClass() ConstraintsClassLike {
	return constraintsClass()
}

// Constructor Methods

func (c *constraintsClass_) Constraints(
	constraint ConstraintLike,
	additionalConstraints col.Sequential[AdditionalConstraintLike],
) ConstraintsLike {
	if uti.IsUndefined(constraint) {
		panic("The \"constraint\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalConstraints) {
		panic("The \"additionalConstraints\" attribute is required by this class.")
	}
	var instance = &constraints_{
		// Initialize the instance attributes.
		constraint_:            constraint,
		additionalConstraints_: additionalConstraints,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *constraints_) GetClass() ConstraintsClassLike {
	return constraintsClass()
}

// Attribute Methods

func (v *constraints_) GetConstraint() ConstraintLike {
	return v.constraint_
}

func (v *constraints_) GetAdditionalConstraints() col.Sequential[AdditionalConstraintLike] {
	return v.additionalConstraints_
}

// PROTECTED INTERFACE

// Instance Structure

type constraints_ struct {
	// Declare the instance attributes.
	constraint_            ConstraintLike
	additionalConstraints_ col.Sequential[AdditionalConstraintLike]
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
