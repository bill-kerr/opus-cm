import { natsWrapper } from './nats-wrapper';

async function start() {
  try {
    await natsWrapper.connect('opuscm', 'users', 'http://nats-srv:4222');
  } catch (error) {
    console.error(error);
  }
}

start();
