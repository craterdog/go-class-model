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

package grammar

import (
	ast "github.com/craterdog/go-class-model/v7/ast"
)

// CLASS INTERFACE

// Access Function

func ProcessorClass() ProcessorClassLike {
	return processorClass()
}

// Constructor Methods

func (c *processorClass_) Processor() ProcessorLike {
	var instance = &processor_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *processor_) GetClass() ProcessorClassLike {
	return processorClass()
}

// Methodical Methods

func (v *processor_) ProcessComment(
	comment string,
) {
}

func (v *processor_) ProcessDelimiter(
	delimiter string,
) {
}

func (v *processor_) ProcessName(
	name string,
) {
}

func (v *processor_) ProcessNewline(
	newline string,
) {
}

func (v *processor_) ProcessPath(
	path string,
) {
}

func (v *processor_) ProcessPrefix(
	prefix string,
) {
}

func (v *processor_) ProcessSpace(
	space string,
) {
}

func (v *processor_) PreprocessAbstraction(
	abstraction ast.AbstractionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessAbstraction(
	abstraction ast.AbstractionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessAbstractionSlot(
	abstraction ast.AbstractionLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessAdditionalArgumentSlot(
	additionalArgument ast.AdditionalArgumentLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessAdditionalConstraint(
	additionalConstraint ast.AdditionalConstraintLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessAdditionalConstraint(
	additionalConstraint ast.AdditionalConstraintLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessAdditionalConstraintSlot(
	additionalConstraint ast.AdditionalConstraintLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessAdditionalValueSlot(
	additionalValue ast.AdditionalValueLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessArgument(
	argument ast.ArgumentLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessArgument(
	argument ast.ArgumentLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessArgumentSlot(
	argument ast.ArgumentLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessArguments(
	arguments ast.ArgumentsLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessArguments(
	arguments ast.ArgumentsLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessArgumentsSlot(
	arguments ast.ArgumentsLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessArray(
	array ast.ArrayLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessArray(
	array ast.ArrayLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessArraySlot(
	array ast.ArrayLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessAspectDeclaration(
	aspectDeclaration ast.AspectDeclarationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessAspectDeclaration(
	aspectDeclaration ast.AspectDeclarationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessAspectDeclarationSlot(
	aspectDeclaration ast.AspectDeclarationLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessAspectInterface(
	aspectInterface ast.AspectInterfaceLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessAspectInterface(
	aspectInterface ast.AspectInterfaceLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessAspectInterfaceSlot(
	aspectInterface ast.AspectInterfaceLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessAspectMethod(
	aspectMethod ast.AspectMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessAspectMethod(
	aspectMethod ast.AspectMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessAspectMethodSlot(
	aspectMethod ast.AspectMethodLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessAspectSection(
	aspectSection ast.AspectSectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessAspectSection(
	aspectSection ast.AspectSectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessAspectSectionSlot(
	aspectSection ast.AspectSectionLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessAspectSubsection(
	aspectSubsection ast.AspectSubsectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessAspectSubsection(
	aspectSubsection ast.AspectSubsectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessAspectSubsectionSlot(
	aspectSubsection ast.AspectSubsectionLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessAttributeMethod(
	attributeMethod ast.AttributeMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessAttributeMethod(
	attributeMethod ast.AttributeMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessAttributeMethodSlot(
	attributeMethod ast.AttributeMethodLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessAttributeSubsection(
	attributeSubsection ast.AttributeSubsectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessAttributeSubsection(
	attributeSubsection ast.AttributeSubsectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessAttributeSubsectionSlot(
	attributeSubsection ast.AttributeSubsectionLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessChannel(
	channel ast.ChannelLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessChannel(
	channel ast.ChannelLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessChannelSlot(
	channel ast.ChannelLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessClassDeclaration(
	classDeclaration ast.ClassDeclarationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessClassDeclaration(
	classDeclaration ast.ClassDeclarationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessClassDeclarationSlot(
	classDeclaration ast.ClassDeclarationLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessClassMethods(
	classMethods ast.ClassMethodsLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessClassMethods(
	classMethods ast.ClassMethodsLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessClassMethodsSlot(
	classMethods ast.ClassMethodsLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessClassSection(
	classSection ast.ClassSectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessClassSection(
	classSection ast.ClassSectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessClassSectionSlot(
	classSection ast.ClassSectionLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessConstantMethod(
	constantMethod ast.ConstantMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessConstantMethod(
	constantMethod ast.ConstantMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessConstantMethodSlot(
	constantMethod ast.ConstantMethodLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessConstantSubsection(
	constantSubsection ast.ConstantSubsectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessConstantSubsection(
	constantSubsection ast.ConstantSubsectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessConstantSubsectionSlot(
	constantSubsection ast.ConstantSubsectionLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessConstraint(
	constraint ast.ConstraintLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessConstraint(
	constraint ast.ConstraintLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessConstraintSlot(
	constraint ast.ConstraintLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessConstraints(
	constraints ast.ConstraintsLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessConstraints(
	constraints ast.ConstraintsLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessConstraintsSlot(
	constraints ast.ConstraintsLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessConstructorMethod(
	constructorMethod ast.ConstructorMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessConstructorMethod(
	constructorMethod ast.ConstructorMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessConstructorMethodSlot(
	constructorMethod ast.ConstructorMethodLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessConstructorSubsection(
	constructorSubsection ast.ConstructorSubsectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessConstructorSubsection(
	constructorSubsection ast.ConstructorSubsectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessConstructorSubsectionSlot(
	constructorSubsection ast.ConstructorSubsectionLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessDeclaration(
	declaration ast.DeclarationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessDeclaration(
	declaration ast.DeclarationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessDeclarationSlot(
	declaration ast.DeclarationLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessEnumeration(
	enumeration ast.EnumerationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessEnumeration(
	enumeration ast.EnumerationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessEnumerationSlot(
	enumeration ast.EnumerationLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessFunctionMethod(
	functionMethod ast.FunctionMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessFunctionMethod(
	functionMethod ast.FunctionMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessFunctionMethodSlot(
	functionMethod ast.FunctionMethodLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessFunctionSubsection(
	functionSubsection ast.FunctionSubsectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessFunctionSubsection(
	functionSubsection ast.FunctionSubsectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessFunctionSubsectionSlot(
	functionSubsection ast.FunctionSubsectionLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessFunctionalDeclaration(
	functionalDeclaration ast.FunctionalDeclarationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessFunctionalDeclaration(
	functionalDeclaration ast.FunctionalDeclarationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessFunctionalDeclarationSlot(
	functionalDeclaration ast.FunctionalDeclarationLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessFunctionalSection(
	functionalSection ast.FunctionalSectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessFunctionalSection(
	functionalSection ast.FunctionalSectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessFunctionalSectionSlot(
	functionalSection ast.FunctionalSectionLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessGetterMethod(
	getterMethod ast.GetterMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessGetterMethod(
	getterMethod ast.GetterMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessGetterMethodSlot(
	getterMethod ast.GetterMethodLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessImportList(
	importList ast.ImportListLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessImportList(
	importList ast.ImportListLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessImportListSlot(
	importList ast.ImportListLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessImportedPackage(
	importedPackage ast.ImportedPackageLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessImportedPackage(
	importedPackage ast.ImportedPackageLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessImportedPackageSlot(
	importedPackage ast.ImportedPackageLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessInstanceDeclaration(
	instanceDeclaration ast.InstanceDeclarationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessInstanceDeclaration(
	instanceDeclaration ast.InstanceDeclarationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessInstanceDeclarationSlot(
	instanceDeclaration ast.InstanceDeclarationLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessInstanceMethods(
	instanceMethods ast.InstanceMethodsLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessInstanceMethods(
	instanceMethods ast.InstanceMethodsLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessInstanceMethodsSlot(
	instanceMethods ast.InstanceMethodsLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessInstanceSection(
	instanceSection ast.InstanceSectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessInstanceSection(
	instanceSection ast.InstanceSectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessInstanceSectionSlot(
	instanceSection ast.InstanceSectionLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessInterfaceDeclarations(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessInterfaceDeclarations(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessInterfaceDeclarationsSlot(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessLegalNotice(
	legalNotice ast.LegalNoticeLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessLegalNotice(
	legalNotice ast.LegalNoticeLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessLegalNoticeSlot(
	legalNotice ast.LegalNoticeLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessMap(
	map_ ast.MapLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessMap(
	map_ ast.MapLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessMapSlot(
	map_ ast.MapLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessMethod(
	method ast.MethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessMethod(
	method ast.MethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessMethodSlot(
	method ast.MethodLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessModel(
	model ast.ModelLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessModel(
	model ast.ModelLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessModelSlot(
	model ast.ModelLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessMultivalue(
	multivalue ast.MultivalueLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessMultivalue(
	multivalue ast.MultivalueLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessMultivalueSlot(
	multivalue ast.MultivalueLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessNone(
	none ast.NoneLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessNone(
	none ast.NoneLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessNoneSlot(
	none ast.NoneLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessPackageDeclaration(
	packageDeclaration ast.PackageDeclarationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessPackageDeclaration(
	packageDeclaration ast.PackageDeclarationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessPackageDeclarationSlot(
	packageDeclaration ast.PackageDeclarationLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessPackageHeader(
	packageHeader ast.PackageHeaderLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessPackageHeader(
	packageHeader ast.PackageHeaderLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessPackageHeaderSlot(
	packageHeader ast.PackageHeaderLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessPackageImports(
	packageImports ast.PackageImportsLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessPackageImports(
	packageImports ast.PackageImportsLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessPackageImportsSlot(
	packageImports ast.PackageImportsLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessParameter(
	parameter ast.ParameterLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessParameter(
	parameter ast.ParameterLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessParameterSlot(
	parameter ast.ParameterLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessParameterList(
	parameterList ast.ParameterListLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessParameterList(
	parameterList ast.ParameterListLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessParameterListSlot(
	parameterList ast.ParameterListLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessPrimitiveDeclarations(
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessPrimitiveDeclarations(
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessPrimitiveDeclarationsSlot(
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessPrincipalMethod(
	principalMethod ast.PrincipalMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessPrincipalMethod(
	principalMethod ast.PrincipalMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessPrincipalMethodSlot(
	principalMethod ast.PrincipalMethodLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessPrincipalSubsection(
	principalSubsection ast.PrincipalSubsectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessPrincipalSubsection(
	principalSubsection ast.PrincipalSubsectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessPrincipalSubsectionSlot(
	principalSubsection ast.PrincipalSubsectionLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessResult(
	result ast.ResultLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessResult(
	result ast.ResultLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessResultSlot(
	result ast.ResultLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessSetterMethod(
	setterMethod ast.SetterMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessSetterMethod(
	setterMethod ast.SetterMethodLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessSetterMethodSlot(
	setterMethod ast.SetterMethodLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessStar(
	star ast.StarLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessStar(
	star ast.StarLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessStarSlot(
	star ast.StarLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessTypeDeclaration(
	typeDeclaration ast.TypeDeclarationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessTypeDeclaration(
	typeDeclaration ast.TypeDeclarationLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessTypeDeclarationSlot(
	typeDeclaration ast.TypeDeclarationLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessTypeSection(
	typeSection ast.TypeSectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessTypeSection(
	typeSection ast.TypeSectionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessTypeSectionSlot(
	typeSection ast.TypeSectionLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessValue(
	value ast.ValueLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessValue(
	value ast.ValueLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessValueSlot(
	value ast.ValueLike,
	slot_ uint,
) {
}

func (v *processor_) PreprocessWrapper(
	wrapper ast.WrapperLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) PostprocessWrapper(
	wrapper ast.WrapperLike,
	index_ uint,
	count_ uint,
) {
}

func (v *processor_) ProcessWrapperSlot(
	wrapper ast.WrapperLike,
	slot_ uint,
) {
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type processor_ struct {
	// Declare the instance attributes.
}

// Class Structure

type processorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func processorClass() *processorClass_ {
	return processorClassReference_
}

var processorClassReference_ = &processorClass_{
	// Initialize the class constants.
}
