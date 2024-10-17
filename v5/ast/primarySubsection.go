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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func PrimarySubsection() PrimarySubsectionClassLike {
	return primarySubsectionReference()
}

// Constructor Methods

func (c *primarySubsectionClass_) Make(
	primaryMethods abs.Sequential[PrimaryMethodLike],
) PrimarySubsectionLike {
	if uti.IsUndefined(primaryMethods) {
		panic("The \"primaryMethods\" attribute is required by this class.")
	}
	var instance = &primarySubsection_{
		// Initialize the instance attributes.
		primaryMethods_: primaryMethods,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *primarySubsection_) GetClass() PrimarySubsectionClassLike {
	return primarySubsectionReference()
}

// Attribute Methods

func (v *primarySubsection_) GetPrimaryMethods() abs.Sequential[PrimaryMethodLike] {
	return v.primaryMethods_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type primarySubsection_ struct {
	// Declare the instance attributes.
	primaryMethods_ abs.Sequential[PrimaryMethodLike]
}

// Class Structure

type primarySubsectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func primarySubsectionReference() *primarySubsectionClass_ {
	return primarySubsectionReference_
}

var primarySubsectionReference_ = &primarySubsectionClass_{
	// Initialize the class constants.
}
