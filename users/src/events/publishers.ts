import { Stan } from 'node-nats-streaming';
import { Event, Subject, UserCreatedEvent, UserRoleChangedEvent } from './events';

export abstract class Publisher<T extends Event> {
  abstract subject: T['subject'];

  constructor(private client: Stan) {}

  publish(data: Record<string, any>): Promise<void> {
    return new Promise((resolve, reject) => {
      this.client.publish(this.subject, JSON.stringify(data), err => {
        if (err) {
          return reject(err);
        }
        console.log('Event published to subject', this.subject);
        resolve();
      });
    });
  }
}

export class UserCreatedPublisher extends Publisher<UserCreatedEvent> {
  subject: Subject.UserCreated = Subject.UserCreated;
}

export class UserRoleChangedPublisher extends Publisher<UserRoleChangedEvent> {
  subject: Subject.UserRoleChanged = Subject.UserRoleChanged;
}
