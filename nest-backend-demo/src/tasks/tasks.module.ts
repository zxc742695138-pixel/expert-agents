import { Module } from '@nestjs/common';
import { TasksController } from './tasks.controller';
import { TasksService } from './tasks.service';
import { InMemoryTaskRepository } from './repository/task.repository';
import { TASK_REPOSITORY } from './tasks.constants';

@Module({
  controllers: [TasksController],
  providers: [
    TasksService,
    { provide: TASK_REPOSITORY, useClass: InMemoryTaskRepository },
  ],
})
export class TasksModule {}
