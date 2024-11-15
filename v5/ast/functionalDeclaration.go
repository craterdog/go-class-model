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

func FunctionalDeclaration() FunctionalDeclarationClassLike {
	return functionalDeclarationReference()
}

// Constructor Methods

func (c *functionalDeclarationClass_) Make(
	declaration DeclarationLike,
	parameters abs.Sequential[ParameterLike],
	result ResultLike,
) FunctionalDeclarationLike {
	if uti.IsUndefined(declaration) {
		panic("The \"declaration\" attribute is required by this class.")
	}
	if uti.IsUndefined(parameters) {
		panic("The \"parameters\" attribute is required by this class.")
	}
	if uti.IsUndefined(result) {
		panic("The \"result\" attribute is required by this class.")
	}
	var instance = &functionalDeclaration_{
		// Initialize the instance attributes.
		declaration_: declaration,
		parameters_:  parameters,
		result_:      result,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *functionalDeclaration_) GetClass() FunctionalDeclarationClassLike {
	return functionalDeclarationReference()
}

// Attribute Methods

func (v *functionalDeclaration_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *functionalDeclaration_) GetParameters() abs.Sequential[ParameterLike] {
	return v.parameters_
}

func (v *functionalDeclaration_) GetResult() ResultLike {
	return v.result_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type functionalDeclaration_ struct {
	// Declare the instance attributes.
	declaration_ DeclarationLike
	parameters_  abs.Sequential[ParameterLike]
	result_      ResultLike
}

// Class Structure

type functionalDeclarationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func functionalDeclarationReference() *functionalDeclarationClass_ {
	return functionalDeclarationReference_
}

var functionalDeclarationReference_ = &functionalDeclarationClass_{
	// Initialize the class constants.
}
