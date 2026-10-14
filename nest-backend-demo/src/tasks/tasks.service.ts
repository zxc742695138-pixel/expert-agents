import { Inject, Injectable, NotFoundException } from '@nestjs/common';
import { Task } from './entities/task.entity';
import { CreateTaskDto } from './dto/create-task.dto';
import { UpdateTaskDto } from './dto/update-task.dto';
import { TaskRepository } from './repository/task.repository';
import { TASK_REPOSITORY } from './tasks.constants';

// Service chứa business logic, không biết gì về HTTP (không import @nestjs/common's Req/Res).
@Injectable()
export class TasksService {
  constructor(
    @Inject(TASK_REPOSITORY) private readonly repository: TaskRepository,
  ) {}

  create(dto: CreateTaskDto): Task {
    return this.repository.create(dto.title);
  }

  findAll(): Task[] {
    return this.repository.findAll();
  }

  findOne(id: number): Task {
    const task = this.repository.findById(id);
    if (!task) {
      throw new NotFoundException('task not found');
    }
    return task;
  }

  update(id: number, dto: UpdateTaskDto): Task {
    const task = this.repository.update(id, dto);
    if (!task) {
      throw new NotFoundException('task not found');
    }
    return task;
  }

  remove(id: number): void {
    const deleted = this.repository.delete(id);
    if (!deleted) {
      throw new NotFoundException('task not found');
    }
  }
}
