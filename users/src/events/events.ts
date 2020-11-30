import { Role } from '../models/role';
import { User } from '../models/user';

export interface Event {
  subject: Subject;
  data: any;
}

export interface RoleChange {
  id: string;
  role: Role;
}

export enum Subject {
  UserCreated = 'user:created',
  UserUpdated = 'user:updated',
  UserRoleChanged = 'user:role_changed',
}

export interface UserCreatedEvent {
  subject: Subject.UserCreated;
  data: User;
}

export interface UserRoleChangedEvent {
  subject: Subject.UserRoleChanged;
  data: RoleChange;
}
