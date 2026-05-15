
import { Context } from './Context'


class NhlApiDocumentationError extends Error {

  isNhlApiDocumentationError = true

  sdk = 'NhlApiDocumentation'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  NhlApiDocumentationError
}

