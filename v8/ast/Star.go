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
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func StarClass() StarClassLike {
	return starClass()
}

// Constructor Methods

func (c *starClass_) Star(
	delimiter string,
) StarLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	var instance = &star_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *star_) GetClass() StarClassLike {
	return starClass()
}

// Attribute Methods

func (v *star_) GetDelimiter() string {
	return v.delimiter_
}

// PROTECTED INTERFACE

// Instance Structure

type star_ struct {
	// Declare the instance attributes.
	delimiter_ string
}

// Class Structure

type starClass_ struct {
	// Declare the class constants.
}

// Class Reference

func starClass() *starClass_ {
	return starClassReference_
}

var starClassReference_ = &starClass_{
	// Initialize the class constants.
}
