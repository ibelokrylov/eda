import { PrismaClient } from '@prisma/client';

export class DefaultServiceClass {
  protected prisma: PrismaClient;
  constructor() {
    this.prisma = new PrismaClient();
  }
}
