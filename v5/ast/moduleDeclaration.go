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
	notice NoticeLike,
	header HeaderLike,
	optionalImports ImportsLike,
) ModuleDeclarationLike {
	if uti.IsUndefined(notice) {
		panic("The \"notice\" attribute is required by this class.")
	}
	if uti.IsUndefined(header) {
		panic("The \"header\" attribute is required by this class.")
	}
	var instance = &moduleDeclaration_{
		// Initialize the instance attributes.
		notice_:          notice,
		header_:          header,
		optionalImports_: optionalImports,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *moduleDeclaration_) GetClass() ModuleDeclarationClassLike {
	return moduleDeclarationReference()
}

// Attribute Methods

func (v *moduleDeclaration_) GetNotice() NoticeLike {
	return v.notice_
}

func (v *moduleDeclaration_) GetHeader() HeaderLike {
	return v.header_
}

func (v *moduleDeclaration_) GetOptionalImports() ImportsLike {
	return v.optionalImports_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type moduleDeclaration_ struct {
	// Declare the instance attributes.
	notice_          NoticeLike
	header_          HeaderLike
	optionalImports_ ImportsLike
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
