import { Body, Controller, Get, OnModuleInit, Post } from '@nestjs/common';
import { ICreateTodoRequest } from './interfaces/ICreateTodoRequest';
import { Observable } from 'rxjs';
import { ITodo } from './interfaces/ITodo';
import { Client, ClientGrpc, Transport } from '@nestjs/microservices';
import { join } from 'path';

interface TodoService {
  create(data: ICreateTodoRequest): Promise<ITodo[]>;
  findAll({}): Promise<ITodo[]>
}

@Controller('todo')
export class TodoController implements OnModuleInit {
  @Client({
    transport: Transport.GRPC,
    options: {
      url: '0.0.0.0:9000',
      package: 'todo',
      protoPath: join(__dirname, '../../proto/todo.proto'),
    }
  })
  private readonly client: ClientGrpc;

  private todoService: TodoService;

  onModuleInit() {
    this.todoService = this.client.getService<TodoService>('TodoService')
  }

  @Post()
  create(@Body() data: ICreateTodoRequest): Promise<ITodo[]> {
    return this.todoService.create(data);
  }

  @Get()
  findAll(): Promise<ITodo[]> {
    return this.todoService.findAll({});
  }
}
