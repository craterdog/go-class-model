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
│   Updates to any section other than the Private Methods may be overwritten.  │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import ()

// CLASS INTERFACE

// Access Function

func Array() ArrayClassLike {
	return arrayReference()
}

// Constructor Methods

func (c *arrayClass_) Make() ArrayLike {
	var instance = &array_{
		// Initialize the instance attributes.
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *array_) GetClass() ArrayClassLike {
	return arrayReference()
}

// Attribute Methods

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type array_ struct {
	// Declare the instance attributes.
}

// Class Structure

type arrayClass_ struct {
	// Declare the class constants.
}

// Class Reference

func arrayReference() *arrayClass_ {
	return arrayReference_
}

var arrayReference_ = &arrayClass_{
	// Initialize the class constants.
}
