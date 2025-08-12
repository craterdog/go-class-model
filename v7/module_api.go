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
│         This "module_api.go" file was automatically generated using:         │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-class-model/wiki
*/
package module

import (
	ast "github.com/craterdog/go-class-model/v7/ast"
	gra "github.com/craterdog/go-class-model/v7/grammar"
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE ALIASES

// Ast

type (
	AbstractionClassLike           = ast.AbstractionClassLike
	AdditionalArgumentClassLike    = ast.AdditionalArgumentClassLike
	AdditionalConstraintClassLike  = ast.AdditionalConstraintClassLike
	AdditionalValueClassLike       = ast.AdditionalValueClassLike
	ArgumentClassLike              = ast.ArgumentClassLike
	ArgumentsClassLike             = ast.ArgumentsClassLike
	ArrayClassLike                 = ast.ArrayClassLike
	AspectDeclarationClassLike     = ast.AspectDeclarationClassLike
	AspectInterfaceClassLike       = ast.AspectInterfaceClassLike
	AspectMethodClassLike          = ast.AspectMethodClassLike
	AspectSectionClassLike         = ast.AspectSectionClassLike
	AspectSubsectionClassLike      = ast.AspectSubsectionClassLike
	AttributeMethodClassLike       = ast.AttributeMethodClassLike
	AttributeSubsectionClassLike   = ast.AttributeSubsectionClassLike
	ChannelClassLike               = ast.ChannelClassLike
	ClassDeclarationClassLike      = ast.ClassDeclarationClassLike
	ClassMethodsClassLike          = ast.ClassMethodsClassLike
	ClassSectionClassLike          = ast.ClassSectionClassLike
	ConstantMethodClassLike        = ast.ConstantMethodClassLike
	ConstantSubsectionClassLike    = ast.ConstantSubsectionClassLike
	ConstraintClassLike            = ast.ConstraintClassLike
	ConstraintsClassLike           = ast.ConstraintsClassLike
	ConstructorMethodClassLike     = ast.ConstructorMethodClassLike
	ConstructorSubsectionClassLike = ast.ConstructorSubsectionClassLike
	DeclarationClassLike           = ast.DeclarationClassLike
	EnumerationClassLike           = ast.EnumerationClassLike
	FunctionMethodClassLike        = ast.FunctionMethodClassLike
	FunctionSubsectionClassLike    = ast.FunctionSubsectionClassLike
	FunctionalDeclarationClassLike = ast.FunctionalDeclarationClassLike
	FunctionalSectionClassLike     = ast.FunctionalSectionClassLike
	GetterMethodClassLike          = ast.GetterMethodClassLike
	ImportListClassLike            = ast.ImportListClassLike
	ImportedPackageClassLike       = ast.ImportedPackageClassLike
	InstanceDeclarationClassLike   = ast.InstanceDeclarationClassLike
	InstanceMethodsClassLike       = ast.InstanceMethodsClassLike
	InstanceSectionClassLike       = ast.InstanceSectionClassLike
	InterfaceDeclarationsClassLike = ast.InterfaceDeclarationsClassLike
	LegalNoticeClassLike           = ast.LegalNoticeClassLike
	MapClassLike                   = ast.MapClassLike
	MethodClassLike                = ast.MethodClassLike
	ModelClassLike                 = ast.ModelClassLike
	MultivalueClassLike            = ast.MultivalueClassLike
	NoneClassLike                  = ast.NoneClassLike
	PackageDeclarationClassLike    = ast.PackageDeclarationClassLike
	PackageHeaderClassLike         = ast.PackageHeaderClassLike
	PackageImportsClassLike        = ast.PackageImportsClassLike
	ParameterClassLike             = ast.ParameterClassLike
	ParameterListClassLike         = ast.ParameterListClassLike
	PrimitiveDeclarationsClassLike = ast.PrimitiveDeclarationsClassLike
	PrincipalMethodClassLike       = ast.PrincipalMethodClassLike
	PrincipalSubsectionClassLike   = ast.PrincipalSubsectionClassLike
	ResultClassLike                = ast.ResultClassLike
	SetterMethodClassLike          = ast.SetterMethodClassLike
	StarClassLike                  = ast.StarClassLike
	TypeDeclarationClassLike       = ast.TypeDeclarationClassLike
	TypeSectionClassLike           = ast.TypeSectionClassLike
	ValueClassLike                 = ast.ValueClassLike
	WrapperClassLike               = ast.WrapperClassLike
)

type (
	AbstractionLike           = ast.AbstractionLike
	AdditionalArgumentLike    = ast.AdditionalArgumentLike
	AdditionalConstraintLike  = ast.AdditionalConstraintLike
	AdditionalValueLike       = ast.AdditionalValueLike
	ArgumentLike              = ast.ArgumentLike
	ArgumentsLike             = ast.ArgumentsLike
	ArrayLike                 = ast.ArrayLike
	AspectDeclarationLike     = ast.AspectDeclarationLike
	AspectInterfaceLike       = ast.AspectInterfaceLike
	AspectMethodLike          = ast.AspectMethodLike
	AspectSectionLike         = ast.AspectSectionLike
	AspectSubsectionLike      = ast.AspectSubsectionLike
	AttributeMethodLike       = ast.AttributeMethodLike
	AttributeSubsectionLike   = ast.AttributeSubsectionLike
	ChannelLike               = ast.ChannelLike
	ClassDeclarationLike      = ast.ClassDeclarationLike
	ClassMethodsLike          = ast.ClassMethodsLike
	ClassSectionLike          = ast.ClassSectionLike
	ConstantMethodLike        = ast.ConstantMethodLike
	ConstantSubsectionLike    = ast.ConstantSubsectionLike
	ConstraintLike            = ast.ConstraintLike
	ConstraintsLike           = ast.ConstraintsLike
	ConstructorMethodLike     = ast.ConstructorMethodLike
	ConstructorSubsectionLike = ast.ConstructorSubsectionLike
	DeclarationLike           = ast.DeclarationLike
	EnumerationLike           = ast.EnumerationLike
	FunctionMethodLike        = ast.FunctionMethodLike
	FunctionSubsectionLike    = ast.FunctionSubsectionLike
	FunctionalDeclarationLike = ast.FunctionalDeclarationLike
	FunctionalSectionLike     = ast.FunctionalSectionLike
	GetterMethodLike          = ast.GetterMethodLike
	ImportListLike            = ast.ImportListLike
	ImportedPackageLike       = ast.ImportedPackageLike
	InstanceDeclarationLike   = ast.InstanceDeclarationLike
	InstanceMethodsLike       = ast.InstanceMethodsLike
	InstanceSectionLike       = ast.InstanceSectionLike
	InterfaceDeclarationsLike = ast.InterfaceDeclarationsLike
	LegalNoticeLike           = ast.LegalNoticeLike
	MapLike                   = ast.MapLike
	MethodLike                = ast.MethodLike
	ModelLike                 = ast.ModelLike
	MultivalueLike            = ast.MultivalueLike
	NoneLike                  = ast.NoneLike
	PackageDeclarationLike    = ast.PackageDeclarationLike
	PackageHeaderLike         = ast.PackageHeaderLike
	PackageImportsLike        = ast.PackageImportsLike
	ParameterLike             = ast.ParameterLike
	ParameterListLike         = ast.ParameterListLike
	PrimitiveDeclarationsLike = ast.PrimitiveDeclarationsLike
	PrincipalMethodLike       = ast.PrincipalMethodLike
	PrincipalSubsectionLike   = ast.PrincipalSubsectionLike
	ResultLike                = ast.ResultLike
	SetterMethodLike          = ast.SetterMethodLike
	StarLike                  = ast.StarLike
	TypeDeclarationLike       = ast.TypeDeclarationLike
	TypeSectionLike           = ast.TypeSectionLike
	ValueLike                 = ast.ValueLike
	WrapperLike               = ast.WrapperLike
)

// Grammar

type (
	TokenType = gra.TokenType
)

const (
	ErrorToken     = gra.ErrorToken
	CommentToken   = gra.CommentToken
	DelimiterToken = gra.DelimiterToken
	NameToken      = gra.NameToken
	NewlineToken   = gra.NewlineToken
	PathToken      = gra.PathToken
	PrefixToken    = gra.PrefixToken
	SpaceToken     = gra.SpaceToken
)

type (
	FormatterClassLike = gra.FormatterClassLike
	ParserClassLike    = gra.ParserClassLike
	ProcessorClassLike = gra.ProcessorClassLike
	ScannerClassLike   = gra.ScannerClassLike
	TokenClassLike     = gra.TokenClassLike
	ValidatorClassLike = gra.ValidatorClassLike
	VisitorClassLike   = gra.VisitorClassLike
)

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike   = gra.ScannerLike
	TokenLike     = gra.TokenLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
)

