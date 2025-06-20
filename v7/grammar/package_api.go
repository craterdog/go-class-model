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
│            This "package_api.go" file was automatically generated.           │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘

Package "grammar" provides the following grammar classes that operate on the
abstract syntax tree (AST) for this module:
  - Token captures the attributes associated with a parsed token.
  - Scanner is used to scan the source byte stream and recognize matching tokens.
  - Parser is used to process the token stream and generate the AST.
  - Validator is used to validate the semantics associated with an AST.
  - Formatter is used to format an AST back into a canonical version of its source.
  - Visitor walks the AST and calls processor methods for each node in the tree.
  - Processor provides empty processor methods to be inherited by the processors.

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
package grammar

import (
	ast "github.com/craterdog/go-class-model/v7/ast"
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE DECLARATIONS

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	CommentToken
	DelimiterToken
	NameToken
	NewlineToken
	PathToken
	PrefixToken
	SpaceToken
)

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
FormatterClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructor Methods
	Formatter() FormatterLike
}

/*
ParserClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructor Methods
	Parser() ParserLike
}

/*
ProcessorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete processor-like class.
*/
type ProcessorClassLike interface {
	// Constructor Methods
	Processor() ProcessorLike
}

/*
ScannerClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

FormatType() returns the string version of the token type.

MatchesType() determines whether or not a token value is of a specified type.
*/
type ScannerClassLike interface {
	// Constructor Methods
	Scanner(
		source string,
		tokens fra.QueueLike[TokenLike],
	) ScannerLike

	// Function Methods
	FormatToken(
		token TokenLike,
	) string
	FormatType(
		tokenType TokenType,
	) string
	MatchesType(
		tokenValue string,
		tokenType TokenType,
	) bool
}

