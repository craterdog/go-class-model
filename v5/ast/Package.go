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
Package "ast" provides the abstract syntax tree (AST) classes for this module
based on the "Syntax.cdsn" grammar for the module.  Each AST class manages the
attributes associated with its corresponding rule definition found in the
grammar.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-class-model/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-class-model/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package ast

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// Type Declarations

// Functional Declarations

// Class Declarations

/*
AbstractionClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete abstraction-like class.
*/
type AbstractionClassLike interface {
	// Constructor Methods
	Make(
		optionalPrefix PrefixLike,
		name string,
		optionalSuffix SuffixLike,
		optionalArguments ArgumentsLike,
	) AbstractionLike
}

/*
AdditionalArgumentClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete additional-argument-like class.
*/
type AdditionalArgumentClassLike interface {
	// Constructor Methods
	Make(
		argument ArgumentLike,
	) AdditionalArgumentLike
}

/*
AdditionalConstraintClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete additional-constraint-like class.
*/
type AdditionalConstraintClassLike interface {
	// Constructor Methods
	Make(
		constraint ConstraintLike,
	) AdditionalConstraintLike
}

/*
AdditionalValueClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete additional-value-like class.
*/
type AdditionalValueClassLike interface {
	// Constructor Methods
	Make(
		name string,
	) AdditionalValueLike
}

/*
ArgumentClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete argument-like class.
*/
type ArgumentClassLike interface {
	// Constructor Methods
	Make(
		abstraction AbstractionLike,
	) ArgumentLike
}

/*
ArgumentsClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete arguments-like class.
*/
type ArgumentsClassLike interface {
	// Constructor Methods
	Make(
		argument ArgumentLike,
		additionalArguments abs.Sequential[AdditionalArgumentLike],
	) ArgumentsLike
}

/*
ArrayClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete array-like class.
*/
type ArrayClassLike interface {
	// Constructor Methods
	Make() ArrayLike
}

/*
AspectDeclarationClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete aspect-declaration-like class.
*/
type AspectDeclarationClassLike interface {
	// Constructor Methods
	Make(
		declaration DeclarationLike,
		aspectMethods abs.Sequential[AspectMethodLike],
	) AspectDeclarationLike
}

/*
AspectInterfaceClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete aspect-interface-like class.
*/
type AspectInterfaceClassLike interface {
	// Constructor Methods
	Make(
		abstraction AbstractionLike,
	) AspectInterfaceLike
}

/*
AspectMethodClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete aspect-method-like class.
*/
type AspectMethodClassLike interface {
	// Constructor Methods
	Make(
		method MethodLike,
	) AspectMethodLike
}

/*
AspectSectionClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete aspect-section-like class.
*/
type AspectSectionClassLike interface {
	// Constructor Methods
	Make(
		aspectDeclarations abs.Sequential[AspectDeclarationLike],
	) AspectSectionLike
}

/*
AspectSubsectionClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete aspect-subsection-like class.
*/
type AspectSubsectionClassLike interface {
	// Constructor Methods
	Make(
		aspectInterfaces abs.Sequential[AspectInterfaceLike],
	) AspectSubsectionLike
}

/*
AttributeMethodClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete attribute-method-like class.
*/
type AttributeMethodClassLike interface {
	// Constructor Methods
	Make(
		any_ any,
	) AttributeMethodLike
}

/*
AttributeSubsectionClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete attribute-subsection-like class.
*/
type AttributeSubsectionClassLike interface {
	// Constructor Methods
	Make(
		attributeMethods abs.Sequential[AttributeMethodLike],
	) AttributeSubsectionLike
}

/*
ChannelClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete channel-like class.
*/
type ChannelClassLike interface {
	// Constructor Methods
	Make() ChannelLike
}

/*
ClassDeclarationClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete class-declaration-like class.
*/
type ClassDeclarationClassLike interface {
	// Constructor Methods
	Make(
		declaration DeclarationLike,
		classMethods ClassMethodsLike,
	) ClassDeclarationLike
}

