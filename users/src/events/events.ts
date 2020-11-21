export interface Event {
  subject: Subject;
  data: any;
}

export enum Subject {
  UserCreated = 'user:created',
  UserUpdated = 'user:updated',
}

export interface UserCreatedEvent {
  subject: Subject.UserCreated;
  data: string;
}
