import { Exclude, Expose, Transform } from 'class-transformer';
import { IsEmail, IsIn, Length } from 'class-validator';
import { Role } from './role';

@Exclude()
export class User {
  @Expose({ groups: ['http'] })
  object = 'user';

  @Expose({ groups: ['http', 'event'] })
  id: string;

  @Expose({ groups: ['http', 'event', 'create'] })
  @IsEmail({}, { message: 'The email field must contain a valid email.', groups: ['create'] })
  email: string;

  @Expose({ groups: ['http', 'event', 'update'] })
  @IsIn(Object.values(Role).map(role => role.toString()), {
    message: "The role field must contain one of 'SYS_ADMIN', 'PRJ_ADMIN', or 'PRJ_USER'",
    groups: ['update'],
  })
  @Transform((val: string) => val.toUpperCase(), { toClassOnly: true })
  role: Role;

  @Expose({ groups: ['create'] })
  @Length(6, undefined, {
    message: `Passwords must be at least 6 characters long.`,
    groups: ['create'],
  })
  password: string;
}