/*
ClassMethodsClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete class-methods-like class.
*/
type ClassMethodsClassLike interface {
	// Constructor Methods
	Make(
		constructorSubsection ConstructorSubsectionLike,
		optionalConstantSubsection ConstantSubsectionLike,
		optionalFunctionSubsection FunctionSubsectionLike,
	) ClassMethodsLike
}

/*
ClassSectionClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete class-section-like class.
*/
type ClassSectionClassLike interface {
	// Constructor Methods
	Make(
		classDeclarations abs.Sequential[ClassDeclarationLike],
	) ClassSectionLike
}

/*
ConstantMethodClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete constant-method-like class.
*/
type ConstantMethodClassLike interface {
	// Constructor Methods
	Make(
		name string,
		abstraction AbstractionLike,
	) ConstantMethodLike
}

/*
ConstantSubsectionClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete constant-subsection-like class.
*/
type ConstantSubsectionClassLike interface {
	// Constructor Methods
	Make(
		constantMethods abs.Sequential[ConstantMethodLike],
	) ConstantSubsectionLike
}

/*
ConstraintClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete constraint-like class.
*/
type ConstraintClassLike interface {
	// Constructor Methods
	Make(
		name string,
		abstraction AbstractionLike,
	) ConstraintLike
}

/*
ConstraintsClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete constraints-like class.
*/
type ConstraintsClassLike interface {
	// Constructor Methods
	Make(
		constraint ConstraintLike,
		additionalConstraints abs.Sequential[AdditionalConstraintLike],
	) ConstraintsLike
}

/*
ConstructorMethodClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete constructor-method-like class.
*/
type ConstructorMethodClassLike interface {
	// Constructor Methods
	Make(
		name string,
		parameters abs.Sequential[ParameterLike],
		abstraction AbstractionLike,
	) ConstructorMethodLike
}

/*
ConstructorSubsectionClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete constructor-subsection-like class.
*/
type ConstructorSubsectionClassLike interface {
	// Constructor Methods
	Make(
		constructorMethods abs.Sequential[ConstructorMethodLike],
	) ConstructorSubsectionLike
}

/*
DeclarationClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete declaration-like class.
*/
type DeclarationClassLike interface {
	// Constructor Methods
	Make(
		comment string,
		name string,
		optionalConstraints ConstraintsLike,
	) DeclarationLike
}

/*
EnumerationClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete enumeration-like class.
*/
type EnumerationClassLike interface {
	// Constructor Methods
	Make(
		value ValueLike,
		additionalValues abs.Sequential[AdditionalValueLike],
	) EnumerationLike
}

/*
FunctionMethodClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete function-method-like class.
*/
type FunctionMethodClassLike interface {
	// Constructor Methods
	Make(
		name string,
		parameters abs.Sequential[ParameterLike],
		result ResultLike,
	) FunctionMethodLike
}

/*
FunctionSubsectionClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete function-subsection-like class.
*/
type FunctionSubsectionClassLike interface {
	// Constructor Methods
	Make(
		functionMethods abs.Sequential[FunctionMethodLike],
	) FunctionSubsectionLike
}

/*
FunctionalDeclarationClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete functional-declaration-like class.
*/
type FunctionalDeclarationClassLike interface {
	// Constructor Methods
	Make(
		declaration DeclarationLike,
		parameters abs.Sequential[ParameterLike],
		result ResultLike,
	) FunctionalDeclarationLike
}

/*
FunctionalSectionClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete functional-section-like class.
*/
type FunctionalSectionClassLike interface {
	// Constructor Methods
	Make(
		functionalDeclarations abs.Sequential[FunctionalDeclarationLike],
	) FunctionalSectionLike
}

/*
GetterMethodClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete getter-method-like class.
*/
type GetterMethodClassLike interface {
	// Constructor Methods
	Make(
		name string,
		abstraction AbstractionLike,
	) GetterMethodLike
}

/*
ImportedPackageClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete imported-package-like class.
*/
type ImportedPackageClassLike interface {
	// Constructor Methods
	Make(
		name string,
		path string,
	) ImportedPackageLike
}

