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
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	ast "github.com/craterdog/go-class-model/v8/ast"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func VisitorClass() VisitorClassLike {
	return visitorClass()
}

// Constructor Methods

func (c *visitorClass_) Visitor(
	processor Methodical,
) VisitorLike {
	if uti.IsUndefined(processor) {
		panic("The \"processor\" attribute is required by this class.")
	}
	var instance = &visitor_{
		// Initialize the instance attributes.
		processor_: processor,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *visitor_) GetClass() VisitorClassLike {
	return visitorClass()
}

func (v *visitor_) VisitModel(
	model ast.ModelLike,
) {
	v.processor_.PreprocessModel(
		model,
		0,
		0,
	)
	v.visitModel(model)
	v.processor_.PostprocessModel(
		model,
		0,
		0,
	)
}

// PROTECTED INTERFACE

// Private Methods

func (v *visitor_) visitAbstraction(
	abstraction ast.AbstractionLike,
) {
	var optionalWrapper = abstraction.GetOptionalWrapper()
	if uti.IsDefined(optionalWrapper) {
		v.processor_.PreprocessWrapper(
			optionalWrapper,
			0,
			0,
		)
		v.visitWrapper(optionalWrapper)
		v.processor_.PostprocessWrapper(
			optionalWrapper,
			0,
			0,
		)
	}
	// Visit slot 1 between terms.
	v.processor_.ProcessAbstractionSlot(
		abstraction,
		1,
	)

	var type_ = abstraction.GetType()
	v.processor_.PreprocessType(
		type_,
		0,
		0,
	)
	v.visitType(type_)
	v.processor_.PostprocessType(
		type_,
		0,
		0,
	)
}

func (v *visitor_) visitAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
) {
	var delimiter = additionalArgument.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessAdditionalArgumentSlot(
		additionalArgument,
		1,
	)

	var argument = additionalArgument.GetArgument()
	v.processor_.PreprocessArgument(
		argument,
		0,
		0,
	)
	v.visitArgument(argument)
	v.processor_.PostprocessArgument(
		argument,
		0,
		0,
	)
}

func (v *visitor_) visitAdditionalConstraint(
	additionalConstraint ast.AdditionalConstraintLike,
) {
	var delimiter = additionalConstraint.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessAdditionalConstraintSlot(
		additionalConstraint,
		1,
	)

	var constraint = additionalConstraint.GetConstraint()
	v.processor_.PreprocessConstraint(
		constraint,
		0,
		0,
	)
	v.visitConstraint(constraint)
	v.processor_.PostprocessConstraint(
		constraint,
		0,
		0,
	)
}

func (v *visitor_) visitAdditionalValue(
	additionalValue ast.AdditionalValueLike,
) {
	var name = additionalValue.GetName()
	v.processor_.ProcessName(name)
}

func (v *visitor_) visitArgument(
	argument ast.ArgumentLike,
) {
	var abstraction = argument.GetAbstraction()
	v.processor_.PreprocessAbstraction(
		abstraction,
		0,
		0,
	)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(
		abstraction,
		0,
		0,
	)
}

