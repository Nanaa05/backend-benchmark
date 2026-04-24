import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule, {
    logger: false,
  });
  await app.listen(8080, '127.0.0.1');
  console.log('Listening on http://127.0.0.1:8080');
}

void bootstrap();