type (
	Methodical = gra.Methodical
)

// CLASS ACCESSORS

// Ast

func AbstractionClass() AbstractionClassLike {
	return ast.AbstractionClass()
}

func Abstraction(
	optionalWrapper ast.WrapperLike,
	optionalPrefix string,
	name string,
	optionalArguments ast.ArgumentsLike,
) AbstractionLike {
	return AbstractionClass().Abstraction(
		optionalWrapper,
		optionalPrefix,
		name,
		optionalArguments,
	)
}

func AdditionalArgumentClass() AdditionalArgumentClassLike {
	return ast.AdditionalArgumentClass()
}

func AdditionalArgument(
	delimiter string,
	argument ast.ArgumentLike,
) AdditionalArgumentLike {
	return AdditionalArgumentClass().AdditionalArgument(
		delimiter,
		argument,
	)
}

func AdditionalConstraintClass() AdditionalConstraintClassLike {
	return ast.AdditionalConstraintClass()
}

func AdditionalConstraint(
	delimiter string,
	constraint ast.ConstraintLike,
) AdditionalConstraintLike {
	return AdditionalConstraintClass().AdditionalConstraint(
		delimiter,
		constraint,
	)
}

func AdditionalValueClass() AdditionalValueClassLike {
	return ast.AdditionalValueClass()
}

