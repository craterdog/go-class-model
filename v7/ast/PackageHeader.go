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

func PackageHeaderClass() PackageHeaderClassLike {
	return packageHeaderClass()
}

// Constructor Methods

func (c *packageHeaderClass_) PackageHeader(
	comment string,
	delimiter string,
	name string,
) PackageHeaderLike {
	if uti.IsUndefined(comment) {
		panic("The \"comment\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	var instance = &packageHeader_{
		// Initialize the instance attributes.
		comment_:   comment,
		delimiter_: delimiter,
		name_:      name,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *packageHeader_) GetClass() PackageHeaderClassLike {
	return packageHeaderClass()
}

// Attribute Methods

func (v *packageHeader_) GetComment() string {
	return v.comment_
}

func (v *packageHeader_) GetDelimiter() string {
	return v.delimiter_
}

func (v *packageHeader_) GetName() string {
	return v.name_
}

// PROTECTED INTERFACE

// Instance Structure

type packageHeader_ struct {
	// Declare the instance attributes.
	comment_   string
	delimiter_ string
	name_      string
}

// Class Structure

type packageHeaderClass_ struct {
	// Declare the class constants.
}

// Class Reference

func packageHeaderClass() *packageHeaderClass_ {
	return packageHeaderClassReference_
}

var packageHeaderClassReference_ = &packageHeaderClass_{
	// Initialize the class constants.
}
