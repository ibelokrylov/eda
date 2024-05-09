import { createHmac } from 'node:crypto';

export const hashPasswordHelper = async (dto: {
  password: string;
  hash: string;
}): Promise<string> => {
  return createHmac('sha256', dto.hash).update(dto.password).digest('hex');
};
