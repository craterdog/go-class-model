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
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func DotsClass() DotsClassLike {
	return dotsClass()
}

// Constructor Methods

func (c *dotsClass_) Dots(
	delimiter string,
) DotsLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	var instance = &dots_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *dots_) GetClass() DotsClassLike {
	return dotsClass()
}

// Attribute Methods

func (v *dots_) GetDelimiter() string {
	return v.delimiter_
}

// PROTECTED INTERFACE

// Instance Structure

type dots_ struct {
	// Declare the instance attributes.
	delimiter_ string
}

// Class Structure

type dotsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func dotsClass() *dotsClass_ {
	return dotsClassReference_
}

var dotsClassReference_ = &dotsClass_{
	// Initialize the class constants.
}
