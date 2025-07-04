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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ConstantSubsectionClass() ConstantSubsectionClassLike {
	return constantSubsectionClass()
}

// Constructor Methods

func (c *constantSubsectionClass_) ConstantSubsection(
	delimiter string,
	constantMethods fra.ListLike[ConstantMethodLike],
) ConstantSubsectionLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(constantMethods) {
		panic("The \"constantMethods\" attribute is required by this class.")
	}
	var instance = &constantSubsection_{
		// Initialize the instance attributes.
		delimiter_:       delimiter,
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

func (v *constantSubsection_) GetDelimiter() string {
	return v.delimiter_
}

func (v *constantSubsection_) GetConstantMethods() fra.ListLike[ConstantMethodLike] {
	return v.constantMethods_
}

// PROTECTED INTERFACE

// Instance Structure

type constantSubsection_ struct {
	// Declare the instance attributes.
	delimiter_       string
	constantMethods_ fra.ListLike[ConstantMethodLike]
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