func AdditionalValue(
	name string,
) AdditionalValueLike {
	return AdditionalValueClass().AdditionalValue(
		name,
	)
}

func ArgumentClass() ArgumentClassLike {
	return ast.ArgumentClass()
}

func Argument(
	abstraction ast.AbstractionLike,
) ArgumentLike {
	return ArgumentClass().Argument(
		abstraction,
	)
}

func ArgumentsClass() ArgumentsClassLike {
	return ast.ArgumentsClass()
}

func Arguments(
	delimiter1 string,
	argument ast.ArgumentLike,
	additionalArguments fra.Sequential[ast.AdditionalArgumentLike],
	delimiter2 string,
) ArgumentsLike {
	return ArgumentsClass().Arguments(
		delimiter1,
		argument,
		additionalArguments,
		delimiter2,
	)
}

func ArrayClass() ArrayClassLike {
	return ast.ArrayClass()
}

func Array(
	delimiter1 string,
	delimiter2 string,
) ArrayLike {
	return ArrayClass().Array(
		delimiter1,
		delimiter2,
	)
}

func AspectDeclarationClass() AspectDeclarationClassLike {
	return ast.AspectDeclarationClass()
}

func AspectDeclaration(
	declaration ast.DeclarationLike,
	delimiter1 string,
	delimiter2 string,
	aspectMethods fra.Sequential[ast.AspectMethodLike],
	delimiter3 string,
) AspectDeclarationLike {
	return AspectDeclarationClass().AspectDeclaration(
		declaration,
		delimiter1,
		delimiter2,
		aspectMethods,
		delimiter3,
	)
}

func AspectInterfaceClass() AspectInterfaceClassLike {
	return ast.AspectInterfaceClass()
}

