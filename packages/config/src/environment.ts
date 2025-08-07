// Environment types
export type Environment = 'development' | 'staging' | 'production' | 'test';

// Environment detection
export const getEnvironment = (): Environment => {
  if (typeof process !== 'undefined' && process.env.NODE_ENV) {
    return process.env.NODE_ENV as Environment;
  }
  return 'development';
};

export const isDevelopment = (): boolean => getEnvironment() === 'development';
export const isProduction = (): boolean => getEnvironment() === 'production';
export const isStaging = (): boolean => getEnvironment() === 'staging';
export const isTest = (): boolean => getEnvironment() === 'test';

// Environment-specific configurations
export const getApiBaseUrl = (): string => {
  const env = getEnvironment();
  
  switch (env) {
    case 'production':
      return process.env.NEXT_PUBLIC_API_URL || 'https://api.company.com';
    case 'staging':
      return process.env.NEXT_PUBLIC_API_URL || 'https://api-staging.company.com';
    case 'test':
      return 'http://localhost:3001';
    default:
      return process.env.NEXT_PUBLIC_API_URL || 'http://localhost:3001';
  }
};

export const getWebSocketUrl = (): string => {
  const env = getEnvironment();
  
  switch (env) {
    case 'production':
      return process.env.NEXT_PUBLIC_WS_URL || 'wss://ws.company.com';
    case 'staging':
      return process.env.NEXT_PUBLIC_WS_URL || 'wss://ws-staging.company.com';
    default:
      return process.env.NEXT_PUBLIC_WS_URL || 'ws://localhost:3001';
  }
};

// Feature flags based on environment
export const getFeatureFlags = () => {
  const env = getEnvironment();
  
  return {
    enableDevTools: env === 'development',
    enableAnalytics: env === 'production' || env === 'staging',
    enableLogging: env !== 'production',
    enableMockData: env === 'development' || env === 'test',
    enableErrorReporting: env === 'production' || env === 'staging',
  };
};

// Environment configuration
export const config = {
  app: {
    name: process.env.NEXT_PUBLIC_APP_NAME || 'Company App',
    version: process.env.NEXT_PUBLIC_APP_VERSION || '1.0.0',
  },
  api: {
    baseUrl: getApiBaseUrl(),
    timeout: parseInt(process.env.NEXT_PUBLIC_API_TIMEOUT || '10000'),
  },
  auth: {
    tokenKey: 'auth_token',
    refreshTokenKey: 'refresh_token',
  },
  features: getFeatureFlags(),
  environment: getEnvironment(),
};

// Runtime environment checks
export const requireEnvVar = (name: string): string => {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set`);
  }
  return value;
};

export const getEnvVar = (name: string, defaultValue?: string): string | undefined => {
  return process.env[name] || defaultValue;
};