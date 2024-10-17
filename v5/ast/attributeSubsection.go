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

func AttributeSubsection() AttributeSubsectionClassLike {
	return attributeSubsectionReference()
}

// Constructor Methods

func (c *attributeSubsectionClass_) Make(
	attributeMethods abs.Sequential[AttributeMethodLike],
) AttributeSubsectionLike {
	if uti.IsUndefined(attributeMethods) {
		panic("The \"attributeMethods\" attribute is required by this class.")
	}
	var instance = &attributeSubsection_{
		// Initialize the instance attributes.
		attributeMethods_: attributeMethods,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *attributeSubsection_) GetClass() AttributeSubsectionClassLike {
	return attributeSubsectionReference()
}

// Attribute Methods

func (v *attributeSubsection_) GetAttributeMethods() abs.Sequential[AttributeMethodLike] {
	return v.attributeMethods_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type attributeSubsection_ struct {
	// Declare the instance attributes.
	attributeMethods_ abs.Sequential[AttributeMethodLike]
}

// Class Structure

type attributeSubsectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func attributeSubsectionReference() *attributeSubsectionClass_ {
	return attributeSubsectionReference_
}

var attributeSubsectionReference_ = &attributeSubsectionClass_{
	// Initialize the class constants.
}
