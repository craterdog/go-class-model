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

func PrimaryMethod() PrimaryMethodClassLike {
	return primaryMethodReference()
}

// Constructor Methods

func (c *primaryMethodClass_) Make(
	method MethodLike,
) PrimaryMethodLike {
	if uti.IsUndefined(method) {
		panic("The \"method\" attribute is required by this class.")
	}
	var instance = &primaryMethod_{
		// Initialize the instance attributes.
		method_: method,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *primaryMethod_) GetClass() PrimaryMethodClassLike {
	return primaryMethodReference()
}

// Attribute Methods

func (v *primaryMethod_) GetMethod() MethodLike {
	return v.method_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type primaryMethod_ struct {
	// Declare the instance attributes.
	method_ MethodLike
}

// Class Structure

type primaryMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func primaryMethodReference() *primaryMethodClass_ {
	return primaryMethodReference_
}

var primaryMethodReference_ = &primaryMethodClass_{
	// Initialize the class constants.
}
