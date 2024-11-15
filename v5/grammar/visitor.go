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

package grammar

import (
	fmt "fmt"
	ast "github.com/craterdog/go-class-model/v5/ast"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Visitor() VisitorClassLike {
	return visitorReference()
}

// Constructor Methods

func (c *visitorClass_) Make(
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

// Primary Methods

func (v *visitor_) GetClass() VisitorClassLike {
	return visitorReference()
}

func (v *visitor_) VisitModel(
	model ast.ModelLike,
) {
	v.processor_.PreprocessModel(model)
	v.visitModel(model)
	v.processor_.PostprocessModel(model)
}

// PROTECTED INTERFACE

// Private Methods

func (v *visitor_) visitAbstraction(abstraction ast.AbstractionLike) {
	// Visit an optional prefix rule.
	var optionalPrefix = abstraction.GetOptionalPrefix()
	if uti.IsDefined(optionalPrefix) {
		v.processor_.PreprocessPrefix(optionalPrefix)
		v.visitPrefix(optionalPrefix)
		v.processor_.PostprocessPrefix(optionalPrefix)
	}

	// Visit slot 1 between references.
	v.processor_.ProcessAbstractionSlot(1)

	// Visit a single name token.
	var name = abstraction.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 2 between references.
	v.processor_.ProcessAbstractionSlot(2)

	// Visit an optional suffix rule.
	var optionalSuffix = abstraction.GetOptionalSuffix()
	if uti.IsDefined(optionalSuffix) {
		v.processor_.PreprocessSuffix(optionalSuffix)
		v.visitSuffix(optionalSuffix)
		v.processor_.PostprocessSuffix(optionalSuffix)
	}

	// Visit slot 3 between references.
	v.processor_.ProcessAbstractionSlot(3)

	// Visit an optional arguments rule.
	var optionalArguments = abstraction.GetOptionalArguments()
	if uti.IsDefined(optionalArguments) {
		v.processor_.PreprocessArguments(optionalArguments)
		v.visitArguments(optionalArguments)
		v.processor_.PostprocessArguments(optionalArguments)
	}
}

func (v *visitor_) visitAdditionalArgument(additionalArgument ast.AdditionalArgumentLike) {
	// Visit a single argument rule.
	var argument = additionalArgument.GetArgument()
	v.processor_.PreprocessArgument(argument)
	v.visitArgument(argument)
	v.processor_.PostprocessArgument(argument)
}

func (v *visitor_) visitAdditionalConstraint(additionalConstraint ast.AdditionalConstraintLike) {
	// Visit a single constraint rule.
	var constraint = additionalConstraint.GetConstraint()
	v.processor_.PreprocessConstraint(constraint)
	v.visitConstraint(constraint)
	v.processor_.PostprocessConstraint(constraint)
}

func (v *visitor_) visitAdditionalValue(additionalValue ast.AdditionalValueLike) {
	// Visit a single name token.
	var name = additionalValue.GetName()
	v.processor_.ProcessName(name)
}

func (v *visitor_) visitArgument(argument ast.ArgumentLike) {
	// Visit a single abstraction rule.
	var abstraction = argument.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}

func (v *visitor_) visitArguments(arguments ast.ArgumentsLike) {
	// Visit a single argument rule.
	var argument = arguments.GetArgument()
	v.processor_.PreprocessArgument(argument)
	v.visitArgument(argument)
	v.processor_.PostprocessArgument(argument)

	// Visit slot 1 between references.
	v.processor_.ProcessArgumentsSlot(1)

	// Visit each additionalArgument rule.
	var additionalArgumentIndex uint
	var additionalArguments = arguments.GetAdditionalArguments().GetIterator()
	var additionalArgumentsSize = uint(additionalArguments.GetSize())
	for additionalArguments.HasNext() {
		additionalArgumentIndex++
		var additionalArgument = additionalArguments.GetNext()
		v.processor_.PreprocessAdditionalArgument(
			additionalArgument,
			additionalArgumentIndex,
			additionalArgumentsSize,
		)
		v.visitAdditionalArgument(additionalArgument)
		v.processor_.PostprocessAdditionalArgument(
			additionalArgument,
			additionalArgumentIndex,
			additionalArgumentsSize,
		)
	}
}

func (v *visitor_) visitArray(array ast.ArrayLike) {}

func (v *visitor_) visitAspectDeclaration(aspectDeclaration ast.AspectDeclarationLike) {
	// Visit a single declaration rule.
	var declaration = aspectDeclaration.GetDeclaration()
	v.processor_.PreprocessDeclaration(declaration)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(declaration)

	// Visit slot 1 between references.
	v.processor_.ProcessAspectDeclarationSlot(1)

	// Visit each aspectMethod rule.
	var aspectMethodIndex uint
	var aspectMethods = aspectDeclaration.GetAspectMethods().GetIterator()
	var aspectMethodsSize = uint(aspectMethods.GetSize())
	for aspectMethods.HasNext() {
		aspectMethodIndex++
		var aspectMethod = aspectMethods.GetNext()
		v.processor_.PreprocessAspectMethod(
			aspectMethod,
			aspectMethodIndex,
			aspectMethodsSize,
		)
		v.visitAspectMethod(aspectMethod)
		v.processor_.PostprocessAspectMethod(
			aspectMethod,
			aspectMethodIndex,
			aspectMethodsSize,
		)
	}
}

func (v *visitor_) visitAspectInterface(aspectInterface ast.AspectInterfaceLike) {
	// Visit a single abstraction rule.
	var abstraction = aspectInterface.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}

func (v *visitor_) visitAspectMethod(aspectMethod ast.AspectMethodLike) {
	// Visit a single method rule.
	var method = aspectMethod.GetMethod()
	v.processor_.PreprocessMethod(method)
	v.visitMethod(method)
	v.processor_.PostprocessMethod(method)
}

func (v *visitor_) visitAspectSection(aspectSection ast.AspectSectionLike) {
	// Visit each aspectDeclaration rule.
	var aspectDeclarationIndex uint
	var aspectDeclarations = aspectSection.GetAspectDeclarations().GetIterator()
	var aspectDeclarationsSize = uint(aspectDeclarations.GetSize())
	for aspectDeclarations.HasNext() {
		aspectDeclarationIndex++
		var aspectDeclaration = aspectDeclarations.GetNext()
		v.processor_.PreprocessAspectDeclaration(
			aspectDeclaration,
			aspectDeclarationIndex,
			aspectDeclarationsSize,
		)
		v.visitAspectDeclaration(aspectDeclaration)
		v.processor_.PostprocessAspectDeclaration(
			aspectDeclaration,
			aspectDeclarationIndex,
			aspectDeclarationsSize,
		)
	}
}

func (v *visitor_) visitAspectSubsection(aspectSubsection ast.AspectSubsectionLike) {
	// Visit each aspectInterface rule.
	var aspectInterfaceIndex uint
	var aspectInterfaces = aspectSubsection.GetAspectInterfaces().GetIterator()
	var aspectInterfacesSize = uint(aspectInterfaces.GetSize())
	for aspectInterfaces.HasNext() {
		aspectInterfaceIndex++
		var aspectInterface = aspectInterfaces.GetNext()
		v.processor_.PreprocessAspectInterface(
			aspectInterface,
			aspectInterfaceIndex,
			aspectInterfacesSize,
		)
		v.visitAspectInterface(aspectInterface)
		v.processor_.PostprocessAspectInterface(
			aspectInterface,
			aspectInterfaceIndex,
			aspectInterfacesSize,
		)
	}
}

func (v *visitor_) visitAttributeMethod(attributeMethod ast.AttributeMethodLike) {
	// Visit the possible attributeMethod types.
	switch actual := attributeMethod.GetAny().(type) {
	case ast.GetterMethodLike:
		v.processor_.PreprocessGetterMethod(actual)
		v.visitGetterMethod(actual)
		v.processor_.PostprocessGetterMethod(actual)
	case ast.SetterMethodLike:
		v.processor_.PreprocessSetterMethod(actual)
		v.visitSetterMethod(actual)
		v.processor_.PostprocessSetterMethod(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitAttributeSubsection(attributeSubsection ast.AttributeSubsectionLike) {
	// Visit each attributeMethod rule.
	var attributeMethodIndex uint
	var attributeMethods = attributeSubsection.GetAttributeMethods().GetIterator()
	var attributeMethodsSize = uint(attributeMethods.GetSize())
	for attributeMethods.HasNext() {
		attributeMethodIndex++
		var attributeMethod = attributeMethods.GetNext()
		v.processor_.PreprocessAttributeMethod(
			attributeMethod,
			attributeMethodIndex,
			attributeMethodsSize,
		)
		v.visitAttributeMethod(attributeMethod)
		v.processor_.PostprocessAttributeMethod(
			attributeMethod,
			attributeMethodIndex,
			attributeMethodsSize,
		)
	}
}

func (v *visitor_) visitChannel(channel ast.ChannelLike) {}

func (v *visitor_) visitClassDeclaration(classDeclaration ast.ClassDeclarationLike) {
	// Visit a single declaration rule.
	var declaration = classDeclaration.GetDeclaration()
	v.processor_.PreprocessDeclaration(declaration)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(declaration)

	// Visit slot 1 between references.
	v.processor_.ProcessClassDeclarationSlot(1)

	// Visit a single classMethods rule.
	var classMethods = classDeclaration.GetClassMethods()
	v.processor_.PreprocessClassMethods(classMethods)
	v.visitClassMethods(classMethods)
	v.processor_.PostprocessClassMethods(classMethods)
}

func (v *visitor_) visitClassMethods(classMethods ast.ClassMethodsLike) {
	// Visit a single constructorSubsection rule.
	var constructorSubsection = classMethods.GetConstructorSubsection()
	v.processor_.PreprocessConstructorSubsection(constructorSubsection)
	v.visitConstructorSubsection(constructorSubsection)
	v.processor_.PostprocessConstructorSubsection(constructorSubsection)

	// Visit slot 1 between references.
	v.processor_.ProcessClassMethodsSlot(1)

	// Visit an optional constantSubsection rule.
	var optionalConstantSubsection = classMethods.GetOptionalConstantSubsection()
	if uti.IsDefined(optionalConstantSubsection) {
		v.processor_.PreprocessConstantSubsection(optionalConstantSubsection)
		v.visitConstantSubsection(optionalConstantSubsection)
		v.processor_.PostprocessConstantSubsection(optionalConstantSubsection)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessClassMethodsSlot(2)

	// Visit an optional functionSubsection rule.
	var optionalFunctionSubsection = classMethods.GetOptionalFunctionSubsection()
	if uti.IsDefined(optionalFunctionSubsection) {
		v.processor_.PreprocessFunctionSubsection(optionalFunctionSubsection)
		v.visitFunctionSubsection(optionalFunctionSubsection)
		v.processor_.PostprocessFunctionSubsection(optionalFunctionSubsection)
	}
}

func (v *visitor_) visitClassSection(classSection ast.ClassSectionLike) {
	// Visit each classDeclaration rule.
	var classDeclarationIndex uint
	var classDeclarations = classSection.GetClassDeclarations().GetIterator()
	var classDeclarationsSize = uint(classDeclarations.GetSize())
	for classDeclarations.HasNext() {
		classDeclarationIndex++
		var classDeclaration = classDeclarations.GetNext()
		v.processor_.PreprocessClassDeclaration(
			classDeclaration,
			classDeclarationIndex,
			classDeclarationsSize,
		)
		v.visitClassDeclaration(classDeclaration)
		v.processor_.PostprocessClassDeclaration(
			classDeclaration,
			classDeclarationIndex,
			classDeclarationsSize,
		)
	}
}

func (v *visitor_) visitConstantMethod(constantMethod ast.ConstantMethodLike) {
	// Visit a single name token.
	var name = constantMethod.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessConstantMethodSlot(1)

	// Visit a single abstraction rule.
	var abstraction = constantMethod.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}

func (v *visitor_) visitConstantSubsection(constantSubsection ast.ConstantSubsectionLike) {
	// Visit each constantMethod rule.
	var constantMethodIndex uint
	var constantMethods = constantSubsection.GetConstantMethods().GetIterator()
	var constantMethodsSize = uint(constantMethods.GetSize())
	for constantMethods.HasNext() {
		constantMethodIndex++
		var constantMethod = constantMethods.GetNext()
		v.processor_.PreprocessConstantMethod(
			constantMethod,
			constantMethodIndex,
			constantMethodsSize,
		)
		v.visitConstantMethod(constantMethod)
		v.processor_.PostprocessConstantMethod(
			constantMethod,
			constantMethodIndex,
			constantMethodsSize,
		)
	}
}

func (v *visitor_) visitConstraint(constraint ast.ConstraintLike) {
	// Visit a single name token.
	var name = constraint.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessConstraintSlot(1)

	// Visit a single abstraction rule.
	var abstraction = constraint.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}

func (v *visitor_) visitConstraints(constraints ast.ConstraintsLike) {
	// Visit a single constraint rule.
	var constraint = constraints.GetConstraint()
	v.processor_.PreprocessConstraint(constraint)
	v.visitConstraint(constraint)
	v.processor_.PostprocessConstraint(constraint)

	// Visit slot 1 between references.
	v.processor_.ProcessConstraintsSlot(1)

	// Visit each additionalConstraint rule.
	var additionalConstraintIndex uint
	var additionalConstraints = constraints.GetAdditionalConstraints().GetIterator()
	var additionalConstraintsSize = uint(additionalConstraints.GetSize())
	for additionalConstraints.HasNext() {
		additionalConstraintIndex++
		var additionalConstraint = additionalConstraints.GetNext()
		v.processor_.PreprocessAdditionalConstraint(
			additionalConstraint,
			additionalConstraintIndex,
			additionalConstraintsSize,
		)
		v.visitAdditionalConstraint(additionalConstraint)
		v.processor_.PostprocessAdditionalConstraint(
			additionalConstraint,
			additionalConstraintIndex,
			additionalConstraintsSize,
		)
	}
}

func (v *visitor_) visitConstructorMethod(constructorMethod ast.ConstructorMethodLike) {
	// Visit a single name token.
	var name = constructorMethod.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessConstructorMethodSlot(1)

	// Visit each parameter rule.
	var parameterIndex uint
	var parameters = constructorMethod.GetParameters().GetIterator()
	var parametersSize = uint(parameters.GetSize())
	for parameters.HasNext() {
		parameterIndex++
		var parameter = parameters.GetNext()
		v.processor_.PreprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
		v.visitParameter(parameter)
		v.processor_.PostprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessConstructorMethodSlot(2)

	// Visit a single abstraction rule.
	var abstraction = constructorMethod.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}

func (v *visitor_) visitConstructorSubsection(constructorSubsection ast.ConstructorSubsectionLike) {
	// Visit each constructorMethod rule.
	var constructorMethodIndex uint
	var constructorMethods = constructorSubsection.GetConstructorMethods().GetIterator()
	var constructorMethodsSize = uint(constructorMethods.GetSize())
	for constructorMethods.HasNext() {
		constructorMethodIndex++
		var constructorMethod = constructorMethods.GetNext()
		v.processor_.PreprocessConstructorMethod(
			constructorMethod,
			constructorMethodIndex,
			constructorMethodsSize,
		)
		v.visitConstructorMethod(constructorMethod)
		v.processor_.PostprocessConstructorMethod(
			constructorMethod,
			constructorMethodIndex,
			constructorMethodsSize,
		)
	}
}

func (v *visitor_) visitDeclaration(declaration ast.DeclarationLike) {
	// Visit a single comment token.
	var comment = declaration.GetComment()
	v.processor_.ProcessComment(comment)

	// Visit slot 1 between references.
	v.processor_.ProcessDeclarationSlot(1)

	// Visit a single name token.
	var name = declaration.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 2 between references.
	v.processor_.ProcessDeclarationSlot(2)

	// Visit an optional constraints rule.
	var optionalConstraints = declaration.GetOptionalConstraints()
	if uti.IsDefined(optionalConstraints) {
		v.processor_.PreprocessConstraints(optionalConstraints)
		v.visitConstraints(optionalConstraints)
		v.processor_.PostprocessConstraints(optionalConstraints)
	}
}

func (v *visitor_) visitEnumeration(enumeration ast.EnumerationLike) {
	// Visit a single value rule.
	var value = enumeration.GetValue()
	v.processor_.PreprocessValue(value)
	v.visitValue(value)
	v.processor_.PostprocessValue(value)

	// Visit slot 1 between references.
	v.processor_.ProcessEnumerationSlot(1)

	// Visit each additionalValue rule.
	var additionalValueIndex uint
	var additionalValues = enumeration.GetAdditionalValues().GetIterator()
	var additionalValuesSize = uint(additionalValues.GetSize())
	for additionalValues.HasNext() {
		additionalValueIndex++
		var additionalValue = additionalValues.GetNext()
		v.processor_.PreprocessAdditionalValue(
			additionalValue,
			additionalValueIndex,
			additionalValuesSize,
		)
		v.visitAdditionalValue(additionalValue)
		v.processor_.PostprocessAdditionalValue(
			additionalValue,
			additionalValueIndex,
			additionalValuesSize,
		)
	}
}

func (v *visitor_) visitFunctionMethod(functionMethod ast.FunctionMethodLike) {
	// Visit a single name token.
	var name = functionMethod.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessFunctionMethodSlot(1)

	// Visit each parameter rule.
	var parameterIndex uint
	var parameters = functionMethod.GetParameters().GetIterator()
	var parametersSize = uint(parameters.GetSize())
	for parameters.HasNext() {
		parameterIndex++
		var parameter = parameters.GetNext()
		v.processor_.PreprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
		v.visitParameter(parameter)
		v.processor_.PostprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessFunctionMethodSlot(2)

	// Visit a single result rule.
	var result = functionMethod.GetResult()
	v.processor_.PreprocessResult(result)
	v.visitResult(result)
	v.processor_.PostprocessResult(result)
}

func (v *visitor_) visitFunctionSubsection(functionSubsection ast.FunctionSubsectionLike) {
	// Visit each functionMethod rule.
	var functionMethodIndex uint
	var functionMethods = functionSubsection.GetFunctionMethods().GetIterator()
	var functionMethodsSize = uint(functionMethods.GetSize())
	for functionMethods.HasNext() {
		functionMethodIndex++
		var functionMethod = functionMethods.GetNext()
		v.processor_.PreprocessFunctionMethod(
			functionMethod,
			functionMethodIndex,
			functionMethodsSize,
		)
		v.visitFunctionMethod(functionMethod)
		v.processor_.PostprocessFunctionMethod(
			functionMethod,
			functionMethodIndex,
			functionMethodsSize,
		)
	}
}

func (v *visitor_) visitFunctionalDeclaration(functionalDeclaration ast.FunctionalDeclarationLike) {
	// Visit a single declaration rule.
	var declaration = functionalDeclaration.GetDeclaration()
	v.processor_.PreprocessDeclaration(declaration)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(declaration)

	// Visit slot 1 between references.
	v.processor_.ProcessFunctionalDeclarationSlot(1)

	// Visit each parameter rule.
	var parameterIndex uint
	var parameters = functionalDeclaration.GetParameters().GetIterator()
	var parametersSize = uint(parameters.GetSize())
	for parameters.HasNext() {
		parameterIndex++
		var parameter = parameters.GetNext()
		v.processor_.PreprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
		v.visitParameter(parameter)
		v.processor_.PostprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessFunctionalDeclarationSlot(2)

	// Visit a single result rule.
	var result = functionalDeclaration.GetResult()
	v.processor_.PreprocessResult(result)
	v.visitResult(result)
	v.processor_.PostprocessResult(result)
}

func (v *visitor_) visitFunctionalSection(functionalSection ast.FunctionalSectionLike) {
	// Visit each functionalDeclaration rule.
	var functionalDeclarationIndex uint
	var functionalDeclarations = functionalSection.GetFunctionalDeclarations().GetIterator()
	var functionalDeclarationsSize = uint(functionalDeclarations.GetSize())
	for functionalDeclarations.HasNext() {
		functionalDeclarationIndex++
		var functionalDeclaration = functionalDeclarations.GetNext()
		v.processor_.PreprocessFunctionalDeclaration(
			functionalDeclaration,
			functionalDeclarationIndex,
			functionalDeclarationsSize,
		)
		v.visitFunctionalDeclaration(functionalDeclaration)
		v.processor_.PostprocessFunctionalDeclaration(
			functionalDeclaration,
			functionalDeclarationIndex,
			functionalDeclarationsSize,
		)
	}
}

func (v *visitor_) visitGetterMethod(getterMethod ast.GetterMethodLike) {
	// Visit a single name token.
	var name = getterMethod.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessGetterMethodSlot(1)

	// Visit a single abstraction rule.
	var abstraction = getterMethod.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}

func (v *visitor_) visitImportedPackage(importedPackage ast.ImportedPackageLike) {
	// Visit a single name token.
	var name = importedPackage.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessImportedPackageSlot(1)

	// Visit a single path token.
	var path = importedPackage.GetPath()
	v.processor_.ProcessPath(path)
}

func (v *visitor_) visitInstanceDeclaration(instanceDeclaration ast.InstanceDeclarationLike) {
	// Visit a single declaration rule.
	var declaration = instanceDeclaration.GetDeclaration()
	v.processor_.PreprocessDeclaration(declaration)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(declaration)

	// Visit slot 1 between references.
	v.processor_.ProcessInstanceDeclarationSlot(1)

	// Visit a single instanceMethods rule.
	var instanceMethods = instanceDeclaration.GetInstanceMethods()
	v.processor_.PreprocessInstanceMethods(instanceMethods)
	v.visitInstanceMethods(instanceMethods)
	v.processor_.PostprocessInstanceMethods(instanceMethods)
}

func (v *visitor_) visitInstanceMethods(instanceMethods ast.InstanceMethodsLike) {
	// Visit a single primarySubsection rule.
	var primarySubsection = instanceMethods.GetPrimarySubsection()
	v.processor_.PreprocessPrimarySubsection(primarySubsection)
	v.visitPrimarySubsection(primarySubsection)
	v.processor_.PostprocessPrimarySubsection(primarySubsection)

	// Visit slot 1 between references.
	v.processor_.ProcessInstanceMethodsSlot(1)

	// Visit an optional attributeSubsection rule.
	var optionalAttributeSubsection = instanceMethods.GetOptionalAttributeSubsection()
	if uti.IsDefined(optionalAttributeSubsection) {
		v.processor_.PreprocessAttributeSubsection(optionalAttributeSubsection)
		v.visitAttributeSubsection(optionalAttributeSubsection)
		v.processor_.PostprocessAttributeSubsection(optionalAttributeSubsection)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessInstanceMethodsSlot(2)

	// Visit an optional aspectSubsection rule.
	var optionalAspectSubsection = instanceMethods.GetOptionalAspectSubsection()
	if uti.IsDefined(optionalAspectSubsection) {
		v.processor_.PreprocessAspectSubsection(optionalAspectSubsection)
		v.visitAspectSubsection(optionalAspectSubsection)
		v.processor_.PostprocessAspectSubsection(optionalAspectSubsection)
	}
}

func (v *visitor_) visitInstanceSection(instanceSection ast.InstanceSectionLike) {
	// Visit each instanceDeclaration rule.
	var instanceDeclarationIndex uint
	var instanceDeclarations = instanceSection.GetInstanceDeclarations().GetIterator()
	var instanceDeclarationsSize = uint(instanceDeclarations.GetSize())
	for instanceDeclarations.HasNext() {
		instanceDeclarationIndex++
		var instanceDeclaration = instanceDeclarations.GetNext()
		v.processor_.PreprocessInstanceDeclaration(
			instanceDeclaration,
			instanceDeclarationIndex,
			instanceDeclarationsSize,
		)
		v.visitInstanceDeclaration(instanceDeclaration)
		v.processor_.PostprocessInstanceDeclaration(
			instanceDeclaration,
			instanceDeclarationIndex,
			instanceDeclarationsSize,
		)
	}
}

func (v *visitor_) visitInterfaceDeclarations(interfaceDeclarations ast.InterfaceDeclarationsLike) {
	// Visit a single classSection rule.
	var classSection = interfaceDeclarations.GetClassSection()
	v.processor_.PreprocessClassSection(classSection)
	v.visitClassSection(classSection)
	v.processor_.PostprocessClassSection(classSection)

	// Visit slot 1 between references.
	v.processor_.ProcessInterfaceDeclarationsSlot(1)

	// Visit a single instanceSection rule.
	var instanceSection = interfaceDeclarations.GetInstanceSection()
	v.processor_.PreprocessInstanceSection(instanceSection)
	v.visitInstanceSection(instanceSection)
	v.processor_.PostprocessInstanceSection(instanceSection)

	// Visit slot 2 between references.
	v.processor_.ProcessInterfaceDeclarationsSlot(2)

	// Visit a single aspectSection rule.
	var aspectSection = interfaceDeclarations.GetAspectSection()
	v.processor_.PreprocessAspectSection(aspectSection)
	v.visitAspectSection(aspectSection)
	v.processor_.PostprocessAspectSection(aspectSection)
}

func (v *visitor_) visitLegalNotice(legalNotice ast.LegalNoticeLike) {
	// Visit a single comment token.
	var comment = legalNotice.GetComment()
	v.processor_.ProcessComment(comment)
}

func (v *visitor_) visitMap(map_ ast.MapLike) {
	// Visit a single name token.
	var name = map_.GetName()
	v.processor_.ProcessName(name)
}

func (v *visitor_) visitMethod(method ast.MethodLike) {
	// Visit a single name token.
	var name = method.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessMethodSlot(1)

	// Visit each parameter rule.
	var parameterIndex uint
	var parameters = method.GetParameters().GetIterator()
	var parametersSize = uint(parameters.GetSize())
	for parameters.HasNext() {
		parameterIndex++
		var parameter = parameters.GetNext()
		v.processor_.PreprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
		v.visitParameter(parameter)
		v.processor_.PostprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessMethodSlot(2)

	// Visit an optional result rule.
	var optionalResult = method.GetOptionalResult()
	if uti.IsDefined(optionalResult) {
		v.processor_.PreprocessResult(optionalResult)
		v.visitResult(optionalResult)
		v.processor_.PostprocessResult(optionalResult)
	}
}

func (v *visitor_) visitModel(model ast.ModelLike) {
	// Visit a single moduleDeclaration rule.
	var moduleDeclaration = model.GetModuleDeclaration()
	v.processor_.PreprocessModuleDeclaration(moduleDeclaration)
	v.visitModuleDeclaration(moduleDeclaration)
	v.processor_.PostprocessModuleDeclaration(moduleDeclaration)

	// Visit slot 1 between references.
	v.processor_.ProcessModelSlot(1)

	// Visit a single primitiveDeclarations rule.
	var primitiveDeclarations = model.GetPrimitiveDeclarations()
	v.processor_.PreprocessPrimitiveDeclarations(primitiveDeclarations)
	v.visitPrimitiveDeclarations(primitiveDeclarations)
	v.processor_.PostprocessPrimitiveDeclarations(primitiveDeclarations)

	// Visit slot 2 between references.
	v.processor_.ProcessModelSlot(2)

	// Visit a single interfaceDeclarations rule.
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	v.processor_.PreprocessInterfaceDeclarations(interfaceDeclarations)
	v.visitInterfaceDeclarations(interfaceDeclarations)
	v.processor_.PostprocessInterfaceDeclarations(interfaceDeclarations)
}

func (v *visitor_) visitModuleDeclaration(moduleDeclaration ast.ModuleDeclarationLike) {
	// Visit a single legalNotice rule.
	var legalNotice = moduleDeclaration.GetLegalNotice()
	v.processor_.PreprocessLegalNotice(legalNotice)
	v.visitLegalNotice(legalNotice)
	v.processor_.PostprocessLegalNotice(legalNotice)

	// Visit slot 1 between references.
	v.processor_.ProcessModuleDeclarationSlot(1)

	// Visit a single moduleHeader rule.
	var moduleHeader = moduleDeclaration.GetModuleHeader()
	v.processor_.PreprocessModuleHeader(moduleHeader)
	v.visitModuleHeader(moduleHeader)
	v.processor_.PostprocessModuleHeader(moduleHeader)

	// Visit slot 2 between references.
	v.processor_.ProcessModuleDeclarationSlot(2)

	// Visit a single moduleImports rule.
	var moduleImports = moduleDeclaration.GetModuleImports()
	v.processor_.PreprocessModuleImports(moduleImports)
	v.visitModuleImports(moduleImports)
	v.processor_.PostprocessModuleImports(moduleImports)
}

func (v *visitor_) visitModuleHeader(moduleHeader ast.ModuleHeaderLike) {
	// Visit a single comment token.
	var comment = moduleHeader.GetComment()
	v.processor_.ProcessComment(comment)

	// Visit slot 1 between references.
	v.processor_.ProcessModuleHeaderSlot(1)

	// Visit a single name token.
	var name = moduleHeader.GetName()
	v.processor_.ProcessName(name)
}

func (v *visitor_) visitModuleImports(moduleImports ast.ModuleImportsLike) {
	// Visit each importedPackage rule.
	var importedPackageIndex uint
	var importedPackages = moduleImports.GetImportedPackages().GetIterator()
	var importedPackagesSize = uint(importedPackages.GetSize())
	for importedPackages.HasNext() {
		importedPackageIndex++
		var importedPackage = importedPackages.GetNext()
		v.processor_.PreprocessImportedPackage(
			importedPackage,
			importedPackageIndex,
			importedPackagesSize,
		)
		v.visitImportedPackage(importedPackage)
		v.processor_.PostprocessImportedPackage(
			importedPackage,
			importedPackageIndex,
			importedPackagesSize,
		)
	}
}

func (v *visitor_) visitMultivalue(multivalue ast.MultivalueLike) {
	// Visit each parameter rule.
	var parameterIndex uint
	var parameters = multivalue.GetParameters().GetIterator()
	var parametersSize = uint(parameters.GetSize())
	for parameters.HasNext() {
		parameterIndex++
		var parameter = parameters.GetNext()
		v.processor_.PreprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
		v.visitParameter(parameter)
		v.processor_.PostprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
	}
}

func (v *visitor_) visitNone(none ast.NoneLike) {
	// Visit a single newline token.
	var newline = none.GetNewline()
	v.processor_.ProcessNewline(newline)
}

func (v *visitor_) visitParameter(parameter ast.ParameterLike) {
	// Visit a single name token.
	var name = parameter.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessParameterSlot(1)

	// Visit a single abstraction rule.
	var abstraction = parameter.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}

func (v *visitor_) visitPrefix(prefix ast.PrefixLike) {
	// Visit the possible prefix types.
	switch actual := prefix.GetAny().(type) {
	case ast.ArrayLike:
		v.processor_.PreprocessArray(actual)
		v.visitArray(actual)
		v.processor_.PostprocessArray(actual)
	case ast.MapLike:
		v.processor_.PreprocessMap(actual)
		v.visitMap(actual)
		v.processor_.PostprocessMap(actual)
	case ast.ChannelLike:
		v.processor_.PreprocessChannel(actual)
		v.visitChannel(actual)
		v.processor_.PostprocessChannel(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitPrimaryMethod(primaryMethod ast.PrimaryMethodLike) {
	// Visit a single method rule.
	var method = primaryMethod.GetMethod()
	v.processor_.PreprocessMethod(method)
	v.visitMethod(method)
	v.processor_.PostprocessMethod(method)
}

func (v *visitor_) visitPrimarySubsection(primarySubsection ast.PrimarySubsectionLike) {
	// Visit each primaryMethod rule.
	var primaryMethodIndex uint
	var primaryMethods = primarySubsection.GetPrimaryMethods().GetIterator()
	var primaryMethodsSize = uint(primaryMethods.GetSize())
	for primaryMethods.HasNext() {
		primaryMethodIndex++
		var primaryMethod = primaryMethods.GetNext()
		v.processor_.PreprocessPrimaryMethod(
			primaryMethod,
			primaryMethodIndex,
			primaryMethodsSize,
		)
		v.visitPrimaryMethod(primaryMethod)
		v.processor_.PostprocessPrimaryMethod(
			primaryMethod,
			primaryMethodIndex,
			primaryMethodsSize,
		)
	}
}

func (v *visitor_) visitPrimitiveDeclarations(primitiveDeclarations ast.PrimitiveDeclarationsLike) {
	// Visit a single typeSection rule.
	var typeSection = primitiveDeclarations.GetTypeSection()
	v.processor_.PreprocessTypeSection(typeSection)
	v.visitTypeSection(typeSection)
	v.processor_.PostprocessTypeSection(typeSection)

	// Visit slot 1 between references.
	v.processor_.ProcessPrimitiveDeclarationsSlot(1)

	// Visit a single functionalSection rule.
	var functionalSection = primitiveDeclarations.GetFunctionalSection()
	v.processor_.PreprocessFunctionalSection(functionalSection)
	v.visitFunctionalSection(functionalSection)
	v.processor_.PostprocessFunctionalSection(functionalSection)
}

func (v *visitor_) visitResult(result ast.ResultLike) {
	// Visit the possible result types.
	switch actual := result.GetAny().(type) {
	case ast.NoneLike:
		v.processor_.PreprocessNone(actual)
		v.visitNone(actual)
		v.processor_.PostprocessNone(actual)
	case ast.AbstractionLike:
		v.processor_.PreprocessAbstraction(actual)
		v.visitAbstraction(actual)
		v.processor_.PostprocessAbstraction(actual)
	case ast.MultivalueLike:
		v.processor_.PreprocessMultivalue(actual)
		v.visitMultivalue(actual)
		v.processor_.PostprocessMultivalue(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitSetterMethod(setterMethod ast.SetterMethodLike) {
	// Visit a single name token.
	var name = setterMethod.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessSetterMethodSlot(1)

	// Visit a single parameter rule.
	var parameter = setterMethod.GetParameter()
	if uti.IsDefined(parameter) {
		v.processor_.PreprocessParameter(parameter, 1, 1)
		v.visitParameter(parameter)
		v.processor_.PostprocessParameter(parameter, 1, 1)
	}
}

func (v *visitor_) visitSuffix(suffix ast.SuffixLike) {
	// Visit a single name token.
	var name = suffix.GetName()
	v.processor_.ProcessName(name)
}

func (v *visitor_) visitTypeDeclaration(typeDeclaration ast.TypeDeclarationLike) {
	// Visit a single declaration rule.
	var declaration = typeDeclaration.GetDeclaration()
	v.processor_.PreprocessDeclaration(declaration)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(declaration)

	// Visit slot 1 between references.
	v.processor_.ProcessTypeDeclarationSlot(1)

	// Visit a single abstraction rule.
	var abstraction = typeDeclaration.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)

	// Visit slot 2 between references.
	v.processor_.ProcessTypeDeclarationSlot(2)

	// Visit an optional enumeration rule.
	var optionalEnumeration = typeDeclaration.GetOptionalEnumeration()
	if uti.IsDefined(optionalEnumeration) {
		v.processor_.PreprocessEnumeration(optionalEnumeration)
		v.visitEnumeration(optionalEnumeration)
		v.processor_.PostprocessEnumeration(optionalEnumeration)
	}
}

func (v *visitor_) visitTypeSection(typeSection ast.TypeSectionLike) {
	// Visit each typeDeclaration rule.
	var typeDeclarationIndex uint
	var typeDeclarations = typeSection.GetTypeDeclarations().GetIterator()
	var typeDeclarationsSize = uint(typeDeclarations.GetSize())
	for typeDeclarations.HasNext() {
		typeDeclarationIndex++
		var typeDeclaration = typeDeclarations.GetNext()
		v.processor_.PreprocessTypeDeclaration(
			typeDeclaration,
			typeDeclarationIndex,
			typeDeclarationsSize,
		)
		v.visitTypeDeclaration(typeDeclaration)
		v.processor_.PostprocessTypeDeclaration(
			typeDeclaration,
			typeDeclarationIndex,
			typeDeclarationsSize,
		)
	}
}

func (v *visitor_) visitValue(value ast.ValueLike) {
	// Visit a single name token.
	var name = value.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessValueSlot(1)

	// Visit a single abstraction rule.
	var abstraction = value.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
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

func visitorReference() *visitorClass_ {
	return visitorReference_
}

var visitorReference_ = &visitorClass_{
	// Initialize the class constants.
}
