export abstract class BaseError extends Error {
  abstract statusCode: number;

  constructor(message: string) {
    super(message);
    Object.setPrototypeOf(this, BaseError.prototype);
  }
}

export interface ErrorResponse {
  object: 'error';
  name: string;
  statusCode: number;
  details: ErrorDetail | string;
}

export interface ErrorDetail {
  object: 'error-detail';
  name: string;
  details: string;
}
