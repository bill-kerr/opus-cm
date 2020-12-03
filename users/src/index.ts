import express from 'express';
import helmet from 'helmet';
import cors from 'cors';
import 'express-async-errors';
import 'reflect-metadata';
import { json } from 'body-parser';
import { natsWrapper } from './nats-wrapper';
import { initializeFirebase } from './auth';
import { router } from './routes';
import { errorHandler } from './errors/error-handler';

async function start() {
  const app = express();
  app.use(helmet());
  app.use(cors());
  app.use(json());

  initializeFirebase();

  app.use(router);

  app.get('/users', (req, res) => {
    console.log(req.url, req.baseUrl, req.originalUrl);
    res.send({ message: 'this is the users endpoint' });
  });
  app.get('/', (_, res) => res.send({ message: 'this is the root endpoint!' }));

  try {
    await natsWrapper.connect('opuscm', 'users', 'http://nats-srv:4222');
  } catch (error) {
    console.error(error);
  }

  const subscription = natsWrapper.client.subscribe('test');
  subscription.on('message', msg => console.log(msg.getData()));

  app.use(errorHandler);
  app.listen(3000, () => console.log('Users service listening on port 3000.'));
}

start();
