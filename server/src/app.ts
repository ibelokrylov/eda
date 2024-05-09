import fastify from 'fastify';
import router from './routes';
import prismaPlugin from './lib/plugins/prisma.plugin';
import { TypeBoxTypeProvider } from '@fastify/type-provider-typebox';
import fastifyCookie from '@fastify/cookie';
import fastifySession from '@mgcrea/fastify-session';
import { config } from './config';
import Redis from 'ioredis';
import RedisStore from '@mgcrea/fastify-session-redis-store';

export const app = fastify({
  logger: true,
  ajv: {
    customOptions: {
      coerceTypes: false,
    },
  },
}).withTypeProvider<TypeBoxTypeProvider>();

const redisClient = new Redis({
  host: config.redis.url ?? 'localhost',
  port: config.redis.port ?? 6379,
  password: config.redis.password ?? undefined,
});

app.register(fastifyCookie);
app.register(fastifySession, {
  secret: config.session.secret,
  store: new RedisStore({
    client: redisClient,
    ttl: config.session.ttl,
  }),
  cookie: {
    maxAge: config.session.ttl,
  },
});
// app.addHook('onRequest', async (request, reply) => {
//   request.session.set('')
// })

app.register(router);
app.register(prismaPlugin);
