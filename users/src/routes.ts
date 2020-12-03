import express from 'express';
import { classToPlain } from 'class-transformer';
import { createUser, requireAdmin, requireAuth, setRole } from './auth';
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

router.put('/:id', requireAuth, requireAdmin, validateBody(User, ['update']), async (req, res) => {
  const userId: string = req.params.id;
  const newRole: Role = req.body.role;
  await setRole(userId, newRole);

  await new UserRoleChangedPublisher(natsWrapper.client).publish({
    id: userId,
    role: Role.SYS_ADMIN,
  });
  return res.sendStatus(200);
});

router.post('/superuser', requireAuth, async (req, res) => {
  await setRole(req.userId, Role.SYS_ADMIN);
  res.sendStatus(200);
});

export { router };
