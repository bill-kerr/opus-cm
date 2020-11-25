export abstract class BaseError extends Error {
  public abstract statusCode: number;
  public abstract details: string | ErrorDetail[];

  constructor(message: string) {
    super(message);
    Object.setPrototypeOf(this, BaseError.prototype);
  }
}

export interface ErrorResponse {
  object: 'error';
  name: string;
  statusCode: number;
  details: ErrorDetail[] | string;
}

export interface ErrorDetail {
  object: 'error-detail';
  name: string;
  details: string;
}
