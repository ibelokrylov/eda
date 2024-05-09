import { FastifyReply, FastifyRequest, HookHandlerDoneFunction } from 'fastify';
import { generateErrorHelper } from '../helpers/generate-error.helper';
import { ErrorTypeEnum } from '../enums/error-type.enum';

export const authHook = (
  req: FastifyRequest,
  reply: FastifyReply,
  done: HookHandlerDoneFunction,
) => {
  if (!req.session.data?.user_id) {
    generateErrorHelper(ErrorTypeEnum.UNAUTHORIZED, reply);
    return;
  }
  done();
};