func AspectInterface(
	abstraction ast.AbstractionLike,
) AspectInterfaceLike {
	return AspectInterfaceClass().AspectInterface(
		abstraction,
	)
}

func AspectMethodClass() AspectMethodClassLike {
	return ast.AspectMethodClass()
}

func AspectMethod(
	method ast.MethodLike,
) AspectMethodLike {
	return AspectMethodClass().AspectMethod(
		method,
	)
}

func AspectSectionClass() AspectSectionClassLike {
	return ast.AspectSectionClass()
}

func AspectSection(
	delimiter string,
	aspectDeclarations fra.Sequential[ast.AspectDeclarationLike],
) AspectSectionLike {
	return AspectSectionClass().AspectSection(
		delimiter,
		aspectDeclarations,
	)
}

func AspectSubsectionClass() AspectSubsectionClassLike {
	return ast.AspectSubsectionClass()
}

func AspectSubsection(
	delimiter string,
	aspectInterfaces fra.Sequential[ast.AspectInterfaceLike],
) AspectSubsectionLike {
	return AspectSubsectionClass().AspectSubsection(
		delimiter,
		aspectInterfaces,
	)
}

func AttributeMethodClass() AttributeMethodClassLike {
	return ast.AttributeMethodClass()
}

func AttributeMethod(
	any_ any,
) AttributeMethodLike {
	return AttributeMethodClass().AttributeMethod(
		any_,
	)
}

func AttributeSubsectionClass() AttributeSubsectionClassLike {
	return ast.AttributeSubsectionClass()
}

func AttributeSubsection(
	delimiter string,
	attributeMethods fra.Sequential[ast.AttributeMethodLike],
) AttributeSubsectionLike {
	return AttributeSubsectionClass().AttributeSubsection(
		delimiter,
		attributeMethods,
	)
}

func ChannelClass() ChannelClassLike {
	return ast.ChannelClass()
}

func Channel(
	delimiter string,
) ChannelLike {
	return ChannelClass().Channel(
		delimiter,
	)
}

func ClassDeclarationClass() ClassDeclarationClassLike {
	return ast.ClassDeclarationClass()
}

func ClassDeclaration(
	declaration ast.DeclarationLike,
	delimiter1 string,
	delimiter2 string,
	classMethods ast.ClassMethodsLike,
	delimiter3 string,
) ClassDeclarationLike {
	return ClassDeclarationClass().ClassDeclaration(
		declaration,
		delimiter1,
		delimiter2,
		classMethods,
		delimiter3,
	)
}

func ClassMethodsClass() ClassMethodsClassLike {
	return ast.ClassMethodsClass()
}

func ClassMethods(
	constructorSubsection ast.ConstructorSubsectionLike,
	optionalConstantSubsection ast.ConstantSubsectionLike,
	optionalFunctionSubsection ast.FunctionSubsectionLike,
) ClassMethodsLike {
	return ClassMethodsClass().ClassMethods(
		constructorSubsection,
		optionalConstantSubsection,
		optionalFunctionSubsection,
	)
}

func ClassSectionClass() ClassSectionClassLike {
	return ast.ClassSectionClass()
}

func ClassSection(
	delimiter string,
	classDeclarations fra.Sequential[ast.ClassDeclarationLike],
) ClassSectionLike {
	return ClassSectionClass().ClassSection(
		delimiter,
		classDeclarations,
	)
}

func ConstantMethodClass() ConstantMethodClassLike {
	return ast.ConstantMethodClass()
}

func ConstantMethod(
	name string,
	delimiter1 string,
	delimiter2 string,
	abstraction ast.AbstractionLike,
) ConstantMethodLike {
	return ConstantMethodClass().ConstantMethod(
		name,
		delimiter1,
		delimiter2,
		abstraction,
	)
}

func ConstantSubsectionClass() ConstantSubsectionClassLike {
	return ast.ConstantSubsectionClass()
}

