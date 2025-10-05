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

func NamedClass() NamedClassLike {
	return namedClass()
}

// Constructor Methods

func (c *namedClass_) Named(
	optionalPrefix string,
	name string,
	optionalArguments ArgumentsLike,
) NamedLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	var instance = &named_{
		// Initialize the instance attributes.
		optionalPrefix_:    optionalPrefix,
		name_:              name,
		optionalArguments_: optionalArguments,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *named_) GetClass() NamedClassLike {
	return namedClass()
}

// Attribute Methods

func (v *named_) GetOptionalPrefix() string {
	return v.optionalPrefix_
}

func (v *named_) GetName() string {
	return v.name_
}

func (v *named_) GetOptionalArguments() ArgumentsLike {
	return v.optionalArguments_
}

// PROTECTED INTERFACE

// Instance Structure

type named_ struct {
	// Declare the instance attributes.
	optionalPrefix_    string
	name_              string
	optionalArguments_ ArgumentsLike
}

// Class Structure

type namedClass_ struct {
	// Declare the class constants.
}

// Class Reference

func namedClass() *namedClass_ {
	return namedClassReference_
}

var namedClassReference_ = &namedClass_{
	// Initialize the class constants.
}
