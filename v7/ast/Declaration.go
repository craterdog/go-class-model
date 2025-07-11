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

func DeclarationClass() DeclarationClassLike {
	return declarationClass()
}

// Constructor Methods

func (c *declarationClass_) Declaration(
	comment string,
	delimiter string,
	name string,
	optionalConstraints ConstraintsLike,
) DeclarationLike {
	if uti.IsUndefined(comment) {
		panic("The \"comment\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	var instance = &declaration_{
		// Initialize the instance attributes.
		comment_:             comment,
		delimiter_:           delimiter,
		name_:                name,
		optionalConstraints_: optionalConstraints,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *declaration_) GetClass() DeclarationClassLike {
	return declarationClass()
}

// Attribute Methods

func (v *declaration_) GetComment() string {
	return v.comment_
}

func (v *declaration_) GetDelimiter() string {
	return v.delimiter_
}

func (v *declaration_) GetName() string {
	return v.name_
}

func (v *declaration_) GetOptionalConstraints() ConstraintsLike {
	return v.optionalConstraints_
}

// PROTECTED INTERFACE

// Instance Structure

type declaration_ struct {
	// Declare the instance attributes.
	comment_             string
	delimiter_           string
	name_                string
	optionalConstraints_ ConstraintsLike
}

// Class Structure

type declarationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func declarationClass() *declarationClass_ {
	return declarationClassReference_
}

var declarationClassReference_ = &declarationClass_{
	// Initialize the class constants.
}
