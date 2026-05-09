/*
 * Copyright 2025 Daytona Platforms Inc.
 * SPDX-License-Identifier: AGPL-3.0
 */

export const RECOVERY_ERROR_SUBSTRINGS: string[] = [
  'Can not connect to the Docker daemon',
  'ECONNREFUSED',
  'ECONNRESET',
  'ETIMEDOUT',
  'EHOSTUNREACH',
  'ENETUNREACH',
  'AggregateError',
  'no such container',
  'container is not running',
  'network is unreachable',
]
