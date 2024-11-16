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
│                 THIS CLASS FILE WAS AUTOMATICALLY GENERATED.                 │
│                    ANY UPDATES TO IT WILL BE OVERWRITTEN.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Model() ModelClassLike {
	return modelReference()
}

// Constructor Methods

func (c *modelClass_) Make(
	moduleDeclaration ModuleDeclarationLike,
	primitiveDeclarations PrimitiveDeclarationsLike,
	interfaceDeclarations InterfaceDeclarationsLike,
) ModelLike {
	if uti.IsUndefined(moduleDeclaration) {
		panic("The \"moduleDeclaration\" attribute is required by this class.")
	}
	if uti.IsUndefined(primitiveDeclarations) {
		panic("The \"primitiveDeclarations\" attribute is required by this class.")
	}
	if uti.IsUndefined(interfaceDeclarations) {
		panic("The \"interfaceDeclarations\" attribute is required by this class.")
	}
	var instance = &model_{
		// Initialize the instance attributes.
		moduleDeclaration_:     moduleDeclaration,
		primitiveDeclarations_: primitiveDeclarations,
		interfaceDeclarations_: interfaceDeclarations,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *model_) GetClass() ModelClassLike {
	return modelReference()
}

// Attribute Methods

func (v *model_) GetModuleDeclaration() ModuleDeclarationLike {
	return v.moduleDeclaration_
}

func (v *model_) GetPrimitiveDeclarations() PrimitiveDeclarationsLike {
	return v.primitiveDeclarations_
}

func (v *model_) GetInterfaceDeclarations() InterfaceDeclarationsLike {
	return v.interfaceDeclarations_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type model_ struct {
	// Declare the instance attributes.
	moduleDeclaration_     ModuleDeclarationLike
	primitiveDeclarations_ PrimitiveDeclarationsLike
	interfaceDeclarations_ InterfaceDeclarationsLike
}

// Class Structure

type modelClass_ struct {
	// Declare the class constants.
}

// Class Reference

func modelReference() *modelClass_ {
	return modelReference_
}

var modelReference_ = &modelClass_{
	// Initialize the class constants.
}
