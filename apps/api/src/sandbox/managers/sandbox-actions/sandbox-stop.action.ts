/*
 * Copyright 2025 Daytona Platforms Inc.
 * SPDX-License-Identifier: AGPL-3.0
 */

import { Injectable, Logger } from '@nestjs/common'
import { Sandbox } from '../../entities/sandbox.entity'
import { SandboxState } from '../../enums/sandbox-state.enum'
import { DONT_SYNC_AGAIN, SandboxAction, SyncState, SYNC_AGAIN } from './sandbox.action'
import { BackupState } from '../../enums/backup-state.enum'
import { RunnerState } from '../../enums/runner-state.enum'
import { RunnerService } from '../../services/runner.service'
import { RunnerAdapterFactory } from '../../runner-adapter/runnerAdapter'
import { SandboxRepository } from '../../repositories/sandbox.repository'
import { LockCode, RedisLockProvider } from '../../common/redis-lock.provider'
import { WithSpan } from '../../../common/decorators/otel.decorator'

const CONNECTION_ERROR_CODES = ['ECONNREFUSED', 'ECONNRESET', 'ETIMEDOUT', 'EHOSTUNREACH', 'ENETUNREACH', 'EAI_AGAIN']

@Injectable()
export class SandboxStopAction extends SandboxAction {
  private readonly logger = new Logger(SandboxStopAction.name)

  constructor(
    protected runnerService: RunnerService,
    protected runnerAdapterFactory: RunnerAdapterFactory,
    protected sandboxRepository: SandboxRepository,
    protected redisLockProvider: RedisLockProvider,
  ) {
    super(runnerService, runnerAdapterFactory, sandboxRepository, redisLockProvider)
  }

  @WithSpan()
  async run(sandbox: Sandbox, lockCode: LockCode): Promise<SyncState> {
    const runner = await this.runnerService.findOneOrFail(sandbox.runnerId)
    if (runner.state !== RunnerState.READY) {
      return DONT_SYNC_AGAIN
    }

    const runnerAdapter = await this.runnerAdapterFactory.create(runner)

    try {
      if (sandbox.state === SandboxState.STARTED) {
        await runnerAdapter.stopSandbox(sandbox.id)
        await this.updateSandboxState(sandbox, SandboxState.STOPPING, lockCode)

        return SYNC_AGAIN
      }

      if (sandbox.state !== SandboxState.STOPPING && sandbox.state !== SandboxState.ERROR) {
        return DONT_SYNC_AGAIN
      }

      const sandboxInfo = await runnerAdapter.sandboxInfo(sandbox.id)

      if (sandboxInfo.state === SandboxState.STOPPED) {
        await this.updateSandboxState(
          sandbox,
          SandboxState.STOPPED,
          lockCode,
          undefined,
          undefined,
          undefined,
          BackupState.NONE,
        )
        return DONT_SYNC_AGAIN
      } else if (sandboxInfo.state === SandboxState.ERROR) {
        await this.updateSandboxState(
          sandbox,
          SandboxState.ERROR,
          lockCode,
          undefined,
          'Sandbox is in error state on runner',
        )
        return DONT_SYNC_AGAIN
      }

      return SYNC_AGAIN
    } catch (error) {
      if (error.response?.status === 404 || error.statusCode === 404) {
        this.logger.warn(
          `Runner returned 404 during stop of sandbox ${sandbox.id}, marking as stopped`,
        )
        await this.updateSandboxState(
          sandbox,
          SandboxState.STOPPED,
          lockCode,
          undefined,
          undefined,
          undefined,
          BackupState.NONE,
        )
        return DONT_SYNC_AGAIN
      }

      const isConnectionError = CONNECTION_ERROR_CODES.some(
        (code) => error.code === code || error.message?.includes(code),
      )
      if (isConnectionError && sandbox.state === SandboxState.STOPPING) {
        this.logger.warn(
          `Runner unreachable during stop of sandbox ${sandbox.id} (${error.code || error.message}), will retry on next sync`,
        )
        return DONT_SYNC_AGAIN
      }

      throw error
    }
  }
}
