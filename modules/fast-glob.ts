export default {
  sync: (patterns: any) => [].concat(patterns),
  escapePath: (path: string) => path,
  generateTasks: (patterns: string[]) => {
    return [
      {
        dynamic: false,
        base: '.',
        negative: [],
        positive: ([] as string[]).concat(patterns),
        patterns: ([] as string[]).concat(patterns),
      },
    ]
  },
}
