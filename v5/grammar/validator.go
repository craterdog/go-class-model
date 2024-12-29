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
	fmt "fmt"
	ast "github.com/craterdog/go-class-model/v5/ast"
)

// CLASS INTERFACE

// Access Function

func ValidatorClass() ValidatorClassLike {
	return validatorClass()
}

// Constructor Methods

func (c *validatorClass_) Validator() ValidatorLike {
	var instance = &validator_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	instance.visitor_ = VisitorClass().Visitor(instance)
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *validator_) GetClass() ValidatorClassLike {
	return validatorClass()
}

func (v *validator_) ValidateModel(
	model ast.ModelLike,
) {
	v.visitor_.VisitModel(model)
}

// Methodical Methods

func (v *validator_) ProcessComment(
	comment string,
) {
	v.validateToken(comment, CommentToken)
}

func (v *validator_) ProcessName(
	name string,
) {
	v.validateToken(name, NameToken)
}

func (v *validator_) ProcessNewline(
	newline string,
) {
	v.validateToken(newline, NewlineToken)
}

func (v *validator_) ProcessPath(
	path string,
) {
	v.validateToken(path, PathToken)
}

func (v *validator_) ProcessPrefix(
	prefix string,
) {
	v.validateToken(prefix, PrefixToken)
}

func (v *validator_) ProcessSpace(
	space string,
) {
	v.validateToken(space, SpaceToken)
}

