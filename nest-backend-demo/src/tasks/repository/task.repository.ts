import { Injectable } from '@nestjs/common';
import { Task } from '../entities/task.entity';
import { UpdateTaskDto } from '../dto/update-task.dto';

// Interface tách rời nơi lưu trữ khỏi logic nghiệp vụ.
// Muốn đổi sang Postgres/Mongo chỉ cần viết class khác implement interface này
// và đổi provider trong tasks.module.ts — service/controller không cần sửa.
export interface TaskRepository {
  create(title: string): Task;
  findAll(): Task[];
  findById(id: number): Task | undefined;
  update(id: number, dto: UpdateTaskDto): Task | undefined;
  delete(id: number): boolean;
}

@Injectable()
export class InMemoryTaskRepository implements TaskRepository {
  private tasks = new Map<number, Task>();
  private nextId = 1;

  create(title: string): Task {
    const task: Task = {
      id: this.nextId++,
      title,
      done: false,
      createdAt: new Date(),
    };
    this.tasks.set(task.id, task);
    return task;
  }

  findAll(): Task[] {
    return Array.from(this.tasks.values());
  }

  findById(id: number): Task | undefined {
    return this.tasks.get(id);
  }

  update(id: number, dto: UpdateTaskDto): Task | undefined {
    const task = this.tasks.get(id);
    if (!task) return undefined;

    const updated: Task = {
      ...task,
      title: dto.title ?? task.title,
      done: dto.done ?? task.done,
    };
    this.tasks.set(id, updated);
    return updated;
  }

  delete(id: number): boolean {
    return this.tasks.delete(id);
  }
}
