import { ValidationError as ClassValidationError } from 'class-validator';
import { mapClassValidationErrors } from '../validators';
import { BaseError, ErrorDetail } from './types';

export class InternalServerError extends BaseError {
  statusCode = 500;
  name = 'Internal server error';

  constructor(public details = 'An unknown error occurred.') {
    super(details);
    Object.setPrototypeOf(this, InternalServerError.prototype);
  }
}

export class ValidationError extends BaseError {
  statusCode = 400;
  name = 'Bad request';
  details: string | ErrorDetail[];

  constructor(private errors: ClassValidationError[]) {
    super('Validation errors occurred.');
    Object.setPrototypeOf(this, ValidationError.prototype);

    const messages = mapClassValidationErrors(errors);
    if (messages.length < 2) {
      this.setSingleError(messages);
    } else {
      this.setMultipleErrors(messages);
    }
  }

  private setSingleError(messages: string[]) {
    if (this.errors.length === 0) {
      this.details = 'An unknown validation error occurred.';
      return;
    }
    this.name = 'Validation error';
    this.details = messages[0];
  }

  private setMultipleErrors(messages: string[]) {
    this.details = messages.map(message => ({
      object: 'error-detail',
      name: 'Validation error',
      details: message,
    }));
  }
}

export class BadRequestError extends BaseError {
  statusCode = 400;
  name = 'Bad request error';

  constructor(public details: string) {
    super(details);
    Object.setPrototypeOf(this, BadRequestError.prototype);
  }
}

export class UnauthorizedError extends BaseError {
  statusCode = 401;
  name = 'Unauthorized error';

  constructor(public details: string) {
    super(details);
    Object.setPrototypeOf(this, UnauthorizedError.prototype);
  }
}