func (v *validator_) PreprocessAbstraction(
	abstraction ast.AbstractionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAbstractionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAbstraction(
	abstraction ast.AbstractionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAdditionalArgumentSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalConstraint(
	additionalConstraint ast.AdditionalConstraintLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAdditionalConstraintSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAdditionalConstraint(
	additionalConstraint ast.AdditionalConstraintLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAdditionalValueSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessArgument(
	argument ast.ArgumentLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessArgumentSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessArgument(
	argument ast.ArgumentLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessArguments(
	arguments ast.ArgumentsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessArgumentsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessArguments(
	arguments ast.ArgumentsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessArray(
	array ast.ArrayLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessArraySlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessArray(
	array ast.ArrayLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAspectDeclaration(
	aspectDeclaration ast.AspectDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAspectDeclarationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAspectDeclaration(
	aspectDeclaration ast.AspectDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAspectInterface(
	aspectInterface ast.AspectInterfaceLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAspectInterfaceSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAspectInterface(
	aspectInterface ast.AspectInterfaceLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAspectMethod(
	aspectMethod ast.AspectMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAspectMethodSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAspectMethod(
	aspectMethod ast.AspectMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAspectSection(
	aspectSection ast.AspectSectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAspectSectionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAspectSection(
	aspectSection ast.AspectSectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAspectSubsection(
	aspectSubsection ast.AspectSubsectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAspectSubsectionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAspectSubsection(
	aspectSubsection ast.AspectSubsectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAttributeMethod(
	attributeMethod ast.AttributeMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAttributeMethodSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAttributeMethod(
	attributeMethod ast.AttributeMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAttributeSubsection(
	attributeSubsection ast.AttributeSubsectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAttributeSubsectionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAttributeSubsection(
	attributeSubsection ast.AttributeSubsectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessChannel(
	channel ast.ChannelLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessChannelSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessChannel(
	channel ast.ChannelLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessClassDeclaration(
	classDeclaration ast.ClassDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessClassDeclarationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessClassDeclaration(
	classDeclaration ast.ClassDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessClassMethods(
	classMethods ast.ClassMethodsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessClassMethodsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessClassMethods(
	classMethods ast.ClassMethodsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessClassSection(
	classSection ast.ClassSectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessClassSectionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessClassSection(
	classSection ast.ClassSectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConstantMethod(
	constantMethod ast.ConstantMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessConstantMethodSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConstantMethod(
	constantMethod ast.ConstantMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConstantSubsection(
	constantSubsection ast.ConstantSubsectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessConstantSubsectionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConstantSubsection(
	constantSubsection ast.ConstantSubsectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConstraint(
	constraint ast.ConstraintLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessConstraintSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConstraint(
	constraint ast.ConstraintLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConstraints(
	constraints ast.ConstraintsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessConstraintsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConstraints(
	constraints ast.ConstraintsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConstructorMethod(
	constructorMethod ast.ConstructorMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessConstructorMethodSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConstructorMethod(
	constructorMethod ast.ConstructorMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConstructorSubsection(
	constructorSubsection ast.ConstructorSubsectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessConstructorSubsectionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConstructorSubsection(
	constructorSubsection ast.ConstructorSubsectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessDeclaration(
	declaration ast.DeclarationLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessDeclarationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessDeclaration(
	declaration ast.DeclarationLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessEnumeration(
	enumeration ast.EnumerationLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessEnumerationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessEnumeration(
	enumeration ast.EnumerationLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFunctionMethod(
	functionMethod ast.FunctionMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessFunctionMethodSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessFunctionMethod(
	functionMethod ast.FunctionMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFunctionSubsection(
	functionSubsection ast.FunctionSubsectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessFunctionSubsectionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessFunctionSubsection(
	functionSubsection ast.FunctionSubsectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFunctionalDeclaration(
	functionalDeclaration ast.FunctionalDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessFunctionalDeclarationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessFunctionalDeclaration(
	functionalDeclaration ast.FunctionalDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFunctionalSection(
	functionalSection ast.FunctionalSectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessFunctionalSectionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessFunctionalSection(
	functionalSection ast.FunctionalSectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessGetterMethod(
	getterMethod ast.GetterMethodLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessGetterMethodSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessGetterMethod(
	getterMethod ast.GetterMethodLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessImportedPackage(
	importedPackage ast.ImportedPackageLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessImportedPackageSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessImportedPackage(
	importedPackage ast.ImportedPackageLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInstanceDeclaration(
	instanceDeclaration ast.InstanceDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInstanceDeclarationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInstanceDeclaration(
	instanceDeclaration ast.InstanceDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInstanceMethods(
	instanceMethods ast.InstanceMethodsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInstanceMethodsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInstanceMethods(
	instanceMethods ast.InstanceMethodsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInstanceSection(
	instanceSection ast.InstanceSectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInstanceSectionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInstanceSection(
	instanceSection ast.InstanceSectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInterfaceDeclarations(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInterfaceDeclarationsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInterfaceDeclarations(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLegalNotice(
	legalNotice ast.LegalNoticeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessLegalNoticeSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLegalNotice(
	legalNotice ast.LegalNoticeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMap(
	map_ ast.MapLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMapSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMap(
	map_ ast.MapLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMethod(
	method ast.MethodLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMethodSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMethod(
	method ast.MethodLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessModel(
	model ast.ModelLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessModelSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessModel(
	model ast.ModelLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMultivalue(
	multivalue ast.MultivalueLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMultivalueSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMultivalue(
	multivalue ast.MultivalueLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNone(
	none ast.NoneLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessNoneSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessNone(
	none ast.NoneLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPackageDeclaration(
	packageDeclaration ast.PackageDeclarationLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessPackageDeclarationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPackageDeclaration(
	packageDeclaration ast.PackageDeclarationLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPackageHeader(
	packageHeader ast.PackageHeaderLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessPackageHeaderSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPackageHeader(
	packageHeader ast.PackageHeaderLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPackageImports(
	packageImports ast.PackageImportsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessPackageImportsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPackageImports(
	packageImports ast.PackageImportsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessParameter(
	parameter ast.ParameterLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessParameterSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessParameter(
	parameter ast.ParameterLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPrimitiveDeclarations(
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessPrimitiveDeclarationsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPrimitiveDeclarations(
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPrincipalMethod(
	principalMethod ast.PrincipalMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessPrincipalMethodSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPrincipalMethod(
	principalMethod ast.PrincipalMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPrincipalSubsection(
	principalSubsection ast.PrincipalSubsectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessPrincipalSubsectionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPrincipalSubsection(
	principalSubsection ast.PrincipalSubsectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessResult(
	result ast.ResultLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessResultSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessResult(
	result ast.ResultLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSetterMethod(
	setterMethod ast.SetterMethodLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessSetterMethodSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessSetterMethod(
	setterMethod ast.SetterMethodLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessTypeDeclaration(
	typeDeclaration ast.TypeDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessTypeDeclarationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessTypeDeclaration(
	typeDeclaration ast.TypeDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessTypeSection(
	typeSection ast.TypeSectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessTypeSectionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessTypeSection(
	typeSection ast.TypeSectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessValue(
	value ast.ValueLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessValueSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessValue(
	value ast.ValueLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessWrapper(
	wrapper ast.WrapperLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessWrapperSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessWrapper(
	wrapper ast.WrapperLike,
) {
	// TBD - Add any validation checks.
}

// PROTECTED INTERFACE

// Private Methods

func (v *validator_) validateToken(
	tokenValue string,
	tokenType TokenType,
) {
	var scannerClass = ScannerClass()
	if !scannerClass.MatchesType(tokenValue, tokenType) {
		var message = fmt.Sprintf(
			"The following token value is not of type %v: %v",
			scannerClass.FormatType(tokenType),
			tokenValue,
		)
		panic(message)
	}
}

// Instance Structure

type validator_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type validatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func validatorClass() *validatorClass_ {
	return validatorClassReference_
}

var validatorClassReference_ = &validatorClass_{
	// Initialize the class constants.
}
