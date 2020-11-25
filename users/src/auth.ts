import firebase from 'firebase-admin';
import { v4 } from 'uuid';
import { Role } from './models/role';
import { User, UserClaims } from './models/user';

export function initializeFirebase() {
  firebase.initializeApp({
    credential: firebase.credential.applicationDefault(),
  });
  console.log('Users Service connected to Firebase.');
}

export async function createUser(email: string, password: string): Promise<User | null> {
  try {
    const userRecord = await firebase.auth().createUser({ email, password, uid: v4() });
    return mapUser(userRecord);
  } catch (error) {
    return null;
  }
}

export async function setRole(userId: string, role: Role): Promise<boolean> {
  try {
    await firebase.auth().setCustomUserClaims(userId, { role });
  } catch (error) {
    console.error(error);
    return false;
  }
  return true;
}

export async function getClaims(userId: string) {
  const user = await firebase.auth().getUser(userId);
  return mapClaims(user.customClaims);
}

function mapUser(userRecord: firebase.auth.UserRecord): User {
  return {
    id: userRecord.uid,
    email: userRecord.email || '',
  };
}

function mapClaims(claims: { [key: string]: any } | undefined): UserClaims {
  if (!claims || !claims.role) {
    return { role: Role.PROJECT_USER }; // Or whatever the default should be.
  }
  const role = claims.role as Role;
  return { role };
}
