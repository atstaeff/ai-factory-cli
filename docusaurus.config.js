// Docusaurus sources live in docs/. GitHub Pages serves the generated build/ artifact.
const config = {
  title: 'AI Factory',
  tagline: 'Process-driven AI work orchestration',
  favicon: 'img/favicon.svg',
  url: 'https://atstaeff.github.io',
  baseUrl: '/ai-factory-cli/',
  organizationName: 'atstaeff',
  projectName: 'ai-factory-cli',
  trailingSlash: false,
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  i18n: { defaultLocale: 'en', locales: ['en'] },
  presets: [
    ['classic', {
      docs: {
        sidebarPath: './sidebars.js',
        editUrl: 'https://github.com/atstaeff/ai-factory-cli/edit/main/',
      },
      blog: false,
      theme: { customCss: './src/css/custom.css' },
    }],
  ],
  themeConfig: {
    navbar: {
      title: 'AI Factory',
      logo: { alt: 'AI Factory', src: 'img/logo.svg' },
      items: [
        { type: 'docSidebar', sidebarId: 'tutorialSidebar', label: 'Documentation', position: 'left' },
        { to: '/docs/getting-started', label: 'Quickstart', position: 'left' },
        { href: 'https://github.com/atstaeff/ai-factory-cli', label: 'GitHub', position: 'right' },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        { title: 'Learn', items: [{ label: 'Getting started', to: '/docs/getting-started' }, { label: 'Examples', to: '/docs/examples' }] },
        { title: 'Project', items: [{ label: 'Source code', href: 'https://github.com/atstaeff/ai-factory-cli' }] },
      ],
      copyright: 'AI Factory — open source work orchestration in progress',
    },
    prism: { additionalLanguages: ['bash', 'yaml', 'go'] },
  },
};
export default config;
