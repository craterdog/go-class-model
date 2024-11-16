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

func ModuleDeclaration() ModuleDeclarationClassLike {
	return moduleDeclarationReference()
}

// Constructor Methods

func (c *moduleDeclarationClass_) Make(
	legalNotice LegalNoticeLike,
	moduleHeader ModuleHeaderLike,
	moduleImports ModuleImportsLike,
) ModuleDeclarationLike {
	if uti.IsUndefined(legalNotice) {
		panic("The \"legalNotice\" attribute is required by this class.")
	}
	if uti.IsUndefined(moduleHeader) {
		panic("The \"moduleHeader\" attribute is required by this class.")
	}
	if uti.IsUndefined(moduleImports) {
		panic("The \"moduleImports\" attribute is required by this class.")
	}
	var instance = &moduleDeclaration_{
		// Initialize the instance attributes.
		legalNotice_:   legalNotice,
		moduleHeader_:  moduleHeader,
		moduleImports_: moduleImports,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *moduleDeclaration_) GetClass() ModuleDeclarationClassLike {
	return moduleDeclarationReference()
}

// Attribute Methods

func (v *moduleDeclaration_) GetLegalNotice() LegalNoticeLike {
	return v.legalNotice_
}

func (v *moduleDeclaration_) GetModuleHeader() ModuleHeaderLike {
	return v.moduleHeader_
}

func (v *moduleDeclaration_) GetModuleImports() ModuleImportsLike {
	return v.moduleImports_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type moduleDeclaration_ struct {
	// Declare the instance attributes.
	legalNotice_   LegalNoticeLike
	moduleHeader_  ModuleHeaderLike
	moduleImports_ ModuleImportsLike
}

// Class Structure

type moduleDeclarationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func moduleDeclarationReference() *moduleDeclarationClass_ {
	return moduleDeclarationReference_
}

var moduleDeclarationReference_ = &moduleDeclarationClass_{
	// Initialize the class constants.
}
