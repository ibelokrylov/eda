import { FastifyRequest, FastifyReply, HookHandlerDoneFunction } from 'fastify';
import { ErrorTypeEnum } from '../enums/error-type.enum';
import { generateErrorHelper } from '../helpers/generate-error.helper';

export const isNotAuthHook = (
  req: FastifyRequest,
  reply: FastifyReply,
  done: HookHandlerDoneFunction,
) => {
  if (req.session.data?.user_id) {
    generateErrorHelper(ErrorTypeEnum.AUTHENTICATED, reply);
    return;
  }
  done();
};
