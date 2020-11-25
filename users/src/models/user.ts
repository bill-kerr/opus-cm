import { IsEmail, Length } from 'class-validator';
import { Role } from './role';

export interface User {
  id: string;
  email: string;
}

export interface UserClaims {
  role: Role;
}

export class UserCreate {
  @IsEmail()
  email: string;

  @Length(6)
  password: string;
}
