import express from 'express';
import { classToPlain } from 'class-transformer';
import { createUser, getClaims, requireAdmin, requireAuth, setRole } from './auth';
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
  return res.sendStatus(200);
});

export { router };
