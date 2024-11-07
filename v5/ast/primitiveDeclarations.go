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

// CLASS INTERFACE

// Access Function

func PrimitiveDeclarations() PrimitiveDeclarationsClassLike {
	return primitiveDeclarationsReference()
}

// Constructor Methods

func (c *primitiveDeclarationsClass_) Make(
	optionalTypeSection TypeSectionLike,
	optionalFunctionalSection FunctionalSectionLike,
) PrimitiveDeclarationsLike {
	var instance = &primitiveDeclarations_{
		// Initialize the instance attributes.
		optionalTypeSection_:       optionalTypeSection,
		optionalFunctionalSection_: optionalFunctionalSection,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *primitiveDeclarations_) GetClass() PrimitiveDeclarationsClassLike {
	return primitiveDeclarationsReference()
}

// Attribute Methods

func (v *primitiveDeclarations_) GetOptionalTypeSection() TypeSectionLike {
	return v.optionalTypeSection_
}

func (v *primitiveDeclarations_) GetOptionalFunctionalSection() FunctionalSectionLike {
	return v.optionalFunctionalSection_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type primitiveDeclarations_ struct {
	// Declare the instance attributes.
	optionalTypeSection_       TypeSectionLike
	optionalFunctionalSection_ FunctionalSectionLike
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