func ConstantSubsection(
	delimiter string,
	constantMethods fra.Sequential[ast.ConstantMethodLike],
) ConstantSubsectionLike {
	return ConstantSubsectionClass().ConstantSubsection(
		delimiter,
		constantMethods,
	)
}

func ConstraintClass() ConstraintClassLike {
	return ast.ConstraintClass()
}

func Constraint(
	name string,
	abstraction ast.AbstractionLike,
) ConstraintLike {
	return ConstraintClass().Constraint(
		name,
		abstraction,
	)
}

func ConstraintsClass() ConstraintsClassLike {
	return ast.ConstraintsClass()
}

func Constraints(
	delimiter1 string,
	constraint ast.ConstraintLike,
	additionalConstraints fra.Sequential[ast.AdditionalConstraintLike],
	delimiter2 string,
) ConstraintsLike {
	return ConstraintsClass().Constraints(
		delimiter1,
		constraint,
		additionalConstraints,
		delimiter2,
	)
}

func ConstructorMethodClass() ConstructorMethodClassLike {
	return ast.ConstructorMethodClass()
}

func ConstructorMethod(
	name string,
	delimiter1 string,
	optionalParameterList ast.ParameterListLike,
	delimiter2 string,
	abstraction ast.AbstractionLike,
) ConstructorMethodLike {
	return ConstructorMethodClass().ConstructorMethod(
		name,
		delimiter1,
		optionalParameterList,
		delimiter2,
		abstraction,
	)
}

func ConstructorSubsectionClass() ConstructorSubsectionClassLike {
	return ast.ConstructorSubsectionClass()
}

func ConstructorSubsection(
	delimiter string,
	constructorMethods fra.Sequential[ast.ConstructorMethodLike],
) ConstructorSubsectionLike {
	return ConstructorSubsectionClass().ConstructorSubsection(
		delimiter,
		constructorMethods,
	)
}

func DeclarationClass() DeclarationClassLike {
	return ast.DeclarationClass()
}

func Declaration(
	comment string,
	delimiter string,
	name string,
	optionalConstraints ast.ConstraintsLike,
) DeclarationLike {
	return DeclarationClass().Declaration(
		comment,
		delimiter,
		name,
		optionalConstraints,
	)
}

func EnumerationClass() EnumerationClassLike {
	return ast.EnumerationClass()
}

func Enumeration(
	delimiter1 string,
	delimiter2 string,
	value ast.ValueLike,
	additionalValues fra.Sequential[ast.AdditionalValueLike],
	delimiter3 string,
) EnumerationLike {
	return EnumerationClass().Enumeration(
		delimiter1,
		delimiter2,
		value,
		additionalValues,
		delimiter3,
	)
}

func FunctionMethodClass() FunctionMethodClassLike {
	return ast.FunctionMethodClass()
}

func FunctionMethod(
	name string,
	delimiter1 string,
	optionalParameterList ast.ParameterListLike,
	delimiter2 string,
	result ast.ResultLike,
) FunctionMethodLike {
	return FunctionMethodClass().FunctionMethod(
		name,
		delimiter1,
		optionalParameterList,
		delimiter2,
		result,
	)
}

func FunctionSubsectionClass() FunctionSubsectionClassLike {
	return ast.FunctionSubsectionClass()
}

func FunctionSubsection(
	delimiter string,
	functionMethods fra.Sequential[ast.FunctionMethodLike],
) FunctionSubsectionLike {
	return FunctionSubsectionClass().FunctionSubsection(
		delimiter,
		functionMethods,
	)
}

func FunctionalDeclarationClass() FunctionalDeclarationClassLike {
	return ast.FunctionalDeclarationClass()
}

func FunctionalDeclaration(
	declaration ast.DeclarationLike,
	delimiter1 string,
	delimiter2 string,
	optionalParameterList ast.ParameterListLike,
	delimiter3 string,
	result ast.ResultLike,
) FunctionalDeclarationLike {
	return FunctionalDeclarationClass().FunctionalDeclaration(
		declaration,
		delimiter1,
		delimiter2,
		optionalParameterList,
		delimiter3,
		result,
	)
}

