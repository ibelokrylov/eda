import { FastifyInstance } from 'fastify';
import { ErrorTypeEnum } from '../../lib/enums/error-type.enum';
import { generateErrorHelper } from '../../lib/helpers/generate-error.helper';
import { authHook } from '../../lib/hook/auth.hook';
import { isNotAuthHook } from '../../lib/hook/is-not-auth.hook';
import { DefaultResponseTypes } from '../../lib/types/default-response.types';
import { CreateUserDto, CreateUserDtoType } from './dto/create-user.dto';
import { PublicUser } from './user.entity';
import { UserService } from './user.service';

export async function userController(fastify: FastifyInstance) {
  const userService = new UserService();
  fastify.get<{
    Reply: DefaultResponseTypes<PublicUser>;
  }>('/user', {
    preHandler: authHook,
    handler: async (req, reply) => {
      try {
        const user = await userService.getUserProfile({
          id: req.session.data.user_id as string,
        });
        return reply.send({
          data: user,
        });
      } catch (error) {
        return generateErrorHelper(ErrorTypeEnum.DEFAULT, reply);
      }
    },
  });

  fastify.post<{
    Body: CreateUserDtoType;
    Reply: DefaultResponseTypes<PublicUser>;
  }>('/user', {
    schema: {
      body: CreateUserDto,
    },
    preHandler: isNotAuthHook,
    handler: async (req, reply) => {
      try {
        const user_data = await userService.createUser({
          dto: req.body,
        });
        req.session.set('user_id', user_data.id);
        return reply.send({
          data: user_data,
        });
      } catch (error) {
        return generateErrorHelper(ErrorTypeEnum.DEFAULT, reply, error.message);
      }
    },
  });
}
