/** @type {import('next').NextConfig} */
const nextConfig = {
  transpilePackages: ['@company/ui', '@company/utils', '@company/config', '@company/types'],
  experimental: {
    optimizePackageImports: ['@company/ui'],
  },
};

module.exports = nextConfig;