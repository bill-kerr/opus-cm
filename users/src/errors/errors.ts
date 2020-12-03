import { ValidationError as ClassValidationError } from 'class-validator';
import { mapClassValidationErrors } from '../validators';
import { BaseError, ErrorDetail } from './types';

export class InternalServerError extends BaseError {
  statusCode = 500;
  name = 'Internal Server Error';

  constructor(public details = 'An unknown error occurred.') {
    super(details);
    Object.setPrototypeOf(this, InternalServerError.prototype);
  }
}

export class ValidationError extends BaseError {
  statusCode = 400;
  name = 'Bad Request Error';
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
    this.name = 'Validation Error';
    this.details = messages[0];
  }

  private setMultipleErrors(messages: string[]) {
    this.details = messages.map(message => ({
      object: 'error-detail',
      name: 'Validation Error',
      details: message,
    }));
  }
}

export class BadRequestError extends BaseError {
  statusCode = 400;
  name = 'Bad Request Error';

  constructor(public details: string) {
    super(details);
    Object.setPrototypeOf(this, BadRequestError.prototype);
  }
}

export class UnauthorizedError extends BaseError {
  statusCode = 401;
  name = 'Unauthorized Error';

  constructor(public details: string) {
    super(details);
    Object.setPrototypeOf(this, UnauthorizedError.prototype);
  }
}

export class InsufficientPermissionsError extends BaseError {
  statusCode = 403;
  name = 'Insufficient Permissions Error';

  constructor(
    public details: string = 'You do not have the requisite permissions to perform this operation.'
  ) {
    super(details);
    Object.setPrototypeOf(this, InsufficientPermissionsError.prototype);
  }
}

export class NotFoundError extends BaseError {
  statusCode = 404;
  name = 'Not Found Error';

  constructor(public details: string = 'The requested resource was not found.') {
    super(details);
    Object.setPrototypeOf(this, NotFoundError.prototype);
  }
}
