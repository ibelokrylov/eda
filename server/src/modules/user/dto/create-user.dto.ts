import { Static, Type } from '@sinclair/typebox';

export const CreateUserDto = Type.Object({
  username: Type.String({ minLength: 3, format: 'email' }),
  password: Type.String({ minLength: 8 }),
  passwordConfirm: Type.String({ minLength: 8 }),
  firstName: Type.String({ minLength: 2 }),
  lastName: Type.String({ minLength: 2 }),
});

export type CreateUserDtoType = Static<typeof CreateUserDto>;
