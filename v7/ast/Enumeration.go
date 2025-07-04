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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func EnumerationClass() EnumerationClassLike {
	return enumerationClass()
}

// Constructor Methods

func (c *enumerationClass_) Enumeration(
	delimiter1 string,
	delimiter2 string,
	value ValueLike,
	additionalValues fra.ListLike[AdditionalValueLike],
	delimiter3 string,
) EnumerationLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(value) {
		panic("The \"value\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalValues) {
		panic("The \"additionalValues\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter3) {
		panic("The \"delimiter3\" attribute is required by this class.")
	}
	var instance = &enumeration_{
		// Initialize the instance attributes.
		delimiter1_:       delimiter1,
		delimiter2_:       delimiter2,
		value_:            value,
		additionalValues_: additionalValues,
		delimiter3_:       delimiter3,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *enumeration_) GetClass() EnumerationClassLike {
	return enumerationClass()
}

// Attribute Methods

func (v *enumeration_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *enumeration_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *enumeration_) GetValue() ValueLike {
	return v.value_
}

func (v *enumeration_) GetAdditionalValues() fra.ListLike[AdditionalValueLike] {
	return v.additionalValues_
}

func (v *enumeration_) GetDelimiter3() string {
	return v.delimiter3_
}

// PROTECTED INTERFACE

// Instance Structure

type enumeration_ struct {
	// Declare the instance attributes.
	delimiter1_       string
	delimiter2_       string
	value_            ValueLike
	additionalValues_ fra.ListLike[AdditionalValueLike]
	delimiter3_       string
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
