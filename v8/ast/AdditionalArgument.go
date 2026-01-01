/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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

func AdditionalArgumentClass() AdditionalArgumentClassLike {
	return additionalArgumentClass()
}

// Constructor Methods

func (c *additionalArgumentClass_) AdditionalArgument(
	delimiter string,
	argument ArgumentLike,
) AdditionalArgumentLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(argument) {
		panic("The \"argument\" attribute is required by this class.")
	}
	var instance = &additionalArgument_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		argument_:  argument,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *additionalArgument_) GetClass() AdditionalArgumentClassLike {
	return additionalArgumentClass()
}

// Attribute Methods

func (v *additionalArgument_) GetDelimiter() string {
	return v.delimiter_
}

func (v *additionalArgument_) GetArgument() ArgumentLike {
	return v.argument_
}

// PROTECTED INTERFACE

// Instance Structure

type additionalArgument_ struct {
	// Declare the instance attributes.
	delimiter_ string
	argument_  ArgumentLike
}

// Class Structure

type additionalArgumentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func additionalArgumentClass() *additionalArgumentClass_ {
	return additionalArgumentClassReference_
}

var additionalArgumentClassReference_ = &additionalArgumentClass_{
	// Initialize the class constants.
}
