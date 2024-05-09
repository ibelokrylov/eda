import { Static, Type } from '@sinclair/typebox';

export const CreateCategoryDto = Type.Object({
  name: Type.String({ minLength: 3 }),
  description: Type.Optional(Type.String({ minLength: 3 })),
});

export type CreateCategoryDtoType = Static<typeof CreateCategoryDto>;
