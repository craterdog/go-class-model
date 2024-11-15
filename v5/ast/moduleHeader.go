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

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func ModuleHeader() ModuleHeaderClassLike {
	return moduleHeaderReference()
}

// Constructor Methods

func (c *moduleHeaderClass_) Make(
	comment string,
	name string,
) ModuleHeaderLike {
	if uti.IsUndefined(comment) {
		panic("The \"comment\" attribute is required by this class.")
	}
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	var instance = &moduleHeader_{
		// Initialize the instance attributes.
		comment_: comment,
		name_:    name,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *moduleHeader_) GetClass() ModuleHeaderClassLike {
	return moduleHeaderReference()
}

// Attribute Methods

func (v *moduleHeader_) GetComment() string {
	return v.comment_
}

func (v *moduleHeader_) GetName() string {
	return v.name_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type moduleHeader_ struct {
	// Declare the instance attributes.
	comment_ string
	name_    string
}

// Class Structure

type moduleHeaderClass_ struct {
	// Declare the class constants.
}

// Class Reference

func moduleHeaderReference() *moduleHeaderClass_ {
	return moduleHeaderReference_
}

var moduleHeaderReference_ = &moduleHeaderClass_{
	// Initialize the class constants.
}
