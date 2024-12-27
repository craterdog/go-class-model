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
	abs "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func EnumerationClass() EnumerationClassLike {
	return enumerationClass()
}

// Constructor Methods

func (c *enumerationClass_) Enumeration(
	value ValueLike,
	additionalValues abs.Sequential[AdditionalValueLike],
) EnumerationLike {
	if uti.IsUndefined(value) {
		panic("The \"value\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalValues) {
		panic("The \"additionalValues\" attribute is required by this class.")
	}
	var instance = &enumeration_{
		// Initialize the instance attributes.
		value_:            value,
		additionalValues_: additionalValues,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *enumeration_) GetClass() EnumerationClassLike {
	return enumerationClass()
}

// Attribute Methods

func (v *enumeration_) GetValue() ValueLike {
	return v.value_
}

func (v *enumeration_) GetAdditionalValues() abs.Sequential[AdditionalValueLike] {
	return v.additionalValues_
}

// PROTECTED INTERFACE

// Instance Structure

type enumeration_ struct {
	// Declare the instance attributes.
	value_            ValueLike
	additionalValues_ abs.Sequential[AdditionalValueLike]
}

// Class Structure

type enumerationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func enumerationClass() *enumerationClass_ {
	return enumerationClassReference_
}

var enumerationClassReference_ = &enumerationClass_{
	// Initialize the class constants.
}
