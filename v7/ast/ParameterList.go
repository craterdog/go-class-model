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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ParameterListClass() ParameterListClassLike {
	return parameterListClass()
}

// Constructor Methods

func (c *parameterListClass_) ParameterList(
	parameters fra.ListLike[ParameterLike],
) ParameterListLike {
	if uti.IsUndefined(parameters) {
		panic("The \"parameters\" attribute is required by this class.")
	}
	var instance = &parameterList_{
		// Initialize the instance attributes.
		parameters_: parameters,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parameterList_) GetClass() ParameterListClassLike {
	return parameterListClass()
}

// Attribute Methods

func (v *parameterList_) GetParameters() fra.ListLike[ParameterLike] {
	return v.parameters_
}

// PROTECTED INTERFACE

// Instance Structure

type parameterList_ struct {
	// Declare the instance attributes.
	parameters_ fra.ListLike[ParameterLike]
}

// Class Structure

type parameterListClass_ struct {
	// Declare the class constants.
}

// Class Reference

func parameterListClass() *parameterListClass_ {
	return parameterListClassReference_
}

var parameterListClassReference_ = &parameterListClass_{
	// Initialize the class constants.
}
