import firebase from 'firebase-admin';
import { v4 } from 'uuid';
import { BadRequestError, InternalServerError, ValidationError } from './errors/errors';
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

export async function setRole(userId: string, role: Role): Promise<boolean> {
  try {
    await firebase.auth().setCustomUserClaims(userId, { role });
  } catch (error) {
    handleAuthError(error);
    return false;
  }
  return true;
}

export async function getClaims(userId: string) {
  const user = await firebase.auth().getUser(userId);
  return mapClaims(user.customClaims);
}

async function mapUser(userRecord: firebase.auth.UserRecord | undefined | null): Promise<User> {
  const user = new User();
  user.id = userRecord?.uid || '';
  user.email = userRecord?.email || '';
  user.role = (await getClaims(user.id)).role;
  return user;
}

function mapClaims(claims: { [key: string]: any } | undefined): UserClaims {
  if (!claims || !claims.role) {
    return { role: Role.PROJECT_USER }; // Or whatever the default should be.
  }
  const role = claims.role as Role;
  return { role };
}

function handleAuthError(error: FirebaseError) {
  switch (error.errorInfo.code) {
    case 'auth/email-already-exists':
      throw new BadRequestError('A user with that email already exists.');
    case 'auth/invalid-email':
      throw new BadRequestError('Invalid email address.');
    default:
      throw new InternalServerError();
  }
}
