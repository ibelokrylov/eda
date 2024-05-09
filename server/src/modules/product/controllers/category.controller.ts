import { FastifyInstance } from 'fastify';
import { ErrorTypeEnum } from '../../../lib/enums/error-type.enum';
import { generateErrorHelper } from '../../../lib/helpers/generate-error.helper';
import { authHook } from '../../../lib/hook/auth.hook';
import { DefaultResponseTypes } from '../../../lib/types/default-response.types';
import { CreateCategoryDto } from '../dto/create-category.dto';
import { CreateCategoryDtoType } from '../dto/create-product.dto';
import { CategoryProductEntity } from '../entities/category-product.entity';
import { CategoryService } from '../services/category.service';

export const CategoryController = (fastify: FastifyInstance) => {
  const service = new CategoryService();

  fastify.post<{
    Body: CreateCategoryDtoType;
    Reply: DefaultResponseTypes<CategoryProductEntity>;
  }>('/category', {
    preHandler: authHook,
    handler: async (req, reply) => {
      try {
        const category = await service.createCategory(req.body);
        return reply.send({
          data: category,
        });
      } catch (error) {
        console.log('🚀 ~ handler: ~ error:', error);
        return generateErrorHelper(ErrorTypeEnum.DEFAULT, reply);
      }
    },
    schema: {
      body: CreateCategoryDto,
    },
  });
};