/*
InstanceDeclarationClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete instance-declaration-like class.
*/
type InstanceDeclarationClassLike interface {
	// Constructor Methods
	Make(
		declaration DeclarationLike,
		instanceMethods InstanceMethodsLike,
	) InstanceDeclarationLike
}

/*
InstanceMethodsClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete instance-methods-like class.
*/
type InstanceMethodsClassLike interface {
	// Constructor Methods
	Make(
		primarySubsection PrimarySubsectionLike,
		optionalAttributeSubsection AttributeSubsectionLike,
		optionalAspectSubsection AspectSubsectionLike,
	) InstanceMethodsLike
}

/*
InstanceSectionClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete instance-section-like class.
*/
type InstanceSectionClassLike interface {
	// Constructor Methods
	Make(
		instanceDeclarations abs.Sequential[InstanceDeclarationLike],
	) InstanceSectionLike
}

/*
InterfaceDeclarationsClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete interface-declarations-like class.
*/
type InterfaceDeclarationsClassLike interface {
	// Constructor Methods
	Make(
		classSection ClassSectionLike,
		instanceSection InstanceSectionLike,
		aspectSection AspectSectionLike,
	) InterfaceDeclarationsLike
}

/*
LegalNoticeClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete legal-notice-like class.
*/
type LegalNoticeClassLike interface {
	// Constructor Methods
	Make(
		comment string,
	) LegalNoticeLike
}

/*
MapClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete map-like class.
*/
type MapClassLike interface {
	// Constructor Methods
	Make(
		name string,
	) MapLike
}

/*
MethodClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete method-like class.
*/
type MethodClassLike interface {
	// Constructor Methods
	Make(
		name string,
		parameters abs.Sequential[ParameterLike],
		optionalResult ResultLike,
	) MethodLike
}

/*
ModelClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete model-like class.
*/
type ModelClassLike interface {
	// Constructor Methods
	Make(
		moduleDeclaration ModuleDeclarationLike,
		primitiveDeclarations PrimitiveDeclarationsLike,
		interfaceDeclarations InterfaceDeclarationsLike,
	) ModelLike
}

/*
ModuleDeclarationClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete module-declaration-like class.
*/
type ModuleDeclarationClassLike interface {
	// Constructor Methods
	Make(
		legalNotice LegalNoticeLike,
		moduleHeader ModuleHeaderLike,
		moduleImports ModuleImportsLike,
	) ModuleDeclarationLike
}

/*
ModuleHeaderClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete module-header-like class.
*/
type ModuleHeaderClassLike interface {
	// Constructor Methods
	Make(
		comment string,
		name string,
	) ModuleHeaderLike
}

/*
ModuleImportsClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete module-imports-like class.
*/
type ModuleImportsClassLike interface {
	// Constructor Methods
	Make(
		importedPackages abs.Sequential[ImportedPackageLike],
	) ModuleImportsLike
}

/*
MultivalueClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete multivalue-like class.
*/
type MultivalueClassLike interface {
	// Constructor Methods
	Make(
		parameters abs.Sequential[ParameterLike],
	) MultivalueLike
}

/*
NoneClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete none-like class.
*/
type NoneClassLike interface {
	// Constructor Methods
	Make(
		newline string,
	) NoneLike
}

/*
ParameterClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete parameter-like class.
*/
type ParameterClassLike interface {
	// Constructor Methods
	Make(
		name string,
		abstraction AbstractionLike,
	) ParameterLike
}

/*
PrefixClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete prefix-like class.
*/
type PrefixClassLike interface {
	// Constructor Methods
	Make(
		any_ any,
	) PrefixLike
}

/*
PrimaryMethodClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete primary-method-like class.
*/
type PrimaryMethodClassLike interface {
	// Constructor Methods
	Make(
		method MethodLike,
	) PrimaryMethodLike
}

/*
PrimarySubsectionClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete primary-subsection-like class.
*/
type PrimarySubsectionClassLike interface {
	// Constructor Methods
	Make(
		primaryMethods abs.Sequential[PrimaryMethodLike],
	) PrimarySubsectionLike
}

/*
PrimitiveDeclarationsClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete primitive-declarations-like class.
*/
type PrimitiveDeclarationsClassLike interface {
	// Constructor Methods
	Make(
		typeSection TypeSectionLike,
		functionalSection FunctionalSectionLike,
	) PrimitiveDeclarationsLike
}

