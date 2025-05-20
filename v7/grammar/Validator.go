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
	ast "github.com/craterdog/go-class-model/v7/ast"
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
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAbstraction(
	abstraction ast.AbstractionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalConstraint(
	additionalConstraint ast.AdditionalConstraintLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAdditionalConstraint(
	additionalConstraint ast.AdditionalConstraintLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessArgument(
	argument ast.ArgumentLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessArgument(
	argument ast.ArgumentLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessArguments(
	arguments ast.ArgumentsLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessArguments(
	arguments ast.ArgumentsLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessArray(
	array ast.ArrayLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessArray(
	array ast.ArrayLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAspectDeclaration(
	aspectDeclaration ast.AspectDeclarationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAspectDeclaration(
	aspectDeclaration ast.AspectDeclarationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAspectInterface(
	aspectInterface ast.AspectInterfaceLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAspectInterface(
	aspectInterface ast.AspectInterfaceLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAspectMethod(
	aspectMethod ast.AspectMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAspectMethod(
	aspectMethod ast.AspectMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAspectSection(
	aspectSection ast.AspectSectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAspectSection(
	aspectSection ast.AspectSectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAspectSubsection(
	aspectSubsection ast.AspectSubsectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAspectSubsection(
	aspectSubsection ast.AspectSubsectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAttributeMethod(
	attributeMethod ast.AttributeMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAttributeMethod(
	attributeMethod ast.AttributeMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAttributeSubsection(
	attributeSubsection ast.AttributeSubsectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAttributeSubsection(
	attributeSubsection ast.AttributeSubsectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessChannel(
	channel ast.ChannelLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessChannel(
	channel ast.ChannelLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessClassDeclaration(
	classDeclaration ast.ClassDeclarationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessClassDeclaration(
	classDeclaration ast.ClassDeclarationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessClassMethods(
	classMethods ast.ClassMethodsLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessClassMethods(
	classMethods ast.ClassMethodsLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessClassSection(
	classSection ast.ClassSectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessClassSection(
	classSection ast.ClassSectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConstantMethod(
	constantMethod ast.ConstantMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConstantMethod(
	constantMethod ast.ConstantMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConstantSubsection(
	constantSubsection ast.ConstantSubsectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConstantSubsection(
	constantSubsection ast.ConstantSubsectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConstraint(
	constraint ast.ConstraintLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConstraint(
	constraint ast.ConstraintLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConstraints(
	constraints ast.ConstraintsLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConstraints(
	constraints ast.ConstraintsLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConstructorMethod(
	constructorMethod ast.ConstructorMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConstructorMethod(
	constructorMethod ast.ConstructorMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConstructorSubsection(
	constructorSubsection ast.ConstructorSubsectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConstructorSubsection(
	constructorSubsection ast.ConstructorSubsectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessDeclaration(
	declaration ast.DeclarationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessDeclaration(
	declaration ast.DeclarationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessEnumeration(
	enumeration ast.EnumerationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessEnumeration(
	enumeration ast.EnumerationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFunctionMethod(
	functionMethod ast.FunctionMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessFunctionMethod(
	functionMethod ast.FunctionMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFunctionSubsection(
	functionSubsection ast.FunctionSubsectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessFunctionSubsection(
	functionSubsection ast.FunctionSubsectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFunctionalDeclaration(
	functionalDeclaration ast.FunctionalDeclarationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessFunctionalDeclaration(
	functionalDeclaration ast.FunctionalDeclarationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFunctionalSection(
	functionalSection ast.FunctionalSectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessFunctionalSection(
	functionalSection ast.FunctionalSectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessGetterMethod(
	getterMethod ast.GetterMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessGetterMethod(
	getterMethod ast.GetterMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessImportList(
	importList ast.ImportListLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessImportList(
	importList ast.ImportListLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessImportedPackage(
	importedPackage ast.ImportedPackageLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessImportedPackage(
	importedPackage ast.ImportedPackageLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInstanceDeclaration(
	instanceDeclaration ast.InstanceDeclarationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInstanceDeclaration(
	instanceDeclaration ast.InstanceDeclarationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInstanceMethods(
	instanceMethods ast.InstanceMethodsLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInstanceMethods(
	instanceMethods ast.InstanceMethodsLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInstanceSection(
	instanceSection ast.InstanceSectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInstanceSection(
	instanceSection ast.InstanceSectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInterfaceDeclarations(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInterfaceDeclarations(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLegalNotice(
	legalNotice ast.LegalNoticeLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLegalNotice(
	legalNotice ast.LegalNoticeLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMap(
	map_ ast.MapLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMap(
	map_ ast.MapLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMethod(
	method ast.MethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMethod(
	method ast.MethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessModel(
	model ast.ModelLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessModel(
	model ast.ModelLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMultivalue(
	multivalue ast.MultivalueLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMultivalue(
	multivalue ast.MultivalueLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNone(
	none ast.NoneLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessNone(
	none ast.NoneLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPackageDeclaration(
	packageDeclaration ast.PackageDeclarationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPackageDeclaration(
	packageDeclaration ast.PackageDeclarationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPackageHeader(
	packageHeader ast.PackageHeaderLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPackageHeader(
	packageHeader ast.PackageHeaderLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPackageImports(
	packageImports ast.PackageImportsLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPackageImports(
	packageImports ast.PackageImportsLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessParameter(
	parameter ast.ParameterLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessParameter(
	parameter ast.ParameterLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessParameterList(
	parameterList ast.ParameterListLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessParameterList(
	parameterList ast.ParameterListLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPrimitiveDeclarations(
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPrimitiveDeclarations(
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPrincipalMethod(
	principalMethod ast.PrincipalMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPrincipalMethod(
	principalMethod ast.PrincipalMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPrincipalSubsection(
	principalSubsection ast.PrincipalSubsectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPrincipalSubsection(
	principalSubsection ast.PrincipalSubsectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessResult(
	result ast.ResultLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessResult(
	result ast.ResultLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSetterMethod(
	setterMethod ast.SetterMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessSetterMethod(
	setterMethod ast.SetterMethodLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessTypeDeclaration(
	typeDeclaration ast.TypeDeclarationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessTypeDeclaration(
	typeDeclaration ast.TypeDeclarationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessTypeSection(
	typeSection ast.TypeSectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessTypeSection(
	typeSection ast.TypeSectionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessValue(
	value ast.ValueLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessValue(
	value ast.ValueLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessWrapper(
	wrapper ast.WrapperLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessWrapper(
	wrapper ast.WrapperLike,
	index uint,
	count uint,
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
