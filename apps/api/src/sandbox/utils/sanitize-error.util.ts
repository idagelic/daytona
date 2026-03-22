/*
 * Copyright 2025 Daytona Platforms Inc.
 * SPDX-License-Identifier: AGPL-3.0
 */

const RECOVERABLE_ERROR_CODES = ['ECONNREFUSED', 'ECONNRESET', 'ETIMEDOUT', 'EHOSTUNREACH', 'ENETUNREACH', 'EAI_AGAIN']

function isTransientNetworkError(error: any): boolean {
  if (!error || typeof error !== 'object') return false

  const code = error.code || ''
  if (RECOVERABLE_ERROR_CODES.includes(code)) return true

  const message = error.message || ''
  if (RECOVERABLE_ERROR_CODES.some((c) => message.includes(c))) return true

  if (error.name === 'AggregateError' || error.constructor?.name === 'AggregateError') return true

  return false
}

export function sanitizeSandboxError(error: any): { recoverable: boolean; errorReason: string } {
  if (typeof error === 'object' && error !== null && isTransientNetworkError(error)) {
    return { recoverable: true, errorReason: error.message || String(error) }
  }

  if (typeof error === 'string') {
    try {
      const errObj = JSON.parse(error) as { recoverable: boolean; errorReason: string }
      return { recoverable: errObj.recoverable, errorReason: errObj.errorReason }
    } catch {
      if (RECOVERABLE_ERROR_CODES.some((c) => error.includes(c))) {
        return { recoverable: true, errorReason: error }
      }
      return { recoverable: false, errorReason: error }
    }
  } else if (typeof error === 'object' && error !== null && 'recoverable' in error && 'errorReason' in error) {
    return { recoverable: error.recoverable, errorReason: error.errorReason }
  } else if (typeof error === 'object' && error.message) {
    return sanitizeSandboxError(error.message)
  }

  return { recoverable: false, errorReason: String(error) }
}