/*
ResultClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete result-like class.
*/
type ResultClassLike interface {
	// Constructor Methods
	Make(
		any_ any,
	) ResultLike
}

/*
SetterMethodClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete setter-method-like class.
*/
type SetterMethodClassLike interface {
	// Constructor Methods
	Make(
		name string,
		parameter ParameterLike,
	) SetterMethodLike
}

/*
SuffixClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete suffix-like class.
*/
type SuffixClassLike interface {
	// Constructor Methods
	Make(
		name string,
	) SuffixLike
}

/*
TypeDeclarationClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete type-declaration-like class.
*/
type TypeDeclarationClassLike interface {
	// Constructor Methods
	Make(
		declaration DeclarationLike,
		abstraction AbstractionLike,
		optionalEnumeration EnumerationLike,
	) TypeDeclarationLike
}

/*
TypeSectionClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete type-section-like class.
*/
type TypeSectionClassLike interface {
	// Constructor Methods
	Make(
		typeDeclarations abs.Sequential[TypeDeclarationLike],
	) TypeSectionLike
}

/*
ValueClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete value-like class.
*/
type ValueClassLike interface {
	// Constructor Methods
	Make(
		name string,
		abstraction AbstractionLike,
	) ValueLike
}

// Instance Declarations

/*
AbstractionLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete abstraction-like class.
*/
type AbstractionLike interface {
	// Primary Methods
	GetClass() AbstractionClassLike

	// Attribute Methods
	GetOptionalPrefix() PrefixLike
	GetName() string
	GetOptionalSuffix() SuffixLike
	GetOptionalArguments() ArgumentsLike
}

/*
AdditionalArgumentLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete additional-argument-like class.
*/
type AdditionalArgumentLike interface {
	// Primary Methods
	GetClass() AdditionalArgumentClassLike

	// Attribute Methods
	GetArgument() ArgumentLike
}

/*
AdditionalConstraintLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete additional-constraint-like class.
*/
type AdditionalConstraintLike interface {
	// Primary Methods
	GetClass() AdditionalConstraintClassLike

	// Attribute Methods
	GetConstraint() ConstraintLike
}

/*
AdditionalValueLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete additional-value-like class.
*/
type AdditionalValueLike interface {
	// Primary Methods
	GetClass() AdditionalValueClassLike

	// Attribute Methods
	GetName() string
}

/*
ArgumentLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete argument-like class.
*/
type ArgumentLike interface {
	// Primary Methods
	GetClass() ArgumentClassLike

	// Attribute Methods
	GetAbstraction() AbstractionLike
}

/*
ArgumentsLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete arguments-like class.
*/
type ArgumentsLike interface {
	// Primary Methods
	GetClass() ArgumentsClassLike

	// Attribute Methods
	GetArgument() ArgumentLike
	GetAdditionalArguments() abs.Sequential[AdditionalArgumentLike]
}

/*
ArrayLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete array-like class.
*/
type ArrayLike interface {
	// Primary Methods
	GetClass() ArrayClassLike
}

/*
AspectDeclarationLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete aspect-declaration-like class.
*/
type AspectDeclarationLike interface {
	// Primary Methods
	GetClass() AspectDeclarationClassLike

	// Attribute Methods
	GetDeclaration() DeclarationLike
	GetAspectMethods() abs.Sequential[AspectMethodLike]
}

/*
AspectInterfaceLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete aspect-interface-like class.
*/
type AspectInterfaceLike interface {
	// Primary Methods
	GetClass() AspectInterfaceClassLike

	// Attribute Methods
	GetAbstraction() AbstractionLike
}

/*
AspectMethodLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete aspect-method-like class.
*/
type AspectMethodLike interface {
	// Primary Methods
	GetClass() AspectMethodClassLike

	// Attribute Methods
	GetMethod() MethodLike
}

/*
AspectSectionLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete aspect-section-like class.
*/
type AspectSectionLike interface {
	// Primary Methods
	GetClass() AspectSectionClassLike

	// Attribute Methods
	GetAspectDeclarations() abs.Sequential[AspectDeclarationLike]
}

