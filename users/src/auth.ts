import { NextFunction, Request, Response } from 'express';
import firebase from 'firebase-admin';
import { v4 } from 'uuid';
import {
  BadRequestError,
  InsufficientPermissionsError,
  InternalServerError,
  NotFoundError,
  UnauthorizedError,
} from './errors/errors';
import { Role } from './models/role';
import { User } from './models/user';

interface FirebaseError {
  errorInfo: { code: string; message: string };
  codePrefix: string;
}

interface UserClaims {
  role: Role;
}

export function initializeFirebase() {
  firebase.initializeApp({
    credential: firebase.credential.applicationDefault(),
  });
  console.log('Users Service connected to Firebase.');
}

export async function createUser(email: string, password: string): Promise<User> {
  try {
    const userRecord = await firebase.auth().createUser({ email, password, uid: v4() });
    await setRole(userRecord.uid, Role.PROJECT_USER);
    return mapUser(userRecord || null);
  } catch (error) {
    handleAuthError(error);
    return mapUser(null);
  }
}

export async function setRole(userId: string, role: Role) {
  try {
    await firebase.auth().setCustomUserClaims(userId, { role });
  } catch (error) {
    handleAuthError(error, userId);
  }
}

export async function getClaims(userId: string) {
  try {
    const user = await firebase.auth().getUser(userId);
    return mapClaims(user.customClaims);
  } catch (error) {
    handleAuthError(error, userId);
    return mapClaims({});
  }
}

export async function requireAuth(req: Request, _res: Response, next: NextFunction) {
  const authHeader = req.headers.authorization;
  if (!authHeader) {
    throw new UnauthorizedError('The Authorization header must be set.');
  }

  const [bearer, authToken] = authHeader.split(' ');
  if (bearer !== 'Bearer') {
    throw new UnauthorizedError(
      "The Authorization header must be formatted as 'Bearer <token>' where <token> is a valid auth key."
    );
  }

  let token: firebase.auth.DecodedIdToken;
  try {
    token = await firebase.auth().verifyIdToken(authToken);
  } catch (err) {
    throw new UnauthorizedError('The provided authentication token is not valid.');
  }

  if (!token) {
    throw new UnauthorizedError('You are not authorized to access this resource.');
  }

  req.userId = token.uid;
  req.userRole = (await getClaims(token.uid)).role;
  next();
}

export async function requireAdmin(req: Request, _res: Response, next: NextFunction) {
  if (req.userRole !== Role.SYS_ADMIN) {
    throw new InsufficientPermissionsError();
  }
  next();
}

async function mapUser(userRecord: firebase.auth.UserRecord | undefined | null): Promise<User> {
  const user = new User();
  user.id = userRecord?.uid || '';
  user.email = userRecord?.email || '';
  user.role = userRecord ? mapClaims(userRecord.customClaims).role : Role.PROJECT_USER;
  return user;
}

function mapClaims(claims: { [key: string]: any } | undefined): UserClaims {
  if (!claims || !claims.role) {
    return { role: Role.PROJECT_USER }; // Or whatever the default should be.
  }
  const role = claims.role as Role;
  return { role };
}

function handleAuthError(error: FirebaseError, userId?: string) {
  switch (error.errorInfo.code) {
    case 'auth/email-already-exists':
      throw new BadRequestError('A user with that email already exists.');
    case 'auth/invalid-email':
      throw new BadRequestError('Invalid email address.');
    case 'auth/user-not-found':
      throw new NotFoundError(`A user with an id of ${userId} does not exist.`);
    default:
      throw new InternalServerError();
  }
}
