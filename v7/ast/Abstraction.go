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

func AbstractionClass() AbstractionClassLike {
	return abstractionClass()
}

// Constructor Methods

func (c *abstractionClass_) Abstraction(
	optionalWrapper WrapperLike,
	type_ TypeLike,
) AbstractionLike {
	if uti.IsUndefined(type_) {
		panic("The \"type\" attribute is required by this class.")
	}
	var instance = &abstraction_{
		// Initialize the instance attributes.
		optionalWrapper_: optionalWrapper,
		type_:            type_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *abstraction_) GetClass() AbstractionClassLike {
	return abstractionClass()
}

// Attribute Methods

func (v *abstraction_) GetOptionalWrapper() WrapperLike {
	return v.optionalWrapper_
}

func (v *abstraction_) GetType() TypeLike {
	return v.type_
}

// PROTECTED INTERFACE

// Instance Structure

type abstraction_ struct {
	// Declare the instance attributes.
	optionalWrapper_ WrapperLike
	type_            TypeLike
}

// Class Structure

type abstractionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func abstractionClass() *abstractionClass_ {
	return abstractionClassReference_
}

var abstractionClassReference_ = &abstractionClass_{
	// Initialize the class constants.
}
