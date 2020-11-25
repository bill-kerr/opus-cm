import { BaseError } from './types';

export class InternalServerError extends BaseError {
  statusCode = 500;
  name = 'Internal server error';

  constructor() {
    super('An unknown error occurred.');
    Object.setPrototypeOf(this, InternalServerError.prototype);
  }
}