/*
AspectSubsectionLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete aspect-subsection-like class.
*/
type AspectSubsectionLike interface {
	// Primary Methods
	GetClass() AspectSubsectionClassLike

	// Attribute Methods
	GetAspectInterfaces() abs.Sequential[AspectInterfaceLike]
}

/*
AttributeMethodLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete attribute-method-like class.
*/
type AttributeMethodLike interface {
	// Primary Methods
	GetClass() AttributeMethodClassLike

	// Attribute Methods
	GetAny() any
}

/*
AttributeSubsectionLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete attribute-subsection-like class.
*/
type AttributeSubsectionLike interface {
	// Primary Methods
	GetClass() AttributeSubsectionClassLike

	// Attribute Methods
	GetAttributeMethods() abs.Sequential[AttributeMethodLike]
}

/*
ChannelLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete channel-like class.
*/
type ChannelLike interface {
	// Primary Methods
	GetClass() ChannelClassLike
}

/*
ClassDeclarationLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete class-declaration-like class.
*/
type ClassDeclarationLike interface {
	// Primary Methods
	GetClass() ClassDeclarationClassLike

	// Attribute Methods
	GetDeclaration() DeclarationLike
	GetClassMethods() ClassMethodsLike
}

/*
ClassMethodsLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete class-methods-like class.
*/
type ClassMethodsLike interface {
	// Primary Methods
	GetClass() ClassMethodsClassLike

	// Attribute Methods
	GetConstructorSubsection() ConstructorSubsectionLike
	GetOptionalConstantSubsection() ConstantSubsectionLike
	GetOptionalFunctionSubsection() FunctionSubsectionLike
}

/*
ClassSectionLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete class-section-like class.
*/
type ClassSectionLike interface {
	// Primary Methods
	GetClass() ClassSectionClassLike

	// Attribute Methods
	GetClassDeclarations() abs.Sequential[ClassDeclarationLike]
}

/*
ConstantMethodLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete constant-method-like class.
*/
type ConstantMethodLike interface {
	// Primary Methods
	GetClass() ConstantMethodClassLike

	// Attribute Methods
	GetName() string
	GetAbstraction() AbstractionLike
}

/*
ConstantSubsectionLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete constant-subsection-like class.
*/
type ConstantSubsectionLike interface {
	// Primary Methods
	GetClass() ConstantSubsectionClassLike

	// Attribute Methods
	GetConstantMethods() abs.Sequential[ConstantMethodLike]
}

/*
ConstraintLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete constraint-like class.
*/
type ConstraintLike interface {
	// Primary Methods
	GetClass() ConstraintClassLike

	// Attribute Methods
	GetName() string
	GetAbstraction() AbstractionLike
}

/*
ConstraintsLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete constraints-like class.
*/
type ConstraintsLike interface {
	// Primary Methods
	GetClass() ConstraintsClassLike

	// Attribute Methods
	GetConstraint() ConstraintLike
	GetAdditionalConstraints() abs.Sequential[AdditionalConstraintLike]
}

/*
ConstructorMethodLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete constructor-method-like class.
*/
type ConstructorMethodLike interface {
	// Primary Methods
	GetClass() ConstructorMethodClassLike

	// Attribute Methods
	GetName() string
	GetParameters() abs.Sequential[ParameterLike]
	GetAbstraction() AbstractionLike
}

/*
ConstructorSubsectionLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete constructor-subsection-like class.
*/
type ConstructorSubsectionLike interface {
	// Primary Methods
	GetClass() ConstructorSubsectionClassLike

	// Attribute Methods
	GetConstructorMethods() abs.Sequential[ConstructorMethodLike]
}

/*
DeclarationLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete declaration-like class.
*/
type DeclarationLike interface {
	// Primary Methods
	GetClass() DeclarationClassLike

	// Attribute Methods
	GetComment() string
	GetName() string
	GetOptionalConstraints() ConstraintsLike
}

/*
EnumerationLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete enumeration-like class.
*/
type EnumerationLike interface {
	// Primary Methods
	GetClass() EnumerationClassLike

	// Attribute Methods
	GetValue() ValueLike
	GetAdditionalValues() abs.Sequential[AdditionalValueLike]
}

