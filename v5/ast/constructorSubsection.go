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
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func ConstructorSubsectionClass() ConstructorSubsectionClassLike {
	return constructorSubsectionClass()
}

// Constructor Methods

func (c *constructorSubsectionClass_) ConstructorSubsection(
	constructorMethods col.Sequential[ConstructorMethodLike],
) ConstructorSubsectionLike {
	if uti.IsUndefined(constructorMethods) {
		panic("The \"constructorMethods\" attribute is required by this class.")
	}
	var instance = &constructorSubsection_{
		// Initialize the instance attributes.
		constructorMethods_: constructorMethods,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *constructorSubsection_) GetClass() ConstructorSubsectionClassLike {
	return constructorSubsectionClass()
}

// Attribute Methods

func (v *constructorSubsection_) GetConstructorMethods() col.Sequential[ConstructorMethodLike] {
	return v.constructorMethods_
}

// PROTECTED INTERFACE

// Instance Structure

type constructorSubsection_ struct {
	// Declare the instance attributes.
	constructorMethods_ col.Sequential[ConstructorMethodLike]
}

// Class Structure

type constructorSubsectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func constructorSubsectionClass() *constructorSubsectionClass_ {
	return constructorSubsectionClassReference_
}

var constructorSubsectionClassReference_ = &constructorSubsectionClass_{
	// Initialize the class constants.
}
