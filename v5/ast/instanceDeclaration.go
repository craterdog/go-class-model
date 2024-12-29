/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
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

func InstanceDeclarationClass() InstanceDeclarationClassLike {
	return instanceDeclarationClass()
}

// Constructor Methods

func (c *instanceDeclarationClass_) InstanceDeclaration(
	declaration DeclarationLike,
	instanceMethods InstanceMethodsLike,
) InstanceDeclarationLike {
	if uti.IsUndefined(declaration) {
		panic("The \"declaration\" attribute is required by this class.")
	}
	if uti.IsUndefined(instanceMethods) {
		panic("The \"instanceMethods\" attribute is required by this class.")
	}
	var instance = &instanceDeclaration_{
		// Initialize the instance attributes.
		declaration_:     declaration,
		instanceMethods_: instanceMethods,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *instanceDeclaration_) GetClass() InstanceDeclarationClassLike {
	return instanceDeclarationClass()
}

// Attribute Methods

func (v *instanceDeclaration_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *instanceDeclaration_) GetInstanceMethods() InstanceMethodsLike {
	return v.instanceMethods_
}

// PROTECTED INTERFACE

// Instance Structure

type instanceDeclaration_ struct {
	// Declare the instance attributes.
	declaration_     DeclarationLike
	instanceMethods_ InstanceMethodsLike
}

// Class Structure

type instanceDeclarationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func instanceDeclarationClass() *instanceDeclarationClass_ {
	return instanceDeclarationClassReference_
}

var instanceDeclarationClassReference_ = &instanceDeclarationClass_{
	// Initialize the class constants.
}