func FunctionalSectionClass() FunctionalSectionClassLike {
	return ast.FunctionalSectionClass()
}

func FunctionalSection(
	delimiter string,
	functionalDeclarations fra.Sequential[ast.FunctionalDeclarationLike],
) FunctionalSectionLike {
	return FunctionalSectionClass().FunctionalSection(
		delimiter,
		functionalDeclarations,
	)
}

func GetterMethodClass() GetterMethodClassLike {
	return ast.GetterMethodClass()
}

func GetterMethod(
	name string,
	delimiter1 string,
	delimiter2 string,
	abstraction ast.AbstractionLike,
) GetterMethodLike {
	return GetterMethodClass().GetterMethod(
		name,
		delimiter1,
		delimiter2,
		abstraction,
	)
}

func ImportListClass() ImportListClassLike {
	return ast.ImportListClass()
}

func ImportList(
	importedPackages fra.Sequential[ast.ImportedPackageLike],
) ImportListLike {
	return ImportListClass().ImportList(
		importedPackages,
	)
}

func ImportedPackageClass() ImportedPackageClassLike {
	return ast.ImportedPackageClass()
}

func ImportedPackage(
	name string,
	path string,
) ImportedPackageLike {
	return ImportedPackageClass().ImportedPackage(
		name,
		path,
	)
}

func InstanceDeclarationClass() InstanceDeclarationClassLike {
	return ast.InstanceDeclarationClass()
}

func InstanceDeclaration(
	declaration ast.DeclarationLike,
	delimiter1 string,
	delimiter2 string,
	instanceMethods ast.InstanceMethodsLike,
	delimiter3 string,
) InstanceDeclarationLike {
	return InstanceDeclarationClass().InstanceDeclaration(
		declaration,
		delimiter1,
		delimiter2,
		instanceMethods,
		delimiter3,
	)
}

func InstanceMethodsClass() InstanceMethodsClassLike {
	return ast.InstanceMethodsClass()
}

func InstanceMethods(
	principalSubsection ast.PrincipalSubsectionLike,
	optionalAttributeSubsection ast.AttributeSubsectionLike,
	optionalAspectSubsection ast.AspectSubsectionLike,
) InstanceMethodsLike {
	return InstanceMethodsClass().InstanceMethods(
		principalSubsection,
		optionalAttributeSubsection,
		optionalAspectSubsection,
	)
}

func InstanceSectionClass() InstanceSectionClassLike {
	return ast.InstanceSectionClass()
}

func InstanceSection(
	delimiter string,
	instanceDeclarations fra.Sequential[ast.InstanceDeclarationLike],
) InstanceSectionLike {
	return InstanceSectionClass().InstanceSection(
		delimiter,
		instanceDeclarations,
	)
}

func InterfaceDeclarationsClass() InterfaceDeclarationsClassLike {
	return ast.InterfaceDeclarationsClass()
}

func InterfaceDeclarations(
	classSection ast.ClassSectionLike,
	instanceSection ast.InstanceSectionLike,
	aspectSection ast.AspectSectionLike,
) InterfaceDeclarationsLike {
	return InterfaceDeclarationsClass().InterfaceDeclarations(
		classSection,
		instanceSection,
		aspectSection,
	)
}

func LegalNoticeClass() LegalNoticeClassLike {
	return ast.LegalNoticeClass()
}

func LegalNotice(
	comment string,
) LegalNoticeLike {
	return LegalNoticeClass().LegalNotice(
		comment,
	)
}

func MapClass() MapClassLike {
	return ast.MapClass()
}

func Map(
	delimiter1 string,
	delimiter2 string,
	name string,
	delimiter3 string,
) MapLike {
	return MapClass().Map(
		delimiter1,
		delimiter2,
		name,
		delimiter3,
	)
}

func MethodClass() MethodClassLike {
	return ast.MethodClass()
}