func (v *visitor_) visitArguments(
	arguments ast.ArgumentsLike,
) {
	var delimiter1 = arguments.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessArgumentsSlot(
		arguments,
		1,
	)

	var argument = arguments.GetArgument()
	v.processor_.PreprocessArgument(
		argument,
		0,
		0,
	)
	v.visitArgument(argument)
	v.processor_.PostprocessArgument(
		argument,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessArgumentsSlot(
		arguments,
		2,
	)

	var additionalArgumentsIndex uint
	var additionalArguments = arguments.GetAdditionalArguments().GetIterator()
	var additionalArgumentsCount = uint(additionalArguments.GetSize())
	for additionalArguments.HasNext() {
		additionalArgumentsIndex++
		var rule = additionalArguments.GetNext()
		v.processor_.PreprocessAdditionalArgument(
			rule,
			additionalArgumentsIndex,
			additionalArgumentsCount,
		)
		v.visitAdditionalArgument(rule)
		v.processor_.PostprocessAdditionalArgument(
			rule,
			additionalArgumentsIndex,
			additionalArgumentsCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessArgumentsSlot(
		arguments,
		3,
	)

	var delimiter2 = arguments.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitArray(
	array ast.ArrayLike,
) {
	var delimiter1 = array.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessArraySlot(
		array,
		1,
	)

	var delimiter2 = array.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitAspectDeclaration(
	aspectDeclaration ast.AspectDeclarationLike,
) {
	var declaration = aspectDeclaration.GetDeclaration()
	v.processor_.PreprocessDeclaration(
		declaration,
		0,
		0,
	)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(
		declaration,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessAspectDeclarationSlot(
		aspectDeclaration,
		1,
	)

	var delimiter1 = aspectDeclaration.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 2 between terms.
	v.processor_.ProcessAspectDeclarationSlot(
		aspectDeclaration,
		2,
	)

	var delimiter2 = aspectDeclaration.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessAspectDeclarationSlot(
		aspectDeclaration,
		3,
	)

	var aspectMethodsIndex uint
	var aspectMethods = aspectDeclaration.GetAspectMethods().GetIterator()
	var aspectMethodsCount = uint(aspectMethods.GetSize())
	for aspectMethods.HasNext() {
		aspectMethodsIndex++
		var rule = aspectMethods.GetNext()
		v.processor_.PreprocessAspectMethod(
			rule,
			aspectMethodsIndex,
			aspectMethodsCount,
		)
		v.visitAspectMethod(rule)
		v.processor_.PostprocessAspectMethod(
			rule,
			aspectMethodsIndex,
			aspectMethodsCount,
		)
	}
	// Visit slot 4 between terms.
	v.processor_.ProcessAspectDeclarationSlot(
		aspectDeclaration,
		4,
	)

	var delimiter3 = aspectDeclaration.GetDelimiter3()
	v.processor_.ProcessDelimiter(delimiter3)
}

func (v *visitor_) visitAspectInterface(
	aspectInterface ast.AspectInterfaceLike,
) {
	var abstraction = aspectInterface.GetAbstraction()
	v.processor_.PreprocessAbstraction(
		abstraction,
		0,
		0,
	)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(
		abstraction,
		0,
		0,
	)
}

func (v *visitor_) visitAspectMethod(
	aspectMethod ast.AspectMethodLike,
) {
	var method = aspectMethod.GetMethod()
	v.processor_.PreprocessMethod(
		method,
		0,
		0,
	)
	v.visitMethod(method)
	v.processor_.PostprocessMethod(
		method,
		0,
		0,
	)
}

func (v *visitor_) visitAspectSection(
	aspectSection ast.AspectSectionLike,
) {
	var delimiter = aspectSection.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessAspectSectionSlot(
		aspectSection,
		1,
	)

	var aspectDeclarationsIndex uint
	var aspectDeclarations = aspectSection.GetAspectDeclarations().GetIterator()
	var aspectDeclarationsCount = uint(aspectDeclarations.GetSize())
	for aspectDeclarations.HasNext() {
		aspectDeclarationsIndex++
		var rule = aspectDeclarations.GetNext()
		v.processor_.PreprocessAspectDeclaration(
			rule,
			aspectDeclarationsIndex,
			aspectDeclarationsCount,
		)
		v.visitAspectDeclaration(rule)
		v.processor_.PostprocessAspectDeclaration(
			rule,
			aspectDeclarationsIndex,
			aspectDeclarationsCount,
		)
	}
}

func (v *visitor_) visitAspectSubsection(
	aspectSubsection ast.AspectSubsectionLike,
) {
	var delimiter = aspectSubsection.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessAspectSubsectionSlot(
		aspectSubsection,
		1,
	)

	var aspectInterfacesIndex uint
	var aspectInterfaces = aspectSubsection.GetAspectInterfaces().GetIterator()
	var aspectInterfacesCount = uint(aspectInterfaces.GetSize())
	for aspectInterfaces.HasNext() {
		aspectInterfacesIndex++
		var rule = aspectInterfaces.GetNext()
		v.processor_.PreprocessAspectInterface(
			rule,
			aspectInterfacesIndex,
			aspectInterfacesCount,
		)
		v.visitAspectInterface(rule)
		v.processor_.PostprocessAspectInterface(
			rule,
			aspectInterfacesIndex,
			aspectInterfacesCount,
		)
	}
}

func (v *visitor_) visitAttributeMethod(
	attributeMethod ast.AttributeMethodLike,
) {
	// Visit the possible attributeMethod rule types.
	switch actual := attributeMethod.GetAny().(type) {
	case ast.GetterMethodLike:
		v.processor_.PreprocessGetterMethod(
			actual,
			0,
			0,
		)
		v.visitGetterMethod(actual)
		v.processor_.PostprocessGetterMethod(
			actual,
			0,
			0,
		)
	case ast.SetterMethodLike:
		v.processor_.PreprocessSetterMethod(
			actual,
			0,
			0,
		)
		v.visitSetterMethod(actual)
		v.processor_.PostprocessSetterMethod(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitAttributeSubsection(
	attributeSubsection ast.AttributeSubsectionLike,
) {
	var delimiter = attributeSubsection.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessAttributeSubsectionSlot(
		attributeSubsection,
		1,
	)

	var attributeMethodsIndex uint
	var attributeMethods = attributeSubsection.GetAttributeMethods().GetIterator()
	var attributeMethodsCount = uint(attributeMethods.GetSize())
	for attributeMethods.HasNext() {
		attributeMethodsIndex++
		var rule = attributeMethods.GetNext()
		v.processor_.PreprocessAttributeMethod(
			rule,
			attributeMethodsIndex,
			attributeMethodsCount,
		)
		v.visitAttributeMethod(rule)
		v.processor_.PostprocessAttributeMethod(
			rule,
			attributeMethodsIndex,
			attributeMethodsCount,
		)
	}
}

func (v *visitor_) visitChannel(
	channel ast.ChannelLike,
) {
	var delimiter = channel.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
}

func (v *visitor_) visitClassDeclaration(
	classDeclaration ast.ClassDeclarationLike,
) {
	var declaration = classDeclaration.GetDeclaration()
	v.processor_.PreprocessDeclaration(
		declaration,
		0,
		0,
	)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(
		declaration,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessClassDeclarationSlot(
		classDeclaration,
		1,
	)

	var delimiter1 = classDeclaration.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 2 between terms.
	v.processor_.ProcessClassDeclarationSlot(
		classDeclaration,
		2,
	)

	var delimiter2 = classDeclaration.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessClassDeclarationSlot(
		classDeclaration,
		3,
	)

	var classMethods = classDeclaration.GetClassMethods()
	v.processor_.PreprocessClassMethods(
		classMethods,
		0,
		0,
	)
	v.visitClassMethods(classMethods)
	v.processor_.PostprocessClassMethods(
		classMethods,
		0,
		0,
	)
	// Visit slot 4 between terms.
	v.processor_.ProcessClassDeclarationSlot(
		classDeclaration,
		4,
	)

	var delimiter3 = classDeclaration.GetDelimiter3()
	v.processor_.ProcessDelimiter(delimiter3)
}

func (v *visitor_) visitClassMethods(
	classMethods ast.ClassMethodsLike,
) {
	var constructorSubsection = classMethods.GetConstructorSubsection()
	v.processor_.PreprocessConstructorSubsection(
		constructorSubsection,
		0,
		0,
	)
	v.visitConstructorSubsection(constructorSubsection)
	v.processor_.PostprocessConstructorSubsection(
		constructorSubsection,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessClassMethodsSlot(
		classMethods,
		1,
	)

	var optionalConstantSubsection = classMethods.GetOptionalConstantSubsection()
	if uti.IsDefined(optionalConstantSubsection) {
		v.processor_.PreprocessConstantSubsection(
			optionalConstantSubsection,
			0,
			0,
		)
		v.visitConstantSubsection(optionalConstantSubsection)
		v.processor_.PostprocessConstantSubsection(
			optionalConstantSubsection,
			0,
			0,
		)
	}
	// Visit slot 2 between terms.
	v.processor_.ProcessClassMethodsSlot(
		classMethods,
		2,
	)

	var optionalFunctionSubsection = classMethods.GetOptionalFunctionSubsection()
	if uti.IsDefined(optionalFunctionSubsection) {
		v.processor_.PreprocessFunctionSubsection(
			optionalFunctionSubsection,
			0,
			0,
		)
		v.visitFunctionSubsection(optionalFunctionSubsection)
		v.processor_.PostprocessFunctionSubsection(
			optionalFunctionSubsection,
			0,
			0,
		)
	}
}

func (v *visitor_) visitClassSection(
	classSection ast.ClassSectionLike,
) {
	var delimiter = classSection.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessClassSectionSlot(
		classSection,
		1,
	)

	var classDeclarationsIndex uint
	var classDeclarations = classSection.GetClassDeclarations().GetIterator()
	var classDeclarationsCount = uint(classDeclarations.GetSize())
	for classDeclarations.HasNext() {
		classDeclarationsIndex++
		var rule = classDeclarations.GetNext()
		v.processor_.PreprocessClassDeclaration(
			rule,
			classDeclarationsIndex,
			classDeclarationsCount,
		)
		v.visitClassDeclaration(rule)
		v.processor_.PostprocessClassDeclaration(
			rule,
			classDeclarationsIndex,
			classDeclarationsCount,
		)
	}
}

func (v *visitor_) visitConstantMethod(
	constantMethod ast.ConstantMethodLike,
) {
	var name = constantMethod.GetName()
	v.processor_.ProcessName(name)
	// Visit slot 1 between terms.
	v.processor_.ProcessConstantMethodSlot(
		constantMethod,
		1,
	)

	var delimiter1 = constantMethod.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 2 between terms.
	v.processor_.ProcessConstantMethodSlot(
		constantMethod,
		2,
	)

	var delimiter2 = constantMethod.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessConstantMethodSlot(
		constantMethod,
		3,
	)

	var abstraction = constantMethod.GetAbstraction()
	v.processor_.PreprocessAbstraction(
		abstraction,
		0,
		0,
	)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(
		abstraction,
		0,
		0,
	)
}

func (v *visitor_) visitConstantSubsection(
	constantSubsection ast.ConstantSubsectionLike,
) {
	var delimiter = constantSubsection.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessConstantSubsectionSlot(
		constantSubsection,
		1,
	)

	var constantMethodsIndex uint
	var constantMethods = constantSubsection.GetConstantMethods().GetIterator()
	var constantMethodsCount = uint(constantMethods.GetSize())
	for constantMethods.HasNext() {
		constantMethodsIndex++
		var rule = constantMethods.GetNext()
		v.processor_.PreprocessConstantMethod(
			rule,
			constantMethodsIndex,
			constantMethodsCount,
		)
		v.visitConstantMethod(rule)
		v.processor_.PostprocessConstantMethod(
			rule,
			constantMethodsIndex,
			constantMethodsCount,
		)
	}
}

func (v *visitor_) visitConstraint(
	constraint ast.ConstraintLike,
) {
	var name = constraint.GetName()
	v.processor_.ProcessName(name)
	// Visit slot 1 between terms.
	v.processor_.ProcessConstraintSlot(
		constraint,
		1,
	)

	var abstraction = constraint.GetAbstraction()
	v.processor_.PreprocessAbstraction(
		abstraction,
		0,
		0,
	)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(
		abstraction,
		0,
		0,
	)
}

func (v *visitor_) visitConstraints(
	constraints ast.ConstraintsLike,
) {
	var delimiter1 = constraints.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessConstraintsSlot(
		constraints,
		1,
	)

	var constraint = constraints.GetConstraint()
	v.processor_.PreprocessConstraint(
		constraint,
		0,
		0,
	)
	v.visitConstraint(constraint)
	v.processor_.PostprocessConstraint(
		constraint,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessConstraintsSlot(
		constraints,
		2,
	)

	var additionalConstraintsIndex uint
	var additionalConstraints = constraints.GetAdditionalConstraints().GetIterator()
	var additionalConstraintsCount = uint(additionalConstraints.GetSize())
	for additionalConstraints.HasNext() {
		additionalConstraintsIndex++
		var rule = additionalConstraints.GetNext()
		v.processor_.PreprocessAdditionalConstraint(
			rule,
			additionalConstraintsIndex,
			additionalConstraintsCount,
		)
		v.visitAdditionalConstraint(rule)
		v.processor_.PostprocessAdditionalConstraint(
			rule,
			additionalConstraintsIndex,
			additionalConstraintsCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessConstraintsSlot(
		constraints,
		3,
	)

	var delimiter2 = constraints.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitConstructorMethod(
	constructorMethod ast.ConstructorMethodLike,
) {
	var name = constructorMethod.GetName()
	v.processor_.ProcessName(name)
	// Visit slot 1 between terms.
	v.processor_.ProcessConstructorMethodSlot(
		constructorMethod,
		1,
	)

	var delimiter1 = constructorMethod.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 2 between terms.
	v.processor_.ProcessConstructorMethodSlot(
		constructorMethod,
		2,
	)

	var optionalParameterList = constructorMethod.GetOptionalParameterList()
	if uti.IsDefined(optionalParameterList) {
		v.processor_.PreprocessParameterList(
			optionalParameterList,
			0,
			0,
		)
		v.visitParameterList(optionalParameterList)
		v.processor_.PostprocessParameterList(
			optionalParameterList,
			0,
			0,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessConstructorMethodSlot(
		constructorMethod,
		3,
	)

	var delimiter2 = constructorMethod.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 4 between terms.
	v.processor_.ProcessConstructorMethodSlot(
		constructorMethod,
		4,
	)

	var abstraction = constructorMethod.GetAbstraction()
	v.processor_.PreprocessAbstraction(
		abstraction,
		0,
		0,
	)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(
		abstraction,
		0,
		0,
	)
}

func (v *visitor_) visitConstructorSubsection(
	constructorSubsection ast.ConstructorSubsectionLike,
) {
	var delimiter = constructorSubsection.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessConstructorSubsectionSlot(
		constructorSubsection,
		1,
	)

	var constructorMethodsIndex uint
	var constructorMethods = constructorSubsection.GetConstructorMethods().GetIterator()
	var constructorMethodsCount = uint(constructorMethods.GetSize())
	for constructorMethods.HasNext() {
		constructorMethodsIndex++
		var rule = constructorMethods.GetNext()
		v.processor_.PreprocessConstructorMethod(
			rule,
			constructorMethodsIndex,
			constructorMethodsCount,
		)
		v.visitConstructorMethod(rule)
		v.processor_.PostprocessConstructorMethod(
			rule,
			constructorMethodsIndex,
			constructorMethodsCount,
		)
	}
}

func (v *visitor_) visitDeclaration(
	declaration ast.DeclarationLike,
) {
	var comment = declaration.GetComment()
	v.processor_.ProcessComment(comment)
	// Visit slot 1 between terms.
	v.processor_.ProcessDeclarationSlot(
		declaration,
		1,
	)

	var delimiter = declaration.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 2 between terms.
	v.processor_.ProcessDeclarationSlot(
		declaration,
		2,
	)

	var name = declaration.GetName()
	v.processor_.ProcessName(name)
	// Visit slot 3 between terms.
	v.processor_.ProcessDeclarationSlot(
		declaration,
		3,
	)

	var optionalConstraints = declaration.GetOptionalConstraints()
	if uti.IsDefined(optionalConstraints) {
		v.processor_.PreprocessConstraints(
			optionalConstraints,
			0,
			0,
		)
		v.visitConstraints(optionalConstraints)
		v.processor_.PostprocessConstraints(
			optionalConstraints,
			0,
			0,
		)
	}
}

func (v *visitor_) visitDots(
	dots ast.DotsLike,
) {
	var delimiter = dots.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
}

func (v *visitor_) visitEnumeration(
	enumeration ast.EnumerationLike,
) {
	var delimiter1 = enumeration.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessEnumerationSlot(
		enumeration,
		1,
	)

	var delimiter2 = enumeration.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 2 between terms.
	v.processor_.ProcessEnumerationSlot(
		enumeration,
		2,
	)

	var value = enumeration.GetValue()
	v.processor_.PreprocessValue(
		value,
		0,
		0,
	)
	v.visitValue(value)
	v.processor_.PostprocessValue(
		value,
		0,
		0,
	)
	// Visit slot 3 between terms.
	v.processor_.ProcessEnumerationSlot(
		enumeration,
		3,
	)

	var additionalValuesIndex uint
	var additionalValues = enumeration.GetAdditionalValues().GetIterator()
	var additionalValuesCount = uint(additionalValues.GetSize())
	for additionalValues.HasNext() {
		additionalValuesIndex++
		var rule = additionalValues.GetNext()
		v.processor_.PreprocessAdditionalValue(
			rule,
			additionalValuesIndex,
			additionalValuesCount,
		)
		v.visitAdditionalValue(rule)
		v.processor_.PostprocessAdditionalValue(
			rule,
			additionalValuesIndex,
			additionalValuesCount,
		)
	}
	// Visit slot 4 between terms.
	v.processor_.ProcessEnumerationSlot(
		enumeration,
		4,
	)

	var delimiter3 = enumeration.GetDelimiter3()
	v.processor_.ProcessDelimiter(delimiter3)
}

func (v *visitor_) visitFunctionMethod(
	functionMethod ast.FunctionMethodLike,
) {
	var name = functionMethod.GetName()
	v.processor_.ProcessName(name)
	// Visit slot 1 between terms.
	v.processor_.ProcessFunctionMethodSlot(
		functionMethod,
		1,
	)

	var delimiter1 = functionMethod.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 2 between terms.
	v.processor_.ProcessFunctionMethodSlot(
		functionMethod,
		2,
	)

	var optionalParameterList = functionMethod.GetOptionalParameterList()
	if uti.IsDefined(optionalParameterList) {
		v.processor_.PreprocessParameterList(
			optionalParameterList,
			0,
			0,
		)
		v.visitParameterList(optionalParameterList)
		v.processor_.PostprocessParameterList(
			optionalParameterList,
			0,
			0,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessFunctionMethodSlot(
		functionMethod,
		3,
	)

	var delimiter2 = functionMethod.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 4 between terms.
	v.processor_.ProcessFunctionMethodSlot(
		functionMethod,
		4,
	)

	var result = functionMethod.GetResult()
	v.processor_.PreprocessResult(
		result,
		0,
		0,
	)
	v.visitResult(result)
	v.processor_.PostprocessResult(
		result,
		0,
		0,
	)
}

func (v *visitor_) visitFunctionSubsection(
	functionSubsection ast.FunctionSubsectionLike,
) {
	var delimiter = functionSubsection.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessFunctionSubsectionSlot(
		functionSubsection,
		1,
	)

	var functionMethodsIndex uint
	var functionMethods = functionSubsection.GetFunctionMethods().GetIterator()
	var functionMethodsCount = uint(functionMethods.GetSize())
	for functionMethods.HasNext() {
		functionMethodsIndex++
		var rule = functionMethods.GetNext()
		v.processor_.PreprocessFunctionMethod(
			rule,
			functionMethodsIndex,
			functionMethodsCount,
		)
		v.visitFunctionMethod(rule)
		v.processor_.PostprocessFunctionMethod(
			rule,
			functionMethodsIndex,
			functionMethodsCount,
		)
	}
}

func (v *visitor_) visitFunctional(
	functional ast.FunctionalLike,
) {
	var delimiter1 = functional.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessFunctionalSlot(
		functional,
		1,
	)

	var delimiter2 = functional.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 2 between terms.
	v.processor_.ProcessFunctionalSlot(
		functional,
		2,
	)

	var optionalParameterList = functional.GetOptionalParameterList()
	if uti.IsDefined(optionalParameterList) {
		v.processor_.PreprocessParameterList(
			optionalParameterList,
			0,
			0,
		)
		v.visitParameterList(optionalParameterList)
		v.processor_.PostprocessParameterList(
			optionalParameterList,
			0,
			0,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessFunctionalSlot(
		functional,
		3,
	)

	var delimiter3 = functional.GetDelimiter3()
	v.processor_.ProcessDelimiter(delimiter3)
	// Visit slot 4 between terms.
	v.processor_.ProcessFunctionalSlot(
		functional,
		4,
	)

	var optionalResult = functional.GetOptionalResult()
	if uti.IsDefined(optionalResult) {
		v.processor_.PreprocessResult(
			optionalResult,
			0,
			0,
		)
		v.visitResult(optionalResult)
		v.processor_.PostprocessResult(
			optionalResult,
			0,
			0,
		)
	}
}

func (v *visitor_) visitFunctionalDeclaration(
	functionalDeclaration ast.FunctionalDeclarationLike,
) {
	var declaration = functionalDeclaration.GetDeclaration()
	v.processor_.PreprocessDeclaration(
		declaration,
		0,
		0,
	)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(
		declaration,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessFunctionalDeclarationSlot(
		functionalDeclaration,
		1,
	)

	var functional = functionalDeclaration.GetFunctional()
	v.processor_.PreprocessFunctional(
		functional,
		0,
		0,
	)
	v.visitFunctional(functional)
	v.processor_.PostprocessFunctional(
		functional,
		0,
		0,
	)
}

func (v *visitor_) visitFunctionalSection(
	functionalSection ast.FunctionalSectionLike,
) {
	var delimiter = functionalSection.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessFunctionalSectionSlot(
		functionalSection,
		1,
	)

	var functionalDeclarationsIndex uint
	var functionalDeclarations = functionalSection.GetFunctionalDeclarations().GetIterator()
	var functionalDeclarationsCount = uint(functionalDeclarations.GetSize())
	for functionalDeclarations.HasNext() {
		functionalDeclarationsIndex++
		var rule = functionalDeclarations.GetNext()
		v.processor_.PreprocessFunctionalDeclaration(
			rule,
			functionalDeclarationsIndex,
			functionalDeclarationsCount,
		)
		v.visitFunctionalDeclaration(rule)
		v.processor_.PostprocessFunctionalDeclaration(
			rule,
			functionalDeclarationsIndex,
			functionalDeclarationsCount,
		)
	}
}

func (v *visitor_) visitGetterMethod(
	getterMethod ast.GetterMethodLike,
) {
	var name = getterMethod.GetName()
	v.processor_.ProcessName(name)
	// Visit slot 1 between terms.
	v.processor_.ProcessGetterMethodSlot(
		getterMethod,
		1,
	)

	var delimiter1 = getterMethod.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 2 between terms.
	v.processor_.ProcessGetterMethodSlot(
		getterMethod,
		2,
	)

	var delimiter2 = getterMethod.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessGetterMethodSlot(
		getterMethod,
		3,
	)

	var abstraction = getterMethod.GetAbstraction()
	v.processor_.PreprocessAbstraction(
		abstraction,
		0,
		0,
	)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(
		abstraction,
		0,
		0,
	)
}

func (v *visitor_) visitImportList(
	importList ast.ImportListLike,
) {
	var importedPackagesIndex uint
	var importedPackages = importList.GetImportedPackages().GetIterator()
	var importedPackagesCount = uint(importedPackages.GetSize())
	for importedPackages.HasNext() {
		importedPackagesIndex++
		var rule = importedPackages.GetNext()
		v.processor_.PreprocessImportedPackage(
			rule,
			importedPackagesIndex,
			importedPackagesCount,
		)
		v.visitImportedPackage(rule)
		v.processor_.PostprocessImportedPackage(
			rule,
			importedPackagesIndex,
			importedPackagesCount,
		)
	}
}

func (v *visitor_) visitImportedPackage(
	importedPackage ast.ImportedPackageLike,
) {
	var name = importedPackage.GetName()
	v.processor_.ProcessName(name)
	// Visit slot 1 between terms.
	v.processor_.ProcessImportedPackageSlot(
		importedPackage,
		1,
	)

	var path = importedPackage.GetPath()
	v.processor_.ProcessPath(path)
}

func (v *visitor_) visitInstanceDeclaration(
	instanceDeclaration ast.InstanceDeclarationLike,
) {
	var declaration = instanceDeclaration.GetDeclaration()
	v.processor_.PreprocessDeclaration(
		declaration,
		0,
		0,
	)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(
		declaration,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessInstanceDeclarationSlot(
		instanceDeclaration,
		1,
	)

	var delimiter1 = instanceDeclaration.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 2 between terms.
	v.processor_.ProcessInstanceDeclarationSlot(
		instanceDeclaration,
		2,
	)

	var delimiter2 = instanceDeclaration.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessInstanceDeclarationSlot(
		instanceDeclaration,
		3,
	)

	var instanceMethods = instanceDeclaration.GetInstanceMethods()
	v.processor_.PreprocessInstanceMethods(
		instanceMethods,
		0,
		0,
	)
	v.visitInstanceMethods(instanceMethods)
	v.processor_.PostprocessInstanceMethods(
		instanceMethods,
		0,
		0,
	)
	// Visit slot 4 between terms.
	v.processor_.ProcessInstanceDeclarationSlot(
		instanceDeclaration,
		4,
	)

	var delimiter3 = instanceDeclaration.GetDelimiter3()
	v.processor_.ProcessDelimiter(delimiter3)
}

func (v *visitor_) visitInstanceMethods(
	instanceMethods ast.InstanceMethodsLike,
) {
	var principalSubsection = instanceMethods.GetPrincipalSubsection()
	v.processor_.PreprocessPrincipalSubsection(
		principalSubsection,
		0,
		0,
	)
	v.visitPrincipalSubsection(principalSubsection)
	v.processor_.PostprocessPrincipalSubsection(
		principalSubsection,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessInstanceMethodsSlot(
		instanceMethods,
		1,
	)

	var optionalAttributeSubsection = instanceMethods.GetOptionalAttributeSubsection()
	if uti.IsDefined(optionalAttributeSubsection) {
		v.processor_.PreprocessAttributeSubsection(
			optionalAttributeSubsection,
			0,
			0,
		)
		v.visitAttributeSubsection(optionalAttributeSubsection)
		v.processor_.PostprocessAttributeSubsection(
			optionalAttributeSubsection,
			0,
			0,
		)
	}
	// Visit slot 2 between terms.
	v.processor_.ProcessInstanceMethodsSlot(
		instanceMethods,
		2,
	)

	var optionalAspectSubsection = instanceMethods.GetOptionalAspectSubsection()
	if uti.IsDefined(optionalAspectSubsection) {
		v.processor_.PreprocessAspectSubsection(
			optionalAspectSubsection,
			0,
			0,
		)
		v.visitAspectSubsection(optionalAspectSubsection)
		v.processor_.PostprocessAspectSubsection(
			optionalAspectSubsection,
			0,
			0,
		)
	}
}

func (v *visitor_) visitInstanceSection(
	instanceSection ast.InstanceSectionLike,
) {
	var delimiter = instanceSection.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessInstanceSectionSlot(
		instanceSection,
		1,
	)

	var instanceDeclarationsIndex uint
	var instanceDeclarations = instanceSection.GetInstanceDeclarations().GetIterator()
	var instanceDeclarationsCount = uint(instanceDeclarations.GetSize())
	for instanceDeclarations.HasNext() {
		instanceDeclarationsIndex++
		var rule = instanceDeclarations.GetNext()
		v.processor_.PreprocessInstanceDeclaration(
			rule,
			instanceDeclarationsIndex,
			instanceDeclarationsCount,
		)
		v.visitInstanceDeclaration(rule)
		v.processor_.PostprocessInstanceDeclaration(
			rule,
			instanceDeclarationsIndex,
			instanceDeclarationsCount,
		)
	}
}

func (v *visitor_) visitInterfaceDeclarations(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
) {
	var classSection = interfaceDeclarations.GetClassSection()
	v.processor_.PreprocessClassSection(
		classSection,
		0,
		0,
	)
	v.visitClassSection(classSection)
	v.processor_.PostprocessClassSection(
		classSection,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessInterfaceDeclarationsSlot(
		interfaceDeclarations,
		1,
	)

	var instanceSection = interfaceDeclarations.GetInstanceSection()
	v.processor_.PreprocessInstanceSection(
		instanceSection,
		0,
		0,
	)
	v.visitInstanceSection(instanceSection)
	v.processor_.PostprocessInstanceSection(
		instanceSection,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessInterfaceDeclarationsSlot(
		interfaceDeclarations,
		2,
	)

	var aspectSection = interfaceDeclarations.GetAspectSection()
	v.processor_.PreprocessAspectSection(
		aspectSection,
		0,
		0,
	)
	v.visitAspectSection(aspectSection)
	v.processor_.PostprocessAspectSection(
		aspectSection,
		0,
		0,
	)
}

func (v *visitor_) visitLegalNotice(
	legalNotice ast.LegalNoticeLike,
) {
	var comment = legalNotice.GetComment()
	v.processor_.ProcessComment(comment)
}

func (v *visitor_) visitMap(
	map_ ast.MapLike,
) {
	var delimiter1 = map_.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessMapSlot(
		map_,
		1,
	)

	var delimiter2 = map_.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 2 between terms.
	v.processor_.ProcessMapSlot(
		map_,
		2,
	)

	var name = map_.GetName()
	v.processor_.ProcessName(name)
	// Visit slot 3 between terms.
	v.processor_.ProcessMapSlot(
		map_,
		3,
	)

	var delimiter3 = map_.GetDelimiter3()
	v.processor_.ProcessDelimiter(delimiter3)
}

func (v *visitor_) visitMethod(
	method ast.MethodLike,
) {
	var name = method.GetName()
	v.processor_.ProcessName(name)
	// Visit slot 1 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		1,
	)

	var delimiter1 = method.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 2 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		2,
	)

	var optionalParameterList = method.GetOptionalParameterList()
	if uti.IsDefined(optionalParameterList) {
		v.processor_.PreprocessParameterList(
			optionalParameterList,
			0,
			0,
		)
		v.visitParameterList(optionalParameterList)
		v.processor_.PostprocessParameterList(
			optionalParameterList,
			0,
			0,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		3,
	)

	var delimiter2 = method.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 4 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		4,
	)

	var result = method.GetResult()
	v.processor_.PreprocessResult(
		result,
		0,
		0,
	)
	v.visitResult(result)
	v.processor_.PostprocessResult(
		result,
		0,
		0,
	)
}

func (v *visitor_) visitModel(
	model ast.ModelLike,
) {
	var packageDeclaration = model.GetPackageDeclaration()
	v.processor_.PreprocessPackageDeclaration(
		packageDeclaration,
		0,
		0,
	)
	v.visitPackageDeclaration(packageDeclaration)
	v.processor_.PostprocessPackageDeclaration(
		packageDeclaration,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessModelSlot(
		model,
		1,
	)

	var primitiveDeclarations = model.GetPrimitiveDeclarations()
	v.processor_.PreprocessPrimitiveDeclarations(
		primitiveDeclarations,
		0,
		0,
	)
	v.visitPrimitiveDeclarations(primitiveDeclarations)
	v.processor_.PostprocessPrimitiveDeclarations(
		primitiveDeclarations,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessModelSlot(
		model,
		2,
	)

	var interfaceDeclarations = model.GetInterfaceDeclarations()
	v.processor_.PreprocessInterfaceDeclarations(
		interfaceDeclarations,
		0,
		0,
	)
	v.visitInterfaceDeclarations(interfaceDeclarations)
	v.processor_.PostprocessInterfaceDeclarations(
		interfaceDeclarations,
		0,
		0,
	)
}

func (v *visitor_) visitMultivalue(
	multivalue ast.MultivalueLike,
) {
	var delimiter1 = multivalue.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessMultivalueSlot(
		multivalue,
		1,
	)

	var parameterList = multivalue.GetParameterList()
	v.processor_.PreprocessParameterList(
		parameterList,
		0,
		0,
	)
	v.visitParameterList(parameterList)
	v.processor_.PostprocessParameterList(
		parameterList,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessMultivalueSlot(
		multivalue,
		2,
	)

	var delimiter2 = multivalue.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitNamed(
	named ast.NamedLike,
) {
	var optionalPrefix = named.GetOptionalPrefix()
	if uti.IsDefined(optionalPrefix) {
		v.processor_.ProcessPrefix(optionalPrefix)
	}
	// Visit slot 1 between terms.
	v.processor_.ProcessNamedSlot(
		named,
		1,
	)

	var name = named.GetName()
	v.processor_.ProcessName(name)
	// Visit slot 2 between terms.
	v.processor_.ProcessNamedSlot(
		named,
		2,
	)

	var optionalArguments = named.GetOptionalArguments()
	if uti.IsDefined(optionalArguments) {
		v.processor_.PreprocessArguments(
			optionalArguments,
			0,
			0,
		)
		v.visitArguments(optionalArguments)
		v.processor_.PostprocessArguments(
			optionalArguments,
			0,
			0,
		)
	}
}

func (v *visitor_) visitNone(
	none ast.NoneLike,
) {
	var newline = none.GetNewline()
	v.processor_.ProcessNewline(newline)
}

func (v *visitor_) visitPackageDeclaration(
	packageDeclaration ast.PackageDeclarationLike,
) {
	var legalNotice = packageDeclaration.GetLegalNotice()
	v.processor_.PreprocessLegalNotice(
		legalNotice,
		0,
		0,
	)
	v.visitLegalNotice(legalNotice)
	v.processor_.PostprocessLegalNotice(
		legalNotice,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessPackageDeclarationSlot(
		packageDeclaration,
		1,
	)

	var packageHeader = packageDeclaration.GetPackageHeader()
	v.processor_.PreprocessPackageHeader(
		packageHeader,
		0,
		0,
	)
	v.visitPackageHeader(packageHeader)
	v.processor_.PostprocessPackageHeader(
		packageHeader,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessPackageDeclarationSlot(
		packageDeclaration,
		2,
	)

	var packageImports = packageDeclaration.GetPackageImports()
	v.processor_.PreprocessPackageImports(
		packageImports,
		0,
		0,
	)
	v.visitPackageImports(packageImports)
	v.processor_.PostprocessPackageImports(
		packageImports,
		0,
		0,
	)
}

func (v *visitor_) visitPackageHeader(
	packageHeader ast.PackageHeaderLike,
) {
	var comment = packageHeader.GetComment()
	v.processor_.ProcessComment(comment)
	// Visit slot 1 between terms.
	v.processor_.ProcessPackageHeaderSlot(
		packageHeader,
		1,
	)

	var delimiter = packageHeader.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 2 between terms.
	v.processor_.ProcessPackageHeaderSlot(
		packageHeader,
		2,
	)

	var name = packageHeader.GetName()
	v.processor_.ProcessName(name)
}

func (v *visitor_) visitPackageImports(
	packageImports ast.PackageImportsLike,
) {
	var delimiter1 = packageImports.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessPackageImportsSlot(
		packageImports,
		1,
	)

	var delimiter2 = packageImports.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 2 between terms.
	v.processor_.ProcessPackageImportsSlot(
		packageImports,
		2,
	)

	var optionalImportList = packageImports.GetOptionalImportList()
	if uti.IsDefined(optionalImportList) {
		v.processor_.PreprocessImportList(
			optionalImportList,
			0,
			0,
		)
		v.visitImportList(optionalImportList)
		v.processor_.PostprocessImportList(
			optionalImportList,
			0,
			0,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessPackageImportsSlot(
		packageImports,
		3,
	)

	var delimiter3 = packageImports.GetDelimiter3()
	v.processor_.ProcessDelimiter(delimiter3)
}

func (v *visitor_) visitParameter(
	parameter ast.ParameterLike,
) {
	var name = parameter.GetName()
	v.processor_.ProcessName(name)
	// Visit slot 1 between terms.
	v.processor_.ProcessParameterSlot(
		parameter,
		1,
	)

	var abstraction = parameter.GetAbstraction()
	v.processor_.PreprocessAbstraction(
		abstraction,
		0,
		0,
	)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(
		abstraction,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessParameterSlot(
		parameter,
		2,
	)

	var delimiter = parameter.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
}

func (v *visitor_) visitParameterList(
	parameterList ast.ParameterListLike,
) {
	var parametersIndex uint
	var parameters = parameterList.GetParameters().GetIterator()
	var parametersCount = uint(parameters.GetSize())
	for parameters.HasNext() {
		parametersIndex++
		var rule = parameters.GetNext()
		v.processor_.PreprocessParameter(
			rule,
			parametersIndex,
			parametersCount,
		)
		v.visitParameter(rule)
		v.processor_.PostprocessParameter(
			rule,
			parametersIndex,
			parametersCount,
		)
	}
}

func (v *visitor_) visitPrimitiveDeclarations(
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
) {
	var typeSection = primitiveDeclarations.GetTypeSection()
	v.processor_.PreprocessTypeSection(
		typeSection,
		0,
		0,
	)
	v.visitTypeSection(typeSection)
	v.processor_.PostprocessTypeSection(
		typeSection,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessPrimitiveDeclarationsSlot(
		primitiveDeclarations,
		1,
	)

	var functionalSection = primitiveDeclarations.GetFunctionalSection()
	v.processor_.PreprocessFunctionalSection(
		functionalSection,
		0,
		0,
	)
	v.visitFunctionalSection(functionalSection)
	v.processor_.PostprocessFunctionalSection(
		functionalSection,
		0,
		0,
	)
}

func (v *visitor_) visitPrincipalMethod(
	principalMethod ast.PrincipalMethodLike,
) {
	var method = principalMethod.GetMethod()
	v.processor_.PreprocessMethod(
		method,
		0,
		0,
	)
	v.visitMethod(method)
	v.processor_.PostprocessMethod(
		method,
		0,
		0,
	)
}

func (v *visitor_) visitPrincipalSubsection(
	principalSubsection ast.PrincipalSubsectionLike,
) {
	var delimiter = principalSubsection.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessPrincipalSubsectionSlot(
		principalSubsection,
		1,
	)

	var principalMethodsIndex uint
	var principalMethods = principalSubsection.GetPrincipalMethods().GetIterator()
	var principalMethodsCount = uint(principalMethods.GetSize())
	for principalMethods.HasNext() {
		principalMethodsIndex++
		var rule = principalMethods.GetNext()
		v.processor_.PreprocessPrincipalMethod(
			rule,
			principalMethodsIndex,
			principalMethodsCount,
		)
		v.visitPrincipalMethod(rule)
		v.processor_.PostprocessPrincipalMethod(
			rule,
			principalMethodsIndex,
			principalMethodsCount,
		)
	}
}

func (v *visitor_) visitResult(
	result ast.ResultLike,
) {
	// Visit the possible result rule types.
	switch actual := result.GetAny().(type) {
	case ast.NoneLike:
		v.processor_.PreprocessNone(
			actual,
			0,
			0,
		)
		v.visitNone(actual)
		v.processor_.PostprocessNone(
			actual,
			0,
			0,
		)
	case ast.AbstractionLike:
		v.processor_.PreprocessAbstraction(
			actual,
			0,
			0,
		)
		v.visitAbstraction(actual)
		v.processor_.PostprocessAbstraction(
			actual,
			0,
			0,
		)
	case ast.MultivalueLike:
		v.processor_.PreprocessMultivalue(
			actual,
			0,
			0,
		)
		v.visitMultivalue(actual)
		v.processor_.PostprocessMultivalue(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitSetterMethod(
	setterMethod ast.SetterMethodLike,
) {
	var name = setterMethod.GetName()
	v.processor_.ProcessName(name)
	// Visit slot 1 between terms.
	v.processor_.ProcessSetterMethodSlot(
		setterMethod,
		1,
	)

	var delimiter1 = setterMethod.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 2 between terms.
	v.processor_.ProcessSetterMethodSlot(
		setterMethod,
		2,
	)

	var parameter = setterMethod.GetParameter()
	v.processor_.PreprocessParameter(
		parameter,
		0,
		0,
	)
	v.visitParameter(parameter)
	v.processor_.PostprocessParameter(
		parameter,
		0,
		0,
	)
	// Visit slot 3 between terms.
	v.processor_.ProcessSetterMethodSlot(
		setterMethod,
		3,
	)

	var delimiter2 = setterMethod.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitStar(
	star ast.StarLike,
) {
	var delimiter = star.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
}

func (v *visitor_) visitType(
	type_ ast.TypeLike,
) {
	// Visit the possible type rule types.
	switch actual := type_.GetAny().(type) {
	case ast.NamedLike:
		v.processor_.PreprocessNamed(
			actual,
			0,
			0,
		)
		v.visitNamed(actual)
		v.processor_.PostprocessNamed(
			actual,
			0,
			0,
		)
	case ast.FunctionalLike:
		v.processor_.PreprocessFunctional(
			actual,
			0,
			0,
		)
		v.visitFunctional(actual)
		v.processor_.PostprocessFunctional(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitTypeDeclaration(
	typeDeclaration ast.TypeDeclarationLike,
) {
	var declaration = typeDeclaration.GetDeclaration()
	v.processor_.PreprocessDeclaration(
		declaration,
		0,
		0,
	)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(
		declaration,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessTypeDeclarationSlot(
		typeDeclaration,
		1,
	)

	var abstraction = typeDeclaration.GetAbstraction()
	v.processor_.PreprocessAbstraction(
		abstraction,
		0,
		0,
	)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(
		abstraction,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessTypeDeclarationSlot(
		typeDeclaration,
		2,
	)

	var optionalEnumeration = typeDeclaration.GetOptionalEnumeration()
	if uti.IsDefined(optionalEnumeration) {
		v.processor_.PreprocessEnumeration(
			optionalEnumeration,
			0,
			0,
		)
		v.visitEnumeration(optionalEnumeration)
		v.processor_.PostprocessEnumeration(
			optionalEnumeration,
			0,
			0,
		)
	}
}

func (v *visitor_) visitTypeSection(
	typeSection ast.TypeSectionLike,
) {
	var delimiter = typeSection.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessTypeSectionSlot(
		typeSection,
		1,
	)

	var typeDeclarationsIndex uint
	var typeDeclarations = typeSection.GetTypeDeclarations().GetIterator()
	var typeDeclarationsCount = uint(typeDeclarations.GetSize())
	for typeDeclarations.HasNext() {
		typeDeclarationsIndex++
		var rule = typeDeclarations.GetNext()
		v.processor_.PreprocessTypeDeclaration(
			rule,
			typeDeclarationsIndex,
			typeDeclarationsCount,
		)
		v.visitTypeDeclaration(rule)
		v.processor_.PostprocessTypeDeclaration(
			rule,
			typeDeclarationsIndex,
			typeDeclarationsCount,
		)
	}
}

func (v *visitor_) visitValue(
	value ast.ValueLike,
) {
	var name = value.GetName()
	v.processor_.ProcessName(name)
	// Visit slot 1 between terms.
	v.processor_.ProcessValueSlot(
		value,
		1,
	)

	var abstraction = value.GetAbstraction()
	v.processor_.PreprocessAbstraction(
		abstraction,
		0,
		0,
	)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(
		abstraction,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessValueSlot(
		value,
		2,
	)

	var delimiter1 = value.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 3 between terms.
	v.processor_.ProcessValueSlot(
		value,
		3,
	)

	var delimiter2 = value.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitWrapper(
	wrapper ast.WrapperLike,
) {
	// Visit the possible wrapper rule types.
	switch actual := wrapper.GetAny().(type) {
	case ast.DotsLike:
		v.processor_.PreprocessDots(
			actual,
			0,
			0,
		)
		v.visitDots(actual)
		v.processor_.PostprocessDots(
			actual,
			0,
			0,
		)
	case ast.StarLike:
		v.processor_.PreprocessStar(
			actual,
			0,
			0,
		)
		v.visitStar(actual)
		v.processor_.PostprocessStar(
			actual,
			0,
			0,
		)
	case ast.ArrayLike:
		v.processor_.PreprocessArray(
			actual,
			0,
			0,
		)
		v.visitArray(actual)
		v.processor_.PostprocessArray(
			actual,
			0,
			0,
		)
	case ast.ChannelLike:
		v.processor_.PreprocessChannel(
			actual,
			0,
			0,
		)
		v.visitChannel(actual)
		v.processor_.PostprocessChannel(
			actual,
			0,
			0,
		)
	case ast.MapLike:
		v.processor_.PreprocessMap(
			actual,
			0,
			0,
		)
		v.visitMap(actual)
		v.processor_.PostprocessMap(
			actual,
			0,
			0,
		)
	}
}

// Instance Structure

type visitor_ struct {
	// Declare the instance attributes.
	processor_ Methodical
}

// Class Structure

type visitorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func visitorClass() *visitorClass_ {
	return visitorClassReference_
}

var visitorClassReference_ = &visitorClass_{
	// Initialize the class constants.
}
