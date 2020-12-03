import { Exclude, Expose } from 'class-transformer';
import { IsBoolean, IsEmail, Length } from 'class-validator';

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
  @IsBoolean({
    message: 'The admin field must contain a boolean value.',
    groups: ['update'],
  })
  admin: boolean;

  @Expose({ groups: ['create'] })
  @Length(6, undefined, {
    message: `Passwords must be at least 6 characters long.`,
    groups: ['create'],
  })
  password: string;
}
