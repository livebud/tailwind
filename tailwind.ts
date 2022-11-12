import tailwindcss, { Config } from 'tailwindcss'
import autoprefixer from 'autoprefixer'
import postcss from 'postcss'
import cssnano from 'cssnano'

export function process(config: Config): Promise<string> {
  return postcss(
    tailwindcss(config),
    autoprefixer({
      overrideBrowserslist: ["defaults"]
    }),
    cssnano()
  ).process(`
    @tailwind base;
    @tailwind components;
    @tailwind utilities;
  `, {
    from: undefined,
  }).then(result => result.css)
}
