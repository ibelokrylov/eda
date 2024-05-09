import { FastifyInstance } from 'fastify';
import { authController } from '../modules/auth/auth.controller';
import { userController } from '../modules/user/user.controller';

export default async function router(fastifyInstance: FastifyInstance) {
  fastifyInstance.get('/', {
    handler: async (_, reply) => {
      return reply.send({
        data: 'Worked',
      });
    },
  });
  userController(fastifyInstance);
  authController(fastifyInstance);
}