/*
TokenClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete visitor-like class.
*/
type TokenClassLike interface {
	// Constructor Methods
	Token(
		line uint,
		position uint,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
ValidatorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructor Methods
	Validator() ValidatorLike
}

/*
VisitorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete visitor-like class.
*/
type VisitorClassLike interface {
	// Constructor Methods
	Visitor(
		processor Methodical,
	) VisitorLike
}

// INSTANCE DECLARATIONS

/*
FormatterLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Principal Methods
	GetClass() FormatterClassLike
	FormatModel(
		model ast.ModelLike,
	) string

	// Aspect Interfaces
	Methodical
}

/*
ParserLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Principal Methods
	GetClass() ParserClassLike
	ParseSource(
		source string,
	) ast.ModelLike
}

/*
ProcessorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete processor-like class.
*/
type ProcessorLike interface {
	// Principal Methods
	GetClass() ProcessorClassLike

	// Aspect Interfaces
	Methodical
}

/*
ScannerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
	// Principal Methods
	GetClass() ScannerClassLike
}

/*
TokenLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Principal Methods
	GetClass() TokenClassLike

	// Attribute Methods
	GetLine() uint
	GetPosition() uint
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Principal Methods
	GetClass() ValidatorClassLike
	ValidateModel(
		model ast.ModelLike,
	)

	// Aspect Interfaces
	Methodical
}

/*
VisitorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete visitor-like class.
*/
type VisitorLike interface {
	// Principal Methods
	GetClass() VisitorClassLike
	VisitModel(
		model ast.ModelLike,
	)
}

// ASPECT DECLARATIONS

/*
Methodical declares the set of method signatures that must be supported by
all methodical processors.
*/
type Methodical interface {
	ProcessComment(
		comment string,
	)
	ProcessDelimiter(
		delimiter string,
	)
	ProcessName(
		name string,
	)
	ProcessNewline(
		newline string,
	)
	ProcessPath(
		path string,
	)
	ProcessPrefix(
		prefix string,
	)
	ProcessSpace(
		space string,
	)
	PreprocessAbstraction(
		abstraction ast.AbstractionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAbstraction(
		abstraction ast.AbstractionLike,
		index_ uint,
		count_ uint,
	)
	ProcessAbstractionSlot(
		abstraction ast.AbstractionLike,
		slot_ uint,
	)
	PreprocessAdditionalArgument(
		additionalArgument ast.AdditionalArgumentLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAdditionalArgument(
		additionalArgument ast.AdditionalArgumentLike,
		index_ uint,
		count_ uint,
	)
	ProcessAdditionalArgumentSlot(
		additionalArgument ast.AdditionalArgumentLike,
		slot_ uint,
	)
	PreprocessAdditionalConstraint(
		additionalConstraint ast.AdditionalConstraintLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAdditionalConstraint(
		additionalConstraint ast.AdditionalConstraintLike,
		index_ uint,
		count_ uint,
	)
	ProcessAdditionalConstraintSlot(
		additionalConstraint ast.AdditionalConstraintLike,
		slot_ uint,
	)
	PreprocessAdditionalValue(
		additionalValue ast.AdditionalValueLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAdditionalValue(
		additionalValue ast.AdditionalValueLike,
		index_ uint,
		count_ uint,
	)
	ProcessAdditionalValueSlot(
		additionalValue ast.AdditionalValueLike,
		slot_ uint,
	)
	PreprocessArgument(
		argument ast.ArgumentLike,
		index_ uint,
		count_ uint,
	)
	PostprocessArgument(
		argument ast.ArgumentLike,
		index_ uint,
		count_ uint,
	)
	ProcessArgumentSlot(
		argument ast.ArgumentLike,
		slot_ uint,
	)
	PreprocessArguments(
		arguments ast.ArgumentsLike,
		index_ uint,
		count_ uint,
	)
	PostprocessArguments(
		arguments ast.ArgumentsLike,
		index_ uint,
		count_ uint,
	)
	ProcessArgumentsSlot(
		arguments ast.ArgumentsLike,
		slot_ uint,
	)
	PreprocessArray(
		array ast.ArrayLike,
		index_ uint,
		count_ uint,
	)
	PostprocessArray(
		array ast.ArrayLike,
		index_ uint,
		count_ uint,
	)
	ProcessArraySlot(
		array ast.ArrayLike,
		slot_ uint,
	)
	PreprocessAspectDeclaration(
		aspectDeclaration ast.AspectDeclarationLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAspectDeclaration(
		aspectDeclaration ast.AspectDeclarationLike,
		index_ uint,
		count_ uint,
	)
	ProcessAspectDeclarationSlot(
		aspectDeclaration ast.AspectDeclarationLike,
		slot_ uint,
	)
	PreprocessAspectInterface(
		aspectInterface ast.AspectInterfaceLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAspectInterface(
		aspectInterface ast.AspectInterfaceLike,
		index_ uint,
		count_ uint,
	)
	ProcessAspectInterfaceSlot(
		aspectInterface ast.AspectInterfaceLike,
		slot_ uint,
	)
	PreprocessAspectMethod(
		aspectMethod ast.AspectMethodLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAspectMethod(
		aspectMethod ast.AspectMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessAspectMethodSlot(
		aspectMethod ast.AspectMethodLike,
		slot_ uint,
	)
	PreprocessAspectSection(
		aspectSection ast.AspectSectionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAspectSection(
		aspectSection ast.AspectSectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessAspectSectionSlot(
		aspectSection ast.AspectSectionLike,
		slot_ uint,
	)
	PreprocessAspectSubsection(
		aspectSubsection ast.AspectSubsectionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAspectSubsection(
		aspectSubsection ast.AspectSubsectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessAspectSubsectionSlot(
		aspectSubsection ast.AspectSubsectionLike,
		slot_ uint,
	)
	PreprocessAttributeMethod(
		attributeMethod ast.AttributeMethodLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAttributeMethod(
		attributeMethod ast.AttributeMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessAttributeMethodSlot(
		attributeMethod ast.AttributeMethodLike,
		slot_ uint,
	)
	PreprocessAttributeSubsection(
		attributeSubsection ast.AttributeSubsectionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAttributeSubsection(
		attributeSubsection ast.AttributeSubsectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessAttributeSubsectionSlot(
		attributeSubsection ast.AttributeSubsectionLike,
		slot_ uint,
	)
	PreprocessChannel(
		channel ast.ChannelLike,
		index_ uint,
		count_ uint,
	)
	PostprocessChannel(
		channel ast.ChannelLike,
		index_ uint,
		count_ uint,
	)
	ProcessChannelSlot(
		channel ast.ChannelLike,
		slot_ uint,
	)
	PreprocessClassDeclaration(
		classDeclaration ast.ClassDeclarationLike,
		index_ uint,
		count_ uint,
	)
	PostprocessClassDeclaration(
		classDeclaration ast.ClassDeclarationLike,
		index_ uint,
		count_ uint,
	)
	ProcessClassDeclarationSlot(
		classDeclaration ast.ClassDeclarationLike,
		slot_ uint,
	)
	PreprocessClassMethods(
		classMethods ast.ClassMethodsLike,
		index_ uint,
		count_ uint,
	)
	PostprocessClassMethods(
		classMethods ast.ClassMethodsLike,
		index_ uint,
		count_ uint,
	)
	ProcessClassMethodsSlot(
		classMethods ast.ClassMethodsLike,
		slot_ uint,
	)
	PreprocessClassSection(
		classSection ast.ClassSectionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessClassSection(
		classSection ast.ClassSectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessClassSectionSlot(
		classSection ast.ClassSectionLike,
		slot_ uint,
	)
	PreprocessConstantMethod(
		constantMethod ast.ConstantMethodLike,
		index_ uint,
		count_ uint,
	)
	PostprocessConstantMethod(
		constantMethod ast.ConstantMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstantMethodSlot(
		constantMethod ast.ConstantMethodLike,
		slot_ uint,
	)
	PreprocessConstantSubsection(
		constantSubsection ast.ConstantSubsectionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessConstantSubsection(
		constantSubsection ast.ConstantSubsectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstantSubsectionSlot(
		constantSubsection ast.ConstantSubsectionLike,
		slot_ uint,
	)
	PreprocessConstraint(
		constraint ast.ConstraintLike,
		index_ uint,
		count_ uint,
	)
	PostprocessConstraint(
		constraint ast.ConstraintLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstraintSlot(
		constraint ast.ConstraintLike,
		slot_ uint,
	)
	PreprocessConstraints(
		constraints ast.ConstraintsLike,
		index_ uint,
		count_ uint,
	)
	PostprocessConstraints(
		constraints ast.ConstraintsLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstraintsSlot(
		constraints ast.ConstraintsLike,
		slot_ uint,
	)
	PreprocessConstructorMethod(
		constructorMethod ast.ConstructorMethodLike,
		index_ uint,
		count_ uint,
	)
	PostprocessConstructorMethod(
		constructorMethod ast.ConstructorMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstructorMethodSlot(
		constructorMethod ast.ConstructorMethodLike,
		slot_ uint,
	)
	PreprocessConstructorSubsection(
		constructorSubsection ast.ConstructorSubsectionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessConstructorSubsection(
		constructorSubsection ast.ConstructorSubsectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstructorSubsectionSlot(
		constructorSubsection ast.ConstructorSubsectionLike,
		slot_ uint,
	)
	PreprocessDeclaration(
		declaration ast.DeclarationLike,
		index_ uint,
		count_ uint,
	)
	PostprocessDeclaration(
		declaration ast.DeclarationLike,
		index_ uint,
		count_ uint,
	)
	ProcessDeclarationSlot(
		declaration ast.DeclarationLike,
		slot_ uint,
	)
	PreprocessEnumeration(
		enumeration ast.EnumerationLike,
		index_ uint,
		count_ uint,
	)
	PostprocessEnumeration(
		enumeration ast.EnumerationLike,
		index_ uint,
		count_ uint,
	)
	ProcessEnumerationSlot(
		enumeration ast.EnumerationLike,
		slot_ uint,
	)
	PreprocessFunctionMethod(
		functionMethod ast.FunctionMethodLike,
		index_ uint,
		count_ uint,
	)
	PostprocessFunctionMethod(
		functionMethod ast.FunctionMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessFunctionMethodSlot(
		functionMethod ast.FunctionMethodLike,
		slot_ uint,
	)
	PreprocessFunctionSubsection(
		functionSubsection ast.FunctionSubsectionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessFunctionSubsection(
		functionSubsection ast.FunctionSubsectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessFunctionSubsectionSlot(
		functionSubsection ast.FunctionSubsectionLike,
		slot_ uint,
	)
	PreprocessFunctionalDeclaration(
		functionalDeclaration ast.FunctionalDeclarationLike,
		index_ uint,
		count_ uint,
	)
	PostprocessFunctionalDeclaration(
		functionalDeclaration ast.FunctionalDeclarationLike,
		index_ uint,
		count_ uint,
	)
	ProcessFunctionalDeclarationSlot(
		functionalDeclaration ast.FunctionalDeclarationLike,
		slot_ uint,
	)
	PreprocessFunctionalSection(
		functionalSection ast.FunctionalSectionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessFunctionalSection(
		functionalSection ast.FunctionalSectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessFunctionalSectionSlot(
		functionalSection ast.FunctionalSectionLike,
		slot_ uint,
	)
	PreprocessGetterMethod(
		getterMethod ast.GetterMethodLike,
		index_ uint,
		count_ uint,
	)
	PostprocessGetterMethod(
		getterMethod ast.GetterMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessGetterMethodSlot(
		getterMethod ast.GetterMethodLike,
		slot_ uint,
	)
	PreprocessImportList(
		importList ast.ImportListLike,
		index_ uint,
		count_ uint,
	)
	PostprocessImportList(
		importList ast.ImportListLike,
		index_ uint,
		count_ uint,
	)
	ProcessImportListSlot(
		importList ast.ImportListLike,
		slot_ uint,
	)
	PreprocessImportedPackage(
		importedPackage ast.ImportedPackageLike,
		index_ uint,
		count_ uint,
	)
	PostprocessImportedPackage(
		importedPackage ast.ImportedPackageLike,
		index_ uint,
		count_ uint,
	)
	ProcessImportedPackageSlot(
		importedPackage ast.ImportedPackageLike,
		slot_ uint,
	)
	PreprocessInstanceDeclaration(
		instanceDeclaration ast.InstanceDeclarationLike,
		index_ uint,
		count_ uint,
	)
	PostprocessInstanceDeclaration(
		instanceDeclaration ast.InstanceDeclarationLike,
		index_ uint,
		count_ uint,
	)
	ProcessInstanceDeclarationSlot(
		instanceDeclaration ast.InstanceDeclarationLike,
		slot_ uint,
	)
	PreprocessInstanceMethods(
		instanceMethods ast.InstanceMethodsLike,
		index_ uint,
		count_ uint,
	)
	PostprocessInstanceMethods(
		instanceMethods ast.InstanceMethodsLike,
		index_ uint,
		count_ uint,
	)
	ProcessInstanceMethodsSlot(
		instanceMethods ast.InstanceMethodsLike,
		slot_ uint,
	)
	PreprocessInstanceSection(
		instanceSection ast.InstanceSectionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessInstanceSection(
		instanceSection ast.InstanceSectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessInstanceSectionSlot(
		instanceSection ast.InstanceSectionLike,
		slot_ uint,
	)
	PreprocessInterfaceDeclarations(
		interfaceDeclarations ast.InterfaceDeclarationsLike,
		index_ uint,
		count_ uint,
	)
	PostprocessInterfaceDeclarations(
		interfaceDeclarations ast.InterfaceDeclarationsLike,
		index_ uint,
		count_ uint,
	)
	ProcessInterfaceDeclarationsSlot(
		interfaceDeclarations ast.InterfaceDeclarationsLike,
		slot_ uint,
	)
	PreprocessLegalNotice(
		legalNotice ast.LegalNoticeLike,
		index_ uint,
		count_ uint,
	)
	PostprocessLegalNotice(
		legalNotice ast.LegalNoticeLike,
		index_ uint,
		count_ uint,
	)
	ProcessLegalNoticeSlot(
		legalNotice ast.LegalNoticeLike,
		slot_ uint,
	)
	PreprocessMap(
		map_ ast.MapLike,
		index_ uint,
		count_ uint,
	)
	PostprocessMap(
		map_ ast.MapLike,
		index_ uint,
		count_ uint,
	)
	ProcessMapSlot(
		map_ ast.MapLike,
		slot_ uint,
	)
	PreprocessMethod(
		method ast.MethodLike,
		index_ uint,
		count_ uint,
	)
	PostprocessMethod(
		method ast.MethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessMethodSlot(
		method ast.MethodLike,
		slot_ uint,
	)
	PreprocessModel(
		model ast.ModelLike,
		index_ uint,
		count_ uint,
	)
	PostprocessModel(
		model ast.ModelLike,
		index_ uint,
		count_ uint,
	)
	ProcessModelSlot(
		model ast.ModelLike,
		slot_ uint,
	)
	PreprocessMultivalue(
		multivalue ast.MultivalueLike,
		index_ uint,
		count_ uint,
	)
	PostprocessMultivalue(
		multivalue ast.MultivalueLike,
		index_ uint,
		count_ uint,
	)
	ProcessMultivalueSlot(
		multivalue ast.MultivalueLike,
		slot_ uint,
	)
	PreprocessNone(
		none ast.NoneLike,
		index_ uint,
		count_ uint,
	)
	PostprocessNone(
		none ast.NoneLike,
		index_ uint,
		count_ uint,
	)
	ProcessNoneSlot(
		none ast.NoneLike,
		slot_ uint,
	)
	PreprocessPackageDeclaration(
		packageDeclaration ast.PackageDeclarationLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPackageDeclaration(
		packageDeclaration ast.PackageDeclarationLike,
		index_ uint,
		count_ uint,
	)
	ProcessPackageDeclarationSlot(
		packageDeclaration ast.PackageDeclarationLike,
		slot_ uint,
	)
	PreprocessPackageHeader(
		packageHeader ast.PackageHeaderLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPackageHeader(
		packageHeader ast.PackageHeaderLike,
		index_ uint,
		count_ uint,
	)
	ProcessPackageHeaderSlot(
		packageHeader ast.PackageHeaderLike,
		slot_ uint,
	)
	PreprocessPackageImports(
		packageImports ast.PackageImportsLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPackageImports(
		packageImports ast.PackageImportsLike,
		index_ uint,
		count_ uint,
	)
	ProcessPackageImportsSlot(
		packageImports ast.PackageImportsLike,
		slot_ uint,
	)
	PreprocessParameter(
		parameter ast.ParameterLike,
		index_ uint,
		count_ uint,
	)
	PostprocessParameter(
		parameter ast.ParameterLike,
		index_ uint,
		count_ uint,
	)
	ProcessParameterSlot(
		parameter ast.ParameterLike,
		slot_ uint,
	)
	PreprocessParameterList(
		parameterList ast.ParameterListLike,
		index_ uint,
		count_ uint,
	)
	PostprocessParameterList(
		parameterList ast.ParameterListLike,
		index_ uint,
		count_ uint,
	)
	ProcessParameterListSlot(
		parameterList ast.ParameterListLike,
		slot_ uint,
	)
	PreprocessPrimitiveDeclarations(
		primitiveDeclarations ast.PrimitiveDeclarationsLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPrimitiveDeclarations(
		primitiveDeclarations ast.PrimitiveDeclarationsLike,
		index_ uint,
		count_ uint,
	)
	ProcessPrimitiveDeclarationsSlot(
		primitiveDeclarations ast.PrimitiveDeclarationsLike,
		slot_ uint,
	)
	PreprocessPrincipalMethod(
		principalMethod ast.PrincipalMethodLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPrincipalMethod(
		principalMethod ast.PrincipalMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessPrincipalMethodSlot(
		principalMethod ast.PrincipalMethodLike,
		slot_ uint,
	)
	PreprocessPrincipalSubsection(
		principalSubsection ast.PrincipalSubsectionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPrincipalSubsection(
		principalSubsection ast.PrincipalSubsectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessPrincipalSubsectionSlot(
		principalSubsection ast.PrincipalSubsectionLike,
		slot_ uint,
	)
	PreprocessResult(
		result ast.ResultLike,
		index_ uint,
		count_ uint,
	)
	PostprocessResult(
		result ast.ResultLike,
		index_ uint,
		count_ uint,
	)
	ProcessResultSlot(
		result ast.ResultLike,
		slot_ uint,
	)
	PreprocessSetterMethod(
		setterMethod ast.SetterMethodLike,
		index_ uint,
		count_ uint,
	)
	PostprocessSetterMethod(
		setterMethod ast.SetterMethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessSetterMethodSlot(
		setterMethod ast.SetterMethodLike,
		slot_ uint,
	)
	PreprocessStar(
		star ast.StarLike,
		index_ uint,
		count_ uint,
	)
	PostprocessStar(
		star ast.StarLike,
		index_ uint,
		count_ uint,
	)
	ProcessStarSlot(
		star ast.StarLike,
		slot_ uint,
	)
	PreprocessTypeDeclaration(
		typeDeclaration ast.TypeDeclarationLike,
		index_ uint,
		count_ uint,
	)
	PostprocessTypeDeclaration(
		typeDeclaration ast.TypeDeclarationLike,
		index_ uint,
		count_ uint,
	)
	ProcessTypeDeclarationSlot(
		typeDeclaration ast.TypeDeclarationLike,
		slot_ uint,
	)
	PreprocessTypeSection(
		typeSection ast.TypeSectionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessTypeSection(
		typeSection ast.TypeSectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessTypeSectionSlot(
		typeSection ast.TypeSectionLike,
		slot_ uint,
	)
	PreprocessValue(
		value ast.ValueLike,
		index_ uint,
		count_ uint,
	)
	PostprocessValue(
		value ast.ValueLike,
		index_ uint,
		count_ uint,
	)
	ProcessValueSlot(
		value ast.ValueLike,
		slot_ uint,
	)
	PreprocessWrapper(
		wrapper ast.WrapperLike,
		index_ uint,
		count_ uint,
	)
	PostprocessWrapper(
		wrapper ast.WrapperLike,
		index_ uint,
		count_ uint,
	)
	ProcessWrapperSlot(
		wrapper ast.WrapperLike,
		slot_ uint,
	)
}
