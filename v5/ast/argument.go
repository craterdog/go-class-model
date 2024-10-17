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

func Argument() ArgumentClassLike {
	return argumentReference()
}

// Constructor Methods

func (c *argumentClass_) Make(
	abstraction AbstractionLike,
) ArgumentLike {
	if uti.IsUndefined(abstraction) {
		panic("The \"abstraction\" attribute is required by this class.")
	}
	var instance = &argument_{
		// Initialize the instance attributes.
		abstraction_: abstraction,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *argument_) GetClass() ArgumentClassLike {
	return argumentReference()
}

// Attribute Methods

func (v *argument_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type argument_ struct {
	// Declare the instance attributes.
	abstraction_ AbstractionLike
}

// Class Structure

type argumentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func argumentReference() *argumentClass_ {
	return argumentReference_
}

var argumentReference_ = &argumentClass_{
	// Initialize the class constants.
}
