import { plainToClass } from 'class-transformer';
import { ClassType } from 'class-transformer/ClassTransformer';
import { validate, ValidationError as ClassValidationError } from 'class-validator';
import { NextFunction, Request, Response } from 'express';
import { ValidationError } from './errors/errors';

export function validateBody<T>(targetClass: ClassType<T>) {
  return async (req: Request, _res: Response, next: NextFunction) => {
    const instance = plainToClass(targetClass, req.body);
    const errors = await validate(instance, {
      forbidUnknownValues: true,
    });
    if (errors.length > 0) {
      throw new ValidationError(errors);
    }
    next();
  };
}

export function mapClassValidationErrors(errors: ClassValidationError[]): string[] {
  const messages: string[] = [];
  errors.forEach(error => {
    if (error.constraints) {
      const fieldErrors = Object.values(error.constraints).map(message => message);
      messages.push(...fieldErrors);
    }
  });
  return messages;
}