/*
FunctionMethodLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete function-method-like class.
*/
type FunctionMethodLike interface {
	// Primary Methods
	GetClass() FunctionMethodClassLike

	// Attribute Methods
	GetName() string
	GetParameters() abs.Sequential[ParameterLike]
	GetResult() ResultLike
}

/*
FunctionSubsectionLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete function-subsection-like class.
*/
type FunctionSubsectionLike interface {
	// Primary Methods
	GetClass() FunctionSubsectionClassLike

	// Attribute Methods
	GetFunctionMethods() abs.Sequential[FunctionMethodLike]
}

/*
FunctionalDeclarationLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete functional-declaration-like class.
*/
type FunctionalDeclarationLike interface {
	// Primary Methods
	GetClass() FunctionalDeclarationClassLike

	// Attribute Methods
	GetDeclaration() DeclarationLike
	GetParameters() abs.Sequential[ParameterLike]
	GetResult() ResultLike
}

/*
FunctionalSectionLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete functional-section-like class.
*/
type FunctionalSectionLike interface {
	// Primary Methods
	GetClass() FunctionalSectionClassLike

	// Attribute Methods
	GetFunctionalDeclarations() abs.Sequential[FunctionalDeclarationLike]
}

/*
GetterMethodLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete getter-method-like class.
*/
type GetterMethodLike interface {
	// Primary Methods
	GetClass() GetterMethodClassLike

	// Attribute Methods
	GetName() string
	GetAbstraction() AbstractionLike
}

/*
ImportedPackageLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete imported-package-like class.
*/
type ImportedPackageLike interface {
	// Primary Methods
	GetClass() ImportedPackageClassLike

	// Attribute Methods
	GetName() string
	GetPath() string
}

/*
InstanceDeclarationLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete instance-declaration-like class.
*/
type InstanceDeclarationLike interface {
	// Primary Methods
	GetClass() InstanceDeclarationClassLike

	// Attribute Methods
	GetDeclaration() DeclarationLike
	GetInstanceMethods() InstanceMethodsLike
}

/*
InstanceMethodsLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete instance-methods-like class.
*/
type InstanceMethodsLike interface {
	// Primary Methods
	GetClass() InstanceMethodsClassLike

	// Attribute Methods
	GetPrimarySubsection() PrimarySubsectionLike
	GetOptionalAttributeSubsection() AttributeSubsectionLike
	GetOptionalAspectSubsection() AspectSubsectionLike
}

/*
InstanceSectionLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete instance-section-like class.
*/
type InstanceSectionLike interface {
	// Primary Methods
	GetClass() InstanceSectionClassLike

	// Attribute Methods
	GetInstanceDeclarations() abs.Sequential[InstanceDeclarationLike]
}

/*
InterfaceDeclarationsLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete interface-declarations-like class.
*/
type InterfaceDeclarationsLike interface {
	// Primary Methods
	GetClass() InterfaceDeclarationsClassLike

	// Attribute Methods
	GetClassSection() ClassSectionLike
	GetInstanceSection() InstanceSectionLike
	GetAspectSection() AspectSectionLike
}

/*
LegalNoticeLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete legal-notice-like class.
*/
type LegalNoticeLike interface {
	// Primary Methods
	GetClass() LegalNoticeClassLike

	// Attribute Methods
	GetComment() string
}

/*
MapLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete map-like class.
*/
type MapLike interface {
	// Primary Methods
	GetClass() MapClassLike

	// Attribute Methods
	GetName() string
}

/*
MethodLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete method-like class.
*/
type MethodLike interface {
	// Primary Methods
	GetClass() MethodClassLike

	// Attribute Methods
	GetName() string
	GetParameters() abs.Sequential[ParameterLike]
	GetOptionalResult() ResultLike
}

/*
ModelLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete model-like class.
*/
type ModelLike interface {
	// Primary Methods
	GetClass() ModelClassLike

	// Attribute Methods
	GetModuleDeclaration() ModuleDeclarationLike
	GetPrimitiveDeclarations() PrimitiveDeclarationsLike
	GetInterfaceDeclarations() InterfaceDeclarationsLike
}

