import { NextFunction, Request, Response } from 'express';
import { BaseError, ErrorResponse } from './types';

const defaultError: ErrorResponse = {
  object: 'error',
  name: 'Internal server error',
  statusCode: 500,
  details: 'An unknown error occurred',
};

export function errorHandler(err: Error, req: Request, res: Response, _next: NextFunction) {
  if (err instanceof BaseError) {
    const errorResponse: ErrorResponse = {
      object: 'error',
      name: err.name,
      statusCode: err.statusCode,
      details: 'Error details go here',
    };
    return res.status(err.statusCode).json(errorResponse);
  }

  if (err instanceof SyntaxError) {
    const errorResponse: ErrorResponse = {
      object: 'error',
      name: 'Bad Request',
      statusCode: 400,
      details: 'The request contained invalid JSON.',
    };
    return res.status(400).json(errorResponse);
  }

  console.error(err);
  res.status(defaultError.statusCode).json(defaultError);
}
