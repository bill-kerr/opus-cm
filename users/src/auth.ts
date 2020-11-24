import firebase from 'firebase-admin';

export function initializeFirebase() {
  firebase.initializeApp({
    credential: firebase.credential.applicationDefault(),
  });
  console.log('Users Service connected to Firebase.');
}

export async function createUser(email: string, password: string) {
  const userRecord = await firebase.auth().createUser({ email, password });
  return userRecord;
}
