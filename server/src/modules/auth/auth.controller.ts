import { FastifyInstance } from 'fastify';
import { ErrorTypeEnum } from '../../lib/enums/error-type.enum';
import { generateErrorHelper } from '../../lib/helpers/generate-error.helper';
import { isNotAuthHook } from '../../lib/hook/is-not-auth.hook';
import { DefaultResponseTypes } from '../../lib/types/default-response.types';
import { PublicUser } from '../user/user.entity';
import { AuthService } from './auth.service';
import { AuthDto, AuthDtoType } from './dto/auth.dto';

export const authController = (fastifyInstance: FastifyInstance) => {
  const authService = new AuthService();

  fastifyInstance.post<{
    Body: AuthDtoType;
    Reply: DefaultResponseTypes<PublicUser>;
  }>('/auth', {
    preHandler: isNotAuthHook,
    handler: async (req, reply) => {
      try {
        const dto = req.body;
        const user_data = await authService.login(dto);
        req.session.set('user_id', user_data.id);
        return reply.send({
          data: user_data,
        });
      } catch (error) {
        return generateErrorHelper(ErrorTypeEnum.DEFAULT, reply, error.message);
      }
    },
    schema: {
      body: AuthDto,
    },
  });
};
