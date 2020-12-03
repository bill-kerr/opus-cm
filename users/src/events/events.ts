import { User } from '../models/user';

export interface Event {
  subject: Subject;
  data: any;
}

export interface AdminStatusChange {
  id: string;
  admin: boolean;
}

export enum Subject {
  UserCreated = 'user:created',
  UserUpdated = 'user:updated',
  AdminStatusChanged = 'user:admin_status_changed',
}

export interface UserCreatedEvent {
  subject: Subject.UserCreated;
  data: User;
}

export interface AdminStatusChangedEvent {
  subject: Subject.AdminStatusChanged;
  data: AdminStatusChange;
}
