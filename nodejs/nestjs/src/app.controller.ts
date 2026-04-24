import {
  Controller,
  Get,
  Post,
  Body,
  Req,
  Res,
  HttpStatus,
} from '@nestjs/common';
import type { Request, Response } from 'express';

@Controller()
export class AppController {
  /* ---------- /ping ---------- */
  @Get('/ping')
  ping(@Res() res: Response) {
    res.type('text/plain').status(200).send('ok');
  }

  /* ---------- /json ---------- */
  @Get('/json')
  json() {
    return {
      status: 'ok',
      service: 'nest',
      timestamp: Math.floor(Date.now() / 1000),
    };
  }

  /* ---------- /echo ---------- */
  @Post('/echo')
  echo(@Req() req: Request, @Body() body: any, @Res() res: Response) {
    if (!req.is('application/json')) {
      return res
        .status(HttpStatus.BAD_REQUEST)
        .json({ error: 'expected application/json' });
    }

    return res.json(body);
  }
}
