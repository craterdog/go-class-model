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
│ Updates to any section other than the Methodical Methods may be overwritten. │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	ast "github.com/craterdog/go-class-model/v7/ast"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func FormatterClass() FormatterClassLike {
	return formatterClass()
}

// Constructor Methods

func (c *formatterClass_) Formatter() FormatterLike {
	var instance = &formatter_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	instance.visitor_ = VisitorClass().Visitor(instance)
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *formatter_) GetClass() FormatterClassLike {
	return formatterClass()
}

func (v *formatter_) FormatModel(model ast.ModelLike) string {
	v.visitor_.VisitModel(model)
	return v.getResult()
}

// Methodical Methods

func (v *formatter_) ProcessComment(
	comment string,
) {
	v.appendString(comment)
}

func (v *formatter_) ProcessDelimiter(
	delimiter string,
) {
	v.appendString(delimiter)
}

func (v *formatter_) ProcessName(
	name string,
) {
	v.appendString(name)
}

func (v *formatter_) ProcessPath(
	path string,
) {
	v.appendString(path)
}

func (v *formatter_) ProcessPrefix(
	prefix string,
) {
	v.appendString(prefix)
}

func (v *formatter_) ProcessSpace(
	space string,
) {
	v.appendString(space)
}

func (v *formatter_) ProcessAdditionalArgumentSlot(
	additionalArgument ast.AdditionalArgumentLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessAdditionalConstraintSlot(
	additionalConstraint ast.AdditionalConstraintLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessAspectDeclaration(
	aspectDeclaration ast.AspectDeclarationLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PostprocessAspectDeclaration(
	aspectDeclaration ast.AspectDeclarationLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessAspectDeclarationSlot(
	aspectDeclaration ast.AspectDeclarationLike,
	slot_ uint,
) {
	switch slot_ {
	case 1, 2:
		v.appendString(" ")
	case 4:
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessAspectInterface(
	aspectInterface ast.AspectInterfaceLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessAspectMethod(
	aspectMethod ast.AspectMethodLike,
	index_ uint,
	count_ uint,
) {
	v.depth_++
}

func (v *formatter_) PostprocessAspectMethod(
	aspectMethod ast.AspectMethodLike,
	index_ uint,
	count_ uint,
) {
	v.depth_--
}

func (v *formatter_) ProcessAspectSectionSlot(
	aspectSection ast.AspectSectionLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessAspectSubsection(
	aspectSubsection ast.AspectSubsectionLike,
	index_ uint,
	count_ uint,
) {
	v.depth_++
	v.appendNewline()
}

func (v *formatter_) PostprocessAspectSubsection(
	aspectSubsection ast.AspectSubsectionLike,
	index_ uint,
	count_ uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PreprocessAttributeSubsection(
	attributeSubsection ast.AttributeSubsectionLike,
	index_ uint,
	count_ uint,
) {
	v.depth_++
	v.appendNewline()
}

func (v *formatter_) PostprocessAttributeSubsection(
	attributeSubsection ast.AttributeSubsectionLike,
	index_ uint,
	count_ uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PreprocessClassDeclaration(
	classDeclaration ast.ClassDeclarationLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PostprocessClassDeclaration(
	classDeclaration ast.ClassDeclarationLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessClassDeclarationSlot(
	classDeclaration ast.ClassDeclarationLike,
	slot_ uint,
) {
	switch slot_ {
	case 1, 2:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessClassSectionSlot(
	classSection ast.ClassSectionLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessConstantMethod(
	constantMethod ast.ConstantMethodLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessConstantMethodSlot(
	constantMethod ast.ConstantMethodLike,
	slot_ uint,
) {
	switch slot_ {
	case 3:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessConstantSubsection(
	constantSubsection ast.ConstantSubsectionLike,
	index_ uint,
	count_ uint,
) {
	v.depth_++
	v.appendNewline()
}

func (v *formatter_) PostprocessConstantSubsection(
	constantSubsection ast.ConstantSubsectionLike,
	index_ uint,
	count_ uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) ProcessConstraintSlot(
	constraint ast.ConstraintLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessConstructorMethod(
	constructorMethod ast.ConstructorMethodLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessConstructorMethodSlot(
	constructorMethod ast.ConstructorMethodLike,
	slot_ uint,
) {
	switch slot_ {
	case 4:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessConstructorSubsection(
	constructorSubsection ast.ConstructorSubsectionLike,
	index_ uint,
	count_ uint,
) {
	v.depth_++
	v.appendNewline()
}

func (v *formatter_) PostprocessConstructorSubsection(
	constructorSubsection ast.ConstructorSubsectionLike,
	index_ uint,
	count_ uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) ProcessDeclarationSlot(
	declaration ast.DeclarationLike,
	slot_ uint,
) {
	switch slot_ {
	case 2:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessEnumeration(
	enumeration ast.EnumerationLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PostprocessEnumeration(
	enumeration ast.EnumerationLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessEnumerationSlot(
	enumeration ast.EnumerationLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendString(" ")
	case 2:
		v.depth_++
		v.appendNewline()
	case 4:
		v.depth_--
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessFunctionMethod(
	functionMethod ast.FunctionMethodLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessFunctionSubsection(
	functionSubsection ast.FunctionSubsectionLike,
	index_ uint,
	count_ uint,
) {
	v.depth_++
	v.appendNewline()
}

func (v *formatter_) PostprocessFunctionSubsection(
	functionSubsection ast.FunctionSubsectionLike,
	index_ uint,
	count_ uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PreprocessFunctionalDeclaration(
	functionalDeclaration ast.FunctionalDeclarationLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PostprocessFunctionalDeclaration(
	functionalDeclaration ast.FunctionalDeclarationLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessFunctionalDeclarationSlot(
	functionalDeclaration ast.FunctionalDeclarationLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessFunctionalSectionSlot(
	functionalSection ast.FunctionalSectionLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessGetterMethod(
	getterMethod ast.GetterMethodLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessGetterMethodSlot(
	getterMethod ast.GetterMethodLike,
	slot_ uint,
) {
	switch slot_ {
	case 3:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessImportedPackage(
	importedPackage ast.ImportedPackageLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessImportedPackageSlot(
	importedPackage ast.ImportedPackageLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessImportList(
	parameterList ast.ImportListLike,
	index_ uint,
	count_ uint,
) {
	v.depth_++
}

func (v *formatter_) PostprocessImportList(
	parameterList ast.ImportListLike,
	index_ uint,
	count_ uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PreprocessInstanceDeclaration(
	instanceDeclaration ast.InstanceDeclarationLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PostprocessInstanceDeclaration(
	instanceDeclaration ast.InstanceDeclarationLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessInstanceDeclarationSlot(
	instanceDeclaration ast.InstanceDeclarationLike,
	slot_ uint,
) {
	switch slot_ {
	case 1, 2:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessInstanceSectionSlot(
	instanceSection ast.InstanceSectionLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessInterfaceDeclarations(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessInterfaceDeclarationsSlot(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessMethod(
	method ast.MethodLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PostprocessPackageDeclaration(
	packageDeclaration ast.PackageDeclarationLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessPackageDeclarationSlot(
	packageDeclaration ast.PackageDeclarationLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) ProcessPackageHeaderSlot(
	packageHeader ast.PackageHeaderLike,
	slot_ uint,
) {
	switch slot_ {
	case 2:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessPackageImports(
	packageImports ast.PackageImportsLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessPackageImportsSlot(
	packageImports ast.PackageImportsLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessParameter(
	parameter ast.ParameterLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessParameterSlot(
	parameter ast.ParameterLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessParameterList(
	parameterList ast.ParameterListLike,
	index_ uint,
	count_ uint,
) {
	v.depth_++
}

func (v *formatter_) PostprocessParameterList(
	parameterList ast.ParameterListLike,
	index_ uint,
	count_ uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PreprocessPrimitiveDeclarations(
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessPrimitiveDeclarationsSlot(
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessPrincipalSubsection(
	principalSubsection ast.PrincipalSubsectionLike,
	index_ uint,
	count_ uint,
) {
	v.depth_++
	v.appendNewline()
}

func (v *formatter_) PostprocessPrincipalSubsection(
	principalSubsection ast.PrincipalSubsectionLike,
	index_ uint,
	count_ uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PreprocessResult(
	result ast.ResultLike,
	index_ uint,
	count_ uint,
) {
	switch result.GetAny().(type) {
	case ast.NoneLike:
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessSetterMethod(
	setterMethod ast.SetterMethodLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessSetterMethodSlot(
	setterMethod ast.SetterMethodLike,
	slot_ uint,
) {
	switch slot_ {
	case 2:
		v.depth_++
	case 3:
		v.depth_--
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessTypeDeclaration(
	typeDeclaration ast.TypeDeclarationLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessTypeDeclarationSlot(
	typeDeclaration ast.TypeDeclarationLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendString(" ")
	case 2:
		v.appendNewline()
	}
}

func (v *formatter_) ProcessTypeSectionSlot(
	typeSection ast.TypeSectionLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) ProcessValueSlot(
	value ast.ValueLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendString(" ")
	}
}

const _indentation = "\t"

// PROTECTED INTERFACE

// Private Methods

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var level uint
	for ; level < v.depth_; level++ {
		newline += _indentation
	}
	v.appendString(newline)
}

func (v *formatter_) appendString(
	string_ string,
) {
	v.result_.WriteString(string_)
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}

// Instance Structure

type formatter_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike
	depth_   uint
	result_  sts.Builder

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type formatterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func formatterClass() *formatterClass_ {
	return formatterClassReference_
}

var formatterClassReference_ = &formatterClass_{
	// Initialize the class constants.
}
