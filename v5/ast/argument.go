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

func ArgumentClass() ArgumentClassLike {
	return argumentClass()
}

// Constructor Methods

func (c *argumentClass_) Argument(
	abstraction AbstractionLike,
) ArgumentLike {
	if uti.IsUndefined(abstraction) {
		panic("The \"abstraction\" attribute is required by this class.")
	}
	var instance = &argument_{
		// Initialize the instance attributes.
		abstraction_: abstraction,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *argument_) GetClass() ArgumentClassLike {
	return argumentClass()
}

// Attribute Methods

func (v *argument_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// PROTECTED INTERFACE

// Instance Structure

type argument_ struct {
	// Declare the instance attributes.
	abstraction_ AbstractionLike
}

// Class Structure

type argumentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func argumentClass() *argumentClass_ {
	return argumentClassReference_
}

var argumentClassReference_ = &argumentClass_{
	// Initialize the class constants.
}
