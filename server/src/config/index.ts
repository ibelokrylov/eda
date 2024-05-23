import 'dotenv/config';
import { generateRandomStringHelper } from '../lib/helpers/generated-random-string.helper';

export const config = {
  app: {
    port: Number(process.env.FASTIFY_PORT) || 3006,
    address: process.env.FASTIFY_ADDRESS || '::',
  },
  prisma: {
    client: {
      url: `postgresql://${process.env.POSTGRES_USER}:${process.env.POSTGRES_PASSWORD}@${process.env.POSTGRES_DB_HOST}:${process.env.POSTGRES_DB_PORT}/${process.env.POSTGRES_DB}?schema=public`,
    },
  },
  session: {
    secret: process.env.SESSION_SECRET ?? generateRandomStringHelper(32),
    salt: process.env.SESSION_SALT ?? generateRandomStringHelper(16),
    ttl: Number(process.env.SESSION_TTL) || 86400,
  },
  redis: {
    url: process.env.REDIS_URL,
    password: process.env.REDIS_PASSWORD,
    port: Number(process.env.REDIS_PORT) || 6379,
    username: process.env.REDIS_USERNAME,
  },
};
