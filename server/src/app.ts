import fastifyCookie, { FastifyCookieOptions } from '@fastify/cookie';
import cors from '@fastify/cors';
import { TypeBoxTypeProvider } from '@fastify/type-provider-typebox';
import fastifySession from '@mgcrea/fastify-session';
import RedisStore from '@mgcrea/fastify-session-redis-store';
import fastify from 'fastify';
import Redis from 'ioredis';
import { config } from './config';
import prismaPlugin from './lib/plugins/prisma.plugin';
import router from './routes';

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

app.register(fastifyCookie, {
  secret: config.session.secret,
  setOptions: {
    path: '/',
    secure: true,
    httpOnly: true,
    sameSite: 'None',
  },
} as FastifyCookieOptions);
app.register(fastifySession, {
  secret: config.session.secret,
  store: new RedisStore({
    client: redisClient,
    ttl: config.session.ttl,
  }),
  cookie: {
    maxAge: config.session.ttl,
    sameSite: 'none',
    secure: false,
    httpOnly: true,
    path: '/',
  },
});

app.register(cors, {
  origin: ['localhost', '92.118.114.163'],
  credentials: true,
  methods: ['GET', 'POST', 'PUT', 'DELETE', 'OPTIONS'],
});
// app.addHook('onRequest', async (request, reply) => {
//   request.session.set('')
// })

app.register(router);
app.register(prismaPlugin);
