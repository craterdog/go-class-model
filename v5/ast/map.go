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

func MapClass() MapClassLike {
	return mapClass()
}

// Constructor Methods

func (c *mapClass_) Map(
	name string,
) MapLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	var instance = &map_{
		// Initialize the instance attributes.
		name_: name,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *map_) GetClass() MapClassLike {
	return mapClass()
}

// Attribute Methods

func (v *map_) GetName() string {
	return v.name_
}

// PROTECTED INTERFACE

// Instance Structure

type map_ struct {
	// Declare the instance attributes.
	name_ string
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
