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

func (v *formatter_) PreprocessAbstraction(
	abstraction ast.AbstractionLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PostprocessAbstraction(
	abstraction ast.AbstractionLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessAdditionalArgumentSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessAdditionalConstraintSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) PostprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessArgument(
	argument ast.ArgumentLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PostprocessArgument(
	argument ast.ArgumentLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessArguments(
	arguments ast.ArgumentsLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PostprocessArguments(
	arguments ast.ArgumentsLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessArray(
	array ast.ArrayLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PostprocessArray(
	array ast.ArrayLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessAspectDeclaration(
	aspectDeclaration ast.AspectDeclarationLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessAspectDeclarationSlot(
	slot uint,
) {
	switch slot {
	case 1, 2:
		v.appendString(" ")
	case 4:
		v.appendNewline()
	}
}

func (v *formatter_) PostprocessAspectDeclaration(
	aspectDeclaration ast.AspectDeclarationLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessAspectInterface(
	aspectInterface ast.AspectInterfaceLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessAspectMethod(
	aspectMethod ast.AspectMethodLike,
	index uint,
	count uint,
) {
	v.depth_++
}

func (v *formatter_) PostprocessAspectMethod(
	aspectMethod ast.AspectMethodLike,
	index uint,
	count uint,
) {
	v.depth_--
}

func (v *formatter_) ProcessAspectSectionSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessAspectSubsection(
	aspectSubsection ast.AspectSubsectionLike,
	index uint,
	count uint,
) {
	v.depth_++
	v.appendNewline()
}

func (v *formatter_) PostprocessAspectSubsection(
	aspectSubsection ast.AspectSubsectionLike,
	index uint,
	count uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PreprocessAttributeMethod(
	attributeMethod ast.AttributeMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PostprocessAttributeMethod(
	attributeMethod ast.AttributeMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessAttributeSubsection(
	attributeSubsection ast.AttributeSubsectionLike,
	index uint,
	count uint,
) {
	v.depth_++
	v.appendNewline()
}

func (v *formatter_) PostprocessAttributeSubsection(
	attributeSubsection ast.AttributeSubsectionLike,
	index uint,
	count uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PreprocessChannel(
	channel ast.ChannelLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PostprocessChannel(
	channel ast.ChannelLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessClassDeclaration(
	classDeclaration ast.ClassDeclarationLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessClassDeclarationSlot(
	slot uint,
) {
	switch slot {
	case 1, 2:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessClassDeclaration(
	classDeclaration ast.ClassDeclarationLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessClassSectionSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessConstantMethod(
	constantMethod ast.ConstantMethodLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessConstantMethodSlot(
	slot uint,
) {
	switch slot {
	case 3:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessConstantSubsection(
	constantSubsection ast.ConstantSubsectionLike,
	index uint,
	count uint,
) {
	v.depth_++
	v.appendNewline()
}

func (v *formatter_) PostprocessConstantSubsection(
	constantSubsection ast.ConstantSubsectionLike,
	index uint,
	count uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PreprocessConstraint(
	constraint ast.ConstraintLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PostprocessConstraint(
	constraint ast.ConstraintLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessConstraintSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessConstructorMethod(
	constructorMethod ast.ConstructorMethodLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessConstructorMethodSlot(
	slot uint,
) {
	switch slot {
	case 4:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessConstructorSubsection(
	constructorSubsection ast.ConstructorSubsectionLike,
	index uint,
	count uint,
) {
	v.depth_++
	v.appendNewline()
}

func (v *formatter_) PostprocessConstructorSubsection(
	constructorSubsection ast.ConstructorSubsectionLike,
	index uint,
	count uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) ProcessDeclarationSlot(
	slot uint,
) {
	switch slot {
	case 2:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessEnumeration(
	enumeration ast.EnumerationLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessEnumerationSlot(
	slot uint,
) {
	switch slot {
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

func (v *formatter_) PostprocessEnumeration(
	enumeration ast.EnumerationLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessFunctionMethod(
	functionMethod ast.FunctionMethodLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessFunctionSubsection(
	functionSubsection ast.FunctionSubsectionLike,
	index uint,
	count uint,
) {
	v.depth_++
	v.appendNewline()
}

func (v *formatter_) PostprocessFunctionSubsection(
	functionSubsection ast.FunctionSubsectionLike,
	index uint,
	count uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PreprocessFunctionalDeclaration(
	functionalDeclaration ast.FunctionalDeclarationLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessFunctionalDeclarationSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessFunctionalDeclaration(
	functionalDeclaration ast.FunctionalDeclarationLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessFunctionalSectionSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessGetterMethod(
	getterMethod ast.GetterMethodLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessGetterMethodSlot(
	slot uint,
) {
	switch slot {
	case 3:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessImportedPackage(
	importedPackage ast.ImportedPackageLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessImportedPackageSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessImportList(
	parameterList ast.ImportListLike,
	index uint,
	count uint,
) {
	v.depth_++
}

func (v *formatter_) PostprocessImportList(
	parameterList ast.ImportListLike,
	index uint,
	count uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PreprocessInstanceDeclaration(
	instanceDeclaration ast.InstanceDeclarationLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessInstanceDeclarationSlot(
	slot uint,
) {
	switch slot {
	case 1, 2:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessInstanceDeclaration(
	instanceDeclaration ast.InstanceDeclarationLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessInstanceSectionSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessInterfaceDeclarations(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessInterfaceDeclarationsSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) PostprocessInterfaceDeclarations(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
	index uint,
	count uint,
) {
	//v.appendNewline()
}

func (v *formatter_) PreprocessLegalNotice(
	legalNotice ast.LegalNoticeLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PostprocessLegalNotice(
	legalNotice ast.LegalNoticeLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessMap(
	map_ ast.MapLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PostprocessMap(
	map_ ast.MapLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessMethod(
	method ast.MethodLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessMultivalue(
	multivalue ast.MultivalueLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PostprocessMultivalue(
	multivalue ast.MultivalueLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessNone(
	none ast.NoneLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PostprocessNone(
	none ast.NoneLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessPackageDeclaration(
	packageDeclaration ast.PackageDeclarationLike,
	index uint,
	count uint,
) {
	//v.appendNewline()
}

func (v *formatter_) ProcessPackageDeclarationSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) PostprocessPackageDeclaration(
	packageDeclaration ast.PackageDeclarationLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessPackageHeader(
	packageHeader ast.PackageHeaderLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessPackageHeaderSlot(
	slot uint,
) {
	switch slot {
	case 2:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessPackageHeader(
	packageHeader ast.PackageHeaderLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessPackageImports(
	packageImports ast.PackageImportsLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessPackageImportsSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessParameter(
	parameter ast.ParameterLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessParameterSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessParameterList(
	parameterList ast.ParameterListLike,
	index uint,
	count uint,
) {
	v.depth_++
}

func (v *formatter_) PostprocessParameterList(
	parameterList ast.ParameterListLike,
	index uint,
	count uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PreprocessPrimitiveDeclarations(
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessPrimitiveDeclarationsSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) PostprocessPrimitiveDeclarations(
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
	index uint,
	count uint,
) {
	//v.appendNewline()
}

func (v *formatter_) PreprocessPrincipalSubsection(
	principalSubsection ast.PrincipalSubsectionLike,
	index uint,
	count uint,
) {
	v.depth_++
	v.appendNewline()
}

func (v *formatter_) PostprocessPrincipalSubsection(
	principalSubsection ast.PrincipalSubsectionLike,
	index uint,
	count uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PreprocessResult(
	result ast.ResultLike,
	index uint,
	count uint,
) {
	v.appendString(" ")
}

func (v *formatter_) PreprocessSetterMethod(
	setterMethod ast.SetterMethodLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessSetterMethodSlot(
	slot uint,
) {
	switch slot {
	case 2:
		v.depth_++
	case 3:
		v.depth_--
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessTypeDeclaration(
	typeDeclaration ast.TypeDeclarationLike,
	index uint,
	count uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessTypeDeclarationSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" ")
	case 2:
		v.appendNewline()
	}
}

func (v *formatter_) PostprocessTypeDeclaration(
	typeDeclaration ast.TypeDeclarationLike,
	index uint,
	count uint,
) {
	//v.appendNewline()
}

func (v *formatter_) ProcessTypeSectionSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessValue(
	value ast.ValueLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessValueSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessValue(
	value ast.ValueLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessWrapper(
	wrapper ast.WrapperLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PostprocessWrapper(
	wrapper ast.WrapperLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
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
