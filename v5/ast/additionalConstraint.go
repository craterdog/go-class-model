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

func AdditionalConstraintClass() AdditionalConstraintClassLike {
	return additionalConstraintClass()
}

// Constructor Methods

func (c *additionalConstraintClass_) AdditionalConstraint(
	constraint ConstraintLike,
) AdditionalConstraintLike {
	if uti.IsUndefined(constraint) {
		panic("The \"constraint\" attribute is required by this class.")
	}
	var instance = &additionalConstraint_{
		// Initialize the instance attributes.
		constraint_: constraint,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *additionalConstraint_) GetClass() AdditionalConstraintClassLike {
	return additionalConstraintClass()
}

// Attribute Methods

func (v *additionalConstraint_) GetConstraint() ConstraintLike {
	return v.constraint_
}

// PROTECTED INTERFACE

// Instance Structure

type additionalConstraint_ struct {
	// Declare the instance attributes.
	constraint_ ConstraintLike
}

// Class Structure

type additionalConstraintClass_ struct {
	// Declare the class constants.
}

// Class Reference

func additionalConstraintClass() *additionalConstraintClass_ {
	return additionalConstraintClassReference_
}

var additionalConstraintClassReference_ = &additionalConstraintClass_{
	// Initialize the class constants.
}
