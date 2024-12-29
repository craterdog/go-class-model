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
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func MultivalueClass() MultivalueClassLike {
	return multivalueClass()
}

// Constructor Methods

func (c *multivalueClass_) Multivalue(
	parameters col.Sequential[ParameterLike],
) MultivalueLike {
	if uti.IsUndefined(parameters) {
		panic("The \"parameters\" attribute is required by this class.")
	}
	var instance = &multivalue_{
		// Initialize the instance attributes.
		parameters_: parameters,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *multivalue_) GetClass() MultivalueClassLike {
	return multivalueClass()
}

// Attribute Methods

func (v *multivalue_) GetParameters() col.Sequential[ParameterLike] {
	return v.parameters_
}

// PROTECTED INTERFACE

// Instance Structure

type multivalue_ struct {
	// Declare the instance attributes.
	parameters_ col.Sequential[ParameterLike]
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
