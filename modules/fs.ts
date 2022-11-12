// @ts-ignore
import preflight from 'tailwindcss/lib/css/preflight.css'

export function readFileSync(path: string) {
  if (path === "css/preflight.css") {
    return preflight
  }
  throw new Error(`modules/fs.readFileSync: unable to read: ${path}`)
}

let i = 0
export function statSync(path: string) {
  return { mtimeMs: ++i }
}

export function existsSync(path: string): boolean {
  return false
}

export const promises = {
  readFile(path: string) {
    return Promise.resolve(readFileSync(path))
  }
}

export default {
  readFileSync,
  statSync,
  promises,
}
