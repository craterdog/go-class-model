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
│         This "package_api.go" file was automatically generated using:        │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘

Package "ast" provides the abstract syntax tree (AST) classes for this module
based on the "syntax.cdsn" grammar for the module.  Each AST class manages the
attributes associated with its corresponding rule definition found in the
grammar.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-class-model/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package ast

import (
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
AbstractionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete abstraction-like class.
*/
type AbstractionClassLike interface {
	// Constructor Methods
	Abstraction(
		optionalWrapper WrapperLike,
		optionalPrefix string,
		name string,
		optionalArguments ArgumentsLike,
	) AbstractionLike
}

/*
AdditionalArgumentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete additional-argument-like class.
*/
type AdditionalArgumentClassLike interface {
	// Constructor Methods
	AdditionalArgument(
		delimiter string,
		argument ArgumentLike,
	) AdditionalArgumentLike
}

/*
AdditionalConstraintClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete additional-constraint-like class.
*/
type AdditionalConstraintClassLike interface {
	// Constructor Methods
	AdditionalConstraint(
		delimiter string,
		constraint ConstraintLike,
	) AdditionalConstraintLike
}

/*
AdditionalValueClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete additional-value-like class.
*/
type AdditionalValueClassLike interface {
	// Constructor Methods
	AdditionalValue(
		name string,
	) AdditionalValueLike
}

/*
ArgumentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete argument-like class.
*/
type ArgumentClassLike interface {
	// Constructor Methods
	Argument(
		abstraction AbstractionLike,
	) ArgumentLike
}

/*
ArgumentsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete arguments-like class.
*/
type ArgumentsClassLike interface {
	// Constructor Methods
	Arguments(
		delimiter1 string,
		argument ArgumentLike,
		additionalArguments fra.ListLike[AdditionalArgumentLike],
		delimiter2 string,
	) ArgumentsLike
}

/*
ArrayClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete array-like class.
*/
type ArrayClassLike interface {
	// Constructor Methods
	Array(
		delimiter1 string,
		delimiter2 string,
	) ArrayLike
}

/*
AspectDeclarationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete aspect-declaration-like class.
*/
type AspectDeclarationClassLike interface {
	// Constructor Methods
	AspectDeclaration(
		declaration DeclarationLike,
		delimiter1 string,
		delimiter2 string,
		aspectMethods fra.ListLike[AspectMethodLike],
		delimiter3 string,
	) AspectDeclarationLike
}

/*
AspectInterfaceClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete aspect-interface-like class.
*/
type AspectInterfaceClassLike interface {
	// Constructor Methods
	AspectInterface(
		abstraction AbstractionLike,
	) AspectInterfaceLike
}

/*
AspectMethodClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete aspect-method-like class.
*/
type AspectMethodClassLike interface {
	// Constructor Methods
	AspectMethod(
		method MethodLike,
	) AspectMethodLike
}

/*
AspectSectionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete aspect-section-like class.
*/
type AspectSectionClassLike interface {
	// Constructor Methods
	AspectSection(
		delimiter string,
		aspectDeclarations fra.ListLike[AspectDeclarationLike],
	) AspectSectionLike
}

/*
AspectSubsectionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete aspect-subsection-like class.
*/
type AspectSubsectionClassLike interface {
	// Constructor Methods
	AspectSubsection(
		delimiter string,
		aspectInterfaces fra.ListLike[AspectInterfaceLike],
	) AspectSubsectionLike
}

/*
AttributeMethodClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete attribute-method-like class.
*/
type AttributeMethodClassLike interface {
	// Constructor Methods
	AttributeMethod(
		any_ any,
	) AttributeMethodLike
}

/*
AttributeSubsectionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete attribute-subsection-like class.
*/
type AttributeSubsectionClassLike interface {
	// Constructor Methods
	AttributeSubsection(
		delimiter string,
		attributeMethods fra.ListLike[AttributeMethodLike],
	) AttributeSubsectionLike
}

/*
ChannelClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete channel-like class.
*/
type ChannelClassLike interface {
	// Constructor Methods
	Channel(
		delimiter string,
	) ChannelLike
}

/*
ClassDeclarationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete class-declaration-like class.
*/
type ClassDeclarationClassLike interface {
	// Constructor Methods
	ClassDeclaration(
		declaration DeclarationLike,
		delimiter1 string,
		delimiter2 string,
		classMethods ClassMethodsLike,
		delimiter3 string,
	) ClassDeclarationLike
}

/*
ClassMethodsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete class-methods-like class.
*/
type ClassMethodsClassLike interface {
	// Constructor Methods
	ClassMethods(
		constructorSubsection ConstructorSubsectionLike,
		optionalConstantSubsection ConstantSubsectionLike,
		optionalFunctionSubsection FunctionSubsectionLike,
	) ClassMethodsLike
}

/*
ClassSectionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete class-section-like class.
*/
type ClassSectionClassLike interface {
	// Constructor Methods
	ClassSection(
		delimiter string,
		classDeclarations fra.ListLike[ClassDeclarationLike],
	) ClassSectionLike
}

/*
ConstantMethodClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete constant-method-like class.
*/
type ConstantMethodClassLike interface {
	// Constructor Methods
	ConstantMethod(
		name string,
		delimiter1 string,
		delimiter2 string,
		abstraction AbstractionLike,
	) ConstantMethodLike
}

/*
ConstantSubsectionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete constant-subsection-like class.
*/
type ConstantSubsectionClassLike interface {
	// Constructor Methods
	ConstantSubsection(
		delimiter string,
		constantMethods fra.ListLike[ConstantMethodLike],
	) ConstantSubsectionLike
}

/*
ConstraintClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete constraint-like class.
*/
type ConstraintClassLike interface {
	// Constructor Methods
	Constraint(
		name string,
		abstraction AbstractionLike,
	) ConstraintLike
}

/*
ConstraintsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete constraints-like class.
*/
type ConstraintsClassLike interface {
	// Constructor Methods
	Constraints(
		delimiter1 string,
		constraint ConstraintLike,
		additionalConstraints fra.ListLike[AdditionalConstraintLike],
		delimiter2 string,
	) ConstraintsLike
}

/*
ConstructorMethodClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete constructor-method-like class.
*/
type ConstructorMethodClassLike interface {
	// Constructor Methods
	ConstructorMethod(
		name string,
		delimiter1 string,
		optionalParameterList ParameterListLike,
		delimiter2 string,
		abstraction AbstractionLike,
	) ConstructorMethodLike
}

/*
ConstructorSubsectionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete constructor-subsection-like class.
*/
type ConstructorSubsectionClassLike interface {
	// Constructor Methods
	ConstructorSubsection(
		delimiter string,
		constructorMethods fra.ListLike[ConstructorMethodLike],
	) ConstructorSubsectionLike
}

/*
DeclarationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete declaration-like class.
*/
type DeclarationClassLike interface {
	// Constructor Methods
	Declaration(
		comment string,
		delimiter string,
		name string,
		optionalConstraints ConstraintsLike,
	) DeclarationLike
}

/*
EnumerationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete enumeration-like class.
*/
type EnumerationClassLike interface {
	// Constructor Methods
	Enumeration(
		delimiter1 string,
		delimiter2 string,
		value ValueLike,
		additionalValues fra.ListLike[AdditionalValueLike],
		delimiter3 string,
	) EnumerationLike
}

/*
FunctionMethodClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete function-method-like class.
*/
type FunctionMethodClassLike interface {
	// Constructor Methods
	FunctionMethod(
		name string,
		delimiter1 string,
		optionalParameterList ParameterListLike,
		delimiter2 string,
		result ResultLike,
	) FunctionMethodLike
}

/*
FunctionSubsectionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete function-subsection-like class.
*/
type FunctionSubsectionClassLike interface {
	// Constructor Methods
	FunctionSubsection(
		delimiter string,
		functionMethods fra.ListLike[FunctionMethodLike],
	) FunctionSubsectionLike
}

/*
FunctionalDeclarationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete functional-declaration-like class.
*/
type FunctionalDeclarationClassLike interface {
	// Constructor Methods
	FunctionalDeclaration(
		declaration DeclarationLike,
		delimiter1 string,
		delimiter2 string,
		optionalParameterList ParameterListLike,
		delimiter3 string,
		result ResultLike,
	) FunctionalDeclarationLike
}

/*
FunctionalSectionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete functional-section-like class.
*/
type FunctionalSectionClassLike interface {
	// Constructor Methods
	FunctionalSection(
		delimiter string,
		functionalDeclarations fra.ListLike[FunctionalDeclarationLike],
	) FunctionalSectionLike
}

/*
GetterMethodClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete getter-method-like class.
*/
type GetterMethodClassLike interface {
	// Constructor Methods
	GetterMethod(
		name string,
		delimiter1 string,
		delimiter2 string,
		abstraction AbstractionLike,
	) GetterMethodLike
}

/*
ImportListClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete import-list-like class.
*/
type ImportListClassLike interface {
	// Constructor Methods
	ImportList(
		importedPackages fra.ListLike[ImportedPackageLike],
	) ImportListLike
}

/*
ImportedPackageClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete imported-package-like class.
*/
type ImportedPackageClassLike interface {
	// Constructor Methods
	ImportedPackage(
		name string,
		path string,
	) ImportedPackageLike
}

/*
InstanceDeclarationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete instance-declaration-like class.
*/
type InstanceDeclarationClassLike interface {
	// Constructor Methods
	InstanceDeclaration(
		declaration DeclarationLike,
		delimiter1 string,
		delimiter2 string,
		instanceMethods InstanceMethodsLike,
		delimiter3 string,
	) InstanceDeclarationLike
}

/*
InstanceMethodsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete instance-methods-like class.
*/
type InstanceMethodsClassLike interface {
	// Constructor Methods
	InstanceMethods(
		principalSubsection PrincipalSubsectionLike,
		optionalAttributeSubsection AttributeSubsectionLike,
		optionalAspectSubsection AspectSubsectionLike,
	) InstanceMethodsLike
}

/*
InstanceSectionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete instance-section-like class.
*/
type InstanceSectionClassLike interface {
	// Constructor Methods
	InstanceSection(
		delimiter string,
		instanceDeclarations fra.ListLike[InstanceDeclarationLike],
	) InstanceSectionLike
}

/*
InterfaceDeclarationsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete interface-declarations-like class.
*/
type InterfaceDeclarationsClassLike interface {
	// Constructor Methods
	InterfaceDeclarations(
		classSection ClassSectionLike,
		instanceSection InstanceSectionLike,
		aspectSection AspectSectionLike,
	) InterfaceDeclarationsLike
}

/*
LegalNoticeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete legal-notice-like class.
*/
type LegalNoticeClassLike interface {
	// Constructor Methods
	LegalNotice(
		comment string,
	) LegalNoticeLike
}

/*
MapClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete map-like class.
*/
type MapClassLike interface {
	// Constructor Methods
	Map(
		delimiter1 string,
		delimiter2 string,
		name string,
		delimiter3 string,
	) MapLike
}

/*
MethodClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete method-like class.
*/
type MethodClassLike interface {
	// Constructor Methods
	Method(
		name string,
		delimiter1 string,
		optionalParameterList ParameterListLike,
		delimiter2 string,
		result ResultLike,
	) MethodLike
}

/*
ModelClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete model-like class.
*/
type ModelClassLike interface {
	// Constructor Methods
	Model(
		packageDeclaration PackageDeclarationLike,
		primitiveDeclarations PrimitiveDeclarationsLike,
		interfaceDeclarations InterfaceDeclarationsLike,
	) ModelLike
}

/*
MultivalueClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete multivalue-like class.
*/
type MultivalueClassLike interface {
	// Constructor Methods
	Multivalue(
		delimiter1 string,
		parameterList ParameterListLike,
		delimiter2 string,
	) MultivalueLike
}

/*
NoneClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete none-like class.
*/
type NoneClassLike interface {
	// Constructor Methods
	None(
		newline string,
	) NoneLike
}

/*
PackageDeclarationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete package-declaration-like class.
*/
type PackageDeclarationClassLike interface {
	// Constructor Methods
	PackageDeclaration(
		legalNotice LegalNoticeLike,
		packageHeader PackageHeaderLike,
		packageImports PackageImportsLike,
	) PackageDeclarationLike
}

/*
PackageHeaderClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete package-header-like class.
*/
type PackageHeaderClassLike interface {
	// Constructor Methods
	PackageHeader(
		comment string,
		delimiter string,
		name string,
	) PackageHeaderLike
}

/*
PackageImportsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete package-imports-like class.
*/
type PackageImportsClassLike interface {
	// Constructor Methods
	PackageImports(
		delimiter1 string,
		delimiter2 string,
		optionalImportList ImportListLike,
		delimiter3 string,
	) PackageImportsLike
}

/*
ParameterClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete parameter-like class.
*/
type ParameterClassLike interface {
	// Constructor Methods
	Parameter(
		name string,
		abstraction AbstractionLike,
		delimiter string,
	) ParameterLike
}

/*
ParameterListClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete parameter-list-like class.
*/
type ParameterListClassLike interface {
	// Constructor Methods
	ParameterList(
		parameters fra.ListLike[ParameterLike],
	) ParameterListLike
}

/*
PrimitiveDeclarationsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete primitive-declarations-like class.
*/
type PrimitiveDeclarationsClassLike interface {
	// Constructor Methods
	PrimitiveDeclarations(
		typeSection TypeSectionLike,
		functionalSection FunctionalSectionLike,
	) PrimitiveDeclarationsLike
}

/*
PrincipalMethodClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete principal-method-like class.
*/
type PrincipalMethodClassLike interface {
	// Constructor Methods
	PrincipalMethod(
		method MethodLike,
	) PrincipalMethodLike
}

/*
PrincipalSubsectionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete principal-subsection-like class.
*/
type PrincipalSubsectionClassLike interface {
	// Constructor Methods
	PrincipalSubsection(
		delimiter string,
		principalMethods fra.ListLike[PrincipalMethodLike],
	) PrincipalSubsectionLike
}

/*
ResultClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete result-like class.
*/
type ResultClassLike interface {
	// Constructor Methods
	Result(
		any_ any,
	) ResultLike
}

/*
SetterMethodClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete setter-method-like class.
*/
type SetterMethodClassLike interface {
	// Constructor Methods
	SetterMethod(
		name string,
		delimiter1 string,
		parameter ParameterLike,
		delimiter2 string,
	) SetterMethodLike
}

/*
StarClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete star-like class.
*/
type StarClassLike interface {
	// Constructor Methods
	Star(
		delimiter string,
	) StarLike
}

/*
TypeDeclarationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete type-declaration-like class.
*/
type TypeDeclarationClassLike interface {
	// Constructor Methods
	TypeDeclaration(
		declaration DeclarationLike,
		abstraction AbstractionLike,
		optionalEnumeration EnumerationLike,
	) TypeDeclarationLike
}

/*
TypeSectionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete type-section-like class.
*/
type TypeSectionClassLike interface {
	// Constructor Methods
	TypeSection(
		delimiter string,
		typeDeclarations fra.ListLike[TypeDeclarationLike],
	) TypeSectionLike
}

/*
ValueClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete value-like class.
*/
type ValueClassLike interface {
	// Constructor Methods
	Value(
		name string,
		abstraction AbstractionLike,
		delimiter1 string,
		delimiter2 string,
	) ValueLike
}

/*
WrapperClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete wrapper-like class.
*/
type WrapperClassLike interface {
	// Constructor Methods
	Wrapper(
		any_ any,
	) WrapperLike
}

// INSTANCE DECLARATIONS

/*
AbstractionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete abstraction-like class.
*/
type AbstractionLike interface {
	// Principal Methods
	GetClass() AbstractionClassLike

	// Attribute Methods
	GetOptionalWrapper() WrapperLike
	GetOptionalPrefix() string
	GetName() string
	GetOptionalArguments() ArgumentsLike
}

/*
AdditionalArgumentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete additional-argument-like class.
*/
type AdditionalArgumentLike interface {
	// Principal Methods
	GetClass() AdditionalArgumentClassLike

	// Attribute Methods
	GetDelimiter() string
	GetArgument() ArgumentLike
}

/*
AdditionalConstraintLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete additional-constraint-like class.
*/
type AdditionalConstraintLike interface {
	// Principal Methods
	GetClass() AdditionalConstraintClassLike

	// Attribute Methods
	GetDelimiter() string
	GetConstraint() ConstraintLike
}

/*
AdditionalValueLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete additional-value-like class.
*/
type AdditionalValueLike interface {
	// Principal Methods
	GetClass() AdditionalValueClassLike

	// Attribute Methods
	GetName() string
}

/*
ArgumentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete argument-like class.
*/
type ArgumentLike interface {
	// Principal Methods
	GetClass() ArgumentClassLike

	// Attribute Methods
	GetAbstraction() AbstractionLike
}

/*
ArgumentsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete arguments-like class.
*/
type ArgumentsLike interface {
	// Principal Methods
	GetClass() ArgumentsClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetArgument() ArgumentLike
	GetAdditionalArguments() fra.ListLike[AdditionalArgumentLike]
	GetDelimiter2() string
}

/*
ArrayLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete array-like class.
*/
type ArrayLike interface {
	// Principal Methods
	GetClass() ArrayClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
}

/*
AspectDeclarationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete aspect-declaration-like class.
*/
type AspectDeclarationLike interface {
	// Principal Methods
	GetClass() AspectDeclarationClassLike

	// Attribute Methods
	GetDeclaration() DeclarationLike
	GetDelimiter1() string
	GetDelimiter2() string
	GetAspectMethods() fra.ListLike[AspectMethodLike]
	GetDelimiter3() string
}

/*
AspectInterfaceLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete aspect-interface-like class.
*/
type AspectInterfaceLike interface {
	// Principal Methods
	GetClass() AspectInterfaceClassLike

	// Attribute Methods
	GetAbstraction() AbstractionLike
}

/*
AspectMethodLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete aspect-method-like class.
*/
type AspectMethodLike interface {
	// Principal Methods
	GetClass() AspectMethodClassLike

	// Attribute Methods
	GetMethod() MethodLike
}

/*
AspectSectionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete aspect-section-like class.
*/
type AspectSectionLike interface {
	// Principal Methods
	GetClass() AspectSectionClassLike

	// Attribute Methods
	GetDelimiter() string
	GetAspectDeclarations() fra.ListLike[AspectDeclarationLike]
}

/*
AspectSubsectionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete aspect-subsection-like class.
*/
type AspectSubsectionLike interface {
	// Principal Methods
	GetClass() AspectSubsectionClassLike

	// Attribute Methods
	GetDelimiter() string
	GetAspectInterfaces() fra.ListLike[AspectInterfaceLike]
}

/*
AttributeMethodLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete attribute-method-like class.
*/
type AttributeMethodLike interface {
	// Principal Methods
	GetClass() AttributeMethodClassLike

	// Attribute Methods
	GetAny() any
}

/*
AttributeSubsectionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete attribute-subsection-like class.
*/
type AttributeSubsectionLike interface {
	// Principal Methods
	GetClass() AttributeSubsectionClassLike

	// Attribute Methods
	GetDelimiter() string
	GetAttributeMethods() fra.ListLike[AttributeMethodLike]
}

/*
ChannelLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete channel-like class.
*/
type ChannelLike interface {
	// Principal Methods
	GetClass() ChannelClassLike

	// Attribute Methods
	GetDelimiter() string
}

/*
ClassDeclarationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete class-declaration-like class.
*/
type ClassDeclarationLike interface {
	// Principal Methods
	GetClass() ClassDeclarationClassLike

	// Attribute Methods
	GetDeclaration() DeclarationLike
	GetDelimiter1() string
	GetDelimiter2() string
	GetClassMethods() ClassMethodsLike
	GetDelimiter3() string
}

/*
ClassMethodsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete class-methods-like class.
*/
type ClassMethodsLike interface {
	// Principal Methods
	GetClass() ClassMethodsClassLike

	// Attribute Methods
	GetConstructorSubsection() ConstructorSubsectionLike
	GetOptionalConstantSubsection() ConstantSubsectionLike
	GetOptionalFunctionSubsection() FunctionSubsectionLike
}

/*
ClassSectionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete class-section-like class.
*/
type ClassSectionLike interface {
	// Principal Methods
	GetClass() ClassSectionClassLike

	// Attribute Methods
	GetDelimiter() string
	GetClassDeclarations() fra.ListLike[ClassDeclarationLike]
}

/*
ConstantMethodLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete constant-method-like class.
*/
type ConstantMethodLike interface {
	// Principal Methods
	GetClass() ConstantMethodClassLike

	// Attribute Methods
	GetName() string
	GetDelimiter1() string
	GetDelimiter2() string
	GetAbstraction() AbstractionLike
}

/*
ConstantSubsectionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete constant-subsection-like class.
*/
type ConstantSubsectionLike interface {
	// Principal Methods
	GetClass() ConstantSubsectionClassLike

	// Attribute Methods
	GetDelimiter() string
	GetConstantMethods() fra.ListLike[ConstantMethodLike]
}

/*
ConstraintLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete constraint-like class.
*/
type ConstraintLike interface {
	// Principal Methods
	GetClass() ConstraintClassLike

	// Attribute Methods
	GetName() string
	GetAbstraction() AbstractionLike
}

/*
ConstraintsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete constraints-like class.
*/
type ConstraintsLike interface {
	// Principal Methods
	GetClass() ConstraintsClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetConstraint() ConstraintLike
	GetAdditionalConstraints() fra.ListLike[AdditionalConstraintLike]
	GetDelimiter2() string
}

/*
ConstructorMethodLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete constructor-method-like class.
*/
type ConstructorMethodLike interface {
	// Principal Methods
	GetClass() ConstructorMethodClassLike

	// Attribute Methods
	GetName() string
	GetDelimiter1() string
	GetOptionalParameterList() ParameterListLike
	GetDelimiter2() string
	GetAbstraction() AbstractionLike
}

/*
ConstructorSubsectionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete constructor-subsection-like class.
*/
type ConstructorSubsectionLike interface {
	// Principal Methods
	GetClass() ConstructorSubsectionClassLike

	// Attribute Methods
	GetDelimiter() string
	GetConstructorMethods() fra.ListLike[ConstructorMethodLike]
}

/*
DeclarationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete declaration-like class.
*/
type DeclarationLike interface {
	// Principal Methods
	GetClass() DeclarationClassLike

	// Attribute Methods
	GetComment() string
	GetDelimiter() string
	GetName() string
	GetOptionalConstraints() ConstraintsLike
}

/*
EnumerationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete enumeration-like class.
*/
type EnumerationLike interface {
	// Principal Methods
	GetClass() EnumerationClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
	GetValue() ValueLike
	GetAdditionalValues() fra.ListLike[AdditionalValueLike]
	GetDelimiter3() string
}

/*
FunctionMethodLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete function-method-like class.
*/
type FunctionMethodLike interface {
	// Principal Methods
	GetClass() FunctionMethodClassLike

	// Attribute Methods
	GetName() string
	GetDelimiter1() string
	GetOptionalParameterList() ParameterListLike
	GetDelimiter2() string
	GetResult() ResultLike
}

/*
FunctionSubsectionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete function-subsection-like class.
*/
type FunctionSubsectionLike interface {
	// Principal Methods
	GetClass() FunctionSubsectionClassLike

	// Attribute Methods
	GetDelimiter() string
	GetFunctionMethods() fra.ListLike[FunctionMethodLike]
}

/*
FunctionalDeclarationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete functional-declaration-like class.
*/
type FunctionalDeclarationLike interface {
	// Principal Methods
	GetClass() FunctionalDeclarationClassLike

	// Attribute Methods
	GetDeclaration() DeclarationLike
	GetDelimiter1() string
	GetDelimiter2() string
	GetOptionalParameterList() ParameterListLike
	GetDelimiter3() string
	GetResult() ResultLike
}

/*
FunctionalSectionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete functional-section-like class.
*/
type FunctionalSectionLike interface {
	// Principal Methods
	GetClass() FunctionalSectionClassLike

	// Attribute Methods
	GetDelimiter() string
	GetFunctionalDeclarations() fra.ListLike[FunctionalDeclarationLike]
}

/*
GetterMethodLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete getter-method-like class.
*/
type GetterMethodLike interface {
	// Principal Methods
	GetClass() GetterMethodClassLike

	// Attribute Methods
	GetName() string
	GetDelimiter1() string
	GetDelimiter2() string
	GetAbstraction() AbstractionLike
}

/*
ImportListLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete import-list-like class.
*/
type ImportListLike interface {
	// Principal Methods
	GetClass() ImportListClassLike

	// Attribute Methods
	GetImportedPackages() fra.ListLike[ImportedPackageLike]
}

/*
ImportedPackageLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete imported-package-like class.
*/
type ImportedPackageLike interface {
	// Principal Methods
	GetClass() ImportedPackageClassLike

	// Attribute Methods
	GetName() string
	GetPath() string
}

/*
InstanceDeclarationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete instance-declaration-like class.
*/
type InstanceDeclarationLike interface {
	// Principal Methods
	GetClass() InstanceDeclarationClassLike

	// Attribute Methods
	GetDeclaration() DeclarationLike
	GetDelimiter1() string
	GetDelimiter2() string
	GetInstanceMethods() InstanceMethodsLike
	GetDelimiter3() string
}

/*
InstanceMethodsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete instance-methods-like class.
*/
type InstanceMethodsLike interface {
	// Principal Methods
	GetClass() InstanceMethodsClassLike

	// Attribute Methods
	GetPrincipalSubsection() PrincipalSubsectionLike
	GetOptionalAttributeSubsection() AttributeSubsectionLike
	GetOptionalAspectSubsection() AspectSubsectionLike
}

/*
InstanceSectionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete instance-section-like class.
*/
type InstanceSectionLike interface {
	// Principal Methods
	GetClass() InstanceSectionClassLike

	// Attribute Methods
	GetDelimiter() string
	GetInstanceDeclarations() fra.ListLike[InstanceDeclarationLike]
}

/*
InterfaceDeclarationsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete interface-declarations-like class.
*/
type InterfaceDeclarationsLike interface {
	// Principal Methods
	GetClass() InterfaceDeclarationsClassLike

	// Attribute Methods
	GetClassSection() ClassSectionLike
	GetInstanceSection() InstanceSectionLike
	GetAspectSection() AspectSectionLike
}

/*
LegalNoticeLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete legal-notice-like class.
*/
type LegalNoticeLike interface {
	// Principal Methods
	GetClass() LegalNoticeClassLike

	// Attribute Methods
	GetComment() string
}

/*
MapLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete map-like class.
*/
type MapLike interface {
	// Principal Methods
	GetClass() MapClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
	GetName() string
	GetDelimiter3() string
}

/*
MethodLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete method-like class.
*/
type MethodLike interface {
	// Principal Methods
	GetClass() MethodClassLike

	// Attribute Methods
	GetName() string
	GetDelimiter1() string
	GetOptionalParameterList() ParameterListLike
	GetDelimiter2() string
	GetResult() ResultLike
}

/*
ModelLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete model-like class.
*/
type ModelLike interface {
	// Principal Methods
	GetClass() ModelClassLike

	// Attribute Methods
	GetPackageDeclaration() PackageDeclarationLike
	GetPrimitiveDeclarations() PrimitiveDeclarationsLike
	GetInterfaceDeclarations() InterfaceDeclarationsLike
}

/*
MultivalueLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete multivalue-like class.
*/
type MultivalueLike interface {
	// Principal Methods
	GetClass() MultivalueClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetParameterList() ParameterListLike
	GetDelimiter2() string
}

/*
NoneLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete none-like class.
*/
type NoneLike interface {
	// Principal Methods
	GetClass() NoneClassLike

	// Attribute Methods
	GetNewline() string
}

/*
PackageDeclarationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete package-declaration-like class.
*/
type PackageDeclarationLike interface {
	// Principal Methods
	GetClass() PackageDeclarationClassLike

	// Attribute Methods
	GetLegalNotice() LegalNoticeLike
	GetPackageHeader() PackageHeaderLike
	GetPackageImports() PackageImportsLike
}

/*
PackageHeaderLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete package-header-like class.
*/
type PackageHeaderLike interface {
	// Principal Methods
	GetClass() PackageHeaderClassLike

	// Attribute Methods
	GetComment() string
	GetDelimiter() string
	GetName() string
}

/*
PackageImportsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete package-imports-like class.
*/
type PackageImportsLike interface {
	// Principal Methods
	GetClass() PackageImportsClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
	GetOptionalImportList() ImportListLike
	GetDelimiter3() string
}

/*
ParameterLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete parameter-like class.
*/
type ParameterLike interface {
	// Principal Methods
	GetClass() ParameterClassLike

	// Attribute Methods
	GetName() string
	GetAbstraction() AbstractionLike
	GetDelimiter() string
}

/*
ParameterListLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete parameter-list-like class.
*/
type ParameterListLike interface {
	// Principal Methods
	GetClass() ParameterListClassLike

	// Attribute Methods
	GetParameters() fra.ListLike[ParameterLike]
}

/*
PrimitiveDeclarationsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete primitive-declarations-like class.
*/
type PrimitiveDeclarationsLike interface {
	// Principal Methods
	GetClass() PrimitiveDeclarationsClassLike

	// Attribute Methods
	GetTypeSection() TypeSectionLike
	GetFunctionalSection() FunctionalSectionLike
}

/*
PrincipalMethodLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete principal-method-like class.
*/
type PrincipalMethodLike interface {
	// Principal Methods
	GetClass() PrincipalMethodClassLike

	// Attribute Methods
	GetMethod() MethodLike
}

/*
PrincipalSubsectionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete principal-subsection-like class.
*/
type PrincipalSubsectionLike interface {
	// Principal Methods
	GetClass() PrincipalSubsectionClassLike

	// Attribute Methods
	GetDelimiter() string
	GetPrincipalMethods() fra.ListLike[PrincipalMethodLike]
}

/*
ResultLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete result-like class.
*/
type ResultLike interface {
	// Principal Methods
	GetClass() ResultClassLike

	// Attribute Methods
	GetAny() any
}

/*
SetterMethodLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete setter-method-like class.
*/
type SetterMethodLike interface {
	// Principal Methods
	GetClass() SetterMethodClassLike

	// Attribute Methods
	GetName() string
	GetDelimiter1() string
	GetParameter() ParameterLike
	GetDelimiter2() string
}

/*
StarLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete star-like class.
*/
type StarLike interface {
	// Principal Methods
	GetClass() StarClassLike

	// Attribute Methods
	GetDelimiter() string
}

/*
TypeDeclarationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete type-declaration-like class.
*/
type TypeDeclarationLike interface {
	// Principal Methods
	GetClass() TypeDeclarationClassLike

	// Attribute Methods
	GetDeclaration() DeclarationLike
	GetAbstraction() AbstractionLike
	GetOptionalEnumeration() EnumerationLike
}

/*
TypeSectionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete type-section-like class.
*/
type TypeSectionLike interface {
	// Principal Methods
	GetClass() TypeSectionClassLike

	// Attribute Methods
	GetDelimiter() string
	GetTypeDeclarations() fra.ListLike[TypeDeclarationLike]
}

/*
ValueLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete value-like class.
*/
type ValueLike interface {
	// Principal Methods
	GetClass() ValueClassLike

	// Attribute Methods
	GetName() string
	GetAbstraction() AbstractionLike
	GetDelimiter1() string
	GetDelimiter2() string
}

/*
WrapperLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete wrapper-like class.
*/
type WrapperLike interface {
	// Principal Methods
	GetClass() WrapperClassLike

	// Attribute Methods
	GetAny() any
}

// ASPECT DECLARATIONS
