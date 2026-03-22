/*
 * Copyright Daytona Platforms Inc.
 * SPDX-License-Identifier: AGPL-3.0
 */

const CONNECTION_ERROR_CODES = ['ECONNREFUSED', 'ECONNRESET', 'ETIMEDOUT', 'EHOSTUNREACH', 'ENETUNREACH', 'EAI_AGAIN']

export class RunnerApiError extends Error {
  constructor(
    message: string,
    public readonly statusCode?: number,
    public readonly code?: string,
  ) {
    super(message)
    this.name = 'RunnerApiError'
  }

  isConnectionError(): boolean {
    return CONNECTION_ERROR_CODES.some((c) => this.code === c || this.message?.includes(c))
  }
}
