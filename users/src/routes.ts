import express from 'express';
import { classToClass, classToPlain } from 'class-transformer';
import { createUser, getClaims, requireAuth, setRole } from './auth';
import { InternalServerError } from './errors/errors';
import { UserCreatedPublisher, UserRoleChangedPublisher } from './events/publishers';
import { Role } from './models/role';
import { User } from './models/user';
import { natsWrapper } from './nats-wrapper';
import { validateBody } from './validators';

const router = express.Router();

router.post('/', validateBody(User, ['create']), async (req, res) => {
  const user = await createUser(req.body.email, req.body.password);
  if (!user) {
    throw new InternalServerError();
  }
  await new UserCreatedPublisher(natsWrapper.client).publish(
    classToPlain(user, { groups: ['event'] })
  );
  res.status(201).json(classToPlain(user, { groups: ['http'] }));
});

router.post('/superuser', requireAuth, async (req, res) => {
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
    id: userId,
    role: Role.SYS_ADMIN,
  });
  res.sendStatus(200);
});

export { router };
