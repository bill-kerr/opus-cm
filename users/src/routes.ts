import express from 'express';
import { createUser } from './auth';

const router = express.Router();

router.post('/', async (req, res) => {
  const user = await createUser(req.body.email, req.body.password);
  res.json(user);
});

export { router };
