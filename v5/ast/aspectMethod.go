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

func AspectMethod() AspectMethodClassLike {
	return aspectMethodReference()
}

// Constructor Methods

func (c *aspectMethodClass_) Make(
	method MethodLike,
) AspectMethodLike {
	if uti.IsUndefined(method) {
		panic("The \"method\" attribute is required by this class.")
	}
	var instance = &aspectMethod_{
		// Initialize the instance attributes.
		method_: method,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *aspectMethod_) GetClass() AspectMethodClassLike {
	return aspectMethodReference()
}

// Attribute Methods

func (v *aspectMethod_) GetMethod() MethodLike {
	return v.method_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type aspectMethod_ struct {
	// Declare the instance attributes.
	method_ MethodLike
}

// Class Structure

type aspectMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func aspectMethodReference() *aspectMethodClass_ {
	return aspectMethodReference_
}

var aspectMethodReference_ = &aspectMethodClass_{
	// Initialize the class constants.
}
