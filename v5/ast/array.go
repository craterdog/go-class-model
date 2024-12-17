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

import ()

// CLASS INTERFACE

// Access Function

func ArrayClass() ArrayClassLike {
	return arrayClassReference()
}

// Constructor Methods

func (c *arrayClass_) Array() ArrayLike {
	var instance = &array_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *array_) GetClass() ArrayClassLike {
	return arrayClassReference()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type array_ struct {
	// Declare the instance attributes.
}

// Class Structure

type arrayClass_ struct {
	// Declare the class constants.
}

// Class Reference

func arrayClassReference() *arrayClass_ {
	return arrayClassReference_
}

var arrayClassReference_ = &arrayClass_{
	// Initialize the class constants.
}
