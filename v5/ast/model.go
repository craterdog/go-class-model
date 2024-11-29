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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Model() ModelClassLike {
	return modelReference()
}

// Constructor Methods

func (c *modelClass_) Make(
	packageDeclaration PackageDeclarationLike,
	primitiveDeclarations PrimitiveDeclarationsLike,
	interfaceDeclarations InterfaceDeclarationsLike,
) ModelLike {
	if uti.IsUndefined(packageDeclaration) {
		panic("The \"packageDeclaration\" attribute is required by this class.")
	}
	if uti.IsUndefined(primitiveDeclarations) {
		panic("The \"primitiveDeclarations\" attribute is required by this class.")
	}
	if uti.IsUndefined(interfaceDeclarations) {
		panic("The \"interfaceDeclarations\" attribute is required by this class.")
	}
	var instance = &model_{
		// Initialize the instance attributes.
		packageDeclaration_:    packageDeclaration,
		primitiveDeclarations_: primitiveDeclarations,
		interfaceDeclarations_: interfaceDeclarations,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *model_) GetClass() ModelClassLike {
	return modelReference()
}

// Attribute Methods

func (v *model_) GetPackageDeclaration() PackageDeclarationLike {
	return v.packageDeclaration_
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
	packageDeclaration_    PackageDeclarationLike
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
