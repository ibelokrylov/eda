import { Static, Type } from '@sinclair/typebox';

export const CreateProductDto = Type.Object({
  name: Type.String({ minLength: 3 }),
  categoriesProductId: Type.Optional(Type.String()),
  protein: Type.Optional(Type.Number({ default: 0 })),
  fat: Type.Optional(Type.Number({ default: 0 })),
  carbohydrate: Type.Optional(Type.Number({ default: 0 })),
  kcal: Type.Optional(Type.Number({ default: 0 })),
  description: Type.Optional(Type.String({ minLength: 3 })),
});

export type CreateCategoryDtoType = Static<typeof CreateProductDto>;
