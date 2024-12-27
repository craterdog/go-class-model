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

import (
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func FunctionalDeclarationClass() FunctionalDeclarationClassLike {
	return functionalDeclarationClass()
}

// Constructor Methods

func (c *functionalDeclarationClass_) FunctionalDeclaration(
	declaration DeclarationLike,
	parameters col.Sequential[ParameterLike],
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

// INSTANCE INTERFACE

// Principal Methods

func (v *functionalDeclaration_) GetClass() FunctionalDeclarationClassLike {
	return functionalDeclarationClass()
}

// Attribute Methods

func (v *functionalDeclaration_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *functionalDeclaration_) GetParameters() col.Sequential[ParameterLike] {
	return v.parameters_
}

func (v *functionalDeclaration_) GetResult() ResultLike {
	return v.result_
}

// PROTECTED INTERFACE

// Instance Structure

type functionalDeclaration_ struct {
	// Declare the instance attributes.
	declaration_ DeclarationLike
	parameters_  col.Sequential[ParameterLike]
	result_      ResultLike
}

// Class Structure

type functionalDeclarationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func functionalDeclarationClass() *functionalDeclarationClass_ {
	return functionalDeclarationClassReference_
}

var functionalDeclarationClassReference_ = &functionalDeclarationClass_{
	// Initialize the class constants.
}