func Method(
	name string,
	delimiter1 string,
	optionalParameterList ast.ParameterListLike,
	delimiter2 string,
	result ast.ResultLike,
) MethodLike {
	return MethodClass().Method(
		name,
		delimiter1,
		optionalParameterList,
		delimiter2,
		result,
	)
}

func ModelClass() ModelClassLike {
	return ast.ModelClass()
}

func Model(
	packageDeclaration ast.PackageDeclarationLike,
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
	interfaceDeclarations ast.InterfaceDeclarationsLike,
) ModelLike {
	return ModelClass().Model(
		packageDeclaration,
		primitiveDeclarations,
		interfaceDeclarations,
	)
}

func MultivalueClass() MultivalueClassLike {
	return ast.MultivalueClass()
}

func Multivalue(
	delimiter1 string,
	parameterList ast.ParameterListLike,
	delimiter2 string,
) MultivalueLike {
	return MultivalueClass().Multivalue(
		delimiter1,
		parameterList,
		delimiter2,
	)
}

func NoneClass() NoneClassLike {
	return ast.NoneClass()
}

func None(
	newline string,
) NoneLike {
	return NoneClass().None(
		newline,
	)
}

func PackageDeclarationClass() PackageDeclarationClassLike {
	return ast.PackageDeclarationClass()
}

func PackageDeclaration(
	legalNotice ast.LegalNoticeLike,
	packageHeader ast.PackageHeaderLike,
	packageImports ast.PackageImportsLike,
) PackageDeclarationLike {
	return PackageDeclarationClass().PackageDeclaration(
		legalNotice,
		packageHeader,
		packageImports,
	)
}

func PackageHeaderClass() PackageHeaderClassLike {
	return ast.PackageHeaderClass()
}

func PackageHeader(
	comment string,
	delimiter string,
	name string,
) PackageHeaderLike {
	return PackageHeaderClass().PackageHeader(
		comment,
		delimiter,
		name,
	)
}

func PackageImportsClass() PackageImportsClassLike {
	return ast.PackageImportsClass()
}

func PackageImports(
	delimiter1 string,
	delimiter2 string,
	optionalImportList ast.ImportListLike,
	delimiter3 string,
) PackageImportsLike {
	return PackageImportsClass().PackageImports(
		delimiter1,
		delimiter2,
		optionalImportList,
		delimiter3,
	)
}

func ParameterClass() ParameterClassLike {
	return ast.ParameterClass()
}

func Parameter(
	name string,
	abstraction ast.AbstractionLike,
	delimiter string,
) ParameterLike {
	return ParameterClass().Parameter(
		name,
		abstraction,
		delimiter,
	)
}

func ParameterListClass() ParameterListClassLike {
	return ast.ParameterListClass()
}

func ParameterList(
	parameters fra.Sequential[ast.ParameterLike],
) ParameterListLike {
	return ParameterListClass().ParameterList(
		parameters,
	)
}

func PrimitiveDeclarationsClass() PrimitiveDeclarationsClassLike {
	return ast.PrimitiveDeclarationsClass()
}

func PrimitiveDeclarations(
	typeSection ast.TypeSectionLike,
	functionalSection ast.FunctionalSectionLike,
) PrimitiveDeclarationsLike {
	return PrimitiveDeclarationsClass().PrimitiveDeclarations(
		typeSection,
		functionalSection,
	)
}

func PrincipalMethodClass() PrincipalMethodClassLike {
	return ast.PrincipalMethodClass()
}

func PrincipalMethod(
	method ast.MethodLike,
) PrincipalMethodLike {
	return PrincipalMethodClass().PrincipalMethod(
		method,
	)
}

func PrincipalSubsectionClass() PrincipalSubsectionClassLike {
	return ast.PrincipalSubsectionClass()
}

func PrincipalSubsection(
	delimiter string,
	principalMethods fra.Sequential[ast.PrincipalMethodLike],
) PrincipalSubsectionLike {
	return PrincipalSubsectionClass().PrincipalSubsection(
		delimiter,
		principalMethods,
	)
}

func ResultClass() ResultClassLike {
	return ast.ResultClass()
}

