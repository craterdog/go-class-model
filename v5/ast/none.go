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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func NoneClass() NoneClassLike {
	return noneClass()
}

// Constructor Methods

func (c *noneClass_) None(
	newline string,
) NoneLike {
	if uti.IsUndefined(newline) {
		panic("The \"newline\" attribute is required by this class.")
	}
	var instance = &none_{
		// Initialize the instance attributes.
		newline_: newline,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *none_) GetClass() NoneClassLike {
	return noneClass()
}

// Attribute Methods

func (v *none_) GetNewline() string {
	return v.newline_
}

// PROTECTED INTERFACE

// Instance Structure

type none_ struct {
	// Declare the instance attributes.
	newline_ string
}

// Class Structure

type noneClass_ struct {
	// Declare the class constants.
}

// Class Reference

func noneClass() *noneClass_ {
	return noneClassReference_
}

var noneClassReference_ = &noneClass_{
	// Initialize the class constants.
}
