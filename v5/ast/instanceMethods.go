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

func InstanceMethods() InstanceMethodsClassLike {
	return instanceMethodsReference()
}

// Constructor Methods

func (c *instanceMethodsClass_) Make(
	principalSubsection PrincipalSubsectionLike,
	optionalAttributeSubsection AttributeSubsectionLike,
	optionalAspectSubsection AspectSubsectionLike,
) InstanceMethodsLike {
	if uti.IsUndefined(principalSubsection) {
		panic("The \"principalSubsection\" attribute is required by this class.")
	}
	var instance = &instanceMethods_{
		// Initialize the instance attributes.
		principalSubsection_:         principalSubsection,
		optionalAttributeSubsection_: optionalAttributeSubsection,
		optionalAspectSubsection_:    optionalAspectSubsection,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *instanceMethods_) GetClass() InstanceMethodsClassLike {
	return instanceMethodsReference()
}

// Attribute Methods

func (v *instanceMethods_) GetPrincipalSubsection() PrincipalSubsectionLike {
	return v.principalSubsection_
}

func (v *instanceMethods_) GetOptionalAttributeSubsection() AttributeSubsectionLike {
	return v.optionalAttributeSubsection_
}

func (v *instanceMethods_) GetOptionalAspectSubsection() AspectSubsectionLike {
	return v.optionalAspectSubsection_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type instanceMethods_ struct {
	// Declare the instance attributes.
	principalSubsection_         PrincipalSubsectionLike
	optionalAttributeSubsection_ AttributeSubsectionLike
	optionalAspectSubsection_    AspectSubsectionLike
}

// Class Structure

type instanceMethodsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func instanceMethodsReference() *instanceMethodsClass_ {
	return instanceMethodsReference_
}

var instanceMethodsReference_ = &instanceMethodsClass_{
	// Initialize the class constants.
}
