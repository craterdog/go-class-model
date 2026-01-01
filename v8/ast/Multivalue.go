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

func MultivalueClass() MultivalueClassLike {
	return multivalueClass()
}

// Constructor Methods

func (c *multivalueClass_) Multivalue(
	delimiter1 string,
	parameterList ParameterListLike,
	delimiter2 string,
) MultivalueLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(parameterList) {
		panic("The \"parameterList\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &multivalue_{
		// Initialize the instance attributes.
		delimiter1_:    delimiter1,
		parameterList_: parameterList,
		delimiter2_:    delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *multivalue_) GetClass() MultivalueClassLike {
	return multivalueClass()
}

// Attribute Methods

func (v *multivalue_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *multivalue_) GetParameterList() ParameterListLike {
	return v.parameterList_
}

func (v *multivalue_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type multivalue_ struct {
	// Declare the instance attributes.
	delimiter1_    string
	parameterList_ ParameterListLike
	delimiter2_    string
}

// Class Structure

type multivalueClass_ struct {
	// Declare the class constants.
}

// Class Reference

func multivalueClass() *multivalueClass_ {
	return multivalueClassReference_
}

var multivalueClassReference_ = &multivalueClass_{
	// Initialize the class constants.
}
