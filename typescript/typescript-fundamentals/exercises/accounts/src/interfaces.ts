export interface IUser {
  email: string;
  password: string;
  isActive?: boolean;
}

export interface IAdmin extends IUser {
  adminSince?: Date;
}
