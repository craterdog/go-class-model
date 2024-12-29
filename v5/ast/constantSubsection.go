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

func ConstantSubsectionClass() ConstantSubsectionClassLike {
	return constantSubsectionClass()
}

// Constructor Methods

func (c *constantSubsectionClass_) ConstantSubsection(
	constantMethods col.Sequential[ConstantMethodLike],
) ConstantSubsectionLike {
	if uti.IsUndefined(constantMethods) {
		panic("The \"constantMethods\" attribute is required by this class.")
	}
	var instance = &constantSubsection_{
		// Initialize the instance attributes.
		constantMethods_: constantMethods,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *constantSubsection_) GetClass() ConstantSubsectionClassLike {
	return constantSubsectionClass()
}

// Attribute Methods

func (v *constantSubsection_) GetConstantMethods() col.Sequential[ConstantMethodLike] {
	return v.constantMethods_
}

// PROTECTED INTERFACE

// Instance Structure

type constantSubsection_ struct {
	// Declare the instance attributes.
	constantMethods_ col.Sequential[ConstantMethodLike]
}

// Class Structure

type constantSubsectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func constantSubsectionClass() *constantSubsectionClass_ {
	return constantSubsectionClassReference_
}

var constantSubsectionClassReference_ = &constantSubsectionClass_{
	// Initialize the class constants.
}
