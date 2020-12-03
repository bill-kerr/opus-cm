declare namespace Express {
  export interface Request {
    userId: string;
    userIsAdmin: boolean;
  }
}
