import { IsEmail, Length } from 'class-validator';
import { Role } from './role';

export class User {
  object = 'user';
  id: string;
  email: string;
}

export interface UserClaims {
  role: Role;
}

export class UserCreate {
  @IsEmail({}, { message: 'The email field must contain a valid email.' })
  email: string;

  @Length(6, undefined, { message: 'Passwords must be at least 6 characters long.' })
  password: string;
}

export class UserRead {
  object: 'user';
  id: string;
  email: string;
}
