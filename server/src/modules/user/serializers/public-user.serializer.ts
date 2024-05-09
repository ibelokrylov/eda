import { excludeHelper } from '../../../lib/helpers/exclude.helper';
import { UserEntity, PublicUser } from '../user.entity';

export const publicUserSerializer = (user: UserEntity): PublicUser => {
  return excludeHelper(user, ['password', 'hash']);
};