func Result(
	any_ any,
) ResultLike {
	return ResultClass().Result(
		any_,
	)
}

func SetterMethodClass() SetterMethodClassLike {
	return ast.SetterMethodClass()
}

func SetterMethod(
	name string,
	delimiter1 string,
	parameter ast.ParameterLike,
	delimiter2 string,
) SetterMethodLike {
	return SetterMethodClass().SetterMethod(
		name,
		delimiter1,
		parameter,
		delimiter2,
	)
}

func StarClass() StarClassLike {
	return ast.StarClass()
}

func Star(
	delimiter string,
) StarLike {
	return StarClass().Star(
		delimiter,
	)
}

func TypeDeclarationClass() TypeDeclarationClassLike {
	return ast.TypeDeclarationClass()
}

func TypeDeclaration(
	declaration ast.DeclarationLike,
	abstraction ast.AbstractionLike,
	optionalEnumeration ast.EnumerationLike,
) TypeDeclarationLike {
	return TypeDeclarationClass().TypeDeclaration(
		declaration,
		abstraction,
		optionalEnumeration,
	)
}

func TypeSectionClass() TypeSectionClassLike {
	return ast.TypeSectionClass()
}

func TypeSection(
	delimiter string,
	typeDeclarations fra.Sequential[ast.TypeDeclarationLike],
) TypeSectionLike {
	return TypeSectionClass().TypeSection(
		delimiter,
		typeDeclarations,
	)
}

func ValueClass() ValueClassLike {
	return ast.ValueClass()
}

func Value(
	name string,
	abstraction ast.AbstractionLike,
	delimiter1 string,
	delimiter2 string,
) ValueLike {
	return ValueClass().Value(
		name,
		abstraction,
		delimiter1,
		delimiter2,
	)
}

func WrapperClass() WrapperClassLike {
	return ast.WrapperClass()
}

func Wrapper(
	any_ any,
) WrapperLike {
	return WrapperClass().Wrapper(
		any_,
	)
}

// Grammar

func FormatterClass() FormatterClassLike {
	return gra.FormatterClass()
}

func Formatter() FormatterLike {
	return FormatterClass().Formatter()
}

func ParserClass() ParserClassLike {
	return gra.ParserClass()
}

func Parser() ParserLike {
	return ParserClass().Parser()
}

func ProcessorClass() ProcessorClassLike {
	return gra.ProcessorClass()
}

func Processor() ProcessorLike {
	return ProcessorClass().Processor()
}

func ScannerClass() ScannerClassLike {
	return gra.ScannerClass()
}

func Scanner(
	source string,
	tokens fra.QueueLike[gra.TokenLike],
) ScannerLike {
	return ScannerClass().Scanner(
		source,
		tokens,
	)
}

func TokenClass() TokenClassLike {
	return gra.TokenClass()
}

func Token(
	line uint,
	position uint,
	type_ gra.TokenType,
	value string,
) TokenLike {
	return TokenClass().Token(
		line,
		position,
		type_,
		value,
	)
}

func ValidatorClass() ValidatorClassLike {
	return gra.ValidatorClass()
}

func Validator() ValidatorLike {
	return ValidatorClass().Validator()
}

func VisitorClass() VisitorClassLike {
	return gra.VisitorClass()
}

func Visitor(
	processor gra.Methodical,
) VisitorLike {
	return VisitorClass().Visitor(
		processor,
	)
}

// GLOBAL FUNCTIONS

func FormatModel(
	model ModelLike,
) string {
	var formatter = Formatter()
	return formatter.FormatModel(model)
}

func MatchesType(
	tokenValue string,
	tokenType TokenType,
) bool {
	var scannerClass = gra.ScannerClass()
	return scannerClass.MatchesType(tokenValue, tokenType)
}

func ParseSource(
	source string,
) ModelLike {
	var parser = Parser()
	return parser.ParseSource(source)
}

func ValidateModel(
	model ModelLike,
) {
	var validator = Validator()
	validator.ValidateModel(model)
}
