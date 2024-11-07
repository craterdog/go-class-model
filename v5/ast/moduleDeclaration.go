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

func ModuleDeclaration() ModuleDeclarationClassLike {
	return moduleDeclarationReference()
}

// Constructor Methods

func (c *moduleDeclarationClass_) Make(
	legalNotice LegalNoticeLike,
	moduleHeader ModuleHeaderLike,
	optionalModuleImports ModuleImportsLike,
) ModuleDeclarationLike {
	if uti.IsUndefined(legalNotice) {
		panic("The \"legalNotice\" attribute is required by this class.")
	}
	if uti.IsUndefined(moduleHeader) {
		panic("The \"moduleHeader\" attribute is required by this class.")
	}
	var instance = &moduleDeclaration_{
		// Initialize the instance attributes.
		legalNotice_:           legalNotice,
		moduleHeader_:          moduleHeader,
		optionalModuleImports_: optionalModuleImports,
	}
	return instance

}

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

func (v *moduleDeclaration_) GetOptionalModuleImports() ModuleImportsLike {
	return v.optionalModuleImports_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type moduleDeclaration_ struct {
	// Declare the instance attributes.
	legalNotice_           LegalNoticeLike
	moduleHeader_          ModuleHeaderLike
	optionalModuleImports_ ModuleImportsLike
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