/*
ModuleDeclarationLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete module-declaration-like class.
*/
type ModuleDeclarationLike interface {
	// Primary Methods
	GetClass() ModuleDeclarationClassLike

	// Attribute Methods
	GetLegalNotice() LegalNoticeLike
	GetModuleHeader() ModuleHeaderLike
	GetModuleImports() ModuleImportsLike
}

/*
ModuleHeaderLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete module-header-like class.
*/
type ModuleHeaderLike interface {
	// Primary Methods
	GetClass() ModuleHeaderClassLike

	// Attribute Methods
	GetComment() string
	GetName() string
}

/*
ModuleImportsLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete module-imports-like class.
*/
type ModuleImportsLike interface {
	// Primary Methods
	GetClass() ModuleImportsClassLike

	// Attribute Methods
	GetImportedPackages() abs.Sequential[ImportedPackageLike]
}

/*
MultivalueLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete multivalue-like class.
*/
type MultivalueLike interface {
	// Primary Methods
	GetClass() MultivalueClassLike

	// Attribute Methods
	GetParameters() abs.Sequential[ParameterLike]
}

/*
NoneLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete none-like class.
*/
type NoneLike interface {
	// Primary Methods
	GetClass() NoneClassLike

	// Attribute Methods
	GetNewline() string
}

/*
ParameterLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete parameter-like class.
*/
type ParameterLike interface {
	// Primary Methods
	GetClass() ParameterClassLike

	// Attribute Methods
	GetName() string
	GetAbstraction() AbstractionLike
}

/*
PrefixLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete prefix-like class.
*/
type PrefixLike interface {
	// Primary Methods
	GetClass() PrefixClassLike

	// Attribute Methods
	GetAny() any
}

/*
PrimaryMethodLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete primary-method-like class.
*/
type PrimaryMethodLike interface {
	// Primary Methods
	GetClass() PrimaryMethodClassLike

	// Attribute Methods
	GetMethod() MethodLike
}

/*
PrimarySubsectionLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete primary-subsection-like class.
*/
type PrimarySubsectionLike interface {
	// Primary Methods
	GetClass() PrimarySubsectionClassLike

	// Attribute Methods
	GetPrimaryMethods() abs.Sequential[PrimaryMethodLike]
}

/*
PrimitiveDeclarationsLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete primitive-declarations-like class.
*/
type PrimitiveDeclarationsLike interface {
	// Primary Methods
	GetClass() PrimitiveDeclarationsClassLike

	// Attribute Methods
	GetTypeSection() TypeSectionLike
	GetFunctionalSection() FunctionalSectionLike
}

/*
ResultLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete result-like class.
*/
type ResultLike interface {
	// Primary Methods
	GetClass() ResultClassLike

	// Attribute Methods
	GetAny() any
}

/*
SetterMethodLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete setter-method-like class.
*/
type SetterMethodLike interface {
	// Primary Methods
	GetClass() SetterMethodClassLike

	// Attribute Methods
	GetName() string
	GetParameter() ParameterLike
}

/*
SuffixLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete suffix-like class.
*/
type SuffixLike interface {
	// Primary Methods
	GetClass() SuffixClassLike

	// Attribute Methods
	GetName() string
}

/*
TypeDeclarationLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete type-declaration-like class.
*/
type TypeDeclarationLike interface {
	// Primary Methods
	GetClass() TypeDeclarationClassLike

	// Attribute Methods
	GetDeclaration() DeclarationLike
	GetAbstraction() AbstractionLike
	GetOptionalEnumeration() EnumerationLike
}

/*
TypeSectionLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete type-section-like class.
*/
type TypeSectionLike interface {
	// Primary Methods
	GetClass() TypeSectionClassLike

	// Attribute Methods
	GetTypeDeclarations() abs.Sequential[TypeDeclarationLike]
}

/*
ValueLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete value-like class.
*/
type ValueLike interface {
	// Primary Methods
	GetClass() ValueClassLike

	// Attribute Methods
	GetName() string
	GetAbstraction() AbstractionLike
}

// Aspect Declarations
