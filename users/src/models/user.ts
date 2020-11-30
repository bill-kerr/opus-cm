import { Exclude, Expose } from 'class-transformer';
import { IsEmail, Length } from 'class-validator';
import { Role } from './role';

@Exclude()
export class User {
  @Expose({ groups: ['http'] })
  object = 'user';

  @Expose({ groups: ['http', 'event'] })
  id: string;

  @Expose({ groups: ['http', 'event', 'create'] })
  @IsEmail({}, { message: 'The email field must contain a valid email.' })
  email: string;

  @Expose({ groups: ['http', 'event'] })
  role: Role;

  @Expose({ groups: ['create'] })
  @Length(6, undefined, { message: 'Passwords must be at least 6 characters long.' })
  password: string;
}
