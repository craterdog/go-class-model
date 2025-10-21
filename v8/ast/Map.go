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

func MapClass() MapClassLike {
	return mapClass()
}

// Constructor Methods

func (c *mapClass_) Map(
	delimiter1 string,
	delimiter2 string,
	name string,
	delimiter3 string,
) MapLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter3) {
		panic("The \"delimiter3\" attribute is required by this class.")
	}
	var instance = &map_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		delimiter2_: delimiter2,
		name_:       name,
		delimiter3_: delimiter3,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *map_) GetClass() MapClassLike {
	return mapClass()
}

// Attribute Methods

func (v *map_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *map_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *map_) GetName() string {
	return v.name_
}

func (v *map_) GetDelimiter3() string {
	return v.delimiter3_
}

// PROTECTED INTERFACE

// Instance Structure

type map_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	delimiter2_ string
	name_       string
	delimiter3_ string
}

// Class Structure

type mapClass_ struct {
	// Declare the class constants.
}

// Class Reference

func mapClass() *mapClass_ {
	return mapClassReference_
}

var mapClassReference_ = &mapClass_{
	// Initialize the class constants.
}
