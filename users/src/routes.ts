import express from 'express';
import { createUser, getClaims, setRole } from './auth';
import { InternalServerError } from './errors/errors';
import { UserCreatedPublisher, UserRoleChangedPublisher } from './events/publishers';
import { Role } from './models/role';
import { UserCreate } from './models/user';
import { natsWrapper } from './nats-wrapper';
import { validateBody } from './validators';

const router = express.Router();

router.post('/', validateBody(UserCreate), async (req, res) => {
  const user = await createUser(req.body.email, req.body.password);
  if (!user) {
    throw new InternalServerError();
  }
  await new UserCreatedPublisher(natsWrapper.client).publish(user);
  res.status(201).json(user);
});

router.post('/superuser', async (req, res) => {
  const userId: string = req.body.id;

  const claims = await getClaims(userId);
  if (claims.role === Role.SYS_ADMIN) {
    return res.sendStatus(200);
  }

  const success = await setRole(userId, Role.SYS_ADMIN);
  if (!success) {
    throw new InternalServerError();
  }
  await new UserRoleChangedPublisher(natsWrapper.client).publish({
    userId,
    role: Role.SYS_ADMIN,
  });
  res.sendStatus(200);
});

export { router };
