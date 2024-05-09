import { Static, Type } from '@sinclair/typebox';

export const AuthDto = Type.Object({
  username: Type.String({ minLength: 3 }),
  password: Type.String({ minLength: 8 }),
});

export type AuthDtoType = Static<typeof AuthDto>;
