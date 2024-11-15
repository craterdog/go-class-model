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

func PrimitiveDeclarations() PrimitiveDeclarationsClassLike {
	return primitiveDeclarationsReference()
}

// Constructor Methods

func (c *primitiveDeclarationsClass_) Make(
	typeSection TypeSectionLike,
	functionalSection FunctionalSectionLike,
) PrimitiveDeclarationsLike {
	if uti.IsUndefined(typeSection) {
		panic("The \"typeSection\" attribute is required by this class.")
	}
	if uti.IsUndefined(functionalSection) {
		panic("The \"functionalSection\" attribute is required by this class.")
	}
	var instance = &primitiveDeclarations_{
		// Initialize the instance attributes.
		typeSection_:       typeSection,
		functionalSection_: functionalSection,
	}
	return instance

}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *primitiveDeclarations_) GetClass() PrimitiveDeclarationsClassLike {
	return primitiveDeclarationsReference()
}

// Attribute Methods

func (v *primitiveDeclarations_) GetTypeSection() TypeSectionLike {
	return v.typeSection_
}

func (v *primitiveDeclarations_) GetFunctionalSection() FunctionalSectionLike {
	return v.functionalSection_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type primitiveDeclarations_ struct {
	// Declare the instance attributes.
	typeSection_       TypeSectionLike
	functionalSection_ FunctionalSectionLike
}

// Class Structure

type primitiveDeclarationsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func primitiveDeclarationsReference() *primitiveDeclarationsClass_ {
	return primitiveDeclarationsReference_
}

var primitiveDeclarationsReference_ = &primitiveDeclarationsClass_{
	// Initialize the class constants.
}
